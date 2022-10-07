package go2com

import (
	"fmt"
	"github.com/okieraised/go2com/internal/utils"
	"github.com/stretchr/testify/assert"
	"os"
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
	}
	err := utils.CPUProfilingFunc(fn, "./cpu.pprof")
	assert.NoError(err)
}

func TestNewParser(t *testing.T) {
	assert := assert.New(t)
	InitTagDict()
	file, err := os.Open("./test_data/027.dcm")
	assert.NoError(err)

	defer file.Close()
	info, err := file.Stat()
	assert.NoError(err)
	fileSize := info.Size()

	parser, err := NewParser(file, fileSize, true, false)
	assert.NoError(err)
	err = parser.Parse()
	assert.NoError(err)
}

func TestNewParser2(t *testing.T) {
	assert := assert.New(t)
	InitTagDict()
	filePaths, err := utils.ReadDirRecursively("./test_data")
	assert.NoError(err)
	for _, fPath := range filePaths {
		fmt.Println("process:", fPath)
		file, err := os.Open(fPath)
		assert.NoError(err)

		defer file.Close()
		info, err := file.Stat()
		assert.NoError(err)
		fileSize := info.Size()

		parser, err := NewParser(file, fileSize, false, false)
		assert.NoError(err)
		err = parser.Parse()
		assert.NoError(err)
	}
}

func TestNewParser3(t *testing.T) {
	assert := assert.New(t)
	InitTagDict()
	filePaths, err := utils.ReadDirRecursively("/home/tripg/Downloads/mammo_dicoms")
	assert.NoError(err)
	for _, fPath := range filePaths {
		fmt.Println("process:", fPath)
		file, err := os.Open(fPath)
		assert.NoError(err)

		defer file.Close()
		info, err := file.Stat()
		assert.NoError(err)
		fileSize := info.Size()

		parser, err := NewParser(file, fileSize, true, false)
		assert.NoError(err)
		err = parser.Parse()
		assert.NoError(err)
	}
}

//func TestNewParser3(t *testing.T) {
//	assert := assert.New(t)
//	InitTagDict()
//	file, err := os.Open("/home/tripg/Documents/dicom/oct/1.dcm")
//	assert.NoError(err)
//
//	defer file.Close()
//	info, err := file.Stat()
//	assert.NoError(err)
//	fileSize := info.Size()
//
//	parser, err := NewParser(file, fileSize, true, false)
//	assert.NoError(err)
//	err = parser.Parse()
//	assert.NoError(err)
//
//	for _, elem := range parser.dataset.Elements {
//		fmt.Println(elem)
//	}
//}
