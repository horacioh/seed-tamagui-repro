package syncing

import (
	"context"
	"seed/backend/config"
	"seed/backend/core"
	"seed/backend/index"
	"seed/backend/logging"
	"seed/backend/mttnet"
	"seed/backend/storage"
	"seed/backend/util/future"
	"seed/backend/util/must"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestSync(t *testing.T) {
	t.Skip("TODO(hm24): include back when delegations substitutes and changes are implemented")
	/*
	   t.Parallel()

	   alice := makeTestNode(t, "alice")
	   bob := makeTestNode(t, "bob")
	   ctx := context.Background()

	   require.NoError(t, alice.Connect(ctx, bob.AddrInfo()))

	   entity := hyper.NewEntity("foo")

	   	blob, err := entity.CreateChange(entity.NextTimestamp(), alice.Account, getDelegation(ctx, alice.Device, alice.Blobs), map[string]any{
	   		"name": "alice",
	   	})

	   require.NoError(t, err)
	   require.NoError(t, alice.Blobs.SaveBlob(ctx, blob))

	   res, err := bob.Syncer.SyncAll(ctx)
	   require.NoError(t, err)
	   require.Equalf(t, int64(0), res.NumSyncFailed, "unexpected number of sync failures: %v", res.Errs)
	   require.Equal(t, int64(1), res.NumSyncOK, "unexpected number of successful syncs")

	   	{
	   		blk, err := bob.Blobs.IPFSBlockstoreReader().Get(ctx, blob.CID)
	   		require.NoError(t, err)

	   		require.Equal(t, blob.Data, blk.RawData(), "bob must sync alice's change intact")
	   	}
	*/
}

func makeTestNode(t *testing.T, name string) testNode {
	db := storage.MakeTestDB(t)

	idx := index.NewIndex(db, logging.New("seed/hyper", "debug"), nil)

	cfg := config.Default()
	cfg.P2P.Port = 0
	cfg.P2P.NoRelay = true
	cfg.P2P.BootstrapPeers = nil
	cfg.P2P.NoMetrics = true
	store := storage.MakeTestRepo(t)
	n, err := mttnet.New(cfg.P2P, store.Device(), store.KeyStore(), store.DB(), idx, must.Do2(zap.NewDevelopment()).Named(name))
	require.NoError(t, err)

	errc := make(chan error, 1)
	ctx, cancel := context.WithCancel(context.Background())
	f := future.New[*mttnet.Node]()

	require.NoError(t, f.Resolve(n))
	go func() {
		errc <- n.Start(ctx)
	}()

	t.Cleanup(func() {
		require.NoError(t, <-errc)
	})

	select {
	case <-n.Ready():
	case err := <-errc:
		require.NoError(t, err)
	}

	t.Cleanup(cancel)
	accountKey, err := core.NewKeyPairRandom()
	require.NoError(t, err)
	require.NoError(t, store.KeyStore().StoreKey(ctx, "main", accountKey))

	identity := core.NewIdentity(accountKey.PublicKey, store.Device())
	return testNode{
		Node:    n,
		Account: accountKey,
		Index:   idx,
		Device:  identity,
		Syncer:  NewService(cfg.Syncing, must.Do2(zap.NewDevelopment()).Named(name), db, nil, n), // TODO(hm24): add indexer here.
	}
}

type testNode struct {
	*mttnet.Node
	Account core.KeyPair
	Device  core.Identity
	Index   *index.Index
	Syncer  *Service
}
