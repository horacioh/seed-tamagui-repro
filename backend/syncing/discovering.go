package syncing

import (
	"context"
	"fmt"
	"seed/backend/ipfs"
	"sync"
	"time"

	"github.com/multiformats/go-multicodec"
	"go.uber.org/zap"
)

const defaultDiscoveryTimeout = time.Second * 30

// DiscoverObject discovers an object in the network. If not found, then it returns an error
// If found, this function will store the object locally so that it can be gotten like any
// other local object.
func (s *Service) DiscoverObject(ctx context.Context, entityID, version string) error {
	if s.cfg.NoDiscovery {
		return fmt.Errorf("remote content discovery is disabled")
	}
	if version != "" {
		return fmt.Errorf("Discovering by version is not implemented yet")
	}

	ctx, cancel := context.WithTimeout(ctx, defaultDiscoveryTimeout)
	defer cancel()
	c, err := ipfs.NewCID(uint64(multicodec.Raw), uint64(multicodec.Identity), []byte(entityID))
	if err != nil {
		return fmt.Errorf("Couldn't encode eid into CID: %w", err)
	}

	// Arbitrary number of maximum providers
	maxProviders := 15

	// If we are looking for a specific version, we don't need to limit the number of providers,
	// because we will short-circuit as soon as we found the desired version.
	if version != "" {
		maxProviders = 0
	}

	peers := s.bitswap.FindProvidersAsync(ctx, c, maxProviders)

	var wg sync.WaitGroup

	for p := range peers {
		p := p
		wg.Add(1)
		go func() {
			defer wg.Done()
			log := s.log.With(
				zap.String("entity", entityID),
				zap.String("CID", c.String()),
				zap.String("peer", p.String()),
			)
			log.Debug("DiscoveredProvider")
			if err := s.SyncWithPeer(ctx, p.ID, entityID); err != nil {
				log.Debug("FinishedSyncingWithProvider", zap.Error(err))
				return
			}
		}()
	}

	wg.Wait()

	return nil // fmt.Errorf("Found %d providers but could not sync with them:", len(peers))
}