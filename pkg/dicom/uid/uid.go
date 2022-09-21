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
	case ExplicitVRLittleEndian:
		return binary.LittleEndian, false, nil
	case ExplicitVRBigEndian:
		return binary.BigEndian, false, nil
	default:
		return binary.BigEndian, false, fmt.Errorf("invalid or unknown transfer syntax: %v", uid)
	}
}
