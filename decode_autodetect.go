package main

import (
	"errors"
	"fmt"
	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
)

// Converts to UTF-8.
// Charset automatically detected.
func DecodeAutoDetect(src []byte) (string, error) {
	d := chardet.NewHtmlDetector()
	r, err := d.DetectBest(src)
	if err != nil {
		return string(src), err
	}
	e, _ := charset.Lookup(r.Charset)
	if e == nil {
		return string(src), errors.New(fmt.Sprintf("invalid charset [%s]", r.Charset))
	}
	decodeStr, _, err := transform.Bytes(
		e.NewDecoder(),
		src,
	)
	if err != nil {
		return string(src), err
	}
	return string(decodeStr), nil
}
