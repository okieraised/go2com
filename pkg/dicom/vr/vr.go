package vr

const (
	FixedLength   = "fixed"
	MaximumLength = "maximum"
	MinimumLength = "minimum"
	AnyLength     = "any"
	NotApplicable = "na"
)

// Value Representation (VR) constants. Referenced link at:
// https://dicom.nema.org/dicom/2013/output/chtml/part05/sect_6.2.html
const (
	ApplicationEntity   = "AE"
	AgeString           = "AS"
	AttributeTag        = "AT"
	CodeString          = "CS"
	Date                = "DA"
	DecimalString       = "DS"
	DateTime            = "DT"
	FloatingPointSingle = "FL"
	FloatingPointDouble = "FD"
	IntegerString       = "IS"
	LongString          = "LO"
	LongText            = "LT"
	OtherByteString     = "OB"
	OtherDoubleString   = "OD"
	OtherFloatString    = "OF"
	OtherWordString     = "OW"
	PersonName          = "PN"
	ShortString         = "SH"
	SignedLong          = "SL"
	SequenceOfItems     = "SQ"
	SignedShort         = "SS"
	ShortText           = "ST"
	Time                = "TM"
	UniqueIdentifier    = "UI"
	UnsignedLong        = "UL"
	Unknown             = "UN"
	UnsignedShort       = "US"
	UnlimitedText       = "UT"
)

type VRInfo struct {
	Definition        string
	LengthOfValue     uint32
	LengthDescription string
}

