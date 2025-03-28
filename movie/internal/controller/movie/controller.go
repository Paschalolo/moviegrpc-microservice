package movie

import (
	"context"
	"errors"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	metadatamodel "movieexample.com/metadata/pkg/model"
	"movieexample.com/movie/internal/gateway"
	"movieexample.com/movie/pkg/model"
	ratingmodel "movieexample.com/rating/pkg/model"
)

// ErrNotFound is returned when the movie metadata is not // found.
var ErrNotFound = errors.New("movie metadata not found")

// var ErrRatNotFound = errors.New("ratings not found for record")

type ratingGateway interface {
	GetAggregatedRating(ctx context.Context, recordID ratingmodel.RecordID, recordType ratingmodel.RecordType) (float64, error)
	PutRating(ctx context.Context, recordID ratingmodel.RecordID, recordType ratingmodel.RecordType, rating *ratingmodel.Rating) error
}

type MetaDataGateway interface {
	Get(ctx context.Context, id string) (*metadatamodel.MetaData, error)
}

// controller defines a movie sevice controller
type Controller struct {
	ratingGateway   ratingGateway
	metaDataGateway MetaDataGateway
}

// New creates a new movie service controller
func New(ratingGateway ratingGateway, metaDataGateway MetaDataGateway) *Controller {
	return &Controller{ratingGateway, metaDataGateway}
}

// Get returns the movie details including the aggregated rating and movie metadata
func (c *Controller) Get(ctx context.Context, id string) (*model.MovieDetails, error) {
	metadata, err := c.metaDataGateway.Get(ctx, id)
	if err != nil && errors.Is(err, gateway.ErrNotFound) {
		return nil, ErrNotFound
	} else if err != nil {
		log.Println("its here ")
		return nil, err
	}
	details := &model.MovieDetails{Metadata: *metadata}

	rating, err := c.ratingGateway.GetAggregatedRating(ctx, ratingmodel.RecordID(id), ratingmodel.RecordTypeMovie)
	log.Println(err)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.NotFound {

		} else {
			return nil, err
		}

	}
	details.Rating = &rating
	return details, nil

}
