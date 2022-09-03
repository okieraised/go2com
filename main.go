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
	filePath = "./test_data/File 12.dcm"
)

type Parser struct {
	reader   reader.DcmReader
	dataset  dataset.Dataset
	metadata dataset.Dataset
	file     *os.File
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

	// group, err := p.reader.ReadUInt16()
	// if err != nil {
	// 	return err
	// }
	// element, err := p.reader.ReadUInt16()
	// if err != nil {
	// 	return err
	// }

	// res, err := tag.Find(tag.DicomTag{
	// 	Group:   group,
	// 	Element: element,
	// })
	// if err != nil {
	// 	return err
	// }
	// fmt.Println("tag:", res)

	res, err := element.ReadElement(p.reader, false)
	if err != nil {
		return err
	}
	fmt.Println("res:", res)

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

	// data, err := p.reader.Peek(128 + 4)
	// if err != nil {
	// 	return nil, err
	// }
	// if string(data[128:]) != magicWord {
	// 	return nil, nil
	// }

	// buf := make([]byte, 8)
	// reader := bufio.NewReaderSize(file, 4096)
	// reader.Discard(132)

	// for {
	// 	_, err := reader.Read(buf)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		if err == io.EOF {
	// 			break
	// 		} else {
	// 			return
	// 		}
	// 	}
	// 	dicomTag := buf[0:4]
	// 	fmt.Println("dicomTag: ", dicomTag)

	// 	m := binary.LittleEndian.Uint16(dicomTag)

	// 	b := make([]byte, 4)
	// 	binary.BigEndian.PutUint16(b, m)
	// 	fmt.Println(b)
	// 	break
	// }

}

func (p *Parser) GetMetadata() dataset.Dataset {
	return p.metadata
}