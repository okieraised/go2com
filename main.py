from pydicom import dcmread

fpath = "./test_data/File 12948.dcm"
fpath = "./test_data/File 1.dcm"
fpath = "./test_data/File 18001"
fpath = "./test_data/File 32000"
fpath = "./test_data/File 11636.dcm"
fpath = "./test_data/File 32000"
fpath = "./test_data/File 4000.dcm"
fpath = "./test_data/File 12000"
fpath = "./test_data/File 160.dcm"
fpath = "./test_data/File 8000"
fpath = "./test_data/File 12000"
fpath = '/home/tripg/Documents/1.dcm'

ds = dcmread(fpath)
print(ds)