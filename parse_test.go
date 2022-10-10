package go2com

import (
	"fmt"
	"github.com/okieraised/go2com/internal/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"reflect"
	"strings"
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

func TestProfilingParse2(t *testing.T) {
	assert := assert.New(t)
	fn := func() {
		InitTagDict()
		filePaths, err := utils.ReadDirRecursively("/home/tripg/Documents/dicom/test_data")
		assert.NoError(err)
		for _, fPath := range filePaths {
			fmt.Println("process:", fPath)
			file, err := os.Open(fPath)
			assert.NoError(err)
			if err != nil {
				fmt.Println(err)
				return
			}

			defer file.Close()
			info, err := file.Stat()
			assert.NoError(err)
			if err != nil {
				fmt.Println(err)
				return
			}
			fileSize := info.Size()

			parser, err := NewParser(file, fileSize, false, false)
			assert.NoError(err)
			err = parser.Parse()
			assert.NoError(err)
			if err != nil && !strings.Contains(err.Error(), "could not find tag") {
				fmt.Println(err)
				return
			}
		}
	}
	err := utils.CPUProfilingFunc(fn, "./cpu.pprof")
	assert.NoError(err)
}

func TestNewParser(t *testing.T) {
	assert := assert.New(t)
	InitTagDict()
	file, err := os.Open("/home/tripg/Documents/dicom/mammo_dicoms/1.2.840.113619.2.255.10452022879169.3670200508103440.2701.dicom")
	assert.NoError(err)

	defer file.Close()
	info, err := file.Stat()
	assert.NoError(err)
	fileSize := info.Size()

	parser, err := NewParser(file, fileSize, true, false)
	assert.NoError(err)
	err = parser.Parse()
	assert.NoError(err)

	tt := parser.Export(false)

	for key := range tt {
		//fmt.Println(fmt.Sprintf("%v", tt[key]))
		val := tt[key].Value.([]interface{})
		for _, subVal := range val {
			fmt.Println(reflect.ValueOf(subVal).Kind())
		}

	}
}

func TestNewParser2(t *testing.T) {
	assert := assert.New(t)
	InitTagDict()
	filePaths, err := utils.ReadDirRecursively("/home/tripg/Documents/dicom/test_data")
	assert.NoError(err)
	for _, fPath := range filePaths {
		fmt.Println("process:", fPath)
		file, err := os.Open(fPath)
		assert.NoError(err)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer file.Close()
		info, err := file.Stat()
		assert.NoError(err)
		if err != nil {
			fmt.Println(err)
			return
		}
		fileSize := info.Size()

		parser, err := NewParser(file, fileSize, false, false)
		assert.NoError(err)
		err = parser.Parse()
		assert.NoError(err)
		if err != nil && !strings.Contains(err.Error(), "could not find tag") {
			fmt.Println(err)
			return
		}
	}
}

func TestNewParser3(t *testing.T) {
	assert := assert.New(t)
	InitTagDict()
	filePaths, err := utils.ReadDirRecursively("/home/tripg/Documents/dicom/mammo_dicoms")
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

		_ = parser.Export(false)
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

func TestNewParser5(t *testing.T) {
	assert := assert.New(t)
	InitTagDict()
	//file, err := os.Open("/home/tripg/Documents/dicom/ptt1.dcm")
	//file, err := os.Open("/home/tripg/Documents/dicom/dicoms_mr_func/MR.1.3.46.670589.11.38317.5.0.4476.2014042516042547586")
	//file, err := os.Open("/home/tripg/Documents/dicom/dicoms_struct/N2D_0001.dcm")
	file, err := os.Open("/home/tripg/Documents/dicom/test_data/File 10051.dcm")
	assert.NoError(err)

	defer file.Close()
	info, err := file.Stat()
	assert.NoError(err)
	fileSize := info.Size()

	parser, err := NewParser(file, fileSize, false, false)
	assert.NoError(err)
	err = parser.Parse()
	assert.NoError(err)

	for _, elem := range parser.dataset.Elements {
		fmt.Println(elem)
	}
	//parser.Export(false)

	//tt := parser.Export(false)
	//for key := range tt {
	//	//fmt.Println(fmt.Sprintf("%v", tt[key]))
	//	val := tt[key].Value.([]interface{})
	//	for _, subVal := range val {
	//		fmt.Println(key, tt[key], reflect.ValueOf(subVal).Kind())
	//	}
	//
	//}
}
