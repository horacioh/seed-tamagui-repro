import {hmBlockSchema} from '@/editor/schema'
import {API_FILE_UPLOAD_URL, API_FILE_URL} from '@shm/shared'
import {DOMParser as ProseMirrorDOMParser} from '@tiptap/pm/model'
import rehypeStringify from 'rehype-stringify'
import remarkParse from 'remark-parse'
import remarkRehype from 'remark-rehype'
import {unified} from 'unified'
import {Block, BlockNoteEditor, BlockSchema, nodeToBlock} from '../..'

const fileRegex = /\[([^\]]+)\]\(([^)]*) "size=(\d+)"\)/
const videoRegex = /!\[([^\]]+)\]\(([^)]*) "width=(\d*)"\)/
const mathRegex = /\$\$(.*?)\$\$/

const uploadToIpfs = async (file: File): Promise<string> => {
  if (file.size <= 62914560) {
    const formData = new FormData()
    formData.append('file', file)

    try {
      const response = await fetch(API_FILE_UPLOAD_URL, {
        method: 'POST',
        body: formData,
      })
      if (!response.ok) {
        throw new Error('Failed to upload to IPFS')
      }
      const data = await response.text()
      return data
    } catch (error) {
      console.error('Failed to upload to IPFS:', error)
      throw new Error('Failed to upload to IPFS')
    }
  } else {
    throw new Error('The file size exceeds 60 MB')
  }
}

const isWebUrl = (url: string | undefined) => {
  if (!url) return false
  try {
    new URL(url)
    return true
  } catch (_) {
    return false
  }
}

const readMediaFile = async (filePath: string) => {
  try {
    // @ts-ignore
    const response = await window.fileOpen.readMediaFile(filePath)
    return response
  } catch (error) {
    console.error('Error reading media file:', error)
    return
  }
}

export const processMediaMarkdown = async (
  markdownContent: string,
  directoryPath: string,
) => {
  const filePattern = /\[([^\]]+)\]\\\(([^\s]+\.[^\s]+) "size=(\d+)"\)/g
  const videoPattern = /!\\\[([^\]]+)\]\\\(([^\s]+\.[^\s]+) "width=(\d*)"\)/g
  const imagePattern = /!\[([^\]]*)\]\(([^\s]+\.[^\s]+)(?: "([^"]*)")?\)/g

  const mediaMatches = []
  const mediaRegex = new RegExp(
    `${filePattern.source}|${videoPattern.source}|${imagePattern.source}`,
    'g',
  )
  let match
  while ((match = mediaRegex.exec(markdownContent)) !== null) {
    mediaMatches.push(match)
  }

  for (const match of mediaMatches) {
    let url
    if (match[2]) {
      // File pattern
      url = match[2]
    } else if (match[5]) {
      // Video pattern
      url = match[5]
    } else if (match[8]) {
      // Image pattern
      url = match[8]
    }
    if (url && !isWebUrl(url)) {
      try {
        const filePath = directoryPath + '/' + url
        const fileResponse = await readMediaFile(filePath)
        const fileContent = Uint8Array.from(atob(fileResponse.content), (c) =>
          c.charCodeAt(0),
        )
        const file = new File([fileContent], fileResponse.fileName, {
          type: fileResponse.mimeType,
        })
        const ipfsUrl = await uploadToIpfs(file)
        markdownContent = markdownContent.replace(
          url,
          `${API_FILE_URL}/${ipfsUrl}`,
        )
      } catch (error) {
        console.error(`Error processing file ${url}:`, error)
        markdownContent = markdownContent.replace(url, '')
      }
    }
  }

  return markdownContent
}

export const MarkdownToBlocks = async (
  markdown: string,
  editor: BlockNoteEditor,
) => {
  const blocks: Block<BlockSchema>[] = []
  const organizedBlocks: Block<BlockSchema>[] = []

  const file = await unified()
    .use(remarkParse)
    .use(remarkRehype)
    .use(rehypeStringify)
    .process(markdown)

  const parser = new DOMParser()
  const doc = parser.parseFromString(file.value.toString(), 'text/html')

  const {view} = editor._tiptapEditor
  const {state} = view

  // Get ProseMirror fragment from pasted markdown, previously converted to HTML
  const fragment = ProseMirrorDOMParser.fromSchema(view.state.schema).parse(
    doc.body,
  )
  fragment.firstChild!.content.forEach((node) => {
    if (node.type.name !== 'blockContainer') {
      return false
    }
    blocks.push(nodeToBlock(node, hmBlockSchema))
  })

  // Function to determine heading level
  const getHeadingLevel = (block: Block<BlockSchema>) => {
    if (block.type.startsWith('heading')) {
      return parseInt(block.props.level, 10)
    }
    return 0
  }

  // Stack to track heading levels for hierarchy
  const stack: {level: number; block: Block<BlockSchema>}[] = []

  blocks.forEach((block) => {
    const headingLevel = getHeadingLevel(block)

    if (headingLevel > 0) {
      while (stack.length && stack[stack.length - 1].level >= headingLevel) {
        stack.pop()
      }

      if (stack.length) {
        stack[stack.length - 1].block.children.push(block)
      } else {
        organizedBlocks.push(block)
      }

      stack.push({level: headingLevel, block})
    } else {
      let blockToInsert = block
      if (block.type === 'image' && block.props.src.length == 0) {
        blockToInsert.props = {}
      }
      if (block.content.length > 0) {
        const blockContent =
          block.content[0].type === 'link'
            ? block.content[0].content[0].text
            : // @ts-ignore
              block.content[0].text

        if (blockContent.startsWith('!')) {
          const videoMatch = blockContent.match(videoRegex)
          if (videoMatch) {
            let videoProps = {}
            if (videoMatch[2].startsWith(API_FILE_URL)) {
              videoProps = {
                name: videoMatch[1],
                url: videoMatch[2],
                width: videoMatch[3] || '',
              }
            }
            blockToInsert = {
              id: block.id,
              type: 'video',
              props: videoProps,
              content: [],
              children: [],
            }
          }
        } else if (blockContent.startsWith('[')) {
          const fileMatch = blockContent.match(fileRegex)
          if (fileMatch) {
            let fileProps = {}
            if (fileMatch[2].startsWith(API_FILE_URL)) {
              fileProps = {
                name: fileMatch[1],
                url: fileMatch[2],
                size: fileMatch[3],
              }
            }
            blockToInsert = {
              id: block.id,
              type: 'file',
              props: fileProps,
              content: [],
              children: [],
            }
          }
        } else if (mathRegex.test(blockContent)) {
          const mathMatch = blockContent.match(mathRegex)
          if (mathMatch) {
            const mathContent = mathMatch[1]
            blockToInsert = {
              id: block.id,
              type: 'math',
              content: [
                {
                  text: mathContent,
                  type: 'text',
                  styles: {},
                },
              ],
              children: [],
              props: {
                childrenType: 'group',
              },
            }
          }
        }
      }
      if (stack.length) {
        stack[stack.length - 1].block.children.push(blockToInsert)
      } else {
        organizedBlocks.push(blockToInsert)
      }
    }
  })
  return organizedBlocks
}
