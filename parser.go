package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/okieraised/go2com/internal/constants"
	"github.com/okieraised/go2com/pkg/dicom/dataset"
	"github.com/okieraised/go2com/pkg/dicom/element"
	"github.com/okieraised/go2com/pkg/dicom/reader"
	"github.com/okieraised/go2com/pkg/dicom/tag"
	"github.com/okieraised/go2com/pkg/dicom/uid"
	"io"
	"strconv"
)

type Parser struct {
	fileSize    int64
	skipDataset bool
	reader      reader.DcmReader
	dataset     dataset.Dataset
	metadata    dataset.Dataset
}

// NewParser returns a new dicom parser
// readPixel is not used yet
func NewParser(fileReader io.Reader, fileSize int64, readPixel, skipDataset bool) (*Parser, error) {
	dcmReader := reader.NewDcmReader(bufio.NewReader(fileReader), readPixel)
	parser := Parser{
		skipDataset: skipDataset,
		fileSize:    fileSize,
		reader:      dcmReader,
	}
	return &parser, nil
}

func (p *Parser) Parse() error {
	err := p.validateDicom()
	if err != nil {
		return err
	}
	err = p.parseMetadata()
	if err != nil {
		return err
	}
	if p.skipDataset {
		return nil
	}
	err = p.parseDataset()
	if err != nil {
		return err
	}
	return nil
}

func (p *Parser) GetMetadata() dataset.Dataset {
	return p.metadata
}

func (p *Parser) GetDataset() dataset.Dataset {
	return p.dataset
}

func (p *Parser) ConvertToMap() map[string]element.Element {
	res := map[string]element.Element{}
	fileMeta := p.GetMetadata()
	for _, elem := range fileMeta.Elements {
		res[elem.TagName] = *elem
	}

	fileDataset := p.GetDataset()
	for _, elem := range fileDataset.Elements {
		i := 0
		if elem.TagName == constants.PrivateTag {
			res[elem.TagName+"_"+strconv.Itoa(i)] = *elem
			i++
		} else {
			res[elem.TagName] = *elem
		}
	}
	return res
}

//----------------------------------------------------------------------------------------------------------------------
// Unexported methods
//----------------------------------------------------------------------------------------------------------------------

// validateDicom checks if the dicom file follows the standard by having 128 bytes preamble followed by the magic string 'DICM'
func (p *Parser) validateDicom() error {
	preamble, err := p.reader.Peek(128 + 4)
	if err != nil {
		return fmt.Errorf("cannot read the first 132 bytes: %v", err)
	}
	if string(preamble[128:]) != constants.MagicString {
		return fmt.Errorf("file is not in valid dicom format")
	}
	_ = p.reader.Skip(132)
	return nil
}

// parseMetadata parses the file meta information according to
// https://dicom.nema.org/dicom/2013/output/chtml/part10/chapter_7.html
// the File Meta Information shall be encoded using the Explicit VR Little Endian Transfer Syntax
// (UID=1.2.840.10008.1.2.1)
func (p *Parser) parseMetadata() error {
	var transferSyntaxUID string
	var metadata []*element.Element
	res, err := element.ReadElement(p.reader, p.reader.IsImplicit(), p.reader.ByteOrder())
	if err != nil {
		return err
	}
	metaGroupLength, ok := res.Value.(int)
	if !ok {
		return fmt.Errorf("invalid value for tag (0x%x, 0x%x)", res.Tag.Group, res.Tag.Element)
	}
	metadata = append(metadata, res)
	pBytes, err := p.reader.Peek(metaGroupLength)
	if err != nil {
		return err
	}
	br := bytes.NewReader(pBytes)
	subRd := reader.NewDcmReader(bufio.NewReader(br), false)
	for {
		res, err := element.ReadElement(subRd, p.reader.IsImplicit(), p.reader.ByteOrder())
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		if res.Tag.Compare(tag.DicomTag{
			Group:   0x0002,
			Element: 0x0010,
		}) == 0 {
			transferSyntaxUID = (res.Value).(string)
		}
	}
	dicomMetadata := dataset.Dataset{Elements: metadata}
	p.metadata = dicomMetadata
	err = p.reader.Skip(int64(metaGroupLength))
	if err != nil {
		return err
	}
	// Set transfer syntax here for the dataset parser
	binOrder, isImplicit, err := uid.ParseTransferSyntaxUID(transferSyntaxUID)
	if err != nil {
		return err
	}
	p.reader.SetTransferSyntax(binOrder, isImplicit)
	return nil
}

func (p *Parser) parseDataset() error {
	var data []*element.Element
	for {
		res, err := element.ReadElement(p.reader, p.reader.IsImplicit(), p.reader.ByteOrder())
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		data = append(data, res)
	}
	dicomDataset := dataset.Dataset{Elements: data}
	p.dataset = dicomDataset
	return nil
}
