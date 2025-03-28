package grpc

import (
	"context"

	"movieexample.com/gen"
	"movieexample.com/internal/grpcutil"
	"movieexample.com/pkg/discovery"
	ratingmodel "movieexample.com/rating/pkg/model"
)

type Gateway struct {
	registry discovery.Registery
}

// New creates a new grpc gateway for a movie rating  services
func New(registry discovery.Registery) *Gateway {
	return &Gateway{registry: registry}
}

// GetAggregatedRating returns the aggregated rating for a
// record or Errnotfound if there are no rating for it
func (g *Gateway) GetAggregatedRating(ctx context.Context, recordID ratingmodel.RecordID, recordType ratingmodel.RecordType) (float64, error) {
	conn, err := grpcutil.ServiceConnection(ctx, "rating", g.registry)
	if err != nil {
		return 0, err
	}

	defer conn.Close()
	client := gen.NewRatingServiceClient(conn)
	resp, err := client.GetAggregatedRating(ctx, &gen.GetAggregatedRatingRequest{
		RecordId:   string(recordID),
		RecordType: string(recordType),
	})
	if err != nil {
		return 0, err
	}

	return resp.RatingValue, nil
}

func (g *Gateway) PutRating(ctx context.Context, recordID ratingmodel.RecordID, recordType ratingmodel.RecordType, rating *ratingmodel.Rating) error {
	conn, err := grpcutil.ServiceConnection(ctx, "rating", g.registry)
	if err != nil {
		return err
	}
	defer conn.Close()
	client := gen.NewRatingServiceClient(conn)
	_, err = client.PutRating(ctx, &gen.PutRatingRequest{
		RecordId:    string(recordID),
		UserId:      string(rating.UserID),
		RecordType:  string(recordType),
		RatingValue: int32(rating.Value),
	})
	if err != nil {
		return err
	}

	return nil
}
