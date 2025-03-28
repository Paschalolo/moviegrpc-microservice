package model

import model "movieexample.com/metadata/pkg/model"

// Movie Details includes movie metadata its aggregated rating
type MovieDetails struct {
	Rating   *float64       `json:"rating,omitempty"`
	Metadata model.MetaData `json:"metadata"`
}
