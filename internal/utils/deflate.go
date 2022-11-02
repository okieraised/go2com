package utils

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

func DeflateGzip(b []byte) ([]byte, error) {
	br := bytes.NewReader(b)
	g, err := gzip.NewReader(br)
	if err != nil {
		return nil, err
	}
	defer g.Close()

	p, err := ioutil.ReadAll(g)
	if err != nil {
		return nil, err
	}

	return p, nil
}
