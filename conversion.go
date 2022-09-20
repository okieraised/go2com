package go2com

import (
	"fmt"
	"github.com/okieraised/go2com/pkg/dicom/element"
	"reflect"
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
		if vrStr == "OB" || vrStr == "OW" || strings.ToLower(vrStr) == "ox" {
			continue
		}
		res.mapElement(elem)
	}
	return res
}

func (m mappedTag) mapElement(elem *element.Element) {
	tagStr := fmt.Sprintf("%04X%04X", elem.Tag.Group, elem.Tag.Element)
	vrStr := elem.ValueRepresentationStr
	var vl interface{}

	// If VR is SQ then we do type assertion to []*element.Element. If the length of sequence is 0, then do nothing.
	//
	if vrStr == "SQ" {
		subVL := make([]interface{}, 0)
		vlArr, ok := elem.Value.([]*element.Element)
		if ok {
			if len(vlArr) == 0 {
				return
			}
			groupTag := fmt.Sprintf("%04X%04X", vlArr[0].Tag.Group, vlArr[0].Tag.Element)
			subElemGrp := make(mappedTag)
			for index, subVl := range vlArr {
				subTag := fmt.Sprintf("%04X%04X", subVl.Tag.Group, subVl.Tag.Element)
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
		vl = checkType(elem.Value)
	}
	m[tagStr] = TagBrowser{
		VR:    vrStr,
		Value: vl,
	}
}

func checkType(vl interface{}) []interface{} {
	res := make([]interface{}, 0)
	switch reflect.TypeOf(vl).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(vl)
		for i := 0; i < s.Len(); i++ {
			res = append(res, s.Index(i).Interface())
		}
	default:
		res = append(res, vl)
	}
	return res
}
