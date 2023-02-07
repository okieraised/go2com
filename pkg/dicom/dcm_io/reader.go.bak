package dcm_io

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	_ "github.com/okieraised/go2com/internal/system"
	"github.com/okieraised/go2com/internal/utils"
	"github.com/okieraised/go2com/pkg/dicom/tag"
	"github.com/okieraised/go2com/pkg/dicom/uid"
	"github.com/okieraised/go2com/pkg/dicom/vr"
	"io"
	"strings"
)

const (
	MagicString = "DICM"
)

type DcmReader interface {
	io.Reader
	readUInt8() (uint8, error)
	readUInt16() (uint16, error)
	readUInt32() (uint32, error)
	readUInt64() (uint64, error)
	readInt8() (int8, error)
	readInt16() (int16, error)
	readInt32() (int32, error)
	readInt64() (int64, error)
	readFloat32() (float32, error)
	readFloat64() (float64, error)
	IsImplicit() bool
	isTrackingImplicit() bool
	peek(n int) ([]byte, error)
	discard(n int) (int, error)
	skip(n int64) error
	setTransferSyntax(binaryOrder binary.ByteOrder, isImplicit bool)
	setOverallImplicit(isImplicit bool)
	readString(n uint32) (string, error)
	SkipPixelData() bool
	ByteOrder() binary.ByteOrder
	SetFileSize(fileSize int64) error
	GetFileSize() int64
	Parse() error
	GetMetadata() Dataset
	GetDataset() Dataset
	RetrieveFileUID() (*DicomUID, error)
	GetElementByTagString(tagStr string) (interface{}, error)
	ExportDatasetTags(exportMeta bool) MappedTag
	ExportSeriesTags() MappedTag
}

type dcmReader struct {
	reader               *bufio.Reader
	binaryOrder          binary.ByteOrder
	dataset              Dataset
	metadata             Dataset
	allowNonCompliantDcm bool
	isImplicit           bool
	keepTrackImplicit    bool
	skipPixelData        bool
	skipDataset          bool
	fileSize             int64
}

// NewDICOMReader returns a new reader
func NewDICOMReader(reader *bufio.Reader, options ...func(*dcmReader)) DcmReader {
	parser := &dcmReader{
		reader:        reader,
		binaryOrder:   binary.LittleEndian,
		isImplicit:    false,
		skipPixelData: false,
		skipDataset:   false,
	}
	for _, opt := range options {
		opt(parser)
	}
	return parser
}

// WithAllowNonCompliantDcm provides option to keep trying to parse the file even if it's not DICOM compliant
// e.g.: Missing header, missing FileMetaInformationGroupLength,...
func WithAllowNonCompliantDcm(allowNonCompliantDcm bool) func(*dcmReader) {
	return func(s *dcmReader) {
		s.allowNonCompliantDcm = allowNonCompliantDcm
	}
}

// WithSkipPixelData provides option to skip reading pixel data (7FE0,0010).
// If true, pixel data is skipped. If false, pixel data will be read
func WithSkipPixelData(skipPixelData bool) func(*dcmReader) {
	return func(s *dcmReader) {
		s.skipPixelData = skipPixelData
	}
}

// WithSkipDataset provides option to read only the metadata header.
// If true, only the meta header is read, else, the dataset will be read
func WithSkipDataset(skipDataset bool) func(*dcmReader) {
	return func(s *dcmReader) {
		s.skipDataset = skipDataset
	}
}

// WithSetFileSize provides option to set the file size to the reader
func WithSetFileSize(fileSize int64) func(*dcmReader) {
	return func(s *dcmReader) {
		s.fileSize = fileSize
	}
}

//----------------------------------------------------------------------------------------------------------------------
// Exported Methods
//----------------------------------------------------------------------------------------------------------------------

// GetMetadata returns the file meta header
func (r *dcmReader) GetMetadata() Dataset {
	return r.metadata
}

