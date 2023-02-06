package dcm_io

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/okieraised/go2com/internal/system"
	"github.com/okieraised/go2com/pkg/dicom/tag"
	"github.com/okieraised/go2com/pkg/dicom/vr"
	"io"
	"reflect"
	"strconv"
	"strings"
)

const (
	VLUndefinedLength uint32 = 0xFFFFFFFF
)

const (
	PrivateTag = "PrivateTag"
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

// ReadElement reads the DICOM file tag by tag and returns the pointer to the parsed Element
func ReadElement(r DcmReader, isImplicit bool, binOrder binary.ByteOrder) (*Element, error) {
	tagVal, dcmTagInfo, err := readTag(r)
	if err != nil {
		return nil, err
	}

	if *tagVal == tag.ItemDelimitationItem || *tagVal == tag.Item {
		_ = r.skip(4)
		return nil, nil
	}

	if *tagVal == tag.PixelData && r.SkipPixelData() {
		_, err = r.discard(int(r.GetFileSize()))
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
func readTag(r DcmReader) (*tag.DicomTag, *tag.TagInfo, error) {
	group, err := r.readUInt16()
	if err != nil {
		return nil, nil, err
	}
	element, err := r.readUInt16()
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
			Name:   PrivateTag,
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

// readVR returns the value representation of the tag
func readVR(r DcmReader, isImplicit bool, t tag.DicomTag) (string, error) {
	if isImplicit {
		record, err := tag.Find(t)
		if err != nil {
			return vr.Unknown, nil
		}
		return record.VR, nil
	}
	return r.readString(2)
}

// readVL returns the value length of the dicom tag
func readVL(r DcmReader, isImplicit bool, t tag.DicomTag, valueRepresentation string) (uint32, error) {
	if isImplicit {
		return r.readUInt32()
	}

	switch valueRepresentation {
	// if the VR is equal to ‘OB’,’OW’,’OF’,’SQ’,’UI’ or ’UN’,
	// the VR is having an extra 2 bytes trailing to it. These 2 bytes trailing to VR are empty and are not decoded.
	// When VR is having these 2 extra empty bytes the VL will occupy 4 bytes rather than 2 bytes
	case vr.OtherByte, vr.OtherWord, vr.OtherFloat, vr.SequenceOfItems, vr.Unknown, vr.OtherByteOrOtherWord,
		strings.ToLower(vr.OtherByteOrOtherWord), vr.UnlimitedText, vr.UniversalResourceIdentifier,
		vr.UnlimitedCharacters:
		r.skip(2)
		valueLength, err := r.readUInt32()
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
		valueLength, err := r.readUInt16()
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

// readValue returns the value of the dicom tag
func readValue(r DcmReader, t tag.DicomTag, valueRepresentation string, valueLength uint32) (interface{}, error) {
	// Add this here for a special case when the tag is private and the value representation is UN (unknown) but the
	// value is of sequence of items. In this case, we will peek the next 4 bytes and check if it matches the item tag
	// If yes then handles like SQ
	if valueRepresentation == vr.Unknown && t.Group%2 != 0 {
		n, err := r.peek(4)
		if err != nil {
			return nil, err
		}
		if binary.BigEndian.Uint32(n) == 0xFFFEE000 || binary.BigEndian.Uint32(n) == 0xFEFF00E0 || binary.BigEndian.Uint32(n) == VLUndefinedLength {
			r.setTransferSyntax(r.ByteOrder(), true)
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

// switchStringToNumeric convert the decimal string to its appropriate value type
func switchStringToNumeric(in interface{}, valueRepresentation string) interface{} {
	switch valueRepresentation {
	case vr.IntegerString:
		switch reflect.ValueOf(in).Kind() {
		case reflect.Slice:
			ValStrArr, ok := (in).([]string)
			if !ok {
				return in
			}
			res := make([]int, 0, len(ValStrArr))
			for _, sub := range ValStrArr {
				intVar, err := strconv.Atoi(sub)
				if err != nil {
					return in
				}
				res = append(res, intVar)
			}
			return res
		case reflect.String:
			valStr, ok := (in).(string)
			if !ok {
				return in
			}
			intVal, err := strconv.Atoi(valStr)
			if err != nil {
				return in
			}
			return intVal
		}

	case vr.DecimalString, vr.OtherFloat, vr.OtherDouble:
		switch reflect.ValueOf(in).Kind() {
		case reflect.Slice:
			ValStrArr, ok := (in).([]string)
			if !ok {
				return in
			}
			res := make([]float64, 0, len(ValStrArr))
			for _, sub := range ValStrArr {
				flVar, err := strconv.ParseFloat(sub, 64)
				if err != nil {
					return in
				}
				res = append(res, flVar)
			}
			return res
		case reflect.String:
			valStr, ok := (in).(string)
			if !ok {
				return in
			}
			flVar, err := strconv.ParseFloat(valStr, 64)
			if err != nil {
				return in
			}
			return flVar
		}
	default:
	}
	return in
}

// readStringType reads the value as string and strips any zero padding
func readStringType(r DcmReader, t tag.DicomTag, valueRepresentation string, valueLength uint32) (interface{}, error) {
	sep := "\\"
	str, err := r.readString(valueLength)
	if err != nil {
		return str, err
	}
	str = strings.Trim(str, " \000") // There is a space " \000", not "\000"
	if strings.Contains(str, sep) {
		strArr := strings.Split(str, sep)
		res := switchStringToNumeric(strArr, valueRepresentation)
		return res, nil
	}
	res := switchStringToNumeric(str, valueRepresentation)
	return res, nil
}

// readPixelDataType reads the raw pixel data
func readPixelDataType(r DcmReader, t tag.DicomTag, valueRepresentation string, valueLength uint32) (interface{}, error) {
	if valueLength%2 != 0 && valueLength != VLUndefinedLength {
		fmt.Printf("Odd value length encountered for tag: %v with length %d", t.String(), valueLength)
	}
	byteSize := r.GetFileSize()
	if valueLength != VLUndefinedLength {
		byteSize = int64(valueLength)
	}

	bArr := make([]byte, byteSize)
	n, err := io.ReadFull(r, bArr)
	sbArr := bArr[:n]
	if err != nil {
		if err == io.ErrUnexpectedEOF {
			return sbArr, nil
		}
		return nil, err
	}
	return bArr, nil
}

// readByteType reads the value as byte array
func readByteType(r DcmReader, t tag.DicomTag, valueRepresentation string, valueLength uint32) (interface{}, error) {
	switch valueRepresentation {
	case vr.OtherByte, vr.Unknown, vr.OtherByteOrOtherWord, strings.ToLower(vr.OtherByteOrOtherWord):
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
			word, err := r.readUInt16()
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
		_, err := r.discard(int(valueLength))
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

// readIntType reads the value as integer and returns either the value or a slice of value
func readIntType(r DcmReader, t tag.DicomTag, valueRepresentation string, valueLength uint32) (interface{}, error) {
	var subVal int
	retVal := make([]int, 0, valueLength/2)
	n, err := r.peek(int(valueLength))
	if err != nil {
		return nil, err
	}
	subReader := bytes.NewReader(n)
	subRd := NewDICOMReader(bufio.NewReader(subReader), WithSkipPixelData(r.SkipPixelData()))
	byteRead := 0
	for {
		if byteRead >= int(valueLength) {
			break
		}
		switch valueRepresentation {
		case vr.UnsignedShort, vr.SignedShortOrUnsignedShort, strings.ToLower(vr.SignedShortOrUnsignedShort):
			val, err := subRd.readUInt16()
			if err != nil {
				return nil, err
			}
			subVal = int(val)
			byteRead += 2
		case vr.AttributeTag:
			val, err := subRd.readUInt16()
			if err != nil {
				return nil, err
			}
			subVal = int(val)
			byteRead += 2
		case vr.UnsignedLong:
			val, err := subRd.readUInt32()
			if err != nil {
				return nil, err
			}
			subVal = int(val)
			byteRead += 4
		case vr.SignedLong:
			val, err := subRd.readInt32()
			if err != nil {
				return nil, err
			}
			subVal = int(val)
			byteRead += 4
		case vr.SignedShort:
			val, err := subRd.readInt16()
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
	_, _ = r.discard(int(valueLength))
	if len(retVal) == 1 {
		return retVal[0], nil
	}
	return retVal, nil
}

// readFloatType reads the value as float
func readFloatType(r DcmReader, t tag.DicomTag, valueRepresentation string, valueLength uint32) (interface{}, error) {
	var subVal float64
	retVal := make([]float64, 0, valueLength/2)
	n, err := r.peek(int(valueLength))
	if err != nil {
		return nil, err
	}
	subReader := bytes.NewReader(n)
	subRd := NewDICOMReader(bufio.NewReader(subReader), WithSkipPixelData(r.SkipPixelData()))
	byteRead := 0
	for {
		if byteRead >= int(valueLength) {
			break
		}
		switch valueRepresentation {
		case vr.FloatingPointSingle, vr.OtherFloat:
			val, err := subRd.readFloat32()
			if err != nil {
				return nil, err
			}
			subVal = float64(val)
			byteRead += 4
		case vr.FloatingPointDouble:
			val, err := subRd.readFloat64()
			if err != nil {
				return nil, err
			}
			subVal = val
			byteRead += 8
		}
		retVal = append(retVal, subVal)
	}
	_, _ = r.discard(int(valueLength))
	if len(retVal) == 1 {
		return retVal[0], nil
	}

	return retVal, nil
}

// readSequence reads the value as sequence of items
func readSequence(r DcmReader, t tag.DicomTag, valueRepresentation string, valueLength uint32) (interface{}, error) {
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
		if valueRepresentation == vr.Unknown {
			r.setTransferSyntax(r.ByteOrder(), r.isTrackingImplicit())
		}
	} else {
		n, err := r.peek(int(valueLength))
		if err != nil {
			if err == bufio.ErrBufferFull {
				bRaw, err := writeToBuf(r, int(valueLength))
				if err != nil {
					return nil, err
				}
				sequences, err = readDefinedLengthSequences(r, bRaw, valueRepresentation)
				if err != nil {
					return nil, err
				}
				return sequences, nil
			}
			return nil, err
		}
		sequences, err = readDefinedLengthSequences(r, n, valueRepresentation)
		if err != nil {
			return nil, err
		}
		_, _ = r.discard(int(valueLength))
	}
	//r.setTransferSyntax(r.ByteOrder(), r.isTrackingImplicit())
	return sequences, nil

}

func writeToBuf(r DcmReader, n int) ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0, n))

	for i := 0; i < n; i++ {
		word, err := r.readUInt8()
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

func readDefinedLengthSequences(r DcmReader, b []byte, valueRepresentation string) ([]*Element, error) {
	var sequences []*Element
	br := bytes.NewReader(b)
	subRd := NewDICOMReader(bufio.NewReaderSize(br, len(b)), WithSkipPixelData(r.SkipPixelData()))
	_ = subRd.skip(8)
	subRd.setTransferSyntax(r.ByteOrder(), r.IsImplicit())
	for {
		subElement, err := ReadElement(subRd, r.IsImplicit(), r.ByteOrder())
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
	if valueRepresentation == vr.Unknown {
		r.setTransferSyntax(r.ByteOrder(), r.isTrackingImplicit())
	}

	return sequences, nil
}
