package element

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/okieraised/go2com/internal/constants"
	"github.com/okieraised/go2com/internal/system"
	"github.com/okieraised/go2com/pkg/dicom/reader"
	"github.com/okieraised/go2com/pkg/dicom/tag"
	"github.com/okieraised/go2com/pkg/dicom/vr"
	"io"
	"strings"
)

const (
	GroupSeqItem      uint16 = 0xFFFE
	VLUndefinedLength uint32 = 0xFFFFFFFF
)

// Element defines the struct for each dicom tag element info. Ordered as below to decrease the memory footprint
type Element struct {
	Value                  Value
	TagName                string
	ValueRepresentationStr string
	ValueLength            uint32
	Tag                    tag.DicomTag
	ValueRepresentation    vr.VRKind
}

func ReadElement(r reader.DcmReader, isImplicit bool, binOrder binary.ByteOrder) (*Element, error) {
	tagVal, dcmTagInfo, err := readTag(r)
	if err != nil {
		return nil, err
	}

	if *tagVal == tag.ItemDelimitationItem || *tagVal == tag.Item {
		_ = r.Skip(4)
		return nil, nil
	}

	if *tagVal == tag.PixelData && r.SkipPixelData() {
		_, err = r.Discard(int(r.GetFileSize()))
		if err != nil {
			return nil, err
		}
		return nil, nil
	}

	dmcTagName := dcmTagInfo.Name
	dcmVR, err := readVR(r, isImplicit, *tagVal)
	if err != nil {
		return nil, err
	}

	dcmVL, err := readVL(r, isImplicit, *tagVal, dcmVR)
	if err != nil {
		return nil, err
	}

	value, err := readValue(r, *tagVal, dcmVR, dcmVL)
	if err != nil {
		return nil, err
	}
	if n, ok := value.([]byte); ok {
		dcmVL = uint32(len(n))
	}

	elem := Element{
		Tag:                    *tagVal,
		TagName:                dmcTagName,
		ValueRepresentationStr: dcmVR,
		ValueLength:            dcmVL,
		Value:                  Value{RawValue: value},
	}

	return &elem, nil
}

// readTag returns the tag information
func readTag(r reader.DcmReader) (*tag.DicomTag, *tag.TagInfo, error) {
	group, err := r.ReadUInt16()
	if err != nil {
		return nil, nil, err
	}
	element, err := r.ReadUInt16()
	if err != nil {
		return nil, nil, err
	}

	t := tag.DicomTag{
		Group:   group,
		Element: element,
	}

	// Check if tag is private. If yes, just return here
	// Otherwise, find info about the public tag
	if int(group)%2 != 0 {
		tagInfo := tag.TagInfo{
			VR:     "",
			Name:   constants.PrivateTag,
			VM:     "",
			Status: "",
		}
		return &t, &tagInfo, nil
	}

	tagInfo, err := tag.Find(t)
	if err != nil {
		return nil, nil, err
	}
	return &t, &tagInfo, nil
}

// readVR
func readVR(r reader.DcmReader, isImplicit bool, t tag.DicomTag) (string, error) {
	if isImplicit {
		record, err := tag.Find(t)
		if err != nil {
			return vr.Unknown, nil
		}
		return record.VR, nil
	}
	return r.ReadString(2)
}

// readVL
func readVL(r reader.DcmReader, isImplicit bool, t tag.DicomTag, valueRepresentation string) (uint32, error) {
	if isImplicit {
		return r.ReadUInt32()
	}

	switch valueRepresentation {
	// if the VR is equal to ‘OB’,’OW’,’OF’,’SQ’,’UI’ or ’UN’,
	// the VR is having an extra 2 bytes trailing to it. These 2 bytes trailing to VR are empty and are not decoded.
	// When VR is having these 2 extra empty bytes the VL will occupy 4 bytes rather than 2 bytes
	case vr.OtherByte, vr.OtherWord, vr.OtherFloat, vr.SequenceOfItems, vr.Unknown, vr.OtherByteOrOtherWord, strings.ToLower(vr.OtherByteOrOtherWord):
		r.Skip(2)
		valueLength, err := r.ReadUInt32()
		if err != nil {
			return 0, err
		}
		if valueLength == VLUndefinedLength &&
			(valueRepresentation == vr.UnlimitedCharacters ||
				valueRepresentation == vr.UniversalResourceIdentifier ||
				valueRepresentation == vr.UnlimitedText) {
			return 0, errors.New("UC, UR and UT must have defined length")
		}
		return valueLength, nil
	default:
		valueLength, err := r.ReadUInt16()
		if err != nil {
			return 0, err
		}
		vl := uint32(valueLength)
		if vl == 0xffff {
			vl = VLUndefinedLength
		}
		return vl, nil
	}
}

