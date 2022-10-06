package go2com

import (
	"fmt"
	"github.com/okieraised/go2com/internal/utils"
	"github.com/stretchr/testify/assert"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

func TestProfilingParse(t *testing.T) {
	assert := assert.New(t)
	fn := func() {
		InitTagDict()
		file, err := os.Open("./test_data/01.dcm")
		assert.NoError(err)

		defer file.Close()
		info, err := file.Stat()
		assert.NoError(err)
		fileSize := info.Size()

		parser, err := NewParser(file, fileSize, false, false)
		assert.NoError(err)
		err = parser.Parse()
		assert.NoError(err)

		for _, elem := range parser.GetMetadata().Elements {
			fmt.Println(elem)
		}
	}
	err := utils.CPUProfilingFunc(fn, "./cpu_concurrent.pprof")
	assert.NoError(err)
}

func TestNewParser(t *testing.T) {
	assert := assert.New(t)
	InitTagDict()
	file, err := os.Open("./test_data/02.dcm")
	assert.NoError(err)

	defer file.Close()
	info, err := file.Stat()
	assert.NoError(err)
	fileSize := info.Size()

	parser, err := NewParser(file, fileSize, false, false)
	assert.NoError(err)
	err = parser.Parse()
	assert.NoError(err)

	for _, elem := range parser.GetMetadata().Elements {
		fmt.Println(elem)
	}
}

//func TestNewParser2(t *testing.T) {
//	assert := assert.New(t)
//	InitTagDict()
//	file, err := os.Open("/home/tripg/Documents/dicom/dcm/vietnhat/test/2022/04/04/1.2.826.0.1.3680043.2.4852.20220404.183125705.659403268/1.2.826.0.1.3680043.2.4852.20220404.183125713.920602682.dcm")
//	assert.NoError(err)
//
//	defer file.Close()
//	info, err := file.Stat()
//	assert.NoError(err)
//	fileSize := info.Size()
//
//	parser, err := NewParser(file, fileSize, false, false)
//	assert.NoError(err)
//	err = parser.Parse()
//	assert.NoError(err)
//
//	for _, elem := range parser.GetMetadata().Elements {
//		fmt.Println(elem)
//	}
//}

func TestCount(t *testing.T) {
	fPath := "/home/tripg/Documents/dicom/dcm/vietnhat"

	i := 0

	visit := func(path string, f fs.DirEntry, err error) error {
		if !f.IsDir() {
			i++
		}
		return nil
	}
	err := filepath.WalkDir(fPath, visit)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("len", i)
}
