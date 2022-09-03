from pydicom import dcmread

fpath = "./Test_data/File 12.dcm"

ds = dcmread(fpath)
print(ds)