import {
  getFileUrl,
  getMetadataName,
  HMMetadata,
  UnpackedHypermediaId,
  useRouteLink,
} from "@shm/shared";
import {Button} from "@tamagui/button";
import {AlertCircle} from "@tamagui/lucide-icons";
import {YStack} from "@tamagui/stacks";
import {memo} from "react";
import {UIAvatar, UIAvatarProps} from "./avatar";
import {Tooltip} from "./tooltip";

export const Thumbnail = memo(_Thumbnail);

function _Thumbnail({
  id,
  metadata,
  size = 32,
  ...props
}: Omit<UIAvatarProps, "id"> & {
  id: UnpackedHypermediaId;
  metadata?: HMMetadata | null;
  size?: number;
}) {
  if (!id) return null;

  return (
    <UIAvatar
      size={size}
      // id={id.path?.at(-1) || id.uid.slice(2)}
      id={id.id}
      label={metadata?.name}
      url={getFileUrl(metadata?.thumbnail)}
      borderRadius={id.path && id.path.length != 0 ? size / 8 : undefined}
      flexShrink={0}
      flexGrow={0}
      {...props}
    />
  );
}

export function LinkThumbnail({
  id,
  metadata,
  size,
  error,
}: {
  id: UnpackedHypermediaId;
  metadata?: HMMetadata | null;
  size?: number;
  error?: boolean;
}) {
  const linkProps = useRouteLink({key: "document", id});
  let content = (
    <>
      <Thumbnail id={id} size={size} metadata={metadata} />
      <ErrorDot error={error} />
    </>
  );

  return (
    <Tooltip
      content={
        getMetadataName(metadata) ||
        `${id.uid.slice(0, 5)}...${id.uid.slice(-5)}`
      }
    >
      <Button
        className="no-window-drag"
        size="$1"
        backgroundColor="transparent"
        hoverStyle={{backgroundColor: "transparent"}}
        minWidth={20}
        minHeight={20}
        padding={0}
        {...linkProps}
        position="relative"
        height={size}
      >
        {content}
      </Button>
    </Tooltip>
  );
}

export function ErrorDot({error}: {error?: boolean}) {
  return error ? (
    <YStack
      backgroundColor="$red11"
      display="flex"
      position="absolute"
      top={-8}
      left={-8}
      padding={0}
      paddingLeft={-4}
      width={16}
      height={16}
      borderRadius={8}
    >
      <AlertCircle size={16} />
    </YStack>
  ) : null;
}