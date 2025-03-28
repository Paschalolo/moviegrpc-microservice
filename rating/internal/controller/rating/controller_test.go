package rating

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	gen "movieexample.com/gen/mock/rating/repository"
	"movieexample.com/rating/internal/repository"
	"movieexample.com/rating/pkg/model"
)

var Mockdb = []model.Rating{
	{
		Value: 1,
	},
	{
		Value: 89,
	},
	{
		Value: 94,
	},
}

func TestGetController(t *testing.T) {
	tests := []struct {
		name       string
		id         string
		recordType string
		expRepores []model.Rating
		expRaterr  error
		wantfloat  float64
		wantErr    error
	}{
		{
			name:       "not found",
			id:         "1",
			recordType: "movie",
			expRaterr:  repository.ErrNotFound,
			wantErr:    ErrNotFound,
		},
		{
			name:       "unexpected error ",
			id:         "1",
			recordType: "movie",
			expRaterr:  errors.New("not found"),
			wantErr:    errors.New("not found"),
		},
		{
			name:       "success",
			id:         "1",
			recordType: "movie",
			expRepores: Mockdb,
			wantfloat:  61.333333333333336,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repoMock := gen.NewMockratingRepository(ctrl)
			ingestorMock := gen.NewMockratingIngestor(ctrl)
			c := New(repoMock, ingestorMock)
			ctx := context.Background()
			repoMock.EXPECT().Get(ctx, model.RecordID(tt.id), model.RecordType(tt.recordType)).Return(tt.expRepores, tt.expRaterr)
			res, err := c.GetAggregatedRating(ctx, model.RecordID(tt.id), model.RecordType(tt.recordType))
			assert.Equal(t, tt.wantfloat, res, tt.name)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}

}

func TestPutController(t *testing.T) {
	tests := []struct {
		name       string
		recordid   string
		recordType string
		rating     *model.Rating
		expRaterr  error
		wantErr    error
	}{
		{
			name:       "unexpected error ",
			recordid:   "1",
			recordType: "movie",
			rating:     &model.Rating{},
			expRaterr:  errors.New("Put unsuccesful"),
			wantErr:    errors.New("Put unsuccesful"),
		},
		{
			name:       "success",
			recordid:   "1",
			recordType: "movie",
			expRaterr:  nil,
			rating:     &model.Rating{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repoMock := gen.NewMockratingRepository(ctrl)
			ingestorMock := gen.NewMockratingIngestor(ctrl)
			c := New(repoMock, ingestorMock)
			ctx := context.Background()
			repoMock.EXPECT().Put(ctx, model.RecordID(tt.recordid), model.RecordType(tt.recordType), tt.rating).Return(tt.expRaterr)
			err := c.PutRating(ctx, model.RecordID(tt.recordid), model.RecordType(tt.recordType), tt.rating)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}

}

func TestIngestor(t *testing.T) {
	tests := []struct {
		name      string
		ingestCh  []model.RatingEvent
		ingestErr error
		putErr    error
		wantErr   error
	}{
		{
			name:      "Error in ingester ",
			ingestErr: errors.New("error in ingestor "),
			wantErr:   errors.New("error in ingestor "),
		},
		{
			name:      "Put error",
			ingestErr: nil,
			ingestCh: []model.RatingEvent{
				{
					UserID:     "1",
					RecordID:   "1",
					RecordType: "movie",
					Value:      6,
				},
			},
			putErr:  errors.New("errors in Put "),
			wantErr: errors.New("errors in Put "),
		},
		{
			name: "success",
			ingestCh: []model.RatingEvent{
				{RecordID: "1", RecordType: "movie", UserID: "user1", Value: 5},
				{RecordID: "2", RecordType: "series", UserID: "user2", Value: 8},
			},
			wantErr: nil,
		},
		{
			name:     "empty channel",
			ingestCh: []model.RatingEvent{},
			wantErr:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repoMock := gen.NewMockratingRepository(ctrl)
			ingestMock := gen.NewMockratingIngestor(ctrl)
			c := New(repoMock, ingestMock)
			ctx := context.Background()
			ingestChan := make(chan model.RatingEvent, len(tt.ingestCh))
			for _, e := range tt.ingestCh {
				ingestChan <- e
			}
			close(ingestChan)
			ingestMock.EXPECT().Ingest(ctx).Return(ingestChan, tt.ingestErr)
			if tt.ingestErr == nil {
				for _, e := range tt.ingestCh {
					repoMock.EXPECT().Put(ctx, e.RecordID, e.RecordType, &model.Rating{Value: e.Value, UserID: e.UserID}).Return(tt.putErr)
					if tt.putErr != nil {
						break
					}
				}
			}
			err := c.StartIngestion(ctx)
			assert.Equal(t, tt.wantErr, err, tt.name)
		})
	}
}
