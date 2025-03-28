package testutil

import (
	"movieexample.com/gen"
	"movieexample.com/movie/internal/controller/movie"
	metaGateway "movieexample.com/movie/internal/gateway/metadata/grpc"
	ratinggateway "movieexample.com/movie/internal/gateway/rating/grpc"
	grpchandler "movieexample.com/movie/internal/handler/grpc"
	"movieexample.com/pkg/discovery"
)

func NewTestMovieGRPCServer(registry discovery.Registery) gen.MovieServiceServer {
	metaGate := metaGateway.New(registry)
	ratgateway := ratinggateway.New(registry)
	ctrl := movie.New(ratgateway, metaGate)
	return grpchandler.New(ctrl)
}
