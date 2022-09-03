package tag

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/okieraised/go2com/internal/system"
)

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

// Compare checks if 2 tags are equal. If equals, return 0, else returns 1
func (tag DicomTag) Compare(otherTag DicomTag) int {
	if tag.Group == otherTag.Group && tag.Element == otherTag.Element {
		return 0
	}
	return 1
}

// IsPrivateTag return true if the tag is private
func IsPrivateTag(group uint16) bool {
	return group%2 == 1
}

// IsPublicTag return true if the tag is public
func IsPublicTag(group uint16) bool {
	return group%2 == 0
}

// String returns the tag in (gggg, eeee) format
func (tag DicomTag) String() string {
	return fmt.Sprintf("(%04x,%04x)", tag.Group, tag.Element)
}

// Find finds information about the given tag. If the tag is not
// part of the dictionary, raise error
func Find(tag DicomTag) (TagInfo, error) {
	initTag()
	entry, ok := TagDict[tag]
	if !ok {
		if tag.Group%2 == 0 && tag.Element == 0x0000 {
			entry = TagInfo{tag, "UL", "GenericGroupLength", "1", ""}
		} else {
			return TagInfo{}, fmt.Errorf("could not find tag (0x%x, 0x%x)", tag.Group, tag.Element)
		}
	}
	return entry, nil
}

// FindByName searchs for the tag by name
func FindByName(name string) (TagInfo, error) {
	initTag()
	for _, tag := range TagDict {
		if tag.Name == name {
			return tag, nil
		}
	}
	return TagInfo{}, fmt.Errorf("could not find tag %s", name)
}

// ConvertTagtoHex converts the 4 byte of dicom tag to DicomTag struct (gggg, eeee)
func ConvertTagtoHex(bTag []byte) (DicomTag, error) {
	if len(bTag) != 4 {
		return DicomTag{}, fmt.Errorf("invalid tag length of %d bytes", len(bTag))
	}

	b := make([]byte, 4)
	var group uint16
	var element uint16
	switch system.NativeEndian {
	case binary.BigEndian:
		// TODO: case when machine is big endian
	case binary.LittleEndian:
		bTagUint16 := binary.LittleEndian.Uint16(bTag[0:4])
		binary.BigEndian.PutUint16(b, bTagUint16)
		group64, err := strconv.ParseUint(hex.EncodeToString(b[0:2]), 16, 16)
		if err != nil {
			return DicomTag{}, err
		}
		group = uint16(group64)
		element64, err := strconv.ParseUint(hex.EncodeToString(b[2:4]), 16, 16)
		if err != nil {
			return DicomTag{}, err
		}
		element = uint16(element64)

	}
	return DicomTag{
		Group:   group,
		Element: element,
	}, nil
}
