package vr

import (
	"github.com/okieraised/go2com/pkg/dicom/tag"
	"strings"
)

const (
	ApplicationEntity           = "AE"
	AgeString                   = "AS"
	AttributeTag                = "AT"
	CodeString                  = "CS"
	Date                        = "DA"
	DecimalString               = "DS"
	DateTime                    = "DT"
	FloatingPointSingle         = "FL"
	FloatingPointDouble         = "FD"
	IntegerString               = "IS"
	LongString                  = "LO"
	LongText                    = "LT"
	OtherByte                   = "OB"
	OtherDouble                 = "OD"
	OtherFloat                  = "OF"
	OtherLong                   = "OL"
	OtherVeryLong               = "VL"
	OtherWord                   = "OW"
	PersonName                  = "PN"
	ShortString                 = "SH"
	SignedLong                  = "SL"
	SequenceOfItems             = "SQ"
	SignedShort                 = "SS"
	ShortText                   = "ST"
	SignedVeryLong              = "SV"
	Time                        = "TM"
	UnlimitedCharacters         = "UC"
	UniqueIdentifier            = "UI"
	UnsignedLong                = "UL"
	Unknown                     = "UN"
	UniversalResourceIdentifier = "UR"
	UnsignedShort               = "US"
	UnlimitedText               = "UT"
	UnsignedVeryLong            = "UV"
	SignedShortOrUnsignedShort  = "XS"
	OtherByteOrOtherWord        = "OX"
)

type VRKind int

const (
	VRBytes VRKind = iota
	VRString
	VRUInt16
	VRUInt32
	VRInt16
	VRInt32
	VRFloat32
	VRFloat64
	VRSequence
	VRItem
	VRTagList
	VRDate
	VRPixelData
)

// GetVR returns the golang value encoding of an element with <tag, vr>.
func GetVR(dcmTag tag.DicomTag, vr string) VRKind {
	if dcmTag == tag.PixelData {
		return VRPixelData
	}
	switch vr {
	case Date:
		return VRDate
	case AttributeTag:
		return VRTagList
	case OtherWord, OtherByte, Unknown, OtherByteOrOtherWord, strings.ToLower(OtherByteOrOtherWord):
		return VRBytes
	case LongText, UnlimitedText:
		return VRString
	case UnsignedLong:
		return VRUInt32
	case SignedLong:
		return VRInt32
	case UnsignedShort:
		return VRUInt16
	case SignedShort, SignedShortOrUnsignedShort, strings.ToLower(SignedShortOrUnsignedShort):
		return VRInt16
	case FloatingPointSingle:
		return VRFloat32
	case FloatingPointDouble:
		return VRFloat64
	case SequenceOfItems:
		return VRSequence
	default:
		return VRString
	}
}
