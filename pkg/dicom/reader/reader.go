package reader

import (
	"bufio"
	"encoding/binary"
	"io"

	"github.com/okieraised/go2com/internal/system"
	_ "github.com/okieraised/go2com/internal/system"
)

type DcmReader interface {
}

type dcmReader struct {
	reader        *bufio.Reader
	currentOffset int64
	binaryOrder   binary.ByteOrder
	isImplicit    bool
	readPixel     bool
}

func NewDcmReader(reader *bufio.Reader, readPixel bool) DcmReader {
	return &dcmReader{
		reader:        reader,
		currentOffset: 0,
		binaryOrder:   system.NativeEndian,
		isImplicit:    false,
		readPixel:     readPixel,
	}
}

func (r *dcmReader) Read(p []byte) (int, error) {
	r.reader.Read(p)
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
