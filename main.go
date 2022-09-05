package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/okieraised/go2com/internal/constants"
	"github.com/okieraised/go2com/pkg/dicom/dataset"
	"github.com/okieraised/go2com/pkg/dicom/element"
	"github.com/okieraised/go2com/pkg/dicom/reader"
)

const (
	filePath = "./test_data/File 2.dcm"
)

type Parser struct {
	reader   reader.DcmReader
	dataset  dataset.Dataset
	metadata dataset.Dataset
}

func NewParser(fileReader io.Reader, bytesToRead int64, readPixel bool) (*Parser, error) {
	dcmReader := reader.NewDcmReader(bufio.NewReader(fileReader), readPixel)

	parser := Parser{
		reader: dcmReader,
	}

	return &parser, nil
}

func (p *Parser) validateDicom() error {
	preamble, err := p.reader.Peek(128 + 4)
	if err != nil {
		return fmt.Errorf("cannot read the first 132 bytes")
	}
	if string(preamble[128:]) != constants.MAGIC_STRING {
		return fmt.Errorf("not a valid dicom")
	}
	return nil
}

func (p *Parser) readHeader() error {
	p.reader.Skip(132)

	i := 0
	for {
		if i > 90 {
			return nil
		}
		var err error
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}

		res, err := element.ReadElement(p.reader, false)

		if err != nil {
			return err
		}

		// if res.Tag.Group != 0x0002 {
		// 	return nil
		// }

		if res.Tag.Group == 0x7FE0 && res.Tag.Element == 0010 {
			return nil
		}

		fmt.Println("res:", res)

		i++
	}

	return nil
}

func main() {

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
	fmt.Println("File size: ", fileSize)

	parser, err := NewParser(file, fileSize, false)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = parser.validateDicom()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = parser.readHeader()
	if err != nil {
		fmt.Println(err)
		return
	}

}

func (p *Parser) GetMetadata() dataset.Dataset {
	return p.metadata
}
