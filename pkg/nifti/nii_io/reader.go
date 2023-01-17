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
	data        *Nii             // Contains the NIFTI data structure
	version     int              // Define the version of NIFTI image (1 or 2)
}

// NewNiiReader receives a path to the NIFTI file and returns a new reader to parse the file
//
// TODO: this is not efficient when the file is large so we need to find better way to deal with large file size
func NewNiiReader(filePath string) (NiiReader, error) {
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

	niiData := &Nii{}
	niiData.Data = &NiiData{}

	return &niiReader{
		binaryOrder: binary.LittleEndian,
		reader:      bytes.NewReader(bData),
		data:        niiData,
	}, nil
}

// niftiVersion checks the header to determine the NIFTI version
func (r *niiReader) niftiVersion() error {
	var hSize int32

	err := binary.Read(r.reader, r.binaryOrder, &hSize)
	if err != nil {
		return err
	}

	switch hSize {
	case NII1HeaderSize:
		r.version = NIIVersion1
	case NII2HeaderSize:
		r.version = NIIVersion2
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
		case NII1HeaderSize:
			r.version = NIIVersion1
		case NII2HeaderSize:
			r.version = NIIVersion2
		default:
			return errors.New("invalid NIFTI file format")
		}
	}
	r.data.Data.Version = r.version
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
	err := r.niftiVersion()
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
	case NIIVersion1:
		header := new(Nii1Header)
		err = binary.Read(r.reader, r.binaryOrder, header)
		if err != nil {
			return err
		}
		if header.Magic != [4]byte{110, 43, 49, 0} && header.Magic != [4]byte{110, 105, 49, 0} {
			return errors.New("invalid NIFTI-1 magic string")
		}
		r.data.n1Header = header
		dim0 = int64(header.Dim[0])
	case NIIVersion2:
		header := new(Nii2Header)
		err = binary.Read(r.reader, r.binaryOrder, header)
		if err != nil {
			return err
		}
		if header.Magic != [8]byte{110, 43, 50, 0, 13, 10, 26, 10} {
			return errors.New("invalid NIFTI-2 magic string")
		}
		r.data.n2Header = header
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

// setDatatypeSize sets number of bytes per voxel and the swap size for the header datatype
func (r *niiReader) setDatatypeSize() {
	var datatype int32

	switch r.version {
	case NIIVersion1:
		datatype = int32(r.data.n1Header.Datatype)
	case NIIVersion2:
		datatype = int32(r.data.n2Header.Datatype)
	}
	NByPerVoxel, SwapSize := assignDatatypeSize(datatype)
	r.data.Data.NByPer = int32(NByPerVoxel)
	r.data.Data.SwapSize = int32(SwapSize)
}

