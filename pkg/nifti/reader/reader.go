package reader

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/okieraised/go2com/internal/constants"
	"github.com/okieraised/go2com/internal/utils"
	"github.com/okieraised/go2com/pkg/matrix"
	"github.com/okieraised/go2com/pkg/nifti/constant"
)

type NiiReader interface {
	// Parse returns the input NIFTI as header and image data
	Parse() error
	// GetOrientation returns the image orientation
	GetOrientation() [3]string
	// GetSliceCode returns the slice code
	GetSliceCode() string
	// GetDatatype returns the datatype of the NIFTI image
	GetDatatype() string
	// GetSFormCode returns the SForm code string
	GetSFormCode() string
	// GetQFormCode returns the QForm code string
	GetQFormCode() string
	// GetImgShape returns the image shape [x, y, z, t]
	GetImgShape() [4]int64
	// GetAt returns the value at [x, y, z, t] as float64
	GetAt(x, y, z, t int64) float64
	// GetBinaryOrder returns the binary order of the NIFTI image
	GetBinaryOrder() binary.ByteOrder
	// GetUnitsOfMeasurements returns the spatial and temporal units of the NIFTI image
	GetUnitsOfMeasurements() ([2]string, error)
	// GetTimeSeries returns the time series of value at [x, y, z]
	GetTimeSeries(x, y, z int64) ([]float64, error)
	// GetSlice returns the X-Y slice at [z, t]
	GetSlice(z, t int64) ([][]float64, error)
	// GetVolume returns the image data as 3-D matrix
	GetVolume(t int64) ([][][]float64, error)
	// GetAffine returns the affine matrix of the NIFTI image
	GetAffine() matrix.DMat44
	// QuaternToMatrix converts the quarternions parameters to matrix
	QuaternToMatrix() matrix.DMat44
	// GetNiiData returns the raw NIFTI header and image data
	GetNiiData() *Nii
}

type niiReader struct {
	reader      *bytes.Reader
	binaryOrder binary.ByteOrder
	niiData     *Nii
	version     int
}

func NewNiiReader(filePath string) (NiiReader, error) {
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
	return &niiReader{
		binaryOrder: binary.LittleEndian,
		reader:      bytes.NewReader(bData),
		niiData:     &Nii{},
	}, nil
}

// checkNiiVersion checks the header to determine the NIFTI version
func (r *niiReader) checkNiiVersion() error {
	var hSize int32

	err := binary.Read(r.reader, r.binaryOrder, &hSize)
	if err != nil {
		return err
	}

	switch hSize {
	case constants.NII1HeaderSize:
		r.version = constants.NIIVersion1

	case constants.NII2HeaderSize:
		r.version = constants.NIIVersion2
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
			r.version = constants.NIIVersion1
		case constants.NII2HeaderSize:
			r.version = constants.NIIVersion2
		default:
			return errors.New("invalid NIFTI file format")
		}
	}
	return nil
}

func (r *niiReader) MatrixToOrientation(R matrix.DMat44) {
	r.niiData.matrixToOrientation(R)
}

func (r *niiReader) QuaternToMatrix() matrix.DMat44 {
	return r.niiData.quaternToMatrix()
}

func (r *niiReader) GetSliceCode() string {
	return r.niiData.getSliceCode()
}

func (r *niiReader) GetOrientation() [3]string {
	return r.niiData.getOrientation()
}

func (r *niiReader) GetDatatype() string {
	return r.niiData.getDatatype()
}

func (r *niiReader) GetSlice(z, t int64) ([][]float64, error) {
	return r.niiData.getSlice(z, t)
}

func (r *niiReader) GetTimeSeries(x, y, z int64) ([]float64, error) {
	return r.niiData.getTimeSeries(x, y, z)
}

func (r *niiReader) GetVolume(t int64) ([][][]float64, error) {
	return r.niiData.getVolume(t)
}

func (r *niiReader) GetAt(x, y, z, t int64) float64 {
	return r.niiData.getAt(x, y, z, t)
}

func (r *niiReader) GetUnitsOfMeasurements() ([2]string, error) {
	return r.niiData.getUnitsOfMeasurements()
}

func (r *niiReader) GetAffine() matrix.DMat44 {
	return r.niiData.getAffine()
}

