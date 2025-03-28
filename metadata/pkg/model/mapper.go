package pkg

import (
	"movieexample.com/gen"
)

// MetadataToproto converts a Metadata struct into a
// generated proto counterpart
func MetadataToproto(m *MetaData) *gen.Metadata {
	return &gen.Metadata{
		Id:          m.ID,
		Title:       m.Title,
		Description: m.Description,
		Director:    m.Director,
	}
}

// MetadataFromProto converts a generated proto counterpart
//into a Metadata structure

func MetadataFromProto(m *gen.Metadata) *MetaData {
	return &MetaData{
		ID:          m.Id,
		Title:       m.Title,
		Description: m.Description,
		Director:    m.Director,
	}
}
