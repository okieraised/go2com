package go2com

import (
	"fmt"
	"github.com/okieraised/go2com/pkg/nifti/nii_io"
	"github.com/stretchr/testify/assert"
	_ "image/jpeg"
	"testing"
)

//func Test_NiiReaderProfiling(t *testing.T) {
//	assert := assert.New(t)
//	filePath := "/home/tripg/workspace/anim3.nii.gz"
//	fn := func() {
//		rd, err := nii_io.NewNiiReader(filePath, nii_io.WithInMemory(true))
//		assert.NoError(err)
//		err = rd.Parse()
//		assert.NoError(err)
//	}
//
//	err := utils.CPUProfilingFunc(fn, "/home/tripg/workspace/nii_reader.pprof")
//	assert.NoError(err)
//}

func TestNiiWriter_FillVoxel(t *testing.T) {
	assert := assert.New(t)

	filePath := "/home/tripg/workspace/anim3.nii.gz"

	rd, err := nii_io.NewNiiReader(filePath, nii_io.WithInMemory(true))
	assert.NoError(err)
	err = rd.Parse()
	assert.NoError(err)

	_, err = nii_io.MakeEmptyImageFromImg(rd.GetNiiData())
	assert.NoError(err)
	//
	//fmt.Println(rd.GetAt(0, 0, 0, 0))
	////fmt.Println(voxels)
}

func TestNiiWriter_FullAnnotation(t *testing.T) {
	assert := assert.New(t)

	filePath := "/home/tripg/workspace/anim3.nii.gz"

	rd, err := nii_io.NewNiiReader(filePath, nii_io.WithInMemory(true))
	assert.NoError(err)
	err = rd.Parse()
	assert.NoError(err)

	//fmt.Println(rd.GetVolume(8))
	//fmt.Println(rd.GetAt(198, 180, 0, 7))
	//return
	//original := rd.GetNiiData()

	//img, err := nii_io.MakeEmptyImageFromImg(rd.GetNiiData())
	//assert.NoError(err)

	writer, err := nii_io.NewNiiWriter("/home/tripg/workspace/anim3_out_annotation.nii",
		nii_io.WithNIfTIData(rd.GetNiiData()),
		nii_io.WithCompression(true),
	)

	//err = writer.SetVolume(img)
	//assert.NoError(err)

	for x := 0; x < 462; x++ {
		for y := 0; y < 364; y++ {
			for z := 0; z < 50; z++ {
				for tt := 0; tt < 9; tt++ {
					curr := rd.GetAt(int64(x), int64(y), int64(z), int64(tt))
					if curr > -200 {
						err = writer.SetAt(1, int64(x), int64(y), int64(z), int64(tt))
						assert.NoError(err)
					} else if curr >= -720 && curr <= 200 {
						err = writer.SetAt(2, int64(x), int64(y), int64(z), int64(tt))
						assert.NoError(err)
					} else {
						err = writer.SetAt(0, int64(x), int64(y), int64(z), int64(tt))
						assert.NoError(err)
					}
				}
			}
		}
	}
	err = writer.WriteToFile()
	assert.NoError(err)
}

func TestNiiWriter_EmptyImageData_Filled(t *testing.T) {
	assert := assert.New(t)

	filePath := "/home/tripg/workspace/anim3.nii.gz"

	rd, err := nii_io.NewNiiReader(filePath, nii_io.WithInMemory(true))
	assert.NoError(err)
	err = rd.Parse()
	assert.NoError(err)

	img, err := nii_io.MakeEmptyImageFromImg(rd.GetNiiData())
	assert.NoError(err)

	writer, err := nii_io.NewNiiWriter("/home/tripg/workspace/anim3_out_blank_filled.nii",
		nii_io.WithNIfTIData(rd.GetNiiData()),
		nii_io.WithCompression(true),
	)

	err = writer.SetVolume(img)
	assert.NoError(err)

	for x := 200; x <= 250; x++ {
		for y := 100; y <= 150; y++ {
			for z := 0; z <= 49; z++ {
				err = writer.SetAt(1, int64(x), int64(y), int64(z), 0)
				assert.NoError(err)
			}
		}
	}

	for x := 300; x <= 400; x++ {
		for y := 200; y <= 300; y++ {
			for z := 0; z <= 49; z++ {
				err = writer.SetAt(1, int64(x), int64(y), int64(z), 0)
				assert.NoError(err)
			}

		}
	}

	err = writer.WriteToFile()
	assert.NoError(err)
}

