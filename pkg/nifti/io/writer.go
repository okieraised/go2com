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

	nByper, _ := assignDatatypeSize(datatype)
	header.Bitpix = 8 * nByper
	header.Magic = [4]byte{110, 43, 49, 0}

	return header
}

func MakeNewImage(inDim *[8]int16, inDatatype int32, dataFill int) {
	//header := MakeNewNii1Header(inDim, inDatatype)

}

func NiftiHeaderToImage(header *Nii1Header) {
	//var ii, doSwap, iOff int
	//var isNIFTI, isOneFile int

}
