package go2com

import (
	"os"

	"github.com/okieraised/go2com/pkg/dicom/reader"
)

type DcmParser struct {
	reader reader.DcmReader
}

func NewParser(file *os.File) error {
	// file, err := os.Open(filePath)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }
	// defer file.Close()

	return nil

}
