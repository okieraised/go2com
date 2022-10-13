package dataset

import (
	"fmt"
	"github.com/okieraised/go2com/internal/utils"
	"github.com/okieraised/go2com/pkg/dicom/element"
	"github.com/okieraised/go2com/pkg/dicom/tag"
	"strings"
)

type Dataset struct {
	Elements []*element.Element `json:"elements"`
}

type DicomUID struct {
	StudyInstanceUID  string
	SeriesInstanceUID string
	SOPInstanceUID    string
}

func (ds *Dataset) RetrieveFileUID() (*DicomUID, error) {
	res := DicomUID{}
	for _, elem := range ds.Elements {
		if elem.Tag == tag.SOPInstanceUID {
			res.SOPInstanceUID = (elem.Value.RawValue).(string)
		}
		if elem.Tag == tag.SeriesInstanceUID {
			res.SeriesInstanceUID = (elem.Value.RawValue).(string)
		}
		if elem.Tag == tag.StudyInstanceUID {
			res.StudyInstanceUID = (elem.Value.RawValue).(string)
		}
	}
	return &res, nil
}

// FindElementByTagStr returns the corresponding element of the input tag.
// Tag must be in 'ggggeeee' or '(gggg,eeee)' format
func (ds *Dataset) FindElementByTagStr(tagStr string) (*element.Element, error) {
	tagStr = utils.FormatTag(tagStr)
	for _, elem := range ds.Elements {
		if elem.Tag.StringWithoutParentheses() == tagStr {
			return elem, nil
		}
	}
	return nil, fmt.Errorf("cannot find tag %s", tagStr)
}

// FindElementByTagName returns the corresponding element of the input tag name.
func (ds *Dataset) FindElementByTagName(tagName string) (*element.Element, error) {
	tagName = utils.FormatTagName(tagName)
	for _, elem := range ds.Elements {
		if strings.ToLower(elem.TagName) == tagName {
			return elem, nil
		}
	}
	return nil, fmt.Errorf("cannot find tag %s", tagName)
}

// FindElementByTag returns the corresponding element of the input tag name.
func (ds *Dataset) FindElementByTag(tagName tag.DicomTag) (*element.Element, error) {
	for _, elem := range ds.Elements {
		if tagName == elem.Tag {
			return elem, nil
		}
	}
	return nil, fmt.Errorf("cannot find tag %s", tagName)
}
