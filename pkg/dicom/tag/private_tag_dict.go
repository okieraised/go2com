package tag

// Credits to the DCMTK team
var PrivateTagDict = `
#
#  Copyright (C) 1994-2020, OFFIS e.V.
#  All rights reserved.  See COPYRIGHT file for details.
#
#  This software and supporting documentation were developed by
#
#    OFFIS e.V.
#    R&D Division Health
#    Escherweg 2
#    D-26121 Oldenburg, Germany
#
#
#  Module:  dcmdata
#
#  Author:  Andrew Hewett, Marco Eichelberg, Joerg Riesmeier
#
#  Purpose:
#  This is the private tag DICOM data dictionary for the dcmtk class library.
#
#
# Dictionary of Private Tags
#
#  This dictionary contains the private tags defined in the following
#  reference documents (in alphabetical order):
#   - AGFA IMPAX 6.5.x Solution conformance statement
#   - Circle Cardiovascular Imaging cmr42 3.0 conformance statement
#   - David Clunie's dicom3tools package, 2002-04-20 snapshot
#   - Fuji CR console, 3rd release
#   - GE Vivid S70 version 201 conformance statement
#   - Intelerad Medical Systems Inc., Image Server
#   - OCULUS Pentacam 1.17 conformance statement
#   - Philips Digital Diagnost 1.3 conformance statement
#   - Philips Integris H, catheterization laboratory, RIS-interface
#   - Philips Intera Achieva conformance statement
#   - Philips MR Achieva conformance statement
#   - Siemens Somatom syngo VA40B conformance statement
#   - Siemens AXIOM Artis VB30 conformance statement
#   - SonoWand Invite 2.1.1 conformance statement
#   - Swissvision TR4000 conformance statement
#   - private tags for DCMTK anonymizer tool
#
# Each line represents an entry in the data dictionary.  Each line
# has 5 fields (Tag, VR, Name, VM, Version).  Entries need not be
# in ascending tag order.
#
# Entries may override existing entries.
#
# Each field must be separated by a single tab.
# The tag value may take one of two forms:
#   (gggg,"CREATOR",ee)
#   (gggg,"CREATOR",eeee) [eeee >= 1000]
# The first form describes a private tag that may be used with different
# element numbers as reserved by the private creator element.
# The second form describes a private tag that may only occur with a
# certain fixed element number.
# In both cases, the tag values must be in hexadecimal.
# Repeating groups are represented by indicating the range
# (gggg-o-gggg,"CREATOR",ee) or (gggg-o-gggg,"CREATOR",eeee)
# where "-o-" indicates that only odd group numbers match the definition.
# The element part of the tag can also be a range.
#
# Comments have a '#' at the beginning of the line.
#
# Tag				VR	Name			VM	Version / Description
#
(0019,"1.2.840.113681",10)	ST	CRImageParamsCommon	1	PrivateTag
(0019,"1.2.840.113681",11)	ST	CRImageIPParamsSingle	1	PrivateTag
(0019,"1.2.840.113681",12)	ST	CRImageIPParamsLeft	1	PrivateTag
(0019,"1.2.840.113681",13)	ST	CRImageIPParamsRight	1	PrivateTag

(0087,"1.2.840.113708.794.1.1.2.0",10)	CS	MediaType	1	PrivateTag
(0087,"1.2.840.113708.794.1.1.2.0",20)	CS	MediaLocation	1	PrivateTag
(0087,"1.2.840.113708.794.1.1.2.0",50)	IS	EstimatedRetrieveTime	1	PrivateTag

(0009,"ACUSON",00)	IS	Unknown	1	PrivateTag
(0009,"ACUSON",01)	IS	Unknown	1	PrivateTag
(0009,"ACUSON",02)	UN	Unknown	1	PrivateTag
(0009,"ACUSON",03)	UN	Unknown	1	PrivateTag
(0009,"ACUSON",04)	UN	Unknown	1	PrivateTag
(0009,"ACUSON",05)	UN	Unknown	1	PrivateTag
(0009,"ACUSON",06)	UN	Unknown	1	PrivateTag
(0009,"ACUSON",07)	UN	Unknown	1	PrivateTag
(0009,"ACUSON",08)	LT	Unknown	1	PrivateTag
(0009,"ACUSON",09)	LT	Unknown	1	PrivateTag
(0009,"ACUSON",0a)	IS	Unknown	1	PrivateTag
(0009,"ACUSON",0b)	IS	Unknown	1	PrivateTag
(0009,"ACUSON",0c)	IS	Unknown	1	PrivateTag
(0009,"ACUSON",0d)	IS	Unknown	1	PrivateTag
(0009,"ACUSON",0e)	IS	Unknown	1	PrivateTag
(0009,"ACUSON",0f)	UN	Unknown	1	PrivateTag
(0009,"ACUSON",10)	IS	Unknown	1	PrivateTag
(0009,"ACUSON",11)	UN	Unknown	1	PrivateTag
(0009,"ACUSON",12)	IS	Unknown	1	PrivateTag
(0009,"ACUSON",13)	IS	Unknown	1	PrivateTag
(0009,"ACUSON",14)	LT	Unknown	1	PrivateTag
(0009,"ACUSON",15)	UN	Unknown	1	PrivateTag

(0003,"AEGIS_DICOM_2.00",00)	US	Unknown	1-n	PrivateTag
(0005,"AEGIS_DICOM_2.00",00)	US	Unknown	1-n	PrivateTag
(0009,"AEGIS_DICOM_2.00",00)	US	Unknown	1-n	PrivateTag
(0019,"AEGIS_DICOM_2.00",00)	US	Unknown	1-n	PrivateTag
(0029,"AEGIS_DICOM_2.00",00)	US	Unknown	1-n	PrivateTag
(1369,"AEGIS_DICOM_2.00",00)	US	Unknown	1-n	PrivateTag

(0009,"AGFA",10)	LO	Unknown	1	PrivateTag
(0009,"AGFA",11)	LO	Unknown	1	PrivateTag
(0009,"AGFA",13)	LO	Unknown	1	PrivateTag
(0009,"AGFA",14)	LO	Unknown	1	PrivateTag
(0009,"AGFA",15)	LO	Unknown	1	PrivateTag

(0031,"AGFA PACS Archive Mirroring 1.0",00)	CS	StudyStatus	1	PrivateTag
(0031,"AGFA PACS Archive Mirroring 1.0",01)	UL	DateTimeVerified	1	PrivateTag

(0029,"CAMTRONICS IP",10)	LT	Unknown	1	PrivateTag
(0029,"CAMTRONICS IP",20)	UN	Unknown	1	PrivateTag
(0029,"CAMTRONICS IP",30)	UN	Unknown	1	PrivateTag
(0029,"CAMTRONICS IP",40)	UN	Unknown	1	PrivateTag

(0029,"CAMTRONICS",10)	LT	Commentline	1	PrivateTag
(0029,"CAMTRONICS",20)	DS	EdgeEnhancementCoefficient	1	PrivateTag
(0029,"CAMTRONICS",50)	LT	SceneText	1	PrivateTag
(0029,"CAMTRONICS",60)	LT	ImageText	1	PrivateTag
(0029,"CAMTRONICS",70)	IS	PixelShiftHorizontal	1	PrivateTag
(0029,"CAMTRONICS",80)	IS	PixelShiftVertical	1	PrivateTag
(0029,"CAMTRONICS",90)	IS	Unknown	1	PrivateTag

(0009,"CARDIO-D.R. 1.0",00)	UL	FileLocation	1	PrivateTag
(0009,"CARDIO-D.R. 1.0",01)	UL	FileSize	1	PrivateTag
(0009,"CARDIO-D.R. 1.0",40)	SQ	AlternateImageSequence	1	PrivateTag
(0019,"CARDIO-D.R. 1.0",00)	CS	ImageBlankingShape	1	PrivateTag
(0019,"CARDIO-D.R. 1.0",02)	IS	ImageBlankingLeftVerticalEdge	1	PrivateTag
(0019,"CARDIO-D.R. 1.0",04)	IS	ImageBlankingRightVerticalEdge	1	PrivateTag
(0019,"CARDIO-D.R. 1.0",06)	IS	ImageBlankingUpperHorizontalEdge	1	PrivateTag
(0019,"CARDIO-D.R. 1.0",08)	IS	ImageBlankingLowerHorizontalEdge	1	PrivateTag
(0019,"CARDIO-D.R. 1.0",10)	IS	CenterOfCircularImageBlanking	1	PrivateTag
(0019,"CARDIO-D.R. 1.0",12)	IS	RadiusOfCircularImageBlanking	1	PrivateTag
(0019,"CARDIO-D.R. 1.0",30)	UL	MaximumImageFrameSize	1	PrivateTag
(0021,"CARDIO-D.R. 1.0",13)	IS	ImageSequenceNumber	1	PrivateTag
(0029,"CARDIO-D.R. 1.0",00)	SQ	EdgeEnhancementSequence	1	PrivateTag
(0029,"CARDIO-D.R. 1.0",01)	US	ConvolutionKernelSize	2	PrivateTag
(0029,"CARDIO-D.R. 1.0",02)	DS	ConvolutionKernelCoefficients	1-n	PrivateTag
(0029,"CARDIO-D.R. 1.0",03)	DS	EdgeEnhancementGain	1	PrivateTag

(0025,"CMR42 CIRCLECVI",1010)	LO	WorkspaceID	1	PrivateTag
(0025,"CMR42 CIRCLECVI",1020)	LO	WorkspaceTimeString	1	PrivateTag
(0025,"CMR42 CIRCLECVI",1030)	OB	WorkspaceStream	1	PrivateTag

(0009,"DCMTK_ANONYMIZER",00)	SQ	AnonymizerUIDMap	1	PrivateTag
(0009,"DCMTK_ANONYMIZER",10)	UI	AnonymizerUIDKey	1	PrivateTag
(0009,"DCMTK_ANONYMIZER",20)	UI	AnonymizerUIDValue	1	PrivateTag
(0009,"DCMTK_ANONYMIZER",30)	SQ	AnonymizerPatientIDMap	1	PrivateTag
(0009,"DCMTK_ANONYMIZER",40)	LO	AnonymizerPatientIDKey	1	PrivateTag
(0009,"DCMTK_ANONYMIZER",50)	LO	AnonymizerPatientIDValue	1	PrivateTag

(0019,"DIDI TO PCR 1.1",22)	UN	RouteAET	1	PrivateTag
(0019,"DIDI TO PCR 1.1",23)	DS	PCRPrintScale	1	PrivateTag
(0019,"DIDI TO PCR 1.1",24)	UN	PCRPrintJobEnd	1	PrivateTag
(0019,"DIDI TO PCR 1.1",25)	IS	PCRNoFilmCopies	1	PrivateTag
(0019,"DIDI TO PCR 1.1",26)	IS	PCRFilmLayoutPosition	1	PrivateTag
(0019,"DIDI TO PCR 1.1",27)	UN	PCRPrintReportName	1	PrivateTag
(0019,"DIDI TO PCR 1.1",70)	UN	RADProtocolPrinter	1	PrivateTag
(0019,"DIDI TO PCR 1.1",71)	UN	RADProtocolMedium	1	PrivateTag
(0019,"DIDI TO PCR 1.1",90)	LO	UnprocessedFlag	1	PrivateTag
(0019,"DIDI TO PCR 1.1",91)	UN	KeyValues	1	PrivateTag
(0019,"DIDI TO PCR 1.1",92)	UN	DestinationPostprocessingFunction	1	PrivateTag
(0019,"DIDI TO PCR 1.1",A0)	UN	Version	1	PrivateTag
(0019,"DIDI TO PCR 1.1",A1)	UN	RangingMode	1	PrivateTag
(0019,"DIDI TO PCR 1.1",A2)	UN	AbdomenBrightness	1	PrivateTag
(0019,"DIDI TO PCR 1.1",A3)	UN	FixedBrightness	1	PrivateTag
(0019,"DIDI TO PCR 1.1",A4)	UN	DetailContrast	1	PrivateTag
(0019,"DIDI TO PCR 1.1",A5)	UN	ContrastBalance	1	PrivateTag
(0019,"DIDI TO PCR 1.1",A6)	UN	StructureBoost	1	PrivateTag
(0019,"DIDI TO PCR 1.1",A7)	UN	StructurePreference	1	PrivateTag
(0019,"DIDI TO PCR 1.1",A8)	UN	NoiseRobustness	1	PrivateTag
(0019,"DIDI TO PCR 1.1",A9)	UN	NoiseDoseLimit	1	PrivateTag
(0019,"DIDI TO PCR 1.1",AA)	UN	NoiseDoseStep	1	PrivateTag
(0019,"DIDI TO PCR 1.1",AB)	UN	NoiseFrequencyLimit	1	PrivateTag
(0019,"DIDI TO PCR 1.1",AC)	UN	WeakContrastLimit	1	PrivateTag
(0019,"DIDI TO PCR 1.1",AD)	UN	StrongContrastLimit	1	PrivateTag
(0019,"DIDI TO PCR 1.1",AE)	UN	StructureBoostOffset	1	PrivateTag
(0019,"DIDI TO PCR 1.1",AF)	UN	SmoothGain	1	PrivateTag
(0019,"DIDI TO PCR 1.1",B0)	UN	MeasureField1	1	PrivateTag
(0019,"DIDI TO PCR 1.1",B1)	UN	MeasureField2	1	PrivateTag
(0019,"DIDI TO PCR 1.1",B2)	UN	KeyPercentile1	1	PrivateTag
(0019,"DIDI TO PCR 1.1",B3)	UN	KeyPercentile2	1	PrivateTag
(0019,"DIDI TO PCR 1.1",B4)	UN	DensityLUT	1	PrivateTag
(0019,"DIDI TO PCR 1.1",B5)	UN	Brightness	1	PrivateTag
(0019,"DIDI TO PCR 1.1",B6)	UN	Gamma	1	PrivateTag
(0089,"DIDI TO PCR 1.1",10)	SQ	Unknown	1	PrivateTag

(0029,"DIGISCAN IMAGE",31)	US	Unknown	1-n	PrivateTag
(0029,"DIGISCAN IMAGE",32)	US	Unknown	1-n	PrivateTag
(0029,"DIGISCAN IMAGE",33)	LT	Unknown	1	PrivateTag
(0029,"DIGISCAN IMAGE",34)	LT	Unknown	1	PrivateTag

(7001-o-70ff,"DLX_ANNOT_01",04)	ST	TextAnnotation	1	PrivateTag
(7001-o-70ff,"DLX_ANNOT_01",05)	IS	Box	2	PrivateTag
(7001-o-70ff,"DLX_ANNOT_01",07)	IS	ArrowEnd	2	PrivateTag

(0015,"DLX_EXAMS_01",01)	DS	StenosisCalibrationRatio	1	PrivateTag
(0015,"DLX_EXAMS_01",02)	DS	StenosisMagnification	1	PrivateTag
(0015,"DLX_EXAMS_01",03)	DS	CardiacCalibrationRatio	1	PrivateTag

(6001-o-60ff,"DLX_LKUP_01",01)	US	GrayPaletteColorLookupTableDescriptor	3	PrivateTag
(6001-o-60ff,"DLX_LKUP_01",02)	US	GrayPaletteColorLookupTableData	1	PrivateTag

(0011,"DLX_PATNT_01",01)	LT	PatientDOB	1	PrivateTag

(0019,"DLX_SERIE_01",01)	DS	AngleValueLArm	1	PrivateTag
(0019,"DLX_SERIE_01",02)	DS	AngleValuePArm	1	PrivateTag
(0019,"DLX_SERIE_01",03)	DS	AngleValueCArm	1	PrivateTag
(0019,"DLX_SERIE_01",04)	CS	AngleLabelLArm	1	PrivateTag
(0019,"DLX_SERIE_01",05)	CS	AngleLabelPArm	1	PrivateTag
(0019,"DLX_SERIE_01",06)	CS	AngleLabelCArm	1	PrivateTag
(0019,"DLX_SERIE_01",07)	ST	ProcedureName	1	PrivateTag
(0019,"DLX_SERIE_01",08)	ST	ExamName	1	PrivateTag
(0019,"DLX_SERIE_01",09)	SH	PatientSize	1	PrivateTag
(0019,"DLX_SERIE_01",0a)	IS	RecordView	1	PrivateTag
(0019,"DLX_SERIE_01",10)	DS	InjectorDelay	1	PrivateTag
(0019,"DLX_SERIE_01",11)	CS	AutoInject	1	PrivateTag
(0019,"DLX_SERIE_01",14)	IS	AcquisitionMode	1	PrivateTag
(0019,"DLX_SERIE_01",15)	CS	CameraRotationEnabled	1	PrivateTag
(0019,"DLX_SERIE_01",16)	CS	ReverseSweep	1	PrivateTag
(0019,"DLX_SERIE_01",17)	IS	SpatialFilterStrength	1	PrivateTag
(0019,"DLX_SERIE_01",18)	IS	ZoomFactor	1	PrivateTag
(0019,"DLX_SERIE_01",19)	IS	XZoomCenter	1	PrivateTag
(0019,"DLX_SERIE_01",1a)	IS	YZoomCenter	1	PrivateTag
(0019,"DLX_SERIE_01",1b)	DS	Focus	1	PrivateTag
(0019,"DLX_SERIE_01",1c)	CS	Dose	1	PrivateTag
(0019,"DLX_SERIE_01",1d)	IS	SideMark	1	PrivateTag
(0019,"DLX_SERIE_01",1e)	IS	PercentageLandscape	1	PrivateTag
(0019,"DLX_SERIE_01",1f)	DS	ExposureDuration	1	PrivateTag

(00E1,"ELSCINT1",01)	US	DataDictionaryVersion	1	PrivateTag
(00E1,"ELSCINT1",14)	LT	Unknown	1	PrivateTag
(00E1,"ELSCINT1",22)	DS	Unknown	2	PrivateTag
(00E1,"ELSCINT1",23)	DS	Unknown	2	PrivateTag
(00E1,"ELSCINT1",24)	LT	Unknown	1	PrivateTag
(00E1,"ELSCINT1",25)	LT	Unknown	1	PrivateTag
(00E1,"ELSCINT1",40)	SH	OffsetFromCTMRImages	1	PrivateTag
(0601,"ELSCINT1",00)	SH	ImplementationVersion	1	PrivateTag
(0601,"ELSCINT1",20)	DS	RelativeTablePosition	1	PrivateTag
(0601,"ELSCINT1",21)	DS	RelativeTableHeight	1	PrivateTag
(0601,"ELSCINT1",30)	SH	SurviewDirection	1	PrivateTag
(0601,"ELSCINT1",31)	DS	SurviewLength	1	PrivateTag
(0601,"ELSCINT1",50)	SH	ImageViewType	1	PrivateTag
(0601,"ELSCINT1",70)	DS	BatchNumber	1	PrivateTag
(0601,"ELSCINT1",71)	DS	BatchSize	1	PrivateTag
(0601,"ELSCINT1",72)	DS	BatchSliceNumber	1	PrivateTag

(0009,"FDMS 1.0",04)	SH	ImageControlUnit	1	PrivateTag
(0009,"FDMS 1.0",05)	OW	ImageUID	1	PrivateTag
(0009,"FDMS 1.0",06)	OW	RouteImageUID	1	PrivateTag
(0009,"FDMS 1.0",08)	UL	ImageDisplayInformationVersionNo	1	PrivateTag
(0009,"FDMS 1.0",09)	UL	PatientInformationVersionNo	1	PrivateTag
(0009,"FDMS 1.0",0C)	OW	FilmUID	1	PrivateTag
(0009,"FDMS 1.0",10)	CS	ExposureUnitTypeCode	1	PrivateTag
(0009,"FDMS 1.0",80)	LO	KanjiHospitalName	1	PrivateTag
(0009,"FDMS 1.0",90)	ST	DistributionCode	1	PrivateTag
(0009,"FDMS 1.0",92)	SH	KanjiDepartmentName	1	PrivateTag
(0009,"FDMS 1.0",F0)	CS	BlackeningProcessFlag	1	PrivateTag
(0019,"FDMS 1.0",15)	LO	KanjiBodyPartForExposure	1	PrivateTag
(0019,"FDMS 1.0",32)	LO	KanjiMenuName	1	PrivateTag
(0019,"FDMS 1.0",40)	CS	ImageProcessingType	1	PrivateTag
(0019,"FDMS 1.0",50)	CS	EDRMode	1	PrivateTag
(0019,"FDMS 1.0",60)	SH	RadiographersCode	1	PrivateTag
(0019,"FDMS 1.0",70)	IS	SplitExposureFormat	1	PrivateTag
(0019,"FDMS 1.0",71)	IS	NoOfSplitExposureFrames	1	PrivateTag
(0019,"FDMS 1.0",80)	IS	ReadingPositionSpecification	1	PrivateTag
(0019,"FDMS 1.0",81)	IS	ReadingSensitivityCenter	1	PrivateTag
(0019,"FDMS 1.0",90)	SH	FilmAnnotationCharacterString1	1	PrivateTag
(0019,"FDMS 1.0",91)	SH	FilmAnnotationCharacterString2	1	PrivateTag
(0021,"FDMS 1.0",10)	CS	FCRImageID	1	PrivateTag
(0021,"FDMS 1.0",30)	CS	SetNo	1	PrivateTag
(0021,"FDMS 1.0",40)	IS	ImageNoInTheSet	1	PrivateTag
(0021,"FDMS 1.0",50)	CS	PairProcessingInformation	1	PrivateTag
(0021,"FDMS 1.0",80)	OB	EquipmentTypeSpecificInformation	1	PrivateTag
(0023,"FDMS 1.0",10)	SQ	Unknown	1	PrivateTag
(0023,"FDMS 1.0",20)	SQ	Unknown	1	PrivateTag
(0023,"FDMS 1.0",30)	SQ	Unknown	1	PrivateTag
(0025,"FDMS 1.0",10)	US	RelativeLightEmissionAmountSk	1	PrivateTag
(0025,"FDMS 1.0",11)	US	TermOfCorrectionForEachIPTypeSt	1	PrivateTag
(0025,"FDMS 1.0",12)	US	ReadingGainGp	1	PrivateTag
(0025,"FDMS 1.0",13)	US	Unknown	1	PrivateTag
(0025,"FDMS 1.0",15)	CS	Unknown	1	PrivateTag
(0025,"FDMS 1.0",20)	US	Unknown	2	PrivateTag
(0025,"FDMS 1.0",21)	US	Unknown	1	PrivateTag
(0025,"FDMS 1.0",30)	US	Unknown	1	PrivateTag
(0025,"FDMS 1.0",31)	SS	Unknown	1	PrivateTag
(0025,"FDMS 1.0",32)	US	Unknown	1	PrivateTag
(0025,"FDMS 1.0",33)	SS	Unknown	1	PrivateTag
(0025,"FDMS 1.0",34)	SS	Unknown	1	PrivateTag
(0025,"FDMS 1.0",40)	US	Unknown	1	PrivateTag
(0025,"FDMS 1.0",41)	US	Unknown	1	PrivateTag
(0025,"FDMS 1.0",42)	US	Unknown	1	PrivateTag
(0025,"FDMS 1.0",43)	US	Unknown	1	PrivateTag
(0025,"FDMS 1.0",50)	US	Unknown	1	PrivateTag
(0025,"FDMS 1.0",51)	US	Unknown	1	PrivateTag
(0025,"FDMS 1.0",52)	US	Unknown	1	PrivateTag
(0025,"FDMS 1.0",53)	US	Unknown	1	PrivateTag
(0025,"FDMS 1.0",60)	US	Unknown	1	PrivateTag
(0025,"FDMS 1.0",61)	US	Unknown	1	PrivateTag
(0025,"FDMS 1.0",62)	US	Unknown	1	PrivateTag
(0025,"FDMS 1.0",63)	CS	Unknown	1	PrivateTag
(0025,"FDMS 1.0",70)	US	Unknown	1	PrivateTag
(0025,"FDMS 1.0",71)	US	Unknown	1	PrivateTag
(0025,"FDMS 1.0",72)	US	Unknown	1	PrivateTag
(0025,"FDMS 1.0",73)	US	Unknown	1-n	PrivateTag
(0025,"FDMS 1.0",74)	US	Unknown	1-n	PrivateTag
(0025,"FDMS 1.0",80)	US	Unknown	1	PrivateTag
(0025,"FDMS 1.0",81)	US	Unknown	1	PrivateTag
(0025,"FDMS 1.0",82)	US	Unknown	1	PrivateTag
(0025,"FDMS 1.0",83)	US	Unknown	1-n	PrivateTag
(0025,"FDMS 1.0",84)	US	Unknown	1-n	PrivateTag
(0025,"FDMS 1.0",90)	US	Unknown	1	PrivateTag
(0025,"FDMS 1.0",91)	US	Unknown	1	PrivateTag
(0025,"FDMS 1.0",92)	US	Unknown	1	PrivateTag
(0025,"FDMS 1.0",93)	US	Unknown	1	PrivateTag
(0025,"FDMS 1.0",94)	US	Unknown	1	PrivateTag
(0025,"FDMS 1.0",95)	US	Unknown	1	PrivateTag
(0025,"FDMS 1.0",96)	CS	Unknown	1	PrivateTag
(0025,"FDMS 1.0",a0)	US	Unknown	1	PrivateTag
(0025,"FDMS 1.0",a1)	SS	Unknown	1	PrivateTag
(0025,"FDMS 1.0",a2)	US	Unknown	1	PrivateTag
(0025,"FDMS 1.0",a3)	SS	Unknown	1	PrivateTag
(0027,"FDMS 1.0",10)	SQ	Unknown	1	PrivateTag
(0027,"FDMS 1.0",20)	SQ	Unknown	1	PrivateTag
(0027,"FDMS 1.0",30)	SQ	Unknown	1	PrivateTag
(0027,"FDMS 1.0",40)	SQ	Unknown	1	PrivateTag
(0027,"FDMS 1.0",50)	SQ	Unknown	1	PrivateTag
(0027,"FDMS 1.0",60)	SQ	Unknown	1	PrivateTag
(0027,"FDMS 1.0",70)	SQ	Unknown	1	PrivateTag
(0027,"FDMS 1.0",80)	SQ	Unknown	1	PrivateTag
(0027,"FDMS 1.0",a0)	IS	Unknown	1	PrivateTag
(0027,"FDMS 1.0",a1)	CS	Unknown	2	PrivateTag
(0027,"FDMS 1.0",a2)	CS	Unknown	2	PrivateTag
(0027,"FDMS 1.0",a3)	SS	Unknown	1-n	PrivateTag
(0029,"FDMS 1.0",20)	CS	ImageScanningDirection	1	PrivateTag
(0029,"FDMS 1.0",30)	CS	ExtendedReadingSizeValue	1	PrivateTag
(0029,"FDMS 1.0",34)	US	MagnificationReductionRatio	1	PrivateTag
(0029,"FDMS 1.0",44)	CS	LineDensityCode	1	PrivateTag
(0029,"FDMS 1.0",50)	CS	DataCompressionCode	1	PrivateTag
(2011,"FDMS 1.0",11)	CS	ImagePosition SpecifyingFlag	1	PrivateTag
(50F1,"FDMS 1.0",06)	CS	EnergySubtractionParam	1	PrivateTag
(50F1,"FDMS 1.0",07)	CS	SubtractionRegistrationResult	1	PrivateTag
(50F1,"FDMS 1.0",08)	CS	EnergySubtractionParam2	1	PrivateTag
(50F1,"FDMS 1.0",09)	SL	AfinConversionCoefficient	1	PrivateTag
(50F1,"FDMS 1.0",10)	CS	FilmOutputFormat	1	PrivateTag
(50F1,"FDMS 1.0",20)	CS	ImageProcessingModificationFlag	1	PrivateTag

(0009,"FFP DATA",01)	UN	CRHeaderInformation	1	PrivateTag

(0019,"GE ??? From Adantage Review CS",30)	LO	CREDRMode	1	PrivateTag
(0019,"GE ??? From Adantage Review CS",40)	LO	CRLatitude	1	PrivateTag
(0019,"GE ??? From Adantage Review CS",50)	LO	CRGroupNumber	1	PrivateTag
(0019,"GE ??? From Adantage Review CS",70)	LO	CRImageSerialNumber	1	PrivateTag
(0019,"GE ??? From Adantage Review CS",80)	LO	CRBarCodeNumber	1	PrivateTag
(0019,"GE ??? From Adantage Review CS",90)	LO	CRFilmOutputExposures	1	PrivateTag

(0009,"GEMS_ACQU_01",24)	DS	Unknown	1	PrivateTag
(0009,"GEMS_ACQU_01",25)	US	Unknown	1	PrivateTag
(0009,"GEMS_ACQU_01",3e)	US	Unknown	1	PrivateTag
(0009,"GEMS_ACQU_01",3f)	US	Unknown	1	PrivateTag
(0009,"GEMS_ACQU_01",42)	US	Unknown	1	PrivateTag
(0009,"GEMS_ACQU_01",43)	US	Unknown	1	PrivateTag
(0009,"GEMS_ACQU_01",f8)	US	Unknown	1	PrivateTag
(0009,"GEMS_ACQU_01",fb)	IS	Unknown	1	PrivateTag
(0019,"GEMS_ACQU_01",01)	LT	Unknown	1	PrivateTag
(0019,"GEMS_ACQU_01",02)	SL	NumberOfCellsInDetector	1	PrivateTag
(0019,"GEMS_ACQU_01",03)	DS	CellNumberAtTheta	1	PrivateTag
(0019,"GEMS_ACQU_01",04)	DS	CellSpacing	1	PrivateTag
(0019,"GEMS_ACQU_01",05)	LT	Unknown	1	PrivateTag
(0019,"GEMS_ACQU_01",06)	UN	Unknown	1	PrivateTag
(0019,"GEMS_ACQU_01",0e)	US	Unknown	1	PrivateTag
(0019,"GEMS_ACQU_01",0f)	DS	HorizontalFrameOfReference	1	PrivateTag
(0019,"GEMS_ACQU_01",11)	SS	SeriesContrast	1	PrivateTag
(0019,"GEMS_ACQU_01",12)	SS	LastPseq	1	PrivateTag
(0019,"GEMS_ACQU_01",13)	SS	StartNumberForBaseline	1	PrivateTag
(0019,"GEMS_ACQU_01",14)	SS	End NumberForBaseline	1	PrivateTag
(0019,"GEMS_ACQU_01",15)	SS	StartNumberForEnhancedScans	1	PrivateTag
(0019,"GEMS_ACQU_01",16)	SS	EndNumberForEnhancedScans	1	PrivateTag
(0019,"GEMS_ACQU_01",17)	SS	SeriesPlane	1	PrivateTag
(0019,"GEMS_ACQU_01",18)	LO	FirstScanRAS	1	PrivateTag
(0019,"GEMS_ACQU_01",19)	DS	FirstScanLocation	1	PrivateTag
(0019,"GEMS_ACQU_01",1a)	LO	LastScanRAS	1	PrivateTag
(0019,"GEMS_ACQU_01",1b)	DS	LastScanLocation	1	PrivateTag
(0019,"GEMS_ACQU_01",1e)	DS	DisplayFieldOfView	1	PrivateTag
(0019,"GEMS_ACQU_01",20)	DS	Unknown	1	PrivateTag
(0019,"GEMS_ACQU_01",22)	DS	Unknown	1	PrivateTag
(0019,"GEMS_ACQU_01",23)	DS	TableSpeed	1	PrivateTag
(0019,"GEMS_ACQU_01",24)	DS	MidScanTime	1	PrivateTag
(0019,"GEMS_ACQU_01",25)	SS	MidScanFlag	1	PrivateTag
(0019,"GEMS_ACQU_01",26)	SL	DegreesOfAzimuth	1	PrivateTag
(0019,"GEMS_ACQU_01",27)	DS	GantryPeriod	1	PrivateTag
(0019,"GEMS_ACQU_01",2a)	DS	XrayOnPosition	1	PrivateTag
(0019,"GEMS_ACQU_01",2b)	DS	XrayOffPosition	1	PrivateTag
(0019,"GEMS_ACQU_01",2c)	SL	NumberOfTriggers	1	PrivateTag
(0019,"GEMS_ACQU_01",2d)	US	Unknown	1	PrivateTag
(0019,"GEMS_ACQU_01",2e)	DS	AngleOfFirstView	1	PrivateTag
(0019,"GEMS_ACQU_01",2f)	DS	TriggerFrequency	1	PrivateTag
(0019,"GEMS_ACQU_01",39)	SS	ScanFOVType	1	PrivateTag
(0019,"GEMS_ACQU_01",3a)	IS	Unknown	1	PrivateTag
(0019,"GEMS_ACQU_01",3b)	LT	Unknown	1	PrivateTag
(0019,"GEMS_ACQU_01",3c)	UN	Unknown	1	PrivateTag
(0019,"GEMS_ACQU_01",3e)	UN	Unknown	1	PrivateTag
(0019,"GEMS_ACQU_01",3f)	UN	Unknown	1	PrivateTag
(0019,"GEMS_ACQU_01",40)	SS	StatReconFlag	1	PrivateTag
(0019,"GEMS_ACQU_01",41)	SS	ComputeType	1	PrivateTag
(0019,"GEMS_ACQU_01",42)	SS	SegmentNumber	1	PrivateTag
(0019,"GEMS_ACQU_01",43)	SS	TotalSegmentsRequested	1	PrivateTag
(0019,"GEMS_ACQU_01",44)	DS	InterscanDelay	1	PrivateTag
(0019,"GEMS_ACQU_01",47)	SS	ViewCompressionFactor	1	PrivateTag
(0019,"GEMS_ACQU_01",48)	US	Unknown	1	PrivateTag
(0019,"GEMS_ACQU_01",49)	US	Unknown	1	PrivateTag
(0019,"GEMS_ACQU_01",4a)	SS	TotalNumberOfRefChannels	1	PrivateTag
(0019,"GEMS_ACQU_01",4b)	SL	DataSizeForScanData	1	PrivateTag
(0019,"GEMS_ACQU_01",52)	SS	ReconPostProcessingFlag	1	PrivateTag
(0019,"GEMS_ACQU_01",54)	UN	Unknown	1	PrivateTag
(0019,"GEMS_ACQU_01",57)	SS	CTWaterNumber	1	PrivateTag
(0019,"GEMS_ACQU_01",58)	SS	CTBoneNumber	1	PrivateTag
(0019,"GEMS_ACQU_01",5a)	FL	AcquisitionDuration	1	PrivateTag
(0019,"GEMS_ACQU_01",5d)	US	Unknown	1	PrivateTag
(0019,"GEMS_ACQU_01",5e)	SL	NumberOfChannels1To512	1	PrivateTag
(0019,"GEMS_ACQU_01",5f)	SL	IncrementBetweenChannels	1	PrivateTag
(0019,"GEMS_ACQU_01",60)	SL	StartingView	1	PrivateTag
(0019,"GEMS_ACQU_01",61)	SL	NumberOfViews	1	PrivateTag
(0019,"GEMS_ACQU_01",62)	SL	IncrementBetweenViews	1	PrivateTag
(0019,"GEMS_ACQU_01",6a)	SS	DependantOnNumberOfViewsProcessed	1	PrivateTag
(0019,"GEMS_ACQU_01",6b)	SS	FieldOfViewInDetectorCells	1	PrivateTag
(0019,"GEMS_ACQU_01",70)	SS	ValueOfBackProjectionButton	1	PrivateTag
(0019,"GEMS_ACQU_01",71)	SS	SetIfFatqEstimatesWereUsed	1	PrivateTag
(0019,"GEMS_ACQU_01",72)	DS	ZChannelAvgOverViews	1	PrivateTag
(0019,"GEMS_ACQU_01",73)	DS	AvgOfLeftRefChannelsOverViews	1	PrivateTag
(0019,"GEMS_ACQU_01",74)	DS	MaxLeftChannelOverViews	1	PrivateTag
(0019,"GEMS_ACQU_01",75)	DS	AvgOfRightRefChannelsOverViews	1	PrivateTag
(0019,"GEMS_ACQU_01",76)	DS	MaxRightChannelOverViews	1	PrivateTag
(0019,"GEMS_ACQU_01",7d)	DS	SecondEcho	1	PrivateTag
(0019,"GEMS_ACQU_01",7e)	SS	NumberOfEchos	1	PrivateTag
(0019,"GEMS_ACQU_01",7f)	DS	TableDelta	1	PrivateTag
(0019,"GEMS_ACQU_01",81)	SS	Contiguous	1	PrivateTag
(0019,"GEMS_ACQU_01",82)	US	Unknown	1	PrivateTag
(0019,"GEMS_ACQU_01",83)	DS	Unknown	1	PrivateTag
(0019,"GEMS_ACQU_01",84)	DS	PeakSAR	1	PrivateTag
(0019,"GEMS_ACQU_01",85)	SS	MonitorSAR	1	PrivateTag
(0019,"GEMS_ACQU_01",86)	US	Unknown	1	PrivateTag
(0019,"GEMS_ACQU_01",87)	DS	CardiacRepetition Time	1	PrivateTag
(0019,"GEMS_ACQU_01",88)	SS	ImagesPerCardiacCycle	1	PrivateTag
(0019,"GEMS_ACQU_01",8a)	SS	ActualReceiveGainAnalog	1	PrivateTag
(0019,"GEMS_ACQU_01",8b)	SS	ActualReceiveGainDigital	1	PrivateTag
(0019,"GEMS_ACQU_01",8d)	DS	DelayAfterTrigger	1	PrivateTag
(0019,"GEMS_ACQU_01",8f)	SS	SwapPhaseFrequency	1	PrivateTag
(0019,"GEMS_ACQU_01",90)	SS	PauseInterval	1	PrivateTag
(0019,"GEMS_ACQU_01",91)	DS	PulseTime	1	PrivateTag
(0019,"GEMS_ACQU_01",92)	SL	SliceOffsetOnFrequencyAxis	1	PrivateTag
(0019,"GEMS_ACQU_01",93)	DS	CenterFrequency	1	PrivateTag
(0019,"GEMS_ACQU_01",94)	SS	TransmitGain	1	PrivateTag
(0019,"GEMS_ACQU_01",95)	SS	AnalogReceiverGain	1	PrivateTag
(0019,"GEMS_ACQU_01",96)	SS	DigitalReceiverGain	1	PrivateTag
(0019,"GEMS_ACQU_01",97)	SL	BitmapDefiningCVs	1	PrivateTag
(0019,"GEMS_ACQU_01",98)	SS	CenterFrequencyMethod	1	PrivateTag
(0019,"GEMS_ACQU_01",99)	US	Unknown	1	PrivateTag
(0019,"GEMS_ACQU_01",9b)	SS	PulseSequenceMode	1	PrivateTag
(0019,"GEMS_ACQU_01",9c)	LO	PulseSequenceName	1	PrivateTag
(0019,"GEMS_ACQU_01",9d)	DT	PulseSequenceDate	1	PrivateTag
(0019,"GEMS_ACQU_01",9e)	LO	InternalPulseSequenceName	1	PrivateTag
(0019,"GEMS_ACQU_01",9f)	SS	TransmittingCoil	1	PrivateTag
(0019,"GEMS_ACQU_01",a0)	SS	SurfaceCoilType	1	PrivateTag
(0019,"GEMS_ACQU_01",a1)	SS	ExtremityCoilFlag	1	PrivateTag
(0019,"GEMS_ACQU_01",a2)	SL	RawDataRunNumber	1	PrivateTag
(0019,"GEMS_ACQU_01",a3)	UL	CalibratedFieldStrength	1	PrivateTag
(0019,"GEMS_ACQU_01",a4)	SS	SATFatWaterBone	1	PrivateTag
(0019,"GEMS_ACQU_01",a5)	DS	ReceiveBandwidth	1	PrivateTag
(0019,"GEMS_ACQU_01",a7)	DS	UserData0	1	PrivateTag
(0019,"GEMS_ACQU_01",a8)	DS	UserData1	1	PrivateTag
(0019,"GEMS_ACQU_01",a9)	DS	UserData2	1	PrivateTag
(0019,"GEMS_ACQU_01",aa)	DS	UserData3	1	PrivateTag
(0019,"GEMS_ACQU_01",ab)	DS	UserData4	1	PrivateTag
(0019,"GEMS_ACQU_01",ac)	DS	UserData5	1	PrivateTag
(0019,"GEMS_ACQU_01",ad)	DS	UserData6	1	PrivateTag
(0019,"GEMS_ACQU_01",ae)	DS	UserData7	1	PrivateTag
(0019,"GEMS_ACQU_01",af)	DS	UserData8	1	PrivateTag
(0019,"GEMS_ACQU_01",b0)	DS	UserData9	1	PrivateTag
(0019,"GEMS_ACQU_01",b1)	DS	UserData10	1	PrivateTag
(0019,"GEMS_ACQU_01",b2)	DS	UserData11	1	PrivateTag
(0019,"GEMS_ACQU_01",b3)	DS	UserData12	1	PrivateTag
(0019,"GEMS_ACQU_01",b4)	DS	UserData13	1	PrivateTag
(0019,"GEMS_ACQU_01",b5)	DS	UserData14	1	PrivateTag
(0019,"GEMS_ACQU_01",b6)	DS	UserData15	1	PrivateTag
(0019,"GEMS_ACQU_01",b7)	DS	UserData16	1	PrivateTag
(0019,"GEMS_ACQU_01",b8)	DS	UserData17	1	PrivateTag
(0019,"GEMS_ACQU_01",b9)	DS	UserData18	1	PrivateTag
(0019,"GEMS_ACQU_01",ba)	DS	UserData19	1	PrivateTag
(0019,"GEMS_ACQU_01",bb)	DS	UserData20	1	PrivateTag
(0019,"GEMS_ACQU_01",bc)	DS	UserData21	1	PrivateTag
(0019,"GEMS_ACQU_01",bd)	DS	UserData22	1	PrivateTag
(0019,"GEMS_ACQU_01",be)	DS	ProjectionAngle	1	PrivateTag
(0019,"GEMS_ACQU_01",c0)	SS	SaturationPlanes	1	PrivateTag
(0019,"GEMS_ACQU_01",c1)	SS	SurfaceCoilIntensityCorrectionFlag	1	PrivateTag
(0019,"GEMS_ACQU_01",c2)	SS	SATLocationR	1	PrivateTag
(0019,"GEMS_ACQU_01",c3)	SS	SATLocationL	1	PrivateTag
(0019,"GEMS_ACQU_01",c4)	SS	SATLocationA	1	PrivateTag
(0019,"GEMS_ACQU_01",c5)	SS	SATLocationP	1	PrivateTag
(0019,"GEMS_ACQU_01",c6)	SS	SATLocationH	1	PrivateTag
(0019,"GEMS_ACQU_01",c7)	SS	SATLocationF	1	PrivateTag
(0019,"GEMS_ACQU_01",c8)	SS	SATThicknessRL	1	PrivateTag
(0019,"GEMS_ACQU_01",c9)	SS	SATThicknessAP	1	PrivateTag
(0019,"GEMS_ACQU_01",ca)	SS	SATThicknessHF	1	PrivateTag
(0019,"GEMS_ACQU_01",cb)	SS	PrescribedFlowAxis	1	PrivateTag
(0019,"GEMS_ACQU_01",cc)	SS	VelocityEncoding	1	PrivateTag
(0019,"GEMS_ACQU_01",cd)	SS	ThicknessDisclaimer	1	PrivateTag
(0019,"GEMS_ACQU_01",ce)	SS	PrescanType	1	PrivateTag
(0019,"GEMS_ACQU_01",cf)	SS	PrescanStatus	1	PrivateTag
(0019,"GEMS_ACQU_01",d0)	SH	RawDataType	1	PrivateTag
(0019,"GEMS_ACQU_01",d2)	SS	ProjectionAlgorithm	1	PrivateTag
(0019,"GEMS_ACQU_01",d3)	SH	ProjectionAlgorithm	1	PrivateTag
(0019,"GEMS_ACQU_01",d4)	US	Unknown	1	PrivateTag
(0019,"GEMS_ACQU_01",d5)	SS	FractionalEcho	1	PrivateTag
(0019,"GEMS_ACQU_01",d6)	SS	PrepPulse	1	PrivateTag
(0019,"GEMS_ACQU_01",d7)	SS	CardiacPhases	1	PrivateTag
(0019,"GEMS_ACQU_01",d8)	SS	VariableEchoFlag	1	PrivateTag
(0019,"GEMS_ACQU_01",d9)	DS	ConcatenatedSAT	1	PrivateTag
(0019,"GEMS_ACQU_01",da)	SS	ReferenceChannelUsed	1	PrivateTag
(0019,"GEMS_ACQU_01",db)	DS	BackProjectorCoefficient	1	PrivateTag
(0019,"GEMS_ACQU_01",dc)	SS	PrimarySpeedCorrectionUsed	1	PrivateTag
(0019,"GEMS_ACQU_01",dd)	SS	OverrangeCorrectionUsed	1	PrivateTag
(0019,"GEMS_ACQU_01",de)	DS	DynamicZAlphaValue	1	PrivateTag
(0019,"GEMS_ACQU_01",df)	DS	UserData23	1	PrivateTag
(0019,"GEMS_ACQU_01",e0)	DS	UserData24	1	PrivateTag
(0019,"GEMS_ACQU_01",e1)	DS	Unknown	1	PrivateTag
(0019,"GEMS_ACQU_01",e2)	DS	VelocityEncodeScale	1	PrivateTag
(0019,"GEMS_ACQU_01",e3)	LT	Unknown	1	PrivateTag
(0019,"GEMS_ACQU_01",e4)	LT	Unknown	1	PrivateTag
(0019,"GEMS_ACQU_01",e5)	IS	Unknown	1	PrivateTag
(0019,"GEMS_ACQU_01",e6)	US	Unknown	1	PrivateTag
(0019,"GEMS_ACQU_01",e8)	DS	Unknown	1	PrivateTag
(0019,"GEMS_ACQU_01",e9)	DS	Unknown	1	PrivateTag
(0019,"GEMS_ACQU_01",eb)	DS	Unknown	1	PrivateTag
(0019,"GEMS_ACQU_01",ec)	US	Unknown	1	PrivateTag
(0019,"GEMS_ACQU_01",f0)	UN	Unknown	1	PrivateTag
(0019,"GEMS_ACQU_01",f1)	LT	Unknown	1	PrivateTag
(0019,"GEMS_ACQU_01",f2)	SS	FastPhases	1	PrivateTag
(0019,"GEMS_ACQU_01",f3)	LT	Unknown	1	PrivateTag
(0019,"GEMS_ACQU_01",f4)	LT	Unknown	1	PrivateTag
(0019,"GEMS_ACQU_01",f9)	DS	TransmitGain	1	PrivateTag

(0023,"GEMS_ACRQA_1.0 BLOCK1",00)	LO	CRExposureMenuCode	1	PrivateTag
(0023,"GEMS_ACRQA_1.0 BLOCK1",10)	LO	CRExposureMenuString	1	PrivateTag
(0023,"GEMS_ACRQA_1.0 BLOCK1",20)	LO	CREDRMode	1	PrivateTag
(0023,"GEMS_ACRQA_1.0 BLOCK1",30)	LO	CRLatitude	1	PrivateTag
(0023,"GEMS_ACRQA_1.0 BLOCK1",40)	LO	CRGroupNumber	1	PrivateTag
(0023,"GEMS_ACRQA_1.0 BLOCK1",50)	US	CRImageSerialNumber	1	PrivateTag
(0023,"GEMS_ACRQA_1.0 BLOCK1",60)	LO	CRBarCodeNumber	1	PrivateTag
(0023,"GEMS_ACRQA_1.0 BLOCK1",70)	LO	CRFilmOutputExposure	1	PrivateTag
(0023,"GEMS_ACRQA_1.0 BLOCK1",80)	LO	CRFilmFormat	1	PrivateTag
(0023,"GEMS_ACRQA_1.0 BLOCK1",90)	LO	CRSShiftString	1	PrivateTag
(0023,"GEMS_ACRQA_1.0 BLOCK2",00)	US	CRSShift	1	PrivateTag
(0023,"GEMS_ACRQA_1.0 BLOCK2",10)	DS	CRCShift	1	PrivateTag
(0023,"GEMS_ACRQA_1.0 BLOCK2",20)	DS	CRGT	1	PrivateTag
(0023,"GEMS_ACRQA_1.0 BLOCK2",30)	DS	CRGA	1	PrivateTag
(0023,"GEMS_ACRQA_1.0 BLOCK2",40)	DS	CRGC	1	PrivateTag
(0023,"GEMS_ACRQA_1.0 BLOCK2",50)	DS	CRGS	1	PrivateTag
(0023,"GEMS_ACRQA_1.0 BLOCK2",60)	DS	CRRT	1	PrivateTag
(0023,"GEMS_ACRQA_1.0 BLOCK2",70)	DS	CRRE	1	PrivateTag
(0023,"GEMS_ACRQA_1.0 BLOCK2",80)	US	CRRN	1	PrivateTag
(0023,"GEMS_ACRQA_1.0 BLOCK2",90)	DS	CRDRT	1	PrivateTag
(0023,"GEMS_ACRQA_1.0 BLOCK3",00)	DS	CRDRE	1	PrivateTag
(0023,"GEMS_ACRQA_1.0 BLOCK3",10)	US	CRDRN	1	PrivateTag
(0023,"GEMS_ACRQA_1.0 BLOCK3",20)	DS	CRORE	1	PrivateTag
(0023,"GEMS_ACRQA_1.0 BLOCK3",30)	US	CRORN	1	PrivateTag
(0023,"GEMS_ACRQA_1.0 BLOCK3",40)	US	CRORD	1	PrivateTag
(0023,"GEMS_ACRQA_1.0 BLOCK3",50)	LO	CRCassetteSize	1	PrivateTag
(0023,"GEMS_ACRQA_1.0 BLOCK3",60)	LO	CRMachineID	1	PrivateTag
(0023,"GEMS_ACRQA_1.0 BLOCK3",70)	LO	CRMachineType	1	PrivateTag
(0023,"GEMS_ACRQA_1.0 BLOCK3",80)	LO	CRTechnicianCode	1	PrivateTag
(0023,"GEMS_ACRQA_1.0 BLOCK3",90)	LO	CREnergySubtractionParameters	1	PrivateTag
(0023,"GEMS_ACRQA_2.0 BLOCK1",00)	LO	CRExposureMenuCode	1	PrivateTag
(0023,"GEMS_ACRQA_2.0 BLOCK1",10)	LO	CRExposureMenuString	1	PrivateTag
(0023,"GEMS_ACRQA_2.0 BLOCK1",20)	LO	CREDRMode	1	PrivateTag
(0023,"GEMS_ACRQA_2.0 BLOCK1",30)	LO	CRLatitude	1	PrivateTag
(0023,"GEMS_ACRQA_2.0 BLOCK1",40)	LO	CRGroupNumber	1	PrivateTag
(0023,"GEMS_ACRQA_2.0 BLOCK1",50)	US	CRImageSerialNumber	1	PrivateTag
(0023,"GEMS_ACRQA_2.0 BLOCK1",60)	LO	CRBarCodeNumber	1	PrivateTag
(0023,"GEMS_ACRQA_2.0 BLOCK1",70)	LO	CRFilmOutputExposure	1	PrivateTag
(0023,"GEMS_ACRQA_2.0 BLOCK1",80)	LO	CRFilmFormat	1	PrivateTag
(0023,"GEMS_ACRQA_2.0 BLOCK1",90)	LO	CRSShiftString	1	PrivateTag
(0023,"GEMS_ACRQA_2.0 BLOCK2",00)	US	CRSShift	1	PrivateTag
(0023,"GEMS_ACRQA_2.0 BLOCK2",10)	LO	CRCShift	1	PrivateTag
(0023,"GEMS_ACRQA_2.0 BLOCK2",20)	LO	CRGT	1	PrivateTag
(0023,"GEMS_ACRQA_2.0 BLOCK2",30)	DS	CRGA	1	PrivateTag
(0023,"GEMS_ACRQA_2.0 BLOCK2",40)	DS	CRGC	1	PrivateTag
(0023,"GEMS_ACRQA_2.0 BLOCK2",50)	DS	CRGS	1	PrivateTag
(0023,"GEMS_ACRQA_2.0 BLOCK2",60)	LO	CRRT	1	PrivateTag
(0023,"GEMS_ACRQA_2.0 BLOCK2",70)	DS	CRRE	1	PrivateTag
(0023,"GEMS_ACRQA_2.0 BLOCK2",80)	US	CRRN	1	PrivateTag
(0023,"GEMS_ACRQA_2.0 BLOCK2",90)	DS	CRDRT	1	PrivateTag
(0023,"GEMS_ACRQA_2.0 BLOCK3",00)	DS	CRDRE	1	PrivateTag
(0023,"GEMS_ACRQA_2.0 BLOCK3",10)	US	CRDRN	1	PrivateTag
(0023,"GEMS_ACRQA_2.0 BLOCK3",20)	DS	CRORE	1	PrivateTag
(0023,"GEMS_ACRQA_2.0 BLOCK3",30)	US	CRORN	1	PrivateTag
(0023,"GEMS_ACRQA_2.0 BLOCK3",40)	US	CRORD	1	PrivateTag
(0023,"GEMS_ACRQA_2.0 BLOCK3",50)	LO	CRCassetteSize	1	PrivateTag
(0023,"GEMS_ACRQA_2.0 BLOCK3",60)	LO	CRMachineID	1	PrivateTag
(0023,"GEMS_ACRQA_2.0 BLOCK3",70)	LO	CRMachineType	1	PrivateTag
(0023,"GEMS_ACRQA_2.0 BLOCK3",80)	LO	CRTechnicianCode	1	PrivateTag
(0023,"GEMS_ACRQA_2.0 BLOCK3",90)	LO	CREnergySubtractionParameters	1	PrivateTag
(0023,"GEMS_ACRQA_2.0 BLOCK3",f0)	LO	CRDistributionCode	1	PrivateTag
(0023,"GEMS_ACRQA_2.0 BLOCK3",ff)	US	CRShuttersApplied	1	PrivateTag

(0047,"GEMS_ADWSoft_3D1",01)	SQ	Reconstruction Parameters Sequence	1	PrivateTag
(0047,"GEMS_ADWSoft_3D1",50)	UL	VolumeVoxelCount	1	PrivateTag
(0047,"GEMS_ADWSoft_3D1",51)	UL	VolumeSegmentCount	1	PrivateTag
(0047,"GEMS_ADWSoft_3D1",53)	US	VolumeSliceSize	1	PrivateTag
(0047,"GEMS_ADWSoft_3D1",54)	US	VolumeSliceCount	1	PrivateTag
(0047,"GEMS_ADWSoft_3D1",55)	SL	VolumeThresholdValue	1	PrivateTag
(0047,"GEMS_ADWSoft_3D1",57)	DS	VolumeVoxelRatio	1	PrivateTag
(0047,"GEMS_ADWSoft_3D1",58)	DS	VolumeVoxelSize	1	PrivateTag
(0047,"GEMS_ADWSoft_3D1",59)	US	VolumeZPositionSize	1	PrivateTag
(0047,"GEMS_ADWSoft_3D1",60)	DS	VolumeBaseLine	9	PrivateTag
(0047,"GEMS_ADWSoft_3D1",61)	DS	VolumeCenterPoint	3	PrivateTag
(0047,"GEMS_ADWSoft_3D1",63)	SL	VolumeSkewBase	1	PrivateTag
(0047,"GEMS_ADWSoft_3D1",64)	DS	VolumeRegistrationTransformRotationMatrix	9	PrivateTag
(0047,"GEMS_ADWSoft_3D1",65)	DS	VolumeRegistrationTransformTranslationVector	3	PrivateTag
(0047,"GEMS_ADWSoft_3D1",70)	DS	KVPList	1-n	PrivateTag
(0047,"GEMS_ADWSoft_3D1",71)	IS	XRayTubeCurrentList	1-n	PrivateTag
(0047,"GEMS_ADWSoft_3D1",72)	IS	ExposureList	1-n	PrivateTag
(0047,"GEMS_ADWSoft_3D1",80)	LO	AcquisitionDLXIdentifier	1	PrivateTag
(0047,"GEMS_ADWSoft_3D1",85)	SQ	AcquisitionDLX2DSeriesSequence	1	PrivateTag
(0047,"GEMS_ADWSoft_3D1",89)	DS	ContrastAgentVolumeList	1-n	PrivateTag
(0047,"GEMS_ADWSoft_3D1",8A)	US	NumberOfInjections	1	PrivateTag
(0047,"GEMS_ADWSoft_3D1",8B)	US	FrameCount	1	PrivateTag
(0047,"GEMS_ADWSoft_3D1",91)	LO	XA3DReconstructionAlgorithmName	1	PrivateTag
(0047,"GEMS_ADWSoft_3D1",92)	CS	XA3DReconstructionAlgorithmVersion	1	PrivateTag
(0047,"GEMS_ADWSoft_3D1",93)	DA	DLXCalibrationDate	1	PrivateTag
(0047,"GEMS_ADWSoft_3D1",94)	TM	DLXCalibrationTime	1	PrivateTag
(0047,"GEMS_ADWSoft_3D1",95)	CS	DLXCalibrationStatus	1	PrivateTag
(0047,"GEMS_ADWSoft_3D1",96)	IS	UsedFrames	1-n	PrivateTag
(0047,"GEMS_ADWSoft_3D1",98)	US	TransformCount	1	PrivateTag
(0047,"GEMS_ADWSoft_3D1",99)	SQ	TransformSequence	1	PrivateTag
(0047,"GEMS_ADWSoft_3D1",9A)	DS	TransformRotationMatrix	9	PrivateTag
(0047,"GEMS_ADWSoft_3D1",9B)	DS	TransformTranslationVector	3	PrivateTag
(0047,"GEMS_ADWSoft_3D1",9C)	LO	TransformLabel	1	PrivateTag
(0047,"GEMS_ADWSoft_3D1",B0)	SQ	WireframeList	1	PrivateTag
(0047,"GEMS_ADWSoft_3D1",B1)	US	WireframeCount	1	PrivateTag
(0047,"GEMS_ADWSoft_3D1",B2)	US	LocationSystem	1	PrivateTag
(0047,"GEMS_ADWSoft_3D1",B5)	LO	WireframeName	1	PrivateTag
(0047,"GEMS_ADWSoft_3D1",B6)	LO	WireframeGroupName	1	PrivateTag
(0047,"GEMS_ADWSoft_3D1",B7)	LO	WireframeColor	1	PrivateTag
(0047,"GEMS_ADWSoft_3D1",B8)	SL	WireframeAttributes	1	PrivateTag
(0047,"GEMS_ADWSoft_3D1",B9)	SL	WireframePointCount	1	PrivateTag
(0047,"GEMS_ADWSoft_3D1",BA)	SL	WireframeTimestamp	1	PrivateTag
(0047,"GEMS_ADWSoft_3D1",BB)	SQ	WireframePointList	1	PrivateTag
(0047,"GEMS_ADWSoft_3D1",BC)	DS	WireframePointsCoordinates	3	PrivateTag
(0047,"GEMS_ADWSoft_3D1",C0)	DS	VolumeUpperLeftHighCornerRAS	3	PrivateTag
(0047,"GEMS_ADWSoft_3D1",C1)	DS	VolumeSliceToRASRotationMatrix	9	PrivateTag
(0047,"GEMS_ADWSoft_3D1",C2)	DS	VolumeUpperLeftHighCornerTLOC	1	PrivateTag
(0047,"GEMS_ADWSoft_3D1",D1)	OB	VolumeSegmentList	1	PrivateTag
(0047,"GEMS_ADWSoft_3D1",D2)	OB	VolumeGradientList	1	PrivateTag
(0047,"GEMS_ADWSoft_3D1",D3)	OB	VolumeDensityList	1	PrivateTag
(0047,"GEMS_ADWSoft_3D1",D4)	OB	VolumeZPositionList	1	PrivateTag
(0047,"GEMS_ADWSoft_3D1",D5)	OB	VolumeOriginalIndexList	1	PrivateTag
(0039,"GEMS_ADWSoft_DPO",80)	IS	PrivateEntityNumber	1	PrivateTag
(0039,"GEMS_ADWSoft_DPO",85)	DA	PrivateEntityDate	1	PrivateTag
(0039,"GEMS_ADWSoft_DPO",90)	TM	PrivateEntityTime	1	PrivateTag
(0039,"GEMS_ADWSoft_DPO",95)	LO	PrivateEntityLaunchCommand	1	PrivateTag
(0039,"GEMS_ADWSoft_DPO",AA)	CS	PrivateEntityType	1	PrivateTag

(0033,"GEMS_CTHD_01",02)	UN	Unknown	1	PrivateTag

(0037,"GEMS_DRS_1",10)	LO	ReferringDepartment	1	PrivateTag
(0037,"GEMS_DRS_1",20)	US	ScreenNumber	1	PrivateTag
(0037,"GEMS_DRS_1",40)	SH	LeftOrientation	1	PrivateTag
(0037,"GEMS_DRS_1",42)	SH	RightOrientation	1	PrivateTag
(0037,"GEMS_DRS_1",50)	CS	Inversion	1	PrivateTag
(0037,"GEMS_DRS_1",60)	US	DSA	1	PrivateTag

(0009,"GEMS_GENIE_1",10)	LO	Unknown	1	PrivateTag
(0009,"GEMS_GENIE_1",11)	SL	StudyFlags	1	PrivateTag
(0009,"GEMS_GENIE_1",12)	SL	StudyType	1	PrivateTag
(0009,"GEMS_GENIE_1",1e)	UI	Unknown	1	PrivateTag
(0009,"GEMS_GENIE_1",20)	LO	Unknown	1	PrivateTag
(0009,"GEMS_GENIE_1",21)	SL	SeriesFlags	1	PrivateTag
(0009,"GEMS_GENIE_1",22)	SH	UserOrientation	1	PrivateTag
(0009,"GEMS_GENIE_1",23)	SL	InitiationType	1	PrivateTag
(0009,"GEMS_GENIE_1",24)	SL	InitiationDelay	1	PrivateTag
(0009,"GEMS_GENIE_1",25)	SL	InitiationCountRate	1	PrivateTag
(0009,"GEMS_GENIE_1",26)	SL	NumberEnergySets	1	PrivateTag
(0009,"GEMS_GENIE_1",27)	SL	NumberDetectors	1	PrivateTag
(0009,"GEMS_GENIE_1",29)	SL	Unknown	1	PrivateTag
(0009,"GEMS_GENIE_1",2a)	SL	Unknown	1	PrivateTag
(0009,"GEMS_GENIE_1",2c)	LO	SeriesComments	1	PrivateTag
(0009,"GEMS_GENIE_1",2d)	SL	TrackBeatAverage	1	PrivateTag
(0009,"GEMS_GENIE_1",2e)	FD	DistancePrescribed	1	PrivateTag
(0009,"GEMS_GENIE_1",30)	LO	Unknown	1	PrivateTag
(0009,"GEMS_GENIE_1",35)	SL	GantryLocusType	1	PrivateTag
(0009,"GEMS_GENIE_1",37)	SL	StartingHeartRate	1	PrivateTag
(0009,"GEMS_GENIE_1",38)	SL	RRWindowWidth	1	PrivateTag
(0009,"GEMS_GENIE_1",39)	SL	RRWindowOffset	1	PrivateTag
(0009,"GEMS_GENIE_1",3a)	SL	PercentCycleImaged	1	PrivateTag
(0009,"GEMS_GENIE_1",40)	LO	Unknown	1	PrivateTag
(0009,"GEMS_GENIE_1",41)	SL	PatientFlags	1	PrivateTag
(0009,"GEMS_GENIE_1",42)	DA	PatientCreationDate	1	PrivateTag
(0009,"GEMS_GENIE_1",43)	TM	PatientCreationTime	1	PrivateTag
(0011,"GEMS_GENIE_1",0a)	SL	SeriesType	1	PrivateTag
(0011,"GEMS_GENIE_1",0b)	SL	EffectiveSeriesDuration	1	PrivateTag
(0011,"GEMS_GENIE_1",0c)	SL	NumBeats	1	PrivateTag
(0011,"GEMS_GENIE_1",0d)	LO	RadioNuclideName	1	PrivateTag
(0011,"GEMS_GENIE_1",10)	LO	Unknown	1	PrivateTag
(0011,"GEMS_GENIE_1",12)	LO	DatasetName	1	PrivateTag
(0011,"GEMS_GENIE_1",13)	SL	DatasetType	1	PrivateTag
(0011,"GEMS_GENIE_1",15)	SL	DetectorNumber	1	PrivateTag
(0011,"GEMS_GENIE_1",16)	SL	EnergyNumber	1	PrivateTag
(0011,"GEMS_GENIE_1",17)	SL	RRIntervalWindowNumber	1	PrivateTag
(0011,"GEMS_GENIE_1",18)	SL	MGBinNumber	1	PrivateTag
(0011,"GEMS_GENIE_1",19)	FD	RadiusOfRotation	1	PrivateTag
(0011,"GEMS_GENIE_1",1a)	SL	DetectorCountZone	1	PrivateTag
(0011,"GEMS_GENIE_1",1b)	SL	NumEnergyWindows	1	PrivateTag
(0011,"GEMS_GENIE_1",1c)	SL	EnergyOffset	4	PrivateTag
(0011,"GEMS_GENIE_1",1d)	SL	EnergyRange	1	PrivateTag
(0011,"GEMS_GENIE_1",1f)	SL	ImageOrientation	1	PrivateTag
(0011,"GEMS_GENIE_1",23)	SL	UseFOVMask	1	PrivateTag
(0011,"GEMS_GENIE_1",24)	SL	FOVMaskYCutoffAngle	1	PrivateTag
(0011,"GEMS_GENIE_1",25)	SL	FOVMaskCutoffAngle	1	PrivateTag
(0011,"GEMS_GENIE_1",26)	SL	TableOrientation	1	PrivateTag
(0011,"GEMS_GENIE_1",27)	SL	ROITopLeft	2	PrivateTag
(0011,"GEMS_GENIE_1",28)	SL	ROIBottomRight	2	PrivateTag
(0011,"GEMS_GENIE_1",30)	LO	Unknown	1	PrivateTag
(0011,"GEMS_GENIE_1",33)	LO	EnergyCorrectName	1	PrivateTag
(0011,"GEMS_GENIE_1",34)	LO	SpatialCorrectName	1	PrivateTag
(0011,"GEMS_GENIE_1",35)	LO	TuningCalibName	1	PrivateTag
(0011,"GEMS_GENIE_1",36)	LO	UniformityCorrectName	1	PrivateTag
(0011,"GEMS_GENIE_1",37)	LO	AcquisitionSpecificCorrectName	1	PrivateTag
(0011,"GEMS_GENIE_1",38)	SL	ByteOrder	1	PrivateTag
(0011,"GEMS_GENIE_1",3a)	SL	PictureFormat	1	PrivateTag
(0011,"GEMS_GENIE_1",3b)	FD	PixelScale	1	PrivateTag
(0011,"GEMS_GENIE_1",3c)	FD	PixelOffset	1	PrivateTag
(0011,"GEMS_GENIE_1",3e)	SL	FOVShape	1	PrivateTag
(0011,"GEMS_GENIE_1",3f)	SL	DatasetFlags	1	PrivateTag
(0011,"GEMS_GENIE_1",44)	FD	ThresholdCenter	1	PrivateTag
(0011,"GEMS_GENIE_1",45)	FD	ThresholdWidth	1	PrivateTag
(0011,"GEMS_GENIE_1",46)	SL	InterpolationType	1	PrivateTag
(0011,"GEMS_GENIE_1",55)	FD	Period	1	PrivateTag
(0011,"GEMS_GENIE_1",56)	FD	ElapsedTime	1	PrivateTag
(0013,"GEMS_GENIE_1",10)	FD	DigitalFOV	2	PrivateTag
(0013,"GEMS_GENIE_1",11)	SL	Unknown	1	PrivateTag
(0013,"GEMS_GENIE_1",12)	SL	Unknown	1	PrivateTag
(0013,"GEMS_GENIE_1",16)	SL	AutoTrackPeak	1	PrivateTag
(0013,"GEMS_GENIE_1",17)	SL	AutoTrackWidth	1	PrivateTag
(0013,"GEMS_GENIE_1",18)	FD	TransmissionScanTime	1	PrivateTag
(0013,"GEMS_GENIE_1",19)	FD	TransmissionMaskWidth	1	PrivateTag
(0013,"GEMS_GENIE_1",1a)	FD	CopperAttenuatorThickness	1	PrivateTag
(0013,"GEMS_GENIE_1",1c)	FD	Unknown	1	PrivateTag
(0013,"GEMS_GENIE_1",1d)	FD	Unknown	1	PrivateTag
(0013,"GEMS_GENIE_1",1e)	FD	TomoViewOffset	1-n	PrivateTag
(0013,"GEMS_GENIE_1",26)	LT	StudyComments	1	PrivateTag

(0033,"GEMS_GNHD_01",01)	UN	Unknown	1	PrivateTag
(0033,"GEMS_GNHD_01",02)	UN	Unknown	1	PrivateTag

(0009,"GEMS_IDEN_01",01)	LO	FullFidelity	1	PrivateTag
(0009,"GEMS_IDEN_01",02)	SH	SuiteId	1	PrivateTag
(0009,"GEMS_IDEN_01",04)	SH	ProductId	1	PrivateTag
(0009,"GEMS_IDEN_01",17)	LT	Unknown	1	PrivateTag
(0009,"GEMS_IDEN_01",1a)	US	Unknown	1	PrivateTag
(0009,"GEMS_IDEN_01",20)	US	Unknown	1	PrivateTag
(0009,"GEMS_IDEN_01",27)	SL	ImageActualDate	1	PrivateTag
(0009,"GEMS_IDEN_01",2f)	LT	Unknown	1	PrivateTag
(0009,"GEMS_IDEN_01",30)	SH	ServiceId	1	PrivateTag
(0009,"GEMS_IDEN_01",31)	SH	MobileLocationNumber	1	PrivateTag
(0009,"GEMS_IDEN_01",e2)	LT	Unknown	1	PrivateTag
(0009,"GEMS_IDEN_01",e3)	UI	EquipmentUID	1	PrivateTag
(0009,"GEMS_IDEN_01",e6)	SH	GenesisVersionNow	1	PrivateTag
(0009,"GEMS_IDEN_01",e7)	UL	ExamRecordChecksum	1	PrivateTag
(0009,"GEMS_IDEN_01",e8)	UL	Unknown	1	PrivateTag
(0009,"GEMS_IDEN_01",e9)	SL	ActualSeriesDataTimeStamp	1	PrivateTag

(0027,"GEMS_IMAG_01",06)	SL	ImageArchiveFlag	1	PrivateTag
(0027,"GEMS_IMAG_01",10)	SS	ScoutType	1	PrivateTag
(0027,"GEMS_IMAG_01",1c)	SL	VmaMamp	1	PrivateTag
(0027,"GEMS_IMAG_01",1d)	SS	VmaPhase	1	PrivateTag
(0027,"GEMS_IMAG_01",1e)	SL	VmaMod	1	PrivateTag
(0027,"GEMS_IMAG_01",1f)	SL	VmaClip	1	PrivateTag
(0027,"GEMS_IMAG_01",20)	SS	SmartScanOnOffFlag	1	PrivateTag
(0027,"GEMS_IMAG_01",30)	SH	ForeignImageRevision	1	PrivateTag
(0027,"GEMS_IMAG_01",31)	SS	ImagingMode	1	PrivateTag
(0027,"GEMS_IMAG_01",32)	SS	PulseSequence	1	PrivateTag
(0027,"GEMS_IMAG_01",33)	SL	ImagingOptions	1	PrivateTag
(0027,"GEMS_IMAG_01",35)	SS	PlaneType	1	PrivateTag
(0027,"GEMS_IMAG_01",36)	SL	ObliquePlane	1	PrivateTag
(0027,"GEMS_IMAG_01",40)	SH	RASLetterOfImageLocation	1	PrivateTag
(0027,"GEMS_IMAG_01",41)	FL	ImageLocation	1	PrivateTag
(0027,"GEMS_IMAG_01",42)	FL	CenterRCoordOfPlaneImage	1	PrivateTag
(0027,"GEMS_IMAG_01",43)	FL	CenterACoordOfPlaneImage	1	PrivateTag
(0027,"GEMS_IMAG_01",44)	FL	CenterSCoordOfPlaneImage	1	PrivateTag
(0027,"GEMS_IMAG_01",45)	FL	NormalRCoord	1	PrivateTag
(0027,"GEMS_IMAG_01",46)	FL	NormalACoord	1	PrivateTag
(0027,"GEMS_IMAG_01",47)	FL	NormalSCoord	1	PrivateTag
(0027,"GEMS_IMAG_01",48)	FL	RCoordOfTopRightCorner	1	PrivateTag
(0027,"GEMS_IMAG_01",49)	FL	ACoordOfTopRightCorner	1	PrivateTag
(0027,"GEMS_IMAG_01",4a)	FL	SCoordOfTopRightCorner	1	PrivateTag
(0027,"GEMS_IMAG_01",4b)	FL	RCoordOfBottomRightCorner	1	PrivateTag
(0027,"GEMS_IMAG_01",4c)	FL	ACoordOfBottomRightCorner	1	PrivateTag
(0027,"GEMS_IMAG_01",4d)	FL	SCoordOfBottomRightCorner	1	PrivateTag
(0027,"GEMS_IMAG_01",50)	FL	TableStartLocation	1	PrivateTag
(0027,"GEMS_IMAG_01",51)	FL	TableEndLocation	1	PrivateTag
(0027,"GEMS_IMAG_01",52)	SH	RASLetterForSideOfImage	1	PrivateTag
(0027,"GEMS_IMAG_01",53)	SH	RASLetterForAnteriorPosterior	1	PrivateTag
(0027,"GEMS_IMAG_01",54)	SH	RASLetterForScoutStartLoc	1	PrivateTag
(0027,"GEMS_IMAG_01",55)	SH	RASLetterForScoutEndLoc	1	PrivateTag
(0027,"GEMS_IMAG_01",60)	FL	ImageDimensionX	1	PrivateTag
(0027,"GEMS_IMAG_01",61)	FL	ImageDimensionY	1	PrivateTag
(0027,"GEMS_IMAG_01",62)	FL	NumberOfExcitations	1	PrivateTag

(0029,"GEMS_IMPS_01",04)	SL	LowerRangeOfPixels	1	PrivateTag
(0029,"GEMS_IMPS_01",05)	DS	LowerRangeOfPixels	1	PrivateTag
(0029,"GEMS_IMPS_01",06)	DS	LowerRangeOfPixels	1	PrivateTag
(0029,"GEMS_IMPS_01",07)	SL	LowerRangeOfPixels	1	PrivateTag
(0029,"GEMS_IMPS_01",08)	SH	LowerRangeOfPixels	1	PrivateTag
(0029,"GEMS_IMPS_01",09)	SH	LowerRangeOfPixels	1	PrivateTag
(0029,"GEMS_IMPS_01",0a)	SS	LowerRangeOfPixels	1	PrivateTag
(0029,"GEMS_IMPS_01",15)	SL	LowerRangeOfPixels	1	PrivateTag
(0029,"GEMS_IMPS_01",16)	SL	LowerRangeOfPixels	1	PrivateTag
(0029,"GEMS_IMPS_01",17)	SL	LowerRangeOfPixels	1	PrivateTag
(0029,"GEMS_IMPS_01",18)	SL	UpperRangeOfPixels	1	PrivateTag
(0029,"GEMS_IMPS_01",1a)	SL	LengthOfTotalHeaderInBytes	1	PrivateTag
(0029,"GEMS_IMPS_01",26)	SS	VersionOfHeaderStructure	1	PrivateTag
(0029,"GEMS_IMPS_01",34)	SL	AdvantageCompOverflow	1	PrivateTag
(0029,"GEMS_IMPS_01",35)	SL	AdvantageCompUnderflow	1	PrivateTag

(0043,"GEMS_PARM_01",01)	SS	BitmapOfPrescanOptions	1	PrivateTag
(0043,"GEMS_PARM_01",02)	SS	GradientOffsetInX	1	PrivateTag
(0043,"GEMS_PARM_01",03)	SS	GradientOffsetInY	1	PrivateTag
(0043,"GEMS_PARM_01",04)	SS	GradientOffsetInZ	1	PrivateTag
(0043,"GEMS_PARM_01",05)	SS	ImageIsOriginalOrUnoriginal	1	PrivateTag
(0043,"GEMS_PARM_01",06)	SS	NumberOfEPIShots	1	PrivateTag
(0043,"GEMS_PARM_01",07)	SS	ViewsPerSegment	1	PrivateTag
(0043,"GEMS_PARM_01",08)	SS	RespiratoryRateInBPM	1	PrivateTag
(0043,"GEMS_PARM_01",09)	SS	RespiratoryTriggerPoint	1	PrivateTag
(0043,"GEMS_PARM_01",0a)	SS	TypeOfReceiverUsed	1	PrivateTag
(0043,"GEMS_PARM_01",0b)	DS	PeakRateOfChangeOfGradientField	1	PrivateTag
(0043,"GEMS_PARM_01",0c)	DS	LimitsInUnitsOfPercent	1	PrivateTag
(0043,"GEMS_PARM_01",0d)	DS	PSDEstimatedLimit	1	PrivateTag
(0043,"GEMS_PARM_01",0e)	DS	PSDEstimatedLimitInTeslaPerSecond	1	PrivateTag
(0043,"GEMS_PARM_01",0f)	DS	SARAvgHead	1	PrivateTag
(0043,"GEMS_PARM_01",10)	US	WindowValue	1	PrivateTag
(0043,"GEMS_PARM_01",11)	US	TotalInputViews	1	PrivateTag
(0043,"GEMS_PARM_01",12)	SS	XrayChain	3	PrivateTag
(0043,"GEMS_PARM_01",13)	SS	ReconKernelParameters	5	PrivateTag
(0043,"GEMS_PARM_01",14)	SS	CalibrationParameters	3	PrivateTag
(0043,"GEMS_PARM_01",15)	SS	TotalOutputViews	3	PrivateTag
(0043,"GEMS_PARM_01",16)	SS	NumberOfOverranges	5	PrivateTag
(0043,"GEMS_PARM_01",17)	DS	IBHImageScaleFactors	1	PrivateTag
(0043,"GEMS_PARM_01",18)	DS	BBHCoefficients	3	PrivateTag
(0043,"GEMS_PARM_01",19)	SS	NumberOfBBHChainsToBlend	1	PrivateTag
(0043,"GEMS_PARM_01",1a)	SL	StartingChannelNumber	1	PrivateTag
(0043,"GEMS_PARM_01",1b)	SS	PPScanParameters	1	PrivateTag
(0043,"GEMS_PARM_01",1c)	SS	GEImageIntegrity	1	PrivateTag
(0043,"GEMS_PARM_01",1d)	SS	LevelValue	1	PrivateTag
(0043,"GEMS_PARM_01",1e)	DS	DeltaStartTime	1	PrivateTag
(0043,"GEMS_PARM_01",1f)	SL	MaxOverrangesInAView	1	PrivateTag
(0043,"GEMS_PARM_01",20)	DS	AvgOverrangesAllViews	1	PrivateTag
(0043,"GEMS_PARM_01",21)	SS	CorrectedAfterglowTerms	1	PrivateTag
(0043,"GEMS_PARM_01",25)	SS	ReferenceChannels	6	PrivateTag
(0043,"GEMS_PARM_01",26)	US	NoViewsRefChannelsBlocked	6	PrivateTag
(0043,"GEMS_PARM_01",27)	SH	ScanPitchRatio	1	PrivateTag
(0043,"GEMS_PARM_01",28)	OB	UniqueImageIdentifier	1	PrivateTag
(0043,"GEMS_PARM_01",29)	OB	HistogramTables	1	PrivateTag
(0043,"GEMS_PARM_01",2a)	OB	UserDefinedData	1	PrivateTag
(0043,"GEMS_PARM_01",2b)	SS	PrivateScanOptions	4	PrivateTag
(0043,"GEMS_PARM_01",2c)	SS	EffectiveEchoSpacing	1	PrivateTag
(0043,"GEMS_PARM_01",2d)	SH	StringSlopField1	1	PrivateTag
(0043,"GEMS_PARM_01",2e)	SH	StringSlopField2	1	PrivateTag
(0043,"GEMS_PARM_01",2f)	SS	RawDataType	1	PrivateTag
(0043,"GEMS_PARM_01",30)	SS	RawDataType	1	PrivateTag
(0043,"GEMS_PARM_01",31)	DS	RACoordOfTargetReconCentre	2	PrivateTag
(0043,"GEMS_PARM_01",32)	SS	RawDataType	1	PrivateTag
(0043,"GEMS_PARM_01",33)	FL	NegScanSpacing	1	PrivateTag
(0043,"GEMS_PARM_01",34)	IS	OffsetFrequency	1	PrivateTag
(0043,"GEMS_PARM_01",35)	UL	UserUsageTag	1	PrivateTag
(0043,"GEMS_PARM_01",36)	UL	UserFillMapMSW	1	PrivateTag
(0043,"GEMS_PARM_01",37)	UL	UserFillMapLSW	1	PrivateTag
(0043,"GEMS_PARM_01",38)	FL	User25ToUser48	24	PrivateTag
(0043,"GEMS_PARM_01",39)	IS	SlopInteger6ToSlopInteger9	4	PrivateTag
(0043,"GEMS_PARM_01",40)	FL	TriggerOnPosition	4	PrivateTag
(0043,"GEMS_PARM_01",41)	FL	DegreeOfRotation	4	PrivateTag
(0043,"GEMS_PARM_01",42)	SL	DASTriggerSource	4	PrivateTag
(0043,"GEMS_PARM_01",43)	SL	DASFpaGain	4	PrivateTag
(0043,"GEMS_PARM_01",44)	SL	DASOutputSource	4	PrivateTag
(0043,"GEMS_PARM_01",45)	SL	DASAdInput	4	PrivateTag
(0043,"GEMS_PARM_01",46)	SL	DASCalMode	4	PrivateTag
(0043,"GEMS_PARM_01",47)	SL	DASCalFrequency	4	PrivateTag
(0043,"GEMS_PARM_01",48)	SL	DASRegXm	4	PrivateTag
(0043,"GEMS_PARM_01",49)	SL	DASAutoZero	4	PrivateTag
(0043,"GEMS_PARM_01",4a)	SS	StartingChannelOfView	4	PrivateTag
(0043,"GEMS_PARM_01",4b)	SL	DASXmPattern	4	PrivateTag
(0043,"GEMS_PARM_01",4c)	SS	TGGCTriggerMode	4	PrivateTag
(0043,"GEMS_PARM_01",4d)	FL	StartScanToXrayOnDelay	4	PrivateTag
(0043,"GEMS_PARM_01",4e)	FL	DurationOfXrayOn	4	PrivateTag
(0043,"GEMS_PARM_01",60)	IS	SlopInteger10ToSlopInteger17	8	PrivateTag
(0043,"GEMS_PARM_01",61)	UI	ScannerStudyEntityUID	1	PrivateTag
(0043,"GEMS_PARM_01",62)	SH	ScannerStudyID	1	PrivateTag
(0043,"GEMS_PARM_01",6f)	DS	ScannerTableEntry	3	PrivateTag
(0043,"GEMS_PARM_01",70)	LO	ParadigmName	1	PrivateTag
(0043,"GEMS_PARM_01",71)	ST	ParadigmDescription	1	PrivateTag
(0043,"GEMS_PARM_01",72)	UI	ParadigmUID	1	PrivateTag
(0043,"GEMS_PARM_01",73)	US	ExperimentType	1	PrivateTag
(0043,"GEMS_PARM_01",74)	US	NumberOfRestVolumes	1	PrivateTag
(0043,"GEMS_PARM_01",75)	US	NumberOfActiveVolumes	1	PrivateTag
(0043,"GEMS_PARM_01",76)	US	NumberOfDummyScans	1	PrivateTag
(0043,"GEMS_PARM_01",77)	SH	ApplicationName	1	PrivateTag
(0043,"GEMS_PARM_01",78)	SH	ApplicationVersion	1	PrivateTag
(0043,"GEMS_PARM_01",79)	US	SlicesPerVolume	1	PrivateTag
(0043,"GEMS_PARM_01",7a)	US	ExpectedTimePoints	1	PrivateTag
(0043,"GEMS_PARM_01",7b)	FL	RegressorValues	1-n	PrivateTag
(0043,"GEMS_PARM_01",7c)	FL	DelayAfterSliceGroup	1	PrivateTag
(0043,"GEMS_PARM_01",7d)	US	ReconModeFlagWord	1	PrivateTag
(0043,"GEMS_PARM_01",7e)	LO	PACCSpecificInformation	1-n	PrivateTag
(0043,"GEMS_PARM_01",7f)	DS	EDWIScaleFactor	1-n	PrivateTag
(0043,"GEMS_PARM_01",80)	LO	CoilIDData	1-n	PrivateTag
(0043,"GEMS_PARM_01",81)	LO	GECoilName	1	PrivateTag
(0043,"GEMS_PARM_01",82)	LO	SystemConfigurationInformation	1-n	PrivateTag
(0043,"GEMS_PARM_01",83)	DS	AssetRFactors	1-2	PrivateTag
(0043,"GEMS_PARM_01",84)	LO	AdditionalAssetData	5-n	PrivateTag
(0043,"GEMS_PARM_01",85)	UT	DebugDataTextFormat	1	PrivateTag
(0043,"GEMS_PARM_01",86)	OB	DebugDataBinaryFormat	1	PrivateTag
(0043,"GEMS_PARM_01",87)	UT	ScannerSoftwareVersionLongForm	1	PrivateTag
(0043,"GEMS_PARM_01",88)	UI	PUREAcquisitionCalibrationSeriesUID	1	PrivateTag
(0043,"GEMS_PARM_01",89)	LO	GoverningBodydBdtAndSARDefinition	3	PrivateTag
(0043,"GEMS_PARM_01",8a)	CS	PrivateInPlanePhaseEncodingDirection	1	PrivateTag
(0043,"GEMS_PARM_01",8b)	OB	FMRIBinaryDataBlock	1	PrivateTag
(0043,"GEMS_PARM_01",8c)	DS	VoxelLocation	6	PrivateTag
(0043,"GEMS_PARM_01",8d)	DS	SATBandLocations	7-7n	PrivateTag
(0043,"GEMS_PARM_01",8e)	DS	SpectroPrescanValues	3	PrivateTag
(0043,"GEMS_PARM_01",8f)	DS	SpectroParameters	3	PrivateTag
(0043,"GEMS_PARM_01",90)	LO	SARDefinition	1-n	PrivateTag
(0043,"GEMS_PARM_01",91)	DS	SARValue	1-n	PrivateTag
(0043,"GEMS_PARM_01",92)	LO	ImageErrorText	1	PrivateTag
(0043,"GEMS_PARM_01",93)	DS	SpectroQuantitationValues	1-n	PrivateTag
(0043,"GEMS_PARM_01",94)	DS	SpectroRatioValues	1-n	PrivateTag
(0043,"GEMS_PARM_01",95)	LO	PrescanReuseString	1	PrivateTag
(0043,"GEMS_PARM_01",96)	CS	ContentQualification	1	PrivateTag
(0043,"GEMS_PARM_01",97)	LO	ImageFilteringParameters	9	PrivateTag
(0043,"GEMS_PARM_01",98)	UI	ASSETAcquisitionCalibrationSeriesUID	1	PrivateTag
(0043,"GEMS_PARM_01",99)	LO	ExtendedOptions	1-n	PrivateTag
(0043,"GEMS_PARM_01",9a)	IS	RxStackIdentification	1	PrivateTag
(0043,"GEMS_PARM_01",9b)	DS	NPWFactor	1	PrivateTag
(0043,"GEMS_PARM_01",9c)	OB	ResearchTag1	1	PrivateTag
(0043,"GEMS_PARM_01",9d)	OB	ResearchTag2	1	PrivateTag
(0043,"GEMS_PARM_01",9e)	OB	ResearchTag3	1	PrivateTag
(0043,"GEMS_PARM_01",9f)	OB	ResearchTag4	1	PrivateTag

(0011,"GEMS_PATI_01",10)	SS	PatientStatus	1	PrivateTag

(0021,"GEMS_RELA_01",03)	SS	SeriesFromWhichPrescribed	1	PrivateTag
(0021,"GEMS_RELA_01",05)	SH	GenesisVersionNow	1	PrivateTag
(0021,"GEMS_RELA_01",07)	UL	SeriesRecordChecksum	1	PrivateTag
(0021,"GEMS_RELA_01",15)	US	Unknown	1	PrivateTag
(0021,"GEMS_RELA_01",16)	SS	Unknown	1	PrivateTag
(0021,"GEMS_RELA_01",18)	SH	GenesisVersionNow	1	PrivateTag
(0021,"GEMS_RELA_01",19)	UL	AcqReconRecordChecksum	1	PrivateTag
(0021,"GEMS_RELA_01",20)	DS	TableStartLocation	1	PrivateTag
(0021,"GEMS_RELA_01",35)	SS	SeriesFromWhichPrescribed	1	PrivateTag
(0021,"GEMS_RELA_01",36)	SS	ImageFromWhichPrescribed	1	PrivateTag
(0021,"GEMS_RELA_01",37)	SS	ScreenFormat	1	PrivateTag
(0021,"GEMS_RELA_01",4a)	LO	AnatomicalReferenceForScout	1	PrivateTag
(0021,"GEMS_RELA_01",4e)	US	Unknown	1	PrivateTag
(0021,"GEMS_RELA_01",4f)	SS	LocationsInAcquisition	1	PrivateTag
(0021,"GEMS_RELA_01",50)	SS	GraphicallyPrescribed	1	PrivateTag
(0021,"GEMS_RELA_01",51)	DS	RotationFromSourceXRot	1	PrivateTag
(0021,"GEMS_RELA_01",52)	DS	RotationFromSourceYRot	1	PrivateTag
(0021,"GEMS_RELA_01",53)	DS	RotationFromSourceZRot	1	PrivateTag
(0021,"GEMS_RELA_01",54)	SH	ImagePosition	3	PrivateTag
(0021,"GEMS_RELA_01",55)	SH	ImageOrientation	6	PrivateTag
(0021,"GEMS_RELA_01",56)	SL	IntegerSlop	1	PrivateTag
(0021,"GEMS_RELA_01",57)	SL	IntegerSlop	1	PrivateTag
(0021,"GEMS_RELA_01",58)	SL	IntegerSlop	1	PrivateTag
(0021,"GEMS_RELA_01",59)	SL	IntegerSlop	1	PrivateTag
(0021,"GEMS_RELA_01",5a)	SL	IntegerSlop	1	PrivateTag
(0021,"GEMS_RELA_01",5b)	DS	FloatSlop	1	PrivateTag
(0021,"GEMS_RELA_01",5c)	DS	FloatSlop	1	PrivateTag
(0021,"GEMS_RELA_01",5d)	DS	FloatSlop	1	PrivateTag
(0021,"GEMS_RELA_01",5e)	DS	FloatSlop	1	PrivateTag
(0021,"GEMS_RELA_01",5f)	DS	FloatSlop	1	PrivateTag
(0021,"GEMS_RELA_01",70)	LT	Unknown	1	PrivateTag
(0021,"GEMS_RELA_01",71)	LT	Unknown	1	PrivateTag
(0021,"GEMS_RELA_01",81)	DS	AutoWindowLevelAlpha	1	PrivateTag
(0021,"GEMS_RELA_01",82)	DS	AutoWindowLevelBeta	1	PrivateTag
(0021,"GEMS_RELA_01",83)	DS	AutoWindowLevelWindow	1	PrivateTag
(0021,"GEMS_RELA_01",84)	DS	AutoWindowLevelLevel	1	PrivateTag
(0021,"GEMS_RELA_01",90)	SS	TubeFocalSpotPosition	1	PrivateTag
(0021,"GEMS_RELA_01",91)	SS	BiopsyPosition	1	PrivateTag
(0021,"GEMS_RELA_01",92)	FL	BiopsyTLocation	1	PrivateTag
(0021,"GEMS_RELA_01",93)	FL	BiopsyRefLocation	1	PrivateTag

(0045,"GEMS_SENO_02",04)	CS	AES	1	PrivateTag
(0045,"GEMS_SENO_02",06)	DS	Angulation	1	PrivateTag
(0045,"GEMS_SENO_02",09)	DS	RealMagnificationFactor	1	PrivateTag
(0045,"GEMS_SENO_02",0b)	CS	SenographType	1	PrivateTag
(0045,"GEMS_SENO_02",0c)	DS	IntegrationTime	1	PrivateTag
(0045,"GEMS_SENO_02",0d)	DS	ROIOriginXY	1	PrivateTag
(0045,"GEMS_SENO_02",11)	DS	ReceptorSizeCmXY	2	PrivateTag
(0045,"GEMS_SENO_02",12)	IS	ReceptorSizePixelsXY	2	PrivateTag
(0045,"GEMS_SENO_02",13)	ST	Screen	1	PrivateTag
(0045,"GEMS_SENO_02",14)	DS	PixelPitchMicrons	1	PrivateTag
(0045,"GEMS_SENO_02",15)	IS	PixelDepthBits	1	PrivateTag
(0045,"GEMS_SENO_02",16)	IS	BinningFactorXY	2	PrivateTag
(0045,"GEMS_SENO_02",1B)	CS	ClinicalView	1	PrivateTag
(0045,"GEMS_SENO_02",1D)	DS	MeanOfRawGrayLevels	1	PrivateTag
(0045,"GEMS_SENO_02",1E)	DS	MeanOfOffsetGrayLevels	1	PrivateTag
(0045,"GEMS_SENO_02",1F)	DS	MeanOfCorrectedGrayLevels	1	PrivateTag
(0045,"GEMS_SENO_02",20)	DS	MeanOfRegionGrayLevels	1	PrivateTag
(0045,"GEMS_SENO_02",21)	DS	MeanOfLogRegionGrayLevels	1	PrivateTag
(0045,"GEMS_SENO_02",22)	DS	StandardDeviationOfRawGrayLevels	1	PrivateTag
(0045,"GEMS_SENO_02",23)	DS	StandardDeviationOfCorrectedGrayLevels	1	PrivateTag
(0045,"GEMS_SENO_02",24)	DS	StandardDeviationOfRegionGrayLevels	1	PrivateTag
(0045,"GEMS_SENO_02",25)	DS	StandardDeviationOfLogRegionGrayLevels	1	PrivateTag
(0045,"GEMS_SENO_02",26)	OB	MAOBuffer	1	PrivateTag
(0045,"GEMS_SENO_02",27)	IS	SetNumber	1	PrivateTag
(0045,"GEMS_SENO_02",28)	CS	WindowingType	1	PrivateTag
(0045,"GEMS_SENO_02",29)	DS	WindowingParameters	1-n	PrivateTag
(0045,"GEMS_SENO_02",2a)	IS	CrosshairCursorXCoordinates	1	PrivateTag
(0045,"GEMS_SENO_02",2b)	IS	CrosshairCursorYCoordinates	1	PrivateTag
(0045,"GEMS_SENO_02",39)	US	VignetteRows	1	PrivateTag
(0045,"GEMS_SENO_02",3a)	US	VignetteColumns	1	PrivateTag
(0045,"GEMS_SENO_02",3b)	US	VignetteBitsAllocated	1	PrivateTag
(0045,"GEMS_SENO_02",3c)	US	VignetteBitsStored	1	PrivateTag
(0045,"GEMS_SENO_02",3d)	US	VignetteHighBit	1	PrivateTag
(0045,"GEMS_SENO_02",3e)	US	VignettePixelRepresentation	1	PrivateTag
(0045,"GEMS_SENO_02",3f)	OB	VignettePixelData	1	PrivateTag

(0025,"GEMS_SERS_01",06)	SS	LastPulseSequenceUsed	1	PrivateTag
(0025,"GEMS_SERS_01",07)	SL	ImagesInSeries	1	PrivateTag
(0025,"GEMS_SERS_01",10)	SL	LandmarkCounter	1	PrivateTag
(0025,"GEMS_SERS_01",11)	SS	NumberOfAcquisitions	1	PrivateTag
(0025,"GEMS_SERS_01",14)	SL	IndicatesNumberOfUpdatesToHeader	1	PrivateTag
(0025,"GEMS_SERS_01",17)	SL	SeriesCompleteFlag	1	PrivateTag
(0025,"GEMS_SERS_01",18)	SL	NumberOfImagesArchived	1	PrivateTag
(0025,"GEMS_SERS_01",19)	SL	LastImageNumberUsed	1	PrivateTag
(0025,"GEMS_SERS_01",1a)	SH	PrimaryReceiverSuiteAndHost	1	PrivateTag

(0023,"GEMS_STDY_01",01)	SL	NumberOfSeriesInStudy	1	PrivateTag
(0023,"GEMS_STDY_01",02)	SL	NumberOfUnarchivedSeries	1	PrivateTag
(0023,"GEMS_STDY_01",10)	SS	ReferenceImageField	1	PrivateTag
(0023,"GEMS_STDY_01",50)	SS	SummaryImage	1	PrivateTag
(0023,"GEMS_STDY_01",70)	FD	StartTimeSecsInFirstAxial	1	PrivateTag
(0023,"GEMS_STDY_01",74)	SL	NumberOfUpdatesToHeader	1	PrivateTag
(0023,"GEMS_STDY_01",7d)	SS	IndicatesIfStudyHasCompleteInfo	1	PrivateTag

(7fe1,"GEMS_Ultrasound_MovieGroup_001",1060)	px	IllegalPrivatePixelSequence	1	PrivateTag

(0033,"GEMS_YMHD_01",05)	UN	Unknown	1	PrivateTag
(0033,"GEMS_YMHD_01",06)	UN	Unknown	1	PrivateTag

(0019,"GE_GENESIS_REV3.0",39)	SS	AxialType	1	PrivateTag
(0019,"GE_GENESIS_REV3.0",8f)	SS	SwapPhaseFrequency	1	PrivateTag
(0019,"GE_GENESIS_REV3.0",9c)	SS	PulseSequenceName	1	PrivateTag
(0019,"GE_GENESIS_REV3.0",9f)	SS	CoilType	1	PrivateTag
(0019,"GE_GENESIS_REV3.0",a4)	SS	SATFatWaterBone	1	PrivateTag
(0019,"GE_GENESIS_REV3.0",c0)	SS	BitmapOfSATSelections	1	PrivateTag
(0019,"GE_GENESIS_REV3.0",c1)	SS	SurfaceCoilIntensityCorrectionFlag	1	PrivateTag
(0019,"GE_GENESIS_REV3.0",cb)	SS	PhaseContrastFlowAxis	1	PrivateTag
(0019,"GE_GENESIS_REV3.0",cc)	SS	PhaseContrastVelocityEncoding	1	PrivateTag
(0019,"GE_GENESIS_REV3.0",d5)	SS	FractionalEcho	1	PrivateTag
(0019,"GE_GENESIS_REV3.0",d8)	SS	VariableEchoFlag	1	PrivateTag
(0019,"GE_GENESIS_REV3.0",d9)	DS	ConcatenatedSat	1	PrivateTag
(0019,"GE_GENESIS_REV3.0",f2)	SS	NumberOfPhases	1	PrivateTag
(0043,"GE_GENESIS_REV3.0",1e)	DS	DeltaStartTime	1	PrivateTag
(0043,"GE_GENESIS_REV3.0",27)	SH	ScanPitchRatio	1	PrivateTag

(0029,"INTELERAD MEDICAL SYSTEMS",01)	FD	ImageCompressionFraction	1	PrivateTag
(0029,"INTELERAD MEDICAL SYSTEMS",02)	FD	ImageQuality	1	PrivateTag
(0029,"INTELERAD MEDICAL SYSTEMS",03)	FD	ImageBytesTransferred	1	PrivateTag
(0029,"INTELERAD MEDICAL SYSTEMS",10)	SH	J2cParameterType	1	PrivateTag
(0029,"INTELERAD MEDICAL SYSTEMS",11)	US	J2cPixelRepresentation	1	PrivateTag
(0029,"INTELERAD MEDICAL SYSTEMS",12)	US	J2cBitsAllocated	1	PrivateTag
(0029,"INTELERAD MEDICAL SYSTEMS",13)	US	J2cPixelShiftValue	1	PrivateTag
(0029,"INTELERAD MEDICAL SYSTEMS",14)	US	J2cPlanarConfiguration	1	PrivateTag
(0029,"INTELERAD MEDICAL SYSTEMS",15)	DS	J2cRescaleIntercept	1	PrivateTag
(0029,"INTELERAD MEDICAL SYSTEMS",20)	LO	PixelDataMD5SumPerFrame	1	PrivateTag
(0029,"INTELERAD MEDICAL SYSTEMS",21)	US	HistogramPercentileLabels	1	PrivateTag
(0029,"INTELERAD MEDICAL SYSTEMS",22)	FD	HistogramPercentileValues	1	PrivateTag
(3f01,"INTELERAD MEDICAL SYSTEMS",01)	LO	InstitutionCode	1	PrivateTag
(3f01,"INTELERAD MEDICAL SYSTEMS",02)	LO	RoutedTransferAE	1	PrivateTag
(3f01,"INTELERAD MEDICAL SYSTEMS",03)	LO	SourceAE	1	PrivateTag
(3f01,"INTELERAD MEDICAL SYSTEMS",04)	SH	DeferredValidation	1	PrivateTag
(3f01,"INTELERAD MEDICAL SYSTEMS",05)	LO	SeriesOwner	1	PrivateTag
(3f01,"INTELERAD MEDICAL SYSTEMS",06)	LO	OrderGroupNumber	1	PrivateTag
(3f01,"INTELERAD MEDICAL SYSTEMS",07)	SH	StrippedPixelData	1	PrivateTag
(3f01,"INTELERAD MEDICAL SYSTEMS",08)	SH	PendingMoveRequest	1	PrivateTag

(0041,"INTEGRIS 1.0",20)	FL	AccumulatedFluoroscopyDose	1	PrivateTag
(0041,"INTEGRIS 1.0",30)	FL	AccumulatedExposureDose	1	PrivateTag
(0041,"INTEGRIS 1.0",40)	FL	TotalDose	1	PrivateTag
(0041,"INTEGRIS 1.0",41)	FL	TotalNumberOfFrames	1	PrivateTag
(0041,"INTEGRIS 1.0",50)	SQ	ExposureInformationSequence	1	PrivateTag
(0009,"INTEGRIS 1.0",08)	CS	ExposureChannel	1-n	PrivateTag
(0009,"INTEGRIS 1.0",32)	TM	ExposureStartTime	1	PrivateTag
(0019,"INTEGRIS 1.0",00)	LO	APRName	1	PrivateTag
(0019,"INTEGRIS 1.0",40)	DS	FrameRate	1	PrivateTag
(0021,"INTEGRIS 1.0",12)	IS	ExposureNumber	1	PrivateTag
(0029,"INTEGRIS 1.0",08)	IS	NumberOfExposureResults	1	PrivateTag

(0029,"ISG shadow",70)	IS	Unknown	1	PrivateTag
(0029,"ISG shadow",80)	IS	Unknown	1	PrivateTag
(0029,"ISG shadow",90)	IS	Unknown	1	PrivateTag

(0009,"ISI",01)	UN	SIENETGeneralPurposeIMGEF	1	PrivateTag

(0009,"MERGE TECHNOLOGIES, INC.",00)	OB	Unknown	1	PrivateTag

(0029,"OCULUS Optikgeraete GmbH",1010)	OB	OriginalMeasuringData	1	PrivateTag
(0029,"OCULUS Optikgeraete GmbH",1012)	UL	OriginalMeasuringDataLength	1	PrivateTag
(0029,"OCULUS Optikgeraete GmbH",1020)	OB	OriginalMeasuringRawData	1	PrivateTag
(0029,"OCULUS Optikgeraete GmbH",1022)	UL	OriginalMeasuringRawDataLength	1	PrivateTag

(0041,"PAPYRUS 3.0",00)	LT	PapyrusComments	1	PrivateTag
(0041,"PAPYRUS 3.0",10)	SQ	PointerSequence	1	PrivateTag
(0041,"PAPYRUS 3.0",11)	UL	ImagePointer	1	PrivateTag
(0041,"PAPYRUS 3.0",12)	UL	PixelOffset	1	PrivateTag
(0041,"PAPYRUS 3.0",13)	SQ	ImageIdentifierSequence	1	PrivateTag
(0041,"PAPYRUS 3.0",14)	SQ	ExternalFileReferenceSequence	1	PrivateTag
(0041,"PAPYRUS 3.0",15)	US	NumberOfImages	1	PrivateTag
(0041,"PAPYRUS 3.0",21)	UI	ReferencedSOPClassUID	1	PrivateTag
(0041,"PAPYRUS 3.0",22)	UI	ReferencedSOPInstanceUID	1	PrivateTag
(0041,"PAPYRUS 3.0",31)	LT	ReferencedFileName	1	PrivateTag
(0041,"PAPYRUS 3.0",32)	LT	ReferencedFilePath	1-n	PrivateTag
(0041,"PAPYRUS 3.0",41)	UI	ReferencedImageSOPClassUID	1	PrivateTag
(0041,"PAPYRUS 3.0",42)	UI	ReferencedImageSOPInstanceUID	1	PrivateTag
(0041,"PAPYRUS 3.0",50)	SQ	ImageSequence	1	PrivateTag
(6001-o-60ff,"PAPYRUS 3.0",00)	IS	OverlayID	1	PrivateTag
(6001-o-60ff,"PAPYRUS 3.0",01)	LT	LinkedOverlays	1-n	PrivateTag
(6001-o-60ff,"PAPYRUS 3.0",10)	US	OverlayRows	1	PrivateTag
(6001-o-60ff,"PAPYRUS 3.0",11)	US	OverlayColumns	1	PrivateTag
(6001-o-60ff,"PAPYRUS 3.0",40)	LO	OverlayType	1	PrivateTag
(6001-o-60ff,"PAPYRUS 3.0",50)	US	OverlayOrigin	1-n	PrivateTag
(6001-o-60ff,"PAPYRUS 3.0",60)	LO	Editable	1	PrivateTag
(6001-o-60ff,"PAPYRUS 3.0",70)	LO	OverlayFont	1	PrivateTag
(6001-o-60ff,"PAPYRUS 3.0",72)	LO	OverlayStyle	1	PrivateTag
(6001-o-60ff,"PAPYRUS 3.0",74)	US	OverlayFontSize	1	PrivateTag
(6001-o-60ff,"PAPYRUS 3.0",76)	LO	OverlayColor	1	PrivateTag
(6001-o-60ff,"PAPYRUS 3.0",78)	US	ShadowSize	1	PrivateTag
(6001-o-60ff,"PAPYRUS 3.0",80)	LO	FillPattern	1	PrivateTag
(6001-o-60ff,"PAPYRUS 3.0",82)	US	OverlayPenSize	1	PrivateTag
(6001-o-60ff,"PAPYRUS 3.0",a0)	LO	Label	1	PrivateTag
(6001-o-60ff,"PAPYRUS 3.0",a2)	LT	PostItText	1	PrivateTag
(6001-o-60ff,"PAPYRUS 3.0",a4)	US	AnchorPoint	2	PrivateTag
(6001-o-60ff,"PAPYRUS 3.0",b0)	LO	ROIType	1	PrivateTag
(6001-o-60ff,"PAPYRUS 3.0",b2)	LT	AttachedAnnotation	1	PrivateTag
(6001-o-60ff,"PAPYRUS 3.0",ba)	US	ContourPoints	1-n	PrivateTag
(6001-o-60ff,"PAPYRUS 3.0",bc)	US	MaskData	1-n	PrivateTag
(6001-o-60ff,"PAPYRUS 3.0",c0)	SQ	UINOverlaySequence	1	PrivateTag

(0009,"PAPYRUS",00)	LT	OriginalFileName	1	PrivateTag
(0009,"PAPYRUS",10)	LT	OriginalFileLocation	1	PrivateTag
(0009,"PAPYRUS",18)	LT	DataSetIdentifier	1	PrivateTag
(0041,"PAPYRUS",00)	LT	PapyrusComments	1-n	PrivateTag
(0041,"PAPYRUS",10)	US	FolderType	1	PrivateTag
(0041,"PAPYRUS",11)	LT	PatientFolderDataSetID	1	PrivateTag
(0041,"PAPYRUS",20)	LT	FolderName	1	PrivateTag
(0041,"PAPYRUS",30)	DA	CreationDate	1	PrivateTag
(0041,"PAPYRUS",32)	TM	CreationTime	1	PrivateTag
(0041,"PAPYRUS",34)	DA	ModifiedDate	1	PrivateTag
(0041,"PAPYRUS",36)	TM	ModifiedTime	1	PrivateTag
(0041,"PAPYRUS",40)	LT	OwnerName	1-n	PrivateTag
(0041,"PAPYRUS",50)	LT	FolderStatus	1	PrivateTag
(0041,"PAPYRUS",60)	UL	NumberOfImages	1	PrivateTag
(0041,"PAPYRUS",62)	UL	NumberOfOther	1	PrivateTag
(0041,"PAPYRUS",a0)	LT	ExternalFolderElementDSID	1-n	PrivateTag
(0041,"PAPYRUS",a1)	US	ExternalFolderElementDataSetType	1-n	PrivateTag
(0041,"PAPYRUS",a2)	LT	ExternalFolderElementFileLocation	1-n	PrivateTag
(0041,"PAPYRUS",a3)	UL	ExternalFolderElementLength	1-n	PrivateTag
(0041,"PAPYRUS",b0)	LT	InternalFolderElementDSID	1-n	PrivateTag
(0041,"PAPYRUS",b1)	US	InternalFolderElementDataSetType	1-n	PrivateTag
(0041,"PAPYRUS",b2)	UL	InternalOffsetToDataSet	1-n	PrivateTag
(0041,"PAPYRUS",b3)	UL	InternalOffsetToImage	1-n

# Note: Some Philips devices use these private tags with reservation value
# "Philips Imaging DD 001", others use "PHILIPS IMAGING DD 001". All attributes
# should thus be present twice in this dictionary, once for each spelling variant.
#
(2001,"Philips Imaging DD 001",01)	FL	ChemicalShift	1	PrivateTag
(2001,"Philips Imaging DD 001",02)	IS	ChemicalShiftNumberMR	1	PrivateTag
(2001,"Philips Imaging DD 001",03)	FL	DiffusionBFactor	1	PrivateTag
(2001,"Philips Imaging DD 001",04)	CS	DiffusionDirection	1	PrivateTag
(2001,"Philips Imaging DD 001",06)	CS	ImageEnhanced	1	PrivateTag
(2001,"Philips Imaging DD 001",07)	CS	ImageTypeEDES	1	PrivateTag
(2001,"Philips Imaging DD 001",08)	IS	PhaseNumber	1	PrivateTag
(2001,"Philips Imaging DD 001",09)	FL	ImagePrepulseDelay	1	PrivateTag
(2001,"Philips Imaging DD 001",0a)	IS	SliceNumberMR	1	PrivateTag
(2001,"Philips Imaging DD 001",0b)	CS	SliceOrientation	1	PrivateTag
(2001,"Philips Imaging DD 001",0c)	CS	ArrhythmiaRejection	1	PrivateTag
(2001,"Philips Imaging DD 001",0e)	CS	CardiacCycled	1	PrivateTag
(2001,"Philips Imaging DD 001",0f)	SS	CardiacGateWidth	1	PrivateTag
(2001,"Philips Imaging DD 001",10)	CS	CardiacSync	1	PrivateTag
(2001,"Philips Imaging DD 001",11)	FL	DiffusionEchoTime	1	PrivateTag
(2001,"Philips Imaging DD 001",12)	CS	DynamicSeries	1	PrivateTag
(2001,"Philips Imaging DD 001",13)	SL	EPIFactor	1	PrivateTag
(2001,"Philips Imaging DD 001",14)	SL	NumberOfEchoes	1	PrivateTag
(2001,"Philips Imaging DD 001",15)	SS	NumberOfLocations	1	PrivateTag
(2001,"Philips Imaging DD 001",16)	SS	NumberOfPCDirections	1	PrivateTag
(2001,"Philips Imaging DD 001",17)	SL	NumberOfPhasesMR	1	PrivateTag
(2001,"Philips Imaging DD 001",18)	SL	NumberOfSlicesMR	1	PrivateTag
(2001,"Philips Imaging DD 001",19)	CS	PartialMatrixScanned	1	PrivateTag
(2001,"Philips Imaging DD 001",1a)	FL	PCVelocity	1-n	PrivateTag
(2001,"Philips Imaging DD 001",1b)	FL	PrepulseDelay	1	PrivateTag
(2001,"Philips Imaging DD 001",1c)	CS	PrepulseType	1	PrivateTag
(2001,"Philips Imaging DD 001",1d)	IS	ReconstructionNumberMR	1	PrivateTag
(2001,"Philips Imaging DD 001",1f)	CS	RespirationSync	1	PrivateTag
(2001,"Philips Imaging DD 001",20)	LO	ScanningTechnique	1	PrivateTag
(2001,"Philips Imaging DD 001",21)	CS	SPIR	1	PrivateTag
(2001,"Philips Imaging DD 001",22)	FL	WaterFatShift	1	PrivateTag
(2001,"Philips Imaging DD 001",23)	DS	FlipAnglePhilips	1	PrivateTag
(2001,"Philips Imaging DD 001",24)	CS	SeriesIsInteractive	1	PrivateTag
(2001,"Philips Imaging DD 001",25)	SH	EchoTimeDisplayMR	1	PrivateTag
(2001,"Philips Imaging DD 001",26)	CS	PresentationStateSubtractionActive	1	PrivateTag
(2001,"Philips Imaging DD 001",2d)	SS	StackNumberOfSlices	1	PrivateTag
(2001,"Philips Imaging DD 001",32)	FL	StackRadialAngle	1	PrivateTag
(2001,"Philips Imaging DD 001",33)	CS	StackRadialAxis	1	PrivateTag
(2001,"Philips Imaging DD 001",35)	SS	StackSliceNumber	1	PrivateTag
(2001,"Philips Imaging DD 001",36)	CS	StackType	1	PrivateTag
(2001,"Philips Imaging DD 001",3f)	CS	ZoomMode	1	PrivateTag
(2001,"Philips Imaging DD 001",58)	UL	ContrastTransferTaste	1	PrivateTag
(2001,"Philips Imaging DD 001",5f)	SQ	StackSequence	1	PrivateTag
(2001,"Philips Imaging DD 001",60)	SL	NumberOfStacks	1	PrivateTag
(2001,"Philips Imaging DD 001",61)	CS	SeriesTransmitted	1	PrivateTag
(2001,"Philips Imaging DD 001",62)	CS	SeriesCommitted	1	PrivateTag
(2001,"Philips Imaging DD 001",63)	CS	ExaminationSource	1	PrivateTag
(2001,"Philips Imaging DD 001",67)	CS	LinearPresentationGLTrafoShapeSub	1	PrivateTag
(2001,"Philips Imaging DD 001",77)	CS	GLTrafoType	1	PrivateTag
(2001,"Philips Imaging DD 001",7b)	IS	AcquisitionNumber	1	PrivateTag
(2001,"Philips Imaging DD 001",81)	IS	NumberOfDynamicScans	1	PrivateTag
(2001,"Philips Imaging DD 001",9f)	US	PixelProcessingKernelSize	1	PrivateTag
(2001,"Philips Imaging DD 001",a1)	CS	IsRawImage	1	PrivateTag
(2001,"Philips Imaging DD 001",f1)	FL	ProspectiveMotionCorrection	1	PrivateTag
(2001,"Philips Imaging DD 001",f2)	FL	RetrospectiveMotionCorrection	1	PrivateTag

# Note: Some Philips devices use these private tags with reservation value
# "Philips Imaging DD 001", others use "PHILIPS IMAGING DD 001". All attributes
# should thus be present twice in this dictionary, once for each spelling variant.
#
(2001,"PHILIPS IMAGING DD 001",01)	FL	ChemicalShift	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",02)	IS	ChemicalShiftNumberMR	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",03)	FL	DiffusionBFactor	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",04)	CS	DiffusionDirection	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",06)	CS	ImageEnhanced	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",07)	CS	ImageTypeEDES	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",08)	IS	PhaseNumber	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",09)	FL	ImagePrepulseDelay	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",0a)	IS	SliceNumberMR	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",0b)	CS	SliceOrientation	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",0c)	CS	ArrhythmiaRejection	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",0e)	CS	CardiacCycled	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",0f)	SS	CardiacGateWidth	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",10)	CS	CardiacSync	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",11)	FL	DiffusionEchoTime	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",12)	CS	DynamicSeries	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",13)	SL	EPIFactor	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",14)	SL	NumberOfEchoes	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",15)	SS	NumberOfLocations	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",16)	SS	NumberOfPCDirections	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",17)	SL	NumberOfPhasesMR	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",18)	SL	NumberOfSlicesMR	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",19)	CS	PartialMatrixScanned	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",1a)	FL	PCVelocity	1-n	PrivateTag
(2001,"PHILIPS IMAGING DD 001",1b)	FL	PrepulseDelay	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",1c)	CS	PrepulseType	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",1d)	IS	ReconstructionNumberMR	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",1f)	CS	RespirationSync	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",20)	LO	ScanningTechnique	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",21)	CS	SPIR	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",22)	FL	WaterFatShift	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",23)	DS	FlipAnglePhilips	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",24)	CS	SeriesIsInteractive	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",25)	SH	EchoTimeDisplayMR	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",26)	CS	PresentationStateSubtractionActive	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",2d)	SS	StackNumberOfSlices	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",32)	FL	StackRadialAngle	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",33)	CS	StackRadialAxis	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",35)	SS	StackSliceNumber	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",36)	CS	StackType	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",3f)	CS	ZoomMode	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",58)	UL	ContrastTransferTaste	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",5f)	SQ	StackSequence	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",60)	SL	NumberOfStacks	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",61)	CS	SeriesTransmitted	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",62)	CS	SeriesCommitted	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",63)	CS	ExaminationSource	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",67)	CS	LinearPresentationGLTrafoShapeSub	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",77)	CS	GLTrafoType	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",7b)	IS	AcquisitionNumber	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",81)	IS	NumberOfDynamicScans	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",9f)	US	PixelProcessingKernelSize	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",a1)	CS	IsRawImage	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",f1)	FL	ProspectiveMotionCorrection	1	PrivateTag
(2001,"PHILIPS IMAGING DD 001",f2)	FL	RetrospectiveMotionCorrection	1	PrivateTag

# Note: Some Philips devices use these private tags with reservation value
# "Philips MR Imaging DD 001", others use "PHILIPS MR IMAGING DD 001". All attributes
# should thus be present twice in this dictionary, once for each spelling variant.
#
(2005,"Philips MR Imaging DD 001",05)	CS	SynergyReconstructionType	1	PrivateTag
(2005,"Philips MR Imaging DD 001",1e)	SH	MIPProtocol	1	PrivateTag
(2005,"Philips MR Imaging DD 001",1f)	SH	MPRProtocol	1	PrivateTag
(2005,"Philips MR Imaging DD 001",20)	SL	NumberOfChemicalShifts	1	PrivateTag
(2005,"Philips MR Imaging DD 001",2d)	SS	NumberOfStackSlices	1	PrivateTag
(2005,"Philips MR Imaging DD 001",83)	SQ	Unknown	1	PrivateTag
(2005,"Philips MR Imaging DD 001",a1)	CS	SyncraScanType	1	PrivateTag
(2005,"Philips MR Imaging DD 001",b0)	FL	DiffusionDirectionRL	1	PrivateTag
(2005,"Philips MR Imaging DD 001",b1)	FL	DiffusionDirectionAP	1	PrivateTag
(2005,"Philips MR Imaging DD 001",b2)	FL	DiffusionDirectionFH	1	PrivateTag

(2005,"Philips MR Imaging DD 005",02)	SQ	Unknown	1	PrivateTag

# Note: Some Philips devices use these private tags with reservation value
# "Philips MR Imaging DD 001", others use "PHILIPS MR IMAGING DD 001". All attributes
# should thus be present twice in this dictionary, once for each spelling variant.
#
(2005,"PHILIPS MR IMAGING DD 001",05)	CS	SynergyReconstructionType	1	PrivateTag
(2005,"PHILIPS MR IMAGING DD 001",1e)	SH	MIPProtocol	1	PrivateTag
(2005,"PHILIPS MR IMAGING DD 001",1f)	SH	MPRProtocol	1	PrivateTag
(2005,"PHILIPS MR IMAGING DD 001",20)	SL	NumberOfChemicalShifts	1	PrivateTag
(2005,"PHILIPS MR IMAGING DD 001",2d)	SS	NumberOfStackSlices	1	PrivateTag
(2005,"PHILIPS MR IMAGING DD 001",83)	SQ	Unknown	1	PrivateTag
(2005,"PHILIPS MR IMAGING DD 001",a1)	CS	SyncraScanType	1	PrivateTag
(2005,"PHILIPS MR IMAGING DD 001",b0)	FL	DiffusionDirectionRL	1	PrivateTag
(2005,"PHILIPS MR IMAGING DD 001",b1)	FL	DiffusionDirectionAP	1	PrivateTag
(2005,"PHILIPS MR IMAGING DD 001",b2)	FL	DiffusionDirectionFH	1	PrivateTag

(0019,"PHILIPS MR R5.5/PART",1000)	DS	FieldOfView	1	PrivateTag
(0019,"PHILIPS MR R5.6/PART",1000)	DS	FieldOfView	1	PrivateTag

(0019,"PHILIPS MR SPECTRO;1",01)	US	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",02)	US	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",03)	US	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",04)	US	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",05)	US	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",06)	US	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",07)	IS	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",08)	IS	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",09)	IS	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",10)	IS	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",12)	IS	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",13)	IS	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",14)	US	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",15)	US	Unknown	1-n	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",16)	IS	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",17)	IS	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",18)	UN	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",20)	IS	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",21)	IS	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",22)	IS	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",23)	IS	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",24)	IS	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",25)	IS	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",26)	IS	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",27)	IS	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",28)	IS	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",29)	IS	Unknown	1-n	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",31)	US	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",32)	US	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",41)	LT	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",42)	IS	Unknown	2	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",43)	IS	Unknown	2	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",45)	US	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",46)	US	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",47)	IS	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",48)	IS	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",49)	US	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",50)	UN	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",60)	US	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",61)	US	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",70)	UN	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",71)	IS	Unknown	1-n	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",72)	US	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",73)	US	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",74)	US	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",76)	US	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",77)	US	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",78)	US	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",79)	US	Unknown	1	PrivateTag
(0019,"PHILIPS MR SPECTRO;1",80)	IS	Unknown	1	PrivateTag

(0009,"PHILIPS MR",10)	LO	SPIRelease	1	PrivateTag
(0009,"PHILIPS MR",12)	LO	Unknown	1	PrivateTag

(0019,"PHILIPS MR/LAST",09)	DS	MainMagneticField	1	PrivateTag
(0019,"PHILIPS MR/LAST",0e)	IS	FlowCompensation	1	PrivateTag
(0019,"PHILIPS MR/LAST",b1)	IS	MinimumRRInterval	1	PrivateTag
(0019,"PHILIPS MR/LAST",b2)	IS	MaximumRRInterval	1	PrivateTag
(0019,"PHILIPS MR/LAST",b3)	IS	NumberOfRejections	1	PrivateTag
(0019,"PHILIPS MR/LAST",b4)	IS	NumberOfRRIntervals	1-n	PrivateTag
(0019,"PHILIPS MR/LAST",b5)	IS	ArrhythmiaRejection	1	PrivateTag
(0019,"PHILIPS MR/LAST",c0)	DS	Unknown	1-n	PrivateTag
(0019,"PHILIPS MR/LAST",c6)	IS	CycledMultipleSlice	1	PrivateTag
(0019,"PHILIPS MR/LAST",ce)	IS	REST	1	PrivateTag
(0019,"PHILIPS MR/LAST",d5)	DS	Unknown	1	PrivateTag
(0019,"PHILIPS MR/LAST",d6)	IS	FourierInterpolation	1	PrivateTag
(0019,"PHILIPS MR/LAST",d9)	IS	Unknown	1-n	PrivateTag
(0019,"PHILIPS MR/LAST",e0)	IS	Prepulse	1	PrivateTag
(0019,"PHILIPS MR/LAST",e1)	DS	PrepulseDelay	1	PrivateTag
(0019,"PHILIPS MR/LAST",e2)	IS	Unknown	1	PrivateTag
(0019,"PHILIPS MR/LAST",e3)	DS	Unknown	1	PrivateTag
(0019,"PHILIPS MR/LAST",f0)	LT	WSProtocolString1	1	PrivateTag
(0019,"PHILIPS MR/LAST",f1)	LT	WSProtocolString2	1	PrivateTag
(0019,"PHILIPS MR/LAST",f2)	LT	WSProtocolString3	1	PrivateTag
(0019,"PHILIPS MR/LAST",f3)	LT	WSProtocolString4	1	PrivateTag
(0021,"PHILIPS MR/LAST",00)	IS	Unknown	1	PrivateTag
(0021,"PHILIPS MR/LAST",10)	IS	Unknown	1	PrivateTag
(0021,"PHILIPS MR/LAST",20)	IS	Unknown	1	PrivateTag
(0021,"PHILIPS MR/LAST",21)	DS	SliceGap	1	PrivateTag
(0021,"PHILIPS MR/LAST",22)	DS	StackRadialAngle	1	PrivateTag
(0027,"PHILIPS MR/LAST",00)	US	Unknown	1	PrivateTag
(0027,"PHILIPS MR/LAST",11)	US	Unknown	1-n	PrivateTag
(0027,"PHILIPS MR/LAST",12)	DS	Unknown	1-n	PrivateTag
(0027,"PHILIPS MR/LAST",13)	DS	Unknown	1-n	PrivateTag
(0027,"PHILIPS MR/LAST",14)	DS	Unknown	1-n	PrivateTag
(0027,"PHILIPS MR/LAST",15)	DS	Unknown	1-n	PrivateTag
(0027,"PHILIPS MR/LAST",16)	LO	Unknown	1	PrivateTag
(0029,"PHILIPS MR/LAST",10)	DS	FPMin	1	PrivateTag
(0029,"PHILIPS MR/LAST",20)	DS	FPMax	1	PrivateTag
(0029,"PHILIPS MR/LAST",30)	DS	ScaledMinimum	1	PrivateTag
(0029,"PHILIPS MR/LAST",40)	DS	ScaledMaximum	1	PrivateTag
(0029,"PHILIPS MR/LAST",50)	DS	WindowMinimum	1	PrivateTag
(0029,"PHILIPS MR/LAST",60)	DS	WindowMaximum	1	PrivateTag
(0029,"PHILIPS MR/LAST",61)	IS	Unknown	1	PrivateTag
(0029,"PHILIPS MR/LAST",70)	DS	Unknown	1	PrivateTag
(0029,"PHILIPS MR/LAST",71)	DS	Unknown	1	PrivateTag
(0029,"PHILIPS MR/LAST",72)	IS	Unknown	1	PrivateTag
(0029,"PHILIPS MR/LAST",80)	IS	ViewCenter	1	PrivateTag
(0029,"PHILIPS MR/LAST",81)	IS	ViewSize	1	PrivateTag
(0029,"PHILIPS MR/LAST",82)	IS	ViewZoom	1	PrivateTag
(0029,"PHILIPS MR/LAST",83)	IS	ViewTransform	1	PrivateTag
(6001,"PHILIPS MR/LAST",00)	LT	Unknown	1	PrivateTag

(0019,"PHILIPS MR/PART",1000)	DS	FieldOfView	1	PrivateTag
(0019,"PHILIPS MR/PART",1005)	DS	CCAngulation	1	PrivateTag
(0019,"PHILIPS MR/PART",1006)	DS	APAngulation	1	PrivateTag
(0019,"PHILIPS MR/PART",1007)	DS	LRAngulation	1	PrivateTag
(0019,"PHILIPS MR/PART",1008)	IS	PatientPosition	1	PrivateTag
(0019,"PHILIPS MR/PART",1009)	IS	PatientOrientation	1	PrivateTag
(0019,"PHILIPS MR/PART",100a)	IS	SliceOrientation	1	PrivateTag
(0019,"PHILIPS MR/PART",100b)	DS	LROffcenter	1	PrivateTag
(0019,"PHILIPS MR/PART",100c)	DS	CCOffcenter	1	PrivateTag
(0019,"PHILIPS MR/PART",100d)	DS	APOffcenter	1	PrivateTag
(0019,"PHILIPS MR/PART",100e)	DS	Unknown	1	PrivateTag
(0019,"PHILIPS MR/PART",100f)	IS	NumberOfSlices	1	PrivateTag
(0019,"PHILIPS MR/PART",1010)	DS	SliceFactor	1	PrivateTag
(0019,"PHILIPS MR/PART",1011)	DS	EchoTimes	1-n	PrivateTag
(0019,"PHILIPS MR/PART",1015)	IS	DynamicStudy	1	PrivateTag
(0019,"PHILIPS MR/PART",1018)	DS	HeartbeatInterval	1	PrivateTag
(0019,"PHILIPS MR/PART",1019)	DS	RepetitionTimeFFE	1	PrivateTag
(0019,"PHILIPS MR/PART",101a)	DS	FFEFlipAngle	1	PrivateTag
(0019,"PHILIPS MR/PART",101b)	IS	NumberOfScans	1	PrivateTag
(0019,"PHILIPS MR/PART",1021)	DS	Unknown	1-n	PrivateTag
(0019,"PHILIPS MR/PART",1022)	DS	DynamicScanTimeBegin	1	PrivateTag
(0019,"PHILIPS MR/PART",1024)	IS	Unknown	1	PrivateTag
(0019,"PHILIPS MR/PART",1064)	DS	RepetitionTimeSE	1	PrivateTag
(0019,"PHILIPS MR/PART",1065)	DS	RepetitionTimeIR	1	PrivateTag
(0019,"PHILIPS MR/PART",1069)	IS	NumberOfPhases	1	PrivateTag
(0019,"PHILIPS MR/PART",106a)	IS	CardiacFrequency	1	PrivateTag
(0019,"PHILIPS MR/PART",106b)	DS	InversionDelay	1	PrivateTag
(0019,"PHILIPS MR/PART",106c)	DS	GateDelay	1	PrivateTag
(0019,"PHILIPS MR/PART",106d)	DS	GateWidth	1	PrivateTag
(0019,"PHILIPS MR/PART",106e)	DS	TriggerDelayTime	1	PrivateTag
(0019,"PHILIPS MR/PART",1080)	IS	NumberOfChemicalShifts	1	PrivateTag
(0019,"PHILIPS MR/PART",1081)	DS	ChemicalShift	1	PrivateTag
(0019,"PHILIPS MR/PART",1084)	IS	NumberOfRows	1	PrivateTag
(0019,"PHILIPS MR/PART",1085)	IS	NumberOfSamples	1	PrivateTag
(0019,"PHILIPS MR/PART",1094)	LO	MagnetizationTransferContrast	1	PrivateTag
(0019,"PHILIPS MR/PART",1095)	LO	SpectralPresaturationWithInversionRecovery	1	PrivateTag
(0019,"PHILIPS MR/PART",1096)	IS	Unknown	1	PrivateTag
(0019,"PHILIPS MR/PART",1097)	LO	Unknown	1	PrivateTag
(0019,"PHILIPS MR/PART",10a0)	IS	Unknown	1	PrivateTag
(0019,"PHILIPS MR/PART",10a1)	DS	Unknown	1	PrivateTag
(0019,"PHILIPS MR/PART",10a3)	DS	Unknown	1	PrivateTag
(0019,"PHILIPS MR/PART",10a4)	CS	Unknown	1	PrivateTag
(0019,"PHILIPS MR/PART",10c8)	IS	Unknown	1	PrivateTag
(0019,"PHILIPS MR/PART",10c9)	IS	FoldoverDirectionTransverse	1	PrivateTag
(0019,"PHILIPS MR/PART",10ca)	IS	FoldoverDirectionSagittal	1	PrivateTag
(0019,"PHILIPS MR/PART",10cb)	IS	FoldoverDirectionCoronal	1	PrivateTag
(0019,"PHILIPS MR/PART",10cc)	IS	Unknown	1	PrivateTag
(0019,"PHILIPS MR/PART",10cd)	IS	Unknown	1	PrivateTag
(0019,"PHILIPS MR/PART",10ce)	IS	Unknown	1	PrivateTag
(0019,"PHILIPS MR/PART",10cf)	IS	NumberOfEchoes	1	PrivateTag
(0019,"PHILIPS MR/PART",10d0)	IS	ScanResolution	1	PrivateTag
(0019,"PHILIPS MR/PART",10d2)	LO	WaterFatShift	2	PrivateTag
(0019,"PHILIPS MR/PART",10d4)	IS	ArtifactReduction	1	PrivateTag
(0019,"PHILIPS MR/PART",10d5)	IS	Unknown	1	PrivateTag
(0019,"PHILIPS MR/PART",10d6)	IS	Unknown	1	PrivateTag
(0019,"PHILIPS MR/PART",10d7)	DS	ScanPercentage	1	PrivateTag
(0019,"PHILIPS MR/PART",10d8)	IS	Halfscan	1	PrivateTag
(0019,"PHILIPS MR/PART",10d9)	IS	EPIFactor	1	PrivateTag
(0019,"PHILIPS MR/PART",10da)	IS	TurboFactor	1	PrivateTag
(0019,"PHILIPS MR/PART",10db)	IS	Unknown	1	PrivateTag
(0019,"PHILIPS MR/PART",10e0)	IS	PercentageOfScanCompleted	1	PrivateTag
(0019,"PHILIPS MR/PART",10e1)	IS	Unknown	1	PrivateTag
(0019,"PHILIPS MR/PART",1100)	IS	NumberOfStacks	1	PrivateTag
(0019,"PHILIPS MR/PART",1101)	IS	StackType	1-n	PrivateTag
(0019,"PHILIPS MR/PART",1102)	IS	Unknown	1-n	PrivateTag
(0019,"PHILIPS MR/PART",110b)	DS	LROffcenter	1	PrivateTag
(0019,"PHILIPS MR/PART",110c)	DS	CCOffcenter	1	PrivateTag
(0019,"PHILIPS MR/PART",110d)	DS	APOffcenter	1	PrivateTag
(0019,"PHILIPS MR/PART",1145)	IS	ReconstructionResolution	1	PrivateTag
(0019,"PHILIPS MR/PART",11fc)	IS	ResonanceFrequency	1	PrivateTag
(0019,"PHILIPS MR/PART",12c0)	DS	TriggerDelayTimes	1	PrivateTag
(0019,"PHILIPS MR/PART",12e0)	IS	PrepulseType	1	PrivateTag
(0019,"PHILIPS MR/PART",12e1)	DS	PrepulseDelay	1	PrivateTag
(0019,"PHILIPS MR/PART",12e3)	DS	PhaseContrastVelocity	1	PrivateTag
(0021,"PHILIPS MR/PART",1000)	IS	ReconstructionNumber	1	PrivateTag
(0021,"PHILIPS MR/PART",1010)	IS	ImageType	1	PrivateTag
(0021,"PHILIPS MR/PART",1020)	IS	SliceNumber	1	PrivateTag
(0021,"PHILIPS MR/PART",1030)	IS	EchoNumber	1	PrivateTag
(0021,"PHILIPS MR/PART",1031)	DS	PatientReferenceID	1	PrivateTag
(0021,"PHILIPS MR/PART",1035)	IS	ChemicalShiftNumber	1	PrivateTag
(0021,"PHILIPS MR/PART",1040)	IS	PhaseNumber	1	PrivateTag
(0021,"PHILIPS MR/PART",1050)	IS	DynamicScanNumber	1	PrivateTag
(0021,"PHILIPS MR/PART",1060)	IS	NumberOfRowsInObject	1	PrivateTag
(0021,"PHILIPS MR/PART",1061)	IS	RowNumber	1-n	PrivateTag
(0021,"PHILIPS MR/PART",1062)	IS	Unknown	1-n	PrivateTag
(0021,"PHILIPS MR/PART",1100)	DA	ScanDate	1	PrivateTag
(0021,"PHILIPS MR/PART",1110)	TM	ScanTime	1	PrivateTag
(0021,"PHILIPS MR/PART",1221)	IS	SliceGap	1	PrivateTag
(0029,"PHILIPS MR/PART",00)	DS	Unknown	2	PrivateTag
(0029,"PHILIPS MR/PART",04)	US	Unknown	1	PrivateTag
(0029,"PHILIPS MR/PART",10)	DS	Unknown	1	PrivateTag
(0029,"PHILIPS MR/PART",11)	DS	Unknown	1	PrivateTag
(0029,"PHILIPS MR/PART",20)	LO	Unknown	1	PrivateTag
(0029,"PHILIPS MR/PART",31)	DS	Unknown	2	PrivateTag
(0029,"PHILIPS MR/PART",32)	DS	Unknown	2	PrivateTag
(0029,"PHILIPS MR/PART",c3)	IS	ScanResolution	1	PrivateTag
(0029,"PHILIPS MR/PART",c4)	IS	FieldOfView	1	PrivateTag
(0029,"PHILIPS MR/PART",d5)	LT	SliceThickness	1	PrivateTag

(0019,"PHILIPS-MR-1",11)	IS	ChemicalShiftNumber	1	PrivateTag
(0019,"PHILIPS-MR-1",12)	IS	PhaseNumber	1	PrivateTag
(0021,"PHILIPS-MR-1",01)	IS	ReconstructionNumber	1	PrivateTag
(0021,"PHILIPS-MR-1",02)	IS	SliceNumber	1	PrivateTag

(7001,"Picker NM Private Group",01)	UI	Unknown	1	PrivateTag
(7001,"Picker NM Private Group",02)	OB	Unknown	1	PrivateTag

(0019,"SIEMENS CM VA0  ACQU",10)	LT	ParameterFileName	1	PrivateTag
(0019,"SIEMENS CM VA0  ACQU",11)	LO	SequenceFileName	1	PrivateTag
(0019,"SIEMENS CM VA0  ACQU",12)	LT	SequenceFileOwner	1	PrivateTag
(0019,"SIEMENS CM VA0  ACQU",13)	LT	SequenceDescription	1	PrivateTag
(0019,"SIEMENS CM VA0  ACQU",14)	LT	EPIFileName	1	PrivateTag

(0009,"SIEMENS CM VA0  CMS",00)	DS	NumberOfMeasurements	1	PrivateTag
(0009,"SIEMENS CM VA0  CMS",10)	LT	StorageMode	1	PrivateTag
(0009,"SIEMENS CM VA0  CMS",12)	UL	EvaluationMaskImage	1	PrivateTag
(0009,"SIEMENS CM VA0  CMS",26)	DA	LastMoveDate	1	PrivateTag
(0009,"SIEMENS CM VA0  CMS",27)	TM	LastMoveTime	1	PrivateTag
(0011,"SIEMENS CM VA0  CMS",0a)	LT	Unknown	1	PrivateTag
(0011,"SIEMENS CM VA0  CMS",10)	DA	RegistrationDate	1	PrivateTag
(0011,"SIEMENS CM VA0  CMS",11)	TM	RegistrationTime	1	PrivateTag
(0011,"SIEMENS CM VA0  CMS",22)	LT	Unknown	1	PrivateTag
(0011,"SIEMENS CM VA0  CMS",23)	DS	UsedPatientWeight	1	PrivateTag
(0011,"SIEMENS CM VA0  CMS",40)	IS	OrganCode	1	PrivateTag
(0013,"SIEMENS CM VA0  CMS",00)	LT	ModifyingPhysician	1	PrivateTag
(0013,"SIEMENS CM VA0  CMS",10)	DA	ModificationDate	1	PrivateTag
(0013,"SIEMENS CM VA0  CMS",12)	TM	ModificationTime	1	PrivateTag
(0013,"SIEMENS CM VA0  CMS",20)	LO	PatientName	1	PrivateTag
(0013,"SIEMENS CM VA0  CMS",22)	LO	PatientId	1	PrivateTag
(0013,"SIEMENS CM VA0  CMS",30)	DA	PatientBirthdate	1	PrivateTag
(0013,"SIEMENS CM VA0  CMS",31)	DS	PatientWeight	1	PrivateTag
(0013,"SIEMENS CM VA0  CMS",32)	LT	PatientsMaidenName	1	PrivateTag
(0013,"SIEMENS CM VA0  CMS",33)	LT	ReferringPhysician	1	PrivateTag
(0013,"SIEMENS CM VA0  CMS",34)	LT	AdmittingDiagnosis	1	PrivateTag
(0013,"SIEMENS CM VA0  CMS",35)	LO	PatientSex	1	PrivateTag
(0013,"SIEMENS CM VA0  CMS",40)	LO	ProcedureDescription	1	PrivateTag
(0013,"SIEMENS CM VA0  CMS",42)	LO	RestDirection	1	PrivateTag
(0013,"SIEMENS CM VA0  CMS",44)	LO	PatientPosition	1	PrivateTag
(0013,"SIEMENS CM VA0  CMS",46)	LT	ViewDirection	1	PrivateTag
(0013,"SIEMENS CM VA0  CMS",50)	LT	Unknown	1	PrivateTag
(0013,"SIEMENS CM VA0  CMS",51)	LT	Unknown	1	PrivateTag
(0013,"SIEMENS CM VA0  CMS",52)	LT	Unknown	1	PrivateTag
(0013,"SIEMENS CM VA0  CMS",53)	LT	Unknown	1	PrivateTag
(0013,"SIEMENS CM VA0  CMS",54)	LT	Unknown	1	PrivateTag
(0013,"SIEMENS CM VA0  CMS",55)	LT	Unknown	1	PrivateTag
(0013,"SIEMENS CM VA0  CMS",56)	LT	Unknown	1	PrivateTag
(0019,"SIEMENS CM VA0  CMS",10)	DS	NetFrequency	1	PrivateTag
(0019,"SIEMENS CM VA0  CMS",20)	LT	MeasurementMode	1	PrivateTag
(0019,"SIEMENS CM VA0  CMS",30)	LT	CalculationMode	1	PrivateTag
(0019,"SIEMENS CM VA0  CMS",50)	IS	NoiseLevel	1	PrivateTag
(0019,"SIEMENS CM VA0  CMS",60)	IS	NumberOfDataBytes	1	PrivateTag
(0021,"SIEMENS CM VA0  CMS",20)	DS	FoV	2	PrivateTag
(0021,"SIEMENS CM VA0  CMS",22)	DS	ImageMagnificationFactor	1	PrivateTag
(0021,"SIEMENS CM VA0  CMS",24)	DS	ImageScrollOffset	2	PrivateTag
(0021,"SIEMENS CM VA0  CMS",26)	IS	ImagePixelOffset	1	PrivateTag
(0021,"SIEMENS CM VA0  CMS",30)	LT	ViewDirection	1	PrivateTag
(0021,"SIEMENS CM VA0  CMS",32)	CS	PatientRestDirection	1	PrivateTag
(0021,"SIEMENS CM VA0  CMS",60)	DS	ImagePosition	3	PrivateTag
(0021,"SIEMENS CM VA0  CMS",61)	DS	ImageNormal	3	PrivateTag
(0021,"SIEMENS CM VA0  CMS",63)	DS	ImageDistance	1	PrivateTag
(0021,"SIEMENS CM VA0  CMS",65)	US	ImagePositioningHistoryMask	1	PrivateTag
(0021,"SIEMENS CM VA0  CMS",6a)	DS	ImageRow	3	PrivateTag
(0021,"SIEMENS CM VA0  CMS",6b)	DS	ImageColumn	3	PrivateTag
(0021,"SIEMENS CM VA0  CMS",70)	LT	PatientOrientationSet1	3	PrivateTag
(0021,"SIEMENS CM VA0  CMS",71)	LT	PatientOrientationSet2	3	PrivateTag
(0021,"SIEMENS CM VA0  CMS",80)	LT	StudyName	1	PrivateTag
(0021,"SIEMENS CM VA0  CMS",82)	LT	StudyType	3	PrivateTag
(0029,"SIEMENS CM VA0  CMS",10)	LT	WindowStyle	1	PrivateTag
(0029,"SIEMENS CM VA0  CMS",11)	LT	Unknown	1	PrivateTag
(0029,"SIEMENS CM VA0  CMS",13)	LT	Unknown	1	PrivateTag
(0029,"SIEMENS CM VA0  CMS",20)	LT	PixelQualityCode	3	PrivateTag
(0029,"SIEMENS CM VA0  CMS",22)	IS	PixelQualityValue	3	PrivateTag
(0029,"SIEMENS CM VA0  CMS",50)	LT	ArchiveCode	1	PrivateTag
(0029,"SIEMENS CM VA0  CMS",51)	LT	ExposureCode	1	PrivateTag
(0029,"SIEMENS CM VA0  CMS",52)	LT	SortCode	1	PrivateTag
(0029,"SIEMENS CM VA0  CMS",53)	LT	Unknown	1	PrivateTag
(0029,"SIEMENS CM VA0  CMS",60)	LT	Splash	1	PrivateTag
(0051,"SIEMENS CM VA0  CMS",10)	LT	ImageText	1-n	PrivateTag
(6021,"SIEMENS CM VA0  CMS",00)	LT	ImageGraphicsFormatCode	1	PrivateTag
(6021,"SIEMENS CM VA0  CMS",10)	LT	ImageGraphics	1	PrivateTag
(7fe1,"SIEMENS CM VA0  CMS",00)	OB	BinaryData	1-n	PrivateTag

(0009,"SIEMENS CM VA0  LAB",10)	LT	GeneratorIdentificationLabel	1	PrivateTag
(0009,"SIEMENS CM VA0  LAB",11)	LT	GantryIdentificationLabel	1	PrivateTag
(0009,"SIEMENS CM VA0  LAB",12)	LT	X-RayTubeIdentificationLabel	1	PrivateTag
(0009,"SIEMENS CM VA0  LAB",13)	LT	DetectorIdentificationLabel	1	PrivateTag
(0009,"SIEMENS CM VA0  LAB",14)	LT	DASIdentificationLabel	1	PrivateTag
(0009,"SIEMENS CM VA0  LAB",15)	LT	SMIIdentificationLabel	1	PrivateTag
(0009,"SIEMENS CM VA0  LAB",16)	LT	CPUIdentificationLabel	1	PrivateTag
(0009,"SIEMENS CM VA0  LAB",20)	LT	HeaderVersion	1	PrivateTag

(0029,"SIEMENS CSA HEADER",08)	CS	CSAImageHeaderType	1	PrivateTag
(0029,"SIEMENS CSA HEADER",09)	LO	CSAImageHeaderVersion	1	PrivateTag
(0029,"SIEMENS CSA HEADER",10)	OB	CSAImageHeaderInfo	1	PrivateTag
(0029,"SIEMENS CSA HEADER",18)	CS	CSASeriesHeaderType	1	PrivateTag
(0029,"SIEMENS CSA HEADER",19)	LO	CSASeriesHeaderVersion	1	PrivateTag
(0029,"SIEMENS CSA HEADER",20)	OB	CSASeriesHeaderInfo	1	PrivateTag

(0029,"SIEMENS CSA NON-IMAGE",08)	CS	CSADataType	1	PrivateTag
(0029,"SIEMENS CSA NON-IMAGE",09)	LO	CSADataVersion	1	PrivateTag
(0029,"SIEMENS CSA NON-IMAGE",10)	OB	CSADataInfo	1	PrivateTag
(7FE1,"SIEMENS CSA NON-IMAGE",10)	OB	CSAData	1	PrivateTag

(0019,"SIEMENS CT VA0  COAD",10)	DS	DistanceSourceToSourceSideCollimator	1	PrivateTag
(0019,"SIEMENS CT VA0  COAD",11)	DS	DistanceSourceToDetectorSideCollimator	1	PrivateTag
(0019,"SIEMENS CT VA0  COAD",20)	IS	NumberOfPossibleChannels	1	PrivateTag
(0019,"SIEMENS CT VA0  COAD",21)	IS	MeanChannelNumber	1	PrivateTag
(0019,"SIEMENS CT VA0  COAD",22)	DS	DetectorSpacing	1	PrivateTag
(0019,"SIEMENS CT VA0  COAD",23)	DS	DetectorCenter	1	PrivateTag
(0019,"SIEMENS CT VA0  COAD",24)	DS	ReadingIntegrationTime	1	PrivateTag
(0019,"SIEMENS CT VA0  COAD",50)	DS	DetectorAlignment	1	PrivateTag
(0019,"SIEMENS CT VA0  COAD",52)	DS	Unknown	1	PrivateTag
(0019,"SIEMENS CT VA0  COAD",54)	DS	Unknown	1	PrivateTag
(0019,"SIEMENS CT VA0  COAD",60)	DS	FocusAlignment	1	PrivateTag
(0019,"SIEMENS CT VA0  COAD",65)	UL	FocalSpotDeflectionAmplitude	1	PrivateTag
(0019,"SIEMENS CT VA0  COAD",66)	UL	FocalSpotDeflectionPhase	1	PrivateTag
(0019,"SIEMENS CT VA0  COAD",67)	UL	FocalSpotDeflectionOffset	1	PrivateTag
(0019,"SIEMENS CT VA0  COAD",70)	DS	WaterScalingFactor	1	PrivateTag
(0019,"SIEMENS CT VA0  COAD",71)	DS	InterpolationFactor	1	PrivateTag
(0019,"SIEMENS CT VA0  COAD",80)	LT	PatientRegion	1	PrivateTag
(0019,"SIEMENS CT VA0  COAD",82)	LT	PatientPhaseOfLife	1	PrivateTag
(0019,"SIEMENS CT VA0  COAD",90)	DS	OsteoOffset	1	PrivateTag
(0019,"SIEMENS CT VA0  COAD",92)	DS	OsteoRegressionLineSlope	1	PrivateTag
(0019,"SIEMENS CT VA0  COAD",93)	DS	OsteoRegressionLineIntercept	1	PrivateTag
(0019,"SIEMENS CT VA0  COAD",94)	DS	OsteoStandardizationCode	1	PrivateTag
(0019,"SIEMENS CT VA0  COAD",96)	IS	OsteoPhantomNumber	1	PrivateTag
(0019,"SIEMENS CT VA0  COAD",A3)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS CT VA0  COAD",A4)	DS	Unknown	1	PrivateTag
(0019,"SIEMENS CT VA0  COAD",A5)	DS	Unknown	1	PrivateTag
(0019,"SIEMENS CT VA0  COAD",A6)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS CT VA0  COAD",A7)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS CT VA0  COAD",A8)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS CT VA0  COAD",A9)	DS	Unknown	1	PrivateTag
(0019,"SIEMENS CT VA0  COAD",AA)	LT	Unknown	1	PrivateTag
(0019,"SIEMENS CT VA0  COAD",AB)	DS	Unknown	1	PrivateTag
(0019,"SIEMENS CT VA0  COAD",AC)	DS	Unknown	1	PrivateTag
(0019,"SIEMENS CT VA0  COAD",AD)	DS	Unknown	1	PrivateTag
(0019,"SIEMENS CT VA0  COAD",AE)	DS	Unknown	1	PrivateTag
(0019,"SIEMENS CT VA0  COAD",AF)	DS	Unknown	1	PrivateTag
(0019,"SIEMENS CT VA0  COAD",B0)	DS	FeedPerRotation	1	PrivateTag
(0019,"SIEMENS CT VA0  COAD",BD)	IS	PulmoTriggerLevel	1	PrivateTag
(0019,"SIEMENS CT VA0  COAD",BE)	DS	ExpiratoricReserveVolume	1	PrivateTag
(0019,"SIEMENS CT VA0  COAD",BF)	DS	VitalCapacity	1	PrivateTag
(0019,"SIEMENS CT VA0  COAD",C0)	DS	PulmoWater	1	PrivateTag
(0019,"SIEMENS CT VA0  COAD",C1)	DS	PulmoAir	1	PrivateTag
(0019,"SIEMENS CT VA0  COAD",C2)	DA	PulmoDate	1	PrivateTag
(0019,"SIEMENS CT VA0  COAD",C3)	TM	PulmoTime	1	PrivateTag

(0019,"SIEMENS CT VA0  GEN",10)	DS	SourceSideCollimatorAperture	1	PrivateTag
(0019,"SIEMENS CT VA0  GEN",11)	DS	DetectorSideCollimatorAperture	1	PrivateTag
(0019,"SIEMENS CT VA0  GEN",20)	DS	ExposureTime	1	PrivateTag
(0019,"SIEMENS CT VA0  GEN",21)	DS	ExposureCurrent	1	PrivateTag
(0019,"SIEMENS CT VA0  GEN",25)	DS	KVPGeneratorPowerCurrent	1	PrivateTag
(0019,"SIEMENS CT VA0  GEN",26)	DS	GeneratorVoltage	1	PrivateTag
(0019,"SIEMENS CT VA0  GEN",40)	UL	MasterControlMask	1	PrivateTag
(0019,"SIEMENS CT VA0  GEN",42)	US	ProcessingMask	5	PrivateTag
(0019,"SIEMENS CT VA0  GEN",44)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS CT VA0  GEN",45)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS CT VA0  GEN",62)	IS	NumberOfVirtuellChannels	1	PrivateTag
(0019,"SIEMENS CT VA0  GEN",70)	IS	NumberOfReadings	1	PrivateTag
(0019,"SIEMENS CT VA0  GEN",71)	LT	Unknown	1-n	PrivateTag
(0019,"SIEMENS CT VA0  GEN",74)	IS	NumberOfProjections	1	PrivateTag
(0019,"SIEMENS CT VA0  GEN",75)	IS	NumberOfBytes	1	PrivateTag
(0019,"SIEMENS CT VA0  GEN",80)	LT	ReconstructionAlgorithmSet	1	PrivateTag
(0019,"SIEMENS CT VA0  GEN",81)	LT	ReconstructionAlgorithmIndex	1	PrivateTag
(0019,"SIEMENS CT VA0  GEN",82)	LT	RegenerationSoftwareVersion	1	PrivateTag
(0019,"SIEMENS CT VA0  GEN",88)	DS	Unknown	1	PrivateTag
(0021,"SIEMENS CT VA0  GEN",10)	IS	RotationAngle	1	PrivateTag
(0021,"SIEMENS CT VA0  GEN",11)	IS	StartAngle	1	PrivateTag
(0021,"SIEMENS CT VA0  GEN",20)	US	Unknown	1-n	PrivateTag
(0021,"SIEMENS CT VA0  GEN",30)	IS	TopogramTubePosition	1	PrivateTag
(0021,"SIEMENS CT VA0  GEN",32)	DS	LengthOfTopogram	1	PrivateTag
(0021,"SIEMENS CT VA0  GEN",34)	DS	TopogramCorrectionFactor	1	PrivateTag
(0021,"SIEMENS CT VA0  GEN",36)	DS	MaximumTablePosition	1	PrivateTag
(0021,"SIEMENS CT VA0  GEN",40)	IS	TableMoveDirectionCode	1	PrivateTag
(0021,"SIEMENS CT VA0  GEN",45)	IS	VOIStartRow	1	PrivateTag
(0021,"SIEMENS CT VA0  GEN",46)	IS	VOIStopRow	1	PrivateTag
(0021,"SIEMENS CT VA0  GEN",47)	IS	VOIStartColumn	1	PrivateTag
(0021,"SIEMENS CT VA0  GEN",48)	IS	VOIStopColumn	1	PrivateTag
(0021,"SIEMENS CT VA0  GEN",49)	IS	VOIStartSlice	1	PrivateTag
(0021,"SIEMENS CT VA0  GEN",4a)	IS	VOIStopSlice	1	PrivateTag
(0021,"SIEMENS CT VA0  GEN",50)	IS	VectorStartRow	1	PrivateTag
(0021,"SIEMENS CT VA0  GEN",51)	IS	VectorRowStep	1	PrivateTag
(0021,"SIEMENS CT VA0  GEN",52)	IS	VectorStartColumn	1	PrivateTag
(0021,"SIEMENS CT VA0  GEN",53)	IS	VectorColumnStep	1	PrivateTag
(0021,"SIEMENS CT VA0  GEN",60)	IS	RangeTypeCode	1	PrivateTag
(0021,"SIEMENS CT VA0  GEN",62)	IS	ReferenceTypeCode	1	PrivateTag
(0021,"SIEMENS CT VA0  GEN",70)	DS	ObjectOrientation	3	PrivateTag
(0021,"SIEMENS CT VA0  GEN",72)	DS	LightOrientation	3	PrivateTag
(0021,"SIEMENS CT VA0  GEN",75)	DS	LightBrightness	1	PrivateTag
(0021,"SIEMENS CT VA0  GEN",76)	DS	LightContrast	1	PrivateTag
(0021,"SIEMENS CT VA0  GEN",7a)	IS	OverlayThreshold	2	PrivateTag
(0021,"SIEMENS CT VA0  GEN",7b)	IS	SurfaceThreshold	2	PrivateTag
(0021,"SIEMENS CT VA0  GEN",7c)	IS	GreyScaleThreshold	2	PrivateTag
(0021,"SIEMENS CT VA0  GEN",a0)	DS	Unknown	1	PrivateTag
(0021,"SIEMENS CT VA0  GEN",a2)	LT	Unknown	1	PrivateTag
(0021,"SIEMENS CT VA0  GEN",a7)	LT	Unknown	1	PrivateTag

(0009,"SIEMENS CT VA0  IDE",10)	LT	Unknown	1	PrivateTag
(0009,"SIEMENS CT VA0  IDE",30)	LT	Unknown	1	PrivateTag
(0009,"SIEMENS CT VA0  IDE",31)	LT	Unknown	1	PrivateTag
(0009,"SIEMENS CT VA0  IDE",32)	LT	Unknown	1	PrivateTag
(0009,"SIEMENS CT VA0  IDE",34)	LT	Unknown	1	PrivateTag
(0009,"SIEMENS CT VA0  IDE",40)	LT	Unknown	1	PrivateTag
(0009,"SIEMENS CT VA0  IDE",42)	LT	Unknown	1	PrivateTag
(0009,"SIEMENS CT VA0  IDE",50)	LT	Unknown	1	PrivateTag
(0009,"SIEMENS CT VA0  IDE",51)	LT	Unknown	1	PrivateTag

(0009,"SIEMENS CT VA0  ORI",20)	LT	Unknown	1	PrivateTag
(0009,"SIEMENS CT VA0  ORI",30)	LT	Unknown	1	PrivateTag

(6021,"SIEMENS CT VA0  OST",00)	LT	OsteoContourComment	1	PrivateTag
(6021,"SIEMENS CT VA0  OST",10)	US	OsteoContourBuffer	256	PrivateTag

(0021,"SIEMENS CT VA0  RAW",10)	UL	CreationMask	2	PrivateTag
(0021,"SIEMENS CT VA0  RAW",20)	UL	EvaluationMask	2	PrivateTag
(0021,"SIEMENS CT VA0  RAW",30)	US	ExtendedProcessingMask	7	PrivateTag
(0021,"SIEMENS CT VA0  RAW",40)	US	Unknown	1-n	PrivateTag
(0021,"SIEMENS CT VA0  RAW",41)	US	Unknown	1-n	PrivateTag
(0021,"SIEMENS CT VA0  RAW",42)	US	Unknown	1-n	PrivateTag
(0021,"SIEMENS CT VA0  RAW",43)	US	Unknown	1-n	PrivateTag
(0021,"SIEMENS CT VA0  RAW",44)	US	Unknown	1-n	PrivateTag
(0021,"SIEMENS CT VA0  RAW",50)	LT	Unknown	1	PrivateTag

(0009,"SIEMENS DICOM",10)	UN	Unknown	1	PrivateTag
(0009,"SIEMENS DICOM",12)	LT	Unknown	1	PrivateTag

(0019,"SIEMENS DLR.01",10)	LT	MeasurementMode	1	PrivateTag
(0019,"SIEMENS DLR.01",11)	LT	ImageType	1	PrivateTag
(0019,"SIEMENS DLR.01",15)	LT	SoftwareVersion	1	PrivateTag
(0019,"SIEMENS DLR.01",20)	LT	MPMCode	1	PrivateTag
(0019,"SIEMENS DLR.01",21)	LT	Latitude	1	PrivateTag
(0019,"SIEMENS DLR.01",22)	LT	Sensitivity	1	PrivateTag
(0019,"SIEMENS DLR.01",23)	LT	EDR	1	PrivateTag
(0019,"SIEMENS DLR.01",24)	LT	LFix	1	PrivateTag
(0019,"SIEMENS DLR.01",25)	LT	SFix	1	PrivateTag
(0019,"SIEMENS DLR.01",26)	LT	PresetMode	1	PrivateTag
(0019,"SIEMENS DLR.01",27)	LT	Region	1	PrivateTag
(0019,"SIEMENS DLR.01",28)	LT	Subregion	1	PrivateTag
(0019,"SIEMENS DLR.01",30)	LT	Orientation	1	PrivateTag
(0019,"SIEMENS DLR.01",31)	LT	MarkOnFilm	1	PrivateTag
(0019,"SIEMENS DLR.01",32)	LT	RotationOnDRC	1	PrivateTag
(0019,"SIEMENS DLR.01",40)	LT	ReaderType	1	PrivateTag
(0019,"SIEMENS DLR.01",41)	LT	SubModality	1	PrivateTag
(0019,"SIEMENS DLR.01",42)	LT	ReaderSerialNumber	1	PrivateTag
(0019,"SIEMENS DLR.01",50)	LT	CassetteScale	1	PrivateTag
(0019,"SIEMENS DLR.01",51)	LT	CassetteMatrix	1	PrivateTag
(0019,"SIEMENS DLR.01",52)	LT	CassetteSubmatrix	1	PrivateTag
(0019,"SIEMENS DLR.01",53)	LT	Barcode	1	PrivateTag
(0019,"SIEMENS DLR.01",60)	LT	ContrastType	1	PrivateTag
(0019,"SIEMENS DLR.01",61)	LT	RotationAmount	1	PrivateTag
(0019,"SIEMENS DLR.01",62)	LT	RotationCenter	1	PrivateTag
(0019,"SIEMENS DLR.01",63)	LT	DensityShift	1	PrivateTag
(0019,"SIEMENS DLR.01",64)	US	FrequencyRank	1	PrivateTag
(0019,"SIEMENS DLR.01",65)	LT	FrequencyEnhancement	1	PrivateTag
(0019,"SIEMENS DLR.01",66)	LT	FrequencyType	1	PrivateTag
(0019,"SIEMENS DLR.01",67)	LT	KernelLength	1	PrivateTag
(0019,"SIEMENS DLR.01",68)	UL	KernelMode	1	PrivateTag
(0019,"SIEMENS DLR.01",69)	UL	ConvolutionMode	1	PrivateTag
(0019,"SIEMENS DLR.01",70)	LT	PLASource	1	PrivateTag
(0019,"SIEMENS DLR.01",71)	LT	PLADestination	1	PrivateTag
(0019,"SIEMENS DLR.01",75)	LT	UIDOriginalImage	1	PrivateTag
(0019,"SIEMENS DLR.01",76)	LT	Unknown	1	PrivateTag
(0019,"SIEMENS DLR.01",80)	LT	ReaderHeader	1	PrivateTag
(0019,"SIEMENS DLR.01",90)	LT	PLAOfSecondaryDestination	1	PrivateTag
(0019,"SIEMENS DLR.01",a0)	DS	Unknown	1	PrivateTag
(0019,"SIEMENS DLR.01",a1)	DS	Unknown	1	PrivateTag
(0041,"SIEMENS DLR.01",10)	US	NumberOfHardcopies	1	PrivateTag
(0041,"SIEMENS DLR.01",20)	LT	FilmFormat	1	PrivateTag
(0041,"SIEMENS DLR.01",30)	LT	FilmSize	1	PrivateTag
(0041,"SIEMENS DLR.01",31)	LT	FullFilmFormat	1	PrivateTag

(0003,"SIEMENS ISI",08)	US	ISICommandField	1	PrivateTag
(0003,"SIEMENS ISI",11)	US	AttachIDApplicationCode	1	PrivateTag
(0003,"SIEMENS ISI",12)	UL	AttachIDMessageCount	1	PrivateTag
(0003,"SIEMENS ISI",13)	DA	AttachIDDate	1	PrivateTag
(0003,"SIEMENS ISI",14)	TM	AttachIDTime	1	PrivateTag
(0003,"SIEMENS ISI",20)	US	MessageType	1	PrivateTag
(0003,"SIEMENS ISI",30)	DA	MaxWaitingDate	1	PrivateTag
(0003,"SIEMENS ISI",31)	TM	MaxWaitingTime	1	PrivateTag
(0009,"SIEMENS ISI",01)	UN	RISPatientInfoIMGEF	1	PrivateTag
(0011,"SIEMENS ISI",03)	LT	PatientUID	1	PrivateTag
(0011,"SIEMENS ISI",04)	LT	PatientID	1	PrivateTag
(0011,"SIEMENS ISI",0a)	LT	CaseID	1	PrivateTag
(0011,"SIEMENS ISI",22)	LT	RequestID	1	PrivateTag
(0011,"SIEMENS ISI",23)	LT	ExaminationUID	1	PrivateTag
(0011,"SIEMENS ISI",a1)	DA	PatientRegistrationDate	1	PrivateTag
(0011,"SIEMENS ISI",a2)	TM	PatientRegistrationTime	1	PrivateTag
(0011,"SIEMENS ISI",b0)	LT	PatientLastName	1	PrivateTag
(0011,"SIEMENS ISI",b2)	LT	PatientFirstName	1	PrivateTag
(0011,"SIEMENS ISI",b4)	LT	PatientHospitalStatus	1	PrivateTag
(0011,"SIEMENS ISI",bc)	TM	CurrentLocationTime	1	PrivateTag
(0011,"SIEMENS ISI",c0)	LT	PatientInsuranceStatus	1	PrivateTag
(0011,"SIEMENS ISI",d0)	LT	PatientBillingType	1	PrivateTag
(0011,"SIEMENS ISI",d2)	LT	PatientBillingAddress	1	PrivateTag
(0031,"SIEMENS ISI",12)	LT	ExaminationReason	1	PrivateTag
(0031,"SIEMENS ISI",30)	DA	RequestedDate	1	PrivateTag
(0031,"SIEMENS ISI",32)	TM	WorklistRequestStartTime	1	PrivateTag
(0031,"SIEMENS ISI",33)	TM	WorklistRequestEndTime	1	PrivateTag
(0031,"SIEMENS ISI",4a)	TM	RequestedTime	1	PrivateTag
(0031,"SIEMENS ISI",80)	LT	RequestedLocation	1	PrivateTag
(0055,"SIEMENS ISI",46)	LT	CurrentWard	1	PrivateTag
(0193,"SIEMENS ISI",02)	DS	RISKey	1	PrivateTag
(0307,"SIEMENS ISI",01)	UN	RISWorklistIMGEF	1	PrivateTag
(0309,"SIEMENS ISI",01)	UN	RISReportIMGEF	1	PrivateTag
(4009,"SIEMENS ISI",01)	LT	ReportID	1	PrivateTag
(4009,"SIEMENS ISI",20)	LT	ReportStatus	1	PrivateTag
(4009,"SIEMENS ISI",30)	DA	ReportCreationDate	1	PrivateTag
(4009,"SIEMENS ISI",70)	LT	ReportApprovingPhysician	1	PrivateTag
(4009,"SIEMENS ISI",e0)	LT	ReportText	1	PrivateTag
(4009,"SIEMENS ISI",e1)	LT	ReportAuthor	1	PrivateTag
(4009,"SIEMENS ISI",e3)	LT	ReportingRadiologist	1	PrivateTag

(0029,"SIEMENS MED DISPLAY",04)	LT	PhotometricInterpretation	1	PrivateTag
(0029,"SIEMENS MED DISPLAY",10)	US	RowsOfSubmatrix	1	PrivateTag
(0029,"SIEMENS MED DISPLAY",11)	US	ColumnsOfSubmatrix	1	PrivateTag
(0029,"SIEMENS MED DISPLAY",20)	US	Unknown	1	PrivateTag
(0029,"SIEMENS MED DISPLAY",21)	US	Unknown	1	PrivateTag
(0029,"SIEMENS MED DISPLAY",50)	US	OriginOfSubmatrix	1	PrivateTag
(0029,"SIEMENS MED DISPLAY",99)	LT	ShutterType	1	PrivateTag
(0029,"SIEMENS MED DISPLAY",a0)	US	RowsOfRectangularShutter	1	PrivateTag
(0029,"SIEMENS MED DISPLAY",a1)	US	ColumnsOfRectangularShutter	1	PrivateTag
(0029,"SIEMENS MED DISPLAY",a2)	US	OriginOfRectangularShutter	1	PrivateTag
(0029,"SIEMENS MED DISPLAY",b0)	US	RadiusOfCircularShutter	1	PrivateTag
(0029,"SIEMENS MED DISPLAY",b2)	US	OriginOfCircularShutter	1	PrivateTag
(0029,"SIEMENS MED DISPLAY",c1)	US	ContourOfIrregularShutter	1	PrivateTag

(0029,"SIEMENS MED HG",10)	US	ListOfGroupNumbers	1	PrivateTag
(0029,"SIEMENS MED HG",15)	LT	ListOfShadowOwnerCodes	1	PrivateTag
(0029,"SIEMENS MED HG",20)	US	ListOfElementNumbers	1	PrivateTag
(0029,"SIEMENS MED HG",30)	US	ListOfTotalDisplayLength	1	PrivateTag
(0029,"SIEMENS MED HG",40)	LT	ListOfDisplayPrefix	1	PrivateTag
(0029,"SIEMENS MED HG",50)	LT	ListOfDisplayPostfix	1	PrivateTag
(0029,"SIEMENS MED HG",60)	US	ListOfTextPosition	1	PrivateTag
(0029,"SIEMENS MED HG",70)	LT	ListOfTextConcatenation	1	PrivateTag
(0029,"SIEMENS MED MG",10)	US	ListOfGroupNumbers	1	PrivateTag
(0029,"SIEMENS MED MG",15)	LT	ListOfShadowOwnerCodes	1	PrivateTag
(0029,"SIEMENS MED MG",20)	US	ListOfElementNumbers	1	PrivateTag
(0029,"SIEMENS MED MG",30)	US	ListOfTotalDisplayLength	1	PrivateTag
(0029,"SIEMENS MED MG",40)	LT	ListOfDisplayPrefix	1	PrivateTag
(0029,"SIEMENS MED MG",50)	LT	ListOfDisplayPostfix	1	PrivateTag
(0029,"SIEMENS MED MG",60)	US	ListOfTextPosition	1	PrivateTag
(0029,"SIEMENS MED MG",70)	LT	ListOfTextConcatenation	1	PrivateTag

(0009,"SIEMENS MED",10)	LO	RecognitionCode	1	PrivateTag
(0009,"SIEMENS MED",30)	UL	ByteOffsetOfOriginalHeader	1	PrivateTag
(0009,"SIEMENS MED",31)	UL	LengthOfOriginalHeader	1	PrivateTag
(0009,"SIEMENS MED",40)	UL	ByteOffsetOfPixelmatrix	1	PrivateTag
(0009,"SIEMENS MED",41)	UL	LengthOfPixelmatrixInBytes	1	PrivateTag
(0009,"SIEMENS MED",50)	LT	Unknown	1	PrivateTag
(0009,"SIEMENS MED",51)	LT	Unknown	1	PrivateTag
(0009,"SIEMENS MED",f5)	LT	PDMEFIDPlaceholder	1	PrivateTag
(0009,"SIEMENS MED",f6)	LT	PDMDataObjectTypeExtension	1	PrivateTag
(0021,"SIEMENS MED",10)	DS	Zoom	1	PrivateTag
(0021,"SIEMENS MED",11)	DS	Target	2	PrivateTag
(0021,"SIEMENS MED",12)	IS	TubeAngle	1	PrivateTag
(0021,"SIEMENS MED",20)	US	ROIMask	1	PrivateTag
(7001,"SIEMENS MED",10)	LT	Dummy	1	PrivateTag
(7003,"SIEMENS MED",10)	LT	Header	1	PrivateTag
(7005,"SIEMENS MED",10)	LT	Dummy	1	PrivateTag

(0029,"SIEMENS MEDCOM HEADER",08)	CS	MedComHeaderType	1	PrivateTag
(0029,"SIEMENS MEDCOM HEADER",09)	LO	MedComHeaderVersion	1	PrivateTag
(0029,"SIEMENS MEDCOM HEADER",10)	OB	MedComHeaderInfo	1	PrivateTag
(0029,"SIEMENS MEDCOM HEADER",20)	OB	MedComHistoryInformation	1	PrivateTag
(0029,"SIEMENS MEDCOM HEADER",31)	LO	PMTFInformation1	1	PrivateTag
(0029,"SIEMENS MEDCOM HEADER",32)	UL	PMTFInformation2	1	PrivateTag
(0029,"SIEMENS MEDCOM HEADER",33)	UL	PMTFInformation3	1	PrivateTag
(0029,"SIEMENS MEDCOM HEADER",34)	CS	PMTFInformation4	1	PrivateTag
(0029,"SIEMENS MEDCOM HEADER",35)	UL	PMTFInformation5	1	PrivateTag
(0029,"SIEMENS MEDCOM HEADER",40)	SQ	ApplicationHeaderSequence	1	PrivateTag
(0029,"SIEMENS MEDCOM HEADER",41)	CS	ApplicationHeaderType	1	PrivateTag
(0029,"SIEMENS MEDCOM HEADER",42)	LO	ApplicationHeaderID	1	PrivateTag
(0029,"SIEMENS MEDCOM HEADER",43)	LO	ApplicationHeaderVersion	1	PrivateTag
(0029,"SIEMENS MEDCOM HEADER",44)	OB	ApplicationHeaderInfo	1	PrivateTag
(0029,"SIEMENS MEDCOM HEADER",50)	LO	WorkflowControlFlags	8	PrivateTag
(0029,"SIEMENS MEDCOM HEADER",51)	CS	ArchiveManagementFlagKeepOnline	1	PrivateTag
(0029,"SIEMENS MEDCOM HEADER",52)	CS	ArchiveManagementFlagDoNotArchive	1	PrivateTag
(0029,"SIEMENS MEDCOM HEADER",53)	CS	ImageLocationStatus	1	PrivateTag
(0029,"SIEMENS MEDCOM HEADER",54)	DS	EstimatedRetrieveTime	1	PrivateTag
(0029,"SIEMENS MEDCOM HEADER",55)	DS	DataSizeOfRetrievedImages	1	PrivateTag
(0029,"SIEMENS MEDCOM HEADER",70)	SQ	SiemensLinkSequence	1	PrivateTag
(0029,"SIEMENS MEDCOM HEADER",71)	AT	ReferencedTag	1	PrivateTag
(0029,"SIEMENS MEDCOM HEADER",72)	CS	ReferencedTagType	1	PrivateTag
(0029,"SIEMENS MEDCOM HEADER",73)	UL	ReferencedValueLength	1	PrivateTag
(0029,"SIEMENS MEDCOM HEADER",74)	CS	ReferencedObjectDeviceType	1	PrivateTag
(0029,"SIEMENS MEDCOM HEADER",75)	OB	ReferencedObjectDeviceLocation	1	PrivateTag
(0029,"SIEMENS MEDCOM HEADER",76)	OB	ReferencedObjectDeviceID	1	PrivateTag

(0029,"SIEMENS MEDCOM HEADER2",60)	LO	SeriesWorkflowStatus	1	PrivateTag

(0029,"SIEMENS MEDCOM OOG",08)	CS	MEDCOMOOGType	1	PrivateTag
(0029,"SIEMENS MEDCOM OOG",09)	LO	MEDCOMOOGVersion	1	PrivateTag
(0029,"SIEMENS MEDCOM OOG",10)	OB	MEDCOMOOGInfo	1	PrivateTag

(0019,"SIEMENS MR VA0  COAD",12)	DS	MagneticFieldStrength	1	PrivateTag
(0019,"SIEMENS MR VA0  COAD",14)	DS	ADCVoltage	1	PrivateTag
(0019,"SIEMENS MR VA0  COAD",16)	DS	ADCOffset	2	PrivateTag
(0019,"SIEMENS MR VA0  COAD",20)	DS	TransmitterAmplitude	1	PrivateTag
(0019,"SIEMENS MR VA0  COAD",21)	IS	NumberOfTransmitterAmplitudes	1	PrivateTag
(0019,"SIEMENS MR VA0  COAD",22)	DS	TransmitterAttenuator	1	PrivateTag
(0019,"SIEMENS MR VA0  COAD",24)	DS	TransmitterCalibration	1	PrivateTag
(0019,"SIEMENS MR VA0  COAD",26)	DS	TransmitterReference	1	PrivateTag
(0019,"SIEMENS MR VA0  COAD",50)	DS	ReceiverTotalGain	1	PrivateTag
(0019,"SIEMENS MR VA0  COAD",51)	DS	ReceiverAmplifierGain	1	PrivateTag
(0019,"SIEMENS MR VA0  COAD",52)	DS	ReceiverPreamplifierGain	1	PrivateTag
(0019,"SIEMENS MR VA0  COAD",54)	DS	ReceiverCableAttenuation	1	PrivateTag
(0019,"SIEMENS MR VA0  COAD",55)	DS	ReceiverReferenceGain	1	PrivateTag
(0019,"SIEMENS MR VA0  COAD",56)	DS	ReceiverFilterFrequency	1	PrivateTag
(0019,"SIEMENS MR VA0  COAD",60)	DS	ReconstructionScaleFactor	1	PrivateTag
(0019,"SIEMENS MR VA0  COAD",62)	DS	ReferenceScaleFactor	1	PrivateTag
(0019,"SIEMENS MR VA0  COAD",70)	DS	PhaseGradientAmplitude	1	PrivateTag
(0019,"SIEMENS MR VA0  COAD",71)	DS	ReadoutGradientAmplitude	1	PrivateTag
(0019,"SIEMENS MR VA0  COAD",72)	DS	SelectionGradientAmplitude	1	PrivateTag
(0019,"SIEMENS MR VA0  COAD",80)	DS	GradientDelayTime	3	PrivateTag
(0019,"SIEMENS MR VA0  COAD",82)	DS	TotalGradientDelayTime	1	PrivateTag
(0019,"SIEMENS MR VA0  COAD",90)	LT	SensitivityCorrectionLabel	1	PrivateTag
(0019,"SIEMENS MR VA0  COAD",91)	DS	SaturationPhaseEncodingVectorCoronalComponent	6	PrivateTag
(0019,"SIEMENS MR VA0  COAD",92)	DS	SaturationReadoutVectorCoronalComponent	6	PrivateTag
(0019,"SIEMENS MR VA0  COAD",a0)	US	RFWatchdogMask	3	PrivateTag
(0019,"SIEMENS MR VA0  COAD",a1)	DS	EPIReconstructionSlope	1	PrivateTag
(0019,"SIEMENS MR VA0  COAD",a2)	DS	RFPowerErrorIndicator	1	PrivateTag
(0019,"SIEMENS MR VA0  COAD",a5)	DS	SpecificAbsorptionRateWholeBody	3	PrivateTag
(0019,"SIEMENS MR VA0  COAD",a6)	DS	SpecificEnergyDose	3	PrivateTag
(0019,"SIEMENS MR VA0  COAD",b0)	UL	AdjustmentStatusMask	1	PrivateTag
(0019,"SIEMENS MR VA0  COAD",c1)	DS	EPICapacity	6	PrivateTag
(0019,"SIEMENS MR VA0  COAD",c2)	DS	EPIInductance	3	PrivateTag
(0019,"SIEMENS MR VA0  COAD",c3)	IS	EPISwitchConfigurationCode	1-n	PrivateTag
(0019,"SIEMENS MR VA0  COAD",c4)	IS	EPISwitchHardwareCode	1-n	PrivateTag
(0019,"SIEMENS MR VA0  COAD",c5)	DS	EPISwitchDelayTime	1-n	PrivateTag
(0019,"SIEMENS MR VA0  COAD",d1)	DS	FlowSensitivity	1	PrivateTag
(0019,"SIEMENS MR VA0  COAD",d2)	LT	CalculationSubmode	1	PrivateTag
(0019,"SIEMENS MR VA0  COAD",d3)	DS	FieldOfViewRatio	1	PrivateTag
(0019,"SIEMENS MR VA0  COAD",d4)	IS	BaseRawMatrixSize	1	PrivateTag
(0019,"SIEMENS MR VA0  COAD",d5)	IS	2DOversamplingLines	1	PrivateTag
(0019,"SIEMENS MR VA0  COAD",d6)	IS	3DPhaseOversamplingPartitions	1	PrivateTag
(0019,"SIEMENS MR VA0  COAD",d7)	IS	EchoLinePosition	1	PrivateTag
(0019,"SIEMENS MR VA0  COAD",d8)	IS	EchoColumnPosition	1	PrivateTag
(0019,"SIEMENS MR VA0  COAD",d9)	IS	LinesPerSegment	1	PrivateTag
(0019,"SIEMENS MR VA0  COAD",da)	LT	PhaseCodingDirection	1	PrivateTag

(0019,"SIEMENS MR VA0  GEN",10)	DS	TotalMeasurementTimeNominal	1	PrivateTag
(0019,"SIEMENS MR VA0  GEN",11)	DS	TotalMeasurementTimeCurrent	1	PrivateTag
(0019,"SIEMENS MR VA0  GEN",12)	DS	StartDelayTime	1	PrivateTag
(0019,"SIEMENS MR VA0  GEN",13)	DS	DwellTime	1	PrivateTag
(0019,"SIEMENS MR VA0  GEN",14)	IS	NumberOfPhases	1	PrivateTag
(0019,"SIEMENS MR VA0  GEN",16)	UL	SequenceControlMask	2	PrivateTag
(0019,"SIEMENS MR VA0  GEN",18)	UL	MeasurementStatusMask	1	PrivateTag
(0019,"SIEMENS MR VA0  GEN",20)	IS	NumberOfFourierLinesNominal	1	PrivateTag
(0019,"SIEMENS MR VA0  GEN",21)	IS	NumberOfFourierLinesCurrent	1	PrivateTag
(0019,"SIEMENS MR VA0  GEN",26)	IS	NumberOfFourierLinesAfterZero	1	PrivateTag
(0019,"SIEMENS MR VA0  GEN",28)	IS	FirstMeasuredFourierLine	1	PrivateTag
(0019,"SIEMENS MR VA0  GEN",30)	IS	AcquisitionColumns	1	PrivateTag
(0019,"SIEMENS MR VA0  GEN",31)	IS	ReconstructionColumns	1	PrivateTag
(0019,"SIEMENS MR VA0  GEN",40)	IS	ArrayCoilElementNumber	1	PrivateTag
(0019,"SIEMENS MR VA0  GEN",41)	UL	ArrayCoilElementSelectMask	1	PrivateTag
(0019,"SIEMENS MR VA0  GEN",42)	UL	ArrayCoilElementDataMask	1	PrivateTag
(0019,"SIEMENS MR VA0  GEN",43)	IS	ArrayCoilElementToADCConnect	1-n	PrivateTag
(0019,"SIEMENS MR VA0  GEN",44)	DS	ArrayCoilElementNoiseLevel	1-n	PrivateTag
(0019,"SIEMENS MR VA0  GEN",45)	IS	ArrayCoilADCPairNumber	1	PrivateTag
(0019,"SIEMENS MR VA0  GEN",46)	UL	ArrayCoilCombinationMask	1	PrivateTag
(0019,"SIEMENS MR VA0  GEN",50)	IS	NumberOfAverages	1	PrivateTag
(0019,"SIEMENS MR VA0  GEN",60)	DS	FlipAngle	1	PrivateTag
(0019,"SIEMENS MR VA0  GEN",70)	IS	NumberOfPrescans	1	PrivateTag
(0019,"SIEMENS MR VA0  GEN",81)	LT	FilterTypeForRawData	1	PrivateTag
(0019,"SIEMENS MR VA0  GEN",82)	DS	FilterParameterForRawData	1-n	PrivateTag
(0019,"SIEMENS MR VA0  GEN",83)	LT	FilterTypeForImageData	1	PrivateTag
(0019,"SIEMENS MR VA0  GEN",84)	DS	FilterParameterForImageData	1-n	PrivateTag
(0019,"SIEMENS MR VA0  GEN",85)	LT	FilterTypeForPhaseCorrection	1	PrivateTag
(0019,"SIEMENS MR VA0  GEN",86)	DS	FilterParameterForPhaseCorrection	1-n	PrivateTag
(0019,"SIEMENS MR VA0  GEN",87)	LT	NormalizationFilterTypeForImageData	1	PrivateTag
(0019,"SIEMENS MR VA0  GEN",88)	DS	NormalizationFilterParameterForImageData	1-n	PrivateTag
(0019,"SIEMENS MR VA0  GEN",90)	IS	NumberOfSaturationRegions	1	PrivateTag
(0019,"SIEMENS MR VA0  GEN",91)	DS	SaturationPhaseEncodingVectorSagittalComponent	6	PrivateTag
(0019,"SIEMENS MR VA0  GEN",92)	DS	SaturationReadoutVectorSagittalComponent	6	PrivateTag
(0019,"SIEMENS MR VA0  GEN",93)	DS	EPIStimulationMonitorMode	1	PrivateTag
(0019,"SIEMENS MR VA0  GEN",94)	DS	ImageRotationAngle	1	PrivateTag
(0019,"SIEMENS MR VA0  GEN",96)	UL	CoilIDMask	3	PrivateTag
(0019,"SIEMENS MR VA0  GEN",97)	UL	CoilClassMask	2	PrivateTag
(0019,"SIEMENS MR VA0  GEN",98)	DS	CoilPosition	3	PrivateTag
(0019,"SIEMENS MR VA0  GEN",a0)	DS	EPIReconstructionPhase	1	PrivateTag
(0019,"SIEMENS MR VA0  GEN",a1)	DS	EPIReconstructionSlope	1	PrivateTag
(0021,"SIEMENS MR VA0  GEN",20)	IS	PhaseCorrectionRowsSequence	1	PrivateTag
(0021,"SIEMENS MR VA0  GEN",21)	IS	PhaseCorrectionColumnsSequence	1	PrivateTag
(0021,"SIEMENS MR VA0  GEN",22)	IS	PhaseCorrectionRowsReconstruction	1	PrivateTag
(0021,"SIEMENS MR VA0  GEN",24)	IS	PhaseCorrectionColumnsReconstruction	1	PrivateTag
(0021,"SIEMENS MR VA0  GEN",30)	IS	NumberOf3DRawPartitionsNominal	1	PrivateTag
(0021,"SIEMENS MR VA0  GEN",31)	IS	NumberOf3DRawPartitionsCurrent	1	PrivateTag
(0021,"SIEMENS MR VA0  GEN",34)	IS	NumberOf3DImagePartitions	1	PrivateTag
(0021,"SIEMENS MR VA0  GEN",36)	IS	Actual3DImagePartitionNumber	1	PrivateTag
(0021,"SIEMENS MR VA0  GEN",39)	DS	SlabThickness	1	PrivateTag
(0021,"SIEMENS MR VA0  GEN",40)	IS	NumberOfSlicesNominal	1	PrivateTag
(0021,"SIEMENS MR VA0  GEN",41)	IS	NumberOfSlicesCurrent	1	PrivateTag
(0021,"SIEMENS MR VA0  GEN",42)	IS	CurrentSliceNumber	1	PrivateTag
(0021,"SIEMENS MR VA0  GEN",43)	IS	CurrentGroupNumber	1	PrivateTag
(0021,"SIEMENS MR VA0  GEN",44)	DS	CurrentSliceDistanceFactor	1	PrivateTag
(0021,"SIEMENS MR VA0  GEN",45)	IS	MIPStartRow	1	PrivateTag
(0021,"SIEMENS MR VA0  GEN",46)	IS	MIPStopRow	1	PrivateTag
(0021,"SIEMENS MR VA0  GEN",47)	IS	MIPStartColumn	1	PrivateTag
(0021,"SIEMENS MR VA0  GEN",48)	IS	MIPStartColumn	1	PrivateTag
(0021,"SIEMENS MR VA0  GEN",49)	IS	MIPStartSlice Name=	1	PrivateTag
(0021,"SIEMENS MR VA0  GEN",4a)	IS	MIPStartSlice	1	PrivateTag
(0021,"SIEMENS MR VA0  GEN",4f)	LT	OrderofSlices	1	PrivateTag
(0021,"SIEMENS MR VA0  GEN",50)	US	SignalMask	1	PrivateTag
(0021,"SIEMENS MR VA0  GEN",52)	DS	DelayAfterTrigger	1	PrivateTag
(0021,"SIEMENS MR VA0  GEN",53)	IS	RRInterval	1	PrivateTag
(0021,"SIEMENS MR VA0  GEN",54)	DS	NumberOfTriggerPulses	1	PrivateTag
(0021,"SIEMENS MR VA0  GEN",56)	DS	RepetitionTimeEffective	1	PrivateTag
(0021,"SIEMENS MR VA0  GEN",57)	LT	GatePhase	1	PrivateTag
(0021,"SIEMENS MR VA0  GEN",58)	DS	GateThreshold	1	PrivateTag
(0021,"SIEMENS MR VA0  GEN",59)	DS	GatedRatio	1	PrivateTag
(0021,"SIEMENS MR VA0  GEN",60)	IS	NumberOfInterpolatedImages	1	PrivateTag
(0021,"SIEMENS MR VA0  GEN",70)	IS	NumberOfEchoes	1	PrivateTag
(0021,"SIEMENS MR VA0  GEN",72)	DS	SecondEchoTime	1	PrivateTag
(0021,"SIEMENS MR VA0  GEN",73)	DS	SecondRepetitionTime	1	PrivateTag
(0021,"SIEMENS MR VA0  GEN",80)	IS	CardiacCode	1	PrivateTag
(0021,"SIEMENS MR VA0  GEN",91)	DS	SaturationPhaseEncodingVectorTransverseComponent	6	PrivateTag
(0021,"SIEMENS MR VA0  GEN",92)	DS	SaturationReadoutVectorTransverseComponent	6	PrivateTag
(0021,"SIEMENS MR VA0  GEN",93)	DS	EPIChangeValueOfMagnitude	1	PrivateTag
(0021,"SIEMENS MR VA0  GEN",94)	DS	EPIChangeValueOfXComponent	1	PrivateTag
(0021,"SIEMENS MR VA0  GEN",95)	DS	EPIChangeValueOfYComponent	1	PrivateTag
(0021,"SIEMENS MR VA0  GEN",96)	DS	EPIChangeValueOfZComponent	1	PrivateTag

(0021,"SIEMENS MR VA0  RAW",00)	LT	SequenceType	1	PrivateTag
(0021,"SIEMENS MR VA0  RAW",01)	IS	VectorSizeOriginal	1	PrivateTag
(0021,"SIEMENS MR VA0  RAW",02)	IS	VectorSizeExtended	1	PrivateTag
(0021,"SIEMENS MR VA0  RAW",03)	DS	AcquiredSpectralRange	1	PrivateTag
(0021,"SIEMENS MR VA0  RAW",04)	DS	VOIPosition	3	PrivateTag
(0021,"SIEMENS MR VA0  RAW",05)	DS	VOISize	3	PrivateTag
(0021,"SIEMENS MR VA0  RAW",06)	IS	CSIMatrixSizeOriginal	3	PrivateTag
(0021,"SIEMENS MR VA0  RAW",07)	IS	CSIMatrixSizeExtended	3	PrivateTag
(0021,"SIEMENS MR VA0  RAW",08)	DS	SpatialGridShift	3	PrivateTag
(0021,"SIEMENS MR VA0  RAW",09)	DS	SignalLimitsMinimum	1	PrivateTag
(0021,"SIEMENS MR VA0  RAW",10)	DS	SignalLimitsMaximum	1	PrivateTag
(0021,"SIEMENS MR VA0  RAW",11)	DS	SpecInfoMask	1	PrivateTag
(0021,"SIEMENS MR VA0  RAW",12)	DS	EPITimeRateOfChangeOfMagnitude	1	PrivateTag
(0021,"SIEMENS MR VA0  RAW",13)	DS	EPITimeRateOfChangeOfXComponent	1	PrivateTag
(0021,"SIEMENS MR VA0  RAW",14)	DS	EPITimeRateOfChangeOfYComponent	1	PrivateTag
(0021,"SIEMENS MR VA0  RAW",15)	DS	EPITimeRateOfChangeOfZComponent	1	PrivateTag
(0021,"SIEMENS MR VA0  RAW",16)	DS	EPITimeRateOfChangeLegalLimit1	1	PrivateTag
(0021,"SIEMENS MR VA0  RAW",17)	DS	EPIOperationModeFlag	1	PrivateTag
(0021,"SIEMENS MR VA0  RAW",18)	DS	EPIFieldCalculationSafetyFactor	1	PrivateTag
(0021,"SIEMENS MR VA0  RAW",19)	DS	EPILegalLimit1OfChangeValue	1	PrivateTag
(0021,"SIEMENS MR VA0  RAW",20)	DS	EPILegalLimit2OfChangeValue	1	PrivateTag
(0021,"SIEMENS MR VA0  RAW",21)	DS	EPIRiseTime	1	PrivateTag
(0021,"SIEMENS MR VA0  RAW",30)	DS	ArrayCoilADCOffset	16	PrivateTag
(0021,"SIEMENS MR VA0  RAW",31)	DS	ArrayCoilPreamplifierGain	16	PrivateTag
(0021,"SIEMENS MR VA0  RAW",50)	LT	SaturationType	1	PrivateTag
(0021,"SIEMENS MR VA0  RAW",51)	DS	SaturationNormalVector	3	PrivateTag
(0021,"SIEMENS MR VA0  RAW",52)	DS	SaturationPositionVector	3	PrivateTag
(0021,"SIEMENS MR VA0  RAW",53)	DS	SaturationThickness	6	PrivateTag
(0021,"SIEMENS MR VA0  RAW",54)	DS	SaturationWidth	6	PrivateTag
(0021,"SIEMENS MR VA0  RAW",55)	DS	SaturationDistance	6	PrivateTag

(7fe3,"SIEMENS NUMARIS II",00)	LT	ImageGraphicsFormatCode	1	PrivateTag
(7fe3,"SIEMENS NUMARIS II",10)	OB	ImageGraphics	1	PrivateTag
(7fe3,"SIEMENS NUMARIS II",20)	OB	ImageGraphicsDummy	1	PrivateTag

(0011,"SIEMENS RA GEN",20)	SL	FluoroTimer	1	PrivateTag
(0011,"SIEMENS RA GEN",25)	SL	PtopDoseAreaProduct	1	PrivateTag
(0011,"SIEMENS RA GEN",26)	SL	PtopTotalSkinDose	1	PrivateTag
(0011,"SIEMENS RA GEN",30)	LT	Unknown	1	PrivateTag
(0011,"SIEMENS RA GEN",35)	LO	PatientInitialPuckCounter	1	PrivateTag
(0011,"SIEMENS RA GEN",40)	SS	SPIDataObjectType	1	PrivateTag
(0019,"SIEMENS RA GEN",15)	LO	AcquiredPlane	1	PrivateTag
(0019,"SIEMENS RA GEN",1f)	SS	DefaultTableIsoCenterHeight	1	PrivateTag
(0019,"SIEMENS RA GEN",20)	SL	SceneFlag	1	PrivateTag
(0019,"SIEMENS RA GEN",22)	SL	RefPhotofileFlag	1	PrivateTag
(0019,"SIEMENS RA GEN",24)	LO	SceneName	1	PrivateTag
(0019,"SIEMENS RA GEN",26)	SS	AcquisitionIndex	1	PrivateTag
(0019,"SIEMENS RA GEN",28)	SS	MixedPulseMode	1	PrivateTag
(0019,"SIEMENS RA GEN",2a)	SS	NoOfPositions	1	PrivateTag
(0019,"SIEMENS RA GEN",2c)	SS	NoOfPhases	1	PrivateTag
(0019,"SIEMENS RA GEN",2e)	SS	FrameRateForPositions	1-n	PrivateTag
(0019,"SIEMENS RA GEN",30)	SS	NoOfFramesForPositions	1-n	PrivateTag
(0019,"SIEMENS RA GEN",32)	SS	SteppingDirection	1	PrivateTag
(0019,"SIEMENS RA GEN",34)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA GEN",36)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA GEN",38)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA GEN",3a)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA GEN",3c)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA GEN",3e)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA GEN",40)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA GEN",42)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA GEN",44)	SS	ImageTransferDelay	1	PrivateTag
(0019,"SIEMENS RA GEN",46)	SL	InversFlag	1	PrivateTag
(0019,"SIEMENS RA GEN",48)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA GEN",4a)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA GEN",4c)	SS	BlankingCircleDiameter	1	PrivateTag
(0019,"SIEMENS RA GEN",50)	SL	StandDataValid	1	PrivateTag
(0019,"SIEMENS RA GEN",52)	SS	TableTilt	1	PrivateTag
(0019,"SIEMENS RA GEN",54)	SS	TableAxisRotation	1	PrivateTag
(0019,"SIEMENS RA GEN",56)	SS	TableLongitudalPosition	1	PrivateTag
(0019,"SIEMENS RA GEN",58)	SS	TableSideOffset	1	PrivateTag
(0019,"SIEMENS RA GEN",5a)	SS	TableIsoCenterHeight	1	PrivateTag
(0019,"SIEMENS RA GEN",5c)	UN	Unknown	1	PrivateTag
(0019,"SIEMENS RA GEN",5e)	SL	CollimationDataValid	1	PrivateTag
(0019,"SIEMENS RA GEN",60)	SL	PeriSequenceNo	1	PrivateTag
(0019,"SIEMENS RA GEN",62)	SL	PeriTotalScenes	1	PrivateTag
(0019,"SIEMENS RA GEN",64)	SL	PeriOverlapTop	1	PrivateTag
(0019,"SIEMENS RA GEN",66)	SL	PeriOverlapBottom	1	PrivateTag
(0019,"SIEMENS RA GEN",68)	SL	RawImageNumber	1	PrivateTag
(0019,"SIEMENS RA GEN",6a)	SL	XRayDataValid	1	PrivateTag
(0019,"SIEMENS RA GEN",70)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS RA GEN",72)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS RA GEN",74)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS RA GEN",76)	SL	FillingAverageFactor	1	PrivateTag
(0019,"SIEMENS RA GEN",78)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS RA GEN",7a)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS RA GEN",7c)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS RA GEN",7e)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS RA GEN",80)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS RA GEN",82)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS RA GEN",84)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS RA GEN",86)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS RA GEN",88)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS RA GEN",8a)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS RA GEN",8c)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS RA GEN",8e)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS RA GEN",92)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS RA GEN",94)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS RA GEN",96)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS RA GEN",98)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS RA GEN",9a)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS RA GEN",9c)	SL	IntensifierLevelCalibrationFactor	1	PrivateTag
(0019,"SIEMENS RA GEN",9e)	SL	NativeReviewFlag	1	PrivateTag
(0019,"SIEMENS RA GEN",a2)	SL	SceneNumber	1	PrivateTag
(0019,"SIEMENS RA GEN",a4)	SS	AcquisitionMode	1	PrivateTag
(0019,"SIEMENS RA GEN",a5)	SS	AcquisitonFrameRate	1	PrivateTag
(0019,"SIEMENS RA GEN",a6)	SL	ECGFlag	1	PrivateTag
(0019,"SIEMENS RA GEN",a7)	SL	AdditionalSceneData	1	PrivateTag
(0019,"SIEMENS RA GEN",a8)	SL	FileCopyFlag	1	PrivateTag
(0019,"SIEMENS RA GEN",a9)	SL	PhlebovisionFlag	1	PrivateTag
(0019,"SIEMENS RA GEN",aa)	SL	Co2Flag	1	PrivateTag
(0019,"SIEMENS RA GEN",ab)	SS	MaxSpeed	1	PrivateTag
(0019,"SIEMENS RA GEN",ac)	SS	StepWidth	1	PrivateTag
(0019,"SIEMENS RA GEN",ad)	SL	DigitalAcquisitionZoom	1	PrivateTag
(0019,"SIEMENS RA GEN",ff)	SS	Internal	1-n	PrivateTag
(0021,"SIEMENS RA GEN",15)	SS	ImagesInStudy	1	PrivateTag
(0021,"SIEMENS RA GEN",20)	SS	ScenesInStudy	1	PrivateTag
(0021,"SIEMENS RA GEN",25)	SS	ImagesInPhotofile	1	PrivateTag
(0021,"SIEMENS RA GEN",27)	SS	PlaneBImagesExist	1	PrivateTag
(0021,"SIEMENS RA GEN",28)	SS	NoOf2MBChunks	1	PrivateTag
(0021,"SIEMENS RA GEN",30)	SS	ImagesInAllScenes	1	PrivateTag
(0021,"SIEMENS RA GEN",40)	SS	ArchiveSWInternalVersion	1	PrivateTag

(0011,"SIEMENS RA PLANE A",28)	SL	FluoroTimerA	1	PrivateTag
(0011,"SIEMENS RA PLANE A",29)	SL	FluoroSkinDoseA	1	PrivateTag
(0011,"SIEMENS RA PLANE A",2a)	SL	TotalSkinDoseA	1	PrivateTag
(0011,"SIEMENS RA PLANE A",2b)	SL	FluoroDoseAreaProductA	1	PrivateTag
(0011,"SIEMENS RA PLANE A",2c)	SL	TotalDoseAreaProductA	1	PrivateTag
(0019,"SIEMENS RA PLANE A",15)	LT	OfflineUID	1	PrivateTag
(0019,"SIEMENS RA PLANE A",18)	SS	Internal	1	PrivateTag
(0019,"SIEMENS RA PLANE A",19)	SS	Internal	1	PrivateTag
(0019,"SIEMENS RA PLANE A",1a)	SS	Internal	1	PrivateTag
(0019,"SIEMENS RA PLANE A",1b)	SS	Internal	1	PrivateTag
(0019,"SIEMENS RA PLANE A",1c)	SS	Internal	1	PrivateTag
(0019,"SIEMENS RA PLANE A",1d)	SS	Internal	1	PrivateTag
(0019,"SIEMENS RA PLANE A",1e)	SS	Internal	1	PrivateTag
(0019,"SIEMENS RA PLANE A",1f)	SS	Internal	1-n	PrivateTag
(0019,"SIEMENS RA PLANE A",20)	SS	SystemCalibFactorPlaneA	1	PrivateTag
(0019,"SIEMENS RA PLANE A",22)	SS	XRayParameterSetNo	1	PrivateTag
(0019,"SIEMENS RA PLANE A",24)	SS	XRaySystem	1	PrivateTag
(0019,"SIEMENS RA PLANE A",26)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE A",28)	SS	AcquiredDisplayMode	1	PrivateTag
(0019,"SIEMENS RA PLANE A",2a)	SS	AcquisitionDelay	1	PrivateTag
(0019,"SIEMENS RA PLANE A",2c)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE A",2e)	SS	MaxFramesLimit	1	PrivateTag
(0019,"SIEMENS RA PLANE A",30)	US	MaximumFrameSizeNIU	1	PrivateTag
(0019,"SIEMENS RA PLANE A",32)	SS	SubtractedFilterType	1	PrivateTag
(0019,"SIEMENS RA PLANE A",34)	SS	FilterFactorNative	1	PrivateTag
(0019,"SIEMENS RA PLANE A",36)	SS	AnatomicBackgroundFactor	1	PrivateTag
(0019,"SIEMENS RA PLANE A",38)	SS	WindowUpperLimitNative	1	PrivateTag
(0019,"SIEMENS RA PLANE A",3a)	SS	WindowLowerLimitNative	1	PrivateTag
(0019,"SIEMENS RA PLANE A",3c)	SS	WindowBrightnessPhase1	1	PrivateTag
(0019,"SIEMENS RA PLANE A",3e)	SS	WindowBrightnessPhase2	1	PrivateTag
(0019,"SIEMENS RA PLANE A",40)	SS	WindowContrastPhase1	1	PrivateTag
(0019,"SIEMENS RA PLANE A",42)	SS	WindowContrastPhase2	1	PrivateTag
(0019,"SIEMENS RA PLANE A",44)	SS	FilterFactorSub	1	PrivateTag
(0019,"SIEMENS RA PLANE A",46)	SS	PeakOpacified	1	PrivateTag
(0019,"SIEMENS RA PLANE A",48)	SL	MaskFrame	1	PrivateTag
(0019,"SIEMENS RA PLANE A",4a)	SL	BIHFrame	1	PrivateTag
(0019,"SIEMENS RA PLANE A",4c)	SS	CentBeamAngulationCaudCran	1	PrivateTag
(0019,"SIEMENS RA PLANE A",4e)	SS	CentBeamAngulationLRAnterior	1	PrivateTag
(0019,"SIEMENS RA PLANE A",50)	SS	LongitudinalPosition	1	PrivateTag
(0019,"SIEMENS RA PLANE A",52)	SS	SideOffset	1	PrivateTag
(0019,"SIEMENS RA PLANE A",54)	SS	IsoCenterHeight	1	PrivateTag
(0019,"SIEMENS RA PLANE A",56)	SS	ImageTwist	1	PrivateTag
(0019,"SIEMENS RA PLANE A",58)	SS	SourceImageDistance	1	PrivateTag
(0019,"SIEMENS RA PLANE A",5a)	SS	MechanicalMagnificationFactor	1	PrivateTag
(0019,"SIEMENS RA PLANE A",5c)	SL	CalibrationFlag	1	PrivateTag
(0019,"SIEMENS RA PLANE A",5e)	SL	CalibrationAngleCranCaud	1	PrivateTag
(0019,"SIEMENS RA PLANE A",60)	SL	CalibrationAngleRAOLAO	1	PrivateTag
(0019,"SIEMENS RA PLANE A",62)	SL	CalibrationTableToFloorDist	1	PrivateTag
(0019,"SIEMENS RA PLANE A",64)	SL	CalibrationIsocenterToFloorDist	1	PrivateTag
(0019,"SIEMENS RA PLANE A",66)	SL	CalibrationIsocenterToSourceDist	1	PrivateTag
(0019,"SIEMENS RA PLANE A",68)	SL	CalibrationSourceToII	1	PrivateTag
(0019,"SIEMENS RA PLANE A",6a)	SL	CalibrationIIZoom	1	PrivateTag
(0019,"SIEMENS RA PLANE A",6c)	SL	CalibrationIIField	1	PrivateTag
(0019,"SIEMENS RA PLANE A",6e)	SL	CalibrationFactor	1	PrivateTag
(0019,"SIEMENS RA PLANE A",70)	SL	CalibrationObjectToImageDistance	1	PrivateTag
(0019,"SIEMENS RA PLANE A",72)	SL	CalibrationSystemFactor	1-n	PrivateTag
(0019,"SIEMENS RA PLANE A",74)	SL	CalibrationSystemCorrection	1-n	PrivateTag
(0019,"SIEMENS RA PLANE A",76)	SL	CalibrationSystemIIFormats	1-n	PrivateTag
(0019,"SIEMENS RA PLANE A",78)	SL	CalibrationGantryDataValid	1	PrivateTag
(0019,"SIEMENS RA PLANE A",7a)	SS	CollimatorSquareBreadth	1	PrivateTag
(0019,"SIEMENS RA PLANE A",7c)	SS	CollimatorSquareHeight	1	PrivateTag
(0019,"SIEMENS RA PLANE A",7e)	SS	CollimatorSquareDiameter	1	PrivateTag
(0019,"SIEMENS RA PLANE A",80)	SS	CollimaterFingerTurnAngle	1	PrivateTag
(0019,"SIEMENS RA PLANE A",82)	SS	CollimaterFingerPosition	1	PrivateTag
(0019,"SIEMENS RA PLANE A",84)	SS	CollimaterDiaphragmTurnAngle	1	PrivateTag
(0019,"SIEMENS RA PLANE A",86)	SS	CollimaterDiaphragmPosition1	1	PrivateTag
(0019,"SIEMENS RA PLANE A",88)	SS	CollimaterDiaphragmPosition2	1	PrivateTag
(0019,"SIEMENS RA PLANE A",8a)	SS	CollimaterDiaphragmMode	1	PrivateTag
(0019,"SIEMENS RA PLANE A",8c)	SS	CollimaterBeamLimitBreadth	1	PrivateTag
(0019,"SIEMENS RA PLANE A",8e)	SS	CollimaterBeamLimitHeight	1	PrivateTag
(0019,"SIEMENS RA PLANE A",90)	SS	CollimaterBeamLimitDiameter	1	PrivateTag
(0019,"SIEMENS RA PLANE A",92)	SS	X-RayControlMOde	1	PrivateTag
(0019,"SIEMENS RA PLANE A",94)	SS	X-RaySystem	1	PrivateTag
(0019,"SIEMENS RA PLANE A",96)	SS	FocalSpot	1	PrivateTag
(0019,"SIEMENS RA PLANE A",98)	SS	ExposureControl	1	PrivateTag
(0019,"SIEMENS RA PLANE A",9a)	SL	XRayVoltage	1	PrivateTag
(0019,"SIEMENS RA PLANE A",9c)	SL	XRayCurrent	1	PrivateTag
(0019,"SIEMENS RA PLANE A",9e)	SL	XRayCurrentTimeProduct	1	PrivateTag
(0019,"SIEMENS RA PLANE A",a0)	SL	XRayPulseTime	1	PrivateTag
(0019,"SIEMENS RA PLANE A",a2)	SL	XRaySceneTimeFluoroClock	1	PrivateTag
(0019,"SIEMENS RA PLANE A",a4)	SS	MaximumPulseRate	1	PrivateTag
(0019,"SIEMENS RA PLANE A",a6)	SS	PulsesPerScene	1	PrivateTag
(0019,"SIEMENS RA PLANE A",a8)	SL	DoseAreaProductOfScene	1	PrivateTag
(0019,"SIEMENS RA PLANE A",aa)	SS	Dose	1	PrivateTag
(0019,"SIEMENS RA PLANE A",ac)	SS	DoseRate	1	PrivateTag
(0019,"SIEMENS RA PLANE A",ae)	SL	IIToCoverDistance	1	PrivateTag
(0019,"SIEMENS RA PLANE A",b0)	SS	LastFramePhase1	1	PrivateTag
(0019,"SIEMENS RA PLANE A",b1)	SS	FrameRatePhase1	1	PrivateTag
(0019,"SIEMENS RA PLANE A",b2)	SS	LastFramePhase2	1	PrivateTag
(0019,"SIEMENS RA PLANE A",b3)	SS	FrameRatePhase2	1	PrivateTag
(0019,"SIEMENS RA PLANE A",b4)	SS	LastFramePhase3	1	PrivateTag
(0019,"SIEMENS RA PLANE A",b5)	SS	FrameRatePhase3	1	PrivateTag
(0019,"SIEMENS RA PLANE A",b6)	SS	LastFramePhase4	1	PrivateTag
(0019,"SIEMENS RA PLANE A",b7)	SS	FrameRatePhase4	1	PrivateTag
(0019,"SIEMENS RA PLANE A",b8)	SS	GammaOfNativeImage	1	PrivateTag
(0019,"SIEMENS RA PLANE A",b9)	SS	GammaOfTVSystem	1	PrivateTag
(0019,"SIEMENS RA PLANE A",bb)	SL	PixelshiftX	1	PrivateTag
(0019,"SIEMENS RA PLANE A",bc)	SL	PixelshiftY	1	PrivateTag
(0019,"SIEMENS RA PLANE A",bd)	SL	MaskAverageFactor	1	PrivateTag
(0019,"SIEMENS RA PLANE A",be)	SL	BlankingCircleFlag	1	PrivateTag
(0019,"SIEMENS RA PLANE A",bf)	SL	CircleRowStart	1	PrivateTag
(0019,"SIEMENS RA PLANE A",c0)	SL	CircleRowEnd	1	PrivateTag
(0019,"SIEMENS RA PLANE A",c1)	SL	CircleColumnStart	1	PrivateTag
(0019,"SIEMENS RA PLANE A",c2)	SL	CircleColumnEnd	1	PrivateTag
(0019,"SIEMENS RA PLANE A",c3)	SL	CircleDiameter	1	PrivateTag
(0019,"SIEMENS RA PLANE A",c4)	SL	RectangularCollimaterFlag	1	PrivateTag
(0019,"SIEMENS RA PLANE A",c5)	SL	RectangleRowStart	1	PrivateTag
(0019,"SIEMENS RA PLANE A",c6)	SL	RectangleRowEnd	1	PrivateTag
(0019,"SIEMENS RA PLANE A",c7)	SL	RectangleColumnStart	1	PrivateTag
(0019,"SIEMENS RA PLANE A",c8)	SL	RectangleColumnEnd	1	PrivateTag
(0019,"SIEMENS RA PLANE A",c9)	SL	RectangleAngulation	1	PrivateTag
(0019,"SIEMENS RA PLANE A",ca)	SL	IrisCollimatorFlag	1	PrivateTag
(0019,"SIEMENS RA PLANE A",cb)	SL	IrisRowStart	1	PrivateTag
(0019,"SIEMENS RA PLANE A",cc)	SL	IrisRowEnd	1	PrivateTag
(0019,"SIEMENS RA PLANE A",cd)	SL	IrisColumnStart	1	PrivateTag
(0019,"SIEMENS RA PLANE A",ce)	SL	IrisColumnEnd	1	PrivateTag
(0019,"SIEMENS RA PLANE A",cf)	SL	IrisAngulation	1	PrivateTag
(0019,"SIEMENS RA PLANE A",d1)	SS	NumberOfFramesPlane	1	PrivateTag
(0019,"SIEMENS RA PLANE A",d2)	SS	Internal	1	PrivateTag
(0019,"SIEMENS RA PLANE A",d3)	SS	Internal	1	PrivateTag
(0019,"SIEMENS RA PLANE A",d4)	SS	Internal	1	PrivateTag
(0019,"SIEMENS RA PLANE A",d5)	SS	Internal	1	PrivateTag
(0019,"SIEMENS RA PLANE A",d6)	SS	Internal	1-n	PrivateTag
(0019,"SIEMENS RA PLANE A",d7)	SS	Internal	1-n	PrivateTag
(0019,"SIEMENS RA PLANE A",d8)	SS	Internal	1	PrivateTag
(0019,"SIEMENS RA PLANE A",d9)	SS	Internal	1	PrivateTag
(0019,"SIEMENS RA PLANE A",da)	SS	Internal	1	PrivateTag
(0019,"SIEMENS RA PLANE A",db)	SS	Internal	1	PrivateTag
(0019,"SIEMENS RA PLANE A",dc)	SS	Internal	1	PrivateTag
(0019,"SIEMENS RA PLANE A",dd)	SL	AnatomicBackground	1	PrivateTag
(0019,"SIEMENS RA PLANE A",de)	SL	AutoWindowBase	1-n	PrivateTag
(0019,"SIEMENS RA PLANE A",df)	SS	Internal	1	PrivateTag
(0019,"SIEMENS RA PLANE A",e0)	SL	Internal	1	PrivateTag

(0011,"SIEMENS RA PLANE B",28)	SL	FluoroTimerB	1	PrivateTag
(0011,"SIEMENS RA PLANE B",29)	SL	FluoroSkinDoseB	1	PrivateTag
(0011,"SIEMENS RA PLANE B",2a)	SL	TotalSkinDoseB	1	PrivateTag
(0011,"SIEMENS RA PLANE B",2b)	SL	FluoroDoseAreaProductB	1	PrivateTag
(0011,"SIEMENS RA PLANE B",2c)	SL	TotalDoseAreaProductB	1	PrivateTag
(0019,"SIEMENS RA PLANE B",18)	SS	Internal	1	PrivateTag
(0019,"SIEMENS RA PLANE B",19)	SS	Internal	1	PrivateTag
(0019,"SIEMENS RA PLANE B",1a)	SS	Internal	1	PrivateTag
(0019,"SIEMENS RA PLANE B",1b)	SS	Internal	1	PrivateTag
(0019,"SIEMENS RA PLANE B",1c)	SS	Internal	1	PrivateTag
(0019,"SIEMENS RA PLANE B",1d)	SS	Internal	1	PrivateTag
(0019,"SIEMENS RA PLANE B",1e)	SS	Internal	1	PrivateTag
(0019,"SIEMENS RA PLANE B",1f)	SS	Internal	1	PrivateTag
(0019,"SIEMENS RA PLANE B",20)	SL	SystemCalibFactorPlaneB	1-n	PrivateTag
(0019,"SIEMENS RA PLANE B",22)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",24)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",26)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",28)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",2a)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",2c)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",2e)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",30)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",32)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",34)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",36)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",38)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",3a)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",3c)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",3e)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",40)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",42)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",44)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",46)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",48)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",4a)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS RA PLANE B",4c)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS RA PLANE B",4e)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS RA PLANE B",50)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",52)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",54)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",56)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",58)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",5a)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",5c)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS RA PLANE B",5e)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS RA PLANE B",60)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS RA PLANE B",62)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS RA PLANE B",64)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS RA PLANE B",66)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS RA PLANE B",68)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS RA PLANE B",6a)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS RA PLANE B",6c)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS RA PLANE B",6e)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS RA PLANE B",70)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS RA PLANE B",72)	UN	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",74)	UN	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",76)	UN	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",78)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS RA PLANE B",7a)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",7c)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",7e)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",80)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",82)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",84)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",86)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",88)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",8a)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",8c)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",8e)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",90)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",92)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",94)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",96)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",98)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",9a)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS RA PLANE B",9c)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS RA PLANE B",9e)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS RA PLANE B",a0)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS RA PLANE B",a2)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS RA PLANE B",a4)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",a6)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",a8)	US	Unknown	1-n	PrivateTag
(0019,"SIEMENS RA PLANE B",aa)	US	Unknown	1	PrivateTag
(0019,"SIEMENS RA PLANE B",ac)	US	Unknown	1	PrivateTag

(0011,"SIEMENS RIS",10)	LT	PatientUID	1	PrivateTag
(0011,"SIEMENS RIS",11)	LT	PatientID	1	PrivateTag
(0011,"SIEMENS RIS",20)	DA	PatientRegistrationDate	1	PrivateTag
(0011,"SIEMENS RIS",21)	TM	PatientRegistrationTime	1	PrivateTag
(0011,"SIEMENS RIS",30)	LT	PatientnameRIS	1	PrivateTag
(0011,"SIEMENS RIS",31)	LT	PatientprenameRIS	1	PrivateTag
(0011,"SIEMENS RIS",40)	LT	PatientHospitalStatus	1	PrivateTag
(0011,"SIEMENS RIS",41)	LT	MedicalAlerts	1	PrivateTag
(0011,"SIEMENS RIS",42)	LT	ContrastAllergies	1	PrivateTag
(0031,"SIEMENS RIS",10)	LT	RequestUID	1	PrivateTag
(0031,"SIEMENS RIS",45)	LT	RequestingPhysician	1	PrivateTag
(0031,"SIEMENS RIS",50)	LT	RequestedPhysician	1	PrivateTag
(0033,"SIEMENS RIS",10)	LT	PatientStudyUID	1	PrivateTag

(0021,"SIEMENS SMS-AX  ACQ 1.0",00)	US	AcquisitionType	1	PrivateTag
(0021,"SIEMENS SMS-AX  ACQ 1.0",01)	US	AcquisitionMode	1	PrivateTag
(0021,"SIEMENS SMS-AX  ACQ 1.0",02)	US	FootswitchIndex	1	PrivateTag
(0021,"SIEMENS SMS-AX  ACQ 1.0",03)	US	AcquisitionRoom	1	PrivateTag
(0021,"SIEMENS SMS-AX  ACQ 1.0",04)	SL	CurrentTimeProduct	1	PrivateTag
(0021,"SIEMENS SMS-AX  ACQ 1.0",05)	SL	Dose	1	PrivateTag
(0021,"SIEMENS SMS-AX  ACQ 1.0",06)	SL	SkinDosePercent	1	PrivateTag
(0021,"SIEMENS SMS-AX  ACQ 1.0",07)	SL	SkinDoseAccumulation	1	PrivateTag
(0021,"SIEMENS SMS-AX  ACQ 1.0",08)	SL	SkinDoseRate	1	PrivateTag
(0021,"SIEMENS SMS-AX  ACQ 1.0",0A)	UL	CopperFilter	1	PrivateTag
(0021,"SIEMENS SMS-AX  ACQ 1.0",0B)	US	MeasuringField	1	PrivateTag
(0021,"SIEMENS SMS-AX  ACQ 1.0",0C)	SS	PostBlankingCircle	3	PrivateTag
(0021,"SIEMENS SMS-AX  ACQ 1.0",0D)	SS	DynaAngles	2-2n	PrivateTag
(0021,"SIEMENS SMS-AX  ACQ 1.0",0E)	SS	TotalSteps	1	PrivateTag
(0021,"SIEMENS SMS-AX  ACQ 1.0",0F)	SL	DynaXRayInfo	3-3n	PrivateTag
(0021,"SIEMENS SMS-AX  ACQ 1.0",10)	US	ModalityLUTInputGamma	1	PrivateTag
(0021,"SIEMENS SMS-AX  ACQ 1.0",11)	US	ModalityLUTOutputGamma	1	PrivateTag
(0021,"SIEMENS SMS-AX  ACQ 1.0",12)	OB	SH_STPAR	1-n	PrivateTag
(0021,"SIEMENS SMS-AX  ACQ 1.0",13)	US	AcquisitionZoom	1	PrivateTag
(0021,"SIEMENS SMS-AX  ACQ 1.0",14)	SS	DynaAngulationStepWidth	1	PrivateTag
(0021,"SIEMENS SMS-AX  ACQ 1.0",15)	US	Harmonization	1	PrivateTag
(0021,"SIEMENS SMS-AX  ACQ 1.0",16)	US	DRSingleFlag	1	PrivateTag
(0021,"SIEMENS SMS-AX  ACQ 1.0",17)	SL	SourceToIsocenter	1	PrivateTag
(0021,"SIEMENS SMS-AX  ACQ 1.0",18)	US	PressureData	1	PrivateTag
(0021,"SIEMENS SMS-AX  ACQ 1.0",19)	SL	ECGIndexArray	1	PrivateTag
(0021,"SIEMENS SMS-AX  ACQ 1.0",1A)	US	FDFlag	1	PrivateTag
(0021,"SIEMENS SMS-AX  ACQ 1.0",1B)	OB	SH_ZOOM	1	PrivateTag
(0021,"SIEMENS SMS-AX  ACQ 1.0",1C)	OB	SH_COLPAR	1	PrivateTag
(0021,"SIEMENS SMS-AX  ACQ 1.0",1D)	US	K_Factor	1	PrivateTag
(0021,"SIEMENS SMS-AX  ACQ 1.0",1E)	US	EVE	8	PrivateTag
(0021,"SIEMENS SMS-AX  ACQ 1.0",1F)	SL	TotalSceneTime	1	PrivateTag
(0021,"SIEMENS SMS-AX  ACQ 1.0",20)	US	RestoreFlag	1	PrivateTag
(0021,"SIEMENS SMS-AX  ACQ 1.0",21)	US	StandMovementFlag	1	PrivateTag
(0021,"SIEMENS SMS-AX  ACQ 1.0",22)	US	FDRows	1	PrivateTag
(0021,"SIEMENS SMS-AX  ACQ 1.0",23)	US	FDColumns	1	PrivateTag
(0021,"SIEMENS SMS-AX  ACQ 1.0",24)	US	TableMovementFlag	1	PrivateTag
(0021,"SIEMENS SMS-AX  ACQ 1.0",25)	LO	OriginalOrganProgramName	1	PrivateTag
(0021,"SIEMENS SMS-AX  ACQ 1.0",26)	DS	CrispyXPIFilter	1	PrivateTag

(0025,"SIEMENS SMS-AX  ORIGINAL IMAGE INFO 1.0",00)	US	ViewNative	1	PrivateTag
(0025,"SIEMENS SMS-AX  ORIGINAL IMAGE INFO 1.0",01)	US	OriginalSeriesNumber	1	PrivateTag
(0025,"SIEMENS SMS-AX  ORIGINAL IMAGE INFO 1.0",02)	US	OriginalImageNumber	1	PrivateTag
(0025,"SIEMENS SMS-AX  ORIGINAL IMAGE INFO 1.0",03)	US	WinCenter	1	PrivateTag
(0025,"SIEMENS SMS-AX  ORIGINAL IMAGE INFO 1.0",04)	US	WinWidth	1	PrivateTag
(0025,"SIEMENS SMS-AX  ORIGINAL IMAGE INFO 1.0",05)	US	WinBrightness	1	PrivateTag
(0025,"SIEMENS SMS-AX  ORIGINAL IMAGE INFO 1.0",06)	US	WinContrast	1	PrivateTag
(0025,"SIEMENS SMS-AX  ORIGINAL IMAGE INFO 1.0",07)	US	OriginalFrameNumber	1	PrivateTag
(0025,"SIEMENS SMS-AX  ORIGINAL IMAGE INFO 1.0",08)	US	OriginalMaskFrameNumber	1	PrivateTag
(0025,"SIEMENS SMS-AX  ORIGINAL IMAGE INFO 1.0",09)	US	Opac	1	PrivateTag
(0025,"SIEMENS SMS-AX  ORIGINAL IMAGE INFO 1.0",0A)	US	OriginalNumberOfFrames	1	PrivateTag
(0025,"SIEMENS SMS-AX  ORIGINAL IMAGE INFO 1.0",0B)	DS	OriginalSceneDuration	1	PrivateTag
(0025,"SIEMENS SMS-AX  ORIGINAL IMAGE INFO 1.0",0C)	LO	IdentifierLOID	1	PrivateTag
(0025,"SIEMENS SMS-AX  ORIGINAL IMAGE INFO 1.0",0D)	SS	OriginalSceneVFRInfo	1-n	PrivateTag
(0025,"SIEMENS SMS-AX  ORIGINAL IMAGE INFO 1.0",0E)	SS	OriginalFrameECGPosition	1	PrivateTag
(0025,"SIEMENS SMS-AX  ORIGINAL IMAGE INFO 1.0",0F)	SS	OriginalECG1stFrameOffset_retired	1	PrivateTag
(0025,"SIEMENS SMS-AX  ORIGINAL IMAGE INFO 1.0",10)	SS	ZoomFlag	1	PrivateTag
(0025,"SIEMENS SMS-AX  ORIGINAL IMAGE INFO 1.0",11)	US	Flex	1	PrivateTag
(0025,"SIEMENS SMS-AX  ORIGINAL IMAGE INFO 1.0",12)	US	NumberOfMaskFrames	1	PrivateTag
(0025,"SIEMENS SMS-AX  ORIGINAL IMAGE INFO 1.0",13)	US	NumberOfFillFrames	1	PrivateTag
(0025,"SIEMENS SMS-AX  ORIGINAL IMAGE INFO 1.0",14)	US	SeriesNumber	1	PrivateTag
(0025,"SIEMENS SMS-AX  ORIGINAL IMAGE INFO 1.0",15)	IS	ImageNumber	1	PrivateTag

(0023,"SIEMENS SMS-AX  QUANT 1.0",00)	DS	HorizontalCalibrationPixelSize	2	PrivateTag
(0023,"SIEMENS SMS-AX  QUANT 1.0",01)	DS	VerticalCalibrationPixelSize	2	PrivateTag
(0023,"SIEMENS SMS-AX  QUANT 1.0",02)	LO	CalibrationObject	1	PrivateTag
(0023,"SIEMENS SMS-AX  QUANT 1.0",03)	DS	CalibrationObjectSize	1	PrivateTag
(0023,"SIEMENS SMS-AX  QUANT 1.0",04)	LO	CalibrationMethod	1	PrivateTag
(0023,"SIEMENS SMS-AX  QUANT 1.0",05)	ST	Filename	1	PrivateTag
(0023,"SIEMENS SMS-AX  QUANT 1.0",06)	IS	FrameNumber	1	PrivateTag
(0023,"SIEMENS SMS-AX  QUANT 1.0",07)	IS	CalibrationFactorMultiplicity	2	PrivateTag
(0023,"SIEMENS SMS-AX  QUANT 1.0",08)	IS	CalibrationTODValue	1	PrivateTag

(0019,"SIEMENS SMS-AX  VIEW 1.0",00)	US	ReviewMode	1	PrivateTag
(0019,"SIEMENS SMS-AX  VIEW 1.0",01)	US	AnatomicalBackgroundPercent	1	PrivateTag
(0019,"SIEMENS SMS-AX  VIEW 1.0",02)	US	NumberOfPhases	1	PrivateTag
(0019,"SIEMENS SMS-AX  VIEW 1.0",03)	US	ApplyAnatomicalBackground	1	PrivateTag
(0019,"SIEMENS SMS-AX  VIEW 1.0",04)	SS	PixelShiftArray	4-4n	PrivateTag
(0019,"SIEMENS SMS-AX  VIEW 1.0",05)	US	Brightness	1	PrivateTag
(0019,"SIEMENS SMS-AX  VIEW 1.0",06)	US	Contrast	1	PrivateTag
(0019,"SIEMENS SMS-AX  VIEW 1.0",07)	US	Enabled	1	PrivateTag
(0019,"SIEMENS SMS-AX  VIEW 1.0",08)	US	NativeEdgeEnhancementPercentGain	1	PrivateTag
(0019,"SIEMENS SMS-AX  VIEW 1.0",09)	SS	NativeEdgeEnhancementLUTIndex	1	PrivateTag
(0019,"SIEMENS SMS-AX  VIEW 1.0",0A)	SS	NativeEdgeEnhancementKernelSize	1	PrivateTag
(0019,"SIEMENS SMS-AX  VIEW 1.0",0B)	US	SubtrEdgeEnhancementPercentGain	1	PrivateTag
(0019,"SIEMENS SMS-AX  VIEW 1.0",0C)	SS	SubtrEdgeEnhancementLUTIndex	1	PrivateTag
(0019,"SIEMENS SMS-AX  VIEW 1.0",0D)	SS	SubtrEdgeEnhancementKernelSize	1	PrivateTag
(0019,"SIEMENS SMS-AX  VIEW 1.0",0E)	US	FadePercent	1	PrivateTag
(0019,"SIEMENS SMS-AX  VIEW 1.0",0F)	US	FlippedBeforeLateralityApplied	1	PrivateTag
(0019,"SIEMENS SMS-AX  VIEW 1.0",10)	US	ApplyFade	1	PrivateTag
(0019,"SIEMENS SMS-AX  VIEW 1.0",12)	US	Zoom	1	PrivateTag
(0019,"SIEMENS SMS-AX  VIEW 1.0",13)	SS	PanX	1	PrivateTag
(0019,"SIEMENS SMS-AX  VIEW 1.0",14)	SS	PanY	1	PrivateTag
(0019,"SIEMENS SMS-AX  VIEW 1.0",15)	SS	NativeEdgeEnhancementAdvPercGain	1	PrivateTag
(0019,"SIEMENS SMS-AX  VIEW 1.0",16)	SS	SubtrEdgeEnhancementAdvPercGain	1	PrivateTag
(0019,"SIEMENS SMS-AX  VIEW 1.0",17)	US	InvertFlag	1	PrivateTag
(0019,"SIEMENS SMS-AX  VIEW 1.0",1A)	OB	Quant1KOverlay	1	PrivateTag
(0019,"SIEMENS SMS-AX  VIEW 1.0",1B)	US	OriginalResolution	1	PrivateTag
(0019,"SIEMENS SMS-AX  VIEW 1.0",1C)	DS	AutoWindowCenter	1	PrivateTag
(0019,"SIEMENS SMS-AX  VIEW 1.0",1D)	DS	AutoWindowWidth	1	PrivateTag

(0009,"SIENET",01)	US	SIENETCommandField	1	PrivateTag
(0009,"SIENET",14)	LT	ReceiverPLA	1	PrivateTag
(0009,"SIENET",16)	US	TransferPriority	1	PrivateTag
(0009,"SIENET",29)	LT	ActualUser	1	PrivateTag
(0095,"SIENET",01)	LT	ExaminationFolderID	1	PrivateTag
(0095,"SIENET",04)	UL	FolderReportedStatus	1	PrivateTag
(0095,"SIENET",05)	LT	FolderReportingRadiologist	1	PrivateTag
(0095,"SIENET",07)	LT	SIENETISAPLA	1	PrivateTag
(0099,"SIENET",02)	UL	DataObjectAttributes	1	PrivateTag

(0009,"SPI RELEASE 1",10)	LT	Comments	1	PrivateTag
(0009,"SPI RELEASE 1",15)	LO	SPIImageUID	1	PrivateTag
(0009,"SPI RELEASE 1",40)	US	DataObjectType	1	PrivateTag
(0009,"SPI RELEASE 1",41)	LO	DataObjectSubtype	1	PrivateTag
(0011,"SPI RELEASE 1",10)	LO	Organ	1	PrivateTag
(0011,"SPI RELEASE 1",15)	LO	AllergyIndication	1	PrivateTag
(0011,"SPI RELEASE 1",20)	LO	Pregnancy	1	PrivateTag
(0029,"SPI RELEASE 1",60)	LT	CompressionAlgorithm	1	PrivateTag

(0009,"SPI Release 1",10)	LT	Comments	1	PrivateTag
(0009,"SPI Release 1",15)	LO	SPIImageUID	1	PrivateTag
(0009,"SPI Release 1",40)	US	DataObjectType	1	PrivateTag
(0009,"SPI Release 1",41)	LO	DataObjectSubtype	1	PrivateTag
(0011,"SPI Release 1",10)	LO	Organ	1	PrivateTag
(0011,"SPI Release 1",15)	LO	AllergyIndication	1	PrivateTag
(0011,"SPI Release 1",20)	LO	Pregnancy	1	PrivateTag
(0029,"SPI Release 1",60)	LT	CompressionAlgorithm	1	PrivateTag

(0009,"SPI",10)	LO	Comments	1	PrivateTag
(0009,"SPI",15)	LO	SPIImageUID	1	PrivateTag
(0009,"SPI",40)	US	DataObjectType	1	PrivateTag
(0009,"SPI",41)	LT	DataObjectSubtype	1	PrivateTag
(0011,"SPI",10)	LT	Organ	1	PrivateTag
(0011,"SPI",15)	LT	AllergyIndication	1	PrivateTag
(0011,"SPI",20)	LT	Pregnancy	1	PrivateTag
(0029,"SPI",60)	LT	CompressionAlgorithm	1	PrivateTag

(0011,"SPI RELEASE 1",10)	LO	Organ	1	PrivateTag
(0011,"SPI RELEASE 1",15)	LO	AllergyIndication	1	PrivateTag
(0011,"SPI RELEASE 1",20)	LO	Pregnancy	1	PrivateTag

(0009,"SPI-P Release 1",00)	LT	DataObjectRecognitionCode	1	PrivateTag
(0009,"SPI-P Release 1",04)	LO	ImageDataConsistence	1	PrivateTag
(0009,"SPI-P Release 1",08)	US	Unknown	1	PrivateTag
(0009,"SPI-P Release 1",12)	LO	Unknown	1	PrivateTag
(0009,"SPI-P Release 1",15)	LO	UniqueIdentifier	1	PrivateTag
(0009,"SPI-P Release 1",16)	LO	Unknown	1	PrivateTag
(0009,"SPI-P Release 1",18)	LO	Unknown	1	PrivateTag
(0009,"SPI-P Release 1",21)	LT	Unknown	1	PrivateTag
(0009,"SPI-P Release 1",31)	LT	PACSUniqueIdentifier	1	PrivateTag
(0009,"SPI-P Release 1",34)	LT	ClusterUniqueIdentifier	1	PrivateTag
(0009,"SPI-P Release 1",38)	LT	SystemUniqueIdentifier	1	PrivateTag
(0009,"SPI-P Release 1",39)	LT	Unknown	1	PrivateTag
(0009,"SPI-P Release 1",51)	LT	StudyUniqueIdentifier	1	PrivateTag
(0009,"SPI-P Release 1",61)	LT	SeriesUniqueIdentifier	1	PrivateTag
(0009,"SPI-P Release 1",91)	LT	Unknown	1	PrivateTag
(0009,"SPI-P Release 1",f2)	LT	Unknown	1	PrivateTag
(0009,"SPI-P Release 1",f3)	UN	Unknown	1	PrivateTag
(0009,"SPI-P Release 1",f4)	LT	Unknown	1	PrivateTag
(0009,"SPI-P Release 1",f5)	UN	Unknown	1	PrivateTag
(0009,"SPI-P Release 1",f7)	LT	Unknown	1	PrivateTag
(0011,"SPI-P Release 1",10)	LT	PatientEntryID	1	PrivateTag
(0011,"SPI-P Release 1",21)	UN	Unknown	1	PrivateTag
(0011,"SPI-P Release 1",22)	UN	Unknown	1	PrivateTag
(0011,"SPI-P Release 1",31)	UN	Unknown	1	PrivateTag
(0011,"SPI-P Release 1",32)	UN	Unknown	1	PrivateTag
(0019,"SPI-P Release 1",00)	UN	Unknown	1	PrivateTag
(0019,"SPI-P Release 1",01)	UN	Unknown	1	PrivateTag
(0019,"SPI-P Release 1",02)	UN	Unknown	1	PrivateTag
(0019,"SPI-P Release 1",10)	US	MainsFrequency	1	PrivateTag
(0019,"SPI-P Release 1",25)	LT	OriginalPixelDataQuality	1-n	PrivateTag
(0019,"SPI-P Release 1",30)	US	ECGTriggering	1	PrivateTag
(0019,"SPI-P Release 1",31)	UN	ECG1Offset	1	PrivateTag
(0019,"SPI-P Release 1",32)	UN	ECG2Offset1	1	PrivateTag
(0019,"SPI-P Release 1",33)	UN	ECG2Offset2	1	PrivateTag
(0019,"SPI-P Release 1",50)	US	VideoScanMode	1	PrivateTag
(0019,"SPI-P Release 1",51)	US	VideoLineRate	1	PrivateTag
(0019,"SPI-P Release 1",60)	US	XrayTechnique	1	PrivateTag
(0019,"SPI-P Release 1",61)	DS	ImageIdentifierFromat	1	PrivateTag
(0019,"SPI-P Release 1",62)	US	IrisDiaphragm	1	PrivateTag
(0019,"SPI-P Release 1",63)	CS	Filter	1	PrivateTag
(0019,"SPI-P Release 1",64)	CS	CineParallel	1	PrivateTag
(0019,"SPI-P Release 1",65)	CS	CineMaster	1	PrivateTag
(0019,"SPI-P Release 1",70)	US	ExposureChannel	1	PrivateTag
(0019,"SPI-P Release 1",71)	UN	ExposureChannelFirstImage	1	PrivateTag
(0019,"SPI-P Release 1",72)	US	ProcessingChannel	1	PrivateTag
(0019,"SPI-P Release 1",80)	DS	AcquisitionDelay	1	PrivateTag
(0019,"SPI-P Release 1",81)	UN	RelativeImageTime	1	PrivateTag
(0019,"SPI-P Release 1",90)	CS	VideoWhiteCompression	1	PrivateTag
(0019,"SPI-P Release 1",a0)	US	Angulation	1	PrivateTag
(0019,"SPI-P Release 1",a1)	US	Rotation	1	PrivateTag
(0021,"SPI-P Release 1",12)	LT	SeriesUniqueIdentifier	1	PrivateTag
(0021,"SPI-P Release 1",14)	LT	Unknown	1	PrivateTag
(0029,"SPI-P Release 1",00)	DS	Unknown	4	PrivateTag
(0029,"SPI-P Release 1",20)	DS	PixelAspectRatio	1	PrivateTag
(0029,"SPI-P Release 1",25)	LO	ProcessedPixelDataQuality	1-n	PrivateTag
(0029,"SPI-P Release 1",30)	LT	Unknown	1	PrivateTag
(0029,"SPI-P Release 1",38)	US	Unknown	1	PrivateTag
(0029,"SPI-P Release 1",60)	LT	Unknown	1	PrivateTag
(0029,"SPI-P Release 1",61)	LT	Unknown	1	PrivateTag
(0029,"SPI-P Release 1",67)	LT	Unknown	1	PrivateTag
(0029,"SPI-P Release 1",70)	LT	WindowID	1	PrivateTag
(0029,"SPI-P Release 1",71)	CS	VideoInvertSubtracted	1	PrivateTag
(0029,"SPI-P Release 1",72)	CS	VideoInvertNonsubtracted	1	PrivateTag
(0029,"SPI-P Release 1",77)	CS	WindowSelectStatus	1	PrivateTag
(0029,"SPI-P Release 1",78)	LT	ECGDisplayPrintingID	1	PrivateTag
(0029,"SPI-P Release 1",79)	CS	ECGDisplayPrinting	1	PrivateTag
(0029,"SPI-P Release 1",7e)	CS	ECGDisplayPrintingEnableStatus	1	PrivateTag
(0029,"SPI-P Release 1",7f)	CS	ECGDisplayPrintingSelectStatus	1	PrivateTag
(0029,"SPI-P Release 1",80)	LT	PhysiologicalDisplayID	1	PrivateTag
(0029,"SPI-P Release 1",81)	US	PreferredPhysiologicalChannelDisplay	1	PrivateTag
(0029,"SPI-P Release 1",8e)	CS	PhysiologicalDisplayEnableStatus	1	PrivateTag
(0029,"SPI-P Release 1",8f)	CS	PhysiologicalDisplaySelectStatus	1	PrivateTag
(0029,"SPI-P Release 1",c0)	LT	FunctionalShutterID	1	PrivateTag
(0029,"SPI-P Release 1",c1)	US	FieldOfShutter	1	PrivateTag
(0029,"SPI-P Release 1",c5)	LT	FieldOfShutterRectangle	1	PrivateTag
(0029,"SPI-P Release 1",ce)	CS	ShutterEnableStatus	1	PrivateTag
(0029,"SPI-P Release 1",cf)	CS	ShutterSelectStatus	1	PrivateTag
(7FE1,"SPI-P Release 1",10)	ox	PixelData	1	PrivateTag

(0009,"SPI-P Release 1;1",c0)	LT	Unknown	1	PrivateTag
(0009,"SPI-P Release 1;1",c1)	LT	Unknown	1	PrivateTag
(0019,"SPI-P Release 1;1",00)	UN	PhysiologicalDataType	1	PrivateTag
(0019,"SPI-P Release 1;1",01)	UN	PhysiologicalDataChannelAndKind	1	PrivateTag
(0019,"SPI-P Release 1;1",02)	US	SampleBitsAllocated	1	PrivateTag
(0019,"SPI-P Release 1;1",03)	US	SampleBitsStored	1	PrivateTag
(0019,"SPI-P Release 1;1",04)	US	SampleHighBit	1	PrivateTag
(0019,"SPI-P Release 1;1",05)	US	SampleRepresentation	1	PrivateTag
(0019,"SPI-P Release 1;1",06)	UN	SmallestSampleValue	1	PrivateTag
(0019,"SPI-P Release 1;1",07)	UN	LargestSampleValue	1	PrivateTag
(0019,"SPI-P Release 1;1",08)	UN	NumberOfSamples	1	PrivateTag
(0019,"SPI-P Release 1;1",09)	UN	SampleData	1	PrivateTag
(0019,"SPI-P Release 1;1",0a)	UN	SampleRate	1	PrivateTag
(0019,"SPI-P Release 1;1",10)	UN	PhysiologicalDataType2	1	PrivateTag
(0019,"SPI-P Release 1;1",11)	UN	PhysiologicalDataChannelAndKind2	1	PrivateTag
(0019,"SPI-P Release 1;1",12)	US	SampleBitsAllocated2	1	PrivateTag
(0019,"SPI-P Release 1;1",13)	US	SampleBitsStored2	1	PrivateTag
(0019,"SPI-P Release 1;1",14)	US	SampleHighBit2	1	PrivateTag
(0019,"SPI-P Release 1;1",15)	US	SampleRepresentation2	1	PrivateTag
(0019,"SPI-P Release 1;1",16)	UN	SmallestSampleValue2	1	PrivateTag
(0019,"SPI-P Release 1;1",17)	UN	LargestSampleValue2	1	PrivateTag
(0019,"SPI-P Release 1;1",18)	UN	NumberOfSamples2	1	PrivateTag
(0019,"SPI-P Release 1;1",19)	UN	SampleData2	1	PrivateTag
(0019,"SPI-P Release 1;1",1a)	UN	SampleRate2	1	PrivateTag
(0029,"SPI-P Release 1;1",00)	LT	ZoomID	1	PrivateTag
(0029,"SPI-P Release 1;1",01)	DS	ZoomRectangle	1-n	PrivateTag
(0029,"SPI-P Release 1;1",03)	DS	ZoomFactor	1	PrivateTag
(0029,"SPI-P Release 1;1",04)	US	ZoomFunction	1	PrivateTag
(0029,"SPI-P Release 1;1",0e)	CS	ZoomEnableStatus	1	PrivateTag
(0029,"SPI-P Release 1;1",0f)	CS	ZoomSelectStatus	1	PrivateTag
(0029,"SPI-P Release 1;1",40)	LT	MagnifyingGlassID	1	PrivateTag
(0029,"SPI-P Release 1;1",41)	DS	MagnifyingGlassRectangle	1-n	PrivateTag
(0029,"SPI-P Release 1;1",43)	DS	MagnifyingGlassFactor	1	PrivateTag
(0029,"SPI-P Release 1;1",44)	US	MagnifyingGlassFunction	1	PrivateTag
(0029,"SPI-P Release 1;1",4e)	CS	MagnifyingGlassEnableStatus	1	PrivateTag
(0029,"SPI-P Release 1;1",4f)	CS	MagnifyingGlassSelectStatus	1	PrivateTag

(0029,"SPI-P Release 1;2",00)	LT	SubtractionMaskID	1	PrivateTag
(0029,"SPI-P Release 1;2",04)	UN	MaskingFunction	1	PrivateTag
(0029,"SPI-P Release 1;2",0c)	UN	ProprietaryMaskingParameters	1	PrivateTag
(0029,"SPI-P Release 1;2",1e)	CS	SubtractionMaskEnableStatus	1	PrivateTag
(0029,"SPI-P Release 1;2",1f)	CS	SubtractionMaskSelectStatus	1	PrivateTag
(0029,"SPI-P Release 1;3",00)	LT	ImageEnhancementID	1	PrivateTag
(0029,"SPI-P Release 1;3",01)	LT	ImageEnhancement	1	PrivateTag
(0029,"SPI-P Release 1;3",02)	LT	ConvolutionID	1	PrivateTag
(0029,"SPI-P Release 1;3",03)	LT	ConvolutionType	1	PrivateTag
(0029,"SPI-P Release 1;3",04)	LT	ConvolutionKernelSizeID	1	PrivateTag
(0029,"SPI-P Release 1;3",05)	US	ConvolutionKernelSize	2	PrivateTag
(0029,"SPI-P Release 1;3",06)	US	ConvolutionKernel	1-n	PrivateTag
(0029,"SPI-P Release 1;3",0c)	DS	EnhancementGain	1	PrivateTag
(0029,"SPI-P Release 1;3",1e)	CS	ImageEnhancementEnableStatus	1	PrivateTag
(0029,"SPI-P Release 1;3",1f)	CS	ImageEnhancementSelectStatus	1	PrivateTag

(0011,"SPI-P Release 2;1",18)	LT	Unknown	1	PrivateTag
(0023,"SPI-P Release 2;1",0d)	UI	Unknown	1	PrivateTag
(0023,"SPI-P Release 2;1",0e)	UI	Unknown	1	PrivateTag

(0009,"SPI-P-GV-CT Release 1",00)	LO	Unknown	1	PrivateTag
(0009,"SPI-P-GV-CT Release 1",10)	LO	Unknown	1	PrivateTag
(0009,"SPI-P-GV-CT Release 1",20)	LO	Unknown	1	PrivateTag
(0009,"SPI-P-GV-CT Release 1",30)	LO	Unknown	1	PrivateTag
(0009,"SPI-P-GV-CT Release 1",40)	LO	Unknown	1	PrivateTag
(0009,"SPI-P-GV-CT Release 1",50)	LO	Unknown	1	PrivateTag
(0009,"SPI-P-GV-CT Release 1",60)	LO	Unknown	1	PrivateTag
(0009,"SPI-P-GV-CT Release 1",70)	LO	Unknown	1	PrivateTag
(0009,"SPI-P-GV-CT Release 1",75)	LO	Unknown	1	PrivateTag
(0009,"SPI-P-GV-CT Release 1",80)	LO	Unknown	1	PrivateTag
(0009,"SPI-P-GV-CT Release 1",90)	LO	Unknown	1	PrivateTag
(0019,"SPI-P-GV-CT Release 1",08)	IS	Unknown	1	PrivateTag
(0019,"SPI-P-GV-CT Release 1",09)	IS	Unknown	1	PrivateTag
(0019,"SPI-P-GV-CT Release 1",0a)	IS	Unknown	1	PrivateTag
(0019,"SPI-P-GV-CT Release 1",10)	LO	Unknown	1	PrivateTag
(0019,"SPI-P-GV-CT Release 1",20)	TM	Unknown	1	PrivateTag
(0019,"SPI-P-GV-CT Release 1",50)	LO	Unknown	1	PrivateTag
(0019,"SPI-P-GV-CT Release 1",60)	DS	Unknown	1	PrivateTag
(0019,"SPI-P-GV-CT Release 1",61)	US	Unknown	1	PrivateTag
(0019,"SPI-P-GV-CT Release 1",63)	LO	Unknown	1	PrivateTag
(0019,"SPI-P-GV-CT Release 1",64)	US	Unknown	1	PrivateTag
(0019,"SPI-P-GV-CT Release 1",65)	IS	Unknown	1	PrivateTag
(0019,"SPI-P-GV-CT Release 1",70)	LT	Unknown	1	PrivateTag
(0019,"SPI-P-GV-CT Release 1",80)	LO	Unknown	1	PrivateTag
(0019,"SPI-P-GV-CT Release 1",81)	LO	Unknown	1	PrivateTag
(0019,"SPI-P-GV-CT Release 1",90)	LO	Unknown	1	PrivateTag
(0019,"SPI-P-GV-CT Release 1",a0)	LO	Unknown	1	PrivateTag
(0019,"SPI-P-GV-CT Release 1",a1)	US	Unknown	1	PrivateTag
(0019,"SPI-P-GV-CT Release 1",a2)	US	Unknown	1	PrivateTag
(0019,"SPI-P-GV-CT Release 1",a3)	US	Unknown	1	PrivateTag
(0019,"SPI-P-GV-CT Release 1",b0)	LO	Unknown	1	PrivateTag
(0019,"SPI-P-GV-CT Release 1",b1)	LO	Unknown	1	PrivateTag
(0021,"SPI-P-GV-CT Release 1",20)	LO	Unknown	1	PrivateTag
(0021,"SPI-P-GV-CT Release 1",30)	DS	Unknown	1	PrivateTag
(0021,"SPI-P-GV-CT Release 1",40)	LO	Unknown	1	PrivateTag
(0021,"SPI-P-GV-CT Release 1",50)	LO	Unknown	1	PrivateTag
(0021,"SPI-P-GV-CT Release 1",60)	DS	Unknown	1	PrivateTag
(0021,"SPI-P-GV-CT Release 1",70)	DS	Unknown	1	PrivateTag
(0021,"SPI-P-GV-CT Release 1",80)	DS	Unknown	1	PrivateTag
(0021,"SPI-P-GV-CT Release 1",90)	DS	Unknown	1	PrivateTag
(0021,"SPI-P-GV-CT Release 1",a0)	US	Unknown	1	PrivateTag
(0021,"SPI-P-GV-CT Release 1",a1)	DS	Unknown	1	PrivateTag
(0021,"SPI-P-GV-CT Release 1",a2)	DS	Unknown	1	PrivateTag
(0021,"SPI-P-GV-CT Release 1",a3)	LT	Unknown	1	PrivateTag
(0021,"SPI-P-GV-CT Release 1",a4)	LT	Unknown	1	PrivateTag
(0021,"SPI-P-GV-CT Release 1",b0)	LO	Unknown	1	PrivateTag
(0021,"SPI-P-GV-CT Release 1",c0)	LO	Unknown	1	PrivateTag
(0029,"SPI-P-GV-CT Release 1",10)	LO	Unknown	1	PrivateTag
(0029,"SPI-P-GV-CT Release 1",30)	UL	Unknown	1	PrivateTag
(0029,"SPI-P-GV-CT Release 1",31)	UL	Unknown	1	PrivateTag
(0029,"SPI-P-GV-CT Release 1",32)	UL	Unknown	1	PrivateTag
(0029,"SPI-P-GV-CT Release 1",33)	UL	Unknown	1	PrivateTag
(0029,"SPI-P-GV-CT Release 1",80)	LO	Unknown	1	PrivateTag
(0029,"SPI-P-GV-CT Release 1",90)	LO	Unknown	1	PrivateTag
(0029,"SPI-P-GV-CT Release 1",d0)	IS	Unknown	1	PrivateTag
(0029,"SPI-P-GV-CT Release 1",d1)	IS	Unknown	1	PrivateTag

(0019,"SPI-P-PCR Release 2",30)	US	Unknown	1	PrivateTag

(0021,"SPI-P-Private-CWS Release 1",00)	LT	WindowOfImagesID	1	PrivateTag
(0021,"SPI-P-Private-CWS Release 1",01)	CS	WindowOfImagesType	1	PrivateTag
(0021,"SPI-P-Private-CWS Release 1",02)	IS	WindowOfImagesScope	1-n	PrivateTag

(0019,"SPI-P-Private-DCI Release 1",10)	UN	ECGTimeMapDataBitsAllocated	1	PrivateTag
(0019,"SPI-P-Private-DCI Release 1",11)	UN	ECGTimeMapDataBitsStored	1	PrivateTag
(0019,"SPI-P-Private-DCI Release 1",12)	UN	ECGTimeMapDataHighBit	1	PrivateTag
(0019,"SPI-P-Private-DCI Release 1",13)	UN	ECGTimeMapDataRepresentation	1	PrivateTag
(0019,"SPI-P-Private-DCI Release 1",14)	UN	ECGTimeMapDataSmallestDataValue	1	PrivateTag
(0019,"SPI-P-Private-DCI Release 1",15)	UN	ECGTimeMapDataLargestDataValue	1	PrivateTag
(0019,"SPI-P-Private-DCI Release 1",16)	UN	ECGTimeMapDataNumberOfDataValues	1	PrivateTag
(0019,"SPI-P-Private-DCI Release 1",17)	UN	ECGTimeMapData	1	PrivateTag

(0021,"SPI-P-Private_CDS Release 1",40)	IS	Unknown	1	PrivateTag
(0029,"SPI-P-Private_CDS Release 1",00)	UN	Unknown	1	PrivateTag

(0019,"SPI-P-Private_ICS Release 1",30)	DS	Unknown	1	PrivateTag
(0019,"SPI-P-Private_ICS Release 1",31)	LO	Unknown	1	PrivateTag
(0029,"SPI-P-Private_ICS Release 1",08)	SQ	Unknown	1	PrivateTag
(0029,"SPI-P-Private_ICS Release 1",0f)	SQ	Unknown	1	PrivateTag
(0029,"SPI-P-Private_ICS Release 1",10)	SQ	Unknown	1	PrivateTag
(0029,"SPI-P-Private_ICS Release 1",1b)	SQ	Unknown	1	PrivateTag
(0029,"SPI-P-Private_ICS Release 1",1c)	SQ	Unknown	1	PrivateTag
(0029,"SPI-P-Private_ICS Release 1",21)	SQ	Unknown	1	PrivateTag
(0029,"SPI-P-Private_ICS Release 1",43)	SQ	Unknown	1	PrivateTag
(0029,"SPI-P-Private_ICS Release 1",44)	SQ	Unknown	1	PrivateTag
(0029,"SPI-P-Private_ICS Release 1",4C)	SQ	Unknown	1	PrivateTag
(0029,"SPI-P-Private_ICS Release 1",67)	LO	Unknown	1	PrivateTag
(0029,"SPI-P-Private_ICS Release 1",68)	US	Unknown	1	PrivateTag
(0029,"SPI-P-Private_ICS Release 1",6A)	LO	Unknown	1	PrivateTag
(0029,"SPI-P-Private_ICS Release 1",6B)	US	Unknown	1	PrivateTag

(0029,"SPI-P-Private_ICS Release 1;1",00)	SL	Unknown	1	PrivateTag
(0029,"SPI-P-Private_ICS Release 1;1",05)	FL	Unknown	1	PrivateTag
(0029,"SPI-P-Private_ICS Release 1;1",06)	FL	Unknown	1	PrivateTag
(0029,"SPI-P-Private_ICS Release 1;1",20)	FL	Unknown	1	PrivateTag
(0029,"SPI-P-Private_ICS Release 1;1",21)	FL	Unknown	1	PrivateTag
(0029,"SPI-P-Private_ICS Release 1;1",CD)	SQ	Unknown	1	PrivateTag

(0029,"SPI-P-Private_ICS Release 1;2",00)	FD	Unknown	1	PrivateTag
(0029,"SPI-P-Private_ICS Release 1;2",01)	FD	Unknown	1	PrivateTag
(0029,"SPI-P-Private_ICS Release 1;2",02)	FD	Unknown	1	PrivateTag
(0029,"SPI-P-Private_ICS Release 1;2",03)	SL	Unknown	1	PrivateTag
(0029,"SPI-P-Private_ICS Release 1;2",04)	SL	Unknown	1	PrivateTag
(0029,"SPI-P-Private_ICS Release 1;2",05)	SL	Unknown	1	PrivateTag

(0029,"SPI-P-Private_ICS Release 1;3",C0)	SQ	Unknown	1	PrivateTag
(0029,"SPI-P-Private_ICS Release 1;3",C1)	SQ	Unknown	1	PrivateTag
(0029,"SPI-P-Private_ICS Release 1;3",C2)	SQ	Unknown	1	PrivateTag
(0029,"SPI-P-Private_ICS Release 1;3",C3)	SQ	Unknown	1	PrivateTag
(0029,"SPI-P-Private_ICS Release 1;3",C4)	SQ	Unknown	1	PrivateTag
(0029,"SPI-P-Private_ICS Release 1;3",C5)	SQ	Unknown	1	PrivateTag

(0029,"SPI-P-Private_ICS Release 1;4",02)	SQ	Unknown	1	PrivateTag
(0029,"SPI-P-Private_ICS Release 1;4",9A)	SQ	Unknown	1	PrivateTag
(0029,"SPI-P-Private_ICS Release 1;4",E0)	SQ	Unknown	1	PrivateTag

(0029,"SPI-P-Private_ICS Release 1;5",50)	CS	Unknown	1	PrivateTag
(0029,"SPI-P-Private_ICS Release 1;5",55)	CS	Unknown	1	PrivateTag

(0019,"SPI-P-XSB-DCI Release 1",10)	LT	VideoBeamBoost	1	PrivateTag
(0019,"SPI-P-XSB-DCI Release 1",11)	US	ChannelGeneratingVideoSync	1	PrivateTag
(0019,"SPI-P-XSB-DCI Release 1",12)	US	VideoGain	1	PrivateTag
(0019,"SPI-P-XSB-DCI Release 1",13)	US	VideoOffset	1	PrivateTag
(0019,"SPI-P-XSB-DCI Release 1",20)	DS	RTDDataCompressionFactor	1	PrivateTag

(0029,"Silhouette Annot V1.0",11)	IS	AnnotationName	1	PrivateTag
(0029,"Silhouette Annot V1.0",12)	LT	AnnotationFont	1	PrivateTag
(0029,"Silhouette Annot V1.0",13)	LT	AnnotationTextForegroundColor	1	PrivateTag
(0029,"Silhouette Annot V1.0",14)	LT	AnnotationTextBackgroundColor	1	PrivateTag
(0029,"Silhouette Annot V1.0",15)	UL	AnnotationTextBackingMode	1	PrivateTag
(0029,"Silhouette Annot V1.0",16)	UL	AnnotationTextJustification	1	PrivateTag
(0029,"Silhouette Annot V1.0",17)	UL	AnnotationTextLocation	1	PrivateTag
(0029,"Silhouette Annot V1.0",18)	LT	AnnotationTextString	1	PrivateTag
(0029,"Silhouette Annot V1.0",19)	UL	AnnotationTextAttachMode	1	PrivateTag
(0029,"Silhouette Annot V1.0",20)	UL	AnnotationTextCursorMode	1	PrivateTag
(0029,"Silhouette Annot V1.0",21)	UL	AnnotationTextShadowOffsetX	1	PrivateTag
(0029,"Silhouette Annot V1.0",22)	UL	AnnotationTextShadowOffsetY	1	PrivateTag
(0029,"Silhouette Annot V1.0",23)	LT	AnnotationLineColor	1	PrivateTag
(0029,"Silhouette Annot V1.0",24)	UL	AnnotationLineThickness	1	PrivateTag
(0029,"Silhouette Annot V1.0",25)	UL	AnnotationLineType	1	PrivateTag
(0029,"Silhouette Annot V1.0",26)	UL	AnnotationLineStyle	1	PrivateTag
(0029,"Silhouette Annot V1.0",27)	UL	AnnotationLineDashLength	1	PrivateTag
(0029,"Silhouette Annot V1.0",28)	UL	AnnotationLineAttachMode	1	PrivateTag
(0029,"Silhouette Annot V1.0",29)	UL	AnnotationLinePointCount	1	PrivateTag
(0029,"Silhouette Annot V1.0",30)	FD	AnnotationLinePoints	1	PrivateTag
(0029,"Silhouette Annot V1.0",31)	UL	AnnotationLineControlSize	1	PrivateTag
(0029,"Silhouette Annot V1.0",32)	LT	AnnotationMarkerColor	1	PrivateTag
(0029,"Silhouette Annot V1.0",33)	UL	AnnotationMarkerType	1	PrivateTag
(0029,"Silhouette Annot V1.0",34)	UL	AnnotationMarkerSize	1	PrivateTag
(0029,"Silhouette Annot V1.0",35)	FD	AnnotationMarkerLocation	1	PrivateTag
(0029,"Silhouette Annot V1.0",36)	UL	AnnotationMarkerAttachMode	1	PrivateTag
(0029,"Silhouette Annot V1.0",37)	LT	AnnotationGeomColor	1	PrivateTag
(0029,"Silhouette Annot V1.0",38)	UL	AnnotationGeomThickness	1	PrivateTag
(0029,"Silhouette Annot V1.0",39)	UL	AnnotationGeomLineStyle	1	PrivateTag
(0029,"Silhouette Annot V1.0",40)	UL	AnnotationGeomDashLength	1	PrivateTag
(0029,"Silhouette Annot V1.0",41)	UL	AnnotationGeomFillPattern	1	PrivateTag
(0029,"Silhouette Annot V1.0",42)	UL	AnnotationInteractivity	1	PrivateTag
(0029,"Silhouette Annot V1.0",43)	FD	AnnotationArrowLength	1	PrivateTag
(0029,"Silhouette Annot V1.0",44)	FD	AnnotationArrowAngle	1	PrivateTag
(0029,"Silhouette Annot V1.0",45)	UL	AnnotationDontSave	1	PrivateTag

(0029,"Silhouette Graphics Export V1.0",00)	UI	Unknown	1	PrivateTag

(0029,"Silhouette Line V1.0",11)	IS	LineName	1	PrivateTag
(0029,"Silhouette Line V1.0",12)	LT	LineNameFont	1	PrivateTag
(0029,"Silhouette Line V1.0",13)	UL	LineNameDisplay	1	PrivateTag
(0029,"Silhouette Line V1.0",14)	LT	LineNormalColor	1	PrivateTag
(0029,"Silhouette Line V1.0",15)	UL	LineType	1	PrivateTag
(0029,"Silhouette Line V1.0",16)	UL	LineThickness	1	PrivateTag
(0029,"Silhouette Line V1.0",17)	UL	LineStyle	1	PrivateTag
(0029,"Silhouette Line V1.0",18)	UL	LineDashLength	1	PrivateTag
(0029,"Silhouette Line V1.0",19)	UL	LineInteractivity	1	PrivateTag
(0029,"Silhouette Line V1.0",20)	LT	LineMeasurementColor	1	PrivateTag
(0029,"Silhouette Line V1.0",21)	LT	LineMeasurementFont	1	PrivateTag
(0029,"Silhouette Line V1.0",22)	UL	LineMeasurementDashLength	1	PrivateTag
(0029,"Silhouette Line V1.0",23)	UL	LinePointSpace	1	PrivateTag
(0029,"Silhouette Line V1.0",24)	FD	LinePoints	1	PrivateTag
(0029,"Silhouette Line V1.0",25)	UL	LineControlPointSize	1	PrivateTag
(0029,"Silhouette Line V1.0",26)	UL	LineControlPointSpace	1	PrivateTag
(0029,"Silhouette Line V1.0",27)	FD	LineControlPoints	1	PrivateTag
(0029,"Silhouette Line V1.0",28)	LT	LineLabel	1	PrivateTag
(0029,"Silhouette Line V1.0",29)	UL	LineDontSave	1	PrivateTag

(0029,"Silhouette ROI V1.0",11)	IS	ROIName	1	PrivateTag
(0029,"Silhouette ROI V1.0",12)	LT	ROINameFont	1	PrivateTag
(0029,"Silhouette ROI V1.0",13)	LT	ROINormalColor	1	PrivateTag
(0029,"Silhouette ROI V1.0",14)	UL	ROIFillPattern	1	PrivateTag
(0029,"Silhouette ROI V1.0",15)	UL	ROIBpSeg	1	PrivateTag
(0029,"Silhouette ROI V1.0",16)	UN	ROIBpSegPairs	1	PrivateTag
(0029,"Silhouette ROI V1.0",17)	UL	ROISeedSpace	1	PrivateTag
(0029,"Silhouette ROI V1.0",18)	UN	ROISeeds	1	PrivateTag
(0029,"Silhouette ROI V1.0",19)	UL	ROILineThickness	1	PrivateTag
(0029,"Silhouette ROI V1.0",20)	UL	ROILineStyle	1	PrivateTag
(0029,"Silhouette ROI V1.0",21)	UL	ROILineDashLength	1	PrivateTag
(0029,"Silhouette ROI V1.0",22)	UL	ROIInteractivity	1	PrivateTag
(0029,"Silhouette ROI V1.0",23)	UL	ROINamePosition	1	PrivateTag
(0029,"Silhouette ROI V1.0",24)	UL	ROINameDisplay	1	PrivateTag
(0029,"Silhouette ROI V1.0",25)	LT	ROILabel	1	PrivateTag
(0029,"Silhouette ROI V1.0",26)	UL	ROIShape	1	PrivateTag
(0029,"Silhouette ROI V1.0",27)	FD	ROIShapeTilt	1	PrivateTag
(0029,"Silhouette ROI V1.0",28)	UL	ROIShapePointsCount	1	PrivateTag
(0029,"Silhouette ROI V1.0",29)	UL	ROIShapePointsSpace	1	PrivateTag
(0029,"Silhouette ROI V1.0",30)	FD	ROIShapePoints	1	PrivateTag
(0029,"Silhouette ROI V1.0",31)	UL	ROIShapeControlPointsCount	1	PrivateTag
(0029,"Silhouette ROI V1.0",32)	UL	ROIShapeControlPointsSpace	1	PrivateTag
(0029,"Silhouette ROI V1.0",33)	FD	ROIShapeControlPoints	1	PrivateTag
(0029,"Silhouette ROI V1.0",34)	UL	ROIDontSave	1	PrivateTag

(0029,"Silhouette Sequence Ids V1.0",41)	SQ	Unknown	1	PrivateTag
(0029,"Silhouette Sequence Ids V1.0",42)	SQ	Unknown	1	PrivateTag
(0029,"Silhouette Sequence Ids V1.0",43)	SQ	Unknown	1	PrivateTag

(0029,"Silhouette V1.0",13)	UL	Unknown	1	PrivateTag
(0029,"Silhouette V1.0",14)	UL	Unknown	1	PrivateTag
(0029,"Silhouette V1.0",17)	UN	Unknown	1	PrivateTag
(0029,"Silhouette V1.0",18)	UN	Unknown	1	PrivateTag
(0029,"Silhouette V1.0",19)	UL	Unknown	1	PrivateTag
(0029,"Silhouette V1.0",1a)	UN	Unknown	1	PrivateTag
(0029,"Silhouette V1.0",1b)	UL	Unknown	1	PrivateTag
(0029,"Silhouette V1.0",1c)	UL	Unknown	1	PrivateTag
(0029,"Silhouette V1.0",1d)	UN	Unknown	1	PrivateTag
(0029,"Silhouette V1.0",1e)	UN	Unknown	1	PrivateTag
(0029,"Silhouette V1.0",21)	US	Unknown	1	PrivateTag
(0029,"Silhouette V1.0",22)	US	Unknown	1	PrivateTag
(0029,"Silhouette V1.0",23)	US	Unknown	1	PrivateTag
(0029,"Silhouette V1.0",24)	US	Unknown	1	PrivateTag
(0029,"Silhouette V1.0",25)	US	Unknown	1	PrivateTag
(0029,"Silhouette V1.0",27)	UN	Unknown	1	PrivateTag
(0029,"Silhouette V1.0",28)	UN	Unknown	1	PrivateTag
(0029,"Silhouette V1.0",29)	UN	Unknown	1	PrivateTag
(0029,"Silhouette V1.0",30)	UN	Unknown	1	PrivateTag
(0029,"Silhouette V1.0",52)	US	Unknown	1	PrivateTag
(0029,"Silhouette V1.0",53)	LT	Unknown	1	PrivateTag
(0029,"Silhouette V1.0",54)	UN	Unknown	1	PrivateTag
(0029,"Silhouette V1.0",55)	LT	Unknown	1	PrivateTag
(0029,"Silhouette V1.0",56)	LT	Unknown	1	PrivateTag
(0029,"Silhouette V1.0",57)	UN	Unknown	1	PrivateTag

(0135,"SONOWAND AS",10)	LO	UltrasoundScannerName	1	PrivateTag
(0135,"SONOWAND AS",11)	LO	TransducerSerial	1	PrivateTag
(0135,"SONOWAND AS",12)	LO	ProbeApplication	1	PrivateTag

(0017,"SVISION",00)	LO	ExtendedBodyPart	1	PrivateTag
(0017,"SVISION",10)	LO	ExtendedViewPosition	1	PrivateTag
(0017,"SVISION",F0)	IS	ImagesSOPClass	1	PrivateTag
(0019,"SVISION",00)	IS	AECField	1	PrivateTag
(0019,"SVISION",01)	IS	AECFilmScreen	1	PrivateTag
(0019,"SVISION",02)	IS	AECDensity	1	PrivateTag
(0019,"SVISION",10)	IS	PatientThickness	1	PrivateTag
(0019,"SVISION",18)	IS	BeamDistance	1	PrivateTag
(0019,"SVISION",20)	IS	WorkstationNumber	1	PrivateTag
(0019,"SVISION",28)	IS	TubeNumber	1	PrivateTag
(0019,"SVISION",30)	IS	BuckyGrid	1	PrivateTag
(0019,"SVISION",34)	IS	Focus	1	PrivateTag
(0019,"SVISION",38)	IS	Child	1	PrivateTag
(0019,"SVISION",40)	IS	CollimatorDistanceX	1	PrivateTag
(0019,"SVISION",41)	IS	CollimatorDistanceY	1	PrivateTag
(0019,"SVISION",50)	IS	CentralBeamHeight	1	PrivateTag
(0019,"SVISION",60)	IS	BuckyAngle	1	PrivateTag
(0019,"SVISION",68)	IS	CArmAngle	1	PrivateTag
(0019,"SVISION",69)	IS	CollimatorAngle	1	PrivateTag
(0019,"SVISION",70)	IS	FilterNumber	1	PrivateTag
(0019,"SVISION",74)	LO	FilterMaterial1	1	PrivateTag
(0019,"SVISION",75)	LO	FilterMaterial2	1	PrivateTag
(0019,"SVISION",78)	DS	FilterThickness1	1	PrivateTag
(0019,"SVISION",79)	DS	FilterThickness2	1	PrivateTag
(0019,"SVISION",80)	IS	BuckyFormat	1	PrivateTag
(0019,"SVISION",81)	IS	ObjectPosition	1	PrivateTag
(0019,"SVISION",90)	LO	DeskCommand	1	PrivateTag
(0019,"SVISION",A0)	DS	ExtendedExposureTime	1	PrivateTag
(0019,"SVISION",A1)	DS	ActualExposureTime	1	PrivateTag
(0019,"SVISION",A8)	DS	ExtendedXRayTubeCurrent	1	PrivateTag
(0021,"SVISION",00)	DS	NoiseReduction	1	PrivateTag
(0021,"SVISION",01)	DS	ContrastAmplification	1	PrivateTag
(0021,"SVISION",02)	DS	EdgeContrastBoosting	1	PrivateTag
(0021,"SVISION",03)	DS	LatitudeReduction	1	PrivateTag
(0021,"SVISION",10)	LO	FindRangeAlgorithm	1	PrivateTag
(0021,"SVISION",11)	DS	ThresholdCAlgorithm	1	PrivateTag
(0021,"SVISION",20)	LO	SensometricCurve	1	PrivateTag
(0021,"SVISION",30)	DS	LowerWindowOffset	1	PrivateTag
(0021,"SVISION",31)	DS	UpperWindowOffset	1	PrivateTag
(0021,"SVISION",40)	DS	MinPrintableDensity	1	PrivateTag
(0021,"SVISION",41)	DS	MaxPrintableDensity	1	PrivateTag
(0021,"SVISION",90)	DS	Brightness	1	PrivateTag
(0021,"SVISION",91)	DS	Contrast	1	PrivateTag
(0021,"SVISION",92)	DS	ShapeFactor	1	PrivateTag
(0023,"SVISION",00)	LO	ImageLaterality	1	PrivateTag
(0023,"SVISION",01)	IS	LetterPosition	1	PrivateTag
(0023,"SVISION",02)	IS	BurnedInAnnotation	1	PrivateTag
(0023,"SVISION",03)	LO	Unknown	1	PrivateTag
(0023,"SVISION",F0)	IS	ImageSOPClass	1	PrivateTag
(0025,"SVISION",00)	IS	OriginalImage	1	PrivateTag
(0025,"SVISION",01)	IS	NotProcessedImage	1	PrivateTag
(0025,"SVISION",02)	IS	CutOutImage	1	PrivateTag
(0025,"SVISION",03)	IS	DuplicatedImage	1	PrivateTag
(0025,"SVISION",04)	IS	StoredImage	1	PrivateTag
(0025,"SVISION",05)	IS	RetrievedImage	1	PrivateTag
(0025,"SVISION",06)	IS	RemoteImage	1	PrivateTag
(0025,"SVISION",07)	IS	MediaStoredImage	1	PrivateTag
(0025,"SVISION",08)	IS	ImageState	1	PrivateTag
(0025,"SVISION",20)	LO	SourceImageFile	1	PrivateTag
(0025,"SVISION",21)	UI	Unknown	1	PrivateTag
(0027,"SVISION",00)	IS	NumberOfSeries	1	PrivateTag
(0027,"SVISION",01)	IS	NumberOfStudies	1	PrivateTag
(0027,"SVISION",10)	DT	OldestSeries	1	PrivateTag
(0027,"SVISION",11)	DT	NewestSeries	1	PrivateTag
(0027,"SVISION",12)	DT	OldestStudy	1	PrivateTag
(0027,"SVISION",13)	DT	NewestStudy	1	PrivateTag

(0009,"TOSHIBA_MEC_1.0",01)	LT	Unknown	1	PrivateTag
(0009,"TOSHIBA_MEC_1.0",02)	US	Unknown	1-n	PrivateTag
(0009,"TOSHIBA_MEC_1.0",03)	US	Unknown	1-n	PrivateTag
(0009,"TOSHIBA_MEC_1.0",04)	US	Unknown	1-n	PrivateTag
(0011,"TOSHIBA_MEC_1.0",01)	LT	Unknown	1	PrivateTag
(0011,"TOSHIBA_MEC_1.0",02)	US	Unknown	1-n	PrivateTag
(0019,"TOSHIBA_MEC_1.0",01)	US	Unknown	1-n	PrivateTag
(0019,"TOSHIBA_MEC_1.0",02)	US	Unknown	1-n	PrivateTag
(0021,"TOSHIBA_MEC_1.0",01)	US	Unknown	1-n	PrivateTag
(0021,"TOSHIBA_MEC_1.0",02)	US	Unknown	1-n	PrivateTag
(0021,"TOSHIBA_MEC_1.0",03)	US	Unknown	1-n	PrivateTag
(7ff1,"TOSHIBA_MEC_1.0",01)	US	Unknown	1-n	PrivateTag
(7ff1,"TOSHIBA_MEC_1.0",02)	US	Unknown	1-n	PrivateTag
(7ff1,"TOSHIBA_MEC_1.0",03)	US	Unknown	1-n	PrivateTag
(7ff1,"TOSHIBA_MEC_1.0",10)	US	Unknown	1-n	PrivateTag

(0019,"TOSHIBA_MEC_CT_1.0",01)	IS	Unknown	1	PrivateTag
(0019,"TOSHIBA_MEC_CT_1.0",02)	IS	Unknown	1	PrivateTag
(0019,"TOSHIBA_MEC_CT_1.0",03)	US	Unknown	1-n	PrivateTag
(0019,"TOSHIBA_MEC_CT_1.0",04)	LT	Unknown	1	PrivateTag
(0019,"TOSHIBA_MEC_CT_1.0",05)	LT	Unknown	1	PrivateTag
(0019,"TOSHIBA_MEC_CT_1.0",06)	US	Unknown	1-n	PrivateTag
(0019,"TOSHIBA_MEC_CT_1.0",07)	US	Unknown	1-n	PrivateTag
(0019,"TOSHIBA_MEC_CT_1.0",08)	LT	OrientationHeadFeet	1	PrivateTag
(0019,"TOSHIBA_MEC_CT_1.0",09)	LT	ViewDirection	1	PrivateTag
(0019,"TOSHIBA_MEC_CT_1.0",0a)	LT	OrientationSupineProne	1	PrivateTag
(0019,"TOSHIBA_MEC_CT_1.0",0b)	DS	Unknown	1	PrivateTag
(0019,"TOSHIBA_MEC_CT_1.0",0c)	US	Unknown	1-n	PrivateTag
(0019,"TOSHIBA_MEC_CT_1.0",0d)	TM	Time	1	PrivateTag
(0019,"TOSHIBA_MEC_CT_1.0",0e)	DS	Unknown	1	PrivateTag
(7ff1,"TOSHIBA_MEC_CT_1.0",01)	US	Unknown	1-n	PrivateTag
(7ff1,"TOSHIBA_MEC_CT_1.0",02)	US	Unknown	1-n	PrivateTag
(7ff1,"TOSHIBA_MEC_CT_1.0",03)	IS	Unknown	1	PrivateTag
(7ff1,"TOSHIBA_MEC_CT_1.0",04)	IS	Unknown	1	PrivateTag
(7ff1,"TOSHIBA_MEC_CT_1.0",05)	US	Unknown	1-n	PrivateTag
(7ff1,"TOSHIBA_MEC_CT_1.0",07)	US	Unknown	1-n	PrivateTag
(7ff1,"TOSHIBA_MEC_CT_1.0",08)	US	Unknown	1-n	PrivateTag
(7ff1,"TOSHIBA_MEC_CT_1.0",09)	US	Unknown	1-n	PrivateTag
(7ff1,"TOSHIBA_MEC_CT_1.0",0a)	LT	Unknown	1	PrivateTag
(7ff1,"TOSHIBA_MEC_CT_1.0",0b)	US	Unknown	1-n	PrivateTag
(7ff1,"TOSHIBA_MEC_CT_1.0",0c)	US	Unknown	1-n	PrivateTag
(7ff1,"TOSHIBA_MEC_CT_1.0",0d)	US	Unknown	1-n	PrivateTag
#
# end of private.dic
#
`
