package vr

import "github.com/okieraised/go2com/pkg/dicom/tag"

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
)

type VRKind int

const (
	// VRStringList means the element stores a list of strings
	VRStringList VRKind = iota
	VRBytes
	VRString
	VRUInt16List
	VRUInt32List
	VRInt16List
	VRInt32List
	VRFloat32List
	VRFloat64List
	VRSequence
	VRItem
	VRTagList
	VRDate
	VRPixelData
)

// GetVR returns the golang value encoding of an element with <tag, vr>.
func GetVR(dcmTag tag.DicomTag, vr string) VRKind {
	// if dcmTag == Item {
	// 	return VRItem
	// }

	if dcmTag == tag.PixelData {
		return VRPixelData
	}
	switch vr {
	case "DA":
		return VRDate
	case "AT":
		return VRTagList
	case "OW", "OB", "UN":
		return VRBytes
	case "LT", "UT":
		return VRString
	case "UL":
		return VRUInt32List
	case "SL":
		return VRInt32List
	case "US":
		return VRUInt16List
	case "SS":
		return VRInt16List
	case "FL":
		return VRFloat32List
	case "FD":
		return VRFloat64List
	case "SQ":
		return VRSequence
	default:
		return VRStringList
	}
}
