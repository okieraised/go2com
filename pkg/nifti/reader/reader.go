package reader

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/okieraised/go2com/internal/constants"
	"github.com/okieraised/go2com/internal/utils"
	"net/http"
	"os"
)

type NiiReader interface{}

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
		r.version = 1

	case constants.NII2HeaderSize:
		r.version = 2
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
			r.version = 1
		case constants.NII2HeaderSize:
			r.version = 2
		default:
			return errors.New("invalid NIFTI file format")
		}
	}
	return errors.New("cannot determine NIFTI file version")
}
