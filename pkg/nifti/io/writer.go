package io

import (
	"fmt"
	"github.com/okieraised/go2com/pkg/nifti/constant"
)

type NiiWriter interface {
}

type niiWriter struct {
	niiData *Nii
}

func (w *niiWriter) writeToFile(fileName string) {

}

func (w *niiWriter) makeNewHeader() {

}

func (w *niiWriter) makeNewNii1Header(inDim *[8]int16, inDatatype int32) {

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

	// validate dim: if there is any problem, apply default dim
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

	// Validate dtype
	datatype := inDatatype
	if !IsValidDatatype(datatype) {
		datatype = constant.DT_FLOAT32
	}

	// Populate the header struct
	header.SizeofHdr = NII1HeaderSize
	header.Regular = 'r'

	// Init dim and pixdim
	header.Dim[0] = dim[0]
	header.Pixdim[0] = 0.0
	for c := 1; c <= int(dim[0]); c++ {
		header.Dim[c] = dim[c]
		header.Pixdim[c] = 1.0
	}

	header.Datatype = int16(datatype)

	nByper, _ := w.assignDatatypeSize(datatype)
	header.Bitpix = 8 * nByper
	header.Magic = [4]byte{110, 43, 49, 0}

	w.niiData.n1Header = header
}

// assignDatatypeSize sets the number of bytes per voxel and the swapsize based on a datatype code
func (w *niiWriter) assignDatatypeSize(datatype int32) (int16, int16) {
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

func (w *niiWriter) makeNewImage() {

}
