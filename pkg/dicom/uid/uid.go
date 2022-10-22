package uid

import (
	"encoding/binary"
	"fmt"
)

func ParseTransferSyntaxUID(uid string) (bo binary.ByteOrder, implicit bool, err error) {
	switch uid {
	case ImplicitVRLittleEndian:
		return binary.LittleEndian, true, nil
	case DeflatedExplicitVRLittleEndian:
		fallthrough
	case ExplicitVRLittleEndian, JPEGBaselineProcess1, JPEG2000ImageCompressionLosslessOnly,
		JPEGBaselineProcess2And4, JPEGLosslessNonHierarchicalProcesses14, JPEGLSLosslessImageCompression,
		JPEGLSLossyNearLosslessImageCompression, JPEG2000ImageCompression, MPEG4AVCH264highProfile, MPEG4AVCH264BDCompatibleHighProfile,
		JPEGLosslessNonHierarchicalFirstOrderPredictionProcess14, RLELossless:
		return binary.LittleEndian, false, nil
	case ExplicitVRBigEndian:
		return binary.BigEndian, false, nil
	default:
		return binary.BigEndian, false, fmt.Errorf("unsupported transfer syntax: %v", uid)
	}
}
