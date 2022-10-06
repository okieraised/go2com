package tag

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFind(t *testing.T) {
	assert := assert.New(t)
	InitTagDict()

	dTag := DicomTag{
		Group:   0x6010,
		Element: 0x3000,
	}

	x, err := Find(dTag)
	assert.NoError(err)
	fmt.Println(x)

}
