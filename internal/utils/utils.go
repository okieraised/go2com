package utils

import (
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"runtime/pprof"
	"strings"
	"unicode"
)

func AppendToSlice(vl interface{}) []interface{} {
	res := make([]interface{}, 0)
	if vl == nil {
		return res
	}
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

func ReadDirRecursively(path string) ([]string, error) {
	result := []string{}

	visit := func(path string, f fs.DirEntry, err error) error {
		if !f.IsDir() {
			result = append(result, path)
		}
		return nil
	}
	err := filepath.WalkDir(path, visit)
	if err != nil {
		return result, err
	}
	return result, nil
}

func CPUProfilingFunc(fn func(), output string) error {
	f, err := os.Create(output)
	if err != nil {
		return err
	}
	defer f.Close()
	err = pprof.StartCPUProfile(f)
	if err != nil {
		return err
	}
	fn()
	pprof.StopCPUProfile()
	return nil
}
