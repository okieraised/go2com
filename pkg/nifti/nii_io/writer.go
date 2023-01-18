package nii_io

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/okieraised/go2com/internal/system"
	"github.com/okieraised/go2com/pkg/nifti/constant"
	"math"
	"os"
)

// niiWriter define the NIfTI writer structure.
//
// Parameters:
//     - `filePath`         : Export file path to write NIfTI image
//     - `writeHeaderFile`  : Whether to write NIfTI file pair (hdr + img file)
//     - `compression`      : Whether the NIfTI volume will be compressed
//     - `headerCompression`: Whether the NIfTI header will be compressed. This only works when the writeHeaderFile option is set to True
//     - `niiData`          : Input NIfTI data to write to file
//     - `header`           : Input NIfTI header to write to file. If nil, the default header will be constructed
type niiWriter struct {
	filePath          string      // Export file path to write NIfTI image
	writeHeaderFile   bool        // Whether to write NIfTI file pair (hdr + img file)
	compression       bool        // Whether the NIfTI volume will be compressed
	headerCompression bool        // Whether the NIfTI header will be compressed. This only works when the writeHeaderFile option is set to True
	niiData           *Nii        // Input NIfTI data to write to file
	header            *Nii1Header // Input NIfTI header to write to file. If nil, the default header will be constructed
}

// NewNiiWriter returns a new write for export
func NewNiiWriter(filePath string, options ...func(*niiWriter)) (*niiWriter, error) {
	writer := new(niiWriter)

	writer.filePath = filePath
	writer.writeHeaderFile = false   // Default to false. Write to a single file only
	writer.compression = false       // Default to false. No compression
	writer.headerCompression = false // Default to false. No header compression

	// Other options
	for _, opt := range options {
		opt(writer)
	}
	return writer, nil
}

// WithWriteHeaderFile sets the option to write NIfTI image to a header/image (.hdr/.img) file pair
//
// If true, output will be two files for the header and the image. Default is false.
func WithWriteHeaderFile(writeHeaderFile bool) func(*niiWriter) {
	return func(w *niiWriter) {
		w.writeHeaderFile = writeHeaderFile
	}
}

// WithHeaderCompression sets the option to write compressed NIfTI header structure to file (.hdr.gz)
//
// If true, the header will be compressed. It only works when WithWriteHeaderFile is provided with option `true`. Default is false.
func WithHeaderCompression(headerCompression bool) func(*niiWriter) {
	return func(w *niiWriter) {
		w.headerCompression = headerCompression
	}
}

// WithCompression sets the option to write compressed NIfTI image to a single file (.nii.gz)
//
// If true, the whole file will be compressed. Default is false.
func WithCompression(withCompression bool) func(writer *niiWriter) {
	return func(w *niiWriter) {
		w.compression = withCompression
	}
}

// WithHeader sets the option to allow user to provide predefined NIfTI-1 header structure.
//
// If no header provided, the header will be converted from the NIfTI image structure
func WithHeader(header *Nii1Header) func(*niiWriter) {
	return func(w *niiWriter) {
		w.header = header
	}
}

// WithNIfTIData sets the option to allow user to provide predefined NIfTI-1 data structure.
func WithNIfTIData(data *Nii) func(writer *niiWriter) {
	return func(w *niiWriter) {
		w.niiData = data
	}
}

