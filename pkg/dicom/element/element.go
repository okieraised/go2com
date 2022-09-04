package element

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/okieraised/go2com/internal/system"
	"github.com/okieraised/go2com/pkg/dicom/reader"
	"github.com/okieraised/go2com/pkg/dicom/tag"
	"github.com/okieraised/go2com/pkg/dicom/vr"
)

var mapHandleVR map[string]func(r reader.DcmReader, t tag.DicomTag, valueRepresentation string, valueLength uint32)

const (
	GroupSeqItem      uint16 = 0xFFFE
	VLUndefinedLength uint32 = 0xFFFFFFFF
)

type Element struct {
	Tag                    tag.DicomTag
	ValueRepresentation    vr.VRKind
	ValueRepresentationStr string
	ValueLength            uint32
	Value                  Value
}

type SequenceItem struct {
	elements []*Element
}

type SequenceItemSet struct {
	elements []*SequenceItem
}

type Value interface{}

func ReadElement(r reader.DcmReader, isImplicit bool) (*Element, error) {
	dcmTag, err := readTag(r, isImplicit)
	if err != nil {
		return nil, err
	}

	dcmVR, err := readVR(r, isImplicit, *dcmTag)
	if err != nil {
		return nil, err
	}

	dcmVL, err := readVL(r, isImplicit, *dcmTag, dcmVR)
	if err != nil {
		return nil, err
	}

	value, err := readValue(r, isImplicit, *dcmTag, dcmVR, dcmVL)
	if err != nil {
		return nil, err
	}

	elem := Element{
		Tag:                    *dcmTag,
		ValueRepresentationStr: dcmVR,
		ValueLength:            dcmVL,
		Value:                  value,
	}

	return &elem, nil

}

// readTag
func readTag(r reader.DcmReader, isImplicit bool) (*tag.DicomTag, error) {
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

	_, err = tag.Find(t)
	if err != nil {
		return nil, err
	}

	return &t, nil

}

// readVR
func readVR(r reader.DcmReader, isImplicit bool, t tag.DicomTag) (string, error) {
	if isImplicit {
		if record, err := tag.Find(t); err == nil {
			return record.VR, nil
		}
		return vr.Unknown, nil
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
func readValue(r reader.DcmReader, isImplicit bool, t tag.DicomTag, valueRepresentation string, valueLength uint32) (Value, error) {

	vrKind := vr.GetVR(t, valueRepresentation)

	switch vrKind {
	case vr.VRString, vr.VRDate:
		return readStringType(r, t, valueRepresentation, valueLength)
	case vr.VRInt16List, vr.VRInt32List, vr.VRUInt16List, vr.VRUInt32List, vr.VRTagList:
		return readIntType(r, t, valueRepresentation, valueLength)
	case vr.VRBytes:
		return readByteType(r, t, valueRepresentation, valueLength)
	case vr.VRPixelData:
	case vr.VRSequence:
		return readSequence(r, t, valueRepresentation, valueLength)
	default:
		return readStringType(r, t, valueRepresentation, valueLength)
	}
	return nil, nil
}

// readStringType
func readStringType(r reader.DcmReader, t tag.DicomTag, valueRepresentation string, valueLength uint32) (Value, error) {
	str, err := r.ReadString(valueLength)
	if err != nil {
		return str, err
	}

	str = strings.Trim(str, " \000")
	return str, nil
}

// readByteType
func readByteType(r reader.DcmReader, t tag.DicomTag, valueRepresentation string, valueLength uint32) (Value, error) {
	switch valueRepresentation {
	case vr.OtherByte, vr.Unknown:
		byteStr := make([]byte, valueLength)
		_, err := io.ReadFull(r, byteStr)
		if err != nil {
			return nil, err
		}
		return byteStr, nil
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

			if system.NativeEndian == binary.LittleEndian {
				err = binary.Write(buf, binary.LittleEndian, word)
				if err != nil {
					return nil, err
				}
			} else {
				err = binary.Write(buf, binary.BigEndian, word)
				if err != nil {
					return nil, err
				}
			}

		}
	default:

	}
	return nil, nil
}

func readIntType(r reader.DcmReader, t tag.DicomTag, valueRepresentation string, valueLength uint32) (Value, error) {
	switch valueRepresentation {
	case vr.UnsignedShort:
		val, err := r.ReadUInt16()
		if err != nil {
			return nil, err
		}
		return val, nil
	case vr.AttributeTag:
		group, err := r.ReadUInt16()
		if err != nil {
			return nil, err
		}
		elem, err := r.ReadUInt16()
		if err != nil {
			return nil, err
		}
		return tag.DicomTag{
			Group:   group,
			Element: elem,
		}, nil

	case vr.UnsignedLong:
		val, err := r.ReadUInt32()
		if err != nil {
			return nil, err
		}
		return val, nil

	case vr.SignedLong:
		val, err := r.ReadInt32()
		if err != nil {
			return nil, err
		}
		return val, nil

	case vr.SignedShort:
		val, err := r.ReadInt16()
		if err != nil {
			return nil, err
		}
		return val, nil

	default:
		return nil, fmt.Errorf("cannot parse value as integer for tag %v", t)

	}
}

func readSequence(r reader.DcmReader, t tag.DicomTag, valueRepresentation string, valueLength uint32) (Value, error) {
	// var sequences SequenceItemSet
	var sequences []Element

	// Reference: https://dicom.nema.org/dicom/2013/output/chtml/part05/sect_7.5.html

	if valueLength == VLUndefinedLength {
		for {
			subElement, err := ReadElement(r, false)
			if err != nil {
				return nil, err
			}
			fmt.Println("subElement 1", subElement)
			if subElement.Tag == tag.SequenceDelimitationItem {
				break
			}
			// sequences.elements = append(sequences.elements, subElement.Value.(*SequenceItem))
		}

	} else {
		n, err := r.Peek(int(valueLength))
		if err != nil {
			return nil, err
		}
		br := bytes.NewReader(n)
		subRd := reader.NewDcmReader(bufio.NewReader(br), false)
		subRd.Skip(8)
		for {
			subElement, err := ReadElement(subRd, false)
			if err != nil {
				if err == io.EOF {
					break
				} else {
					return nil, err
				}

			}
			sequences = append(sequences, *subElement)

		}
		// fmt.Println("sequences", sequences)
		r.Discard(int(valueLength))
	}

	return sequences, nil

}
