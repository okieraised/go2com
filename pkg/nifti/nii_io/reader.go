package nii_io

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/okieraised/go2com/internal/matrix"
	"io"
	"net/http"
	"os"

	"github.com/okieraised/go2com/internal/utils"
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
	binaryOrder binary.ByteOrder // Default system order
	inMemory    bool             // Whether to read the whole NIfTI image to memory
	data        *Nii             // Contains the NIFTI data structure
	version     int              // Define the version of NIFTI image (1 or 2)
}

// NewNiiReader receives a path to the NIFTI file and returns a new reader to parse the file
//
// TODO: this is not efficient when the file is large so we need to find better way to deal with large file size
func NewNiiReader(filePath string, options ...func(*niiReader)) (NiiReader, error) {
	// This is inefficient since it read the whole file to the memory
	bData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Check the content type to see if the file is gzipped. Do not depend on just the extensions of the file
	mimeType := http.DetectContentType(bData[:512])
	if mimeType == "application/x-gzip" {
		bData, err = utils.DeflateGzip(bData)
		if err != nil {
			return nil, err
		}
	}

	reader := new(niiReader)
	reader.binaryOrder = binary.LittleEndian
	reader.reader = bytes.NewReader(bData)
	reader.data = &Nii{}

	for _, opt := range options {
		opt(reader)
	}

	return reader, nil
}

func WithInMemory(inMemory bool) func(*niiReader) {
	return func(w *niiReader) {
		w.inMemory = inMemory
	}
}

//----------------------------------------------------------------------------------------------------------------------

// getVersion checks the header to determine the NIFTI version
func (r *niiReader) getVersion() error {
	var hSize int32

	err := binary.Read(r.reader, r.binaryOrder, &hSize)
	if err != nil {
		return err
	}

	switch hSize {
	case constant.NII1HeaderSize:
		r.version = constant.NIIVersion1
	case constant.NII2HeaderSize:
		r.version = constant.NIIVersion2
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
		case constant.NII1HeaderSize:
			r.version = constant.NIIVersion1
		case constant.NII2HeaderSize:
			r.version = constant.NIIVersion2
		default:
			return errors.New("invalid NIFTI file format")
		}
	}
	r.data.Version = r.version
	return nil
}

func (r *niiReader) MatrixToOrientation(R matrix.DMat44) {
	r.data.matrixToOrientation(R)
}

func (r *niiReader) QuaternToMatrix() matrix.DMat44 {
	return r.data.quaternToMatrix()
}

func (r *niiReader) GetSliceCode() string {
	return r.data.getSliceCode()
}

func (r *niiReader) GetOrientation() [3]string {
	return r.data.getOrientation()
}

func (r *niiReader) GetDatatype() string {
	return r.data.getDatatype()
}

func (r *niiReader) GetSlice(z, t int64) ([][]float64, error) {
	return r.data.getSlice(z, t)
}

func (r *niiReader) GetTimeSeries(x, y, z int64) ([]float64, error) {
	return r.data.getTimeSeries(x, y, z)
}

func (r *niiReader) GetVolume(t int64) ([][][]float64, error) {
	return r.data.getVolume(t)
}

func (r *niiReader) GetAt(x, y, z, t int64) float64 {
	return r.data.getAt(x, y, z, t)
}

func (r *niiReader) GetUnitsOfMeasurements() ([2]string, error) {
	return r.data.getUnitsOfMeasurements()
}

func (r *niiReader) GetAffine() matrix.DMat44 {
	return r.data.getAffine()
}

func (r *niiReader) GetImgShape() [4]int64 {
	return r.data.getImgShape()
}

func (r *niiReader) GetQFormCode() string {
	return r.data.getQFormCode()
}

func (r *niiReader) GetSFormCode() string {
	return r.data.getSFormCode()
}

func (r *niiReader) GetNiiData() *Nii {
	return r.data
}

func (r *niiReader) GetBinaryOrder() binary.ByteOrder {
	return r.binaryOrder
}

// Parse returns the raw byte array into NIFTI-1 header and dataset structure
func (r *niiReader) Parse() error {
	err := r.getVersion()
	if err != nil {
		return err
	}

	err = r.parseNIfTI()
	if err != nil {
		return err
	}

	return nil
}

