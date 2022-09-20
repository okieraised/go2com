package element

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/okieraised/go2com/internal/constants"
	"io"
	"strings"

	"github.com/okieraised/go2com/internal/system"
	"github.com/okieraised/go2com/pkg/dicom/reader"
	"github.com/okieraised/go2com/pkg/dicom/tag"
	"github.com/okieraised/go2com/pkg/dicom/vr"
)

const (
	GroupSeqItem      uint16 = 0xFFFE
	VLUndefinedLength uint32 = 0xFFFFFFFF
)

type Element struct {
	Tag                    tag.DicomTag
	TagName                string
	ValueRepresentation    vr.VRKind
	ValueRepresentationStr string
	ValueLength            uint32
	Value                  Value
}

func ReadElement(r reader.DcmReader, isImplicit bool, binOrder binary.ByteOrder) (*Element, error) {
	dcmTagInfo, err := readTag(r)
	if err != nil {
		return nil, err
	}

	if dcmTagInfo.Tag == tag.ItemDelimitationItem || dcmTagInfo.Tag == tag.Item {
		_ = r.Skip(4)
		return nil, nil
	}

	if dcmTagInfo.Tag == tag.PixelData && r.SkipPixelData() {
		_, err = r.Discard(int(r.GetFileSize()))
		if err != nil {
			return nil, err
		}
		return nil, nil
	}

	dmcTagName := dcmTagInfo.Name
	dcmVR, err := readVR(r, isImplicit, dcmTagInfo.Tag)
	if err != nil {
		return nil, err
	}

	dcmVL, err := readVL(r, isImplicit, dcmTagInfo.Tag, dcmVR)
	if err != nil {
		return nil, err
	}

	value, err := readValue(r, dcmTagInfo.Tag, dcmVR, dcmVL)
	if err != nil {
		return nil, err
	}

	elem := Element{
		Tag:                    dcmTagInfo.Tag,
		TagName:                dmcTagName,
		ValueRepresentationStr: dcmVR,
		ValueLength:            dcmVL,
		Value:                  value,
	}

	return &elem, nil
}

// readTag returns the tag information
func readTag(r reader.DcmReader) (*tag.TagInfo, error) {
	group, err := r.ReadUInt16()
	if err != nil {
		return nil, err
	}
	element, err := r.ReadUInt16()
	if err != nil {
		return nil, err
	}

	t := tag.DicomTag{
		Group:   group,
		Element: element,
	}

	// Check if tag is private. If yes, just return here
	// Otherwise, find info about the public tag
	if int(group)%2 != 0 {
		tagInfo := tag.TagInfo{
			Tag:    t,
			VR:     "",
			Name:   constants.PrivateTag,
			VM:     "",
			Status: "",
		}
		return &tagInfo, nil
	}

	tagInfo, err := tag.Find(t)
	if err != nil {
		return nil, err
	}
	return &tagInfo, nil
}

// readVR
func readVR(r reader.DcmReader, isImplicit bool, t tag.DicomTag) (string, error) {
	if isImplicit {
		record, err := tag.Find(t)
		if err != nil {
			return vr.Unknown, nil
		}
		return record.VR, nil
	}
	return r.ReadString(2)
}

// readVL
func readVL(r reader.DcmReader, isImplicit bool, t tag.DicomTag, valueRepresentation string) (uint32, error) {
	if isImplicit {
		return r.ReadUInt32()
	}

	switch valueRepresentation {
	// if the VR is equal to ‘OB’,’OW’,’OF’,’SQ’,’UI’ or ’UN’,
	// the VR is having an extra 2 bytes trailing to it. These 2 bytes trailing to VR are empty and are not decoded.
	// When VR is having these 2 extra empty bytes the VL will occupy 4 bytes rather than 2 bytes
	case vr.OtherByte, vr.OtherWord, vr.OtherFloat, vr.SequenceOfItems, vr.Unknown:
		r.Skip(2)
		valueLength, err := r.ReadUInt32()
		if err != nil {
			return 0, err
		}
		if valueLength == VLUndefinedLength &&
			(valueRepresentation == vr.UnlimitedCharacters ||
				valueRepresentation == vr.UniversalResourceIdentifier ||
				valueRepresentation == vr.UnlimitedText) {
			return 0, errors.New("UC, UR and UT must have defined length")
		}
		return valueLength, nil
	default:
		valueLength, err := r.ReadUInt16()
		if err != nil {
			return 0, err
		}
		vl := uint32(valueLength)
		if vl == 0xffff {
			vl = VLUndefinedLength
		}
		return vl, nil
	}
}

// readValue
func readValue(r reader.DcmReader, t tag.DicomTag, valueRepresentation string, valueLength uint32) (Value, error) {
	vrKind := vr.GetVR(t, valueRepresentation)
	switch vrKind {
	case vr.VRString, vr.VRDate:
		return readStringType(r, t, valueRepresentation, valueLength)
	case vr.VRInt16, vr.VRInt32, vr.VRUInt16, vr.VRUInt32, vr.VRTagList:
		return readIntType(r, t, valueRepresentation, valueLength)
	case vr.VRFloat32, vr.VRFloat64:
		return readFloatType(r, t, valueRepresentation, valueLength)
	case vr.VRBytes:
		return readByteType(r, t, valueRepresentation, valueLength)
	case vr.VRPixelData:
		return readByteType(r, t, valueRepresentation, valueLength)
	case vr.VRSequence:
		return readSequence(r, t, valueRepresentation, valueLength)
	default:
		return readStringType(r, t, valueRepresentation, valueLength)
	}
}

