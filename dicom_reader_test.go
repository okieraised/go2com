package go2com

import (
	"bufio"
	"fmt"
	"github.com/okieraised/go2com/internal/utils"
	"github.com/okieraised/go2com/pkg/dicom/dcm_io"
	"github.com/stretchr/testify/assert"
	_ "image/jpeg"
	"os"
	"strings"
	"testing"
)

func Test_ReaderProfiling(t *testing.T) {
	assert := assert.New(t)
	f, err := os.Open("./test_data/01.dcm")
	assert.NoError(err)

	fInfo, err := f.Stat()
	assert.NoError(err)

	fn := func() {
		dcmReader := dcm_io.NewDICOMReader(bufio.NewReader(f), dcm_io.WithSkipPixelData(false), dcm_io.WithSkipDataset(false), dcm_io.WithSetFileSize(fInfo.Size()))
		err = dcmReader.Parse()
		assert.NoError(err)
	}

	err = utils.CPUProfilingFunc(fn, "./cpu1.pprof")
	assert.NoError(err)
}

func Test_NewParser_PDF(t *testing.T) {
	assert := assert.New(t)
	filePaths, err := utils.ReadDirRecursively("/home/tripg/Downloads/IMG-0003-00001 (1).dcm")
	assert.NoError(err)
	for _, fPath := range filePaths {
		fmt.Println("process:", fPath)

		f, err := os.Open(fPath)
		assert.NoError(err)

		fInfo, err := f.Stat()
		assert.NoError(err)

		dcmReader := dcm_io.NewDICOMReader(bufio.NewReader(f), dcm_io.WithSetFileSize(fInfo.Size()))

		err = dcmReader.Parse()
		if err != nil {
			fmt.Println(err)
			return
		}
		assert.NoError(err)

		//for _, elem := range dcmReader.GetDataset().Elements {
		//	fmt.Println(elem)
		//}
		fmt.Println(dcmReader.ExportSeriesTags())

	}
}

func Test_NewParser1(t *testing.T) {
	assert := assert.New(t)
	filePaths, err := utils.ReadDirRecursively("/home/tripg/workspace/dicom/test_data")
	assert.NoError(err)
	for _, fPath := range filePaths {
		fmt.Println("process:", fPath)

		f, err := os.Open(fPath)
		assert.NoError(err)

		fInfo, err := f.Stat()
		assert.NoError(err)

		dcmReader := dcm_io.NewDICOMReader(bufio.NewReader(f), dcm_io.WithSetFileSize(fInfo.Size()))

		err = dcmReader.Parse()
		if err != nil {
			fmt.Println(err)
			return
		}
		assert.NoError(err)
	}
}

func Test_NewParser2(t *testing.T) {
	assert := assert.New(t)
	filePaths, err := utils.ReadDirRecursively("/home/tripg/workspace/dicom/test_full")
	assert.NoError(err)
	for _, fPath := range filePaths {
		fmt.Println("process:", fPath)
		f, err := os.Open(fPath)
		assert.NoError(err)

		fInfo, err := f.Stat()
		assert.NoError(err)

		dcmReader := dcm_io.NewDICOMReader(bufio.NewReader(f), dcm_io.WithSetFileSize(fInfo.Size()))

		err = dcmReader.Parse()
		if err != nil {
			if strings.Contains(err.Error(), "not in valid dicom format") {
				continue
			}
		}
		assert.NoError(err)
	}
}

func Test_NewParser3(t *testing.T) {
	assert := assert.New(t)
	filePaths, err := utils.ReadDirRecursively("/home/tripg/workspace/dicom/mammo_dicoms")
	assert.NoError(err)
	for _, fPath := range filePaths {
		fmt.Println("process:", fPath)

		f, err := os.Open(fPath)
		assert.NoError(err)

		fInfo, err := f.Stat()
		assert.NoError(err)

		dcmReader := dcm_io.NewDICOMReader(bufio.NewReader(f), dcm_io.WithSetFileSize(fInfo.Size()))

		err = dcmReader.Parse()
		assert.NoError(err)

		_ = dcmReader.ExportDatasetTags(false)
	}
}

func Test_NewParser4(t *testing.T) {
	assert := assert.New(t)
	fPath := "/home/tripg/workspace/dicom/mammo_dicoms/1.2.840.113619.2.255.10452022879169.3670200508103440.2701.dicom"
	f, err := os.Open(fPath)
	assert.NoError(err)

	fInfo, err := f.Stat()
	assert.NoError(err)

	parser := dcm_io.NewDICOMReader(bufio.NewReader(f), dcm_io.WithSkipPixelData(true), dcm_io.WithSkipDataset(true), dcm_io.WithSetFileSize(fInfo.Size()))
	err = parser.Parse()
	assert.NoError(err)

	for _, elem := range parser.GetDataset().Elements {
		fmt.Println(elem)
	}
}

func Test_NewParser5(t *testing.T) {
	assert := assert.New(t)
	fPath := "/home/tripg/workspace/10142022/ALI_Technologies/UltraPACS/studies/w0019837/view0001"
	f, err := os.Open(fPath)
	assert.NoError(err)

	fInfo, err := f.Stat()
	assert.NoError(err)

	parser := dcm_io.NewDICOMReader(bufio.NewReader(f), dcm_io.WithSkipPixelData(false), dcm_io.WithSkipDataset(false), dcm_io.WithSetFileSize(fInfo.Size()))
	err = parser.Parse()
	assert.NoError(err)

}

func Test_NewParser6(t *testing.T) {
	assert := assert.New(t)
	f, err := os.Open("/home/tripg/workspace/10142022/ALI_Technologies/UltraPACS/studies/w0055053/view0013")
	//f, err = os.Open("/home/tripg/workspace/10142022/Acuson/Sequoia/EXAMS/EXAM0000/CLIPS/CLIP0031")
	f, err = os.Open("/home/tripg/workspace/10142022/Hamamatsu/Dog_15x15_20x.dcm")
	//f, err = os.Open("/home/tripg/workspace/dicom2/PrivateGEImplicitVRBigEndianTransferSyntax16Bits.dcm")
	assert.NoError(err)

	fInfo, err := f.Stat()
	assert.NoError(err)

	parser := dcm_io.NewDICOMReader(bufio.NewReader(f), dcm_io.WithSkipPixelData(false), dcm_io.WithSkipDataset(false), dcm_io.WithSetFileSize(fInfo.Size()))
	err = parser.Parse()
	assert.NoError(err)
}
