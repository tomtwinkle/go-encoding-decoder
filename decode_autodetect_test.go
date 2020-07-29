package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeAutoDetect(t *testing.T) {
	t.Run("UTF-8", func(t *testing.T) {
		src := MustReadFile(t, "fixture/utf-8.txt")
		expected := MustReadFile(t, "fixture/utf-8.txt")
		actual, err := DecodeAutoDetect(src)
		assert.NoError(t, err)
		assert.Equal(t, string(expected), actual)
	})

	t.Run("UTF-8 BOM", func(t *testing.T) {
		src := MustReadFile(t, "fixture/utf-8bom.txt")
		expected := MustReadFile(t, "fixture/utf-8bom.txt")
		actual, err := DecodeAutoDetect(src)
		assert.NoError(t, err)
		assert.Equal(t, string(expected), actual)
	})

	t.Run("ISO-2022-JP(JIS)", func(t *testing.T) {
		src := MustReadFile(t, "fixture/jis.txt")
		expected := MustReadFile(t, "fixture/utf-8.txt")
		actual, err := DecodeAutoDetect(src)
		assert.NoError(t, err)
		assert.Equal(t, string(expected), actual)
	})

	t.Run("Shift_JIS", func(t *testing.T) {
		src := MustReadFile(t, "fixture/sjis.txt")
		expected := MustReadFile(t, "fixture/utf-8.txt")
		actual, err := DecodeAutoDetect(src)
		assert.NoError(t, err)
		assert.Equal(t, string(expected), actual)
	})

	t.Run("EUC-JP", func(t *testing.T) {
		src := MustReadFile(t, "fixture/eucjp.txt")
		expected := MustReadFile(t, "fixture/utf-8.txt")
		actual, err := DecodeAutoDetect(src)
		assert.NoError(t, err)
		assert.Equal(t, string(expected), actual)
	})
}
