package nifti1

// #include "./nifti1.h"
import "C"
import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/okieraised/go2com/pkg/nifti/constant"
	"github.com/okieraised/go2com/pkg/nifti/matrix"
	"io"
	"math"
)

type Nii1 struct {
	Header *Nii1Header
	Data   *Nii1Data
}

// Nii1Header defines the structure of the NIFTI-1 header
type Nii1Header struct {
	SizeofHdr      int32
	DataTypeUnused [10]uint8
	DbName         [18]uint8
	Extents        int32
	SessionError   int16
	Regular        uint8
	DimInfo        uint8
	Dim            [8]int16
	IntentP1       float32
	IntentP2       float32
	IntentP3       float32
	IntentCode     int16
	Datatype       int16
	Bitpix         int16
	SliceStart     int16
	Pixdim         [8]float32
	VoxOffset      float32
	SclSlope       float32
	SclInter       float32
	SliceEnd       int16
	SliceCode      uint8
	XyztUnits      uint8
	CalMax         float32
	CalMin         float32
	SliceDuration  float32
	Toffset        float32
	Glmax          int32
	Glmin          int32
	Descrip        [80]uint8
	AuxFile        [24]uint8
	QformCode      int16
	SformCode      int16
	QuaternB       float32
	QuaternC       float32
	QuaternD       float32
	QoffsetX       float32
	QoffsetY       float32
	QoffsetZ       float32
	SrowX          [4]float32
	SrowY          [4]float32
	SrowZ          [4]float32
	IntentName     [16]uint8
	Magic          [4]uint8
}

type Nii1Data struct {
	NDim          int              // last dimension greater than 1 (1..7)
	Nx            int              // dimensions of grid array
	Ny            int              // dimensions of grid array
	Nz            int              // dimensions of grid array
	Nt            int              // dimensions of grid array
	Nu            int              // dimensions of grid array
	Nv            int              // dimensions of grid array
	Nw            int              // dimensions of grid array
	Dim           [8]int           // dim[0] = ndim, dim[1] = nx, etc
	NVox          int              // number of voxels = nx*ny*nz*...*nw
	NByPer        int              // bytes per voxel, matches datatype
	DataType      int              // type of data in voxels: DT_* code
	Dx            float64          // grid spacings
	Dy            float64          // grid spacings
	Dz            float64          // grid spacings
	Dt            float64          // grid spacings
	Du            float64          // grid spacings
	Dv            float64          // grid spacings
	Dw            float64          // grid spacings tEStataILSTERIOn
	PixDim        [8]float64       // pixdim[1]=dx, etc
	SclSlope      float32          // scaling parameter: slope
	SclInter      float32          // scaling parameter: intercept
	CalMin        float64          // calibration parameter: minimum
	CalMax        float64          // calibration parameter: maximum
	QformCode     int32            // codes for (x,y,z) space meaning
	SformCode     int32            // codes for (x,y,z) space meaning
	FreqDim       int32            // indices (1,2,3, or 0) for MRI
	PhaseDim      int32            // directions in dim[]/pixdim[]
	SliceDim      int32            // directions in dim[]/pixdim[]
	SliceCode     int32            // code for slice timing pattern
	SliceStart    int32            // index for start of slices
	SliceEnd      int32            // index for end of slices
	SliceDuration float64          // time between individual slices
	QuaternB      float32          // Quaternion transform parameters [when writing a dataset, these are used for qform, NOT qto_xyz]
	QuaternC      float32          // Quaternion transform parameters [when writing a dataset, these are used for qform, NOT qto_xyz]
	QuaternD      float32          // Quaternion transform parameters [when writing a dataset, these are used for qform, NOT qto_xyz]
	QOffsetX      float32          // Quaternion transform parameters [when writing a dataset, these are used for qform, NOT qto_xyz]
	QOffsetY      float32          // Quaternion transform parameters [when writing a dataset, these are used for qform, NOT qto_xyz]
	QOffsetZ      float32          // Quaternion transform parameters [when writing a dataset, these are used for qform, NOT qto_xyz]
	QFac          float64          // Quaternion transform parameters [when writing a dataset, these are used for qform, NOT qto_xyz]
	QtoXYZ        matrix.FMat44    // qform: transform (i,j,k) to (x,y,z)
	QtoIJK        matrix.FMat44    // qform: transform (x,y,z) to (i,j,k)
	StoXYZ        matrix.FMat44    // sform: transform (i,j,k) to (x,y,z)
	StoIJK        matrix.FMat44    // sform: transform (x,y,z) to (i,j,k)
	TOffset       float64          // time coordinate offset
	XYZUnits      int32            // dx,dy,dz units: NIFTI_UNITS_* code
	TimeUnits     int32            // dt units: NIFTI_UNITS_* code
	NiftiType     int32            // 0==Analyze, 1==NIFTI-1 (file), 2==NIFTI-1 (2 files), 3==NIFTI-ASCII (1 file)
	IntentCode    int32            // statistic type (or something)
	IntentP1      float32          // intent parameters
	IntentP2      float32          // intent parameters
	IntentP3      float32          // intent parameters
	IntentName    [16]byte         // optional description of intent data
	Descrip       [80]byte         // optional text to describe dataset
	AuxFile       [24]byte         // auxiliary filename
	FName         *byte            // header filename
	IName         *byte            // image filename
	INameOffset   int32            // offset into IName where data start
	SwapSize      int32            // swap unit in image data (might be 0)
	ByteOrder     binary.ByteOrder // byte order on disk (MSB_ or LSB_FIRST)
	Data          []byte           // slice of data: nbyper*nvox bytes
	NumExt        int32            // number of extensions in extList
	Nifti1Ext     []interface{}    // array of extension structs (with data)
}