func TestNiiWriter_EmptyImageData(t *testing.T) {
	assert := assert.New(t)

	filePath := "/home/tripg/workspace/anim3.nii.gz"

	rd, err := nii_io.NewNiiReader(filePath, nii_io.WithInMemory(true))
	assert.NoError(err)
	err = rd.Parse()
	assert.NoError(err)

	img, err := nii_io.MakeEmptyImageFromImg(rd.GetNiiData())
	assert.NoError(err)

	writer, err := nii_io.NewNiiWriter("/home/tripg/workspace/anim3_out_blank.nii",
		nii_io.WithNIfTIData(rd.GetNiiData()),
		nii_io.WithCompression(true),
	)

	err = writer.SetVolume(img)
	assert.NoError(err)

	err = writer.WriteToFile()
	assert.NoError(err)
}

func TestNiiWriter_Single(t *testing.T) {
	assert := assert.New(t)

	filePath := "/home/tripg/workspace/anim3.nii.gz"

	rd, err := nii_io.NewNiiReader(filePath, nii_io.WithInMemory(true))
	assert.NoError(err)
	err = rd.Parse()
	assert.NoError(err)

	writer, err := nii_io.NewNiiWriter("/home/tripg/workspace/anim3_out.nii",
		nii_io.WithNIfTIData(rd.GetNiiData()),
		nii_io.WithCompression(true),
	)

	assert.NoError(err)
	err = writer.WriteToFile()
	assert.NoError(err)
}

func TestNiiWriter_Pair(t *testing.T) {
	assert := assert.New(t)

	filePath := "/home/tripg/workspace/anim3.nii.gz"

	rd, err := nii_io.NewNiiReader(filePath, nii_io.WithInMemory(true))
	assert.NoError(err)
	err = rd.Parse()
	assert.NoError(err)

	writer, err := nii_io.NewNiiWriter("/home/tripg/workspace/anim3_out.nii",
		nii_io.WithNIfTIData(rd.GetNiiData()),
		nii_io.WithWriteHeaderFile(true),
		nii_io.WithCompression(true),
	)
	assert.NoError(err)
	err = writer.WriteToFile()
	assert.NoError(err)
}

