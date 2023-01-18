package go2com

import (
	"github.com/okieraised/go2com/pkg/dicom/dcm_io"
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

	parser, err := dcm_io.NewDCMFileParser(fPath, dcm_io.WithSkipPixelData(false), dcm_io.WithSkipDataset(false))
	if err != nil {
		log.Fatal(err)
	}
	err = parser.Parse()
	if err != nil {
		log.Fatal(err)
	}
}
