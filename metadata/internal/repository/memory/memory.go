package memory

import (
	"context"
	"sync"

	"movieexample.com/metadata/internal/repository"
	model "movieexample.com/metadata/pkg/model"
)

// Repositoiry defines a memory movie metadat repository
type Repository struct {
	sync.RWMutex
	data map[string]*model.MetaData
}

// New creates a new memory repository
func New() *Repository {
	return &Repository{data: map[string]*model.MetaData{}}
}

// Get retrives movie metadata by movie id.
func (r *Repository) Get(_ context.Context, id string) (*model.MetaData, error) {
	r.RLock()
	defer r.RUnlock()
	m, ok := r.data[id]
	if !ok {
		return nil, repository.ErrorNotFound
	}
	return m, nil
}

// Puts add movie metadata for a given movie id
func (r *Repository) Put(_ context.Context, id string, metadata *model.MetaData) error {
	r.Lock()
	defer r.Unlock()
	r.data[id] = metadata
	return nil
}
