package nifti1

// #include "./nifti1.h"
import "C"
import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/okieraised/go2com/internal/constants"
	"github.com/okieraised/go2com/internal/utils"
	"io"
	"net/http"
	"os"
)

type Nii1Reader interface {
	Parse() error
	GetBinaryOrder() binary.ByteOrder
	GetNiiData() *Nii1
	GetUnitsOfMeasurements() ([2]string, error)
	GetAffine() [4][4]float32
	GetImgShape() [4]int16
	GetAt(x, y, z, t int32) float32
	GetTimeSeries(x, y, z int32) ([]float32, error)
	GetSlice(z, t int32) ([][]float32, error)
	GetHeader() *Nii1Header
}

type nii1Reader struct {
	reader      *bytes.Reader
	binaryOrder binary.ByteOrder
	niiData     *Nii1
}

func NewNii1Reader(filePath string) (Nii1Reader, error) {
	bData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	mimeType := http.DetectContentType(bData[:512])
	if mimeType == "application/x-gzip" {
		bData, err = utils.DeflateGzip(bData)
		if err != nil {
			return nil, err
		}
	}
	return &nii1Reader{
		binaryOrder: binary.LittleEndian,
		reader:      bytes.NewReader(bData),
		niiData:     &Nii1{},
	}, nil
}

func (r *nii1Reader) GetHeader() *Nii1Header {
	return r.niiData.Header
}

func (r *nii1Reader) GetSlice(z, t int32) ([][]float32, error) {
	return r.niiData.getSlice(z, t)
}

func (r *nii1Reader) GetTimeSeries(x, y, z int32) ([]float32, error) {
	return r.niiData.getTimeSeries(x, y, z)
}

func (r *nii1Reader) GetAt(x, y, z, t int32) float32 {
	return r.niiData.getAt(x, y, z, t)
}

func (r *nii1Reader) GetUnitsOfMeasurements() ([2]string, error) {
	return r.niiData.getUnitsOfMeasurements()
}

func (r *nii1Reader) GetAffine() [4][4]float32 {
	return r.niiData.getAffine()
}

func (r *nii1Reader) GetImgShape() [4]int16 {
	return r.niiData.getImgShape()
}

func (r *nii1Reader) GetNiiData() *Nii1 {
	return r.niiData
}

func (r *nii1Reader) GetBinaryOrder() binary.ByteOrder {
	return r.binaryOrder
}

func (r *nii1Reader) checkNiiVersion() error {
	var hSize int32

	err := binary.Read(r.reader, r.binaryOrder, &hSize)
	if err != nil {
		return err
	}

	switch hSize {
	case constants.NII1HeaderSize:
		return nil
	case constants.NII2HeaderSize:
		return errors.New("file is of NIFTI-2 format")
	default:
		r.binaryOrder = binary.BigEndian
		_, err := r.reader.Seek(0, 0)
		if err != nil {
			return err
		}

		var hSize int32

		err = binary.Read(r.reader, r.binaryOrder, &hSize)
		if err != nil {
			return err
		}
		switch hSize {
		case constants.NII1HeaderSize:
			return nil
		case constants.NII2HeaderSize:
			return errors.New("file is of NIFTI-2 format")
		default:
			return errors.New("invalid NIFTI file format")
		}
	}
}

// Parse returns the raw byte array into NIFTI-1 header and dataset structure
func (r *nii1Reader) Parse() error {
	err := r.checkNiiVersion()
	if err != nil {
		return err
	}

	err = r.parseHeader()
	if err != nil {
		return err
	}

	err = r.parseData()
	if err != nil {
		return err
	}

	return nil
}

// parseHeader parses the raw byte array into NIFTI-1 header structure
func (r *nii1Reader) parseHeader() error {
	_, err := r.reader.Seek(0, 0)
	if err != nil {
		return err
	}
	header := new(Nii1Header)

	err = binary.Read(r.reader, r.binaryOrder, header)
	if err != nil {
		return err
	}
	if header.Magic != [4]byte{110, 43, 49, 0} && header.Magic != [4]byte{110, 105, 49, 0} {
		return errors.New("invalid NIFTI magic string")
	}
	if header.Datatype == C.DT_BINARY || header.Datatype == C.DT_UNKNOWN {
		return errors.New("data type is invalid")
	}

	r.niiData.Header = header

	return nil
}

// parseData parse the raw byte array into NIFTI-1 data structure
func (r *nii1Reader) parseData() error {
	var offset int
	statDim := 1
	image := new(Nii1Data)

	header := r.niiData.Header

	if header.Bitpix == 0 {
		return errors.New("number of bits per voxel value (bitpix) is zero")
	}

	image.NDim, image.Dim[0] = int32(header.Dim[0]), int32(header.Dim[0])
	image.Nx, image.Dim[1] = int32(header.Dim[1]), int32(header.Dim[1])
	image.Ny, image.Dim[2] = int32(header.Dim[2]), int32(header.Dim[2])
	image.Nz, image.Dim[3] = int32(header.Dim[3]), int32(header.Dim[3])
	image.Nt, image.Dim[4] = int32(header.Dim[4]), int32(header.Dim[4])
	image.Nu, image.Dim[5] = int32(header.Dim[5]), int32(header.Dim[5])
	image.Nv, image.Dim[6] = int32(header.Dim[6]), int32(header.Dim[6])
	image.Nw, image.Dim[7] = int32(header.Dim[7]), int32(header.Dim[7])
	image.Dx, image.PixDim[1] = float64(header.Pixdim[1]), float64(header.Pixdim[1])
	image.Dx, image.PixDim[2] = float64(header.Pixdim[2]), float64(header.Pixdim[2])
	image.Dx, image.PixDim[3] = float64(header.Pixdim[3]), float64(header.Pixdim[3])
	image.Dx, image.PixDim[4] = float64(header.Pixdim[4]), float64(header.Pixdim[4])
	image.Dx, image.PixDim[5] = float64(header.Pixdim[5]), float64(header.Pixdim[5])
	image.Dx, image.PixDim[6] = float64(header.Pixdim[6]), float64(header.Pixdim[6])
	image.Dx, image.PixDim[7] = float64(header.Pixdim[7]), float64(header.Pixdim[7])
	image.IntentName = header.IntentName
	image.IntentP1 = header.IntentP1
	image.IntentP2 = header.IntentP2
	image.IntentP3 = header.IntentP3
	image.QformCode = int32(header.QformCode)
	image.SformCode = int32(header.SformCode)
	image.SclSlope = header.SclSlope
	image.SclInter = header.SclInter
	image.QuaternB = header.QuaternB
	image.QuaternC = header.QuaternC
	image.QuaternD = header.QuaternD
	image.Descrip = header.Descrip

	image.NVox = 1
	for i := int16(1); i <= header.Dim[0]; i++ {
		image.NVox *= int32(header.Dim[i])
	}
	image.NByPer = int32(header.Bitpix) / 8
	image.ByteOrder = r.binaryOrder

	if image.Dim[5] > 1 {
		statDim = int(image.Dim[5])
	}
	offset = int(header.VoxOffset)
	dataSize := image.Dim[1] * image.Dim[2] * image.Dim[3] * image.Dim[4] * int32(statDim) * (int32(header.Bitpix) / 8)

	_, err := r.reader.Seek(int64(offset), 0)
	if err != nil {
		return err
	}

	buf := make([]byte, dataSize)
	_, err = io.ReadFull(r.reader, buf)
	if err != nil {
		return err
	}
	image.Data = buf
	r.niiData.Data = image

	return nil
}