func TestNii1(t *testing.T) {
	assert := assert.New(t)

	filePath := "/home/tripg/Documents/nifti/Arnow^Corie^Shelvey^OM_segmented.nii"
	filePath = "/home/tripg/Documents/nifti/RGB16_4D.nii.gz"
	filePath = "/home/tripg/Documents/nifti/someones_anatomy.nii.gz"
	filePath = "/home/tripg/Documents/nifti/someones_epi.nii.gz"
	//filePath = "/home/tripg/Documents/nifti/RGB8_4D.nii.gz"
	//filePath = "/home/tripg/Documents/nifti/jaw.nii.gz"
	//filePath = "/home/tripg/Documents/nifti/Arnow^Corie^Shelvey^OM_segmented.nii"
	//filePath = "/home/tripg/Documents/nifti/knee.nii.gz"
	//filePath = "/home/tripg/Documents/nifti/ExBox11/fmri.nii.gz"
	//filePath = "/home/tripg/Documents/nifti/ExBox11/structural.nii.gz"
	//filePath = "/home/tripg/Documents/nifti/ExBox11/structural_brain.nii.gz"
	filePath = "/home/tripg/Documents/nifti/JHU_MNI_SS_T1.nii.gz"
	//filePath = "/home/tripg/Documents/nifti/avg152T1_LR_nifti2.nii.gz"
	//filePath = "/Users/TriPham/Documents/nifti/avg152T1_RL_nifti.nii"
	filePath = "/home/tripg/Downloads/NIfTI-files/images_structural_unstripped/UPENN-GBM-00630_21/UPENN-GBM-00630_21_FLAIR_unstripped.nii.gz"
	filePath = "/home/tripg/workspace/NIfTI-files/automated_segm/UPENN-GBM-00003_11_automated_approx_segm.nii.gz"
	filePath = "/home/tripg/workspace/niivue-images-main/mni152.nii.gz"
	filePath = "/home/tripg/workspace/niivue-images-main/visiblehuman.nii.gz"
	//filePath = "/home/tripg/workspace/VolumeRenderingData-master/seg3D/MR.nii.gz"

	filePath = "/home/tripg/workspace/nifti/Arnow^Corie^Shelvey^OM_T2.nii.gz"
	filePath = "/home/tripg/workspace/NIfTI-files/automated_segm/UPENN-GBM-00003_11_automated_approx_segm.nii.gz"

	niiReader, err := nii_io.NewNiiReader(filePath, nii_io.WithInMemory(true)) // , nii_io.WithRetainHeader(false)
	assert.NoError(err)
	err = niiReader.Parse()
	assert.NoError(err)

	fmt.Println("shape", niiReader.GetImgShape())
	fmt.Println(niiReader.GetUnitsOfMeasurements())
	fmt.Println("dtype", niiReader.GetDatatype())
	fmt.Println("nbyper", niiReader.GetNiiData().NByPer)
	fmt.Println("GetSFromCode()", niiReader.GetSFormCode())
	fmt.Println("GetQFromCode()", niiReader.GetQFormCode())
	fmt.Println("orientation", niiReader.GetOrientation())
	fmt.Println("affine", niiReader.GetAffine())
	fmt.Println("QOffsetX", niiReader.GetNiiData().QoffsetX)
	fmt.Println("QOffsetY", niiReader.GetNiiData().QoffsetY)
	fmt.Println("QOffsetZ", niiReader.GetNiiData().QoffsetZ)
	fmt.Println("QuaternB", niiReader.GetQuaternB())
	fmt.Println("QuaternC", niiReader.GetQuaternC())
	fmt.Println("QuaternD", niiReader.GetQuaternD())
	fmt.Println("Endian", niiReader.GetNiiData().ByteOrder)
	//fmt.Println("get at", niiReader.GetAt(117, 85, 101, 0))
	fmt.Println("Header", niiReader.GetHeader())
	fmt.Println("PixDim", niiReader.GetNiiData().PixDim)
	//fmt.Println("Descrip", strings.ReplaceAll(string(niiReader.GetNiiData().Descrip[:]), "\x00", ""))
	//fmt.Println(niiReader.GetSlice(77, 0))

	fmt.Println("---------------------------------------------------------------------------------------------------")

	filePath = "/home/tripg/workspace/nifti/Arnow^Corie^Shelvey^OM_segmented.nii.gz"
	filePath = "/home/tripg/workspace/NIfTI-files/images_segm/UPENN-GBM-00009_11_segm.nii.gz"
	filePath = "/home/tripg/workspace/NIfTI-files/images_structural_unstripped/UPENN-GBM-00009_11/UPENN-GBM-00009_11_FLAIR_unstripped.nii.gz"
	filePath = "/home/tripg/workspace/test_segment.nii.gz"
	filePath = "/home/tripg/workspace/anim3.nii.gz"
	//filePath = "/home/tripg/workspace/RGB16_4D.nii.gz"
	//filePath = "/home/tripg/workspace/RGB8_4D.nii.gz"
	//filePath = "/home/tripg/workspace/nifti/JHU_MNI_SS_T1_mask.nii.gz"

	niiReader, err = nii_io.NewNiiReader(filePath, nii_io.WithInMemory(true)) // , nii_io.WithRetainHeader(false)
	assert.NoError(err)
	err = niiReader.Parse()
	assert.NoError(err)

	fmt.Println("shape", niiReader.GetImgShape())
	fmt.Println(niiReader.GetUnitsOfMeasurements())
	fmt.Println("dtype", niiReader.GetDatatype())
	fmt.Println("nbyper", niiReader.GetNiiData().NByPer)
	fmt.Println("GetSFromCode()", niiReader.GetSFormCode())
	fmt.Println("GetQFromCode()", niiReader.GetQFormCode())
	fmt.Println("orientation", niiReader.GetOrientation())
	fmt.Println("affine", niiReader.GetAffine())
	fmt.Println("QOffsetX", niiReader.GetNiiData().QoffsetX)
	fmt.Println("QOffsetY", niiReader.GetNiiData().QoffsetY)
	fmt.Println("QOffsetZ", niiReader.GetNiiData().QoffsetZ)
	fmt.Println("QuaternB", niiReader.GetQuaternB())
	fmt.Println("QuaternC", niiReader.GetQuaternC())
	fmt.Println("QuaternD", niiReader.GetQuaternD())
	fmt.Println("Endian", niiReader.GetNiiData().ByteOrder)
	//fmt.Println("get at", niiReader.GetAt(117, 85, 101, 0))
	fmt.Println("Header", niiReader.GetHeader())
	fmt.Println("PixDim", niiReader.GetNiiData().PixDim)

	//fmt.Println(niiReader.GetVolume(8))
	//fmt.Println(niiReader.GetAt(112, 186, 20, 0))

	fmt.Println(niiReader.GetVolume(2))
	//
	//writer, err := nii_io.NewNiiWriter("./out.nii.gz", nii_io.WithNIfTIData(niiReader.GetNiiData()), nii_io.WithCompression(true))
	//assert.NoError(err)
	//err = writer.WriteToFile()
	//assert.NoError(err)
	//
	//filePath = "./out.nii.gz"
	//niiReader, err = nii_io.NewNiiReader(filePath, nii_io.WithInMemory(true))
	//assert.NoError(err)
	//err = niiReader.Parse()
	//assert.NoError(err)
	//
	//shape = niiReader.GetImgShape()
	//fmt.Println("shape", shape)
	//fmt.Println(niiReader.GetUnitsOfMeasurements())
	//fmt.Println("dtype", niiReader.GetDatatype())
	//fmt.Println("nbyper", niiReader.GetNiiData().NByPer)
	//fmt.Println("GetSFromCode()", niiReader.GetSFormCode())
	//fmt.Println("GetQFromCode()", niiReader.GetQFormCode())
	//fmt.Println("orientation", niiReader.GetOrientation())
	//fmt.Println("affine", niiReader.GetAffine())
	//fmt.Println("QOffsetX", niiReader.GetNiiData().QoffsetX)
	//fmt.Println("QOffsetY", niiReader.GetNiiData().QoffsetY)
	//fmt.Println("QOffsetZ", niiReader.GetNiiData().QoffsetZ)

	//res, err := niiReader.GetSlice(1, 0)
	//assert.NoError(err)
	//
	//fmt.Println("len res", len(res))
	//fmt.Println("len res0", len(res[0]))
	//
	////fmt.Println(res)
	//res2, err := niiReader.GetVolume(0)
	//assert.NoError(err)
	//
	//fmt.Println("len res2", len(res2))
	//
	//writer := nii_io.NewNiiWriter("./out.nii", nii_io.WithNIfTIData(niiReader.GetNiiData()))
	//
	//err = writer.WriteToFile()
	//assert.NoError(err)

	//for _, elem := range res {
	//	fmt.Println(elem)
	//
	//}

	//fmt.Println(niiReader.GetTimeSeries(26, 30, 16))

	//for x := 0; x < 57; x++ {
	//	for y := 0; y < 67; y++ {
	//		for z := 0; z < 56; z++ {
	//			fmt.Println(niiReader.GetTimeSeries(int64(x), int64(y), int64(z)))
	//		}
	//	}
	//}
}

