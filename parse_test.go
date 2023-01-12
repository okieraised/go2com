package go2com

import (
	"fmt"
	"github.com/okieraised/go2com/internal/utils"
	"github.com/okieraised/go2com/pkg/dicom/iod"
	"github.com/okieraised/go2com/pkg/dicom/tag"
	"github.com/stretchr/testify/assert"
	_ "image/jpeg"
	"os"
	"strings"
	"testing"
)

func TestProfilingParse1(t *testing.T) {
	assert := assert.New(t)
	fn := func() {
		InitTagDict()
		file, err := os.Open("./test_data/01.dcm")
		assert.NoError(err)

		defer file.Close()
		info, err := file.Stat()
		assert.NoError(err)
		fileSize := info.Size()

		parser, err := NewParser(file, fileSize, false, false)
		assert.NoError(err)
		err = parser.Parse()
		assert.NoError(err)
	}
	err := utils.CPUProfilingFunc(fn, "./cpu1.pprof")
	assert.NoError(err)
}

func TestProfilingParse2(t *testing.T) {
	assert := assert.New(t)
	fn := func() {
		InitTagDict()
		filePaths, err := utils.ReadDirRecursively("/home/tripg/Documents/dicom/mammo_dicoms")
		assert.NoError(err)
		for _, fPath := range filePaths {
			fmt.Println("process:", fPath)
			file, err := os.Open(fPath)
			assert.NoError(err)
			if err != nil {
				fmt.Println(err)
				return
			}

			defer file.Close()
			info, err := file.Stat()
			assert.NoError(err)
			if err != nil {
				fmt.Println(err)
				return
			}
			fileSize := info.Size()

			parser, err := NewParser(file, fileSize, false, false)
			assert.NoError(err)
			err = parser.Parse()
			assert.NoError(err)
			if err != nil && !strings.Contains(err.Error(), "could not find tag") {
				fmt.Println(err)
				return
			}
		}
	}
	err := utils.CPUProfilingFunc(fn, "./cpu2.pprof")
	assert.NoError(err)
}

func TestProfilingParse3(t *testing.T) {
	assert := assert.New(t)
	fn := func() {
		InitTagDict()
		filePaths, err := utils.ReadDirRecursively("/home/tripg/Documents/dicom/test_data")
		assert.NoError(err)
		for _, fPath := range filePaths {
			fmt.Println("process:", fPath)
			file, err := os.Open(fPath)
			assert.NoError(err)
			if err != nil {
				fmt.Println(err)
				return
			}

			defer file.Close()
			info, err := file.Stat()
			assert.NoError(err)
			if err != nil {
				fmt.Println(err)
				return
			}
			fileSize := info.Size()

			parser, err := NewParser(file, fileSize, false, false)
			assert.NoError(err)
			err = parser.Parse()
			assert.NoError(err)
			if err != nil && !strings.Contains(err.Error(), "could not find tag") {
				fmt.Println(err)
				return
			}
		}
	}
	err := utils.CPUProfilingFunc(fn, "./cpu3.pprof")
	assert.NoError(err)
}

func TestNewParser1(t *testing.T) {
	assert := assert.New(t)
	InitTagDict()
	file, err := os.Open("/home/tripg/Documents/dicom/mammo_dicoms/1.2.840.113619.2.255.10452022879169.3670200508103440.2701.dicom")
	assert.NoError(err)

	defer file.Close()
	info, err := file.Stat()
	assert.NoError(err)
	fileSize := info.Size()

	parser, err := NewParser(file, fileSize, true, false)
	assert.NoError(err)
	err = parser.Parse()
	assert.NoError(err)

	parser.Export(false)

	//tt := parser.Export(false)
	//for key := range tt {
	//	val := tt[key].Value.([]interface{})
	//	for _, subVal := range val {
	//		fmt.Println(reflect.ValueOf(subVal).Kind())
	//	}
	//}
}

func TestNewParser2(t *testing.T) {
	assert := assert.New(t)
	InitTagDict()
	filePaths, err := utils.ReadDirRecursively("/home/tripg/Documents/dicom/test_data")
	assert.NoError(err)
	for _, fPath := range filePaths {
		fmt.Println("process:", fPath)
		file, err := os.Open(fPath)
		assert.NoError(err)

		defer file.Close()
		info, err := file.Stat()
		assert.NoError(err)
		fileSize := info.Size()

		parser, err := NewParser(file, fileSize, false, false)
		assert.NoError(err)
		err = parser.Parse()
		if err != nil && !strings.Contains(err.Error(), "could not find tag") {
			fmt.Println(err)
			return
		}
		if err != nil && !strings.Contains(err.Error(), "not in valid dicom format") {
			fmt.Println(err)
			continue
		}
		uids, err := parser.dataset.RetrieveFileUID()
		err = parser.Parse()
		fmt.Println(uids.StudyInstanceUID, uids.SeriesInstanceUID, uids.SOPInstanceUID)
		_ = parser.Export(false)
	}
}

