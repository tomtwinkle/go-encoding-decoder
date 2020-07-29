package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecode(t *testing.T) {
	t.Run("UTF-8", func(t *testing.T) {
		src := MustReadFile(t, "fixture/utf-8.txt")
		expected := MustReadFile(t, "fixture/utf-8.txt")

		actual, err := Decode(src, "utf-8")
		assert.NoError(t, err)
		assert.Equal(t, string(expected), actual)

		actual, err = Decode(src, "UTF-8")
		assert.NoError(t, err)
		assert.Equal(t, string(expected), actual)
	})

	t.Run("ISO-2022-JP(JIS)", func(t *testing.T) {
		src := MustReadFile(t, "fixture/jis.txt")
		expected := MustReadFile(t, "fixture/utf-8.txt")

		actual, err := Decode(src, "iso-2022-jp")
		assert.NoError(t, err)
		assert.Equal(t, string(expected), actual)

		actual, err = Decode(src, "ISO-2022-JP")
		assert.NoError(t, err)
		assert.Equal(t, string(expected), actual)
	})

	t.Run("Shift_JIS", func(t *testing.T) {
		src := MustReadFile(t, "fixture/sjis.txt")
		expected := MustReadFile(t, "fixture/utf-8.txt")

		actual, err := Decode(src, "Shift_JIS")
		assert.NoError(t, err)
		assert.Equal(t, string(expected), actual)

		actual, err = Decode(src, "shift_jis")
		assert.NoError(t, err)
		assert.Equal(t, string(expected), actual)

		actual, err = Decode(src, "sjis")
		assert.NoError(t, err)
		assert.Equal(t, string(expected), actual)
	})

	t.Run("EUC-JP", func(t *testing.T) {
		src := MustReadFile(t, "fixture/eucjp.txt")
		expected := MustReadFile(t, "fixture/utf-8.txt")

		actual, err := Decode(src, "euc-jp")
		assert.NoError(t, err)
		assert.Equal(t, string(expected), actual)

		actual, err = Decode(src, "EUC-JP")
		assert.NoError(t, err)
		assert.Equal(t, string(expected), actual)
	})
}