func (n *Nii1) Parse(reader *bytes.Reader, bo binary.ByteOrder) error {
	err := n.parseHeader(reader, bo)
	if err != nil {
		return err
	}

	err = n.parseData(reader, bo)
	if err != nil {
		return err
	}

	return nil
}

func (n *Nii1) parseHeader(reader *bytes.Reader, bo binary.ByteOrder) error {
	header := new(Nii1Header)
	err := binary.Read(reader, bo, header)
	if err != nil {
		return err
	}
	if header.Magic != [4]byte{110, 43, 49, 0} && header.Magic != [4]byte{110, 105, 49, 0} {
		return errors.New("invalid NIFTI magic string")
	}
	if header.Datatype == C.DT_BINARY || header.Datatype == C.DT_UNKNOWN {
		return errors.New("data type is invalid")
	}

	n.Header = header

	return nil
}

func (n *Nii1) parseData(reader *bytes.Reader, bo binary.ByteOrder) error {
	var offset int
	statDim := 1
	image := new(Nii1Data)

	header := n.Header

	if header.Bitpix == 0 {
		return errors.New("number of bits per voxel value (bitpix) is zero")
	}

	image.NDim, image.Dim[0] = int(header.Dim[0]), int(header.Dim[0])
	image.Nx, image.Dim[1] = int(header.Dim[1]), int(header.Dim[1])
	image.Ny, image.Dim[2] = int(header.Dim[2]), int(header.Dim[2])
	image.Nz, image.Dim[3] = int(header.Dim[3]), int(header.Dim[3])
	image.Nt, image.Dim[4] = int(header.Dim[4]), int(header.Dim[4])
	image.Nu, image.Dim[5] = int(header.Dim[5]), int(header.Dim[5])
	image.Nv, image.Dim[6] = int(header.Dim[6]), int(header.Dim[6])
	image.Nw, image.Dim[7] = int(header.Dim[7]), int(header.Dim[7])
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
		image.NVox *= int(header.Dim[i])
	}
	image.NByPer = int(header.Bitpix) / 8
	image.ByteOrder = bo

	image.NByPer = int(header.Bitpix) / 8

	if image.Dim[5] > 1 {
		statDim = image.Dim[5]
	}
	offset = int(header.VoxOffset)
	dataSize := image.Dim[1] * image.Dim[2] * image.Dim[3] * image.Dim[4] * statDim * (int(header.Bitpix) / 8)

	_, err := reader.Seek(int64(offset), 0)
	if err != nil {
		return err
	}

	buf := make([]byte, dataSize)
	_, err = io.ReadFull(reader, buf)
	if err != nil {
		return err
	}
	image.Data = buf

	fmt.Println(dataSize, offset, reader.Len(), len(buf))

	n.Data = image

	return nil
}

// GetHeader returns the NIFTI header
func (n *Nii1) GetHeader() interface{} {
	return n.Header
}

