package grpc

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"
	"movieexample.com/gen"
	"movieexample.com/metadata/internal/controller/metadata"
	"movieexample.com/metadata/internal/repository"
	model "movieexample.com/metadata/pkg/model"
)

// Handler defines a movie metadata gRPC handler
type Handler struct {
	gen.UnimplementedMetadataServiceServer
	svc *metadata.Controller
}

// New creates a new movie metadata gRPC handler
func New(svc *metadata.Controller) *Handler {
	return &Handler{svc: svc}
}

// GetMetadatbyID returns movie metadta by id
func (h *Handler) GetMetadata(ctx context.Context, req *gen.GetMetaDataRequest) (*gen.GetMetaDataResponse, error) {
	if req == nil || req.MovieId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "nil req or empty id")
	}
	m, err := h.svc.Get(ctx, req.MovieId)
	if err != nil && errors.Is(err, repository.ErrorNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &gen.GetMetaDataResponse{
		Metadata: model.MetadataToproto(m),
	}, nil
}

func (h *Handler) PutMetadata(ctx context.Context, req *gen.PutMetaDataRequest) (*gen.PutMetaDataResponse, error) {
	if req == nil || req.Metadata == nil {
		return nil, status.Errorf(codes.InvalidArgument, "nil req or metadata")
	}
	if err := h.svc.Put(ctx, model.MetadataFromProto(req.Metadata)); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &gen.PutMetaDataResponse{}, nil
}
