package utils

import "reflect"

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
