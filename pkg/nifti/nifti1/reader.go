package nifti1

// #include "./nifti1.h"
import "C"
import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/okieraised/go2com/internal/constants"
	"github.com/okieraised/go2com/internal/utils"
	"github.com/okieraised/go2com/pkg/matrix"
	"github.com/okieraised/go2com/pkg/nifti/constant"
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
	GetDatatype() string
	GetOrientation() [3]string
	GetSliceCode() string
	QuaternToMatrix() matrix.DMat44
	MatrixToOrientation(R matrix.DMat44)
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

func (r *nii1Reader) MatrixToOrientation(R matrix.DMat44) {
	r.niiData.matrixToOrientation(R)
}

func (r *nii1Reader) QuaternToMatrix() matrix.DMat44 {
	return r.niiData.quaternToMatrix()
}

func (r *nii1Reader) GetSliceCode() string {
	return r.niiData.getSliceCode()
}

func (r *nii1Reader) GetOrientation() [3]string {
	return r.niiData.getOrientation()
}

func (r *nii1Reader) GetDatatype() string {
	return r.niiData.getDatatype()
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

	if header.Dim[0] < 0 || header.Dim[0] > 7 {
		if r.binaryOrder == binary.LittleEndian {
			r.binaryOrder = binary.BigEndian
		} else {
			r.binaryOrder = binary.LittleEndian
		}
	}

	return nil
}

// setDatatypeSize sets number of bytes per voxel and the swapsize for the header datatype
func (r *nii1Reader) setDatatypeSize() {
	var NByPerVoxel int32 = 0
	var SwapSize int32 = 0

	switch r.niiData.Header.Datatype {
	case constant.DT_INT8, constant.DT_UINT8:
		NByPerVoxel = 1
		SwapSize = 0
	case constant.DT_INT16, constant.DT_UINT16:
		NByPerVoxel = 2
		SwapSize = 2
	case constant.DT_RGB24:
		NByPerVoxel = 3
		SwapSize = 0
	case constant.DT_RGBA32:
		NByPerVoxel = 4
		SwapSize = 0
	case constant.DT_INT32, constant.DT_UINT32, constant.DT_FLOAT32:
		NByPerVoxel = 4
		SwapSize = 4
	case constant.DT_COMPLEX64:
		NByPerVoxel = 8
		SwapSize = 4
	case constant.DT_FLOAT64, constant.DT_INT64, constant.DT_UINT64:
		NByPerVoxel = 8
		SwapSize = 8
	case constant.DT_FLOAT128:
		NByPerVoxel = 16
		SwapSize = 16
	case constant.DT_COMPLEX128:
		NByPerVoxel = 16
		SwapSize = 8
	case constant.DT_COMPLEX256:
		NByPerVoxel = 32
		SwapSize = 16
	}
	r.niiData.Data.NByPer = NByPerVoxel
	r.niiData.Data.SwapSize = SwapSize
}

// parseData parse the raw byte array into NIFTI-1 data structure
func (r *nii1Reader) parseData() error {
	r.niiData.Data = &Nii1Data{}
	var offset int
	statDim := 1

	r.setDatatypeSize()

	header := r.niiData.Header

	if header.Bitpix == 0 {
		return errors.New("number of bits per voxel value (bitpix) is zero")
	}

	r.niiData.Data.NDim, r.niiData.Data.Dim[0] = int32(header.Dim[0]), int32(header.Dim[0])
	r.niiData.Data.Nx, r.niiData.Data.Dim[1] = int32(header.Dim[1]), int32(header.Dim[1])
	r.niiData.Data.Ny, r.niiData.Data.Dim[2] = int32(header.Dim[2]), int32(header.Dim[2])
	r.niiData.Data.Nz, r.niiData.Data.Dim[3] = int32(header.Dim[3]), int32(header.Dim[3])
	r.niiData.Data.Nt, r.niiData.Data.Dim[4] = int32(header.Dim[4]), int32(header.Dim[4])
	r.niiData.Data.Nu, r.niiData.Data.Dim[5] = int32(header.Dim[5]), int32(header.Dim[5])
	r.niiData.Data.Nv, r.niiData.Data.Dim[6] = int32(header.Dim[6]), int32(header.Dim[6])
	r.niiData.Data.Nw, r.niiData.Data.Dim[7] = int32(header.Dim[7]), int32(header.Dim[7])
	r.niiData.Data.Dx, r.niiData.Data.PixDim[1] = float64(header.Pixdim[1]), float64(header.Pixdim[1])
	r.niiData.Data.Dx, r.niiData.Data.PixDim[2] = float64(header.Pixdim[2]), float64(header.Pixdim[2])
	r.niiData.Data.Dx, r.niiData.Data.PixDim[3] = float64(header.Pixdim[3]), float64(header.Pixdim[3])
	r.niiData.Data.Dx, r.niiData.Data.PixDim[4] = float64(header.Pixdim[4]), float64(header.Pixdim[4])
	r.niiData.Data.Dx, r.niiData.Data.PixDim[5] = float64(header.Pixdim[5]), float64(header.Pixdim[5])
	r.niiData.Data.Dx, r.niiData.Data.PixDim[6] = float64(header.Pixdim[6]), float64(header.Pixdim[6])
	r.niiData.Data.Dx, r.niiData.Data.PixDim[7] = float64(header.Pixdim[7]), float64(header.Pixdim[7])

	r.niiData.Data.IntentName = header.IntentName
	r.niiData.Data.IntentCode = int32(header.IntentCode)
	r.niiData.Data.SliceCode = int32(header.SliceCode)

	r.niiData.Data.IntentP1 = header.IntentP1
	r.niiData.Data.IntentP2 = header.IntentP2
	r.niiData.Data.IntentP3 = header.IntentP3
	r.niiData.Data.QformCode = int32(header.QformCode)
	r.niiData.Data.SformCode = int32(header.SformCode)
	r.niiData.Data.SclSlope = header.SclSlope
	r.niiData.Data.SclInter = header.SclInter
	r.niiData.Data.QuaternB = header.QuaternB
	r.niiData.Data.QuaternC = header.QuaternC
	r.niiData.Data.QuaternD = header.QuaternD
	r.niiData.Data.Descrip = header.Descrip

	r.niiData.Data.QFac = float64(header.Pixdim[0])

	r.niiData.Data.NVox = 1
	for i := int16(1); i <= header.Dim[0]; i++ {
		r.niiData.Data.NVox *= int32(header.Dim[i])
	}

	r.niiData.Data.Datatype = int32(r.niiData.Header.Datatype)
	r.niiData.Data.ByteOrder = r.binaryOrder

	if r.niiData.Data.Dim[5] > 1 {
		statDim = int(r.niiData.Data.Dim[5])
	}
	offset = int(header.VoxOffset)
	dataSize := r.niiData.Data.Dim[1] * r.niiData.Data.Dim[2] * r.niiData.Data.Dim[3] * r.niiData.Data.Dim[4] * int32(statDim) * (int32(header.Bitpix) / 8)

	_, err := r.reader.Seek(int64(offset), 0)
	if err != nil {
		return err
	}

	buf := make([]byte, dataSize)
	_, err = io.ReadFull(r.reader, buf)
	if err != nil {
		return err
	}
	r.niiData.Data.Data = buf

	return nil
}
