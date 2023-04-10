package go2com

import (
	"fmt"
	"github.com/okieraised/go2com/internal/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestParser_Export(t *testing.T) {
	assert := assert.New(t)
	InitTagDict()
	file, err := os.Open("./test_data/026.dcm")
	assert.NoError(err)

	defer file.Close()
	info, err := file.Stat()
	assert.NoError(err)
	fileSize := info.Size()

	parser, err := NewParser(file, fileSize, false, false)
	assert.NoError(err)
	err = parser.Parse()
	assert.NoError(err)

	mapTagVal := parser.Export(false)
	fmt.Println(mapTagVal)
}

func TestParser_Buffer(t *testing.T) {
	assert := assert.New(t)
	InitTagDict()
	file, err := os.Open("/home/tripg/workspace/dicom_buffer/D4CDC297.dcm")
	assert.NoError(err)

	defer file.Close()
	info, err := file.Stat()
	assert.NoError(err)
	fileSize := info.Size()

	parser, err := NewParser(file, fileSize, false, false)
	assert.NoError(err)
	err = parser.Parse()
	assert.NoError(err)

	mapTagVal := parser.Export(false)
	fmt.Println(mapTagVal)
}

func TestParser_MultipleFiles(t *testing.T) {

	assert := assert.New(t)

	filePaths, err := utils.ReadDirRecursively("/home/tripg/workspace/dicom/test_data")
	assert.NoError(err)
	InitTagDict()
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

func TestParser_MultipleFiles_2(t *testing.T) {

	assert := assert.New(t)

	filePaths, err := utils.ReadDirRecursively("/home/tripg/workspace/dicom/test_full")
	assert.NoError(err)
	InitTagDict()
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
