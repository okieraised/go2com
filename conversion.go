package go2com

import (
	"fmt"
	"github.com/okieraised/go2com/internal/utils"
	"github.com/okieraised/go2com/pkg/dicom/element"
	"github.com/okieraised/go2com/pkg/dicom/tag"
	"github.com/okieraised/go2com/pkg/dicom/vr"
	"reflect"
	"strconv"
	"strings"
)

type MappedTag map[string]tag.TagBrowser

// Export returns the mapped tag/(vr,value) dictionary
func (p *Parser) Export(exportMeta bool) MappedTag {
	res := make(MappedTag, len(p.metadata.Elements)+len(p.dataset.Elements))
	if exportMeta {
		mt := p.metadata
		for _, elem := range mt.Elements {
			res.mapElement(elem)
		}
	}

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
	var value interface{}

	// If VR is SQ then we do type assertion to []*element.Element. If the length of sequence is 0, then do nothing.
	// Else, loop through each element in the sequence and extract the info
	if vrStr == "SQ" {
		subVL := make([]interface{}, 0)
		vlArr, ok := (elem.Value.RawValue).([]*element.Element)
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
		value = switchStringToNumeric(elem)
		value = utils.AppendToSlice(value)
	}
	m[tagStr] = tag.TagBrowser{
		VR:    vrStr,
		Value: value,
	}
}

func switchStringToNumeric(elem *element.Element) interface{} {
	switch elem.ValueRepresentationStr {
	case vr.IntegerString:
		switch reflect.ValueOf(elem.Value.RawValue).Kind() {
		case reflect.Slice:
			ValStrArr, ok := (elem.Value.RawValue).([]string)
			if !ok {
				return elem.Value.RawValue
			}
			res := make([]int, 0)
			for _, sub := range ValStrArr {
				intVar, err := strconv.Atoi(sub)
				if err != nil {
					return elem.Value.RawValue
				}
				res = append(res, intVar)
			}
			return res
		case reflect.String:
			valStr, ok := (elem.Value.RawValue).(string)
			if !ok {
				return elem.Value.RawValue
			}
			intVal, err := strconv.Atoi(valStr)
			if err != nil {
				return elem.Value.RawValue
			}
			return intVal
		}
	case vr.DecimalString, vr.OtherFloat, vr.OtherDouble:
		switch reflect.ValueOf(elem.Value.RawValue).Kind() {
		case reflect.Slice:
			ValStrArr, ok := (elem.Value.RawValue).([]string)
			if !ok {
				return elem.Value.RawValue
			}
			res := make([]float64, 0)
			for _, sub := range ValStrArr {
				flVar, err := strconv.ParseFloat(sub, 64)
				if err != nil {
					return elem.Value.RawValue
				}
				res = append(res, flVar)
			}
			return res
		case reflect.String:
			valStr, ok := (elem.Value.RawValue).(string)
			if !ok {
				return elem.Value.RawValue
			}
			flVar, err := strconv.ParseFloat(valStr, 64)
			if err != nil {
				return elem.Value.RawValue
			}
			return flVar
		}
	default:
	}
	return elem.Value.RawValue
}
