package grpc

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"movieexample.com/gen"
	"movieexample.com/rating/internal/controller/rating"
	"movieexample.com/rating/pkg/model"
)

// handler defines a grpc ratinf for api server
type Handler struct {
	gen.UnimplementedRatingServiceServer
	ctrl *rating.Controller
}

// New creates a new movie rating grpc handler
func New(ctrl *rating.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

// GetAggregatedRating returns the aggregated rating for a // record.
func (h *Handler) GetAggregatedRating(ctx context.Context, req *gen.GetAggregatedRatingRequest) (*gen.GetAggregatedRatingResponse, error) {
	if req == nil || req.RecordId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "nilreq or empty id ")
	}
	v, err := h.ctrl.GetAggregatedRating(ctx, model.RecordID(req.RecordId), model.RecordType(req.RecordType))
	if err != nil && errors.Is(err, rating.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, "ratings not found for record")
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &gen.GetAggregatedRatingResponse{RatingValue: v}, nil
}

func (h *Handler) PutRating(ctx context.Context, req *gen.PutRatingRequest) (*gen.PutRatingResponse, error) {
	if req == nil || req.RecordId == "" || req.UserId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "nilreq or user id or record id ")
	}
	if err := h.ctrl.PutRating(ctx, model.RecordID(req.RecordId), model.RecordType(req.RecordType), &model.Rating{UserID: model.UserID(req.UserId), Value: model.RatingValue(req.RatingValue)}); err != nil {
		return nil, err
	}
	return &gen.PutRatingResponse{}, nil
}
