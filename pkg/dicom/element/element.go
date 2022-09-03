package element

import (
	"github.com/okieraised/go2com/pkg/dicom/reader"
	"github.com/okieraised/go2com/pkg/dicom/tag"
	"github.com/okieraised/go2com/pkg/dicom/vr"
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
	elem := Element{
		Tag:                    *dcmTag,
		ValueRepresentationStr: dcmVR,
	}

	return &elem, nil

}

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

func readVR(r reader.DcmReader, isImplicit bool, t tag.DicomTag) (string, error) {
	if isImplicit {
		if record, err := tag.Find(t); err == nil {
			return record.VR, nil
		}
		return vr.VR_UNKNOWN, nil
	}
	return r.ReadString(2)
}
