package uid

import (
	"encoding/binary"
	"fmt"
)

func RetrieveTransferSyntaxUID(uid string) (string, error) {
	// defaults are explicit VR, little endian
	switch uid {
	case ImplicitVRLittleEndian, ExplicitVRLittleEndian, ExplicitVRBigEndian, DeflatedExplicitVRLittleEndian:
		return uid, nil
	default:
		e, err := Lookup(uid)
		if err != nil {
			return "", err
		}
		if e.Type != TypeTransferSyntax {
			return "", fmt.Errorf("dicom.CanonicalTransferSyntaxUID: '%s' is not a transfer syntax (is %s)", uid, e.Type)
		}
		// The default is ExplicitVRLittleEndian
		return ExplicitVRLittleEndian, nil
	}
}

func ParseTransferSyntaxUID(uid string) (bo binary.ByteOrder, isImplicit bool, err error) {
	canonical, err := RetrieveTransferSyntaxUID(uid)
	if err != nil {
		return nil, false, err
	}
	switch canonical {
	case ImplicitVRLittleEndian:
		return binary.LittleEndian, true, nil
	case DeflatedExplicitVRLittleEndian:
		fallthrough
	case ExplicitVRLittleEndian:
		return binary.LittleEndian, false, nil
	case ExplicitVRBigEndian:
		return binary.BigEndian, false, nil
	default:
		return binary.BigEndian, false, fmt.Errorf("invalid or unknown transfer syntax: %v,  %v", canonical, uid)
	}
}
