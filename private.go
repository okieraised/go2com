package go2com

import (
	"bytes"
	"encoding/binary"
	"github.com/okieraised/go2com/pkg/dicom/tag"
	"github.com/okieraised/go2com/pkg/dicom/uid"
	"github.com/okieraised/go2com/pkg/dicom/vr"
	"io"
)

func (r *dcmReader) isTrackingImplicit() bool {
	return r.keepTrackImplicit
}

func (r *dcmReader) setOverallImplicit(isImplicit bool) {
	r.keepTrackImplicit = isImplicit
}

func (r *dcmReader) readUInt8() (uint8, error) {
	var res uint8

	err := binary.Read(r, r.binaryOrder, &res)

	return res, err
}

func (r *dcmReader) readUInt16() (uint16, error) {
	var res uint16
	err := binary.Read(r, r.binaryOrder, &res)
	return res, err
}

func (r *dcmReader) readUInt32() (uint32, error) {
	var res uint32
	err := binary.Read(r, r.binaryOrder, &res)
	return res, err
}

func (r *dcmReader) readUInt64() (uint64, error) {
	var res uint64
	err := binary.Read(r, r.binaryOrder, &res)
	return res, err
}

func (r *dcmReader) readInt8() (int8, error) {
	var res int8
	err := binary.Read(r, r.binaryOrder, &res)
	return res, err
}

func (r *dcmReader) readInt16() (int16, error) {
	var res int16
	err := binary.Read(r, r.binaryOrder, &res)
	return res, err
}

func (r *dcmReader) readInt32() (int32, error) {
	var res int32
	err := binary.Read(r, r.binaryOrder, &res)
	return res, err
}

func (r *dcmReader) readInt64() (int64, error) {
	var res int64
	err := binary.Read(r, r.binaryOrder, &res)
	return res, err
}

func (r *dcmReader) readFloat32() (float32, error) {
	var res float32
	err := binary.Read(r, r.binaryOrder, &res)
	return res, err
}

func (r *dcmReader) readFloat64() (float64, error) {
	var res float64
	err := binary.Read(r, r.binaryOrder, &res)
	return res, err
}

func (r *dcmReader) peek(n int) ([]byte, error) {
	return r.reader.Peek(n)
}

func (r *dcmReader) discard(n int) (int, error) {
	return r.reader.Discard(n)
}

func (r *dcmReader) skip(n int64) error {
	_, err := io.CopyN(io.Discard, r, n)
	return err
}

func (r *dcmReader) parse() error {
	_ = r.SetFileSize(r.fileSize)
	err := r.IsValidDICOM()
	if err != nil {
		return err
	}
	_ = r.skip(132)
	err = r.parseMetadata()
	if err != nil {
		return err
	}

	// IMPORTANT: Additional check is needed here since there are few instances where the DICOM
	// meta header is registered as Explicit Little-Endian, but Implicit Little-Endian is used in the body
	err = r.verifyImplicity()
	if err != nil {
		return nil
	}

	if r.skipDataset {
		return nil
	}
	err = r.parseDataset()
	if err != nil {
		return err
	}
	return nil
}

// parseMetadata parses the file meta information according to
// https://dicom.nema.org/dicom/2013/output/chtml/part10/chapter_7.html
// the File Meta Information shall be encoded using the Explicit VR Little Endian Transfer Syntax
// (UID=1.2.840.10008.1.2.1)
func (r *dcmReader) parseMetadata() error {
	var metadata []*Element
	var transferSyntaxUID string

	for {
		// No longer relied on the MetaInformationGroupLength tag to determine the length of the meta header.
		// We check if the group tag is 0x0002 before proceeding to read the element. If the group tag is not 0x0002,
		// then break the loop
		n, err := r.peek(2)
		if err != nil {
			return err
		}
		if bytes.Compare(n, []byte{0x2, 0x0}) != 0 {
			break
		}

		res, err := ReadElement(r, r.IsImplicit(), r.ByteOrder())
		if err != nil {
			return err
		}
		metadata = append(metadata, res)
		if res.Tag == tag.TransferSyntaxUID {
			transferSyntaxUID = (res.Value.RawValue).(string)
		}
	}
	r.metadata = Dataset{Elements: metadata}

	// Set transfer syntax here for the dataset parser
	binOrder, isImplicit, err := uid.ParseTransferSyntaxUID(transferSyntaxUID)
	if err != nil {
		return err
	}
	r.SetTransferSyntax(binOrder, isImplicit)
	r.setOverallImplicit(isImplicit)

	return nil
}

func (r *dcmReader) verifyImplicity() error {
	// Need to check if the implicit matches between header and body
	n, err := r.peek(6)
	if err != nil {
		return err
	}
	if !vr.VRMapper[string(n[4:6])] && !r.IsImplicit() {
		r.SetTransferSyntax(r.binaryOrder, true)
	}
	if vr.VRMapper[string(n[4:6])] && r.IsImplicit() {
		r.SetTransferSyntax(r.binaryOrder, false)
	}
	return nil
}

// parseDataset parses the file dataset after the file meta header
func (r *dcmReader) parseDataset() error {
	var data []*Element
	for {
		res, err := ReadElement(r, r.IsImplicit(), r.ByteOrder())
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		data = append(data, res)
		//fmt.Println(res)
	}
	dicomDataset := Dataset{Elements: data}
	//r.dataset.Elements = append(r.dataset.Elements, dicomDataset.Elements...)
	r.dataset = dicomDataset
	return nil
}
