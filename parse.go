package go2com

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/okieraised/go2com/internal/utils"
	"github.com/okieraised/go2com/pkg/dicom/dataset"
	"github.com/okieraised/go2com/pkg/dicom/element"
	"github.com/okieraised/go2com/pkg/dicom/reader"
	"github.com/okieraised/go2com/pkg/dicom/tag"
	"github.com/okieraised/go2com/pkg/dicom/vr"
	"os"

	//_ "github.com/okieraised/go2com/pkg/dicom/tag"
	"github.com/okieraised/go2com/pkg/dicom/uid"
	"io"
	"strings"
)

const (
	MagicString = "DICM"
)

// Parser implements the field required to parse the dicom file
type Parser struct {
	filePath      string
	fileContent   []byte
	fileReader    *io.Reader
	reader        reader.DcmReader
	dataset       dataset.Dataset
	metadata      dataset.Dataset
	fileSize      int64
	skipDataset   bool
	skipPixelData bool
}

// Deprecated: this initialization will be triggered by calling init() in tag pkg
func InitTagDict() {
	tag.InitTagDict()
}

// NewParser returns a new dicom parser
// Deprecated: NewParser will be replaced by NewDCMFileParser in future release
func NewParser(fileReader io.Reader, fileSize int64, skipPixelData, skipDataset bool) (*Parser, error) {
	dcmReader := reader.NewDcmReader(bufio.NewReader(fileReader), skipPixelData)
	parser := Parser{
		skipPixelData: skipPixelData,
		skipDataset:   skipDataset,
		fileSize:      fileSize,
		reader:        dcmReader,
	}
	return &parser, nil
}

//----------------------------------------------------------------------------------------------------------------------
// New implementation of DICOM parser
//----------------------------------------------------------------------------------------------------------------------

// NewDCMFileParser creates new parser from input file path with default options and/or with user-specified options
func NewDCMFileParser(filePath string, options ...func(*Parser)) (*Parser, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	fInfo, err := f.Stat()
	if err != nil {
		return nil, err
	}
	dcmReader := reader.NewDICOMReader(bufio.NewReader(f))
	parser := &Parser{
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
func WithSkipPixelData(skipPixelData bool) func(*Parser) {
	return func(s *Parser) {
		s.skipPixelData = skipPixelData
	}
}

// WithSkipDataset provides option to read only the metadata header.
// If true, only the meta header is read, else, the dataset will be read
func WithSkipDataset(skipPixelDataset bool) func(*Parser) {
	return func(s *Parser) {
		s.skipPixelData = skipPixelDataset
	}
}

// IsValidDICOM checks if the dicom file follows the standard by having 128 bytes preamble followed by the magic string 'DICM'
func (p *Parser) IsValidDICOM() error {
	preamble, err := p.reader.Peek(128 + 4)
	if err != nil {
		return fmt.Errorf("cannot read the first 132 bytes: %v", err)
	}
	if string(preamble[128:]) != MagicString {
		return fmt.Errorf("file is not in valid dicom format")
	}
	return nil
}

//----------------------------------------------------------------------------------------------------------------------

// Parse reads the DICOM file and parses it into array of elements
func (p *Parser) Parse() error {
	err := p.parse()
	if err != nil {
		return err
	}
	return nil
}

// GetMetadata returns the file meta header
func (p *Parser) GetMetadata() dataset.Dataset {
	return p.metadata
}

// GetDataset returns the dataset
func (p *Parser) GetDataset() dataset.Dataset {
	return p.dataset
}

// GetElementByTagString returns the element value of the input tag
// Tag should be in (gggg,eeee) or ggggeeee format
func (p *Parser) GetElementByTagString(tagStr string) (interface{}, error) {
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

//----------------------------------------------------------------------------------------------------------------------
// Unexported methods
//----------------------------------------------------------------------------------------------------------------------
func (p *Parser) parse() error {
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
func (p *Parser) setFileSize() {
	_ = p.reader.SetFileSize(p.fileSize)
}

// parseMetadata parses the file meta information according to
// https://dicom.nema.org/dicom/2013/output/chtml/part10/chapter_7.html
// the File Meta Information shall be encoded using the Explicit VR Little Endian Transfer Syntax
// (UID=1.2.840.10008.1.2.1)
func (p *Parser) parseMetadata() error {
	var transferSyntaxUID string
	var metadata []*element.Element

	//// File meta header is mostly always explicit (except for 1.2.840.113619.5.2 [Implicit VR Big Endian DLX (G.E Private)])
	//// where the header is implicit
	//firstHeaderElemBytes, err := p.reader.Peek(4 + 2 + 2 + 4) // 4 bytes tag + 2 bytes VR + 2 bytes VL + 4 bytes value length
	//if err != nil {
	//	return err
	//}
	//subRd1 := reader.NewDcmReader(bufio.NewReader(bytes.NewReader(firstHeaderElemBytes)), p.skipPixelData)
	//res, err := element.ReadElement(subRd1, p.reader.IsImplicit(), p.reader.ByteOrder())
	//if err != nil {
	//	return err
	//}
	//metaGroupLength, ok := (res.Value.RawValue).(int)
	//if !ok {
	//	//fmt.Println("GOT HERE")
	//	p.reader.SetTransferSyntax(p.reader.ByteOrder(), true)
	//	res, err = element.ReadElement(p.reader, p.reader.IsImplicit(), p.reader.ByteOrder())
	//	if err != nil {
	//		return err
	//	}
	//	metaGroupLength, ok = (res.Value.RawValue).(int)
	//	if !ok {
	//		return fmt.Errorf("invalid value for tag (0x%x, 0x%x)", res.Tag.Group, res.Tag.Element)
	//	}
	//} else {
	//	p.reader.Skip(12)
	//}

	//------------------------------------------------------------------------------------------------------------------
	res, err := element.ReadElement(p.reader, p.reader.IsImplicit(), p.reader.ByteOrder())
	if err != nil {
		return err
	}

	metaGroupLength, ok := (res.Value.RawValue).(int)
	if !ok {
		return fmt.Errorf("invalid value for tag (0x%x, 0x%x)", res.Tag.Group, res.Tag.Element)
	}
	//------------------------------------------------------------------------------------------------------------------

	metadata = append(metadata, res)
	// Keep reading the remaining header based on metaGroupLength
	pBytes, err := p.reader.Peek(metaGroupLength)
	if err != nil {
		return err
	}
	subRd := reader.NewDcmReader(bufio.NewReader(bytes.NewReader(pBytes)), p.skipPixelData)
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
		//fmt.Println(res)
	}
	dicomDataset := dataset.Dataset{Elements: data}
	p.dataset = dicomDataset
	return nil
}
