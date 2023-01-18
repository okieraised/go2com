package iod

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/okieraised/go2com/pkg/dicom/dcm_io"
	"github.com/okieraised/go2com/pkg/dicom/tag"
	"github.com/okieraised/go2com/pkg/dicom/uid"
	"io"
)

type PixelDataMacro map[tag.DicomTag]*dcm_io.Element

// GetPixelDataMacroAttributes retrieves the tags and values corresponding to the PixelData macro
//    +------------------------------------------------+
//    | Element                                        |
//    +-------------+---------------------------+------+
//    | Tag         | Keyword                   | Type |
//    +=============+===========================+======+
//    | (0028,0002) | SamplesPerPixel           | 1    |
//    +-------------+---------------------------+------+
//    | (0028,0004) | PhotometricInterpretation | 1    |
//    +-------------+---------------------------+------+
//    | (0028,0006) | PlanarConfiguration       | 1C   |
//    +-------------+---------------------------+------+
//    | (0028,0008) | NumberOfFrames            | 1C   |
//    +-------------+---------------------------+------+
//    | (0028,0010) | Rows                      | 1    |
//    +-------------+---------------------------+------+
//    | (0028,0011) | Columns                   | 1    |
//    +-------------+---------------------------+------+
//    | (0028,0100) | BitsAllocated             | 1    |
//    +-------------+---------------------------+------+
//    | (0028,0101) | BitsStored                | 1    |
//    +-------------+---------------------------+------+
//    | (0028,0103) | PixelRepresentation       | 1    |
//    +-------------+---------------------------+------+
//    | (7FE0,0008) | FloatPixelData            | 1C   |
//    +-------------+---------------------------+------+
//    | (7FE0,0009) | DoubleFloatPixelData      | 1C   |
//    +-------------+---------------------------+------+
//    | (7FE0,0010) | PixelData                 | 1C   |
//    +-------------+---------------------------+------+
func GetPixelDataMacroAttributes(ds, meta dcm_io.Dataset) PixelDataMacro {
	res := make(PixelDataMacro, 0)
	for _, elem := range meta.Elements {
		if elem.Tag == tag.TransferSyntaxUID {
			res[elem.Tag] = elem
		}
	}

	for _, elem := range ds.Elements {
		if elem.Tag == tag.SamplesPerPixel || elem.Tag == tag.PhotometricInterpretation ||
			elem.Tag == tag.Rows || elem.Tag == tag.Columns ||
			elem.Tag == tag.BitsAllocated || elem.Tag == tag.BitsStored ||
			elem.Tag == tag.HighBit || elem.Tag == tag.PixelRepresentation ||
			elem.Tag == tag.PlanarConfiguration || elem.Tag == tag.PixelAspectRatio ||
			elem.Tag == tag.SmallestImagePixelValue || elem.Tag == tag.LargestImagePixelValue ||
			elem.Tag == tag.RedPaletteColorLookupTableDescriptor || elem.Tag == tag.GreenPaletteColorLookupTableDescriptor ||
			elem.Tag == tag.BluePaletteColorLookupTableDescriptor || elem.Tag == tag.RedPaletteColorLookupTableData ||
			elem.Tag == tag.GreenPaletteColorLookupTableData || elem.Tag == tag.BluePaletteColorLookupTableData ||
			elem.Tag == tag.FloatPixelData || elem.Tag == tag.DoubleFloatPixelData ||
			elem.Tag == tag.PixelData || elem.Tag == tag.NumberOfFrames || elem.Tag == tag.ColorSpace {
			res[elem.Tag] = elem
		}
	}
	return res
}

func (px PixelDataMacro) GetExpectedPixelData() int {
	var length, rows, columns, samplesPerPixel, bitsAllocated int
	noOfFrames := 1

	v, ok := px[tag.Rows]
	if ok {
		rows = v.Value.RawValue.(int)
	}

	v, ok = px[tag.Columns]
	if ok {
		columns = v.Value.RawValue.(int)
	}

	v, ok = px[tag.SamplesPerPixel]
	if ok {
		samplesPerPixel = v.Value.RawValue.(int)
	}

	v, ok = px[tag.BitsAllocated]
	if ok {
		bitsAllocated = v.Value.RawValue.(int)
	}

	k, ok := px[tag.NumberOfFrames]
	if ok {
		val := k.Value.RawValue.(int)
		noOfFrames = val
	}

	length = rows * columns * samplesPerPixel

	length *= noOfFrames

	if bitsAllocated == 1 {
		length = length / 8
	} else {
		length *= bitsAllocated / 8
	}

	if px[tag.PhotometricInterpretation].Value.RawValue == "YBR_FULL_422" {
		length = length / 3 * 2
	}

	return length
}

func (px PixelDataMacro) ReadEncapsulatedPixelData() ([]byte, error) {
	pixelDataElem := px[tag.PixelData]
	rawPixel, ok := pixelDataElem.Value.RawValue.([]byte)
	if !ok {
		return nil, fmt.Errorf("cannot convert pixel data to byte array")
	}

	if uid.UncompressedSyntax[px[tag.TransferSyntaxUID].Value.RawValue.(string)] {
		return rawPixel, nil
	}

	bufRd := bufio.NewReaderSize(bytes.NewReader(rawPixel), int(pixelDataElem.ValueLength))
	pixReader := dcm_io.NewDICOMReader(bufRd, dcm_io.WithSkipPixelData(true))
	actualPixelData := make([]byte, 0, int(pixelDataElem.ValueLength))
	index := 0
	for {
		var tGroup, tElem uint16
		tGroup, err := pixReader.readUInt16()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		tElem, err = pixReader.readUInt16()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		tTag := tag.DicomTag{
			Group:   tGroup,
			Element: tElem,
		}

		if tTag == tag.SequenceDelimitationItem {
			break
		}

		tValueLength, err := pixReader.ReadUInt32()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		rawValue, err := pixReader.peek(int(tValueLength))
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		// The first item is the basic offset table, so we skip it
		if index > 0 {
			actualPixelData = append(actualPixelData, rawValue...)
		}
		_, err = pixReader.discard(int(tValueLength))
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		index++
	}

	return actualPixelData, nil
}

func (px PixelDataMacro) ValidatePixelData() bool {
	expected := px.GetExpectedPixelData()
	rawPixel, ok := px[tag.PixelData]
	if !ok {
		return false
	}

	actual, ok := rawPixel.Value.RawValue.([]byte)
	if !ok {
		return false
	}

	rawTransferSyntax, ok := px[tag.TransferSyntaxUID]
	if !ok {
		return false
	}
	transferSyntax, ok := rawTransferSyntax.Value.RawValue.(string)
	if !ok {
		return false
	}

	if uid.UncompressedSyntax[transferSyntax] {
		if expected != len(actual) {
			return false
		}
		return true
	}

	// Other syntax is compressed, so we cannot check if they are equals to the expected bytes
	return true

}
