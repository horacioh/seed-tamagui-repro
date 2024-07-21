package daemon

import (
	"context"
	"seed/backend/core"
	"seed/backend/core/coretest"
	"seed/backend/daemon/api/documents/v2alpha/docmodel"
	daemon "seed/backend/genproto/daemon/v1alpha"
	documents "seed/backend/genproto/documents/v2alpha"
	networking "seed/backend/genproto/networking/v1alpha"
	"seed/backend/mttnet"
	"seed/backend/pkg/must"
	"seed/backend/testutil"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func TestDaemonRegisterKey(t *testing.T) {
	t.Parallel()

	dmn := makeTestApp(t, "alice", makeTestConfig(t), false)
	ctx := context.Background()

	conn, err := grpc.Dial(dmn.GRPCListener.Addr().String(), grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	defer conn.Close()

	dc := daemon.NewDaemonClient(conn)

	seed, err := dc.GenMnemonic(ctx, &daemon.GenMnemonicRequest{})
	require.NoError(t, err)

	reg, err := dc.RegisterKey(ctx, &daemon.RegisterKeyRequest{
		Name:     "main",
		Mnemonic: seed.Mnemonic,
	})
	require.NoError(t, err)
	require.NotNil(t, reg)
	require.NotEqual(t, "", reg.PublicKey, "account ID must be generated after registration")

	_, err = core.DecodePrincipal(reg.PublicKey)
	require.NoError(t, err, "account must have principal encoding")

	me := must.Do2(dmn.Storage.KeyStore().GetKey(ctx, "main"))
	require.Equal(t, me.String(), reg.PublicKey)

	keys, err := dc.ListKeys(ctx, &daemon.ListKeysRequest{})
	require.NoError(t, err)
	require.Len(t, keys.Keys, 1, "there must only be one key")

	{
		seed, err := dc.GenMnemonic(ctx, &daemon.GenMnemonicRequest{})
		require.NoError(t, err)

		reg, err := dc.RegisterKey(ctx, &daemon.RegisterKeyRequest{
			Name:     "secondary",
			Mnemonic: seed.Mnemonic,
		})
		require.NoError(t, err)
		require.NotNil(t, reg)
		require.NotEqual(t, "", reg.PublicKey, "account ID must be generated after registration")

		keys, err := dc.ListKeys(ctx, &daemon.ListKeysRequest{})
		require.NoError(t, err)
		require.Len(t, keys.Keys, 2, "there must only be two keys after registering second key")
	}
}

func TestDaemonUpdateProfile(t *testing.T) {
	t.Parallel()

	dmn := makeTestApp(t, "alice", makeTestConfig(t), true)
	ctx := context.Background()
	alice := coretest.NewTester("alice")

	doc, err := dmn.RPC.DocumentsV2.ChangeProfileDocument(ctx, &documents.ChangeProfileDocumentRequest{
		AccountId: alice.Account.Principal().String(),
		Changes: []*documents.DocumentChange{
			{Op: &documents.DocumentChange_SetMetadata_{
				SetMetadata: &documents.DocumentChange_SetMetadata{Key: "title", Value: "Alice from the Wonderland"},
			}},
			{Op: &documents.DocumentChange_MoveBlock_{
				MoveBlock: &documents.DocumentChange_MoveBlock{BlockId: "b1", Parent: "", LeftSibling: ""},
			}},
			{Op: &documents.DocumentChange_ReplaceBlock{
				ReplaceBlock: &documents.Block{
					Id:   "b1",
					Type: "paragraph",
					Text: "Hello",
				},
			}},
			{Op: &documents.DocumentChange_MoveBlock_{
				MoveBlock: &documents.DocumentChange_MoveBlock{BlockId: "b2", Parent: "b1", LeftSibling: ""},
			}},
			{Op: &documents.DocumentChange_ReplaceBlock{
				ReplaceBlock: &documents.Block{
					Id:   "b2",
					Type: "paragraph",
					Text: "World!",
				},
			}},
		},
	})
	require.NoError(t, err)

	want := &documents.Document{
		Metadata: map[string]string{
			"title": "Alice from the Wonderland",
		},
		Index:   map[string]string{},
		Owner:   alice.Account.Principal().String(),
		Authors: []string{alice.Account.Principal().String()},
		Content: []*documents.BlockNode{
			{
				Block: &documents.Block{
					Id:   "b1",
					Type: "paragraph",
					Text: "Hello",
				},
				Children: []*documents.BlockNode{
					{
						Block: &documents.Block{
							Id:   "b2",
							Type: "paragraph",
							Text: "World!",
						},
					},
				},
			},
		},
	}

	testutil.StructsEqual(want, doc).
		IgnoreFields(documents.Block{}, "Revision").
		IgnoreFields(documents.Document{}, "Id", "CreateTime", "UpdateTime", "Version", "PreviousVersion").
		Compare(t, "profile document must match")
}

func TestDaemonUpdateProfileCompat(t *testing.T) {
	t.Parallel()

	dmn := makeTestApp(t, "alice", makeTestConfig(t), true)
	ctx := context.Background()
	alice := coretest.NewTester("alice")

	doc, err := dmn.RPC.DocumentsV2.CreateDocumentChange(ctx, &documents.CreateDocumentChangeRequest{
		DocumentId: "hm://a/" + alice.Account.Principal().String(),
		Changes: []*documents.DocumentChange{
			{Op: &documents.DocumentChange_SetMetadata_{
				SetMetadata: &documents.DocumentChange_SetMetadata{Key: "title", Value: "Alice from the Wonderland"},
			}},
			{Op: &documents.DocumentChange_MoveBlock_{
				MoveBlock: &documents.DocumentChange_MoveBlock{BlockId: "b1", Parent: "", LeftSibling: ""},
			}},
			{Op: &documents.DocumentChange_ReplaceBlock{
				ReplaceBlock: &documents.Block{
					Id:   "b1",
					Type: "paragraph",
					Text: "Hello",
				},
			}},
			{Op: &documents.DocumentChange_MoveBlock_{
				MoveBlock: &documents.DocumentChange_MoveBlock{BlockId: "b2", Parent: "b1", LeftSibling: ""},
			}},
			{Op: &documents.DocumentChange_ReplaceBlock{
				ReplaceBlock: &documents.Block{
					Id:   "b2",
					Type: "paragraph",
					Text: "World!",
				},
			}},
		},
		SigningKeyName: "main",
	})
	require.NoError(t, err)

	want := &documents.Document{
		Id: "hm://a/" + alice.Account.Principal().String(),
		Metadata: map[string]string{
			"title": "Alice from the Wonderland",
		},
		Index:   map[string]string{},
		Owner:   alice.Account.Principal().String(),
		Authors: []string{alice.Account.Principal().String()},
		Content: []*documents.BlockNode{
			{
				Block: &documents.Block{
					Id:   "b1",
					Type: "paragraph",
					Text: "Hello",
				},
				Children: []*documents.BlockNode{
					{
						Block: &documents.Block{
							Id:   "b2",
							Type: "paragraph",
							Text: "World!",
						},
					},
				},
			},
		},
	}

	testutil.StructsEqual(want, doc).
		IgnoreFields(documents.Block{}, "Revision").
		IgnoreFields(documents.Document{}, "CreateTime", "UpdateTime", "Version", "PreviousVersion").
		Compare(t, "profile document must match")

	// Do another update.
	{
		doc, err := dmn.RPC.DocumentsV2.CreateDocumentChange(ctx, &documents.CreateDocumentChangeRequest{
			DocumentId: "hm://a/" + alice.Account.Principal().String(),
			Changes: []*documents.DocumentChange{
				{Op: &documents.DocumentChange_SetMetadata_{
					SetMetadata: &documents.DocumentChange_SetMetadata{Key: "title", Value: "Just Alice"},
				}},
			},
			SigningKeyName: "main",
		})
		require.NoError(t, err)

		want := &documents.Document{
			Id: "hm://a/" + alice.Account.Principal().String(),
			Metadata: map[string]string{
				"title": "Just Alice",
			},
			Index:   map[string]string{},
			Owner:   alice.Account.Principal().String(),
			Authors: []string{alice.Account.Principal().String()},
			Content: []*documents.BlockNode{
				{
					Block: &documents.Block{
						Id:   "b1",
						Type: "paragraph",
						Text: "Hello",
					},
					Children: []*documents.BlockNode{
						{
							Block: &documents.Block{
								Id:   "b2",
								Type: "paragraph",
								Text: "World!",
							},
						},
					},
				},
			},
		}

		testutil.StructsEqual(want, doc).
			IgnoreFields(documents.Block{}, "Revision").
			IgnoreFields(documents.Document{}, "CreateTime", "UpdateTime", "Version", "PreviousVersion").
			Compare(t, "profile document must match")
	}
}

func TestSubdocuments(t *testing.T) {
	t.Parallel()

	dmn := makeTestApp(t, "alice", makeTestConfig(t), true)
	ctx := context.Background()
	alice := coretest.NewTester("alice")

	parent, err := dmn.RPC.DocumentsV2.CreateDocumentChange(ctx, &documents.CreateDocumentChangeRequest{
		DocumentId: "hm://a/" + alice.Account.Principal().String(),
		Changes: []*documents.DocumentChange{
			{Op: &documents.DocumentChange_SetMetadata_{
				SetMetadata: &documents.DocumentChange_SetMetadata{Key: "title", Value: "Alice from the Wonderland"},
			}},
		},
		SigningKeyName: "main",
	})
	require.NoError(t, err)
	require.NotNil(t, parent)

	subdoc, err := dmn.RPC.DocumentsV2.CreateDocumentChange(ctx, &documents.CreateDocumentChangeRequest{
		DocumentId: "hm://a/" + alice.Account.Principal().String() + "/subdoc",
		Changes: []*documents.DocumentChange{
			{Op: &documents.DocumentChange_SetMetadata_{
				SetMetadata: &documents.DocumentChange_SetMetadata{Key: "title", Value: "Alice's subdoc"},
			}},
			{Op: &documents.DocumentChange_MoveBlock_{
				MoveBlock: &documents.DocumentChange_MoveBlock{BlockId: "b1", Parent: "", LeftSibling: ""},
			}},
			{Op: &documents.DocumentChange_ReplaceBlock{
				ReplaceBlock: &documents.Block{
					Id:   "b1",
					Type: "paragraph",
					Text: "Hello",
				},
			}},
		},
		SigningKeyName: "main",
	})
	require.NoError(t, err)

	want := &documents.Document{
		Id: "hm://a/" + alice.Account.Principal().String() + "/subdoc",
		Metadata: map[string]string{
			"title": "Alice's subdoc",
		},
		Index:   map[string]string{},
		Owner:   alice.Account.Principal().String(),
		Authors: []string{alice.Account.Principal().String()},
		Content: []*documents.BlockNode{
			{
				Block: &documents.Block{
					Id:   "b1",
					Type: "paragraph",
					Text: "Hello",
				},
			},
		},
	}

	testutil.StructsEqual(want, subdoc).
		IgnoreFields(documents.Block{}, "Revision").
		IgnoreFields(documents.Document{}, "CreateTime", "UpdateTime", "Version", "PreviousVersion").
		Compare(t, "subdoc must match")

	// Check parent document has index item.
	{
		parent, err = dmn.RPC.DocumentsV2.GetDocument(ctx, &documents.GetDocumentRequest{
			DocumentId: "hm://a/" + alice.Account.Principal().String(),
		})
		require.NoError(t, err)

		require.NotEmpty(t, parent.Index["subdoc"], "parent must have subdoc index entry")
		_, err = docmodel.Version(parent.Index["subdoc"]).Parse()
		require.NoError(t, err, "subdoc entry must have version as value")
		require.Equal(t, subdoc.Version, parent.Index["subdoc"], "subdoc version must match")
	}

	// Try updating subdocument.
	{
		subdoc, err = dmn.RPC.DocumentsV2.CreateDocumentChange(ctx, &documents.CreateDocumentChangeRequest{
			DocumentId: "hm://a/" + alice.Account.Principal().String() + "/subdoc",
			Changes: []*documents.DocumentChange{
				{Op: &documents.DocumentChange_SetMetadata_{
					SetMetadata: &documents.DocumentChange_SetMetadata{Key: "abstract", Value: "Just some new value to see the update."},
				}},
			},
			SigningKeyName: "main",
		})
		require.NoError(t, err)
		require.NotNil(t, subdoc)

		want := &documents.Document{
			Id: "hm://a/" + alice.Account.Principal().String() + "/subdoc",
			Metadata: map[string]string{
				"title":    "Alice's subdoc",
				"abstract": "Just some new value to see the update.",
			},
			Index:   map[string]string{},
			Owner:   alice.Account.Principal().String(),
			Authors: []string{alice.Account.Principal().String()},
			Content: []*documents.BlockNode{
				{
					Block: &documents.Block{
						Id:   "b1",
						Type: "paragraph",
						Text: "Hello",
					},
				},
			},
		}

		testutil.StructsEqual(want, subdoc).
			IgnoreFields(documents.Block{}, "Revision").
			IgnoreFields(documents.Document{}, "CreateTime", "UpdateTime", "Version", "PreviousVersion").
			Compare(t, "subdoc must match")
	}

	// Remove index entry from parent.
	{
		parent, err := dmn.RPC.DocumentsV2.CreateDocumentChange(ctx, &documents.CreateDocumentChangeRequest{
			DocumentId: "hm://a/" + alice.Account.Principal().String(),
			Changes: []*documents.DocumentChange{
				{Op: &documents.DocumentChange_SetIndex_{SetIndex: &documents.DocumentChange_SetIndex{
					Key:   "subdoc",
					Value: "",
				}}},
			},
			SigningKeyName: "main",
		})
		require.NoError(t, err)

		if _, ok := parent.Index["subdoc"]; ok {
			t.Fatal("parent must not have subdoc index entry after removal")
		}

		// Getting subdocument after removal should fail.
		_, err = dmn.RPC.DocumentsV2.GetDocument(ctx, &documents.GetDocumentRequest{
			DocumentId: "hm://a/" + alice.Account.Principal().String() + "/subdoc",
		})
		require.Error(t, err)
		require.Equal(t, codes.NotFound, status.Code(err), "subdoc must be not found after removal")
	}
}

func TestSyncingProfiles(t *testing.T) {
	t.Parallel()

	alice := makeTestApp(t, "alice", makeTestConfig(t), true)
	ctx := context.Background()
	aliceIdentity := coretest.NewTester("alice")
	bob := makeTestApp(t, "bob", makeTestConfig(t), true)
	bobIdentity := coretest.NewTester("bob")
	doc, err := alice.RPC.DocumentsV2.ChangeProfileDocument(ctx, &documents.ChangeProfileDocumentRequest{
		AccountId: aliceIdentity.Account.Principal().String(),
		Changes: []*documents.DocumentChange{
			{Op: &documents.DocumentChange_SetMetadata_{
				SetMetadata: &documents.DocumentChange_SetMetadata{Key: "title", Value: "Alice from the Wonderland"},
			}},
			{Op: &documents.DocumentChange_MoveBlock_{
				MoveBlock: &documents.DocumentChange_MoveBlock{BlockId: "b1", Parent: "", LeftSibling: ""},
			}},
			{Op: &documents.DocumentChange_ReplaceBlock{
				ReplaceBlock: &documents.Block{
					Id:   "b1",
					Type: "paragraph",
					Text: "Hello",
				},
			}},
			{Op: &documents.DocumentChange_MoveBlock_{
				MoveBlock: &documents.DocumentChange_MoveBlock{BlockId: "b2", Parent: "b1", LeftSibling: ""},
			}},
			{Op: &documents.DocumentChange_ReplaceBlock{
				ReplaceBlock: &documents.Block{
					Id:   "b2",
					Type: "paragraph",
					Text: "World!",
				},
			}},
		},
	})
	require.NoError(t, err)

	_, err = alice.RPC.Networking.Connect(ctx, &networking.ConnectRequest{
		Addrs: mttnet.AddrInfoToStrings(bob.Net.AddrInfo()),
	})
	require.NoError(t, err)

	// _, err = bob.RPC.DocumentsV2.GetProfileDocument(ctx, &documents.GetProfileDocumentRequest{
	//	AccountId: aliceIdentity.Account.Principal().String(),
	// })
	// require.Error(t, err)
	// Since bob implements a syncback policy triggered when Alice connected to him, we don't need
	// to force any syncing just wait for bob to instantly syncs content right after connection.
	//_, err = bob.RPC.Daemon.ForceSync(ctx, &daemon.ForceSyncRequest{})
	//require.NoError(t, err)
	time.Sleep(time.Millisecond * 100)
	doc2, err := bob.RPC.DocumentsV2.GetProfileDocument(ctx, &documents.GetProfileDocumentRequest{
		AccountId: aliceIdentity.Account.Principal().String(),
	})
	require.NoError(t, err)
	require.Equal(t, doc.Content, doc2.Content)

	bobsProfile, err := bob.RPC.DocumentsV2.ChangeProfileDocument(ctx, &documents.ChangeProfileDocumentRequest{
		AccountId: bobIdentity.Account.Principal().String(),
		Changes: []*documents.DocumentChange{
			{Op: &documents.DocumentChange_SetMetadata_{
				SetMetadata: &documents.DocumentChange_SetMetadata{Key: "title", Value: "Bob's land"},
			}},
		},
	})
	require.NoError(t, err)
	docs, err := bob.RPC.DocumentsV2.ListProfileDocuments(ctx, &documents.ListProfileDocumentsRequest{})
	require.NoError(t, err)
	require.Len(t, docs.Documents, 2)
	docs, err = bob.RPC.DocumentsV2.ListProfileDocuments(ctx, &documents.ListProfileDocumentsRequest{
		PageSize:  1,
		PageToken: "",
	})
	require.NoError(t, err)
	require.Len(t, docs.Documents, 1)
	require.Equal(t, bobsProfile.Content, docs.Documents[0].Content)
	docs, err = bob.RPC.DocumentsV2.ListProfileDocuments(ctx, &documents.ListProfileDocumentsRequest{
		PageSize:  1,
		PageToken: docs.NextPageToken,
	})
	require.NoError(t, err)
	require.Len(t, docs.Documents, 1)
	require.Equal(t, doc.Content, docs.Documents[0].Content)
}
