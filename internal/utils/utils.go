package utils

import (
	"reflect"
	"strings"
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
	tagStr = strings.TrimSpace(tagStr)
	tagStr = strings.ToUpper(tagStr)

	return tagStr
}