func TestNewParser3(t *testing.T) {
	assert := assert.New(t)
	filePaths, err := utils.ReadDirRecursively("/home/tripg/Documents/dicom/mammo_dicoms")
	assert.NoError(err)
	for _, fPath := range filePaths {
		InitTagDict()
		fmt.Println("process:", fPath)
		file, err := os.Open(fPath)
		assert.NoError(err)

		defer file.Close()
		info, err := file.Stat()
		assert.NoError(err)
		fileSize := info.Size()

		parser, err := NewParser(file, fileSize, false, false)
		assert.NoError(err)
		err = parser.Parse()
		assert.NoError(err)

		_ = parser.Export(false)

		pixelData := iod.GetPixelDataMacroAttributes(parser.dataset, parser.metadata)
		pixelData.GetExpectedPixelData()
		valid := pixelData.ValidatePixelData()
		if !valid {
			fmt.Println("Invalid", fPath)
		}

	}
}

func TestNewParser4(t *testing.T) {
	assert := assert.New(t)
	InitTagDict()
	file, err := os.Open("/home/tripg/Documents/dicom/oct/1.dcm")
	assert.NoError(err)

	defer file.Close()
	info, err := file.Stat()
	assert.NoError(err)
	fileSize := info.Size()

	parser, err := NewParser(file, fileSize, true, false)
	assert.NoError(err)
	err = parser.Parse()
	assert.NoError(err)

	for _, elem := range parser.dataset.Elements {
		fmt.Println(elem)
	}
}

