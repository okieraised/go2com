package reader

import (
	"bufio"
	"encoding/binary"
	_ "github.com/okieraised/go2com/internal/system"
	"io"
)

type DcmReader interface {
	io.Reader
	ReadUInt8() (uint8, error)
	ReadUInt16() (uint16, error)
	ReadUInt32() (uint32, error)
	ReadUInt64() (uint64, error)
	ReadInt8() (int8, error)
	ReadInt16() (int16, error)
	ReadInt32() (int32, error)
	ReadInt64() (int64, error)
	ReadFloat32() (float32, error)
	ReadFloat64() (float64, error)
	IsImplicit() bool
	IsTrackingImplicit() bool
	SkipPixelData() bool
	Peek(n int) ([]byte, error)
	Discard(n int) (int, error)
	ByteOrder() binary.ByteOrder
	Skip(n int64) error
	SetTransferSyntax(binaryOrder binary.ByteOrder, isImplicit bool)
	SetOverallImplicit(isImplicit bool)
	SetFileSize(fileSize int64) error
	GetFileSize() int64
	ReadString(n uint32) (string, error)
}

type dcmReader struct {
	reader            *bufio.Reader
	binaryOrder       binary.ByteOrder
	isImplicit        bool
	keepTrackImplicit bool
	skipPixelData     bool
	fileSize          int64
}

// NewDICOMReader returns a new reader
func NewDICOMReader(reader *bufio.Reader, options ...func(*dcmReader)) DcmReader {
	parser := &dcmReader{
		reader:        reader,
		binaryOrder:   binary.LittleEndian,
		isImplicit:    false,
		skipPixelData: false,
	}
	for _, opt := range options {
		opt(parser)
	}
	return parser
}

// WithSkipPixelData provides option to skip reading pixel data (7FE0,0010).
// If true, pixel data is skipped. If false, pixel data will be read
func WithSkipPixelData(skipPixelData bool) func(*dcmReader) {
	return func(s *dcmReader) {
		s.skipPixelData = skipPixelData
	}
}

func (r *dcmReader) IsTrackingImplicit() bool {
	return r.keepTrackImplicit
}

func (r *dcmReader) SetOverallImplicit(isImplicit bool) {
	r.keepTrackImplicit = isImplicit
}

func (r *dcmReader) Read(p []byte) (int, error) {
	return r.reader.Read(p)
}

func (r *dcmReader) ReadUInt8() (uint8, error) {
	var res uint8

	err := binary.Read(r, r.binaryOrder, &res)

	return res, err
}

func (r *dcmReader) ReadUInt16() (uint16, error) {
	var res uint16
	err := binary.Read(r, r.binaryOrder, &res)
	return res, err
}

func (r *dcmReader) ReadUInt32() (uint32, error) {
	var res uint32
	err := binary.Read(r, r.binaryOrder, &res)
	return res, err
}

func (r *dcmReader) ReadUInt64() (uint64, error) {
	var res uint64
	err := binary.Read(r, r.binaryOrder, &res)
	return res, err
}

func (r *dcmReader) ReadInt8() (int8, error) {
	var res int8
	err := binary.Read(r, r.binaryOrder, &res)
	return res, err
}

func (r *dcmReader) ReadInt16() (int16, error) {
	var res int16
	err := binary.Read(r, r.binaryOrder, &res)
	return res, err
}

func (r *dcmReader) ReadInt32() (int32, error) {
	var res int32
	err := binary.Read(r, r.binaryOrder, &res)
	return res, err
}

func (r *dcmReader) ReadInt64() (int64, error) {
	var res int64
	err := binary.Read(r, r.binaryOrder, &res)
	return res, err
}

func (r *dcmReader) ReadFloat32() (float32, error) {
	var res float32
	err := binary.Read(r, r.binaryOrder, &res)
	return res, err
}

func (r *dcmReader) ReadFloat64() (float64, error) {
	var res float64
	err := binary.Read(r, r.binaryOrder, &res)
	return res, err
}

func (r *dcmReader) IsImplicit() bool {
	return r.isImplicit
}

func (r *dcmReader) SkipPixelData() bool {
	return r.skipPixelData
}

func (r *dcmReader) Peek(n int) ([]byte, error) {
	return r.reader.Peek(n)
}

func (r *dcmReader) Discard(n int) (int, error) {
	return r.reader.Discard(n)
}

func (r *dcmReader) ByteOrder() binary.ByteOrder {
	return r.binaryOrder
}

func (r *dcmReader) Skip(n int64) error {

	_, err := io.CopyN(io.Discard, r, n)

	return err
}

func (r *dcmReader) SetTransferSyntax(binaryOrder binary.ByteOrder, isImplicit bool) {
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

func (r *dcmReader) ReadString(n uint32) (string, error) {
	data := make([]byte, n)
	_, err := io.ReadFull(r, data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
