package go2com

import (
	"github.com/okieraised/go2com/internal/utils"
	"github.com/okieraised/go2com/pkg/dicom/element"
	"strings"
)

type TagBrowser struct {
	VR    string      `json:"vr"`
	Value interface{} `json:"Value"`
}

type mappedTag map[string]TagBrowser

func (p *Parser) Export() map[string]TagBrowser {
	res := make(mappedTag)
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

// mapElement returns a map[string]interface{} with key as tag and value as the tag values
func (m mappedTag) mapElement(elem *element.Element) {
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
			subElemGrp := make(mappedTag)
			for index, subVl := range vlArr {
				subTag := subVl.Tag.StringWithoutParentheses()
				if subTag == groupTag && index > 0 {
					subVL = append(subVL, subElemGrp)
					subElemGrp = mappedTag{}
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
	m[tagStr] = TagBrowser{
		VR:    vrStr,
		Value: vl,
	}
}
