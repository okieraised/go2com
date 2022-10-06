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
	case ExplicitVRLittleEndian, JPEGBaselineProcess1, JPEG2000ImageCompressionLosslessOnly:
		return binary.LittleEndian, false, nil
	case ExplicitVRBigEndian:
		return binary.BigEndian, false, nil
	default:
		return binary.BigEndian, false, fmt.Errorf("unsupported transfer syntax: %v", uid)
	}
}
