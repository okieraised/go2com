package reader

import (
	"encoding/binary"
	"github.com/okieraised/go2com/pkg/matrix"
)

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
	QtoXYZ        matrix.DMat44    // qform: transform (i,j,k) to (x,y,z)
	QtoIJK        matrix.DMat44    // qform: transform (x,y,z) to (i,j,k)
	StoXYZ        matrix.DMat44    // sform: transform (i,j,k) to (x,y,z)
	StoIJK        matrix.DMat44    // sform: transform (x,y,z) to (i,j,k)
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

// Nii2Header defines the structure of the NIFTI-2 header
type Nii2Header struct {
	SizeofHdr     int32
	Magic         [8]uint8
	Datatype      int16
	Bitpix        int16
	Dim           [8]int64
	IntentP1      float64
	IntentP2      float64
	IntentP3      float64
	Pixdim        [8]float64
	VoxOffset     int64
	SclSlope      float64
	SclInter      float64
	CalMax        float64
	CalMin        float64
	SliceDuration float64
	Toffset       float64
	SliceStart    int64
	SliceEnd      int64
	Descrip       [80]uint8
	AuxFile       [24]uint8
	QformCode     int32
	SformCode     int32
	QuaternB      float64
	QuaternC      float64
	QuaternD      float64
	QoffsetX      float64
	QoffsetY      float64
	QoffsetZ      float64
	SrowX         [4]float64
	SrowY         [4]float64
	SrowZ         [4]float64
	SliceCode     int32
	XyztUnits     int32
	IntentCode    int32
	IntentName    [16]uint8
	DimInfo       uint8
	UnusedStr     [15]uint8
}

type Nii2Data struct {
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
