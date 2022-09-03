package element

import (
	"errors"

	"github.com/okieraised/go2com/pkg/dicom/reader"
	"github.com/okieraised/go2com/pkg/dicom/tag"
	"github.com/okieraised/go2com/pkg/dicom/vr"
)

var mapHandleVER map[string]func(string, string)

const (
	GroupSeqItem      uint16 = 0xFFFE
	VLUndefinedLength uint32 = 0xFFFFFFFF
)

type Element struct {
	Tag                    tag.DicomTag `json:"tag"`
	ValueRepresentation    vr.VRKind    `json:"vr"`
	ValueRepresentationStr string       `json:"vr_str"`
	ValueLength            uint32       `json:"valueLength"`
	Value                  string       `json:"value"`
}

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

	elem := Element{
		Tag:                    *dcmTag,
		ValueRepresentationStr: dcmVR,
		ValueLength:            dcmVL,
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
	case vr.OtherByte, vr.OtherWord, vr.OtherFloat, vr.SequenceOfItems, vr.UniqueIdentifier, vr.Unknown:
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