// readValue
func readValue(r reader.DcmReader, t tag.DicomTag, valueRepresentation string, valueLength uint32) (interface{}, error) {
	// Add this here for a special case when the tag is private and the value representation is UN (unknown) but the
	// value is of sequence of items. In this case, we will peek the next 4 bytes and check if it matches the item tag
	// If yes then handles like SQ
	if valueRepresentation == vr.Unknown && t.Group%2 != 0 {
		n, err := r.Peek(4)
		if err != nil {
			return nil, err
		}
		if binary.BigEndian.Uint32(n) == 0xFFFEE000 || binary.BigEndian.Uint32(n) == 0xFEFF00E0 || binary.BigEndian.Uint32(n) == VLUndefinedLength {
			return readSequence(r, t, valueRepresentation, valueLength)
		}
	}
	vrKind := vr.GetVR(t, valueRepresentation)
	switch vrKind {
	case vr.VRString, vr.VRDate:
		return readStringType(r, t, valueRepresentation, valueLength)
	case vr.VRInt16, vr.VRInt32, vr.VRUInt16, vr.VRUInt32, vr.VRTagList:
		return readIntType(r, t, valueRepresentation, valueLength)
	case vr.VRFloat32, vr.VRFloat64:
		return readFloatType(r, t, valueRepresentation, valueLength)
	case vr.VRBytes:
		return readByteType(r, t, valueRepresentation, valueLength)
	case vr.VRPixelData:
		return readPixelDataType(r, t, valueRepresentation, valueLength)
	case vr.VRSequence:
		return readSequence(r, t, valueRepresentation, valueLength)
	default:
		return readStringType(r, t, valueRepresentation, valueLength)
	}
}

// readStringType
func readStringType(r reader.DcmReader, t tag.DicomTag, valueRepresentation string, valueLength uint32) (interface{}, error) {
	sep := "\\"
	str, err := r.ReadString(valueLength)
	if err != nil {
		return str, err
	}
	str = strings.Trim(str, " \000") // There is a space " \000", not "\000"
	if strings.Contains(str, sep) {
		res := strings.Split(str, sep)
		return res, nil
	}
	return str, nil
}

func readPixelDataType(r reader.DcmReader, t tag.DicomTag, valueRepresentation string, valueLength uint32) (interface{}, error) {
	// FE FF DD E0
	if valueLength%2 != 0 {
		fmt.Printf("Odd value length encountered for tag: %v with length %d", t.String(), valueLength)
	}
	res := make([]byte, 0)

	for {
		bRead, err := r.ReadUInt8()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		res = append(res, bRead)
	}
	return res, nil
}

// readByteType reads the value as byte array or word array
func readByteType(r reader.DcmReader, t tag.DicomTag, valueRepresentation string, valueLength uint32) (interface{}, error) {
	//fmt.Println(valueLength)
	if valueLength == VLUndefinedLength {
		fmt.Println(true)
	}
	switch valueRepresentation {
	case vr.OtherByte, vr.Unknown:
		bArr := make([]byte, valueLength)
		n, err := io.ReadFull(r, bArr)
		sbArr := bArr[:n]
		if err != nil {
			if err == io.ErrUnexpectedEOF {
				return sbArr, nil
			}
			return nil, err
		}
		return bArr, nil
	case vr.OtherWord:
		if valueLength%2 != 0 {
			fmt.Printf("Odd value length encountered for tag: %v with length %d", t.String(), valueLength)
		}
		buf := bytes.NewBuffer(make([]byte, 0, valueLength))
		numWords := int(valueLength / 2)
		for i := 0; i < numWords; i++ {
			word, err := r.ReadUInt16()
			if err != nil {
				// Handle a case when the actual pixel data is less than the value length. Just return what we can
				// read here
				if err == io.EOF {
					err = binary.Write(buf, system.NativeEndian, word)
					if err != nil {
						return nil, err
					}
					r = nil
					return buf.Bytes(), nil
				}
				return nil, err
			}
			err = binary.Write(buf, system.NativeEndian, word)
			if err != nil {
				return nil, err
			}

		}
		return buf.Bytes(), nil
	default:
	}
	return nil, nil
}

