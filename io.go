package go2com

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"github.com/okieraised/go2com/internal/utils"
	"github.com/okieraised/go2com/pkg/dicom/tag"
	"github.com/okieraised/go2com/pkg/dicom/vr"
	"github.com/okieraised/go2com/pkg/plugins/orthanc"
	"io"
	"strings"
)

const (
	MagicString = "DICM"
	PrivateTag  = "PrivateTag"
)

type dcmReader struct {
	reader               *bufio.Reader
	binaryOrder          binary.ByteOrder
	dataset              Dataset
	metadata             Dataset
	allowNonCompliantDcm bool
	isImplicit           bool
	keepTrackImplicit    bool
	skipPixelData        bool
	skipDataset          bool
	fileSize             int64
}

// NewDICOMReader returns a new reader
func NewDICOMReader(reader *bufio.Reader, options ...func(*dcmReader)) *dcmReader {
	parser := &dcmReader{
		reader:        reader,
		binaryOrder:   binary.LittleEndian,
		isImplicit:    false,
		skipPixelData: false,
		skipDataset:   false,
	}
	for _, opt := range options {
		opt(parser)
	}
	return parser
}

// WithAllowNonCompliantDcm provides option to keep trying to parse the file even if it's not DICOM compliant
// e.g.: Missing header, missing FileMetaInformationGroupLength,...
func WithAllowNonCompliantDcm(allowNonCompliantDcm bool) func(*dcmReader) {
	return func(s *dcmReader) {
		s.allowNonCompliantDcm = allowNonCompliantDcm
	}
}

// WithSkipPixelData provides option to skip reading pixel data (7FE0,0010).
// If true, pixel data is skipped. If false, pixel data will be read
func WithSkipPixelData(skipPixelData bool) func(*dcmReader) {
	return func(s *dcmReader) {
		s.skipPixelData = skipPixelData
	}
}

// WithSkipDataset provides option to read only the metadata header.
// If true, only the meta header is read, else, the dataset will be read
func WithSkipDataset(skipDataset bool) func(*dcmReader) {
	return func(s *dcmReader) {
		s.skipDataset = skipDataset
	}
}

// WithSetFileSize provides option to set the file size to the reader
func WithSetFileSize(fileSize int64) func(*dcmReader) {
	return func(s *dcmReader) {
		s.fileSize = fileSize
	}
}

// GetMetadata returns the file meta header
func (r *dcmReader) GetMetadata() Dataset {
	return r.metadata
}

// GetDataset returns the dataset
func (r *dcmReader) GetDataset() Dataset {
	return r.dataset
}

func (r *dcmReader) RetrieveFileUID() (*DicomUID, error) {
	return r.dataset.RetrieveFileUID()
}

// Parse reads the DICOM file and parses it into array of elements
func (r *dcmReader) Parse() error {
	err := r.parse()
	if err != nil {
		return err
	}
	return nil
}

func (r *dcmReader) SkipPixelData() bool {
	return r.skipPixelData
}

func (r *dcmReader) ByteOrder() binary.ByteOrder {
	return r.binaryOrder
}

// IsValidDICOM checks if the dicom file follows the standard by having 128 bytes preamble followed by the magic string 'DICM'
func (r *dcmReader) IsValidDICOM() error {
	preamble, err := r.peek(132)
	if err != nil {
		return fmt.Errorf("cannot read the first 132 bytes: %v", err)
	}
	if string(preamble[128:]) != MagicString {
		return fmt.Errorf("file is not in valid dicom format")
	}
	return nil
}

// GetElementByTagString returns the element value of the input tag
// Tag should be in (gggg,eeee) or ggggeeee format
func (r *dcmReader) GetElementByTagString(tagStr string) (interface{}, error) {
	tagStr = utils.FormatTag(tagStr)

	if strings.HasPrefix(tagStr, "0002") {
		for _, elem := range r.metadata.Elements {
			if tagStr == elem.Tag.StringWithoutParentheses() {
				return elem.Value, nil
			}
		}
		return nil, fmt.Errorf("cannot find tag %s", tagStr)
	} else {
		for _, elem := range r.dataset.Elements {
			if tagStr == elem.Tag.StringWithoutParentheses() {
				return elem.Value, nil
			}
		}
		return nil, fmt.Errorf("cannot find tag %s", tagStr)
	}
}

// ExportDatasetTags returns the mapped tag/(vr,value) dictionary
func (r *dcmReader) ExportDatasetTags(exportMeta bool) MappedTag {
	res := make(MappedTag, len(r.metadata.Elements)+len(r.dataset.Elements))
	if exportMeta {
		mt := r.metadata
		for _, elem := range mt.Elements {
			res.mapElement(elem)
		}
	}

	ds := r.dataset
	//colorImage := false
	for _, elem := range ds.Elements {
		//if elem.Tag == tag.RedPaletteColorLookupTableData || elem.Tag == tag.BluePaletteColorLookupTableData || elem.Tag == tag.GreenPaletteColorLookupTableData {
		//	colorImage = true
		//}
		vrStr := elem.ValueRepresentationStr
		if vrStr == "OB" || vrStr == "OW" || vrStr == "UN" || strings.ToLower(vrStr) == "ox" {
			continue
		}
		res.mapElement(elem)
	}

	//if colorImage {
	//	bulkURIMap := createOrthancURI(ds)
	//	for k, v := range bulkURIMap {
	//		res[k] = v
	//	}
	//}

	return res
}

func (r *dcmReader) ExportSeriesTags() MappedTag {
	res := make(MappedTag, len(r.dataset.Elements))
	var value interface{}
	for _, elem := range r.dataset.Elements {
		if _, ok := orthanc.OrthancGetSeriesOfStudyTags[elem.Tag]; ok {
			tagStr := elem.Tag.StringWithoutParentheses()
			if elem.ValueRepresentationStr == vr.PersonName {
				value = utils.AppendToSlice(map[string]interface{}{
					"Alphabetic": elem.Value.RawValue,
				})
			} else {
				value = utils.AppendToSlice(elem.Value.RawValue)
			}
			res[tagStr] = tag.TagBrowser{
				VR:    elem.ValueRepresentationStr,
				Value: value,
			}

		}
	}

	prefix := "http://127.0.0.1:8042"
	firstSeriesURI := fmt.Sprintf("%s/dicom-web/studies/%s/series/%s", prefix,
		res[tag.StudyInstanceUID.StringWithoutParentheses()].Value.([]interface{})[0].(string),
		res[tag.SeriesInstanceUID.StringWithoutParentheses()].Value.([]interface{})[0].(string),
	)

	res[tag.RetrieveURL.StringWithoutParentheses()] = tag.TagBrowser{
		Value: []interface{}{firstSeriesURI},
		VR:    "UR",
	}
	return res
}

func (r *dcmReader) Read(p []byte) (int, error) {
	return r.reader.Read(p)
}

func (r *dcmReader) IsImplicit() bool {
	return r.isImplicit
}

func (r *dcmReader) SetTransferSyntax(binaryOrder binary.ByteOrder, isImplicit bool) {
	r.binaryOrder = binaryOrder
	r.isImplicit = isImplicit
}

func (r *dcmReader) SetFileSize(fileSize int64) error {
	r.fileSize = fileSize
	return nil
}

func (r *dcmReader) GetFileSize() int64 {
	return r.fileSize
}

func (r *dcmReader) readString(n uint32) (string, error) {
	data := make([]byte, n)
	_, err := io.ReadFull(r, data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
