package tag

import "fmt"

const (
	TYPE_CURRENT = "current"
	TYPE_RETIRED = "retired"
)

type DicomTag struct {
	Group   uint16
	Element uint16
}

type TagInfo struct {
	Tag    DicomTag
	VR     string
	Name   string
	VM     string
	Status string
}

func IsPrivateTag(group uint16) bool {
	return group%2 == 1
}

func (tag DicomTag) String() string {
	return fmt.Sprintf("(%04x,%04x)", tag.Group, tag.Element)
}
