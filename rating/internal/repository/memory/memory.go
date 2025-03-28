package memory

import (
	"context"
	"fmt"
	"log"

	"movieexample.com/rating/internal/repository"
	"movieexample.com/rating/pkg/model"
)

// Repository defines rating repository
type Repository struct {
	data map[model.RecordType]map[model.RecordID][]model.Rating
}

// New creates a new memory repository
func New() *Repository {
	return &Repository{map[model.RecordType]map[model.RecordID][]model.Rating{}}
}

// get retrives all ratings for a given function
func (r *Repository) Get(_ context.Context, recordID model.RecordID, recordType model.RecordType) ([]model.Rating, error) {
	if _, ok := r.data[recordType]; !ok {
		return nil, repository.ErrNotFound
	}
	if ratings, ok := r.data[recordType][recordID]; !ok || len(ratings) == 0 {
		return nil, repository.ErrNotFound
	}
	return r.data[recordType][recordID], nil

}

// Put adds a rating for a given record
func (r *Repository) Put(_ context.Context, recordID model.RecordID, recordType model.RecordType, rating *model.Rating) error {
	if _, ok := r.data[recordType]; !ok {
		log.Println(recordID, recordType)
		r.data[recordType] = make(map[model.RecordID][]model.Rating)
	}
	r.data[recordType][recordID] = append(r.data[recordType][recordID], *rating)
	fmt.Println(r)
	return nil
}
