package constant

const (
	NIFTI_EXT            = ".nii"
	NIFTI_COMPRESSED_EXT = ".gz"
)

const (
	NIIVersion1 = iota + 1
	NIIVersion2
)

const (
	NII1HeaderSize = 348
	NII2HeaderSize = 540
)

//// Possible NIFTI image extension
//const (
//	NIFTI_FTYPE_NIFTI1_1     = ".nii"
//	NIFTI_FTYPE_NIFTI1_2     = ".img"
//	NIFTI_FTYPE_NIFTI1_2_HDR = ".hdr"
//	NIFTI_FTYPE_NIFTI_GZIP   = ".gz"
//)
//
//var ValidNIfTIExtMapper = map[string]bool{
//	NIFTI_FTYPE_NIFTI1_1:     true,
//	NIFTI_FTYPE_NIFTI1_2:     true,
//	NIFTI_FTYPE_NIFTI1_2_HDR: true,
//	NIFTI_FTYPE_NIFTI_GZIP:   true,
//}

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
	NIFTI_UNKNOWN_ORIENT = 0
	NIFTI_L2R            = 1
	NIFTI_R2L            = 2
	NIFTI_P2A            = 3
	NIFTI_A2P            = 4
	NIFTI_I2S            = 5
	NIFTI_S2I            = 6
)

var OrietationToString = map[int]string{
	NIFTI_UNKNOWN_ORIENT: "Unknown",
	NIFTI_L2R:            "Left-to-Right",
	NIFTI_R2L:            "Right-to-Left",
	NIFTI_P2A:            "Posterior-to-Anterior",
	NIFTI_A2P:            "Anterior-to-Posterior",
	NIFTI_I2S:            "Inferior-to-Superior",
	NIFTI_S2I:            "Superior-to-Inferior",
}

const (
	DT_UNKNOWN    int32 = 0    // what it says, dude
	DT_BINARY     int32 = 1    // binary (1 bit/voxel)
	DT_UINT8      int32 = 2    // unsigned char (8 bits/voxel)
	DT_INT16      int32 = 4    // signed short (16 bits/voxel)
	DT_INT32      int32 = 8    // signed int (32 bits/voxel)
	DT_FLOAT32    int32 = 16   // float (32 bits/voxel)
	DT_COMPLEX64  int32 = 32   // complex (64 bits/voxel)
	DT_FLOAT64    int32 = 64   // double (64 bits/voxel)
	DT_RGB24      int32 = 128  // RGB triple (24 bits/voxel)
	DT_ALL        int32 = 255  // not very useful (?)
	DT_INT8       int32 = 256  // signed char (8 bits)
	DT_UINT16     int32 = 512  // unsigned short (16 bits)
	DT_UINT32     int32 = 768  // unsigned int (32 bits)
	DT_INT64      int32 = 1024 // long long (64 bits)
	DT_UINT64     int32 = 1280 // unsigned long long (64 bits)
	DT_FLOAT128   int32 = 1536 // long double (128 bits)
	DT_COMPLEX128 int32 = 1792 // double pair (128 bits)
	DT_COMPLEX256 int32 = 2048 // long double pair (256 bits)
	DT_RGBA32     int32 = 2304
)

var ValidDatatype = map[int32]bool{
	DT_UNKNOWN:    true,
	DT_BINARY:     true,
	DT_INT8:       true,
	DT_UINT8:      true,
	DT_INT16:      true,
	DT_UINT16:     true,
	DT_INT32:      true,
	DT_UINT32:     true,
	DT_INT64:      true,
	DT_UINT64:     true,
	DT_FLOAT32:    true,
	DT_FLOAT64:    true,
	DT_ALL:        true,
	DT_FLOAT128:   true,
	DT_COMPLEX64:  true,
	DT_COMPLEX128: true,
	DT_COMPLEX256: true,
	DT_RGB24:      true,
	DT_RGBA32:     true,
}

var IsDatatypeInt = map[int32]bool{
	DT_UNKNOWN:    false,
	DT_BINARY:     false,
	DT_INT8:       true,
	DT_UINT8:      true,
	DT_INT16:      true,
	DT_UINT16:     true,
	DT_INT32:      true,
	DT_UINT32:     true,
	DT_INT64:      true,
	DT_UINT64:     true,
	DT_FLOAT32:    false,
	DT_FLOAT64:    false,
	DT_FLOAT128:   false,
	DT_COMPLEX64:  false,
	DT_COMPLEX128: false,
	DT_COMPLEX256: false,
	DT_RGB24:      true,
	DT_RGBA32:     true,
}

const (
	NIFTI_XFORM_UNKNOWN      = 0
	NIFTI_XFORM_SCANNER_ANAT = 1
	NIFTI_XFORM_ALIGNED_ANAT = 2
	NIFTI_XFORM_TALAIRACH    = 3
	NIFTI_XFORM_MNI_152      = 4
)

var NiiPatientOrientationInfo = map[int32]string{
	NIFTI_XFORM_UNKNOWN:      "0: Unknown",
	NIFTI_XFORM_SCANNER_ANAT: "1: Scanner Anat",
	NIFTI_XFORM_ALIGNED_ANAT: "2: Aligned Anat",
	NIFTI_XFORM_TALAIRACH:    "3: Talairach",
	NIFTI_XFORM_MNI_152:      "4: MNI",
}

const (
	NIFTI_SLICE_UNKNOWN  int32 = 0
	NIFTI_SLICE_SEQ_INC  int32 = 1
	NIFTI_SLICE_SEQ_DEC  int32 = 2
	NIFTI_SLICE_ALT_INC  int32 = 3
	NIFTI_SLICE_ALT_DEC  int32 = 4
	NIFTI_SLICE_ALT_INC2 int32 = 5 /* 05 May 2005: RWCox */
	NIFTI_SLICE_ALT_DEC2 int32 = 6 /* 05 May 2005: RWCox */
)

var NiiSliceAcquistionInfo = map[int32]string{
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
