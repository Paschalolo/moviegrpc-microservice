package main

import "testing"

func BenchmarkSerializetoJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		serializetoJson(metadata)
	}
}
func BenchmarkSerializetoXML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		serializeToXML(metadata)
	}
}
func BenchmarkSerializetoProto(b *testing.B) {
	for i := 0; i < b.N; i++ {
		serializetoProto(genMetadata)
	}
}