func TestNewParser5(t *testing.T) {
	assert := assert.New(t)
	InitTagDict()
	//file, err := os.Open("/home/tripg/Documents/dicom/ptt1.dcm")
	//file, err := os.Open("/home/tripg/Documents/dicom/dicoms_mr_func/MR.1.3.46.670589.11.38317.5.0.4476.2014042516042547586")
	//file, err := os.Open("/home/tripg/Documents/dicom/dicoms_struct/N2D_0001.dcm")
	//file, err := os.Open("/home/tripg/Documents/dicom/test_data/File 10051.dcm")
	//file, err := os.Open("/home/tripg/Downloads/1.2.840.113619.2.278.3.717616.166.1580339214.7.99.dcm")
	//file, err := os.Open("/home/tripg/Documents/dicom/test_data/40009")
	//file, err := os.Open("/home/tripg/Documents/dicom/test_data/File 12943.dcm")
	//file, err := os.Open("/home/tripg/Documents/dicom/mammo_dicoms/1.3.12.2.1107.5.12.7.3367.30000018112001512650000000209.dicom")
	//file, err := os.Open("/home/tripg/Documents/dicom/2_skull_ct/DICOM/I0")
	//file, err := os.Open("/home/tripg/Documents/dicom/Class-3-malocclusion/Class 3 malocclusion/DICOM/I0")
	//file, err := os.Open("/home/tripg/Documents/img2.dcm")
	//file, err := os.Open("/home/tripg/Documents/dicom/color_dicom/JPEG2000-RGB.dcm")
	//file, err := os.Open("/home/tripg/Documents/dicom/color_dicom/JPEG2000-YBR_FULL.dcm")
	//file, err := os.Open("/home/tripg/Documents/dicom/color_dicom/JPEGLS-RGB.dcm")
	//file, err := os.Open("/home/tripg/Documents/dicom/color_dicom/losslessJPEG-RGB.dcm")
	//file, err := os.Open("/home/tripg/Documents/dicom/MammoTomoUPMC_Case22/Case22 [Case22]/20071030 021043 [ - MAMMOGRAM DIGITAL DX BILAT]/Series 71100000 [MG - L CC]/1.3.6.1.4.1.5962.99.1.2280943358.716200484.1363785608958.480.0.dcm")
	//file, err := os.Open("/home/tripg/Documents/dicom/us_valid_pixel_aspect.dcm")
	//file, err := os.Open("/home/tripg/Documents/dicom/KiTS-00072/04-01-2000-abdomenw-15076/2.000000-arterial-99348/1-001.dcm")
	//file, err := os.Open("/home/tripg/Documents/dicom/us_monochrome2.dcm")
	//file, err := os.Open("/home/tripg/Documents/dicom/seg_abdomen_r1/")
	//file, err := os.Open("/home/tripg/Documents/dicom/perfusion_ct/CT0014")
	//file, err := os.Open("/home/tripg/Documents/dicom/US-RGB-8-esopecho")
	//file, err := os.Open("/home/tripg/Documents/dicom/KiTS-00072/04-01-2000-abdomenw-15076/300.000000-Segmentation-99191/1-1.dcm")
	//file, err := os.Open("/home/tripg/Documents/dicom/9947.LEFT_MLO.dcm")
	//file, err := os.Open("/home/tripg/Documents/dicom/utf8test.dcm")
	//file, err := os.Open("/home/tripg/Documents/dicom/samples-of-mr-images-1.0.0/E1154S7I.dcm")
	//file, err := os.Open("/home/tripg/Documents/dicom/GSPS_Liver/DICOM/IM_0001")
	//file, err := os.Open("/home/tripg/Documents/dicom/us_valid_pixel_aspect.dcm")
	//file, err := os.Open("/home/tripg/Documents/dicom/US-RGB-8-esopecho")
	//file, err := os.Open("/home/tripg/Documents/dicom/Class-3-malocclusion/Class 3 malocclusion/DICOM/I0")
	//file, err := os.Open("/home/tripg/Documents/dicom/MammoTomoUPMC_Case4/Case4 [Case4]/20071218 093012 [ - MAMMOGRAM DIGITAL SCR BILAT]/Series 73100000 [MG - R CC Tomosynthesis Reconstruction]/1.3.6.1.4.1.5962.99.1.2280943358.716200484.1363785608958.589.0.dcm")
	//file, err := os.Open("/home/tripg/Documents/dicom/test_full/063.dcm")
	//file, err := os.Open("/home/tripg/Documents/dicom/test_full/068.dcm")
	file, err := os.Open("/home/tripg/Documents/dicom/10142022/Acuson/Sequoia/EXAMS/EXAM0003/CLIPS/CLIP0039")
	assert.NoError(err)

	defer file.Close()
	info, err := file.Stat()
	assert.NoError(err)
	fileSize := info.Size()

	parser, err := NewParser(file, fileSize, false, false)
	assert.NoError(err)
	err = parser.Parse()
	assert.NoError(err)

	for _, elem := range parser.dataset.Elements {
		if elem.Tag == tag.BluePaletteColorLookupTableData || elem.Tag == tag.RedPaletteColorLookupTableData || elem.Tag == tag.GreenPaletteColorLookupTableData {
			continue
		}
		if elem.ValueRepresentationStr == "OB" || elem.ValueRepresentationStr == "OW" {
			continue
		}
		fmt.Println(elem)
	}
}

func TestNewParser6(t *testing.T) {
	assert := assert.New(t)
	InitTagDict()
	file, err := os.Open("/home/tripg/Documents/dicom/MammoTomoUPMC_Case4/Case4 [Case4]/20071218 093012 [ - MAMMOGRAM DIGITAL SCR BILAT]/Series 73200000 [MG - R CC Breast Tomosynthesis Image]/1.3.6.1.4.1.5962.99.1.2280943358.716200484.1363785608958.597.0.dcm")
	assert.NoError(err)

	defer file.Close()
	info, err := file.Stat()
	assert.NoError(err)
	fileSize := info.Size()

	parser, err := NewParser(file, fileSize, false, false)
	assert.NoError(err)
	err = parser.Parse()
	assert.NoError(err)

	//for _, elem := range parser.dataset.Elements {
	//	fmt.Println(elem)
	//}

	//pixelData := iod.GetPixelDataMacroAttributes(parser.dataset, parser.metadata)
	//pixelData.GetExpectedPixelData()
	//valid := pixelData.ValidatePixelData()
	//fmt.Println(valid)

	//tt := parser.Export(false)
	//for k := range tt {
	//	fmt.Println(k, tt[k])
	//}

}

