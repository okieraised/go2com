# go2com (DICOM and NIFTI image parser)

** This package is under active development and can be in a very broken state. Please use the latest released version **

## TODO
- [ ] Improve NIfTI reader for large file size
- [ ] Improve NIfTI writer to export as NIfTI-2 format
- [ ] Improve DICOM frame parser
- [ ] Support for additional medical image formats
- [ ] Support changing value in NIfTI frame

## Example
To parse a DICOM file
```go

```

To parse a single NIfTI file:
```go

```

## Supported Transfer Syntaxes
```text
ImplicitVRLittleEndian                                     "1.2.840.10008.1.2"
ExplicitVRLittleEndian                                     "1.2.840.10008.1.2.1"
ExplicitVRBigEndian                                        "1.2.840.10008.1.2.2"
DeflatedExplicitVRLittleEndian                             "1.2.840.10008.1.2.1.99"
JPEGBaselineProcess1                                       "1.2.840.10008.1.2.4.50"
JPEGBaselineProcess2And4                                   "1.2.840.10008.1.2.4.51"
JPEGLosslessNonHierarchicalProcesses14                     "1.2.840.10008.1.2.4.57"
JPEGLosslessNonHierarchicalFirstOrderPredictionProcess14   "1.2.840.10008.1.2.4.70"
JPEGLSLosslessImageCompression                             "1.2.840.10008.1.2.4.80"
JPEGLSLossyNearLosslessImageCompression                    "1.2.840.10008.1.2.4.81"
JPEG2000ImageCompressionLosslessOnly                       "1.2.840.10008.1.2.4.90"
JPEG2000ImageCompression                                   "1.2.840.10008.1.2.4.91"
MPEG4AVCH264highProfile                                    "1.2.840.10008.1.2.4.102"
MPEG4AVCH264BDCompatibleHighProfile                        "1.2.840.10008.1.2.4.103"
```

## Benchmark Result
Using 2 cores

#### Parsing result without skipping the PixelData option
```shell
goos: linux
goarch: amd64
pkg: github.com/okieraised/go2com
cpu: Intel(R) Core(TM) i9-9900K CPU @ 3.60GHz
BenchmarkNewParser-2         246           4898625 ns/op
BenchmarkNewParser-2         222           5023251 ns/op
BenchmarkNewParser-2         240           4900945 ns/op
BenchmarkNewParser-2         241           4957440 ns/op
BenchmarkNewParser-2         238           4911706 ns/op
BenchmarkNewParser-2         241           5001905 ns/op
BenchmarkNewParser-2         240           5062585 ns/op
BenchmarkNewParser-2         238           4964253 ns/op
BenchmarkNewParser-2         240           4904890 ns/op
BenchmarkNewParser-2         237           4882366 ns/op
BenchmarkNewParser-2         238           4971147 ns/op
BenchmarkNewParser-2         240           5134180 ns/op
BenchmarkNewParser-2         246           4930679 ns/op
BenchmarkNewParser-2         232           4971505 ns/op
BenchmarkNewParser-2         241           5002399 ns/op
BenchmarkNewParser-2         242           4868371 ns/op
BenchmarkNewParser-2         241           4897682 ns/op
BenchmarkNewParser-2         246           5014649 ns/op
BenchmarkNewParser-2         236           4989994 ns/op
BenchmarkNewParser-2         236           5238820 ns/op
```

#### Parsing result with skipping the PixelData option
```shell
goos: linux
goarch: amd64
pkg: github.com/okieraised/go2com
cpu: Intel(R) Core(TM) i9-9900K CPU @ 3.60GHz
BenchmarkNewParser-2         248           4977196 ns/op
BenchmarkNewParser-2         241           4799562 ns/op
BenchmarkNewParser-2         250           4847411 ns/op
BenchmarkNewParser-2         237           4818379 ns/op
BenchmarkNewParser-2         223           4738318 ns/op
BenchmarkNewParser-2         252           4687902 ns/op
BenchmarkNewParser-2         242           4769615 ns/op
BenchmarkNewParser-2         249           4670712 ns/op
BenchmarkNewParser-2         254           4817976 ns/op
BenchmarkNewParser-2         244           4680175 ns/op
BenchmarkNewParser-2         259           4749270 ns/op
BenchmarkNewParser-2         237           4674772 ns/op
BenchmarkNewParser-2         255           4621554 ns/op
BenchmarkNewParser-2         229           5062034 ns/op
BenchmarkNewParser-2         252           4587018 ns/op
BenchmarkNewParser-2         253           4697444 ns/op
BenchmarkNewParser-2         261           4802300 ns/op
BenchmarkNewParser-2         259           4789614 ns/op
BenchmarkNewParser-2         244           5045803 ns/op
BenchmarkNewParser-2         232           4986901 ns/op
```