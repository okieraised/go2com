package constant

const (
	NIFTI_INTENT_CORREL                     int16 = 2
	NIFTI_INTENT_TTEST                      int16 = 3
	NIFTI_INTENT_FTEST                      int16 = 4
	NIFTI_INTENT_ZSCORE                     int16 = 5
	NIFTI_INTENT_CHISQ                      int16 = 6
	NIFTI_INTENT_BETA                       int16 = 7
	NIFTI_INTENT_BINOM                      int16 = 8
	NIFTI_INTENT_GAMMA                      int16 = 9
	NIFTI_INTENT_POISSON                    int16 = 10
	NIFTI_INTENT_NORMAL                     int16 = 11
	NIFTI_INTENT_FTEST_NONC                 int16 = 12
	NIFTI_INTENT_CHISQ_NONC                 int16 = 13
	NIFTI_INTENT_LOGISTIC                   int16 = 14
	NIFTI_INTENT_LAPLACE                    int16 = 15
	NIFTI_INTENT_UNIFORM                    int16 = 16
	NIFTI_INTENT_TTEST_NONC                 int16 = 17
	NIFTI_INTENT_WEIBULL                    int16 = 18
	NIFTI_INTENT_CHI                        int16 = 19
	NIFTI_INTENT_INVGAUSS                   int16 = 20
	NIFTI_INTENT_EXTVAL                     int16 = 21
	NIFTI_INTENT_PVAL                       int16 = 22
	NIFTI_INTENT_LOGPVAL                    int16 = 23
	NIFTI_INTENT_LOG10PVAL                  int16 = 24
	NIFTI_INTENT_ESTIMATE                   int16 = 1001
	NIFTI_INTENT_LABEL                      int16 = 1002
	NIFTI_INTENT_NEURONAME                  int16 = 1003
	NIFTI_INTENT_GENMATRIX                  int16 = 1004
	NIFTI_INTENT_SYMMATRIX                  int16 = 1005
	NIFTI_INTENT_DISPVECT                   int16 = 1006 /* specifically for displacements */
	NIFTI_INTENT_VECTOR                     int16 = 1007 /* for any other type of vector */
	NIFTI_INTENT_POINTSET                   int16 = 1008
	NIFTI_INTENT_TRIANGLE                   int16 = 1009
	NIFTI_INTENT_QUATERNION                 int16 = 1010
	NIFTI_INTENT_DIMLESS                    int16 = 1011
	NIFTI_INTENT_TIME_SERIES                int16 = 2001
	NIFTI_INTENT_NODE_INDEX                 int16 = 2002
	NIFTI_INTENT_RGB_VECTOR                 int16 = 2003
	NIFTI_INTENT_RGBA_VECTOR                int16 = 2004
	NIFTI_INTENT_SHAPE                      int16 = 2005
	FSL_FNIRT_DISPLACEMENT_FIELD            int16 = 2006
	FSL_CUBIC_SPLINE_COEFFICIENTS           int16 = 2007
	FSL_DCT_COEFFICIENTS                    int16 = 2008
	FSL_QUADRATIC_SPLINE_COEFFICIENTS       int16 = 2009
	FSL_TOPUP_CUBIC_SPLINE_COEFFICIENTS     int16 = 2016
	FSL_TOPUP_QUADRATIC_SPLINE_COEFFICIENTS int16 = 2017
	FSL_TOPUP_FIELD                         int16 = 2018
)

const (
	NIFTI_TYPE_UINT8      = 2
	NIFTI_TYPE_INT16      = 4
	NIFTI_TYPE_INT32      = 8
	NIFTI_TYPE_FLOAT32    = 16
	NIFTI_TYPE_COMPLEX64  = 32
	NIFTI_TYPE_FLOAT64    = 64
	NIFTI_TYPE_RGB24      = 128
	NIFTI_TYPE_INT8       = 256
	NIFTI_TYPE_UINT16     = 512
	NIFTI_TYPE_UINT32     = 768
	NIFTI_TYPE_INT64      = 1024
	NIFTI_TYPE_UINT64     = 1280
	NIFTI_TYPE_FLOAT128   = 1536
	NIFTI_TYPE_COMPLEX128 = 1792
	NIFTI_TYPE_RGBA32     = 2304
	NIFTI_TYPE_COMPLEX256 = 2048
)