// WriteToFile write the header and image to a NIfTI file
func (w *niiWriter) WriteToFile() error {
	// convert image structure to file
	err := w.convertImageToHeader()
	if err != nil {
		return err
	}

	if w.writeHeaderFile { // If user decides to write to a separate hdr/img file pair
		return nil
	} else { // Just one file for both header and the image data

		// We need to check the offset from the end of header file to the start of the voxel dataset
		offsetFromHeaderToVoxel := int(w.header.VoxOffset) - int(w.header.SizeofHdr)
		var offset []byte

		// If the header size is equals to 348 (default), then we need to add 4 bytes or the offset we calculated to make it divisible by 6
		if int(w.header.SizeofHdr) == constant.NII1HeaderSize {
			defaultPadding := 4
			if offsetFromHeaderToVoxel > 0 {
				offset = make([]byte, offsetFromHeaderToVoxel)
			} else {
				// This is for a case where we read the image as .hdr/.img pair but want to write to a single file. We
				// have to update the VoxOffset value
				w.header.VoxOffset = float32(int(w.header.SizeofHdr) + defaultPadding)
				offset = make([]byte, defaultPadding)
			}
		}

		hdrBuf := &bytes.Buffer{}
		err := binary.Write(hdrBuf, system.NativeEndian, w.header)
		if err != nil {
			return err
		}

		bHeader := hdrBuf.Bytes()
		bData := w.niiData.Volume

		if offsetFromHeaderToVoxel > 0 {
			offset = make([]byte, offsetFromHeaderToVoxel)
		}

		dataset := []byte{}
		dataset = append(dataset, bHeader...)
		dataset = append(dataset, offset...)
		dataset = append(dataset, bData...)

		// Create a file
		file, err := os.Create(w.filePath)
		if err != nil {
			return err
		}
		defer file.Close()

		if w.compression { // If the compression is set to true, then write a compressed file
			gzipWriter := gzip.NewWriter(file)
			_, err = gzipWriter.Write(dataset)
			if err != nil {
				return err
			}
			err = gzipWriter.Close()
			if err != nil {
				return err
			}
		} else { // Otherwise, just write normal file
			_, err = file.Write(dataset)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// convertImageToHeader returns the header from a NIfTI image structure
func (w *niiWriter) convertImageToHeader() error {
	if w.niiData == nil {
		return errors.New("image data is nil")
	}

	header := new(Nii1Header)
	header.SizeofHdr = constant.NII1HeaderSize
	header.Regular = 'r'

	header.Dim[0] = int16(w.niiData.NDim)
	header.Dim[1], header.Dim[2], header.Dim[3] = int16(w.niiData.Nx), int16(w.niiData.Ny), int16(w.niiData.Nz)
	header.Dim[4], header.Dim[5], header.Dim[6] = int16(w.niiData.Nt), int16(w.niiData.Nu), int16(w.niiData.Nv)
	header.Dim[7] = int16(w.niiData.Nw)

	header.Pixdim[0] = 0.0
	header.Pixdim[1], header.Pixdim[2], header.Pixdim[3] = float32(math.Abs(w.niiData.Dx)), float32(math.Abs(w.niiData.Dy)), float32(math.Abs(w.niiData.Dz))
	header.Pixdim[4], header.Pixdim[5], header.Pixdim[6] = float32(math.Abs(w.niiData.Dt)), float32(math.Abs(w.niiData.Du)), float32(math.Abs(w.niiData.Dv))
	header.Pixdim[7] = float32(w.niiData.Dw)

	header.Datatype = int16(w.niiData.Datatype)
	header.Bitpix = int16(8 * w.niiData.NByPer)

	if w.niiData.CalMax > w.niiData.CalMin {
		header.CalMin = float32(w.niiData.CalMin)
		header.CalMax = float32(w.niiData.CalMax)
	}

	if w.niiData.SclSlope != 0.0 {
		header.SclSlope = float32(w.niiData.SclSlope)
		header.SclInter = float32(w.niiData.SclInter)
	}

	if w.niiData.Descrip[0] != 0x0 {
		for i := 0; i < 79; i++ {
			header.Descrip[i] = w.niiData.Descrip[i]
		}
		header.Descrip[79] = 0x0
	}

	if w.niiData.AuxFile[0] != 0x0 {
		for i := 0; i < 23; i++ {
			header.AuxFile[i] = w.niiData.AuxFile[i]
		}
		header.AuxFile[23] = 0x0
	}

	// Load NIFTI specific stuff into the header
	if w.writeHeaderFile {
		header.Magic = [4]byte{110, 43, 49, 0} // n+1
	} else {
		header.Magic = [4]byte{110, 105, 49, 0} // ni1
	}

	header.IntentCode = int16(w.niiData.IntentCode)
	header.IntentP1 = float32(w.niiData.IntentP1)
	header.IntentP2 = float32(w.niiData.IntentP2)
	header.IntentP3 = float32(w.niiData.IntentP3)
	if w.niiData.IntentName[0] != 0x0 {
		for i := 0; i < 15; i++ {
			header.IntentName[i] = w.niiData.IntentName[i]
		}
		header.IntentName[15] = 0x0
	}

	header.VoxOffset = float32(w.niiData.VoxOffset)
	header.XyztUnits = convertSpaceTimeToXYZT(w.niiData.XYZUnits, w.niiData.TimeUnits)
	header.Toffset = float32(w.niiData.TOffset)

	if w.niiData.QformCode > 0 {
		header.QformCode = int16(w.niiData.QformCode)
		header.QuaternB = float32(w.niiData.QuaternB)
		header.QuaternC = float32(w.niiData.QuaternC)
		header.QuaternD = float32(w.niiData.QuaternD)

		header.QoffsetX = float32(w.niiData.QoffsetX)
		header.QoffsetY = float32(w.niiData.QoffsetY)
		header.QoffsetZ = float32(w.niiData.QoffsetZ)

		if w.niiData.QFac >= 0 {
			header.Pixdim[0] = 1.0
		} else {
			header.Pixdim[0] = -1.0
		}
	}

	if w.niiData.SformCode > 0 {
		header.SformCode = int16(w.niiData.SformCode)
		header.SrowX[0] = float32(w.niiData.StoXYZ.M[0][0])
		header.SrowX[1] = float32(w.niiData.StoXYZ.M[0][1])
		header.SrowX[2] = float32(w.niiData.StoXYZ.M[0][2])
		header.SrowX[3] = float32(w.niiData.StoXYZ.M[0][3])

		header.SrowY[0] = float32(w.niiData.StoXYZ.M[1][0])
		header.SrowY[1] = float32(w.niiData.StoXYZ.M[1][1])
		header.SrowY[2] = float32(w.niiData.StoXYZ.M[1][2])
		header.SrowY[3] = float32(w.niiData.StoXYZ.M[1][3])

		header.SrowZ[0] = float32(w.niiData.StoXYZ.M[2][0])
		header.SrowZ[1] = float32(w.niiData.StoXYZ.M[2][1])
		header.SrowZ[2] = float32(w.niiData.StoXYZ.M[2][2])
		header.SrowZ[3] = float32(w.niiData.StoXYZ.M[2][3])
	}

	header.DimInfo = convertFPSIntoDimInfo(w.niiData.FreqDim, w.niiData.PhaseDim, w.niiData.SliceDim)

	header.SliceCode = uint8(w.niiData.SliceCode)
	header.SliceStart = int16(w.niiData.SliceStart)
	header.SliceEnd = int16(w.niiData.SliceEnd)
	header.SliceDuration = float32(w.niiData.SliceDuration)

	w.header = header

	return nil
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