// readStringType
func readStringType(r reader.DcmReader, t tag.DicomTag, valueRepresentation string, valueLength uint32) (Value, error) {
	sep := "\\"
	str, err := r.ReadString(valueLength)
	if err != nil {
		return str, err
	}
	str = strings.Trim(str, " \000") // There is a space " \000", not "\000"
	if strings.Contains(str, sep) {
		res := strings.Split(str, sep)
		return res, nil
	}
	return str, nil
}

// readByteType reads the value as byte array or word array
func readByteType(r reader.DcmReader, t tag.DicomTag, valueRepresentation string, valueLength uint32) (Value, error) {
	switch valueRepresentation {
	case vr.OtherByte, vr.Unknown:
		byteStr := make([]byte, valueLength)
		_, err := io.ReadFull(r, byteStr)
		if err != nil {
			return nil, err
		}
		return len(byteStr), nil
	case vr.OtherWord:
		if valueLength%2 != 0 {
			return nil, fmt.Errorf("odd value encountered")
		}
		buf := bytes.NewBuffer(make([]byte, 0, valueLength))
		numWords := int(valueLength / 2)
		for i := 0; i < numWords; i++ {
			word, err := r.ReadUInt16()
			if err != nil {
				return nil, err
			}
			err = binary.Write(buf, system.NativeEndian, word)
			if err != nil {
				return nil, err
			}

		}
		return len(buf.Bytes()), nil
	default:
	}
	return nil, nil
}

// readIntType reads the value as integer
func readIntType(r reader.DcmReader, t tag.DicomTag, valueRepresentation string, valueLength uint32) (Value, error) {
	var subVal int
	retVal := make([]int, 0, valueLength/2)
	n, err := r.Peek(int(valueLength))
	if err != nil {
		return nil, err
	}
	subReader := bytes.NewReader(n)
	subRd := reader.NewDcmReader(bufio.NewReader(subReader), false)
	byteRead := 0
	for {
		if byteRead >= int(valueLength) {
			break
		}
		switch valueRepresentation {
		case vr.UnsignedShort, vr.SignedShortOrUnsignedShort, strings.ToLower(vr.SignedShortOrUnsignedShort):
			val, err := subRd.ReadUInt16()
			if err != nil {
				return nil, err
			}
			subVal = int(val)
			byteRead += 2
		case vr.AttributeTag:
			val, err := subRd.ReadUInt16()
			if err != nil {
				return nil, err
			}
			subVal = int(val)
			byteRead += 2
		case vr.UnsignedLong:
			val, err := subRd.ReadUInt32()
			if err != nil {
				return nil, err
			}
			subVal = int(val)
			byteRead += 4
		case vr.SignedLong:
			val, err := subRd.ReadInt32()
			if err != nil {
				return nil, err
			}
			subVal = int(val)
			byteRead += 4
		case vr.SignedShort:
			val, err := subRd.ReadInt16()
			if err != nil {
				return nil, err
			}
			subVal = int(val)
			byteRead += 2
			//default:
			//	return nil, fmt.Errorf("cannot parse value as integer for tag %v", t)
		}
		retVal = append(retVal, subVal)
	}
	_, _ = r.Discard(int(valueLength))
	if len(retVal) == 1 {
		return retVal[0], nil
	}
	return retVal, nil
}

// readFloatType reads the value as float
func readFloatType(r reader.DcmReader, t tag.DicomTag, valueRepresentation string, valueLength uint32) (Value, error) {
	var subVal float64
	retVal := make([]float64, 0, valueLength/2)
	n, err := r.Peek(int(valueLength))
	if err != nil {
		return nil, err
	}
	subReader := bytes.NewReader(n)
	subRd := reader.NewDcmReader(bufio.NewReader(subReader), false)
	byteRead := 0
	for {
		if byteRead >= int(valueLength) {
			break
		}
		switch valueRepresentation {
		case vr.FloatingPointSingle, vr.OtherFloat:
			val, err := subRd.ReadFloat32()
			if err != nil {
				return nil, err
			}
			subVal = float64(val)
			byteRead += 4
		case vr.FloatingPointDouble:
			val, err := subRd.ReadFloat64()
			if err != nil {
				return nil, err
			}
			subVal = val
			byteRead += 8
		}
		retVal = append(retVal, subVal)
	}
	_, _ = r.Discard(int(valueLength))
	if len(retVal) == 1 {
		return retVal[0], nil
	}

	return retVal, nil
}

// readSequence reads the value as sequence of items
func readSequence(r reader.DcmReader, t tag.DicomTag, valueRepresentation string, valueLength uint32) (Value, error) {
	var sequences []*Element
	// Reference: https://dicom.nema.org/dicom/2013/output/chtml/part05/sect_7.5.html
	if valueLength == VLUndefinedLength {
		for {
			subElement, err := ReadElement(r, r.IsImplicit(), r.ByteOrder())
			if err != nil {
				return nil, err
			}

			if subElement == nil {
				continue
			}

			if subElement.Tag == tag.SequenceDelimitationItem {
				break
			}
			sequences = append(sequences, subElement)
		}

	} else {
		n, err := r.Peek(int(valueLength))
		if err != nil {
			return nil, err
		}
		br := bytes.NewReader(n)
		subRd := reader.NewDcmReader(bufio.NewReader(br), false)
		_ = subRd.Skip(8)
		subRd.SetTransferSyntax(r.ByteOrder(), r.IsImplicit())
		for {
			subElement, err := ReadElement(subRd, subRd.IsImplicit(), subRd.ByteOrder())
			if err != nil {
				if err == io.EOF {
					break
				} else {
					return nil, err
				}
			}
			if subElement == nil {
				continue
			}
			sequences = append(sequences, subElement)
		}
		_, _ = r.Discard(int(valueLength))
	}

	return sequences, nil

}
