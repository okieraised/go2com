from pydicom import dcmread

fpath = "./test_data/File 2.dcm"

ds = dcmread(fpath)
print(ds)