// parseNIfTI parse the NIfTI header and the data
func (r *niiReader) parseNIfTI() error {
	_, err := r.reader.Seek(0, 0)
	if err != nil {
		return err
	}

	var dim0 int64
	var header interface{}

	switch r.version {
	case constant.NIIVersion1:
		n1Header := new(Nii1Header)
		err = binary.Read(r.reader, r.binaryOrder, n1Header)
		if err != nil {
			return err
		}
		if n1Header.Magic != [4]byte{110, 43, 49, 0} && n1Header.Magic != [4]byte{110, 105, 49, 0} {
			return errors.New("invalid NIFTI-1 magic string")
		}
		dim0 = int64(n1Header.Dim[0])

		if dim0 < 0 || dim0 > 7 {
			if r.binaryOrder == binary.LittleEndian {
				r.binaryOrder = binary.BigEndian
			} else {
				r.binaryOrder = binary.LittleEndian
			}
		}
		header = n1Header
	case constant.NIIVersion2:
		n2Header := new(Nii2Header)
		err = binary.Read(r.reader, r.binaryOrder, n2Header)
		if err != nil {
			return err
		}
		if n2Header.Magic != [8]byte{110, 43, 50, 0, 13, 10, 26, 10} {
			return errors.New("invalid NIFTI-2 magic string")
		}
		dim0 = n2Header.Dim[0]

		if dim0 < 0 || dim0 > 7 {
			if r.binaryOrder == binary.LittleEndian {
				r.binaryOrder = binary.BigEndian
			} else {
				r.binaryOrder = binary.LittleEndian
			}
		}
		header = n2Header
	default:
		return errors.New("invalid version")
	}
	err = r.parseData(header)
	if err != nil {
		return err
	}
	return nil
}

