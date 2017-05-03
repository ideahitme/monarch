package manager

import (
	"sync"

	"github.com/ideahitme/monarch/storage"
)

// Manager maintains all necessary information to respond to queries
type Manager struct {
	backend storage.Storage
	sync.Mutex
}

// NewManager creates a state from a slice of bytes typically read from configmap
func NewManager(storage storage.Storage) *Manager {
	return &Manager{
		backend: storage,
	}
}