// GetImg returns the raw NIFTI image data
func (n *Nii1) GetImg() interface{} {
	return n.Data
}

func (n *Nii1) GetUnitsOfMeasurements() ([2]string, error) {
	// The bits 1-3 are used to store the spatial dimensions, the bits 4-6 are for temporal dimensions,
	// and the bits 6 and 7 are not used
	units := [2]string{}
	xyzUnit := n.Header.XyztUnits % 8
	tUnit := n.Header.XyztUnits - xyzUnit

	spatialUnit, ok := constant.NiiMeasurementUnits[xyzUnit]
	if !ok {
		return units, fmt.Errorf("invalid spatial unit %d", xyzUnit)
	}

	temporalUnit, ok := constant.NiiMeasurementUnits[tUnit]
	if !ok {
		return units, fmt.Errorf("invalid temporal unit %d", tUnit)
	}

	units[0] = spatialUnit
	units[1] = temporalUnit

	return units, nil
}

func (n *Nii1) GetAffine() interface{} {
	affine := [4][4]float32{}
	affine[0] = n.Header.SrowX
	affine[1] = n.Header.SrowY
	affine[2] = n.Header.SrowZ
	affine[3] = [4]float32{0, 0, 0, 1}
	return affine
}

func (n *Nii1) GetImgShape() interface{} {
	dim := [4]int16{}
	for index, _ := range dim {
		dim[index] = n.Header.Dim[index+1]
	}
	return dim
}

func (n *Nii1) GetVoxelSize() interface{} {
	size := [4]float32{}
	for index, _ := range size {
		size[index] = n.Header.Pixdim[index+1]
	}
	return size
}

func (n *Nii1) GetAt(x, y, z, t int) float32 {

	tIndex := t * n.Data.Nx * n.Data.Ny * n.Data.Nz
	zIndex := n.Data.Nx * n.Data.Ny * z
	yIndex := n.Data.Nx * y
	xIndex := x
	index := tIndex + zIndex + yIndex + xIndex

	dataPoint := n.Data.Data[index*n.Data.NByPer : (index+1)*n.Data.NByPer]

	var value float32
	switch n.Data.NByPer {
	case 1:
		if len(dataPoint) > 0 {
			value = float32(dataPoint[0])
		}
	case 2:
		v := binary.LittleEndian.Uint16(dataPoint)
		value = float32(v)
	case 4:
		v := binary.LittleEndian.Uint32(dataPoint)
		value = math.Float32frombits(v)
	case 8:
		v := binary.LittleEndian.Uint64(dataPoint)
		value = float32(math.Float64frombits(v))
	default:
	}

	if n.Data.SclSlope != 0 {
		value = n.Data.SclSlope*value + n.Data.SclInter
	}

	return value
}

func (n *Nii1) GetTimeSeries(x, y, z int) interface{} {
	var timeSeries []float32
	volumeN := n.Data.Dim[1] * n.Data.Dim[2] * n.Data.Dim[3]

	timeNum := len(n.Data.Data) / volumeN * n.Data.NByPer

	fmt.Println("timeNum", timeNum, n.Data.NByPer, volumeN)
	for i := 0; i < n.Data.Dim[4]; i++ {
		timeSeries = append(timeSeries, n.GetAt(x, y, z, i))
	}
	return timeSeries
}

func (n *Nii1) GetSlice(z, t int) interface{} {
	sliceX := n.Data.Nx
	sliceY := n.Data.Ny
	slice := make([][]float32, sliceX)
	for i := range slice {
		slice[i] = make([]float32, sliceY)
	}
	for xi := 0; xi < sliceX; xi++ {
		for yi := 0; yi < sliceY; yi++ {
			slice[xi][yi] = n.GetAt(xi, yi, z, t)
		}
	}
	return slice
}

//func (hdr *Nii1Header) GetIntentCode() (string, error) {
//	code := hdr.IntentCode
//	res, ok := constant.
//	if ok {
//		return res, nil
//	}
//	return "", errors.New("key not exist")
//}
//
//// GetSForm returns the coordinate specified in the header filed `SformCode`
//func (hdr *Nii1Header) GetSForm() (string, error) {
//	code := hdr.SformCode
//	res, ok := pkg.OrientationInfo[code]
//	if ok {
//		return res, nil
//	}
//	return "", errors.New("key not exist")
//}
