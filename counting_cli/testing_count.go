package main

import (
	"bytes"
	"testing"
)

func TestCount(t *testing.T) {
	word_test := bytes.NewBufferString("word1 and word3")
	exp := 3
	res := CountingWords(word_test, false)
	if exp != res {
		t.Errorf("Words mismatch")
	}
}