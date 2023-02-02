package nii_io

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/okieraised/go2com/internal/system"
	"github.com/okieraised/go2com/internal/utils"
	"github.com/okieraised/go2com/pkg/nifti/constant"
	"math"
	"net/http"
)

// IsValidDatatype checks whether the datatype is valid for NIFTI format
func IsValidDatatype(datatype int32) bool {
	if constant.ValidDatatype[datatype] {
		return true
	}
	return false
}

// swapNIFTI1Header swaps all NIFTI fields
func swapNIFTI1Header(header *Nii1Header) (*Nii1Header, error) {
	newHeader := new(Nii1Header)
	var err error

	newHeader.SizeofHdr = swapInt32(header.SizeofHdr)
	newHeader.Extents = swapInt32(header.Extents)
	newHeader.SessionError = swapInt16(header.SessionError)
	for i := 0; i < 8; i++ {
		newHeader.Dim[i] = swapInt16(header.Dim[i])
	}

	newHeader.IntentP1, err = swapFloat32(header.IntentP1)
	if err != nil {
		return nil, fmt.Errorf("failed to byte swap IntentP1: %v", err)
	}
	newHeader.IntentP2, err = swapFloat32(header.IntentP2)
	if err != nil {
		return nil, fmt.Errorf("failed to byte swap IntentP2: %v", err)
	}
	newHeader.IntentP3, err = swapFloat32(header.IntentP3)
	if err != nil {
		return nil, fmt.Errorf("failed to byte swap IntentP3: %v", err)
	}

	newHeader.IntentCode = swapInt16(header.IntentCode)
	newHeader.Datatype = swapInt16(header.Datatype)
	newHeader.Bitpix = swapInt16(header.Bitpix)
	newHeader.SliceStart = swapInt16(header.SliceStart)

	for i := 0; i < 8; i++ {
		newHeader.Pixdim[i], err = swapFloat32(header.Pixdim[i])
		if err != nil {
			return nil, fmt.Errorf("failed to byte swap Pixdim[%d]: %v", i, err)
		}
	}

	newHeader.VoxOffset, err = swapFloat32(header.VoxOffset)
	if err != nil {
		return nil, fmt.Errorf("failed to byte swap VoxOffset: %v", err)
	}

	newHeader.SclSlope, err = swapFloat32(header.SclSlope)
	if err != nil {
		return nil, fmt.Errorf("failed to byte swap SclSlope: %v", err)
	}

	newHeader.SclInter, err = swapFloat32(header.SclInter)
	if err != nil {
		return nil, fmt.Errorf("failed to byte swap SclInter: %v", err)
	}

	newHeader.SliceEnd = swapInt16(header.SliceEnd)

	newHeader.CalMin, err = swapFloat32(header.CalMin)
	if err != nil {
		return nil, fmt.Errorf("failed to byte swap CalMin: %v", err)
	}

	newHeader.CalMax, err = swapFloat32(header.CalMax)
	if err != nil {
		return nil, fmt.Errorf("failed to byte swap CalMax: %v", err)
	}

	newHeader.SliceDuration, err = swapFloat32(header.SliceDuration)
	if err != nil {
		return nil, fmt.Errorf("failed to byte swap SliceDuration: %v", err)
	}

	newHeader.Glmin = swapInt32(header.Glmin)
	newHeader.Glmax = swapInt32(header.Glmax)

	newHeader.QformCode = swapInt16(header.QformCode)
	newHeader.SformCode = swapInt16(header.SformCode)

	newHeader.QuaternB, err = swapFloat32(header.QuaternB)
	if err != nil {
		return nil, fmt.Errorf("failed to byte swap QuaternB: %v", err)
	}

	newHeader.QuaternC, err = swapFloat32(header.QuaternC)
	if err != nil {
		return nil, fmt.Errorf("failed to byte swap QuaternC: %v", err)
	}

	newHeader.QuaternD, err = swapFloat32(header.QuaternD)
	if err != nil {
		return nil, fmt.Errorf("failed to byte swap QuaternD: %v", err)
	}

	newHeader.QoffsetX, err = swapFloat32(header.QoffsetX)
	if err != nil {
		return nil, fmt.Errorf("failed to byte swap QoffsetX: %v", err)
	}
	newHeader.QoffsetY, err = swapFloat32(header.QoffsetY)
	if err != nil {
		return nil, fmt.Errorf("failed to byte swap QoffsetY: %v", err)
	}
	newHeader.QoffsetZ, err = swapFloat32(header.QoffsetZ)
	if err != nil {
		return nil, fmt.Errorf("failed to byte swap QoffsetZ: %v", err)
	}

	for i := 0; i < 4; i++ {
		newHeader.SrowX[i], err = swapFloat32(header.SrowX[i])
		if err != nil {
			return nil, fmt.Errorf("failed to byte swap SrowX[%d]: %v", i, err)
		}
		newHeader.SrowY[i], err = swapFloat32(header.SrowY[i])
		if err != nil {
			return nil, fmt.Errorf("failed to byte swap SrowY[%d]: %v", i, err)
		}
		newHeader.SrowZ[i], err = swapFloat32(header.SrowZ[i])
		if err != nil {
			return nil, fmt.Errorf("failed to byte swap SrowZ[%d]: %v", i, err)
		}
	}
	return newHeader, nil
}

