package go2com

import (
	"fmt"
	"github.com/okieraised/go2com/pkg/dicom/dcm_io"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser_Export(t *testing.T) {
	assert := assert.New(t)
	fPath := "./test_data/026.dcm"
	parser, err := dcm_io.NewDCMFileParser(fPath, dcm_io.WithSkipPixelData(true), dcm_io.WithSkipDataset(false))
	assert.NoError(err)
	err = parser.Parse()
	assert.NoError(err)

	mapTagVal := parser.Export(false)
	fmt.Println(mapTagVal)
}