func TestNii2(t *testing.T) {
	assert := assert.New(t)

	filePath := "/home/tripg/Documents/nifti/Arnow^Corie^Shelvey^OM_segmented.nii"
	filePath = "/home/tripg/Documents/nifti/RGB16_4D.nii.gz"
	filePath = "/home/tripg/Documents/nifti/someones_anatomy.nii.gz"
	filePath = "/home/tripg/Documents/nifti/someones_epi.nii.gz"
	//filePath = "/home/tripg/Documents/nifti/RGB8_4D.nii.gz"
	//filePath = "/home/tripg/Documents/nifti/jaw.nii.gz"
	//filePath = "/home/tripg/Documents/nifti/Arnow^Corie^Shelvey^OM_segmented.nii"
	//filePath = "/home/tripg/Documents/nifti/knee.nii.gz"
	//filePath = "/home/tripg/Documents/nifti/ExBox11/fmri.nii.gz"
	//filePath = "/home/tripg/Documents/nifti/ExBox11/structural.nii.gz"
	//filePath = "/home/tripg/Documents/nifti/ExBox11/structural_brain.nii.gz"
	//filePath = "/home/tripg/Documents/nifti/JHU_MNI_SS_T1.nii.gz"
	//filePath = "/home/tripg/Documents/nifti/avg152T1_LR_nifti2.nii.gz"
	filePath = "/home/tripg/Documents/nifti/avg152T1_LR_nifti.img.gz"

	niiReader, err := nii_io.NewNiiReader(filePath, nii_io.WithInMemory(true), nii_io.WithHeaderFile("/home/tripg/Documents/nifti/avg152T1_LR_nifti.hdr.gz"))
	assert.NoError(err)
	err = niiReader.Parse()
	assert.NoError(err)

	shape := niiReader.GetImgShape()
	fmt.Println("shape", shape)
	fmt.Println(niiReader.GetUnitsOfMeasurements())
	fmt.Println("dtype", niiReader.GetDatatype())
	fmt.Println("nbyper", niiReader.GetNiiData().NByPer)
	fmt.Println("GetSFromCode()", niiReader.GetSFormCode())
	fmt.Println("GetQFromCode()", niiReader.GetQFormCode())
	fmt.Println("orientation", niiReader.GetOrientation())
	fmt.Println("affine", niiReader.GetAffine())
	fmt.Println("QOffsetX", niiReader.GetNiiData().QoffsetX)
	fmt.Println("QOffsetY", niiReader.GetNiiData().QoffsetY)
	fmt.Println("QOffsetZ", niiReader.GetNiiData().QoffsetZ)
	fmt.Println("VoxOffset", niiReader.GetNiiData().VoxOffset)

	fmt.Println("---------------------------------------------------------------------------------------------------")

	writer, err := nii_io.NewNiiWriter("./out2.nii.gz", nii_io.WithNIfTIData(niiReader.GetNiiData()), nii_io.WithCompression(true))
	assert.NoError(err)
	err = writer.WriteToFile()
	assert.NoError(err)

	filePath = "./out2.nii.gz"
	niiReader, err = nii_io.NewNiiReader(filePath, nii_io.WithInMemory(true))
	assert.NoError(err)
	err = niiReader.Parse()
	assert.NoError(err)

	shape = niiReader.GetImgShape()
	fmt.Println("shape", shape)
	fmt.Println(niiReader.GetUnitsOfMeasurements())
	fmt.Println("dtype", niiReader.GetDatatype())
	fmt.Println("nbyper", niiReader.GetNiiData().NByPer)
	fmt.Println("GetSFromCode()", niiReader.GetSFormCode())
	fmt.Println("GetQFromCode()", niiReader.GetQFormCode())
	fmt.Println("orientation", niiReader.GetOrientation())
	fmt.Println("affine", niiReader.GetAffine())
	fmt.Println("QOffsetX", niiReader.GetNiiData().QoffsetX)
	fmt.Println("QOffsetY", niiReader.GetNiiData().QoffsetY)
	fmt.Println("QOffsetZ", niiReader.GetNiiData().QoffsetZ)
}

//func TestMagicString(t *testing.T) {
//	fmt.Println([]byte("ni2"))
//	fmt.Println(110, 105, 50, 0, 13, 10, 26, 10)
//	// 0D 0A 1A 0A
//	x := make([]byte, 100, 100)
//
//	fmt.Println(x)
//}

//func TestFloatToBytes(t *testing.T) {
//	var original = []byte{0x00, 0x00, 0x00, 0x11, 0x12, 0x64, 0x80, 0x00, 0x59, 0x57}
//	var x float64 = 1
//	var buf bytes.Buffer
//	err := binary.Write(&buf, binary.LittleEndian, x)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(buf.Bytes())
//
//	fmt.Println("1", original)
//	//bytes.ReplaceAll(original[3:4], []byte{0x11}, []byte{0x16})
//
//	copy(original[3:4], []byte{0x96})
//
//	fmt.Println("2", original)
//
//	y := []int{1, 2, 34, 56}
//
//	fmt.Println(y[:3])
//
//}