// getDatatype returns the appropriate datatype of the NIFTI image
func getDatatype(datatype int32) string {
	switch datatype {
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
	return ILLEGAL
}

// getSliceCode returns the name of the slice code
func getSliceCode(sliceCode int32) string {
	switch sliceCode {
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

	return "UNKNOWN"
}

// assignDatatypeSize sets the number of bytes per voxel and the swapsize based on a datatype code
// returns nByper and swapSize
func assignDatatypeSize(datatype int32) (int16, int16) {
	var nByper, swapSize int16
	switch datatype {
	case constant.DT_INT8, constant.DT_UINT8:
		nByper = 1
		swapSize = 0
	case constant.DT_INT16, constant.DT_UINT16:
		nByper = 2
		swapSize = 2
	case constant.DT_RGB24:
		nByper = 3
		swapSize = 0
	case constant.DT_RGBA32:
		nByper = 4
		swapSize = 0
	case constant.DT_INT32, constant.DT_UINT32, constant.DT_FLOAT32:
		nByper = 4
		swapSize = 4
	case constant.DT_COMPLEX64:
		nByper = 8
		swapSize = 4
	case constant.DT_FLOAT64, constant.DT_INT64, constant.DT_UINT64:
		nByper = 8
		swapSize = 8
	case constant.DT_FLOAT128:
		nByper = 16
		swapSize = 16
	case constant.DT_COMPLEX128:
		nByper = 16
		swapSize = 8
	case constant.DT_COMPLEX256:
		nByper = 32
		swapSize = 16
	default:
	}
	return nByper, swapSize
}

// needHeaderSwap checks whether byte swapping is needed. dim0 should be in [0,7], and headerSize should be accurate.
//
// Returns:
//
// > 0 : needs swap
//
// = 0 : does not need swap
//
// < 0 : error condition
func needHeaderSwap(dim0 int16) int {
	d0 := dim0
	if d0 != 0 {
		if d0 > 0 && d0 < 7 {
			return 0
		}

		d0 = swapInt16(d0)
		if d0 > 0 && d0 < 7 {
			return 1
		}
		return -1
	}
	return -2
}

// swapInt16 swaps int16 from native endian to the other
func swapInt16(in int16) int16 {
	b := make([]byte, 2)

	switch system.NativeEndian {
	case binary.LittleEndian:
		binary.LittleEndian.PutUint16(b, uint16(in))
		return int16(binary.BigEndian.Uint16(b))
	default:
		binary.BigEndian.PutUint16(b, uint16(in))
		return int16(binary.LittleEndian.Uint16(b))
	}
}

// swapInt32 swaps int32 from native endian to the other
func swapInt32(in int32) int32 {
	b := make([]byte, 4)

	switch system.NativeEndian {
	case binary.LittleEndian:
		binary.LittleEndian.PutUint32(b, uint32(in))
		return int32(binary.BigEndian.Uint16(b))
	default:
		binary.BigEndian.PutUint32(b, uint32(in))
		return int32(binary.LittleEndian.Uint32(b))
	}
}

// swapInt64 swaps int64 from native endian to the other
func swapInt64(in int64) int64 {
	b := make([]byte, 8)

	switch system.NativeEndian {
	case binary.LittleEndian:
		binary.LittleEndian.PutUint64(b, uint64(in))
		return int64(binary.BigEndian.Uint64(b))
	default:
		binary.BigEndian.PutUint64(b, uint64(in))
		return int64(binary.LittleEndian.Uint64(b))
	}
}

// swapFloat32 swaps float32 from native endian to the other
func swapFloat32(in float32) (float32, error) {
	buf := new(bytes.Buffer)

	switch system.NativeEndian {
	case binary.LittleEndian:
		err := binary.Write(buf, binary.LittleEndian, in)
		if err != nil {
			return 0, err
		}
		bits := binary.BigEndian.Uint32(buf.Bytes())
		res := math.Float32frombits(bits)
		return res, nil
	default:
		err := binary.Write(buf, binary.BigEndian, in)
		if err != nil {
			return 0, err
		}
		bits := binary.LittleEndian.Uint32(buf.Bytes())
		res := math.Float32frombits(bits)
		return res, nil
	}
}

// swapFloat64 swaps float64 from native endian to the other
func swapFloat64(in float64) (float64, error) {
	buf := new(bytes.Buffer)

	switch system.NativeEndian {
	case binary.LittleEndian:
		err := binary.Write(buf, binary.LittleEndian, in)
		if err != nil {
			return 0, err
		}
		bits := binary.BigEndian.Uint64(buf.Bytes())
		res := math.Float64frombits(bits)
		return res, nil
	default:
		err := binary.Write(buf, binary.BigEndian, in)
		if err != nil {
			return 0, err
		}
		bits := binary.LittleEndian.Uint64(buf.Bytes())
		res := math.Float64frombits(bits)
		return res, nil
	}
}

func convertToF64(ar [4]float32) [4]float64 {
	newar := [4]float64{}
	var v float32
	var i int
	for i, v = range ar {
		newar[i] = float64(v)
	}
	return newar
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

// convertSpaceTimeToXYZT converts xyzUnit, timeUnit back to uint8 representation of XyztUnits field
func convertSpaceTimeToXYZT(xyzUnit, timeUnit int32) uint8 {
	return uint8((xyzUnit & 0x07) | (timeUnit & 0x38))
}

// convertFPSIntoDimInfo converts freqDim, phaseDim, sliceDim back to uint8 representation of DimInfo
func convertFPSIntoDimInfo(freqDim, phaseDim, sliceDim int32) uint8 {
	return uint8((freqDim & 0x03) | ((phaseDim & 0x03) << 2) | ((sliceDim & 0x03) << 4))
}

// Check for valid extension
func validNIfTIFileExt(filePath string) {

}

// deflateFileContent deflates the gzipped binary to its original content
func deflateFileContent(bData []byte) ([]byte, error) {
	var err error
	mimeType := http.DetectContentType(bData[:512])
	if mimeType == "application/x-gzip" {
		bData, err = utils.DeflateGzip(bData)
		if err != nil {
			return nil, err
		}
	}
	return bData, nil
}

func MakeNewNii1Header(inDim *[8]int16, inDatatype int32) *Nii1Header {

	// Default Dim value
	defaultDim := [8]int16{3, 1, 1, 1, 0, 0, 0, 0}

	header := new(Nii1Header)
	var dim [8]int16

	// If no input Dim is provided then we use the default value
	if inDim != nil {
		dim = *inDim
	} else {
		dim = defaultDim
	}

	// validate Dim: if there is any problem, apply default Dim
	if dim[0] < 0 || dim[0] > 7 {
		dim = defaultDim
	} else {
		for c := 1; c <= int(dim[0]); c++ {
			if dim[c] < 1 {
				fmt.Printf("bad dim: %d: %d\n", c, dim[c])
				dim = defaultDim
				break
			}
		}
	}

	// Validate datatype
	datatype := inDatatype
	if !IsValidDatatype(datatype) {
		datatype = constant.DT_FLOAT32
	}

	// Populate the header struct
	header.SizeofHdr = constant.NII1HeaderSize
	header.Regular = 'r'

	// Init dim and pixdim
	header.Dim[0] = dim[0]
	header.Pixdim[0] = 0.0
	for c := 1; c <= int(dim[0]); c++ {
		header.Dim[c] = dim[c]
		header.Pixdim[c] = 1.0
	}

	header.Datatype = int16(datatype)

	nByper, _ := assignDatatypeSize(datatype)
	header.Bitpix = 8 * nByper
	header.Magic = [4]byte{110, 43, 49, 0}

	return header
}

func MakeNewNii2Header(inDim *[8]int64, inDatatype int32) *Nii2Header {
	// Default Dim value
	defaultDim := [8]int64{3, 1, 1, 1, 0, 0, 0, 0}

	header := new(Nii2Header)
	var dim [8]int64

	// If no input Dim is provided then we use the default value
	if inDim != nil {
		dim = *inDim
	} else {
		dim = defaultDim
	}

	// validate Dim: if there is any problem, apply default Dim
	if dim[0] < 0 || dim[0] > 7 {
		dim = defaultDim
	} else {
		for c := 1; c <= int(dim[0]); c++ {
			if dim[c] < 1 {
				fmt.Printf("bad dim: %d: %d\n", c, dim[c])
				dim = defaultDim
				break
			}
		}
	}

	// Validate datatype
	datatype := inDatatype
	if !IsValidDatatype(datatype) {
		datatype = constant.DT_FLOAT32
	}

	// Populate the header struct
	header.SizeofHdr = constant.NII2HeaderSize

	// Init dim and pixdim
	header.Dim[0] = dim[0]
	header.Pixdim[0] = 0.0
	for c := 1; c <= int(dim[0]); c++ {
		header.Dim[c] = dim[c]
		header.Pixdim[c] = 1.0
	}

	header.Datatype = int16(datatype)

	nByper, _ := assignDatatypeSize(datatype)
	header.Bitpix = 8 * nByper
	header.Magic = [8]byte{110, 105, 50, 0, 13, 10, 26, 10}

	return header
}

// MakeEmptyImageFromImg returns a zero-filled byte slice from existing Nii image structure
func MakeEmptyImageFromImg(img *Nii) ([]byte, error) {
	var bDataLength int64

	if img == nil {
		return nil, errors.New("NIfTI image structure nil")
	}

	// Need at least nx, ny
	if img.Nx == 0 {
		return nil, errors.New("x dimension must not be zero")
	}
	if img.Ny == 0 {
		return nil, errors.New("y dimension must not be zero")
	}
	bDataLength = img.Nx * img.Ny

	if img.Nz > 0 {
		bDataLength = bDataLength * img.Nz
	}
	if img.Nt > 0 {
		bDataLength = bDataLength * img.Nt
	}
	if img.Nu > 0 {
		bDataLength = bDataLength * img.Nu
	}
	if img.Nv > 0 {
		bDataLength = bDataLength * img.Nv
	}
	if img.Nw > 0 {
		bDataLength = bDataLength * img.Nw
	}

	nByper, _ := assignDatatypeSize(img.Datatype)
	bDataLength = bDataLength * int64(nByper)

	// Init a slice of bytes with capacity of bDataLength and initial value of 0
	bData := make([]byte, bDataLength, bDataLength)

	return bData, nil
}

// MakeEmptyImageFromHdr initializes a zero-filled byte slice from existing header structure
func MakeEmptyImageFromHdr(hdr *Nii1Header) ([]byte, error) {
	var bDataLength int64

	if hdr == nil {
		return nil, errors.New("NIfTI image structure nil")
	}

	if hdr.Dim[1] == 0 {
		return nil, errors.New("x dimension must not be zero")
	}
	if hdr.Dim[2] == 0 {
		return nil, errors.New("y dimension must not be zero")
	}
	bDataLength = int64(hdr.Dim[1] * hdr.Dim[2])

	if hdr.Dim[3] > 0 {
		bDataLength = bDataLength * int64(hdr.Dim[3])
	}
	if hdr.Dim[4] > 0 {
		bDataLength = bDataLength * int64(hdr.Dim[4])
	}
	if hdr.Dim[5] > 0 {
		bDataLength = bDataLength * int64(hdr.Dim[5])
	}
	if hdr.Dim[6] > 0 {
		bDataLength = bDataLength * int64(hdr.Dim[6])
	}
	if hdr.Dim[7] > 0 {
		bDataLength = bDataLength * int64(hdr.Dim[7])
	}

	nByper, _ := assignDatatypeSize(int32(hdr.Datatype))
	bDataLength = bDataLength * int64(nByper)

	// Init a slice of bytes with capacity of bDataLength and initial value of 0
	bData := make([]byte, bDataLength, bDataLength)

	return bData, nil
}

func uint64ToFloat64(v uint64, datatype int32) float64 {
	var value float64

	switch datatype {
	case constant.DT_FLOAT64:
		value = float64(v)
	case constant.DT_INT64:
		value = float64(int64(v))
	case constant.DT_UINT64:
		value = float64(v)
	case constant.DT_COMPLEX64:
		value = math.Float64frombits(v)
	}
	return value
}

func uint32ToFloat64(v uint32, datatype int32) float64 {
	var value float64

	switch datatype {
	case constant.DT_INT32:
		value = float64(int32(v))
	case constant.DT_UINT32:
		value = float64(v)
	case constant.DT_FLOAT32:
		value = float64(float32(v))
	case constant.DT_RGBA32:
		value = float64(math.Float32frombits(v))
	}
	return value
}