// parseData parse the raw byte array into NIFTI-1 or NIFTI-2 data structure
func (r *niiReader) parseData(header interface{}) error {
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
	case constant.NIIVersion1:
		n1Header := header.(*Nii1Header)

		freqDim = int32(dimInfoToFreqDim(n1Header.DimInfo))
		phaseDim = int32(dimInfoToPhaseDim(n1Header.DimInfo))
		sliceDim = int32(dimInfoToSliceDim(n1Header.DimInfo))

		voxOffset = int64(n1Header.VoxOffset)
		datatype = int32(n1Header.Datatype)

		// The bits 1-3 are used to store the spatial dimensions, the bits 4-6 are for temporal dimensions,
		// and the bits 6 and 7 are not used
		r.data.XYZUnits = int32(n1Header.XyztUnits % 8)
		r.data.TimeUnits = int32(n1Header.XyztUnits) - r.data.XYZUnits

		sliceCode = int32(n1Header.SliceCode)
		sliceStart = int64(n1Header.SliceStart)
		sliceEnd = int64(n1Header.SliceEnd)
		sliceDuration = float64(n1Header.SliceDuration)

		calMin = float64(n1Header.CalMin)
		calMax = float64(n1Header.CalMax)

		qFormCode = int32(n1Header.QformCode)
		sFormCode = int32(n1Header.SformCode)
		pixDim0 = float64(n1Header.Pixdim[0])

		sRowX = convertToF64(n1Header.SrowX)
		sRowY = convertToF64(n1Header.SrowY)
		sRowZ = convertToF64(n1Header.SrowZ)

		sclSlope = float64(n1Header.SclSlope)
		sclInter = float64(n1Header.SclInter)

		intentName = n1Header.IntentName
		intentCode = int32(n1Header.IntentCode)
		intentP1 = float64(n1Header.IntentP1)
		intentP2 = float64(n1Header.IntentP2)
		intentP3 = float64(n1Header.IntentP3)

		quaternB = float64(n1Header.QuaternB)
		quaternC = float64(n1Header.QuaternC)
		quaternD = float64(n1Header.QuaternD)
		descrip = n1Header.Descrip

		// Set the dimension of data array
		r.data.NDim, r.data.Dim[0] = int64(n1Header.Dim[0]), int64(n1Header.Dim[0])
		r.data.Nx, r.data.Dim[1] = int64(n1Header.Dim[1]), int64(n1Header.Dim[1])
		r.data.Ny, r.data.Dim[2] = int64(n1Header.Dim[2]), int64(n1Header.Dim[2])
		r.data.Nz, r.data.Dim[3] = int64(n1Header.Dim[3]), int64(n1Header.Dim[3])
		r.data.Nt, r.data.Dim[4] = int64(n1Header.Dim[4]), int64(n1Header.Dim[4])
		r.data.Nu, r.data.Dim[5] = int64(n1Header.Dim[5]), int64(n1Header.Dim[5])
		r.data.Nv, r.data.Dim[6] = int64(n1Header.Dim[6]), int64(n1Header.Dim[6])
		r.data.Nw, r.data.Dim[7] = int64(n1Header.Dim[7]), int64(n1Header.Dim[7])

		// Set the grid spacing
		r.data.Dx, r.data.PixDim[1] = float64(n1Header.Pixdim[1]), float64(n1Header.Pixdim[1])
		r.data.Dy, r.data.PixDim[2] = float64(n1Header.Pixdim[2]), float64(n1Header.Pixdim[2])
		r.data.Dz, r.data.PixDim[3] = float64(n1Header.Pixdim[3]), float64(n1Header.Pixdim[3])
		r.data.Dt, r.data.PixDim[4] = float64(n1Header.Pixdim[4]), float64(n1Header.Pixdim[4])
		r.data.Du, r.data.PixDim[5] = float64(n1Header.Pixdim[5]), float64(n1Header.Pixdim[5])
		r.data.Dv, r.data.PixDim[6] = float64(n1Header.Pixdim[6]), float64(n1Header.Pixdim[6])
		r.data.Dw, r.data.PixDim[7] = float64(n1Header.Pixdim[7]), float64(n1Header.Pixdim[7])

		bitpix = n1Header.Bitpix

		NByPerVoxel, SwapSize := assignDatatypeSize(datatype)
		r.data.NByPer = int32(NByPerVoxel)
		r.data.SwapSize = int32(SwapSize)

		r.data.QuaternB, r.data.QuaternC, r.data.QuaternD = float64(n1Header.QuaternB), float64(n1Header.QuaternC), float64(n1Header.QuaternD)
		r.data.QoffsetX, r.data.QoffsetY, r.data.QoffsetZ = float64(n1Header.QoffsetX), float64(n1Header.QoffsetY), float64(n1Header.QoffsetZ)

	case constant.NIIVersion2:
		n2Header := header.(*Nii2Header)

		freqDim = int32(dimInfoToFreqDim(n2Header.DimInfo))
		phaseDim = int32(dimInfoToPhaseDim(n2Header.DimInfo))
		sliceDim = int32(dimInfoToSliceDim(n2Header.DimInfo))

		voxOffset = n2Header.VoxOffset
		datatype = int32(n2Header.Datatype)

		// The bits 1-3 are used to store the spatial dimensions, the bits 4-6 are for temporal dimensions,
		// and the bits 6 and 7 are not used
		r.data.XYZUnits = n2Header.XyztUnits % 8
		r.data.TimeUnits = n2Header.XyztUnits - r.data.XYZUnits

		sliceCode = n2Header.SliceCode
		sliceStart = n2Header.SliceStart
		sliceEnd = n2Header.SliceEnd
		sliceDuration = n2Header.SliceDuration

		calMin = n2Header.CalMin
		calMax = n2Header.CalMax

		qFormCode = n2Header.QformCode
		pixDim0 = n2Header.Pixdim[0]
		sFormCode = n2Header.SformCode

		sclSlope = n2Header.SclSlope
		sclInter = n2Header.SclInter

		intentName = n2Header.IntentName
		intentCode = n2Header.IntentCode
		r.data.IntentP1 = n2Header.IntentP1
		r.data.IntentP2 = n2Header.IntentP2
		r.data.IntentP3 = n2Header.IntentP3

		r.data.QuaternB = n2Header.QuaternB
		r.data.QuaternC = n2Header.QuaternC
		r.data.QuaternD = n2Header.QuaternD
		descrip = n2Header.Descrip

		// Set the dimension of data array
		r.data.NDim, r.data.Dim[0] = n2Header.Dim[0], n2Header.Dim[0]
		r.data.Nx, r.data.Dim[1] = n2Header.Dim[1], n2Header.Dim[1]
		r.data.Ny, r.data.Dim[2] = n2Header.Dim[2], n2Header.Dim[2]
		r.data.Nz, r.data.Dim[3] = n2Header.Dim[3], n2Header.Dim[3]
		r.data.Nt, r.data.Dim[4] = n2Header.Dim[4], n2Header.Dim[4]
		r.data.Nu, r.data.Dim[5] = n2Header.Dim[5], n2Header.Dim[5]
		r.data.Nv, r.data.Dim[6] = n2Header.Dim[6], n2Header.Dim[6]
		r.data.Nw, r.data.Dim[7] = n2Header.Dim[7], n2Header.Dim[7]

		// Set the grid spacing
		r.data.Dx, r.data.PixDim[1] = n2Header.Pixdim[1], n2Header.Pixdim[1]
		r.data.Dy, r.data.PixDim[2] = n2Header.Pixdim[2], n2Header.Pixdim[2]
		r.data.Dz, r.data.PixDim[3] = n2Header.Pixdim[3], n2Header.Pixdim[3]
		r.data.Dt, r.data.PixDim[4] = n2Header.Pixdim[4], n2Header.Pixdim[4]
		r.data.Du, r.data.PixDim[5] = n2Header.Pixdim[5], n2Header.Pixdim[5]
		r.data.Dv, r.data.PixDim[6] = n2Header.Pixdim[6], n2Header.Pixdim[6]
		r.data.Dw, r.data.PixDim[7] = n2Header.Pixdim[7], n2Header.Pixdim[7]

		bitpix = n2Header.Bitpix

		// SRowX, SRowY, SRowZ
		sRowX, sRowY, sRowZ = n2Header.SrowX, n2Header.SrowY, n2Header.SrowZ

		NByPerVoxel, SwapSize := assignDatatypeSize(datatype)
		r.data.NByPer = int32(NByPerVoxel)
		r.data.SwapSize = int32(SwapSize)

		r.data.QuaternB, r.data.QuaternC, r.data.QuaternD = n2Header.QuaternB, n2Header.QuaternC, n2Header.QuaternD
		r.data.QoffsetX, r.data.QoffsetY, r.data.QoffsetZ = n2Header.QoffsetX, n2Header.QoffsetY, n2Header.QoffsetZ
	}

	// Fix bad value in header
	if r.data.Nz <= 0 && r.data.Dim[3] <= 0 {
		r.data.Nz = 1
		r.data.Dim[3] = 1
	}
	if r.data.Nt <= 0 && r.data.Dim[4] <= 0 {
		r.data.Nt = 1
		r.data.Dim[4] = 1
	}
	if r.data.Nu <= 0 && r.data.Dim[5] <= 0 {
		r.data.Nu = 1
		r.data.Dim[5] = 1
	}
	if r.data.Nv <= 0 && r.data.Dim[6] <= 0 {
		r.data.Nv = 1
		r.data.Dim[6] = 1
	}
	if r.data.Nw <= 0 && r.data.Dim[7] <= 0 {
		r.data.Nw = 1
		r.data.Dim[7] = 1
	}

	// Set the byte order
	r.data.ByteOrder = r.binaryOrder

	if bitpix == 0 {
		return errors.New("number of bits per voxel value (bitpix) is zero")
	}

	r.data.NVox = 1
	for i := int64(1); i <= r.data.NDim; i++ {
		r.data.NVox *= r.data.Dim[i]
	}

	// compute QToXYK transformation from pixel indexes (i,j,k) to (x,y,z)
	if qFormCode <= 0 {
		r.data.QtoXYZ.M[0][0] = r.data.Dx
		r.data.QtoXYZ.M[1][1] = r.data.Dy
		r.data.QtoXYZ.M[2][2] = r.data.Dz

		// off diagonal is zero
		r.data.QtoXYZ.M[0][1] = 0
		r.data.QtoXYZ.M[0][2] = 0
		r.data.QtoXYZ.M[0][3] = 0

		r.data.QtoXYZ.M[1][0] = 0
		r.data.QtoXYZ.M[1][2] = 0
		r.data.QtoXYZ.M[1][3] = 0

		r.data.QtoXYZ.M[2][0] = 0
		r.data.QtoXYZ.M[2][1] = 0
		r.data.QtoXYZ.M[2][3] = 0

		// last row is [0, 0, 0, 1]
		r.data.QtoXYZ.M[3][0] = 0
		r.data.QtoXYZ.M[3][1] = 0
		r.data.QtoXYZ.M[3][2] = 0
		r.data.QtoXYZ.M[3][3] = 1.0

		r.data.QformCode = constant.NIFTI_XFORM_UNKNOWN
	} else {
		if pixDim0 < 0 {
			r.data.QFac = -1
		} else {
			r.data.QFac = 1
		}
		r.data.QtoXYZ = r.QuaternToMatrix()
		r.data.QformCode = qFormCode
	}

	// Set QToIJK
	r.data.QtoIJK = matrix.Mat44Inverse(r.data.QtoXYZ)

	if sFormCode <= 0 {
		r.data.SformCode = constant.NIFTI_XFORM_UNKNOWN
	} else {
		r.data.StoXYZ.M[0][0] = sRowX[0]
		r.data.StoXYZ.M[0][1] = sRowX[1]
		r.data.StoXYZ.M[0][2] = sRowX[2]
		r.data.StoXYZ.M[0][3] = sRowX[3]

		r.data.StoXYZ.M[1][0] = sRowY[0]
		r.data.StoXYZ.M[1][1] = sRowY[1]
		r.data.StoXYZ.M[1][2] = sRowY[2]
		r.data.StoXYZ.M[1][3] = sRowY[3]

		r.data.StoXYZ.M[2][0] = sRowZ[0]
		r.data.StoXYZ.M[2][1] = sRowZ[0]
		r.data.StoXYZ.M[2][2] = sRowZ[0]
		r.data.StoXYZ.M[2][3] = sRowZ[0]

		r.data.StoXYZ.M[3][0] = 0
		r.data.StoXYZ.M[3][1] = 0
		r.data.StoXYZ.M[3][2] = 0
		r.data.StoXYZ.M[3][3] = 1

		r.data.StoIJK = matrix.Mat44Inverse(r.data.StoXYZ)

		r.data.SformCode = sFormCode
	}

	// Other stuff
	r.data.SclSlope = sclSlope
	r.data.SclInter = sclInter

	r.data.IntentName = intentName
	r.data.IntentCode = intentCode
	r.data.IntentP1 = intentP1
	r.data.IntentP2 = intentP2
	r.data.IntentP3 = intentP3

	r.data.QuaternB = quaternB
	r.data.QuaternC = quaternC
	r.data.QuaternD = quaternD
	r.data.Descrip = descrip

	// Frequency dimension, phase dimension, slice dimension
	r.data.FreqDim = freqDim
	r.data.PhaseDim = phaseDim
	r.data.SliceDim = sliceDim

	r.data.SliceCode = sliceCode
	r.data.SliceStart = sliceStart
	r.data.SliceEnd = sliceEnd
	r.data.SliceDuration = sliceDuration

	r.data.CalMin = calMin
	r.data.CalMax = calMax

	r.data.Datatype = datatype

	if r.data.Dim[5] > 1 {
		statDim = r.data.Dim[5]
	}
	offset = voxOffset
	dataSize := r.data.Dim[1] * r.data.Dim[2] * r.data.Dim[3] * r.data.Dim[4] * statDim * (int64(bitpix) / 8)

	_, err := r.reader.Seek(offset, 0)
	if err != nil {
		return err
	}

	buf := make([]byte, dataSize)
	_, err = io.ReadFull(r.reader, buf)
	if err != nil {
		return err
	}
	r.data.Volume = buf

	affine := matrix.DMat44{}
	affine.M[0] = sRowX
	affine.M[1] = sRowY
	affine.M[2] = sRowZ
	affine.M[3] = [4]float64{0, 0, 0, 1}

	r.data.Affine = affine
	r.data.matrixToOrientation(affine)

	return nil
}
