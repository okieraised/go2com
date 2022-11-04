package reader

import (
	"encoding/binary"
	"errors"
	"github.com/okieraised/go2com/internal/constants"
	"github.com/okieraised/go2com/pkg/matrix"
	"github.com/okieraised/go2com/pkg/nifti/constant"
	"github.com/okieraised/go2com/pkg/nifti/nifti1"
	"github.com/okieraised/go2com/pkg/nifti/nifti2"
	"math"
)

type Nii struct {
	n1Header *nifti1.Nii1Header
	n2Header *nifti2.Nii2Header
	Data     *NiiData
}

// NiiData defines the structure of the NIFTI-1 data for I/O purpose
type NiiData struct {
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
	SclSlope      float64          // scaling parameter: slope
	SclInter      float64          // scaling parameter: intercept
	CalMin        float64          // calibration parameter: minimum
	CalMax        float64          // calibration parameter: maximum
	QformCode     int32            // codes for (x,y,z) space meaning
	SformCode     int32            // codes for (x,y,z) space meaning
	FreqDim       int32            // indices (1,2,3, or 0) for MRI
	PhaseDim      int32            // directions in dim[]/pixdim[]
	SliceDim      int32            // directions in dim[]/pixdim[]
	SliceCode     int32            // code for slice timing pattern
	SliceStart    int64            // index for start of slices
	SliceEnd      int64            // index for end of slices
	SliceDuration float64          // time between individual slices
	QuaternB      float64          // Quaternion transform parameters [when writing a dataset, these are used for qform, NOT qto_xyz]
	QuaternC      float64          // Quaternion transform parameters [when writing a dataset, these are used for qform, NOT qto_xyz]
	QuaternD      float64          // Quaternion transform parameters [when writing a dataset, these are used for qform, NOT qto_xyz]
	QOffsetX      float64          // Quaternion transform parameters [when writing a dataset, these are used for qform, NOT qto_xyz]
	QOffsetY      float64          // Quaternion transform parameters [when writing a dataset, these are used for qform, NOT qto_xyz]
	QOffsetZ      float64          // Quaternion transform parameters [when writing a dataset, these are used for qform, NOT qto_xyz]
	QFac          float64          // Quaternion transform parameters [when writing a dataset, these are used for qform, NOT qto_xyz]
	QtoXYZ        matrix.DMat44    // qform: transform (i,j,k) to (x,y,z)
	QtoIJK        matrix.DMat44    // qform: transform (x,y,z) to (i,j,k)
	StoXYZ        matrix.DMat44    // sform: transform (i,j,k) to (x,y,z)
	StoIJK        matrix.DMat44    // sform: transform (x,y,z) to (i,j,k)
	TOffset       float64          // time coordinate offset
	XYZUnits      int32            // dx,dy,dz units: NIFTI_UNITS_* code
	TimeUnits     int32            // dt units: NIFTI_UNITS_* code
	NiftiType     int32            // 0==Analyze, 1==NIFTI-1 (file), 2==NIFTI-1 (2 files), 3==NIFTI-ASCII (1 file)
	IntentCode    int32            // statistic type (or something)
	IntentP1      float64          // intent parameters
	IntentP2      float64          // intent parameters
	IntentP3      float64          // intent parameters
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

func (n *Nii) getSliceCode() string {
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

func (n *Nii) getQFormCode() string {
	qForm, ok := constant.NiiPatientOrientationInfo[uint8(n.Data.QformCode)]
	if !ok {
		return "Invalid"
	}
	return qForm
}

func (n *Nii) getSFormCode() string {
	sForm, ok := constant.NiiPatientOrientationInfo[uint8(n.Data.SformCode)]
	if !ok {
		return "Invalid"
	}
	return sForm
}

func (n *Nii) getDatatype() string {
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

func (n *Nii) getOrientation() [3]string {
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

// GetAt returns the value at (x, y, z, t) location
func (n *Nii) getAt(x, y, z, t int64) float64 {

	tIndex := t * int64(n.Data.Nx) * int64(n.Data.Ny) * int64(n.Data.Nz)
	zIndex := int64(n.Data.Nx) * int64(n.Data.Ny) * z
	yIndex := int64(n.Data.Nx) * y
	xIndex := x
	index := tIndex + zIndex + yIndex + xIndex
	nByPer := int64(n.Data.NByPer)

	dataPoint := n.Data.Data[index*nByPer : (index+1)*nByPer]

	var value float64
	switch n.Data.NByPer {
	case 0, 1:
		if len(dataPoint) > 0 {
			value = float64(dataPoint[0])
		}
	case 2: // This fits Uint16
		var v uint16
		switch n.Data.ByteOrder {
		case binary.LittleEndian:
			v = binary.LittleEndian.Uint16(dataPoint)
		case binary.BigEndian:
			v = binary.BigEndian.Uint16(dataPoint)
		}
		value = float64(v)
	case 3, 4: // This fits Uint32
		var v uint32
		switch n.Data.ByteOrder {
		case binary.LittleEndian:
			switch len(dataPoint) {
			case 3:
				v = uint32(uint(dataPoint[0]) | uint(dataPoint[1])<<8 | uint(dataPoint[2])<<16)
			case 4:
				v = binary.LittleEndian.Uint32(dataPoint)
			}
		case binary.BigEndian:
			switch len(dataPoint) {
			case 3:
				v = uint32(uint(dataPoint[2]) | uint(dataPoint[1])<<8 | uint(dataPoint[0])<<16)
			case 4:
				v = binary.BigEndian.Uint32(dataPoint)
			}
		}
		value = float64(math.Float32frombits(v))
	case 8:
		var v uint64
		switch n.Data.ByteOrder {
		case binary.LittleEndian:
			v = binary.LittleEndian.Uint64(dataPoint)
		case binary.BigEndian:
			v = binary.BigEndian.Uint64(dataPoint)
		}
		value = math.Float64frombits(v)
	case 16: // Unsupported
	case 32: // Unsupported
	default:
	}

	if n.Data.SclSlope != 0 {
		value = float64(n.Data.SclSlope)*value + float64(n.Data.SclInter)
	}

	return value
}

func (n *Nii) getTimeSeries(x, y, z int64) ([]float64, error) {
	timeSeries := make([]float64, 0, n.Data.Dim[4])

	sliceX := n.Data.Nx
	sliceY := n.Data.Ny
	sliceZ := n.Data.Nx

	if x >= int64(sliceX) {
		return nil, errors.New("invalid x value")
	}

	if y >= int64(sliceY) {
		return nil, errors.New("invalid y value")
	}

	if z >= int64(sliceZ) {
		return nil, errors.New("invalid z value")
	}

	for t := 0; t < int(n.Data.Dim[4]); t++ {
		timeSeries = append(timeSeries, n.getAt(x, y, z, int64(t)))
	}
	return timeSeries, nil
}

func (n *Nii) getSlice(z, t int64) ([][]float64, error) {
	sliceX := n.Data.Nx
	sliceY := n.Data.Ny
	sliceZ := n.Data.Nz
	sliceT := n.Data.Nt

	if z >= int64(sliceZ) {
		return nil, errors.New("invalid z value")
	}

	if t >= int64(sliceT) || t < 0 {
		return nil, errors.New("invalid time value")
	}

	slice := make([][]float64, sliceX)
	for i := range slice {
		slice[i] = make([]float64, sliceY)
	}
	for x := 0; x < int(sliceX); x++ {
		for y := 0; y < int(sliceY); y++ {
			slice[x][y] = n.getAt(int64(x), int64(y), z, t)
		}
	}
	return slice, nil
}

func (n *Nii) getVolume(t int64) ([][][]float64, error) {
	sliceX := n.Data.Nx
	sliceY := n.Data.Ny
	sliceZ := n.Data.Nz
	sliceT := n.Data.Nt

	if t >= int64(sliceT) || t < 0 {
		return nil, errors.New("invalid time value")
	}
	volume := make([][][]float64, sliceX)
	for i := range volume {
		volume[i] = make([][]float64, sliceY)
		for j := range volume[i] {
			volume[i][j] = make([]float64, sliceZ)
		}
	}
	for x := 0; x < int(sliceX); x++ {
		for y := 0; y < int(sliceY); y++ {
			for z := 0; z < int(sliceZ); z++ {
				volume[x][y][z] = n.getAt(int64(x), int64(y), int64(z), t)
			}
		}
	}
	return volume, nil
}
