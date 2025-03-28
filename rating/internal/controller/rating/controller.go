package rating

import (
	"context"
	"errors"

	"movieexample.com/rating/internal/repository"
	"movieexample.com/rating/pkg/model"
)

var ErrNotFound = errors.New("ratings not found for record")

type ratingRepository interface {
	Get(ctx context.Context, recordID model.RecordID, recordType model.RecordType) ([]model.Rating, error)
	Put(ctx context.Context, recordID model.RecordID, recordType model.RecordType, rating *model.Rating) error
}
type ratingIngestor interface {
	Ingest(ctx context.Context) (chan model.RatingEvent, error)
}

// controller defines a rating service controller
type Controller struct {
	repo     ratingRepository
	ingester ratingIngestor
}

// New creates a rating service controller
func New(repo ratingRepository, ingester ratingIngestor) *Controller {
	return &Controller{repo, ingester}
}

// Get AggregatedRating returns the aggregated rating for a
//record or ErrNotFOund if there are no ratings for it

func (c *Controller) GetAggregatedRating(ctx context.Context, recordId model.RecordID, recordType model.RecordType) (float64, error) {
	rating, err := c.repo.Get(ctx, recordId, recordType)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return 0, ErrNotFound
	} else if err != nil {
		return 0, err
	}
	var sum float64 = 0
	for _, r := range rating {
		sum += float64(r.Value)
	}
	return sum / float64(len(rating)), nil
}

// PutRating writes a rating for a given record
func (c *Controller) PutRating(ctx context.Context, recordID model.RecordID, recordType model.RecordType, rating *model.Rating) error {
	return c.repo.Put(ctx, recordID, recordType, rating)
}

// start ingetsion of rating events
func (c *Controller) StartIngestion(ctx context.Context) error {
	ch, err := c.ingester.Ingest(ctx)
	if err != nil {
		return err
	}
	for e := range ch {
		if err := c.PutRating(ctx, e.RecordID, e.RecordType, &model.Rating{UserID: e.UserID, Value: e.Value}); err != nil {
			return err
		}
	}
	return nil
}
