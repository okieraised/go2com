package io

import "github.com/okieraised/go2com/pkg/nifti/constant"

// IsValidDatatype checks whether the datatype is valid for NIFTI format
func IsValidDatatype(datatype int32) bool {
	if constant.ValidDatatype[datatype] {
		return true
	}
	return false
}

// assignDatatypeSize sets the number of bytes per voxel and the swapsize based on a datatype code
func assignDatatypeSize(datatype int32) (int16, int16) {
	var nByper, swapSize int16
	switch datatype {
	case constant.DT_INT8, constant.DT_UINT8:
		nByper = 1
		swapSize = 0
	case constant.DT_INT16, constant.DT_UINT16:
		nByper = 2
		swapSize = 2
	case constant.DT_RGB24:
		nByper = 3
		swapSize = 0
	case constant.DT_RGBA32:
		nByper = 4
		swapSize = 0
	case constant.DT_INT32, constant.DT_UINT32, constant.DT_FLOAT32:
		nByper = 4
		swapSize = 4
	case constant.DT_COMPLEX64:
		nByper = 8
		swapSize = 4
	case constant.DT_FLOAT64, constant.DT_INT64, constant.DT_UINT64:
		nByper = 8
		swapSize = 8
	case constant.DT_FLOAT128:
		nByper = 16
		swapSize = 16
	case constant.DT_COMPLEX128:
		nByper = 16
		swapSize = 8
	case constant.DT_COMPLEX256:
		nByper = 32
		swapSize = 16
	default:
	}
	return nByper, swapSize
}

// needHeaderSwap checks whether byte swapping is needed. dim0 should be in [0,7], and headerSize should be accurate.
//
// Returns:
//
// > 0 : needs swap
//
// = 0 : does not need swap
//
// < 0 : error condition
func needHeaderSwap(dim0 int16, headerSize int) int {
	//d0 := dim0
	//hSize := headerSize
	//
	//if d0 != 0 {
	//	if d0 > 0 && d0 < 7 {
	//		return 0
	//	}
	//
	//}
	return -2
}

// static int need_nhdr_swap( short dim0, int hdrsize )
//{
//   short d0    = dim0;     /* so we won't have to swap them on the stack */
//   int   hsize = hdrsize;
//
//   if( d0 != 0 ){     /* then use it for the check */
//      if( d0 > 0 && d0 <= 7 ) return 0;
//
//      nifti_swap_2bytes(1, &d0);        /* swap? */
//      if( d0 > 0 && d0 <= 7 ) return 1;
//
//      if( g_opts.debug > 1 ){
//         fprintf(stderr,"** NIFTI: bad swapped d0 = %d, unswapped = ", d0);
//         nifti_swap_2bytes(1, &d0);        /* swap? */
//         fprintf(stderr,"%d\n", d0);
//      }
//
//      return -1;        /* bad, naughty d0 */
//   }
//
//   /* dim[0] == 0 should not happen, but could, so try hdrsize */
//   if( hsize == sizeof(nifti_1_header) ) return 0;
//
//   nifti_swap_4bytes(1, &hsize);     /* swap? */
//   if( hsize == sizeof(nifti_1_header) ) return 1;
//
//   if( g_opts.debug > 1 ){
//      fprintf(stderr,"** NIFTI: bad swapped hsize = %d, unswapped = ", hsize);
//      nifti_swap_4bytes(1, &hsize);        /* swap? */
//      fprintf(stderr,"%d\n", hsize);
//   }
//
//   return -2;     /* bad, naughty hsize */
//}

// Swap byte array

// niftiSwap2Bytes swaps 2 bytes at a time
func niftiSwap2Bytes(size int, arr []uint8) {
	for i := 0; i < size; i++ {

	}
}

// niftiSwap4Bytes swaps 4 bytes at a time
func niftiSwap4Bytes(size int, arr []uint8) {
	for i := 0; i < size; i++ {

	}
}

// niftiSwap8Bytes swaps 8 bytes at a time
func niftiSwap8Bytes(size int, arr []uint8) {
	for i := 0; i < size; i++ {

	}
}

// niftiSwap16Bytes swaps 16 bytes at a time
func niftiSwap16Bytes(size int, arr []uint8) {
	for i := 0; i < size; i++ {

	}
}
