package main

import (
	"errors"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
)

var encodings = []string{
	"iso-2022-jp",
	"euc-jp",
	"utf-8",
	"sjis",
}

// Converts the specified charset to UTF-8.
func Decode(src []byte, charSet string) (string, error) {
	e, _ := charset.Lookup(charSet)
	if e == nil {
		return string(src), errors.New(fmt.Sprintf("invalid charset [%s]", charSet))
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