const (
	NIFTI_XFORM_UNKNOWN      = 0
	NIFTI_XFORM_SCANNER_ANAT = 1
	NIFTI_XFORM_ALIGNED_ANAT = 2
	NIFTI_XFORM_TALAIRACH    = 3
	NIFTI_XFORM_MNI_152      = 4
)

var NiiPatientOrientationInfo = map[uint8]string{
	NIFTI_XFORM_UNKNOWN:      "Unknown",
	NIFTI_XFORM_SCANNER_ANAT: "Scanner-based anatomical coordinates",
	NIFTI_XFORM_ALIGNED_ANAT: "Coordinates aligned to another file, or to the truth",
	NIFTI_XFORM_TALAIRACH:    "Coordinates aligned to the Talairach space",
	NIFTI_XFORM_MNI_152:      "Coordinates aligned to the mni space",
}

const (
	NIFTI_SLICE_UNKNOWN  uint8 = 0
	NIFTI_SLICE_SEQ_INC  uint8 = 1
	NIFTI_SLICE_SEQ_DEC  uint8 = 2
	NIFTI_SLICE_ALT_INC  uint8 = 3
	NIFTI_SLICE_ALT_DEC  uint8 = 4
	NIFTI_SLICE_ALT_INC2 uint8 = 5 /* 05 May 2005: RWCox */
	NIFTI_SLICE_ALT_DEC2 uint8 = 6 /* 05 May 2005: RWCox */
)

var NiiSliceAcquistionInfo = map[uint8]string{
	NIFTI_SLICE_UNKNOWN:  "Unknown",
	NIFTI_SLICE_SEQ_INC:  "Sequential, increasing",
	NIFTI_SLICE_SEQ_DEC:  "Sequential, decreasing",
	NIFTI_SLICE_ALT_INC:  "Interleaved, increasing, starting at the 1st mri slice",
	NIFTI_SLICE_ALT_DEC:  "Interleaved, decreasing, starting at the last mri slice",
	NIFTI_SLICE_ALT_INC2: "Interleaved, increasing, starting at the 2nd mri slice",
	NIFTI_SLICE_ALT_DEC2: "Interleaved, decreasing, starting at one before the last mri slice",
}

const (
	NIFTI_UNITS_UNKNOWN uint8 = 0
	NIFTI_UNITS_METER   uint8 = 1
	NIFTI_UNITS_MM      uint8 = 2
	NIFTI_UNITS_MICRON  uint8 = 3
	NIFTI_UNITS_SEC     uint8 = 8
	NIFTI_UNITS_MSEC    uint8 = 16
	NIFTI_UNITS_USEC    uint8 = 24
	NIFTI_UNITS_HZ      uint8 = 32
	NIFTI_UNITS_PPM     uint8 = 40
	NIFTI_UNITS_RADS    uint8 = 48
)

var NiiMeasurementUnits = map[uint8]string{
	NIFTI_UNITS_UNKNOWN: "unknown",
	NIFTI_UNITS_METER:   "meter",
	NIFTI_UNITS_MM:      "millimeter",
	NIFTI_UNITS_MICRON:  "micrometer",
	NIFTI_UNITS_SEC:     "second",
	NIFTI_UNITS_MSEC:    "millisecond",
	NIFTI_UNITS_USEC:    "microsecond",
	NIFTI_UNITS_HZ:      "Hertz",
	NIFTI_UNITS_PPM:     "ppm",
	NIFTI_UNITS_RADS:    "rad/s",
}