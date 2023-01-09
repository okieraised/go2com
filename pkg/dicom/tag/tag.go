package tag

import (
	"encoding/binary"
	"fmt"
	"sync"
)

const (
	TypeCurrent = "current"
	TypeRetired = "retired"
	TagUnknown  = "unknown_tag"
)

type TagBrowser struct {
	Value       interface{} `json:"Value,omitempty" bson:"Value,omitempty"`
	VR          string      `json:"vr" bson:"vr"`
	BulkDataURI string      `json:"BulkDataURI,omitempty" bson:"BulkDataURI,omitempty"`
}

type DicomTag struct {
	Group   uint16
	Element uint16
}

type TagInfo struct {
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

// StringWithoutParentheses returns the tag in ggggeeee format
func (tag DicomTag) StringWithoutParentheses() string {
	return fmt.Sprintf("%04X%04X", tag.Group, tag.Element)
}

// Find finds information about the given tag. If the tag is not
// part of the dictionary, raise error
func Find(tag DicomTag) (TagInfo, error) {
	err := fmt.Errorf("could not find tag (0x%x, 0x%x)", tag.Group, tag.Element)
	entry, ok := TagDict[tag]
	if !ok {
		if tag.Group%2 == 0 {
			if tag.Element == 0x0000 {
				entry = TagInfo{"UL", "GenericGroupLength", "1", ""}
			} else {
				b := make([]byte, 2)
				binary.BigEndian.PutUint16(b, tag.Group)
				switch int(b[0]) {
				case 80:
					entry, kk := TagDict[DicomTag{
						Group:   0x5000,
						Element: tag.Element,
					}]
					if !kk {
						return TagInfo{}, err
					}
					return entry, nil
				case 96:
					entry, kk := TagDict[DicomTag{
						Group:   0x6000,
						Element: tag.Element,
					}]
					if !kk {
						return TagInfo{}, err
					}
					return entry, nil
				default:
					return TagInfo{}, err
				}
			}
		} else {
			return TagInfo{}, err
		}
	}
	return entry, nil
}

// FindByName searchs for the tag by name
func FindByName(name string) (TagInfo, error) {
	for _, tag := range TagDict {
		if tag.Name == name {
			return tag, nil
		}
	}
	return TagInfo{}, fmt.Errorf("could not find tag %s", name)
}

var lock = &sync.RWMutex{}

func InitTagDict() map[DicomTag]TagInfo {
	lock.Lock()
	defer lock.Unlock()
	if TagDict == nil {
		initTag()
	}
	return TagDict
}