func (r *niiReader) GetImgShape() [4]int64 {
	return r.niiData.getImgShape()
}

func (r *niiReader) GetQFormCode() string {
	return r.niiData.getQFormCode()
}

func (r *niiReader) GetSFormCode() string {
	return r.niiData.getSFormCode()
}

func (r *niiReader) GetNiiData() *Nii {
	return r.niiData
}

func (r *niiReader) GetBinaryOrder() binary.ByteOrder {
	return r.binaryOrder
}

// Parse returns the raw byte array into NIFTI-1 header and dataset structure
func (r *niiReader) Parse() error {
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
func (r *niiReader) parseHeader() error {
	_, err := r.reader.Seek(0, 0)
	if err != nil {
		return err
	}

	var dim0 int64

	switch r.version {
	case constants.NIIVersion1:
		header := new(Nii1Header)
		err = binary.Read(r.reader, r.binaryOrder, header)
		if err != nil {
			return err
		}
		if header.Magic != [4]byte{110, 43, 49, 0} && header.Magic != [4]byte{110, 105, 49, 0} {
			return errors.New("invalid NIFTI-1 magic string")
		}
		r.niiData.n1Header = header
		dim0 = int64(header.Dim[0])
	case constants.NIIVersion2:
		header := new(Nii2Header)
		err = binary.Read(r.reader, r.binaryOrder, header)
		if err != nil {
			return err
		}
		if header.Magic != [8]byte{110, 43, 50, 0, 13, 10, 26, 10} {
			return errors.New("invalid NIFTI-2 magic string")
		}
		r.niiData.n2Header = header
		dim0 = header.Dim[0]
	default:
		return errors.New("invalid version")
	}

	if dim0 < 0 || dim0 > 7 {
		if r.binaryOrder == binary.LittleEndian {
			r.binaryOrder = binary.BigEndian
		} else {
			r.binaryOrder = binary.LittleEndian
		}
	}
	return nil
}

// setDatatypeSize sets number of bytes per voxel and the swapsize for the header datatype
func (r *niiReader) setDatatypeSize() {
	var NByPerVoxel int32 = 0
	var SwapSize int32 = 0
	var datatype int16

	switch r.version {
	case constants.NIIVersion1:
		datatype = r.niiData.n1Header.Datatype
	case constants.NIIVersion2:
		datatype = r.niiData.n2Header.Datatype
	}

	switch datatype {
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
func (r *niiReader) parseData() error {
	r.niiData.Data = &NiiData{}
	var offset int64
	var statDim int64 = 1
	var bitpix int16
	var qFormCode, sFormCode, intentCode, sliceCode, datatype int32
	var pixDim0, sclSlope, sclInter, intentP1, intentP2, intentP3, quaternB, quaternC, quaternD, sliceDuration, calMin, calMax float64
	var sRowX, sRowY, sRowZ [4]float64
	var intentName [16]uint8
	var descrip [80]uint8
	var sliceStart, sliceEnd int64
	var voxOffset int64
	var freqDim, phaseDim, sliceDim int32

	switch r.version {
	case constants.NIIVersion1:
		freqDim = int32(dimInfoToFreqDim(r.niiData.n1Header.DimInfo))
		phaseDim = int32(dimInfoToPhaseDim(r.niiData.n1Header.DimInfo))
		sliceDim = int32(dimInfoToSliceDim(r.niiData.n1Header.DimInfo))

		voxOffset = int64(r.niiData.n1Header.VoxOffset)
		datatype = int32(r.niiData.n1Header.Datatype)

		// The bits 1-3 are used to store the spatial dimensions, the bits 4-6 are for temporal dimensions,
		// and the bits 6 and 7 are not used
		r.niiData.Data.XYZUnits = int32(r.niiData.n1Header.XyztUnits % 8)
		r.niiData.Data.TimeUnits = int32(r.niiData.n1Header.XyztUnits) - r.niiData.Data.XYZUnits

		sliceCode = int32(r.niiData.n1Header.SliceCode)
		sliceStart = int64(r.niiData.n1Header.SliceStart)
		sliceEnd = int64(r.niiData.n1Header.SliceEnd)
		sliceDuration = float64(r.niiData.n1Header.SliceDuration)

		calMin = float64(r.niiData.n1Header.CalMin)
		calMax = float64(r.niiData.n1Header.CalMax)

		qFormCode = int32(r.niiData.n1Header.QformCode)
		sFormCode = int32(r.niiData.n1Header.SformCode)
		pixDim0 = float64(r.niiData.n1Header.Pixdim[0])

		sRowX = ConvertToF64(r.niiData.n1Header.SrowX)
		sRowY = ConvertToF64(r.niiData.n1Header.SrowY)
		sRowZ = ConvertToF64(r.niiData.n1Header.SrowZ)

		sclSlope = float64(r.niiData.n1Header.SclSlope)
		sclInter = float64(r.niiData.n1Header.SclInter)

		intentName = r.niiData.n1Header.IntentName
		intentCode = int32(r.niiData.n1Header.IntentCode)
		intentP1 = float64(r.niiData.n1Header.IntentP1)
		intentP2 = float64(r.niiData.n1Header.IntentP2)
		intentP3 = float64(r.niiData.n1Header.IntentP3)

		quaternB = float64(r.niiData.n1Header.QuaternB)
		quaternC = float64(r.niiData.n1Header.QuaternC)
		quaternD = float64(r.niiData.n1Header.QuaternD)
		descrip = r.niiData.n1Header.Descrip

		// Set the dimension of data array
		r.niiData.Data.NDim, r.niiData.Data.Dim[0] = int64(r.niiData.n1Header.Dim[0]), int64(r.niiData.n1Header.Dim[0])
		r.niiData.Data.Nx, r.niiData.Data.Dim[1] = int64(r.niiData.n1Header.Dim[1]), int64(r.niiData.n1Header.Dim[1])
		r.niiData.Data.Ny, r.niiData.Data.Dim[2] = int64(r.niiData.n1Header.Dim[2]), int64(r.niiData.n1Header.Dim[2])
		r.niiData.Data.Nz, r.niiData.Data.Dim[3] = int64(r.niiData.n1Header.Dim[3]), int64(r.niiData.n1Header.Dim[3])
		r.niiData.Data.Nt, r.niiData.Data.Dim[4] = int64(r.niiData.n1Header.Dim[4]), int64(r.niiData.n1Header.Dim[4])
		r.niiData.Data.Nu, r.niiData.Data.Dim[5] = int64(r.niiData.n1Header.Dim[5]), int64(r.niiData.n1Header.Dim[5])
		r.niiData.Data.Nv, r.niiData.Data.Dim[6] = int64(r.niiData.n1Header.Dim[6]), int64(r.niiData.n1Header.Dim[6])
		r.niiData.Data.Nw, r.niiData.Data.Dim[7] = int64(r.niiData.n1Header.Dim[7]), int64(r.niiData.n1Header.Dim[7])

		// Set the grid spacing
		r.niiData.Data.Dx, r.niiData.Data.PixDim[1] = float64(r.niiData.n1Header.Pixdim[1]), float64(r.niiData.n1Header.Pixdim[1])
		r.niiData.Data.Dy, r.niiData.Data.PixDim[2] = float64(r.niiData.n1Header.Pixdim[2]), float64(r.niiData.n1Header.Pixdim[2])
		r.niiData.Data.Dz, r.niiData.Data.PixDim[3] = float64(r.niiData.n1Header.Pixdim[3]), float64(r.niiData.n1Header.Pixdim[3])
		r.niiData.Data.Dt, r.niiData.Data.PixDim[4] = float64(r.niiData.n1Header.Pixdim[4]), float64(r.niiData.n1Header.Pixdim[4])
		r.niiData.Data.Du, r.niiData.Data.PixDim[5] = float64(r.niiData.n1Header.Pixdim[5]), float64(r.niiData.n1Header.Pixdim[5])
		r.niiData.Data.Dv, r.niiData.Data.PixDim[6] = float64(r.niiData.n1Header.Pixdim[6]), float64(r.niiData.n1Header.Pixdim[6])
		r.niiData.Data.Dw, r.niiData.Data.PixDim[7] = float64(r.niiData.n1Header.Pixdim[7]), float64(r.niiData.n1Header.Pixdim[7])

		bitpix = r.niiData.n1Header.Bitpix

	case constants.NIIVersion2:
		freqDim = int32(dimInfoToFreqDim(r.niiData.n2Header.DimInfo))
		phaseDim = int32(dimInfoToPhaseDim(r.niiData.n2Header.DimInfo))
		sliceDim = int32(dimInfoToSliceDim(r.niiData.n2Header.DimInfo))

		voxOffset = r.niiData.n2Header.VoxOffset
		datatype = int32(r.niiData.n2Header.Datatype)

		// The bits 1-3 are used to store the spatial dimensions, the bits 4-6 are for temporal dimensions,
		// and the bits 6 and 7 are not used
		r.niiData.Data.XYZUnits = r.niiData.n2Header.XyztUnits % 8
		r.niiData.Data.TimeUnits = r.niiData.n2Header.XyztUnits - r.niiData.Data.XYZUnits

		sliceCode = r.niiData.n2Header.SliceCode
		sliceStart = r.niiData.n2Header.SliceStart
		sliceEnd = r.niiData.n2Header.SliceEnd
		sliceDuration = r.niiData.n2Header.SliceDuration

		calMin = r.niiData.n2Header.CalMin
		calMax = r.niiData.n2Header.CalMax

		qFormCode = r.niiData.n2Header.QformCode
		pixDim0 = r.niiData.n2Header.Pixdim[0]
		sFormCode = r.niiData.n2Header.SformCode

		sclSlope = r.niiData.n2Header.SclSlope
		sclInter = r.niiData.n2Header.SclInter

		intentName = r.niiData.n2Header.IntentName
		intentCode = r.niiData.n2Header.IntentCode
		r.niiData.Data.IntentP1 = r.niiData.n2Header.IntentP1
		r.niiData.Data.IntentP2 = r.niiData.n2Header.IntentP2
		r.niiData.Data.IntentP3 = r.niiData.n2Header.IntentP3

		r.niiData.Data.QuaternB = r.niiData.n2Header.QuaternB
		r.niiData.Data.QuaternC = r.niiData.n2Header.QuaternC
		r.niiData.Data.QuaternD = r.niiData.n2Header.QuaternD
		descrip = r.niiData.n2Header.Descrip

		// Set the dimension of data array
		r.niiData.Data.NDim, r.niiData.Data.Dim[0] = r.niiData.n2Header.Dim[0], r.niiData.n2Header.Dim[0]
		r.niiData.Data.Nx, r.niiData.Data.Dim[1] = r.niiData.n2Header.Dim[1], r.niiData.n2Header.Dim[1]
		r.niiData.Data.Ny, r.niiData.Data.Dim[2] = r.niiData.n2Header.Dim[2], r.niiData.n2Header.Dim[2]
		r.niiData.Data.Nz, r.niiData.Data.Dim[3] = r.niiData.n2Header.Dim[3], r.niiData.n2Header.Dim[3]
		r.niiData.Data.Nt, r.niiData.Data.Dim[4] = r.niiData.n2Header.Dim[4], r.niiData.n2Header.Dim[4]
		r.niiData.Data.Nu, r.niiData.Data.Dim[5] = r.niiData.n2Header.Dim[5], r.niiData.n2Header.Dim[5]
		r.niiData.Data.Nv, r.niiData.Data.Dim[6] = r.niiData.n2Header.Dim[6], r.niiData.n2Header.Dim[6]
		r.niiData.Data.Nw, r.niiData.Data.Dim[7] = r.niiData.n2Header.Dim[7], r.niiData.n2Header.Dim[7]

		// Set the grid spacing
		r.niiData.Data.Dx, r.niiData.Data.PixDim[1] = r.niiData.n2Header.Pixdim[1], r.niiData.n2Header.Pixdim[1]
		r.niiData.Data.Dy, r.niiData.Data.PixDim[2] = r.niiData.n2Header.Pixdim[2], r.niiData.n2Header.Pixdim[2]
		r.niiData.Data.Dz, r.niiData.Data.PixDim[3] = r.niiData.n2Header.Pixdim[3], r.niiData.n2Header.Pixdim[3]
		r.niiData.Data.Dt, r.niiData.Data.PixDim[4] = r.niiData.n2Header.Pixdim[4], r.niiData.n2Header.Pixdim[4]
		r.niiData.Data.Du, r.niiData.Data.PixDim[5] = r.niiData.n2Header.Pixdim[5], r.niiData.n2Header.Pixdim[5]
		r.niiData.Data.Dv, r.niiData.Data.PixDim[6] = r.niiData.n2Header.Pixdim[6], r.niiData.n2Header.Pixdim[6]
		r.niiData.Data.Dw, r.niiData.Data.PixDim[7] = r.niiData.n2Header.Pixdim[7], r.niiData.n2Header.Pixdim[7]

		bitpix = r.niiData.n2Header.Bitpix

		// SRowX, SRowY, SRowZ
		sRowX, sRowY, sRowZ = r.niiData.n2Header.SrowX, r.niiData.n2Header.SrowY, r.niiData.n2Header.SrowZ
	}

	// Fix bad value in header
	if r.niiData.Data.Nz <= 0 && r.niiData.Data.Dim[3] <= 0 {
		r.niiData.Data.Nz = 1
		r.niiData.Data.Dim[3] = 1
	}
	if r.niiData.Data.Nt <= 0 && r.niiData.Data.Dim[4] <= 0 {
		r.niiData.Data.Nt = 1
		r.niiData.Data.Dim[4] = 1
	}
	if r.niiData.Data.Nu <= 0 && r.niiData.Data.Dim[5] <= 0 {
		r.niiData.Data.Nu = 1
		r.niiData.Data.Dim[5] = 1
	}
	if r.niiData.Data.Nv <= 0 && r.niiData.Data.Dim[6] <= 0 {
		r.niiData.Data.Nv = 1
		r.niiData.Data.Dim[6] = 1
	}
	if r.niiData.Data.Nw <= 0 && r.niiData.Data.Dim[7] <= 0 {
		r.niiData.Data.Nw = 1
		r.niiData.Data.Dim[7] = 1
	}

	// Set the byte order
	r.niiData.Data.ByteOrder = r.binaryOrder

	if bitpix == 0 {
		return errors.New("number of bits per voxel value (bitpix) is zero")
	}

	r.niiData.Data.NVox = 1
	for i := int64(1); i <= r.niiData.Data.NDim; i++ {
		r.niiData.Data.NVox *= r.niiData.Data.Dim[i]
	}

	r.setDatatypeSize()

	// compute QToXYK transformation from pixel indexes (i,j,k) to (x,y,z)
	if qFormCode <= 0 {
		r.niiData.Data.QtoXYZ.M[0][0] = r.niiData.Data.Dx
		r.niiData.Data.QtoXYZ.M[1][1] = r.niiData.Data.Dy
		r.niiData.Data.QtoXYZ.M[2][2] = r.niiData.Data.Dz

		// off diagonal is zero
		r.niiData.Data.QtoXYZ.M[0][1] = 0
		r.niiData.Data.QtoXYZ.M[0][2] = 0
		r.niiData.Data.QtoXYZ.M[0][3] = 0

		r.niiData.Data.QtoXYZ.M[1][0] = 0
		r.niiData.Data.QtoXYZ.M[1][2] = 0
		r.niiData.Data.QtoXYZ.M[1][3] = 0

		r.niiData.Data.QtoXYZ.M[2][0] = 0
		r.niiData.Data.QtoXYZ.M[2][1] = 0
		r.niiData.Data.QtoXYZ.M[2][3] = 0

		// last row is [0, 0, 0, 1]
		r.niiData.Data.QtoXYZ.M[3][0] = 0
		r.niiData.Data.QtoXYZ.M[3][1] = 0
		r.niiData.Data.QtoXYZ.M[3][2] = 0
		r.niiData.Data.QtoXYZ.M[3][3] = 1.0

		r.niiData.Data.QformCode = constant.NIFTI_XFORM_UNKNOWN
	} else {
		if pixDim0 < 0 {
			r.niiData.Data.QFac = -1
		} else {
			r.niiData.Data.QFac = 1
		}
		r.niiData.Data.QtoXYZ = r.QuaternToMatrix()
		r.niiData.Data.QformCode = qFormCode
	}

	// Set QToIJK
	r.niiData.Data.QtoIJK = matrix.Mat44Inverse(r.niiData.Data.QtoXYZ)

	if sFormCode <= 0 {
		r.niiData.Data.SformCode = constant.NIFTI_XFORM_UNKNOWN
	} else {
		r.niiData.Data.StoXYZ.M[0][0] = sRowX[0]
		r.niiData.Data.StoXYZ.M[0][1] = sRowX[1]
		r.niiData.Data.StoXYZ.M[0][2] = sRowX[2]
		r.niiData.Data.StoXYZ.M[0][3] = sRowX[3]

		r.niiData.Data.StoXYZ.M[1][0] = sRowY[0]
		r.niiData.Data.StoXYZ.M[1][1] = sRowY[1]
		r.niiData.Data.StoXYZ.M[1][2] = sRowY[2]
		r.niiData.Data.StoXYZ.M[1][3] = sRowY[3]

		r.niiData.Data.StoXYZ.M[2][0] = sRowZ[0]
		r.niiData.Data.StoXYZ.M[2][1] = sRowZ[0]
		r.niiData.Data.StoXYZ.M[2][2] = sRowZ[0]
		r.niiData.Data.StoXYZ.M[2][3] = sRowZ[0]

		r.niiData.Data.StoXYZ.M[3][0] = 0
		r.niiData.Data.StoXYZ.M[3][1] = 0
		r.niiData.Data.StoXYZ.M[3][2] = 0
		r.niiData.Data.StoXYZ.M[3][3] = 1

		r.niiData.Data.StoIJK = matrix.Mat44Inverse(r.niiData.Data.StoXYZ)

		r.niiData.Data.SformCode = sFormCode
	}

	// Other stuff
	r.niiData.Data.SclSlope = sclSlope
	r.niiData.Data.SclInter = sclInter

	r.niiData.Data.IntentName = intentName
	r.niiData.Data.IntentCode = intentCode
	r.niiData.Data.IntentP1 = intentP1
	r.niiData.Data.IntentP2 = intentP2
	r.niiData.Data.IntentP3 = intentP3

	r.niiData.Data.QuaternB = quaternB
	r.niiData.Data.QuaternC = quaternC
	r.niiData.Data.QuaternD = quaternD
	r.niiData.Data.Descrip = descrip

	// Frequency dimension, phase dimension, slice dimension
	r.niiData.Data.FreqDim = freqDim
	r.niiData.Data.PhaseDim = phaseDim
	r.niiData.Data.SliceDim = sliceDim

	r.niiData.Data.SliceCode = sliceCode
	r.niiData.Data.SliceStart = sliceStart
	r.niiData.Data.SliceEnd = sliceEnd
	r.niiData.Data.SliceDuration = sliceDuration

	r.niiData.Data.CalMin = calMin
	r.niiData.Data.CalMax = calMax

	r.niiData.Data.Datatype = datatype

	if r.niiData.Data.Dim[5] > 1 {
		statDim = r.niiData.Data.Dim[5]
	}
	offset = voxOffset
	dataSize := r.niiData.Data.Dim[1] * r.niiData.Data.Dim[2] * r.niiData.Data.Dim[3] * r.niiData.Data.Dim[4] * statDim * (int64(bitpix) / 8)

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

	affine := matrix.DMat44{}
	affine.M[0] = sRowX
	affine.M[1] = sRowY
	affine.M[2] = sRowZ
	affine.M[3] = [4]float64{0, 0, 0, 1}

	r.niiData.Data.Affine = affine
	r.niiData.matrixToOrientation(affine)

	return nil
}

func dimInfoToFreqDim(DimInfo uint8) uint8 {
	return DimInfo & 0x03
}

func dimInfoToPhaseDim(DimInfo uint8) uint8 {
	return (DimInfo >> 2) & 0x03
}

func dimInfoToSliceDim(DimInfo uint8) uint8 {
	return (DimInfo >> 4) & 0x03
}

func ConvertToF64(ar [4]float32) [4]float64 {
	newar := [4]float64{}
	var v float32
	var i int
	for i, v = range ar {
		newar[i] = float64(v)
	}
	return newar
}
