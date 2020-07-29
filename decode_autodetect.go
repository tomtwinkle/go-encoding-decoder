package main

import (
	"bytes"
	"errors"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
	"unicode/utf8"
)

// Converts to UTF-8.
// Charset (UTF-8, Shift-JIS, EUC-JP, ISO-2022-JP) is automatically detected.
func DecodeAutoDetect(src []byte) (string, error) {
	for _, enc := range encodings {
		e, _ := charset.Lookup(enc)
		if e == nil {
			continue
		}
		var buf bytes.Buffer
		r := transform.NewWriter(&buf, e.NewDecoder())
		_, err := r.Write(src)
		if err != nil {
			continue
		}
		err = r.Close()
		if err != nil {
			continue
		}
		f := buf.Bytes()
		if isInvalidRune(f) {
			continue
		}
		if utf8.Valid(f) {
			if hasBom(f) {
				f = stripBom(f)
			}
			return string(f), nil
		}
	}
	return string(src), errors.New("could not determine character code")
}

var utf8bom = []byte{239, 187, 191}

// check have UTF-8 BOM
func hasBom(in []byte) bool {
	return bytes.HasPrefix(in, utf8bom)
}

// strip UTF-8 BOM
func stripBom(in []byte) []byte {
	return bytes.TrimPrefix(in, utf8bom)
}

func isInvalidRune(in []byte) bool {
	cb := in
	for len(cb) > 0 {
		if utf8.RuneStart(cb[0]) {
			r, size := utf8.DecodeRune(cb)
			if r == utf8.RuneError {
				return true
			}
			cb = cb[size:]
		} else {
			cb = cb[1:]
		}
	}
	return false
}
