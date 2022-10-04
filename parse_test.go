package go2com

import (
	"encoding/json"
	"fmt"
	"github.com/klauspost/compress/zstd"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
	"testing"
)

const (
	//filePath = "./test_data/File 18001"
	//filePath = "./test_data/File 12948.dcm"
	//filePath = "./test_data/File 1.dcm"
	//filePath = "./test_data/File 11636.dcm"
	//filePath = "./test_data/File 32000"
	//filePath = "./test_data/File 4000.dcm"
	//filePath = "./test_data/File 8000"
	//filePath = "./test_data/File 12000"
	//filePath = "./test_data/File 160.dcm"
	//filePath = "/home/tripg/Downloads/1.dcm"
	filePath = "/home/tripg/Documents/dicom/hehe/1ee46cf8-e7a9587a-c2f49d1e-77f30744-570322cf.dcm"
)

func readDir(dir string) ([]string, error) {
	result := []string{}

	visit := func(path string, f fs.DirEntry, err error) error {
		if !f.IsDir() {
			result = append(result, path)
		}
		return nil
	}
	err := filepath.WalkDir(dir, visit)
	if err != nil {
		return result, err
	}
	return result, nil
}

func TestNewParser2(t *testing.T) {
	InitTagDict()
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
	parser, err := NewParser(file, fileSize, false, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = parser.Parse()
	if err != nil {
		fmt.Println(err)
		return
	}

	val2, err := parser.GetElementByTagString("00280010")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("VAL2", val2)

	mapTag := parser.Export(true)
	//for key := range mapTag {
	//	fmt.Println(key, mapTag[key])
	//}

	val, err := mapTag.GetElementByTagString("(0008,0008)")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("VAL", val)

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

	ds := parser.GetDataset()
	val3, err := ds.FindElementByTagName("PixelData")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("VAL3", val3.ValueLength, len(val3.Value.([]byte)))

	res4, err := ds.RetrieveFileUID()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("res4", res4)
}

func TestStress(t *testing.T) {
	dirPath := "/home/tripg/Documents/vietnhat/test/2022/04"
	dicomPath, err := readDir(dirPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	var wg sync.WaitGroup
	for _, pFile := range dicomPath {
		wg.Add(1)
		go func(pFile string) {
			defer wg.Done()
			file, err := os.Open(pFile)
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

			parser, err := NewParser(file, fileSize, false, false)
			if err != nil {
				fmt.Println(err)
				return
			}
			InitTagDict()
			err = parser.Parse()
			if err != nil {
				fmt.Println(err)
				return
			}

			var encoder, _ = zstd.NewWriter(nil, zstd.WithEncoderLevel(zstd.SpeedFastest))

			mapTag := parser.Export(true)

			b, err := json.Marshal(mapTag)
			if err != nil {
				fmt.Println(err)
				return
			}

			compressed := encoder.EncodeAll(b, make([]byte, 0, len(b)))
			fmt.Println("LEN BYTES AFTER", len(compressed))

		}(pFile)

		wg.Wait()
	}

}

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
	parser, err := NewParser(file, fileSize, false, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = parser.Parse()
	if err != nil {
		fmt.Println(err)
		return
	}

	val2, err := parser.GetElementByTagString("00280010")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("VAL2", val2)

	mapTag := parser.Export(true)
	//for key := range mapTag {
	//	fmt.Println(key, mapTag[key])
	//}

	val, err := mapTag.GetElementByTagString("(0008,0008)")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("VAL", val)

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

	ds := parser.GetDataset()
	val3, err := ds.FindElementByTagName("PixelData")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("VAL3", val3.ValueLength, len(val3.Value.([]byte)))

	res4, err := ds.RetrieveFileUID()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("res4", res4)
}
