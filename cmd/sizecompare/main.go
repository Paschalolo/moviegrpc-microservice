package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"google.golang.org/protobuf/proto"
	"movieexample.com/gen"
	model "movieexample.com/metadata/pkg/model"
)

var metadata = &model.MetaData{
	ID:          "123",
	Title:       "The move 2 ",
	Description: "sequel of the legendary the movie",
	Director:    "foo bars ",
}

var genMetadata = &gen.Metadata{
	Id:          "123",
	Title:       "The move 2 ",
	Description: "sequel of the legendary the movie",
	Director:    "foo bars ",
}

func serializetoJson(m *model.MetaData) ([]byte, error) {
	return json.Marshal(m)
}
func serializeToXML(m *model.MetaData) ([]byte, error) {
	return xml.Marshal(m)
}
func serializetoProto(m *gen.Metadata) ([]byte, error) {
	return proto.Marshal(m)
}
func main() {
	jsonBYtes, err := serializetoJson(metadata)
	if err != nil {
		panic(err)
	}
	xmlBytes, err := serializeToXML(metadata)
	if err != nil {
		panic(err)
	}
	protoByte, err := serializetoProto(genMetadata)
	if err != nil {
		panic(err)
	}
	fmt.Printf("JSON size:\t%dB\n", len(jsonBYtes))
	fmt.Printf("xml size:\t%dB\n", len(xmlBytes))
	fmt.Printf("proto size:\t%dB\n", len(protoByte))

}
