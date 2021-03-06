package session

import (
	"context"
	"sync"

	"golang.org/x/xerrors"
)

// StorageMemory implements in-memory session storage.
// Goroutine-safe.
type StorageMemory struct {
	mux  sync.Mutex
	data []byte
}

// LoadSession loads session from memory.
func (s *StorageMemory) LoadSession(ctx context.Context) ([]byte, error) {
	if s == nil {
		return nil, ErrNotFound
	}
	s.mux.Lock()
	defer s.mux.Unlock()

	if len(s.data) == 0 {
		return nil, ErrNotFound
	}

	return s.data, nil
}

// StoreSession stores session to memory.
func (s *StorageMemory) StoreSession(ctx context.Context, data []byte) error {
	if s == nil {
		return xerrors.New("StoreSession called on StorageMemory(nil)")
	}
	s.mux.Lock()
	defer s.mux.Unlock()

	s.data = data
	return nil
}
