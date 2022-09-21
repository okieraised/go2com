package go2com

import (
	"fmt"
	"github.com/okieraised/go2com/internal/utils"
	"github.com/okieraised/go2com/pkg/dicom/element"
	"github.com/okieraised/go2com/pkg/dicom/tag"
	"strings"
)

type MappedTag map[string]tag.TagBrowser

// Export returns the mapped tag/(vr,value) dictionary
func (p *Parser) Export() MappedTag {
	res := make(MappedTag)
	ds := p.dataset
	for _, elem := range ds.Elements {
		vrStr := elem.ValueRepresentationStr
		if vrStr == "OB" || vrStr == "OW" || vrStr == "UN" || strings.ToLower(vrStr) == "ox" {
			continue
		}
		res.mapElement(elem)
	}
	return res
}

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
func (m MappedTag) mapElement(elem *element.Element) {
	tagStr := elem.Tag.StringWithoutParentheses()
	vrStr := elem.ValueRepresentationStr
	var vl interface{}

	// If VR is SQ then we do type assertion to []*element.Element. If the length of sequence is 0, then do nothing.
	// Else, loop through each element in the sequence and extract the info
	if vrStr == "SQ" {
		subVL := make([]interface{}, 0)
		vlArr, ok := elem.Value.([]*element.Element)
		if ok {
			if len(vlArr) == 0 {
				return
			}
			groupTag := vlArr[0].Tag.StringWithoutParentheses()
			subElemGrp := make(MappedTag)
			for index, subVl := range vlArr {
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
		vl = subVL
	} else {
		vl = utils.AppendToSlice(elem.Value)
	}
	m[tagStr] = tag.TagBrowser{
		VR:    vrStr,
		Value: vl,
	}
}
