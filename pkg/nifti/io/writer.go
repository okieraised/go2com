package io

import (
	"errors"
	"fmt"
	"github.com/okieraised/go2com/pkg/nifti/constant"
	"math"
)

type NiiWriter interface {
}

type niiWriter struct {
	niiData *Nii
}

func (w *niiWriter) writeToFile(fileName string) {

}

func (w *niiWriter) makeNewHeader() {

}

func MakeNewNii1Header(inDim *[8]int16, inDatatype int32) *Nii1Header {

	// Default Dim value
	defaultDim := [8]int16{3, 1, 1, 1, 0, 0, 0, 0}

	header := new(Nii1Header)
	var dim [8]int16

	// If no input Dim is provided then we use the default value
	if inDim != nil {
		dim = *inDim
	} else {
		dim = defaultDim
	}

	// validate Dim: if there is any problem, apply default Dim
	if dim[0] < 0 || dim[0] > 7 {
		dim = defaultDim
	} else {
		for c := 1; c <= int(dim[0]); c++ {
			if dim[c] < 1 {
				fmt.Printf("bad dim: %d: %d\n", c, dim[c])
				dim = defaultDim
				break
			}
		}
	}

	// Validate datatype
	datatype := inDatatype
	if !IsValidDatatype(datatype) {
		datatype = constant.DT_FLOAT32
	}

	// Populate the header struct
	header.SizeofHdr = NII1HeaderSize
	header.Regular = 'r'

	// Init dim and pixdim
	header.Dim[0] = dim[0]
	header.Pixdim[0] = 0.0
	for c := 1; c <= int(dim[0]); c++ {
		header.Dim[c] = dim[c]
		header.Pixdim[c] = 1.0
	}

	header.Datatype = int16(datatype)

	nByper, _ := assignDatatypeSize(datatype)
	header.Bitpix = 8 * nByper
	header.Magic = [4]byte{110, 43, 49, 0}

	return header
}

func MakeNewImage(inDim *[8]int16, inDatatype int32, dataFill int) {
	//header := MakeNewNii1Header(inDim, inDatatype)

}

func converHeaderToImage(header *Nii1Header) error {
	// Check if we need to swap bytes
	needSwap := needHeaderSwap(header.Dim[0])
	if needSwap < 0 {
		return fmt.Errorf("bad dim[0] value of %d", header.Dim[0])
	}
	// Swap the header if we need to
	if needSwap > 0 {

	}

	if int32(header.Datatype) == constant.DT_UNKNOWN || int32(header.Datatype) == constant.DT_BINARY {
		return errors.New("bad datatype")
	}

	if header.Dim[0] < 0 {
		return errors.New("bad dim[1]")
	}

	// fix bad []Dim values in the defined dimension range
	for i := 2; i <= int(header.Dim[0]); i++ {
		if header.Dim[i] <= 0 {
			header.Dim[i] = 1
		}
	}

	// fix any remaining bad dim[] values (only values 0 or 1 seem rational, otherwise set to arbitrary 1)
	for i := header.Dim[0] + 1; i <= 7; i++ {
		if header.Dim[i] != 1 && header.Dim[i] != 0 {
			header.Dim[i] = 1
		}
	}

	//  loop backwards until we find a dim bigger than 1
	var ndim int
	for i := 7; i >= 2; i-- {
		if header.Dim[i] > 1 {
			ndim = i
			break
		}
	}
	fmt.Println(ndim)

	// set bad grid spacings to 1.0
	for i := 1; i <= int(header.Dim[0]); i++ {
		if header.Pixdim[i] == 0.0 || math.IsInf(float64(header.Pixdim[i]), 0) {
			header.Pixdim[i] = 1.0
		}
	}

	return nil
}

