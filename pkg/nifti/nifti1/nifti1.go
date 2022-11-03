package nifti1

// #include "./nifti1.h"
import "C"
import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/okieraised/go2com/internal/constants"
	"github.com/okieraised/go2com/pkg/matrix"
	"github.com/okieraised/go2com/pkg/nifti/constant"
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
	NDim          int32            // last dimension greater than 1 (1..7)
	Nx            int32            // dimensions of grid array
	Ny            int32            // dimensions of grid array
	Nz            int32            // dimensions of grid array
	Nt            int32            // dimensions of grid array
	Nu            int32            // dimensions of grid array
	Nv            int32            // dimensions of grid array
	Nw            int32            // dimensions of grid array
	Dim           [8]int32         // dim[0] = ndim, dim[1] = nx, etc
	NVox          int32            // number of voxels = nx*ny*nz*...*nw
	NByPer        int32            // bytes per voxel, matches datatype (Datatype)
	Datatype      int32            // type of data in voxels: DT_* code
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
	Nifti1Ext     []Nifti1Ext      // array of extension structs (with data)
	IJKOrtient    [3]int32         // self-add. Orientation ini, j, k coordinate
}

type Nifti1Ext struct {
	ECode int32
	Edata []byte
	ESize int32
}

// getUnitsOfMeasurements returns the spatial and temporal units of measurements
func (n *Nii1) getUnitsOfMeasurements() ([2]string, error) {
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

// getAffine returns the 4x4 affine matrix
func (n *Nii1) getAffine() [4][4]float32 {
	affine := [4][4]float32{}
	affine[0] = n.Header.SrowX
	affine[1] = n.Header.SrowY
	affine[2] = n.Header.SrowZ
	affine[3] = [4]float32{0, 0, 0, 1}
	return affine
}

// getImgShape returns the image shape in terms of x, y, z, t
func (n *Nii1) getImgShape() [4]int16 {
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

// GetAt returns the value at (x, y, z, t) location
func (n *Nii1) getAt(x, y, z, t int32) float32 {

	tIndex := t * n.Data.Nx * n.Data.Ny * n.Data.Nz
	zIndex := n.Data.Nx * n.Data.Ny * z
	yIndex := n.Data.Nx * y
	xIndex := x
	index := tIndex + zIndex + yIndex + xIndex

	dataPoint := n.Data.Data[index*n.Data.NByPer : (index+1)*n.Data.NByPer]

	var value float32
	switch n.Data.NByPer {
	case constant.NIFTI_TYPE_BOOL:
		if len(dataPoint) > 0 {
			value = float32(dataPoint[0])
		}
	case constant.NIFTI_TYPE_UINT8:
		v := binary.LittleEndian.Uint16(dataPoint)
		value = float32(v)
	case constant.NIFTI_TYPE_INT16:
		v := binary.LittleEndian.Uint32(dataPoint)
		value = math.Float32frombits(v)
	case constant.NIFTI_TYPE_INT32:
		v := binary.LittleEndian.Uint64(dataPoint)
		value = float32(math.Float64frombits(v))
	default:
	}

	if n.Data.SclSlope != 0 {
		value = n.Data.SclSlope*value + n.Data.SclInter
	}

	return value
}

func (n *Nii1) getTimeSeries(x, y, z int32) ([]float32, error) {
	timeSeries := make([]float32, 0, n.Data.Dim[4])

	sliceX := n.Data.Nx
	sliceY := n.Data.Ny
	sliceZ := n.Data.Nx

	if x > sliceX {
		return nil, errors.New("invalid x value")
	}

	if y > sliceY {
		return nil, errors.New("invalid y value")
	}

	if z > sliceZ {
		return nil, errors.New("invalid z value")
	}

	for t := 0; t < int(n.Data.Dim[4]); t++ {
		timeSeries = append(timeSeries, n.getAt(x, y, z, int32(t)))
	}
	return timeSeries, nil
}

func (n *Nii1) getSlice(z, t int32) ([][]float32, error) {
	sliceX := n.Data.Nx
	sliceY := n.Data.Ny
	sliceZ := n.Data.Nx
	sliceT := n.Data.Nt

	if z > sliceZ {
		return nil, errors.New("invalid z value")
	}

	if t > sliceT || t < 0 {
		return nil, errors.New("invalid time value")
	}

	slice := make([][]float32, sliceX)
	for i := range slice {
		slice[i] = make([]float32, sliceY)
	}
	for x := 0; x < int(sliceX); x++ {
		for y := 0; y < int(sliceY); y++ {
			slice[x][y] = n.getAt(int32(x), int32(y), z, t)
		}
	}
	return slice, nil
}

func (n *Nii1) getDatatype() string {
	fmt.Println("n.Data.Datatype", n.Data.Datatype)
	switch int16(n.Data.Datatype) {
	case constant.DT_UNKNOWN:
		return "UNKNOWN"
	case constant.DT_BINARY:
		return "BINARY"
	case constant.DT_INT8:
		return "INT8"
	case constant.DT_UINT8:
		return "UINT8"
	case constant.DT_INT16:
		return "INT16"
	case constant.DT_UINT16:
		return "UINT16"
	case constant.DT_INT32:
		return "INT32"
	case constant.DT_UINT32:
		return "UINT32"
	case constant.DT_INT64:
		return "INT64"
	case constant.DT_UINT64:
		return "UINT64"
	case constant.DT_FLOAT32:
		return "FLOAT32"
	case constant.DT_FLOAT64:
		return "FLOAT64"
	case constant.DT_FLOAT128:
		return "FLOAT128"
	case constant.DT_COMPLEX64:
		return "COMPLEX64"
	case constant.DT_COMPLEX128:
		return "COMPLEX128"
	case constant.DT_COMPLEX256:
		return "COMPLEX256"
	case constant.DT_RGB24:
		return "RGB24"
	case constant.DT_RGBA32:
		return "RGBA32"
	}
	return "ILLEGAL"
}

func (n *Nii1) getOrientation() [3]string {
	res := [3]string{}

	ijk := n.Data.IJKOrtient

	iOrient, ok := constant.OrietationToString[int(ijk[0])]
	if !ok {
		res[0] = constant.OrietationToString[constant.NIFTI_UNKNOWN_ORIENT]
	}
	res[0] = iOrient

	jOrient, ok := constant.OrietationToString[int(ijk[1])]
	if !ok {
		res[1] = constant.OrietationToString[constant.NIFTI_UNKNOWN_ORIENT]
	}
	res[1] = jOrient

	kOrient, ok := constant.OrietationToString[int(ijk[2])]
	if !ok {
		res[2] = constant.OrietationToString[constant.NIFTI_UNKNOWN_ORIENT]
	}
	res[2] = kOrient

	return res
}

func (n *Nii1) getSliceCode() string {
	fmt.Println("n.Data.SliceCode", n.Data.SliceCode)
	switch n.Data.SliceCode {
	case constant.NIFTI_SLICE_UNKNOWN:
		return constant.NiiSliceAcquistionInfo[constant.NIFTI_SLICE_UNKNOWN]
	case constant.NIFTI_SLICE_SEQ_INC:
		return constant.NiiSliceAcquistionInfo[constant.NIFTI_SLICE_SEQ_INC]
	case constant.NIFTI_SLICE_SEQ_DEC:
		return constant.NiiSliceAcquistionInfo[constant.NIFTI_SLICE_SEQ_DEC]
	case constant.NIFTI_SLICE_ALT_INC:
		return constant.NiiSliceAcquistionInfo[constant.NIFTI_SLICE_ALT_INC]
	case constant.NIFTI_SLICE_ALT_DEC:
		return constant.NiiSliceAcquistionInfo[constant.NIFTI_SLICE_ALT_DEC]
	case constant.NIFTI_SLICE_ALT_INC2:
		return constant.NiiSliceAcquistionInfo[constant.NIFTI_SLICE_ALT_INC2]
	case constant.NIFTI_SLICE_ALT_DEC2:
		return constant.NiiSliceAcquistionInfo[constant.NIFTI_SLICE_ALT_DEC2]
	}

	return constants.COMMON_UNKNOWN
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