// GetDataset returns the dataset
func (r *dcmReader) GetDataset() Dataset {
	return r.dataset
}

func (r *dcmReader) RetrieveFileUID() (*DicomUID, error) {
	return r.dataset.RetrieveFileUID()
}

// Parse reads the DICOM file and parses it into array of elements
func (r *dcmReader) Parse() error {
	err := r.parse()
	if err != nil {
		return err
	}
	return nil
}

func (r *dcmReader) SkipPixelData() bool {
	return r.skipPixelData
}

func (r *dcmReader) ByteOrder() binary.ByteOrder {
	return r.binaryOrder
}

// IsValidDICOM checks if the dicom file follows the standard by having 128 bytes preamble followed by the magic string 'DICM'
func (r *dcmReader) IsValidDICOM() error {
	preamble, err := r.peek(128 + 4)
	if err != nil {
		return fmt.Errorf("cannot read the first 132 bytes: %v", err)
	}
	if string(preamble[128:]) != MagicString {
		return fmt.Errorf("file is not in valid dicom format")
	}
	return nil
}

// GetElementByTagString returns the element value of the input tag
// Tag should be in (gggg,eeee) or ggggeeee format
func (r *dcmReader) GetElementByTagString(tagStr string) (interface{}, error) {
	tagStr = utils.FormatTag(tagStr)

	if strings.HasPrefix(tagStr, "0002") {
		for _, elem := range r.metadata.Elements {
			if tagStr == elem.Tag.StringWithoutParentheses() {
				return elem.Value, nil
			}
		}
		return nil, fmt.Errorf("cannot find tag %s", tagStr)
	} else {
		for _, elem := range r.dataset.Elements {
			if tagStr == elem.Tag.StringWithoutParentheses() {
				return elem.Value, nil
			}
		}
		return nil, fmt.Errorf("cannot find tag %s", tagStr)
	}
}

//----------------------------------------------------------------------------------------------------------------------
// Unexported Methods
//----------------------------------------------------------------------------------------------------------------------

func (r *dcmReader) isTrackingImplicit() bool {
	return r.keepTrackImplicit
}

func (r *dcmReader) setOverallImplicit(isImplicit bool) {
	r.keepTrackImplicit = isImplicit
}

func (r *dcmReader) Read(p []byte) (int, error) {
	return r.reader.Read(p)
}

func (r *dcmReader) readUInt8() (uint8, error) {
	var res uint8

	err := binary.Read(r, r.binaryOrder, &res)

	return res, err
}

func (r *dcmReader) readUInt16() (uint16, error) {
	var res uint16
	err := binary.Read(r, r.binaryOrder, &res)
	return res, err
}

func (r *dcmReader) readUInt32() (uint32, error) {
	var res uint32
	err := binary.Read(r, r.binaryOrder, &res)
	return res, err
}

func (r *dcmReader) readUInt64() (uint64, error) {
	var res uint64
	err := binary.Read(r, r.binaryOrder, &res)
	return res, err
}

func (r *dcmReader) readInt8() (int8, error) {
	var res int8
	err := binary.Read(r, r.binaryOrder, &res)
	return res, err
}

func (r *dcmReader) readInt16() (int16, error) {
	var res int16
	err := binary.Read(r, r.binaryOrder, &res)
	return res, err
}

func (r *dcmReader) readInt32() (int32, error) {
	var res int32
	err := binary.Read(r, r.binaryOrder, &res)
	return res, err
}

func (r *dcmReader) readInt64() (int64, error) {
	var res int64
	err := binary.Read(r, r.binaryOrder, &res)
	return res, err
}

func (r *dcmReader) readFloat32() (float32, error) {
	var res float32
	err := binary.Read(r, r.binaryOrder, &res)
	return res, err
}

func (r *dcmReader) readFloat64() (float64, error) {
	var res float64
	err := binary.Read(r, r.binaryOrder, &res)
	return res, err
}

func (r *dcmReader) IsImplicit() bool {
	return r.isImplicit
}

