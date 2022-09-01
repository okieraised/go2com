package transfersyntax

var TransferSyntaxMap = map[string]string{
	"1.2.840.10008.1.2":      "ImplicitVRLittleEndian",
	"1.2.840.10008.1.2.1":    "ExplicitVRLittleEndian",
	"1.2.840.10008.1.2.1.99": "DeflatedExplicitVRLittleEndian",
	"1.2.840.10008.1.2.2":    "ExplicitVRBigEndian",
	"1.2.840.10008.1.2.4.50": "JPEGBaseline8Bit",
	"1.2.840.10008.1.2.4.51": "JPEGExtended12Bit",
	"1.2.840.10008.1.2.4.52": "JPEGExtended35",
	"1.2.840.10008.1.2.4.53": "JPEGSpectralSelectionNonHierarchical68",
	"1.2.840.10008.1.2.4.54": "JPEGSpectralSelectionNonHierarchical79",
	"1.2.840.10008.1.2.4.55": "JPEGSpectralSelectionFullProgression1012",
	"1.2.840.10008.1.2.4.56": "JPEGSpectralSelectionFullProgression1113",
	"1.2.840.10008.1.2.4.57": "JPEGGLossless",
	"1.2.840.10008.1.2.4.58": "JPEGGLosslessNonHierarchical15",
	"1.2.840.10008.1.2.4.59": "JPEGExtendedHierarchical1618",
}

const (

// "1.2.840.10008.1.2.4.60"=	JPEG Extended, Hierarchical (Processes 17 & 19)	Retired
// "1.2.840.10008.1.2.4.61"=	JPEG Spectral Selection, Hierarchical (Processes 20 & 22)	Retired
// "1.2.840.10008.1.2.4.62"=	JPEG Spectral Selection, Hierarchical (Processes 21 & 23)	Retired
// "1.2.840.10008.1.2.4.63"=	JPEG Full Progression, Hierarchical (Processes 24 & 26)	Retired
// "1.2.840.10008.1.2.4.64"=	JPEG Full Progression, Hierarchical (Processes 25 & 27)	Retired
// "1.2.840.10008.1.2.4.65"=	JPEG Lossless, Nonhierarchical (Process 28)	Retired
// "1.2.840.10008.1.2.4.66"=	JPEG Lossless, Nonhierarchical (Process 29)	Retired
// "1.2.840.10008.1.2.4.70"=	JPEG Lossless, Nonhierarchical, First- Order Prediction
// "1.2.840.10008.1.2.4.80"=	JPEG-LS Lossless Image Compression
// "1.2.840.10008.1.2.4.81"=	JPEG-LS Lossy (Near- Lossless) Image Compression
// "1.2.840.10008.1.2.4.90"=	JPEG 2000 Image Compression (Lossless Only)
// "1.2.840.10008.1.2.4.91"=	JPEG 2000 Image Compression
// "1.2.840.10008.1.2.4.92"=	JPEG 2000 Part 2 Multicomponent Image Compression (Lossless Only)
// "1.2.840.10008.1.2.4.93"=	JPEG 2000 Part 2 Multicomponent Image Compression
// "1.2.840.10008.1.2.4.94"=	JPIP Referenced
// "1.2.840.10008.1.2.4.95"=	JPIP Referenced Deflate
// "1.2.840.10008.1.2.5"=	RLE Lossless
// "1.2.840.10008.1.2.6.1"=	RFC 2557 MIME Encapsulation
// "1.2.840.10008.1.2.4.100"=	MPEG2 Main Profile Main Level
// "1.2.840.10008.1.2.4.102"=	MPEG-4 AVC/H.264 High Profile / Level 4.1
// "1.2.840.10008.1.2.4.103"=	"MPEG4​HP41BD"
)
