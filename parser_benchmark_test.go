package go2com

import (
	"log"
	"os"
	"testing"
)

func BenchmarkNewParser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sampleParser()
	}
}

func sampleParser() {
	InitTagDict()
	file, err := os.Open("./test_data/01.dcm")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	info, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	fileSize := info.Size()

	parser, err := NewParser(file, fileSize, false, false)
	if err != nil {
		log.Fatal(err)
	}
	err = parser.Parse()
	if err != nil {
		log.Fatal(err)
	}
}