func (r *dcmReader) peek(n int) ([]byte, error) {
	return r.reader.Peek(n)
}

func (r *dcmReader) discard(n int) (int, error) {
	return r.reader.Discard(n)
}

func (r *dcmReader) skip(n int64) error {

	_, err := io.CopyN(io.Discard, r, n)

	return err
}

func (r *dcmReader) setTransferSyntax(binaryOrder binary.ByteOrder, isImplicit bool) {
	r.binaryOrder = binaryOrder
	r.isImplicit = isImplicit
}

func (r *dcmReader) SetFileSize(fileSize int64) error {
	r.fileSize = fileSize
	return nil
}

func (r *dcmReader) GetFileSize() int64 {
	return r.fileSize
}

func (r *dcmReader) readString(n uint32) (string, error) {
	data := make([]byte, n)
	_, err := io.ReadFull(r, data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (r *dcmReader) parse() error {
	r.setFileSize()
	err := r.IsValidDICOM()
	if err != nil {
		return err
	}
	_ = r.skip(132)
	err = r.parseMetadata()
	if err != nil {
		return err
	}

	// IMPORTANT: Additional check is needed here since there are few instances where the DICOM
	// meta header is registered as Explicit Little-Endian, but Implicit Little-Endian is used in the body
	err = r.checkImplicityAgreement()
	if err != nil {
		return nil
	}

	if r.skipDataset {
		return nil
	}
	err = r.parseDataset()
	if err != nil {
		return err
	}
	return nil
}

// setFileSize sets the file size to the reader
func (r *dcmReader) setFileSize() {
	_ = r.SetFileSize(r.fileSize)
}

// parseMetadata parses the file meta information according to
// https://dicom.nema.org/dicom/2013/output/chtml/part10/chapter_7.html
// the File Meta Information shall be encoded using the Explicit VR Little Endian Transfer Syntax
// (UID=1.2.840.10008.1.2.1)
func (r *dcmReader) parseMetadata() error {
	var metadata []*Element
	var transferSyntaxUID string

	for {
		// No longer relied on the MetaInformationGroupLength tag to determine the length of the meta header.
		// We check if the group tag is 0x0002 before proceeding to read the element. If the group tag is not 0x0002,
		// then break the loop
		n, err := r.peek(2)
		if err != nil {
			return err
		}
		if bytes.Compare(n, []byte{0x2, 0x0}) != 0 {
			break
		}

		res, err := ReadElement(r, r.IsImplicit(), r.ByteOrder())
		if err != nil {
			return err
		}
		metadata = append(metadata, res)
		if res.Tag == tag.TransferSyntaxUID {
			transferSyntaxUID = (res.Value.RawValue).(string)
		}
	}
	r.metadata = Dataset{Elements: metadata}

	// Set transfer syntax here for the dataset parser
	binOrder, isImplicit, err := uid.ParseTransferSyntaxUID(transferSyntaxUID)
	if err != nil {
		return err
	}
	r.setTransferSyntax(binOrder, isImplicit)
	r.setOverallImplicit(isImplicit)

	return nil
}

func (r *dcmReader) checkImplicityAgreement() error {
	// Need to check if the implicit matches between header and body
	n, err := r.peek(6)
	if err != nil {
		return err
	}
	if !vr.VRMapper[string(n[4:6])] && !r.IsImplicit() {
		r.setTransferSyntax(r.binaryOrder, true)
	}
	if vr.VRMapper[string(n[4:6])] && r.IsImplicit() {
		r.setTransferSyntax(r.binaryOrder, false)
	}
	return nil
}

// parseDataset parses the file dataset after the file meta header
func (r *dcmReader) parseDataset() error {
	var data []*Element
	for {
		res, err := ReadElement(r, r.IsImplicit(), r.ByteOrder())
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
	//r.dataset.Elements = append(r.dataset.Elements, dicomDataset.Elements...)
	r.dataset = dicomDataset
	return nil
}