// parseData parse the raw byte array into NIFTI-1 or NIFTI-2 data structure
func (r *niiReader) parseData() error {
	r.data.Data = &NiiData{}
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
	case NIIVersion1:
		freqDim = int32(dimInfoToFreqDim(r.data.n1Header.DimInfo))
		phaseDim = int32(dimInfoToPhaseDim(r.data.n1Header.DimInfo))
		sliceDim = int32(dimInfoToSliceDim(r.data.n1Header.DimInfo))

		voxOffset = int64(r.data.n1Header.VoxOffset)
		datatype = int32(r.data.n1Header.Datatype)

		// The bits 1-3 are used to store the spatial dimensions, the bits 4-6 are for temporal dimensions,
		// and the bits 6 and 7 are not used
		r.data.Data.XYZUnits = int32(r.data.n1Header.XyztUnits % 8)
		r.data.Data.TimeUnits = int32(r.data.n1Header.XyztUnits) - r.data.Data.XYZUnits

		sliceCode = int32(r.data.n1Header.SliceCode)
		sliceStart = int64(r.data.n1Header.SliceStart)
		sliceEnd = int64(r.data.n1Header.SliceEnd)
		sliceDuration = float64(r.data.n1Header.SliceDuration)

		calMin = float64(r.data.n1Header.CalMin)
		calMax = float64(r.data.n1Header.CalMax)

		qFormCode = int32(r.data.n1Header.QformCode)
		sFormCode = int32(r.data.n1Header.SformCode)
		pixDim0 = float64(r.data.n1Header.Pixdim[0])

		sRowX = convertToF64(r.data.n1Header.SrowX)
		sRowY = convertToF64(r.data.n1Header.SrowY)
		sRowZ = convertToF64(r.data.n1Header.SrowZ)

		sclSlope = float64(r.data.n1Header.SclSlope)
		sclInter = float64(r.data.n1Header.SclInter)

		intentName = r.data.n1Header.IntentName
		intentCode = int32(r.data.n1Header.IntentCode)
		intentP1 = float64(r.data.n1Header.IntentP1)
		intentP2 = float64(r.data.n1Header.IntentP2)
		intentP3 = float64(r.data.n1Header.IntentP3)

		quaternB = float64(r.data.n1Header.QuaternB)
		quaternC = float64(r.data.n1Header.QuaternC)
		quaternD = float64(r.data.n1Header.QuaternD)
		descrip = r.data.n1Header.Descrip

		// Set the dimension of data array
		r.data.Data.NDim, r.data.Data.Dim[0] = int64(r.data.n1Header.Dim[0]), int64(r.data.n1Header.Dim[0])
		r.data.Data.Nx, r.data.Data.Dim[1] = int64(r.data.n1Header.Dim[1]), int64(r.data.n1Header.Dim[1])
		r.data.Data.Ny, r.data.Data.Dim[2] = int64(r.data.n1Header.Dim[2]), int64(r.data.n1Header.Dim[2])
		r.data.Data.Nz, r.data.Data.Dim[3] = int64(r.data.n1Header.Dim[3]), int64(r.data.n1Header.Dim[3])
		r.data.Data.Nt, r.data.Data.Dim[4] = int64(r.data.n1Header.Dim[4]), int64(r.data.n1Header.Dim[4])
		r.data.Data.Nu, r.data.Data.Dim[5] = int64(r.data.n1Header.Dim[5]), int64(r.data.n1Header.Dim[5])
		r.data.Data.Nv, r.data.Data.Dim[6] = int64(r.data.n1Header.Dim[6]), int64(r.data.n1Header.Dim[6])
		r.data.Data.Nw, r.data.Data.Dim[7] = int64(r.data.n1Header.Dim[7]), int64(r.data.n1Header.Dim[7])

		// Set the grid spacing
		r.data.Data.Dx, r.data.Data.PixDim[1] = float64(r.data.n1Header.Pixdim[1]), float64(r.data.n1Header.Pixdim[1])
		r.data.Data.Dy, r.data.Data.PixDim[2] = float64(r.data.n1Header.Pixdim[2]), float64(r.data.n1Header.Pixdim[2])
		r.data.Data.Dz, r.data.Data.PixDim[3] = float64(r.data.n1Header.Pixdim[3]), float64(r.data.n1Header.Pixdim[3])
		r.data.Data.Dt, r.data.Data.PixDim[4] = float64(r.data.n1Header.Pixdim[4]), float64(r.data.n1Header.Pixdim[4])
		r.data.Data.Du, r.data.Data.PixDim[5] = float64(r.data.n1Header.Pixdim[5]), float64(r.data.n1Header.Pixdim[5])
		r.data.Data.Dv, r.data.Data.PixDim[6] = float64(r.data.n1Header.Pixdim[6]), float64(r.data.n1Header.Pixdim[6])
		r.data.Data.Dw, r.data.Data.PixDim[7] = float64(r.data.n1Header.Pixdim[7]), float64(r.data.n1Header.Pixdim[7])

		bitpix = r.data.n1Header.Bitpix

	case NIIVersion2:
		freqDim = int32(dimInfoToFreqDim(r.data.n2Header.DimInfo))
		phaseDim = int32(dimInfoToPhaseDim(r.data.n2Header.DimInfo))
		sliceDim = int32(dimInfoToSliceDim(r.data.n2Header.DimInfo))

		voxOffset = r.data.n2Header.VoxOffset
		datatype = int32(r.data.n2Header.Datatype)

		// The bits 1-3 are used to store the spatial dimensions, the bits 4-6 are for temporal dimensions,
		// and the bits 6 and 7 are not used
		r.data.Data.XYZUnits = r.data.n2Header.XyztUnits % 8
		r.data.Data.TimeUnits = r.data.n2Header.XyztUnits - r.data.Data.XYZUnits

		sliceCode = r.data.n2Header.SliceCode
		sliceStart = r.data.n2Header.SliceStart
		sliceEnd = r.data.n2Header.SliceEnd
		sliceDuration = r.data.n2Header.SliceDuration

		calMin = r.data.n2Header.CalMin
		calMax = r.data.n2Header.CalMax

		qFormCode = r.data.n2Header.QformCode
		pixDim0 = r.data.n2Header.Pixdim[0]
		sFormCode = r.data.n2Header.SformCode

		sclSlope = r.data.n2Header.SclSlope
		sclInter = r.data.n2Header.SclInter

		intentName = r.data.n2Header.IntentName
		intentCode = r.data.n2Header.IntentCode
		r.data.Data.IntentP1 = r.data.n2Header.IntentP1
		r.data.Data.IntentP2 = r.data.n2Header.IntentP2
		r.data.Data.IntentP3 = r.data.n2Header.IntentP3

		r.data.Data.QuaternB = r.data.n2Header.QuaternB
		r.data.Data.QuaternC = r.data.n2Header.QuaternC
		r.data.Data.QuaternD = r.data.n2Header.QuaternD
		descrip = r.data.n2Header.Descrip

		// Set the dimension of data array
		r.data.Data.NDim, r.data.Data.Dim[0] = r.data.n2Header.Dim[0], r.data.n2Header.Dim[0]
		r.data.Data.Nx, r.data.Data.Dim[1] = r.data.n2Header.Dim[1], r.data.n2Header.Dim[1]
		r.data.Data.Ny, r.data.Data.Dim[2] = r.data.n2Header.Dim[2], r.data.n2Header.Dim[2]
		r.data.Data.Nz, r.data.Data.Dim[3] = r.data.n2Header.Dim[3], r.data.n2Header.Dim[3]
		r.data.Data.Nt, r.data.Data.Dim[4] = r.data.n2Header.Dim[4], r.data.n2Header.Dim[4]
		r.data.Data.Nu, r.data.Data.Dim[5] = r.data.n2Header.Dim[5], r.data.n2Header.Dim[5]
		r.data.Data.Nv, r.data.Data.Dim[6] = r.data.n2Header.Dim[6], r.data.n2Header.Dim[6]
		r.data.Data.Nw, r.data.Data.Dim[7] = r.data.n2Header.Dim[7], r.data.n2Header.Dim[7]

		// Set the grid spacing
		r.data.Data.Dx, r.data.Data.PixDim[1] = r.data.n2Header.Pixdim[1], r.data.n2Header.Pixdim[1]
		r.data.Data.Dy, r.data.Data.PixDim[2] = r.data.n2Header.Pixdim[2], r.data.n2Header.Pixdim[2]
		r.data.Data.Dz, r.data.Data.PixDim[3] = r.data.n2Header.Pixdim[3], r.data.n2Header.Pixdim[3]
		r.data.Data.Dt, r.data.Data.PixDim[4] = r.data.n2Header.Pixdim[4], r.data.n2Header.Pixdim[4]
		r.data.Data.Du, r.data.Data.PixDim[5] = r.data.n2Header.Pixdim[5], r.data.n2Header.Pixdim[5]
		r.data.Data.Dv, r.data.Data.PixDim[6] = r.data.n2Header.Pixdim[6], r.data.n2Header.Pixdim[6]
		r.data.Data.Dw, r.data.Data.PixDim[7] = r.data.n2Header.Pixdim[7], r.data.n2Header.Pixdim[7]

		bitpix = r.data.n2Header.Bitpix

		// SRowX, SRowY, SRowZ
		sRowX, sRowY, sRowZ = r.data.n2Header.SrowX, r.data.n2Header.SrowY, r.data.n2Header.SrowZ
	}

	// Fix bad value in header
	if r.data.Data.Nz <= 0 && r.data.Data.Dim[3] <= 0 {
		r.data.Data.Nz = 1
		r.data.Data.Dim[3] = 1
	}
	if r.data.Data.Nt <= 0 && r.data.Data.Dim[4] <= 0 {
		r.data.Data.Nt = 1
		r.data.Data.Dim[4] = 1
	}
	if r.data.Data.Nu <= 0 && r.data.Data.Dim[5] <= 0 {
		r.data.Data.Nu = 1
		r.data.Data.Dim[5] = 1
	}
	if r.data.Data.Nv <= 0 && r.data.Data.Dim[6] <= 0 {
		r.data.Data.Nv = 1
		r.data.Data.Dim[6] = 1
	}
	if r.data.Data.Nw <= 0 && r.data.Data.Dim[7] <= 0 {
		r.data.Data.Nw = 1
		r.data.Data.Dim[7] = 1
	}

	// Set the byte order
	r.data.Data.ByteOrder = r.binaryOrder

	if bitpix == 0 {
		return errors.New("number of bits per voxel value (bitpix) is zero")
	}

	r.data.Data.NVox = 1
	for i := int64(1); i <= r.data.Data.NDim; i++ {
		r.data.Data.NVox *= r.data.Data.Dim[i]
	}

	r.setDatatypeSize()

	// compute QToXYK transformation from pixel indexes (i,j,k) to (x,y,z)
	if qFormCode <= 0 {
		r.data.Data.QtoXYZ.M[0][0] = r.data.Data.Dx
		r.data.Data.QtoXYZ.M[1][1] = r.data.Data.Dy
		r.data.Data.QtoXYZ.M[2][2] = r.data.Data.Dz

		// off diagonal is zero
		r.data.Data.QtoXYZ.M[0][1] = 0
		r.data.Data.QtoXYZ.M[0][2] = 0
		r.data.Data.QtoXYZ.M[0][3] = 0

		r.data.Data.QtoXYZ.M[1][0] = 0
		r.data.Data.QtoXYZ.M[1][2] = 0
		r.data.Data.QtoXYZ.M[1][3] = 0

		r.data.Data.QtoXYZ.M[2][0] = 0
		r.data.Data.QtoXYZ.M[2][1] = 0
		r.data.Data.QtoXYZ.M[2][3] = 0

		// last row is [0, 0, 0, 1]
		r.data.Data.QtoXYZ.M[3][0] = 0
		r.data.Data.QtoXYZ.M[3][1] = 0
		r.data.Data.QtoXYZ.M[3][2] = 0
		r.data.Data.QtoXYZ.M[3][3] = 1.0

		r.data.Data.QformCode = constant.NIFTI_XFORM_UNKNOWN
	} else {
		if pixDim0 < 0 {
			r.data.Data.QFac = -1
		} else {
			r.data.Data.QFac = 1
		}
		r.data.Data.QtoXYZ = r.QuaternToMatrix()
		r.data.Data.QformCode = qFormCode
	}

	// Set QToIJK
	r.data.Data.QtoIJK = matrix.Mat44Inverse(r.data.Data.QtoXYZ)

	if sFormCode <= 0 {
		r.data.Data.SformCode = constant.NIFTI_XFORM_UNKNOWN
	} else {
		r.data.Data.StoXYZ.M[0][0] = sRowX[0]
		r.data.Data.StoXYZ.M[0][1] = sRowX[1]
		r.data.Data.StoXYZ.M[0][2] = sRowX[2]
		r.data.Data.StoXYZ.M[0][3] = sRowX[3]

		r.data.Data.StoXYZ.M[1][0] = sRowY[0]
		r.data.Data.StoXYZ.M[1][1] = sRowY[1]
		r.data.Data.StoXYZ.M[1][2] = sRowY[2]
		r.data.Data.StoXYZ.M[1][3] = sRowY[3]

		r.data.Data.StoXYZ.M[2][0] = sRowZ[0]
		r.data.Data.StoXYZ.M[2][1] = sRowZ[0]
		r.data.Data.StoXYZ.M[2][2] = sRowZ[0]
		r.data.Data.StoXYZ.M[2][3] = sRowZ[0]

		r.data.Data.StoXYZ.M[3][0] = 0
		r.data.Data.StoXYZ.M[3][1] = 0
		r.data.Data.StoXYZ.M[3][2] = 0
		r.data.Data.StoXYZ.M[3][3] = 1

		r.data.Data.StoIJK = matrix.Mat44Inverse(r.data.Data.StoXYZ)

		r.data.Data.SformCode = sFormCode
	}

	// Other stuff
	r.data.Data.SclSlope = sclSlope
	r.data.Data.SclInter = sclInter

	r.data.Data.IntentName = intentName
	r.data.Data.IntentCode = intentCode
	r.data.Data.IntentP1 = intentP1
	r.data.Data.IntentP2 = intentP2
	r.data.Data.IntentP3 = intentP3

	r.data.Data.QuaternB = quaternB
	r.data.Data.QuaternC = quaternC
	r.data.Data.QuaternD = quaternD
	r.data.Data.Descrip = descrip

	// Frequency dimension, phase dimension, slice dimension
	r.data.Data.FreqDim = freqDim
	r.data.Data.PhaseDim = phaseDim
	r.data.Data.SliceDim = sliceDim

	r.data.Data.SliceCode = sliceCode
	r.data.Data.SliceStart = sliceStart
	r.data.Data.SliceEnd = sliceEnd
	r.data.Data.SliceDuration = sliceDuration

	r.data.Data.CalMin = calMin
	r.data.Data.CalMax = calMax

	r.data.Data.Datatype = datatype

	if r.data.Data.Dim[5] > 1 {
		statDim = r.data.Data.Dim[5]
	}
	offset = voxOffset
	dataSize := r.data.Data.Dim[1] * r.data.Data.Dim[2] * r.data.Data.Dim[3] * r.data.Data.Dim[4] * statDim * (int64(bitpix) / 8)

	_, err := r.reader.Seek(int64(offset), 0)
	if err != nil {
		return err
	}

	buf := make([]byte, dataSize)
	_, err = io.ReadFull(r.reader, buf)
	if err != nil {
		return err
	}
	r.data.Data.Data = buf

	affine := matrix.DMat44{}
	affine.M[0] = sRowX
	affine.M[1] = sRowY
	affine.M[2] = sRowZ
	affine.M[3] = [4]float64{0, 0, 0, 1}

	r.data.Data.Affine = affine
	r.data.matrixToOrientation(affine)

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
