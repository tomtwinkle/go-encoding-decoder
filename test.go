package main

import (
	"io/ioutil"
	"testing"
)

func MustReadFile(t *testing.T, filename string) []byte {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	return bytes
}
