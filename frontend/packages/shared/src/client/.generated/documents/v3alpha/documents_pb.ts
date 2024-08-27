// @generated by protoc-gen-es v1.4.1 with parameter "target=ts,import_extension=none"
// @generated from file documents/v3alpha/documents.proto (package com.seed.documents.v3alpha, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";

/**
 * Request for getting a single document.
 *
 * @generated from message com.seed.documents.v3alpha.GetDocumentRequest
 */
export class GetDocumentRequest extends Message<GetDocumentRequest> {
  /**
   * Required. The ID of the account where the document is located.
   *
   * @generated from field: string account = 1;
   */
  account = "";

  /**
   * Required. Path of the document.
   * Empty string means root document.
   *
   * @generated from field: string path = 2;
   */
  path = "";

  /**
   * Optional. Exact version of the document to retrieve.
   *
   * @generated from field: string version = 3;
   */
  version = "";

  constructor(data?: PartialMessage<GetDocumentRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "com.seed.documents.v3alpha.GetDocumentRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "account", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "path", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "version", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetDocumentRequest {
    return new GetDocumentRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetDocumentRequest {
    return new GetDocumentRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetDocumentRequest {
    return new GetDocumentRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetDocumentRequest | PlainMessage<GetDocumentRequest> | undefined, b: GetDocumentRequest | PlainMessage<GetDocumentRequest> | undefined): boolean {
    return proto3.util.equals(GetDocumentRequest, a, b);
  }
}

/**
 * Request to create a new document change.
 *
 * @generated from message com.seed.documents.v3alpha.CreateDocumentChangeRequest
 */
export class CreateDocumentChangeRequest extends Message<CreateDocumentChangeRequest> {
  /**
   * Required. The ID of the account where the document is located.
   *
   * @generated from field: string account = 1;
   */
  account = "";

  /**
   * Required. Path of the document to create change for.
   * If document doesn't exist it will be created.
   * Empty string means root document.
   *
   * @generated from field: string path = 2;
   */
  path = "";

  /**
   * Required. Version of the document to apply changes to.
   * Can be empty when creating a new document.
   *
   * @generated from field: string base_version = 3;
   */
  baseVersion = "";

  /**
   * Required. Changes to be applied to the document.
   *
   * @generated from field: repeated com.seed.documents.v3alpha.DocumentChange changes = 4;
   */
  changes: DocumentChange[] = [];

  /**
   * Required. Name of the key to use for signing.
   * Use the Daemon API to list and manage keys.
   *
   * @generated from field: string signing_key_name = 5;
   */
  signingKeyName = "";

  /**
   * Optional. ID of the capability that allows signing key to write on behalf of the account
   * for this particular path.
   *
   * @generated from field: string capability = 6;
   */
  capability = "";

  constructor(data?: PartialMessage<CreateDocumentChangeRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "com.seed.documents.v3alpha.CreateDocumentChangeRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "account", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "path", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "base_version", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "changes", kind: "message", T: DocumentChange, repeated: true },
    { no: 5, name: "signing_key_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "capability", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateDocumentChangeRequest {
    return new CreateDocumentChangeRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateDocumentChangeRequest {
    return new CreateDocumentChangeRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateDocumentChangeRequest {
    return new CreateDocumentChangeRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateDocumentChangeRequest | PlainMessage<CreateDocumentChangeRequest> | undefined, b: CreateDocumentChangeRequest | PlainMessage<CreateDocumentChangeRequest> | undefined): boolean {
    return proto3.util.equals(CreateDocumentChangeRequest, a, b);
  }
}

/**
 * @generated from message com.seed.documents.v3alpha.DeleteDocumentRequest
 */
export class DeleteDocumentRequest extends Message<DeleteDocumentRequest> {
  /**
   * Required. ID of the account to delete the document from.
   *
   * @generated from field: string account = 1;
   */
  account = "";

  /**
   * Required. Path of the document to delete.
   *
   * @generated from field: string path = 2;
   */
  path = "";

  constructor(data?: PartialMessage<DeleteDocumentRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "com.seed.documents.v3alpha.DeleteDocumentRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "account", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "path", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteDocumentRequest {
    return new DeleteDocumentRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteDocumentRequest {
    return new DeleteDocumentRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteDocumentRequest {
    return new DeleteDocumentRequest().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteDocumentRequest | PlainMessage<DeleteDocumentRequest> | undefined, b: DeleteDocumentRequest | PlainMessage<DeleteDocumentRequest> | undefined): boolean {
    return proto3.util.equals(DeleteDocumentRequest, a, b);
  }
}

/**
 * Request for listing root documents.
 *
 * @generated from message com.seed.documents.v3alpha.ListRootDocumentsRequest
 */
export class ListRootDocumentsRequest extends Message<ListRootDocumentsRequest> {
  /**
   * Optional. Number of results per page. Default is defined by the server.
   *
   * @generated from field: int32 page_size = 1;
   */
  pageSize = 0;

  /**
   * Optional. Value from next_page_token obtained from a previous response.
   *
   * @generated from field: string page_token = 2;
   */
  pageToken = "";

  constructor(data?: PartialMessage<ListRootDocumentsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "com.seed.documents.v3alpha.ListRootDocumentsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "page_size", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "page_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListRootDocumentsRequest {
    return new ListRootDocumentsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListRootDocumentsRequest {
    return new ListRootDocumentsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListRootDocumentsRequest {
    return new ListRootDocumentsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ListRootDocumentsRequest | PlainMessage<ListRootDocumentsRequest> | undefined, b: ListRootDocumentsRequest | PlainMessage<ListRootDocumentsRequest> | undefined): boolean {
    return proto3.util.equals(ListRootDocumentsRequest, a, b);
  }
}

/**
 * Response for listing root documents.
 *
 * @generated from message com.seed.documents.v3alpha.ListRootDocumentsResponse
 */
export class ListRootDocumentsResponse extends Message<ListRootDocumentsResponse> {
  /**
   * List of root documents.
   *
   * @generated from field: repeated com.seed.documents.v3alpha.DocumentListItem documents = 1;
   */
  documents: DocumentListItem[] = [];

  /**
   * Token for the next page if there're more results.
   *
   * @generated from field: string next_page_token = 2;
   */
  nextPageToken = "";

  constructor(data?: PartialMessage<ListRootDocumentsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "com.seed.documents.v3alpha.ListRootDocumentsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "documents", kind: "message", T: DocumentListItem, repeated: true },
    { no: 2, name: "next_page_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListRootDocumentsResponse {
    return new ListRootDocumentsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListRootDocumentsResponse {
    return new ListRootDocumentsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListRootDocumentsResponse {
    return new ListRootDocumentsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ListRootDocumentsResponse | PlainMessage<ListRootDocumentsResponse> | undefined, b: ListRootDocumentsResponse | PlainMessage<ListRootDocumentsResponse> | undefined): boolean {
    return proto3.util.equals(ListRootDocumentsResponse, a, b);
  }
}

/**
 * Request for listing documents.
 *
 * @generated from message com.seed.documents.v3alpha.ListDocumentsRequest
 */
export class ListDocumentsRequest extends Message<ListDocumentsRequest> {
  /**
   * Required. ID of the account to list documents for.
   *
   * @generated from field: string account = 1;
   */
  account = "";

  /**
   * Optional. Number of results per page. Default is defined by the server.
   *
   * @generated from field: int32 page_size = 2;
   */
  pageSize = 0;

  /**
   * Optional. Value from next_page_token obtained from a previous response.
   *
   * @generated from field: string page_token = 3;
   */
  pageToken = "";

  constructor(data?: PartialMessage<ListDocumentsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "com.seed.documents.v3alpha.ListDocumentsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "account", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "page_size", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "page_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListDocumentsRequest {
    return new ListDocumentsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListDocumentsRequest {
    return new ListDocumentsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListDocumentsRequest {
    return new ListDocumentsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ListDocumentsRequest | PlainMessage<ListDocumentsRequest> | undefined, b: ListDocumentsRequest | PlainMessage<ListDocumentsRequest> | undefined): boolean {
    return proto3.util.equals(ListDocumentsRequest, a, b);
  }
}

/**
 * Response with list of documents.
 *
 * @generated from message com.seed.documents.v3alpha.ListDocumentsResponse
 */
export class ListDocumentsResponse extends Message<ListDocumentsResponse> {
  /**
   * List of documents.
   *
   * @generated from field: repeated com.seed.documents.v3alpha.DocumentListItem documents = 1;
   */
  documents: DocumentListItem[] = [];

  /**
   * Token for the next page if there're more results.
   *
   * @generated from field: string next_page_token = 2;
   */
  nextPageToken = "";

  constructor(data?: PartialMessage<ListDocumentsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "com.seed.documents.v3alpha.ListDocumentsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "documents", kind: "message", T: DocumentListItem, repeated: true },
    { no: 2, name: "next_page_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListDocumentsResponse {
    return new ListDocumentsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListDocumentsResponse {
    return new ListDocumentsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListDocumentsResponse {
    return new ListDocumentsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ListDocumentsResponse | PlainMessage<ListDocumentsResponse> | undefined, b: ListDocumentsResponse | PlainMessage<ListDocumentsResponse> | undefined): boolean {
    return proto3.util.equals(ListDocumentsResponse, a, b);
  }
}

/**
 * Basic data about a document that is returned in list responses.
 * Content is omitted for efficiency reasons.
 *
 * @generated from message com.seed.documents.v3alpha.DocumentListItem
 */
export class DocumentListItem extends Message<DocumentListItem> {
  /**
   * Account to which the document belongs.
   *
   * @generated from field: string account = 1;
   */
  account = "";

  /**
   * Path of the document within the account.
   * Empty string means root document.
   *
   * @generated from field: string path = 2;
   */
  path = "";

  /**
   * User-defined metadata attributes of the document.
   *
   * @generated from field: map<string, string> metadata = 3;
   */
  metadata: { [key: string]: string } = {};

  /**
   * Every author ID who has modified this document's version.
   *
   * @generated from field: repeated string authors = 4;
   */
  authors: string[] = [];

  /**
   * Time when the document was created.
   *
   * @generated from field: google.protobuf.Timestamp create_time = 5;
   */
  createTime?: Timestamp;

  /**
   * Time when the document was updated.
   *
   * @generated from field: google.protobuf.Timestamp update_time = 6;
   */
  updateTime?: Timestamp;

  /**
   * Current version of the document.
   *
   * @generated from field: string version = 8;
   */
  version = "";

  constructor(data?: PartialMessage<DocumentListItem>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "com.seed.documents.v3alpha.DocumentListItem";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "account", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "path", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "metadata", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 9 /* ScalarType.STRING */} },
    { no: 4, name: "authors", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 5, name: "create_time", kind: "message", T: Timestamp },
    { no: 6, name: "update_time", kind: "message", T: Timestamp },
    { no: 8, name: "version", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DocumentListItem {
    return new DocumentListItem().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DocumentListItem {
    return new DocumentListItem().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DocumentListItem {
    return new DocumentListItem().fromJsonString(jsonString, options);
  }

  static equals(a: DocumentListItem | PlainMessage<DocumentListItem> | undefined, b: DocumentListItem | PlainMessage<DocumentListItem> | undefined): boolean {
    return proto3.util.equals(DocumentListItem, a, b);
  }
}

/**
 * Document represents metadata and content of a document.
 *
 * @generated from message com.seed.documents.v3alpha.Document
 */
export class Document extends Message<Document> {
  /**
   * Account to which the document belongs.
   *
   * @generated from field: string account = 1;
   */
  account = "";

  /**
   * Path of the document within the account.
   * Empty string means root document.
   *
   * @generated from field: string path = 2;
   */
  path = "";

  /**
   * Metadata values for a document.
   *
   * @generated from field: map<string, string> metadata = 3;
   */
  metadata: { [key: string]: string } = {};

  /**
   * Output only. Every account ID who has modified the document.
   * Includes the original author as well.
   *
   * @generated from field: repeated string authors = 5;
   */
  authors: string[] = [];

  /**
   * Blocks content of the document.
   *
   * @generated from field: repeated com.seed.documents.v3alpha.BlockNode content = 6;
   */
  content: BlockNode[] = [];

  /**
   * Output only. Time when document was created.
   *
   * @generated from field: google.protobuf.Timestamp create_time = 7;
   */
  createTime?: Timestamp;

  /**
   * Output only. Time when document was updated.
   *
   * @generated from field: google.protobuf.Timestamp update_time = 8;
   */
  updateTime?: Timestamp;

  /**
   * Output only. Current version of the document.
   *
   * @generated from field: string version = 9;
   */
  version = "";

  /**
   * Output only. Previous version of the document,
   * Empty if this is the first version.
   *
   * @generated from field: string previous_version = 10;
   */
  previousVersion = "";

  constructor(data?: PartialMessage<Document>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "com.seed.documents.v3alpha.Document";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "account", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "path", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "metadata", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 9 /* ScalarType.STRING */} },
    { no: 5, name: "authors", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 6, name: "content", kind: "message", T: BlockNode, repeated: true },
    { no: 7, name: "create_time", kind: "message", T: Timestamp },
    { no: 8, name: "update_time", kind: "message", T: Timestamp },
    { no: 9, name: "version", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 10, name: "previous_version", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Document {
    return new Document().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Document {
    return new Document().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Document {
    return new Document().fromJsonString(jsonString, options);
  }

  static equals(a: Document | PlainMessage<Document> | undefined, b: Document | PlainMessage<Document> | undefined): boolean {
    return proto3.util.equals(Document, a, b);
  }
}

/**
 * Content block with children.
 *
 * @generated from message com.seed.documents.v3alpha.BlockNode
 */
export class BlockNode extends Message<BlockNode> {
  /**
   * Content block.
   *
   * @generated from field: com.seed.documents.v3alpha.Block block = 1;
   */
  block?: Block;

  /**
   * Child blocks.
   *
   * @generated from field: repeated com.seed.documents.v3alpha.BlockNode children = 2;
   */
  children: BlockNode[] = [];

  constructor(data?: PartialMessage<BlockNode>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "com.seed.documents.v3alpha.BlockNode";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "block", kind: "message", T: Block },
    { no: 2, name: "children", kind: "message", T: BlockNode, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): BlockNode {
    return new BlockNode().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): BlockNode {
    return new BlockNode().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): BlockNode {
    return new BlockNode().fromJsonString(jsonString, options);
  }

  static equals(a: BlockNode | PlainMessage<BlockNode> | undefined, b: BlockNode | PlainMessage<BlockNode> | undefined): boolean {
    return proto3.util.equals(BlockNode, a, b);
  }
}

/**
 * Content block.
 *
 * @generated from message com.seed.documents.v3alpha.Block
 */
export class Block extends Message<Block> {
  /**
   * Block ID. Must be unique within the document.
   *
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * Type of the block. Specific to the renderer.
   *
   * @generated from field: string type = 2;
   */
  type = "";

  /**
   * Text of the content block.
   *
   * @generated from field: string text = 3;
   */
  text = "";

  /**
   * Optional. The hyperlink to an external resource.
   * Must be a valid URL.
   *
   * @generated from field: string ref = 7;
   */
  ref = "";

  /**
   * Arbitrary attributes of the block.
   *
   * @generated from field: map<string, string> attributes = 4;
   */
  attributes: { [key: string]: string } = {};

  /**
   * Annotation "layers" of the block.
   *
   * @generated from field: repeated com.seed.documents.v3alpha.Annotation annotations = 5;
   */
  annotations: Annotation[] = [];

  /**
   * Output only. Current revision of the block. It's the ID of the last Change that modified this block.
   * Additional information about the Change can be obtained using the Changes service.
   *
   * @generated from field: string revision = 6;
   */
  revision = "";

  constructor(data?: PartialMessage<Block>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "com.seed.documents.v3alpha.Block";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "type", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "text", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "ref", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "attributes", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 9 /* ScalarType.STRING */} },
    { no: 5, name: "annotations", kind: "message", T: Annotation, repeated: true },
    { no: 6, name: "revision", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Block {
    return new Block().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Block {
    return new Block().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Block {
    return new Block().fromJsonString(jsonString, options);
  }

  static equals(a: Block | PlainMessage<Block> | undefined, b: Block | PlainMessage<Block> | undefined): boolean {
    return proto3.util.equals(Block, a, b);
  }
}

/**
 * Conceptual annotation "layer" that is applied to arbitrary spans of block text.
 * An "identity" of the layer should be derived deterministically based on its type
 * attributes. Spans inside the same annotation can't overlap.
 *
 * Spans are stored inside the Annotation in a "columnar" format,
 * i.e. StructureOfArrays instead of ArrayOfStructures. See: https://en.wikipedia.org/wiki/AoS_and_SoA.
 * This is useful to reduce the number of allocations and offers more compact serialization, because
 * protobuf is able to "pack" primitive repeated fields more efficiently.
 *
 * @generated from message com.seed.documents.v3alpha.Annotation
 */
export class Annotation extends Message<Annotation> {
  /**
   * Type of the annotation.
   *
   * @generated from field: string type = 1;
   */
  type = "";

  /**
   * Optional. A hyperlink to an external resource.
   * Must be a valid URL.
   *
   * @generated from field: string ref = 5;
   */
  ref = "";

  /**
   * Arbitrary key-value attributes of the annotation.
   *
   * @generated from field: map<string, string> attributes = 2;
   */
  attributes: { [key: string]: string } = {};

  /**
   * Start offsets of possibly disjoint spans of text for which this annotation is applied.
   * Must be sorted and have the same number of items as `ends` list.
   *
   * @generated from field: repeated int32 starts = 3;
   */
  starts: number[] = [];

  /**
   * End offsets of possibly disjoint spans of text for which this annotation is applied.
   * Must be sorted and have the same number of items as `starts` list.
   *
   * @generated from field: repeated int32 ends = 4;
   */
  ends: number[] = [];

  constructor(data?: PartialMessage<Annotation>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "com.seed.documents.v3alpha.Annotation";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "type", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "ref", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "attributes", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 9 /* ScalarType.STRING */} },
    { no: 3, name: "starts", kind: "scalar", T: 5 /* ScalarType.INT32 */, repeated: true },
    { no: 4, name: "ends", kind: "scalar", T: 5 /* ScalarType.INT32 */, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Annotation {
    return new Annotation().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Annotation {
    return new Annotation().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Annotation {
    return new Annotation().fromJsonString(jsonString, options);
  }

  static equals(a: Annotation | PlainMessage<Annotation> | undefined, b: Annotation | PlainMessage<Annotation> | undefined): boolean {
    return proto3.util.equals(Annotation, a, b);
  }
}

/**
 * Granular document change.
 *
 * @generated from message com.seed.documents.v3alpha.DocumentChange
 */
export class DocumentChange extends Message<DocumentChange> {
  /**
   * @generated from oneof com.seed.documents.v3alpha.DocumentChange.op
   */
  op: {
    /**
     * New metadata to set on the document.
     *
     * @generated from field: com.seed.documents.v3alpha.DocumentChange.SetMetadata set_metadata = 1;
     */
    value: DocumentChange_SetMetadata;
    case: "setMetadata";
  } | {
    /**
     * Move operation that creates/moves a block within the document hierarchy.
     *
     * @generated from field: com.seed.documents.v3alpha.DocumentChange.MoveBlock move_block = 2;
     */
    value: DocumentChange_MoveBlock;
    case: "moveBlock";
  } | {
    /**
     * New block state that replaces an existing block.
     *
     * @generated from field: com.seed.documents.v3alpha.Block replace_block = 3;
     */
    value: Block;
    case: "replaceBlock";
  } | {
    /**
     * ID of a block to delete.
     *
     * @generated from field: string delete_block = 4;
     */
    value: string;
    case: "deleteBlock";
  } | { case: undefined; value?: undefined } = { case: undefined };

  constructor(data?: PartialMessage<DocumentChange>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "com.seed.documents.v3alpha.DocumentChange";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "set_metadata", kind: "message", T: DocumentChange_SetMetadata, oneof: "op" },
    { no: 2, name: "move_block", kind: "message", T: DocumentChange_MoveBlock, oneof: "op" },
    { no: 3, name: "replace_block", kind: "message", T: Block, oneof: "op" },
    { no: 4, name: "delete_block", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "op" },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DocumentChange {
    return new DocumentChange().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DocumentChange {
    return new DocumentChange().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DocumentChange {
    return new DocumentChange().fromJsonString(jsonString, options);
  }

  static equals(a: DocumentChange | PlainMessage<DocumentChange> | undefined, b: DocumentChange | PlainMessage<DocumentChange> | undefined): boolean {
    return proto3.util.equals(DocumentChange, a, b);
  }
}

/**
 * Operation to move an existing block to a different place in the document.
 * Move and Create operations are both expressed with this.
 * Conceptually new blocks are moved out of nowhere into the document.
 *
 * @generated from message com.seed.documents.v3alpha.DocumentChange.MoveBlock
 */
export class DocumentChange_MoveBlock extends Message<DocumentChange_MoveBlock> {
  /**
   * ID of the block to move.
   *
   * @generated from field: string block_id = 1;
   */
  blockId = "";

  /**
   * ID of the new parent for the block being moved.
   *
   * @generated from field: string parent = 2;
   */
  parent = "";

  /**
   * ID of the new left sibling for the block being moved.
   *
   * @generated from field: string left_sibling = 3;
   */
  leftSibling = "";

  constructor(data?: PartialMessage<DocumentChange_MoveBlock>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "com.seed.documents.v3alpha.DocumentChange.MoveBlock";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "block_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "parent", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "left_sibling", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DocumentChange_MoveBlock {
    return new DocumentChange_MoveBlock().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DocumentChange_MoveBlock {
    return new DocumentChange_MoveBlock().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DocumentChange_MoveBlock {
    return new DocumentChange_MoveBlock().fromJsonString(jsonString, options);
  }

  static equals(a: DocumentChange_MoveBlock | PlainMessage<DocumentChange_MoveBlock> | undefined, b: DocumentChange_MoveBlock | PlainMessage<DocumentChange_MoveBlock> | undefined): boolean {
    return proto3.util.equals(DocumentChange_MoveBlock, a, b);
  }
}

/**
 * Operation to replace a metadata field with a new value
 *
 * @generated from message com.seed.documents.v3alpha.DocumentChange.SetMetadata
 */
export class DocumentChange_SetMetadata extends Message<DocumentChange_SetMetadata> {
  /**
   * Metadata key to set.
   *
   * @generated from field: string key = 1;
   */
  key = "";

  /**
   * Metadata value to set.
   *
   * @generated from field: string value = 2;
   */
  value = "";

  constructor(data?: PartialMessage<DocumentChange_SetMetadata>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "com.seed.documents.v3alpha.DocumentChange.SetMetadata";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "key", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "value", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DocumentChange_SetMetadata {
    return new DocumentChange_SetMetadata().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DocumentChange_SetMetadata {
    return new DocumentChange_SetMetadata().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DocumentChange_SetMetadata {
    return new DocumentChange_SetMetadata().fromJsonString(jsonString, options);
  }

  static equals(a: DocumentChange_SetMetadata | PlainMessage<DocumentChange_SetMetadata> | undefined, b: DocumentChange_SetMetadata | PlainMessage<DocumentChange_SetMetadata> | undefined): boolean {
    return proto3.util.equals(DocumentChange_SetMetadata, a, b);
  }
}

