package go2com

import (
	"fmt"
	"github.com/okieraised/go2com/pkg/nifti/nifti1"
	"github.com/stretchr/testify/assert"
	_ "image/jpeg"
	"testing"
)

func TestNii1(t *testing.T) {
	assert := assert.New(t)

	filePath := "/home/tripg/Documents/nifti/Arnow^Corie^Shelvey^OM_segmented.nii"
	filePath = "/home/tripg/Documents/nifti/RGB16_4D.nii.gz"
	filePath = "/home/tripg/Documents/nifti/someones_anatomy.nii.gz"
	//filePath = "/home/tripg/Documents/nifti/someones_epi.nii.gz"
	//filePath = "/home/tripg/Documents/nifti/RGB8_4D.nii.gz"
	//filePath = "/home/tripg/Documents/nifti/jaw.nii.gz"
	//filePath = "/home/tripg/Documents/nifti/Arnow^Corie^Shelvey^OM_segmented.nii"

	niiReader, err := nifti1.NewNii1Reader(filePath)
	assert.NoError(err)
	err = niiReader.Parse()
	assert.NoError(err)

	fmt.Println(niiReader.GetUnitsOfMeasurements())
	fmt.Println(niiReader.GetDatatype())
	fmt.Println(niiReader.GetNiiData().Header)
	fmt.Println(niiReader.GetNiiData().Header.Bitpix)
	fmt.Println("GetSFromCode()", niiReader.GetSFromCode())
	fmt.Println("GetQFromCode()", niiReader.GetQFromCode())

	//fmt.Println(niiReader.GetSliceCode())
	//mat := niiReader.QuaternToMatrix()
	//fmt.Println(mat.M)
	//
	//niiReader.MatrixToOrientation(mat)
	//
	//fmt.Println(niiReader.GetOrientation())
	//
	//fmt.Println(niiReader.GetUnitsOfMeasurements())
	//
	//fmt.Println(niiReader.GetNiiData().Header)

	fmt.Println(niiReader.GetSlice(13, 0))

	//fmt.Println(niiReader.GetTimeSeries(26, 30, 16))

	//for x := 0; x < 57; x++ {
	//	for y := 0; y < 67; y++ {
	//		for z := 0; z < 56; z++ {
	//			fmt.Println(niiReader.GetTimeSeries(int64(x), int64(y), int64(z)))
	//		}
	//	}
	//}

	shape := niiReader.GetImgShape()
	fmt.Println(shape)
}
