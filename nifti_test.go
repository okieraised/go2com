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

	niiReader, err := nifti1.NewNii1Reader(filePath)
	assert.NoError(err)
	err = niiReader.Parse()
	assert.NoError(err)

	fmt.Println(niiReader.GetDatatype())
	fmt.Println(niiReader.GetSliceCode())
	mat := niiReader.QuaternToMatrix()
	fmt.Println(mat.M)

	niiReader.MatrixToOrientation(mat)

	fmt.Println(niiReader.GetOrientation())

	fmt.Println(niiReader.GetUnitsOfMeasurements())

	fmt.Println(niiReader.GetNiiData().Data.Data)

	shape := niiReader.GetImgShape()
	fmt.Println(shape)

	fmt.Println(niiReader.GetSlice(28, 0))
}
