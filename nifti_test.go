package go2com

import (
	"fmt"
	_ "image/jpeg"
	"strings"
	"testing"

	"github.com/okieraised/go2com/pkg/nifti/nii_io"
	"github.com/stretchr/testify/assert"
)

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

	niiReader, err := nii_io.NewNiiReader(filePath, nii_io.WithInMemory(true)) // , nii_io.WithRetainHeader(false)
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
	fmt.Println("QuaternB", niiReader.GetQuaternB())
	fmt.Println("QuaternC", niiReader.GetQuaternC())
	fmt.Println("QuaternD", niiReader.GetQuaternD())
	fmt.Println("Endian", niiReader.GetNiiData().ByteOrder)
	//fmt.Println("get at", niiReader.GetAt(117, 85, 101, 0))
	fmt.Println("Header", niiReader.GetHeader())
	fmt.Println("PixDim", niiReader.GetNiiData().PixDim)
	fmt.Println("Descrip", strings.ReplaceAll(string(niiReader.GetNiiData().Descrip[:]), "\x00", ""))
	//fmt.Println(niiReader.GetSlice(77, 0))

	//fmt.Println("---------------------------------------------------------------------------------------------------")
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

	shape = niiReader.GetImgShape()
	fmt.Println(shape)

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

func TestMagicString(t *testing.T) {
	fmt.Println([]byte("ni1"))
}