func TestNewParser7(t *testing.T) {
	assert := assert.New(t)
	filePaths, err := utils.ReadDirRecursively("/home/tripg/workspace/dicom/test_full")
	assert.NoError(err)
	for _, fPath := range filePaths {
		InitTagDict()
		fmt.Println("process:", fPath)
		if strings.Contains(fPath, "1706") {
			continue
		}
		file, err := os.Open(fPath)
		assert.NoError(err)

		defer file.Close()
		info, err := file.Stat()
		assert.NoError(err)
		fileSize := info.Size()

		parser, err := NewParser(file, fileSize, false, false)
		assert.NoError(err)
		err = parser.Parse()
		assert.NoError(err)

		_ = parser.Export(false)

		pixelData := iod.GetPixelDataMacroAttributes(parser.dataset, parser.metadata)
		pixelData.GetExpectedPixelData()
		valid := pixelData.ValidatePixelData()
		if !valid {
			fmt.Println("Invalid", fPath)
		}
	}
}

func TestNewParser8(t *testing.T) {
	assert := assert.New(t)
	InitTagDict()
	file, err := os.Open("/home/tripg/workspace/10142022/ALI_Technologies/UltraPACS/studies/w0019837/view0001")
	assert.NoError(err)

	defer file.Close()
	info, err := file.Stat()
	assert.NoError(err)
	fileSize := info.Size()

	parser, err := NewParser(file, fileSize, false, false)
	assert.NoError(err)
	err = parser.Parse()
	assert.NoError(err)

	//seriesTag := parser.ExportSeriesTags()
	//for k := range seriesTag {
	//	fmt.Println(k, seriesTag[k])
	//}

	for _, elem := range parser.dataset.Elements {
		fmt.Println(elem)
	}

	//pixelData := iod.GetPixelDataMacroAttributes(parser.dataset, parser.metadata)
	//pixelData.GetExpectedPixelData()
	//valid := pixelData.ValidatePixelData()
	//fmt.Println(valid)

	//tt := parser.Export(false)
	//for k := range tt {
	//	fmt.Println(k, tt[k])
	//}

}

func TestNewParser9(t *testing.T) {
	assert := assert.New(t)
	filePaths, err := utils.ReadDirRecursively("/home/tripg/workspace/dicom/vinlab/Mini-batch0")
	assert.NoError(err)
	for _, fPath := range filePaths {
		InitTagDict()
		fmt.Println("process:", fPath)
		file, err := os.Open(fPath)
		assert.NoError(err)

		defer file.Close()
		info, err := file.Stat()
		assert.NoError(err)
		fileSize := info.Size()

		parser, err := NewParser(file, fileSize, false, false)
		assert.NoError(err)
		err = parser.Parse()
		assert.NoError(err)

		_ = parser.Export(false)

		pixelData := iod.GetPixelDataMacroAttributes(parser.dataset, parser.metadata)
		pixelData.GetExpectedPixelData()
		valid := pixelData.ValidatePixelData()
		if !valid {
			fmt.Println("Invalid", fPath)
		}
	}
}

func TestNewParser10(t *testing.T) {
	assert := assert.New(t)
	InitTagDict()
	file, err := os.Open("/home/tripg/workspace/dicom/test_full/1706.dcm")
	assert.NoError(err)

	defer file.Close()
	info, err := file.Stat()
	assert.NoError(err)
	fileSize := info.Size()

	parser, err := NewParser(file, fileSize, false, false)
	assert.NoError(err)
	err = parser.Parse()
	assert.NoError(err)
}

func TestNewParser11(t *testing.T) {
	assert := assert.New(t)
	InitTagDict()
	file, err := os.Open("/home/tripg/Downloads/123.241606668321866.1722978010541148/DICOM/1.2.840.113619.2.427.84108138632.1643160910.120.dicom")
	file, err = os.Open("/home/tripg/workspace/10142022/ALI_Technologies/UltraPACS/studies/w0055053/view0013")
	file, err = os.Open("/home/tripg/workspace/10142022/Acuson/Sequoia/EXAMS/EXAM0000/CLIPS/CLIP0031")
	//file, err = os.Open("/home/tripg/workspace/10142022/Hamamatsu/Dog_15x15_20x.dcm")
	//file, err = os.Open("/home/tripg/Downloads/N2D0027.dcm")
	//file, err = os.Open("/home/tripg/Downloads/123.241606668321866.1724728615648318_en.dcm")
	//file, err = os.Open("/home/tripg/Downloads/1-1.dcm")
	//file, err = os.Open("/home/tripg/workspace/dicom2/PrivateGEImplicitVRBigEndianTransferSyntax16Bits.dcm")
	assert.NoError(err)

	defer file.Close()
	info, err := file.Stat()
	assert.NoError(err)
	fileSize := info.Size()

	parser, err := NewParser(file, fileSize, true, false)
	assert.NoError(err)
	err = parser.Parse()
	assert.NoError(err)

	for _, elem := range parser.dataset.Elements {
		fmt.Println(elem)
	}
}
