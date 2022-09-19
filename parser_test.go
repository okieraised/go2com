package parser

import (
	"fmt"
	"os"
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
	filePath = "./test_data/File 160.dcm"
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

	//mt := parser.GetMetadata()
	//ds := parser.GetDataset()

	//for _, d := range mt.Elements {
	//	fmt.Println("res", d)
	//}
	//
	//for _, d := range ds.Elements {
	//	fmt.Println("res", d)
	//}

	//res := parser.ConvertToMap()
	//fmt.Println("res", res)
}
