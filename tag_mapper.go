package go2com

import (
	"fmt"
	"github.com/okieraised/go2com/internal/utils"
	"github.com/okieraised/go2com/pkg/dicom/tag"
	"github.com/okieraised/go2com/pkg/dicom/vr"
	"strings"
)

type MappedTag map[string]tag.TagBrowser

// GetElementByTagString returns the element value of the input tag
// Tag should be in (gggg,eeee) or ggggeeee format
func (m MappedTag) GetElementByTagString(tagStr string) (interface{}, error) {
	tagStr = utils.FormatTag(tagStr)

	result, ok := m[tagStr]
	if !ok {
		return nil, fmt.Errorf("tag not found: %s", tagStr)
	}
	return result.Value, nil
}

// mapElement returns a map[string]interface{} with key as tag and value as the tag values
func (m MappedTag) mapElement(elem *Element) {
	tagStr := elem.Tag.StringWithoutParentheses()
	vrStr := elem.ValueRepresentationStr
	var value interface{}

	// If VR is SQ then we do type assertion to []*element.Element. If the length of sequence is 0, then do nothing.
	// Else, loop through each element in the sequence and extract the info
	if vrStr == "SQ" {
		subVL := make([]interface{}, 0)
		vlArr, ok := (elem.Value.RawValue).([]*Element)
		if ok {
			if len(vlArr) == 0 {
				return
			}
			groupTag := vlArr[0].Tag.StringWithoutParentheses()
			subElemGrp := make(MappedTag)
			for index, subVl := range vlArr {
				subVRStr := subVl.ValueRepresentationStr
				if subVRStr == "OB" || subVRStr == "OW" || subVRStr == "UN" || strings.ToLower(subVRStr) == "ox" {
					continue
				}
				subTag := subVl.Tag.StringWithoutParentheses()
				if subTag == groupTag && index > 0 {
					subVL = append(subVL, subElemGrp)
					subElemGrp = MappedTag{}
				}
				subElemGrp.mapElement(subVl)
				if index == len(vlArr)-1 {
					subVL = append(subVL, subElemGrp)
				}
			}
		}
		value = subVL
	} else {
		if elem.ValueRepresentationStr == vr.PersonName {
			value = utils.AppendToSlice(map[string]interface{}{
				"Alphabetic": elem.Value.RawValue,
			})
		} else {
			value = utils.AppendToSlice(elem.Value.RawValue)
		}

	}
	m[tagStr] = tag.TagBrowser{
		VR:    vrStr,
		Value: value,
	}
}

func createOrthancURI(ds Dataset) MappedTag {
	res := make(MappedTag, 3)
	prefix := "http://127.0.0.1:8042"
	uids, err := ds.RetrieveFileUID()
	if err != nil {
		return res
	}
	for _, elem := range ds.Elements {
		var bulkURI string
		tagStr := elem.Tag.StringWithoutParentheses()
		if elem.Tag == tag.RedPaletteColorLookupTableData || elem.Tag == tag.BluePaletteColorLookupTableData || elem.Tag == tag.GreenPaletteColorLookupTableData {
			bulkURI = fmt.Sprintf("%s/dicom-web/studies/%s/series/%s/instances/%s/bulk/%s", prefix, uids.StudyInstanceUID, uids.SeriesInstanceUID, uids.SOPInstanceUID, tagStr)
			res[tagStr] = tag.TagBrowser{
				VR:          "US",
				BulkDataURI: bulkURI,
			}
		}
	}
	return res
}
