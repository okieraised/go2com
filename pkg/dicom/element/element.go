package element

type Element struct {
	Tag                 string `json:"tag"`
	ValueRepresentation string `json:"vr"`
	ValueLength         uint32 `json:"valueLength"`
	Value               string `json:"value"`
}
