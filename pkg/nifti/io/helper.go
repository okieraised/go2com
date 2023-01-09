package io

import "github.com/okieraised/go2com/pkg/nifti/constant"

func IsValidDatatype(dType int32) bool {
	if constant.ValidDatatype[dType] {
		return true
	}
	return false
}
