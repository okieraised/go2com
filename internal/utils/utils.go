package utils

import (
	"reflect"
	"strings"
	"unicode"
)

func AppendToSlice(vl interface{}) []interface{} {
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

func FormatTag(tagStr string) string {
	tagStr = strings.ReplaceAll(tagStr, "(", "")
	tagStr = strings.ReplaceAll(tagStr, ")", "")
	tagStr = strings.ReplaceAll(tagStr, ",", "")
	tagStr = strings.ToUpper(tagStr)
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, tagStr)
}

func FormatTagName(tagName string) string {
	tagName = strings.ToLower(tagName)
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, tagName)
}