var VRMapper = map[string]VRInfo{
	"AE": {
		Definition:        "A string of characters that identifies an Application Entity with leading and trailing spaces (20H) being non-significant. A value consisting solely of spaces shall not be used.",
		LengthOfValue:     16,
		LengthDescription: MaximumLength,
	},
	"AS": {
		Definition:        "A string of characters with one of the following formats -- nnnD, nnnW, nnnM, nnnY; where nnn shall contain the number of days for D, weeks for W, months for M, or years for Y.",
		LengthOfValue:     4,
		LengthDescription: FixedLength,
	},
	"AT": {
		Definition:        "Ordered pair of 16-bit unsigned integers that is the value of a Data Element Tag.",
		LengthOfValue:     4,
		LengthDescription: FixedLength,
	},
	"CS": {
		Definition:        "A string of characters with leading or trailing spaces (20H) being non-significant.",
		LengthOfValue:     16,
		LengthDescription: MaximumLength,
	},
	"DA": {
		Definition:        "A string of characters of the format YYYYMMDD; where YYYY shall contain year, MM shall contain the month, and DD shall contain the day, interpreted as a date of the Gregorian calendar system.",
		LengthOfValue:     8,
		LengthDescription: FixedLength,
	},
	"DS": {
		Definition:        "A string of characters representing either a fixed point number or a floating point number. A fixed point number shall contain only the characters 0-9 with an optional leading \"+\" or \"-\" and an optional \".\" to mark the decimal point. A floating point number shall be conveyed as defined in ANSI X3.9, with an \"E\" or \"e\" to indicate the start of the exponent. Decimal Strings may be padded with leading or trailing spaces. Embedded spaces are not allowed.",
		LengthOfValue:     16,
		LengthDescription: MaximumLength,
	},
	"DT": {
		Definition:        "A concatenated date-time character string in the format: YYYYMMDDHHMMSS.FFFFFF&ZZXX",
		LengthOfValue:     26,
		LengthDescription: MaximumLength,
	},
	"FL": {
		Definition:        "Single precision binary floating point number represented in IEEE 754:1985 32-bit Floating Point Number Format.",
		LengthOfValue:     4,
		LengthDescription: FixedLength,
	},
	"FD": {
		Definition:        "Double precision binary floating point number represented in IEEE 754:1985 64-bit Floating Point Number Format.",
		LengthOfValue:     8,
		LengthDescription: FixedLength,
	},
	"IS": {
		Definition:        "A string of characters representing an Integer in base-10 (decimal), shall contain only the characters 0 - 9, with an optional leading \"+\" or \"-\". It may be padded with leading and/or trailing spaces. Embedded spaces are not allowed.",
		LengthOfValue:     12,
		LengthDescription: MaximumLength,
	},
	"LO": {
		Definition:        "A character string that may be padded with leading and/or trailing spaces. The character code 5CH (the BACKSLASH \"\\\" in ISO-IR 6) shall not be present, as it is used as the delimiter between values in multiple valued data elements. The string shall not have Control Characters except for ESC.",
		LengthOfValue:     64,
		LengthDescription: MinimumLength,
	},
	"LT": {
		Definition:        "A character string that may contain one or more paragraphs. It may contain the Graphic Character set and the Control Characters, CR, LF, FF, and ESC. It may be padded with trailing spaces, which may be ignored, but leading spaces are considered to be significant. Data Elements with this VR shall not be multi-valued and therefore character code 5CH (the BACKSLASH \"\\\" in ISO-IR 6) may be used.",
		LengthOfValue:     10240,
		LengthDescription: MaximumLength,
	},
	"OB": {
		Definition:        "A string of bytes where the encoding of the contents is specified by the negotiated Transfer Syntax. OB is a VR that is insensitive to Little/Big Endian byte ordering. The string of bytes shall be padded with a single trailing NULL byte value (00H) when necessary to achieve even length.",
		LengthOfValue:     0,
		LengthDescription: NotApplicable,
	},
	"OD": {
		Definition:        "A string of 64-bit IEEE 754:1985 floating point words. OD is a VR that requires byte swapping within each 64-bit word when changing between Little Endian and Big Endian byte ordering",
		LengthOfValue:     4294967288,
		LengthDescription: MaximumLength,
	},
	"OF": {
		Definition:        "A string of 32-bit IEEE 754:1985 floating point words. OF is a VR that requires byte swapping within each 32-bit word when changing between Little Endian and Big Endian byte ordering",
		LengthOfValue:     4294967292,
		LengthDescription: MaximumLength,
	},
	"OW": {
		Definition:        "A string of 16-bit words where the encoding of the contents is specified by the negotiated Transfer Syntax. OW is a VR that requires byte swapping within each word when changing between Little Endian and Big Endian byte ordering",
		LengthOfValue:     0,
		LengthDescription: NotApplicable,
	},
	"PN": {
		Definition:        "A character string encoded using a 5 component convention. The character code 5CH (the BACKSLASH \"\\\" in ISO-IR 6) shall not be present, as it is used as the delimiter between values in multiple valued data elements. The string may be padded with trailing spaces. For human use, the five components in their order of occurrence are: family name complex, given name complex, middle name, name prefix, name suffix.",
		LengthOfValue:     64,
		LengthDescription: NotApplicable,
	},
	"SH": {
		Definition:        "A character string that may be padded with leading and/or trailing spaces. The character code 05CH (the BACKSLASH \"\\\" in ISO-IR 6) shall not be present, as it is used as the delimiter between values for multiple data elements. The string shall not have Control Characters except ESC.",
		LengthOfValue:     16,
		LengthDescription: FixedLength,
	},
	"SL": {
		Definition:        "Signed binary integer 32 bits long in 2's complement form. Represents an integer, n, in the range: -2^31 <= n <= 2^31-1.",
		LengthOfValue:     4,
		LengthDescription: FixedLength,
	},
	"SQ": {
		Definition:        "Value is a Sequence of zero or more Items",
		LengthOfValue:     0,
		LengthDescription: NotApplicable,
	},
	"SS": {
		Definition:        "Signed binary integer 16 bits long in 2's complement form. Represents an integer n in the range: -2^15 <= n <= 2^15-1.",
		LengthOfValue:     2,
		LengthDescription: FixedLength,
	},
	"ST": {
		Definition:        "A character string that may contain one or more paragraphs. It may contain the Graphic Character set and the Control Characters, CR, LF, FF, and ESC. It may be padded with trailing spaces, which may be ignored, but leading spaces are considered to be significant. Data Elements with this VR shall not be multi-valued and therefore character code 5CH (the BACKSLASH \"\\\" in ISO-IR 6) may be used.",
		LengthOfValue:     1024,
		LengthDescription: MaximumLength,
	},
	"TM": {
		Definition:        "A string of characters of the format HHMMSS.FFFFFF; where HH contains hours (range \"00\" - \"23\"), MM contains minutes (range \"00\" - \"59\"), SS contains seconds (range \"00\" - \"60\"), and FFFFFF contains a fractional part of a second as small as 1 millionth of a second (range \"000000\" - \"999999\"). A 24-hour clock is used. Midnight shall be represented by only \"0000\" since \"2400\" would violate the hour range. The string may be padded with trailing spaces. Leading and embedded spaces are not allowed.",
		LengthOfValue:     16,
		LengthDescription: MaximumLength,
	},
	"UI": {
		Definition:        "A character string containing a UID that is used to uniquely identify a wide variety of items. The UID is a series of numeric components separated by the period \".\" character. If a Value Field containing one or more UIDs is an odd number of bytes in length, the Value Field shall be padded with a single trailing NULL (00H) character to ensure that the Value Field is an even number of bytes in length.",
		LengthOfValue:     64,
		LengthDescription: MaximumLength,
	},
	"UL": {
		Definition:        "Unsigned binary integer 32 bits long. Represents an integer n in the range: 0 <= n < 2^32.",
		LengthOfValue:     4,
		LengthDescription: FixedLength,
	},
	"UN": {
		Definition:        "A string of bytes where the encoding of the contents is unknown",
		LengthOfValue:     0,
		LengthDescription: AnyLength,
	},
	"US": {
		Definition:        "Unsigned binary integer 16 bits long. Represents integer n in the range: 0 <= n < 2^16.",
		LengthOfValue:     2,
		LengthDescription: FixedLength,
	},
	"UT": {
		Definition:        "A character string that may contain one or more paragraphs. It may contain the Graphic Character set and the Control Characters, CR, LF, FF, and ESC. It may be padded with trailing spaces, which may be ignored, but leading spaces are considered to be significant. Data Elements with this VR shall not be multi-valued and therefore character code 5CH (the BACKSLASH \"\\\" in ISO-IR 6) may be used.",
		LengthOfValue:     4294967294,
		LengthDescription: MaximumLength,
	},
}
