package go2com

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestParser_Export(t *testing.T) {
	assert := assert.New(t)
	InitTagDict()
	file, err := os.Open("./test_data/026.dcm")
	assert.NoError(err)

	defer file.Close()
	info, err := file.Stat()
	assert.NoError(err)
	fileSize := info.Size()

	parser, err := NewParser(file, fileSize, false, false)
	assert.NoError(err)
	err = parser.Parse()
	assert.NoError(err)

	mapTagVal := parser.Export(false)
	fmt.Println(mapTagVal)
}
