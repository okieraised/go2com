package dataset

import "github.com/okieraised/go2com/pkg/dicom/element"

type Dataset struct {
	Elements []*element.Element `json:"elements"`
}
