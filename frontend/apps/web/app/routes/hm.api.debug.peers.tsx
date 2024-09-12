import {toPlainMessage} from "@bufbuild/protobuf";
import type {LoaderFunction} from "@remix-run/node";
import {json} from "@remix-run/node";
import {queryClient} from "~/client";

export const loader: LoaderFunction = async () => {
  const peers = await queryClient.networking.listPeers({});
  return json({
    addrs: toPlainMessage(peers),
  });
};