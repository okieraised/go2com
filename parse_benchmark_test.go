package go2com

import (
	"log"
	"testing"
)

func BenchmarkNewParser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sampleParser()
	}
}

func sampleParser() {
	fPath := "./test_data/01.dcm"

	parser, err := NewDCMFileParser(fPath, WithSkipPixelData(false), WithSkipDataset(false))
	if err != nil {
		log.Fatal(err)
	}
	err = parser.Parse()
	if err != nil {
		log.Fatal(err)
	}
}
