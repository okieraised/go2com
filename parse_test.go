package go2com

import (
	"encoding/json"
	"fmt"
	"github.com/klauspost/compress/zstd"
	"os"
	"testing"
)

const (
	//filePath = "./test_data/File 18001"
	//filePath = "./test_data/File 12948.dcm"
	//filePath = "./test_data/File 1.dcm"
	//filePath = "./test_data/File 11636.dcm"
	//filePath = "./test_data/File 32000"
	filePath = "./test_data/File 4000.dcm"
	//filePath = "./test_data/File 8000"
	//filePath = "./test_data/File 12000"
	//filePath = "./test_data/File 160.dcm"
)

func TestNewParser(t *testing.T) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	info, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	fileSize := info.Size()

	//------------------------------------------------------------------------------------------------------------------
	parser, err := NewParser(file, fileSize, true, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = parser.Parse()
	if err != nil {
		fmt.Println(err)
		return
	}
	mapTag := parser.Export()
	//for key := range mapTag {
	//	fmt.Println(key, mapTag[key])
	//}

	//------------------------------------------------------------------------------------------------------------------
	b, err := json.Marshal(mapTag)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("LEN BYTES BEFORE", len(b))

	var encoder, _ = zstd.NewWriter(nil, zstd.WithEncoderLevel(zstd.SpeedFastest))

	compressed := encoder.EncodeAll(b, make([]byte, 0, len(b)))
	fmt.Println("LEN BYTES AFTER", len(compressed))

	var decoder, _ = zstd.NewReader(nil, zstd.WithDecoderLowmem(true))
	deflated, err := decoder.DecodeAll(compressed, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("LEN BYTES AFTER 2", len(deflated))

	type dst map[string]interface{}
	x := &dst{}
	json.Unmarshal(deflated, x)
	fmt.Println(x)

	//mt := parser.GetMetadata()
	//
	//for _, d := range mt.Elements {
	//	fmt.Println("res", d)
	//}
	//

	//ds := parser.GetDataset()
	//for _, d := range ds.Elements {
	//	fmt.Println("res", d)
	//}
}
