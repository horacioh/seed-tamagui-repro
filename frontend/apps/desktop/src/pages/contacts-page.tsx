import Footer from '@/components/footer'
import {ListItem} from '@/components/list-item'
import {MainWrapper, MainWrapperNoScroll} from '@/components/main-wrapper'
import {useListProfileDocuments} from '@/models/documents'
import {getFileUrl} from '@/utils/account-url'
import {useNavigate} from '@/utils/useNavigate'
import {PlainMessage} from '@bufbuild/protobuf'
import {DocumentListItem, hmId} from '@shm/shared'
import {
  Container,
  List,
  PageHeading,
  Spinner,
  Text,
  UIAvatar,
  YStack,
} from '@shm/ui'
import {useRef} from 'react'
import {useShowTitleObserver} from './app-title'

function ErrorPage({}: {error: any}) {
  // todo, this!
  return (
    <MainWrapper>
      <Container>
        <Text fontFamily="$body" fontSize="$3">
          Error
        </Text>
      </Container>
    </MainWrapper>
  )
}

export default function ContactsPage() {
  const contacts = useListProfileDocuments()
  const navigate = useNavigate('push')
  const ref = useRef(null)
  useShowTitleObserver(ref.current)
  if (contacts.isLoading) {
    return (
      <MainWrapper>
        <Container>
          <Spinner />
        </Container>
      </MainWrapper>
    )
  }
  if (contacts.error) {
    return <ErrorPage error={contacts.error} />
  }
  if (!contacts.data?.length) {
    return (
      <>
        <MainWrapper>
          <Container>
            <YStack gap="$5" paddingVertical="$8">
              <Text fontFamily="$body" fontSize="$3">
                You have no Contacts yet.
              </Text>
            </YStack>
          </Container>
        </MainWrapper>
        <Footer />
      </>
    )
  }
  return (
    <>
      <MainWrapperNoScroll>
        <List
          header={
            <Container>
              <PageHeading ref={ref}>Contacts</PageHeading>
            </Container>
          }
          items={contacts.data!}
          renderItem={({item}: {item: PlainMessage<DocumentListItem>}) => {
            console.log(
              `== ~ ContactsPageasdkjahds kajsdh  ~ item:`,
              getFileUrl(item.metadata.thumbnail),
            )
            return (
              <ListItem
                title={item.metadata.name}
                icon={
                  <UIAvatar
                    url={getFileUrl(item.metadata.thumbnail)}
                    label={item.metadata.name}
                    size={24}
                  />
                }
                onPress={() => {
                  navigate({
                    key: 'document',
                    id: hmId('d', item.account),
                  })
                }}
              />
            )
          }}
        />
      </MainWrapperNoScroll>
      {/* {copyDialogContent}
      {deleteEntity.content} */}
      <Footer />
    </>
  )
}
