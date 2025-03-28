package metadata

import (
	"context"
	"errors"

	"movieexample.com/metadata/internal/repository"
	model "movieexample.com/metadata/pkg/model"
)

type metadataRepository interface {
	Get(ctx context.Context, id string) (*model.MetaData, error)
	Put(ctx context.Context, id string, metadata *model.MetaData) error
}

// controller defines a metadata service controller
type Controller struct {
	repo metadataRepository
}

// New creates a metadata service controller
func New(repo metadataRepository) *Controller {
	return &Controller{repo}
}

// get return movie metadata by id
func (c *Controller) Get(ctx context.Context, id string) (*model.MetaData, error) {
	res, err := c.repo.Get(ctx, id)
	if err != nil && errors.Is(err, repository.ErrorNotFound) {
		return nil, repository.ErrorNotFound
	} else if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Controller) Put(ctx context.Context, m *model.MetaData) error {
	return c.repo.Put(ctx, m.ID, m)
}
