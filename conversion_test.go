package go2com

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser_Export(t *testing.T) {
	assert := assert.New(t)
	fPath := "./test_data/026.dcm"
	parser, err := NewDCMFileParser(fPath, WithSkipPixelData(true), WithSkipDataset(false))
	assert.NoError(err)
	err = parser.Parse()
	assert.NoError(err)

	mapTagVal := parser.Export(false)
	fmt.Println(mapTagVal)
}
