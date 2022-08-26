package tag

type DicomTag struct {
	Group   uint16
	Element uint16
}

func IsPrivateTag(group uint16) bool {
	return group%2 == 1
}
