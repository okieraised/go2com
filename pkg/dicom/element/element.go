package element

import (
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
