package testutil

import (
	"movieexample.com/gen"
	"movieexample.com/pkg/discovery"
	"movieexample.com/rating/internal/controller/rating"
	grpchandler "movieexample.com/rating/internal/handler/grpc"
	"movieexample.com/rating/internal/repository/memory"
)

func NewTestRatingGRPCServer(registry discovery.Registery) gen.RatingServiceServer {
	repo := memory.New()
	ctrl := rating.New(repo, nil)
	return grpchandler.New(ctrl)
}
