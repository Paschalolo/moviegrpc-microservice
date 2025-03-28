package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"movieexample.com/gen"
	"movieexample.com/internal/grpcutil"
	model "movieexample.com/metadata/pkg/model"
	"movieexample.com/pkg/discovery"
)

type Gateway struct {
	registry discovery.Registery
}

// New creates a new grpc gateway for a movie metadata services
func New(registry discovery.Registery) *Gateway {
	return &Gateway{registry: registry}
}

// Get return movie metadaat by a movie id
func (g *Gateway) Get(ctx context.Context, id string) (*model.MetaData, error) {
	conn, err := grpcutil.ServiceConnection(ctx, "metadata", g.registry)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := gen.NewMetadataServiceClient(conn)
	var resp *gen.GetMetaDataResponse
	const maxRetires = 5
	for i := 0; i < maxRetires; i++ {
		resp, err = client.GetMetadata(ctx, &gen.GetMetaDataRequest{MovieId: id})
		if err != nil {
			if shouldRetry(err) {
				continue
			}
			return nil, err
		}
		return model.MetadataFromProto(resp.Metadata), nil

	}
	return nil, err

}

func shouldRetry(err error) bool {
	e, ok := status.FromError(err)
	if !ok {
		return false
	}
	return e.Code() == codes.DeadlineExceeded || e.Code() == codes.ResourceExhausted || e.Code() == codes.Unavailable
}
