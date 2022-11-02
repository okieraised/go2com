package nii_reader

// #include "../nifti1/nifti1.h"
import "C"

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/okieraised/go2com/internal/constants"
	"github.com/okieraised/go2com/internal/utils"
	"github.com/okieraised/go2com/pkg/nifti/nifti1"
	"net/http"
	"os"
)

type NiiReader interface {
	Parse() error
	GetBinaryOrder() binary.ByteOrder
	GetNiiVersion() int
	GetUnitsOfMeasurements() ([2]string, error)
}

type NiiData interface {
	GetUnitsOfMeasurements() ([2]string, error)
	GetHeader() interface{}
	GetImg() interface{}
	GetAffine() interface{}
	GetImgShape() interface{}
	GetVoxelSize() interface{}
	GetTimeSeries(x, y, z int) interface{}
	GetSlice(z, t int) interface{}
}

type niiReader struct {
	reader      *bytes.Reader
	binaryOrder binary.ByteOrder
	niiVersion  int
	niiData     NiiData
}

// NewNiiReader creates a new NIFTI reader
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
	}, nil
}

func (r *niiReader) GetUnitsOfMeasurements() ([2]string, error) {
	return r.niiData.GetUnitsOfMeasurements()
}

func (r *niiReader) Parse() error {
	// The field short dim[8] contains the size of the image array. The first element (dim[0]) contains the number of
	// dimensions (1-7). If dim[0] is not in this interval, the data is assumed to have opposite endianness and so,
	// should be byte-swapped (the nifti standard does not specify a specific field for endianness,
	// but encourages the use of dim[0] for this purpose)
	err := r.checkNiiVersion()
	if err != nil {
		return err
	}
	// Reset the reader back to the original position
	_, err = r.reader.Seek(0, 0)
	if err != nil {
		return err
	}

	switch r.niiVersion {
	case constants.NIIVersion1:
		niiData := new(nifti1.Nii1)
		err := niiData.Parse(r.reader, r.binaryOrder)
		if err != nil {
			return err
		}
		r.niiData = niiData

	case constants.NIIVersion2:
	default:
		return errors.New("invalid/unsupported NIFTI format")
	}

	fmt.Println(r.niiData.GetUnitsOfMeasurements())
	fmt.Println(r.niiData.GetHeader())
	fmt.Println(r.niiData.GetAffine())
	fmt.Println(r.niiData.GetImgShape())
	fmt.Println(r.niiData.GetVoxelSize())
	//fmt.Println(r.niiData.GetSlice(12, 0))
	fmt.Println(r.niiData.GetTimeSeries(1, 1, 1))

	return nil
}

func (r *niiReader) GetBinaryOrder() binary.ByteOrder {
	return r.binaryOrder
}

func (r *niiReader) GetNiiVersion() int {
	return r.niiVersion
}

func (r *niiReader) checkNiiVersion() error {
	var hSize int32

	err := binary.Read(r.reader, r.binaryOrder, &hSize)
	if err != nil {
		return err
	}

	switch hSize {
	case constants.NII1HeaderSize:
		r.setNiiVersion(constants.NIIVersion1)
	case constants.NII2HeaderSize:
		r.setNiiVersion(constants.NIIVersion2)
	default:
		r.binaryOrder = binary.BigEndian
		var hSize int32

		err = binary.Read(r.reader, r.binaryOrder, &hSize)
		if err != nil {
			return err
		}
		switch hSize {
		case constants.NII1HeaderSize:
			r.setNiiVersion(constants.NIIVersion1)
		case constants.NII2HeaderSize:
			r.setNiiVersion(constants.NIIVersion2)
		default:
			return errors.New("invalid NIFTI file format")
		}
	}
	return nil
}

func (r *niiReader) setNiiVersion(version int) {
	r.niiVersion = version
}
