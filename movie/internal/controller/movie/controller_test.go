package movie

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	gen "movieexample.com/gen/mock/movie/repository"
	metadatamodel "movieexample.com/metadata/pkg/model"
	"movieexample.com/movie/internal/gateway"
	"movieexample.com/movie/pkg/model"
	ratingmodel "movieexample.com/rating/pkg/model"
)

var ratingNumber float64 = 63.999

func TestController(t *testing.T) {
	tests := []struct {
		name         string
		id           string
		expRatingres float64
		expRatingErr error
		expMeta      *metadatamodel.MetaData
		expMetaErr   error
		wantRes      *model.MovieDetails
		wantErr      error
	}{
		{
			name:         "rating not found",
			id:           "1",
			expRatingErr: gateway.ErrNotFound,
			expMeta:      &metadatamodel.MetaData{ID: "1", Title: "the house of thieves", Description: "a film about loyalty among thieves ", Director: "stevcen spilberg"},
			expMetaErr:   nil,
			wantRes:      &model.MovieDetails{Metadata: metadatamodel.MetaData{ID: "1", Title: "the house of thieves", Description: "a film about loyalty among thieves ", Director: "stevcen spilberg"}},
		},
		{
			name:       "meta not found",
			id:         "1",
			expMetaErr: gateway.ErrNotFound,
			wantErr:    ErrNotFound,
		},
		{
			name:         "unexpected rating error",
			id:           "1",
			expRatingErr: errors.New("unexpected rating error "),
			expMeta:      &metadatamodel.MetaData{ID: "1", Title: "the house of thieves", Description: "a film about loyalty among thieves ", Director: "stevcen spilberg"},
			expMetaErr:   nil,
			wantErr:      errors.New("unexpected rating error "),
		},
		{
			name:       "unexpected meta error",
			id:         "1",
			expMetaErr: errors.New("unexpected meta error"),
			wantErr:    errors.New("unexpected meta error"),
		},
		{
			name:         "success",
			id:           "1",
			expMeta:      &metadatamodel.MetaData{ID: "1", Title: "the house of thieves", Description: "a film about loyalty among thieves ", Director: "stevcen spilberg"},
			expRatingres: ratingNumber,
			wantRes:      &model.MovieDetails{Rating: &ratingNumber, Metadata: metadatamodel.MetaData{ID: "1", Title: "the house of thieves", Description: "a film about loyalty among thieves ", Director: "stevcen spilberg"}},
			wantErr:      nil, // No error expected
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			ratgate := gen.NewMockratingGateway(ctrl)
			Metagate := gen.NewMockMetaDataGateway(ctrl)
			ctx := context.Background()
			c := New(ratgate, Metagate)
			Metagate.EXPECT().Get(ctx, tt.id).Return(tt.expMeta, tt.expMetaErr)
			if tt.expMetaErr == nil {
				ratgate.EXPECT().GetAggregatedRating(ctx, ratingmodel.RecordID(tt.id), ratingmodel.RecordTypeMovie).Return(tt.expRatingres, tt.expRatingErr)
			}
			res, err := c.Get(ctx, tt.id)
			assert.Equal(t, tt.wantRes, res, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}
