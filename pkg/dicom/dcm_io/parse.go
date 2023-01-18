package dcm_io

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/okieraised/go2com/internal/utils"
	"github.com/okieraised/go2com/pkg/dicom/element"
	"github.com/okieraised/go2com/pkg/dicom/reader"
	"github.com/okieraised/go2com/pkg/dicom/tag"
	"github.com/okieraised/go2com/pkg/dicom/vr"
	"io"
	"os"
	"strings"

	//_ "github.com/okieraised/go2com/pkg/dicom/tag"
	"github.com/okieraised/go2com/pkg/dicom/uid"
)

const (
	MagicString = "DICM"
)

// DcmParser implements the field required to parse the dicom file
type DcmParser struct {
	filePath      string
	fileContent   []byte
	fileReader    *io.Reader
	reader        reader.DcmReader
	dataset       Dataset
	metadata      Dataset
	fileSize      int64
	skipDataset   bool
	skipPixelData bool
}

// NewDCMFileParser creates new parser from input file path with default options and/or with user-specified options
func NewDCMFileParser(filePath string, options ...func(*DcmParser)) (*DcmParser, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	fInfo, err := f.Stat()
	if err != nil {
		return nil, err
	}
	dcmReader := reader.NewDICOMReader(bufio.NewReader(f))
	parser := &DcmParser{
		reader:   dcmReader,
		fileSize: fInfo.Size(),
	}
	for _, opt := range options {
		opt(parser)
	}
	return parser, nil
}

// WithSkipPixelData provides option to skip reading pixel data (7FE0,0010).
// If true, pixel data is skipped. If false, pixel data will be read
func WithSkipPixelData(skipPixelData bool) func(*DcmParser) {
	return func(s *DcmParser) {
		s.skipPixelData = skipPixelData
	}
}

// WithSkipDataset provides option to read only the metadata header.
// If true, only the meta header is read, else, the dataset will be read
func WithSkipDataset(skipPixelDataset bool) func(*DcmParser) {
	return func(s *DcmParser) {
		s.skipPixelData = skipPixelDataset
	}
}

// IsValidDICOM checks if the dicom file follows the standard by having 128 bytes preamble followed by the magic string 'DICM'
func (p *DcmParser) IsValidDICOM() error {
	preamble, err := p.reader.Peek(128 + 4)
	if err != nil {
		return fmt.Errorf("cannot read the first 132 bytes: %v", err)
	}
	if string(preamble[128:]) != MagicString {
		return fmt.Errorf("file is not in valid dicom format")
	}
	return nil
}

// Parse reads the DICOM file and parses it into array of elements
func (p *DcmParser) Parse() error {
	err := p.parse()
	if err != nil {
		return err
	}
	return nil
}

// GetMetadata returns the file meta header
func (p *DcmParser) GetMetadata() Dataset {
	return p.metadata
}

// GetDataset returns the dataset
func (p *DcmParser) GetDataset() Dataset {
	return p.dataset
}

// GetElementByTagString returns the element value of the input tag
// Tag should be in (gggg,eeee) or ggggeeee format
func (p *DcmParser) GetElementByTagString(tagStr string) (interface{}, error) {
	tagStr = utils.FormatTag(tagStr)

	if strings.HasPrefix(tagStr, "0002") {
		for _, elem := range p.metadata.Elements {
			if tagStr == elem.Tag.StringWithoutParentheses() {
				return elem.Value, nil
			}
		}
		return nil, fmt.Errorf("cannot find tag %s", tagStr)
	} else {
		for _, elem := range p.dataset.Elements {
			if tagStr == elem.Tag.StringWithoutParentheses() {
				return elem.Value, nil
			}
		}
		return nil, fmt.Errorf("cannot find tag %s", tagStr)
	}
}

func (p *DcmParser) RetrieveFileUID() (*DicomUID, error) {
	return p.dataset.RetrieveFileUID()
}

//----------------------------------------------------------------------------------------------------------------------
// Unexported methods
//----------------------------------------------------------------------------------------------------------------------
func (p *DcmParser) parse() error {
	p.setFileSize()
	err := p.IsValidDICOM()
	if err != nil {
		return err
	}
	_ = p.reader.Skip(132)
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

// setFileSize sets the file size to the reader
func (p *DcmParser) setFileSize() {
	_ = p.reader.SetFileSize(p.fileSize)
}

// parseMetadata parses the file meta information according to
// https://dicom.nema.org/dicom/2013/output/chtml/part10/chapter_7.html
// the File Meta Information shall be encoded using the Explicit VR Little Endian Transfer Syntax
// (UID=1.2.840.10008.1.2.1)
func (p *DcmParser) parseMetadata() error {
	var transferSyntaxUID string
	var metadata []*element.Element

	res, err := element.ReadElement(p.reader, p.reader.IsImplicit(), p.reader.ByteOrder())
	if err != nil {
		return err
	}

	metaGroupLength, ok := (res.Value.RawValue).(int)
	if !ok {
		return fmt.Errorf("invalid value for tag (0x%x, 0x%x)", res.Tag.Group, res.Tag.Element)
	}

	metadata = append(metadata, res)
	// Keep reading the remaining header based on metaGroupLength
	pBytes, err := p.reader.Peek(metaGroupLength)
	if err != nil {
		return err
	}

	subRd := reader.NewDICOMReader(bufio.NewReader(bytes.NewReader(pBytes)), reader.WithSkipPixelData(p.skipPixelData))
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
			transferSyntaxUID = (res.Value.RawValue).(string)
		}
		metadata = append(metadata, res)
		//fmt.Println(res)
	}
	dicomMetadata := Dataset{Elements: metadata}
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
	p.reader.SetOverallImplicit(isImplicit)

	// IMPORTANT: Additional check is needed here since there are few instances where the DICOM
	// meta header is registered as Explicit Little-Endian, but Implicit Little-Endian is used in the body
	if transferSyntaxUID == uid.ExplicitVRLittleEndian {
		firstElem, err := p.reader.Peek(6)
		if err != nil {
			return err
		}
		if !vr.VRMapper[string(firstElem[4:6])] {
			p.reader.SetTransferSyntax(binOrder, true)
		}
	}

	return nil
}

// parseDataset parses the file dataset after the file meta header
func (p *DcmParser) parseDataset() error {
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
		//fmt.Println(res)
	}
	dicomDataset := Dataset{Elements: data}
	p.dataset = dicomDataset
	return nil
}
