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
	//filePath = "/home/tripg/Documents/nifti/RGB16_4D.nii.gz"

	niiReader, err := nifti1.NewNii1Reader(filePath)
	assert.NoError(err)

	//err = niiReader.ParseHeader()
	//assert.NoError(err)
	//err = niiReader.ParseData()
	//assert.NoError(err)
	err = niiReader.Parse()
	assert.NoError(err)

	fmt.Println(niiReader.GetHeader().Datatype)

	//x, err := niiReader.GetTimeSeries(111, 256, 12)
	//assert.NoError(err)
	//fmt.Println(x)
	//
	//shape := niiReader.GetImgShape()
	//fmt.Println(shape)
	//
	//slices, err := niiReader.GetSlice(1, 0)
	//assert.NoError(err)
	//fmt.Println(slices)
}
