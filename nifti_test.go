package go2com

import (
	"github.com/okieraised/go2com/pkg/nifti/nii_reader"
	"github.com/stretchr/testify/assert"
	_ "image/jpeg"
	"testing"
)

func TestNii1(t *testing.T) {
	assert := assert.New(t)

	filePath := "/home/tripg/Documents/nifti/Arnow^Corie^Shelvey^OM_segmented.nii"
	//filePath = "/home/tripg/Documents/nifti/RGB16_4D.nii.gz"

	niiReader, err := nii_reader.NewNiiReader(filePath)
	assert.NoError(err)

	//err = niiReader.ParseHeader()
	//assert.NoError(err)
	//err = niiReader.ParseData()
	//assert.NoError(err)
	err = niiReader.Parse()
	assert.NoError(err)

}
