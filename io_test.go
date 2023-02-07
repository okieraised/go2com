package go2com

import (
	"bufio"
	"fmt"
	"github.com/okieraised/go2com/internal/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewDICOMReader_1(t *testing.T) {
	assert := assert.New(t)
	filePaths, err := utils.ReadDirRecursively("/home/tripg/workspace/dicom/test_data")
	assert.NoError(err)
	for _, fPath := range filePaths {
		fmt.Println("process:", fPath)

		f, err := os.Open(fPath)
		assert.NoError(err)

		fInfo, err := f.Stat()
		assert.NoError(err)

		rd := NewDICOMReader(bufio.NewReader(f), WithSetFileSize(fInfo.Size()))

		err = rd.Parse()
		if err != nil {
			fmt.Println(err)
			return
		}
		assert.NoError(err)
	}
}

func TestNewDICOMReader_2(t *testing.T) {
	assert := assert.New(t)
	filePaths, err := utils.ReadDirRecursively("/home/tripg/workspace/dicom/test_full")
	assert.NoError(err)
	for _, fPath := range filePaths {
		fmt.Println("process:", fPath)

		f, err := os.Open(fPath)
		assert.NoError(err)

		fInfo, err := f.Stat()
		assert.NoError(err)

		rd := NewDICOMReader(bufio.NewReader(f), WithSetFileSize(fInfo.Size()))

		err = rd.Parse()
		if err != nil {
			fmt.Println(err)
			return
		}
		assert.NoError(err)
	}
}

func TestNewDICOMReader_3(t *testing.T) {
	assert := assert.New(t)
	filePaths, err := utils.ReadDirRecursively("/home/tripg/workspace/dicom/mammo_dicoms")
	assert.NoError(err)
	for _, fPath := range filePaths {
		fmt.Println("process:", fPath)

		f, err := os.Open(fPath)
		assert.NoError(err)

		fInfo, err := f.Stat()
		assert.NoError(err)

		rd := NewDICOMReader(bufio.NewReader(f), WithSetFileSize(fInfo.Size()))

		err = rd.Parse()
		if err != nil {
			fmt.Println(err)
			return
		}
		assert.NoError(err)
	}
}

func TestNewDICOMReader_4(t *testing.T) {
	assert := assert.New(t)
	filePaths, err := utils.ReadDirRecursively("/home/tripg/workspace/dicom/mammo_dicoms")
	assert.NoError(err)
	for _, fPath := range filePaths {
		fmt.Println("process:", fPath)

		f, err := os.Open(fPath)
		assert.NoError(err)

		fInfo, err := f.Stat()
		assert.NoError(err)

		rd := NewDICOMReader(bufio.NewReader(f), WithSetFileSize(fInfo.Size()))

		err = rd.Parse()
		if err != nil {
			fmt.Println(err)
			return
		}
		assert.NoError(err)
	}
}

func TestNewDICOMReader_5(t *testing.T) {
	assert := assert.New(t)
	filePaths, err := utils.ReadDirRecursively("/home/tripg/workspace/color_dicom")
	assert.NoError(err)
	for _, fPath := range filePaths {
		fmt.Println("process:", fPath)

		f, err := os.Open(fPath)
		assert.NoError(err)

		fInfo, err := f.Stat()
		assert.NoError(err)

		rd := NewDICOMReader(bufio.NewReader(f), WithSetFileSize(fInfo.Size()))

		err = rd.Parse()
		assert.NoError(err)

		_ = rd.ExportDatasetTags(false)
	}
}

func TestNewDICOMReader_6(t *testing.T) {
	assert := assert.New(t)
	filePaths, err := utils.ReadDirRecursively("/home/tripg/workspace/pydicom_dcm")
	assert.NoError(err)
	for _, fPath := range filePaths {
		fmt.Println("process:", fPath)

		f, err := os.Open(fPath)
		assert.NoError(err)

		fInfo, err := f.Stat()
		assert.NoError(err)

		rd := NewDICOMReader(bufio.NewReader(f), WithSetFileSize(fInfo.Size()))

		err = rd.Parse()
		assert.NoError(err)

		_ = rd.ExportDatasetTags(false)
	}
}

func TestNewDICOMReader_7(t *testing.T) {
	assert := assert.New(t)
	filePaths, err := utils.ReadDirRecursively("/home/tripg/workspace/dicom/dicoms_mr_func")
	assert.NoError(err)
	for _, fPath := range filePaths {
		fmt.Println("process:", fPath)

		f, err := os.Open(fPath)
		assert.NoError(err)

		fInfo, err := f.Stat()
		assert.NoError(err)

		rd := NewDICOMReader(bufio.NewReader(f), WithSetFileSize(fInfo.Size()))

		err = rd.Parse()
		assert.NoError(err)

		_ = rd.ExportDatasetTags(false)
	}
}
