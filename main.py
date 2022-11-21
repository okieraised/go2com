from pydicom import dcmread
import pydicom
fpath = "./test_data/File 12948.dcm"
# fpath = "./test_data/File 1.dcm"
fpath = "./test_data/File 18001"
fpath = "./test_data/File 32000"
fpath = "./test_data/File 11636.dcm"
fpath = "./test_data/File 32000"
fpath = "./test_data/File 4000.dcm"
fpath = "./test_data/File 12000"
fpath = "./test_data/File 160.dcm"
fpath = "./test_data/File 8000"
# fpath = "./test_data/File 12000"
fpath = '/home/tripg/Documents/1.dcm'
fpath = '/home/tripg/Downloads/mammo_dicoms/1.3.12.2.1107.5.12.7.5054.30000019090500530000000000556.dicom'
fpath = '/home/tripg/Documents/dicom/test_data/400014'
fpath = '/home/tripg/Documents/dicom/mammo_dicoms/1.2.840.113619.2.255.10452022879169.3670200508103440.2701.dicom'
fpath = '/home/tripg/Documents/dicom/test_data/File 12986'
fpath = '/home/tripg/Documents/dicom/ptt1.dcm'
fpath = '/home/tripg/Documents/dicom/test_data/File 10000.dcm'
fpath = '/home/tripg/Documents/dicom/dicoms_mr_func/MR.1.3.46.670589.11.38317.5.0.4476.2014042516042547586'
fpath = '/home/tripg/Documents/dicom/dicoms_struct/N2D_0001.dcm'
fpath = '/home/tripg/Documents/dicom/test_data/File 10051.dcm'
fpath = '/home/tripg/Documents/dicom/test_data/40009'
fpath = '/home/tripg/Documents/dicom/test_data/File 12943.dcm'
fpath = '/home/tripg/Documents/dicom/mammo_dicoms/1.3.12.2.1107.5.12.7.3367.30000018112001512650000000209.dicom'
fpath = '/home/tripg/Documents/dicom/2_skull_ct/DICOM/I0'
fpath = '/home/tripg/Documents/dicom/Class-3-malocclusion/Class 3 malocclusion/DICOM/I0'
fpath = '/home/tripg/Documents/dicom/ptt1.dcm'
fpath = '/home/tripg/Documents/dicom/us_valid_pixel_aspect.dcm'
fpath = '/home/tripg/Documents/dicom/color_dicom/JPEG2000-RGB.dcm'
fpath = '/home/tripg/Documents/dicom/055829-00000000.dcm'
fpath = '/home/tripg/Documents/dicom/us_monochrome2.dcm'
# fpath = '/home/tripg/Documents/dicom/KiTS-00072/04-01-2000-abdomenw-15076/300.000000-Segmentation-99191/1-1.dcm'
# fpath = '/home/tripg/Documents/dicom/MammoTomoUPMC_Case22/Case22 [Case22]/20071030 022108 [ - BREAST IMAGING TOMOSYNTHESIS]/Series 003 [SR]/1.3.6.1.4.1.5962.99.1.2280943358.716200484.1363785608958.476.0.dcm'
# fpath = '/home/tripg/Documents/dicom/MammoTomoUPMC_Case4/Case4 [Case4]/20071218 093012 [ - MAMMOGRAM DIGITAL SCR BILAT]/Series 73100000 [MG - R CC Tomosynthesis Reconstruction]/1.3.6.1.4.1.5962.99.1.2280943358.716200484.1363785608958.589.0.dcm'
# fpath = '/home/tripg/Documents/dicom/test_full/063.dcm'
fpath = '/home/tripg/Documents/dicom/MammoTomoUPMC_Case4/Case4 [Case4]/20071218 093012 [ - MAMMOGRAM DIGITAL SCR BILAT]/Series 73100000 [MG - R CC Tomosynthesis Reconstruction]/1.3.6.1.4.1.5962.99.1.2280943358.716200484.1363785608958.589.0.dcm'
fpath = '/home/tripg/Documents/dicom/invalid_pixel.dcm'
fpath = '/home/tripg/Documents/dicom/Le_Cong_Giuong/1.2.840.113704.9.1000.16.2.20220421110622828000200020001.dcm'
fpath = '/home/tripg/Documents/dicom/dcm/vietnhat/test/2022/04/03/1.2.392.200036.9123.100.12.12.31738.90220403093121413382737448/1.2.392.200036.9123.100.12.12.31738.90220403094010178790824826.dcm'
fpath = '/home/tripg/Documents/dicom/10142022/Acuson/Sequoia/EXAMS/EXAM0003/CLIPS/CLIP0039'
fpath = '/home/tripg/Documents/x2.dcm'
fpath = '/home/tripg/Documents/dicom/test_full/023.dcm'
fpath = '/home/tripg/Documents/dicom/10142022/ALI_Technologies/UltraPACS/studies/w0019837/view0001'
fpath = '/home/tripg/Documents/dicom/vinlab/Mini-batch0/1.3.12.2.1107.5.1.7.112561.30000019122607094739800003704/DICOM/1.3.12.2.1107.5.1.7.112561.30000019122622575003400000100.dcm'
fpath = '/home/tripg/Downloads/11.10.1106.1133.119.1227.20190308102324.222.dicom_2515042_1668816018440596.dcm'
fpath = '/home/tripg/Downloads/11.10.1106.1133.119.1227.20190308102324.222.dicom_2515042_1668816281813791.dcm'

ds = dcmread(fpath, force=True)
print(ds)

expected = pydicom.pixel_data_handlers.util.get_expected_length(ds)
print(expected)