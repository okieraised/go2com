package dataset

import (
	"fmt"
	"github.com/okieraised/go2com/internal/utils"
	"github.com/okieraised/go2com/pkg/dicom/element"
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
		if strings.ToLower(elem.TagName) == "sopinstanceuid" {
			res.SOPInstanceUID = (elem.Value.RawValue).(string)
		}
		if strings.ToLower(elem.TagName) == "seriesinstanceuid" {
			res.SeriesInstanceUID = (elem.Value.RawValue).(string)
		}
		if strings.ToLower(elem.TagName) == "studyinstanceuid" {
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