// is_onefile = is_nifti && NIFTI_ONEFILE(nhdr) ;
//
//  if( is_nifti ) nim->nifti_type = (is_onefile) ? NIFTI_FTYPE_NIFTI1_1
//                                                : NIFTI_FTYPE_NIFTI1_2 ;
//  else           nim->nifti_type = NIFTI_FTYPE_ANALYZE ;
//
//  ii = nifti_short_order() ;
//  if( doswap )   nim->byteorder = REVERSE_ORDER(ii) ;
//  else           nim->byteorder = ii ;
//
//
//  /**- set dimensions of data array */
//
//  nim->ndim = nim->dim[0] = nhdr.dim[0];
//  nim->nx   = nim->dim[1] = nhdr.dim[1];
//  nim->ny   = nim->dim[2] = nhdr.dim[2];
//  nim->nz   = nim->dim[3] = nhdr.dim[3];
//  nim->nt   = nim->dim[4] = nhdr.dim[4];
//  nim->nu   = nim->dim[5] = nhdr.dim[5];
//  nim->nv   = nim->dim[6] = nhdr.dim[6];
//  nim->nw   = nim->dim[7] = nhdr.dim[7];
//
//  for( ii=1, nim->nvox=1; ii <= nhdr.dim[0]; ii++ )
//     nim->nvox *= nhdr.dim[ii];
//
//  /**- set the type of data in voxels and how many bytes per voxel */
//
//  nim->datatype = nhdr.datatype ;
//
//  nifti_datatype_sizes( nim->datatype , &(nim->nbyper) , &(nim->swapsize) ) ;
//  if( nim->nbyper == 0 ){ free(nim); ERREX("bad datatype"); }
//
//  /**- set the grid spacings */
//
//  nim->dx = nim->pixdim[1] = nhdr.pixdim[1] ;
//  nim->dy = nim->pixdim[2] = nhdr.pixdim[2] ;
//  nim->dz = nim->pixdim[3] = nhdr.pixdim[3] ;
//  nim->dt = nim->pixdim[4] = nhdr.pixdim[4] ;
//  nim->du = nim->pixdim[5] = nhdr.pixdim[5] ;
//  nim->dv = nim->pixdim[6] = nhdr.pixdim[6] ;
//  nim->dw = nim->pixdim[7] = nhdr.pixdim[7] ;
//
//  /**- compute qto_xyz transformation from pixel indexes (i,j,k) to (x,y,z) */
//
//  if( !is_nifti || nhdr.qform_code <= 0 ){
//    /**- if not nifti or qform_code <= 0, use grid spacing for qto_xyz */
//
//    nim->qto_xyz.m[0][0] = nim->dx ;  /* grid spacings */
//    nim->qto_xyz.m[1][1] = nim->dy ;  /* along diagonal */
//    nim->qto_xyz.m[2][2] = nim->dz ;
//
//    /* off diagonal is zero */
//
//    nim->qto_xyz.m[0][1]=nim->qto_xyz.m[0][2]=nim->qto_xyz.m[0][3] = 0.0f;
//    nim->qto_xyz.m[1][0]=nim->qto_xyz.m[1][2]=nim->qto_xyz.m[1][3] = 0.0f;
//    nim->qto_xyz.m[2][0]=nim->qto_xyz.m[2][1]=nim->qto_xyz.m[2][3] = 0.0f;
//
//    /* last row is always [ 0 0 0 1 ] */
//
//    nim->qto_xyz.m[3][0]=nim->qto_xyz.m[3][1]=nim->qto_xyz.m[3][2] = 0.0f;
//    nim->qto_xyz.m[3][3]= 1.0f ;
//
//    nim->qform_code = NIFTI_XFORM_UNKNOWN ;
//
//    if( g_opts.debug > 1 ) fprintf(stderr,"-d no qform provided\n");
//  } else {
//    /**- else NIFTI: use the quaternion-specified transformation */
//
//    nim->quatern_b = FIXED_FLOAT( nhdr.quatern_b ) ;
//    nim->quatern_c = FIXED_FLOAT( nhdr.quatern_c ) ;
//    nim->quatern_d = FIXED_FLOAT( nhdr.quatern_d ) ;
//
//    nim->qoffset_x = FIXED_FLOAT(nhdr.qoffset_x) ;
//    nim->qoffset_y = FIXED_FLOAT(nhdr.qoffset_y) ;
//    nim->qoffset_z = FIXED_FLOAT(nhdr.qoffset_z) ;
//
//    nim->qfac = (nhdr.pixdim[0] < 0.0) ? -1.0f : 1.0f ;  /* left-handedness? */
//
//    nim->qto_xyz = nifti_quatern_to_mat44(
//                      nim->quatern_b, nim->quatern_c, nim->quatern_d,
//                      nim->qoffset_x, nim->qoffset_y, nim->qoffset_z,
//                      nim->dx       , nim->dy       , nim->dz       ,
//                      nim->qfac                                      ) ;
//
//    nim->qform_code = nhdr.qform_code ;
//
//    if( g_opts.debug > 1 )
//       nifti_disp_matrix_orient("-d qform orientations:\n", nim->qto_xyz);
//  }
//
//  /**- load inverse transformation (x,y,z) -> (i,j,k) */
//
//  nim->qto_ijk = nifti_mat44_inverse( nim->qto_xyz ) ;
//
//  /**- load sto_xyz affine transformation, if present */
//
//  if( !is_nifti || nhdr.sform_code <= 0 ){
//    /**- if not nifti or sform_code <= 0, then no sto transformation */
//
//    nim->sform_code = NIFTI_XFORM_UNKNOWN ;
//
//    if( g_opts.debug > 1 ) fprintf(stderr,"-d no sform provided\n");
//
//  } else {
//    /**- else set the sto transformation from srow_*[] */
//
//    nim->sto_xyz.m[0][0] = nhdr.srow_x[0] ;
//    nim->sto_xyz.m[0][1] = nhdr.srow_x[1] ;
//    nim->sto_xyz.m[0][2] = nhdr.srow_x[2] ;
//    nim->sto_xyz.m[0][3] = nhdr.srow_x[3] ;
//
//    nim->sto_xyz.m[1][0] = nhdr.srow_y[0] ;
//    nim->sto_xyz.m[1][1] = nhdr.srow_y[1] ;
//    nim->sto_xyz.m[1][2] = nhdr.srow_y[2] ;
//    nim->sto_xyz.m[1][3] = nhdr.srow_y[3] ;
//
//    nim->sto_xyz.m[2][0] = nhdr.srow_z[0] ;
//    nim->sto_xyz.m[2][1] = nhdr.srow_z[1] ;
//    nim->sto_xyz.m[2][2] = nhdr.srow_z[2] ;
//    nim->sto_xyz.m[2][3] = nhdr.srow_z[3] ;
//
//    /* last row is always [ 0 0 0 1 ] */
//
//    nim->sto_xyz.m[3][0]=nim->sto_xyz.m[3][1]=nim->sto_xyz.m[3][2] = 0.0f;
//    nim->sto_xyz.m[3][3]= 1.0f ;
//
//    nim->sto_ijk = nifti_mat44_inverse( nim->sto_xyz ) ;
//
//    nim->sform_code = nhdr.sform_code ;
//
//    if( g_opts.debug > 1 )
//       nifti_disp_matrix_orient("-d sform orientations:\n", nim->sto_xyz);
//  }
//
//  /**- set miscellaneous NIFTI stuff */
//
//  if( is_nifti ){
//    nim->scl_slope   = FIXED_FLOAT( nhdr.scl_slope ) ;
//    nim->scl_inter   = FIXED_FLOAT( nhdr.scl_inter ) ;
//
//    nim->intent_code = nhdr.intent_code ;
//
//    nim->intent_p1 = FIXED_FLOAT( nhdr.intent_p1 ) ;
//    nim->intent_p2 = FIXED_FLOAT( nhdr.intent_p2 ) ;
//    nim->intent_p3 = FIXED_FLOAT( nhdr.intent_p3 ) ;
//
//    nim->toffset   = FIXED_FLOAT( nhdr.toffset ) ;
//
//    memcpy(nim->intent_name,nhdr.intent_name,15); nim->intent_name[15] = '\0';
//
//    nim->xyz_units  = XYZT_TO_SPACE(nhdr.xyzt_units) ;
//    nim->time_units = XYZT_TO_TIME (nhdr.xyzt_units) ;
//
//    nim->freq_dim  = DIM_INFO_TO_FREQ_DIM ( nhdr.dim_info ) ;
//    nim->phase_dim = DIM_INFO_TO_PHASE_DIM( nhdr.dim_info ) ;
//    nim->slice_dim = DIM_INFO_TO_SLICE_DIM( nhdr.dim_info ) ;
//
//    nim->slice_code     = nhdr.slice_code  ;
//    nim->slice_start    = nhdr.slice_start ;
//    nim->slice_end      = nhdr.slice_end   ;
//    nim->slice_duration = FIXED_FLOAT(nhdr.slice_duration) ;
//  }
//
//  /**- set Miscellaneous ANALYZE stuff */
//
//  nim->cal_min = FIXED_FLOAT(nhdr.cal_min) ;
//  nim->cal_max = FIXED_FLOAT(nhdr.cal_max) ;
//
//  memcpy(nim->descrip ,nhdr.descrip ,79) ; nim->descrip [79] = '\0' ;
//  memcpy(nim->aux_file,nhdr.aux_file,23) ; nim->aux_file[23] = '\0' ;
//
//   /**- set ioff from vox_offset (but at least sizeof(header)) */
//
//   is_onefile = is_nifti && NIFTI_ONEFILE(nhdr) ;
//
//   if( is_onefile ){
//     ioff = (int)nhdr.vox_offset ;
//     if( ioff < (int) sizeof(nhdr) ) ioff = (int) sizeof(nhdr) ;
//   } else {
//     ioff = (int)nhdr.vox_offset ;
//   }
//   nim->iname_offset = ioff ;
//
//
//   /**- deal with file names if set */
//   if (fname!=NULL) {
//       nifti_set_filenames(nim,fname,0,0);
//       if (nim->iname==NULL)  { ERREX("bad filename"); }
//   } else {
//     nim->fname = NULL;
//     nim->iname = NULL;
//   }
//
//   /* clear extension fields */
//   nim->num_ext = 0;
//   nim->ext_list = NULL;
//
//   return nim;