// readIntType reads the value as integer
func readIntType(r reader.DcmReader, t tag.DicomTag, valueRepresentation string, valueLength uint32) (interface{}, error) {
	var subVal int
	retVal := make([]int, 0, valueLength/2)
	n, err := r.Peek(int(valueLength))
	if err != nil {
		return nil, err
	}
	subReader := bytes.NewReader(n)
	subRd := reader.NewDcmReader(bufio.NewReader(subReader), r.SkipPixelData())
	byteRead := 0
	for {
		if byteRead >= int(valueLength) {
			break
		}
		switch valueRepresentation {
		case vr.UnsignedShort, vr.SignedShortOrUnsignedShort, strings.ToLower(vr.SignedShortOrUnsignedShort):
			val, err := subRd.ReadUInt16()
			if err != nil {
				return nil, err
			}
			subVal = int(val)
			byteRead += 2
		case vr.AttributeTag:
			val, err := subRd.ReadUInt16()
			if err != nil {
				return nil, err
			}
			subVal = int(val)
			byteRead += 2
		case vr.UnsignedLong:
			val, err := subRd.ReadUInt32()
			if err != nil {
				return nil, err
			}
			subVal = int(val)
			byteRead += 4
		case vr.SignedLong:
			val, err := subRd.ReadInt32()
			if err != nil {
				return nil, err
			}
			subVal = int(val)
			byteRead += 4
		case vr.SignedShort:
			val, err := subRd.ReadInt16()
			if err != nil {
				return nil, err
			}
			subVal = int(val)
			byteRead += 2
			//default:
			//	return nil, fmt.Errorf("cannot parse value as integer for tag %v", t)
		}
		retVal = append(retVal, subVal)
	}
	_, _ = r.Discard(int(valueLength))
	if len(retVal) == 1 {
		return retVal[0], nil
	}
	return retVal, nil
}

// readFloatType reads the value as float
func readFloatType(r reader.DcmReader, t tag.DicomTag, valueRepresentation string, valueLength uint32) (interface{}, error) {
	var subVal float64
	retVal := make([]float64, 0, valueLength/2)
	n, err := r.Peek(int(valueLength))
	if err != nil {
		return nil, err
	}
	subReader := bytes.NewReader(n)
	subRd := reader.NewDcmReader(bufio.NewReader(subReader), r.SkipPixelData())
	byteRead := 0
	for {
		if byteRead >= int(valueLength) {
			break
		}
		switch valueRepresentation {
		case vr.FloatingPointSingle, vr.OtherFloat:
			val, err := subRd.ReadFloat32()
			if err != nil {
				return nil, err
			}
			subVal = float64(val)
			byteRead += 4
		case vr.FloatingPointDouble:
			val, err := subRd.ReadFloat64()
			if err != nil {
				return nil, err
			}
			subVal = val
			byteRead += 8
		}
		retVal = append(retVal, subVal)
	}
	_, _ = r.Discard(int(valueLength))
	if len(retVal) == 1 {
		return retVal[0], nil
	}

	return retVal, nil
}

// readSequence reads the value as sequence of items
func readSequence(r reader.DcmReader, t tag.DicomTag, valueRepresentation string, valueLength uint32) (interface{}, error) {
	var sequences []*Element
	// Reference: https://dicom.nema.org/dicom/2013/output/chtml/part05/sect_7.5.html
	if valueLength == VLUndefinedLength {
		for {
			subElement, err := ReadElement(r, r.IsImplicit(), r.ByteOrder())
			if err != nil {
				return nil, err
			}

			if subElement == nil {
				continue
			}

			if subElement.Tag == tag.SequenceDelimitationItem {
				break
			}
			sequences = append(sequences, subElement)
		}
	} else {
		n, err := r.Peek(int(valueLength))
		if err != nil {
			if err == bufio.ErrBufferFull {
				bRaw, err := writeToBuf(r, int(valueLength))
				if err != nil {
					return nil, err
				}
				sequences, err = readDefinedLengthSequences(r, bRaw)
				if err != nil {
					return nil, err
				}
				return sequences, nil
			}
			return nil, err
		}
		sequences, err = readDefinedLengthSequences(r, n)
		if err != nil {
			return nil, err
		}
		_, _ = r.Discard(int(valueLength))
	}

	return sequences, nil

}

func writeToBuf(r reader.DcmReader, n int) ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0, n))

	for i := 0; i < n; i++ {
		word, err := r.ReadUInt8()
		if err != nil {
			return nil, err
		}
		err = binary.Write(buf, system.NativeEndian, word)
		if err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

func readDefinedLengthSequences(r reader.DcmReader, b []byte) ([]*Element, error) {
	var sequences []*Element
	br := bytes.NewReader(b)
	subRd := reader.NewDcmReader(bufio.NewReaderSize(br, len(b)), r.SkipPixelData())
	_ = subRd.Skip(8)
	subRd.SetTransferSyntax(r.ByteOrder(), r.IsImplicit())
	for {
		subElement, err := ReadElement(subRd, subRd.IsImplicit(), subRd.ByteOrder())
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
		if subElement == nil {
			continue
		}
		sequences = append(sequences, subElement)
	}
	return sequences, nil
}
