package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestPrintMD5(t *testing.T) {
	in := strings.NewReader("go") // io.Reader
	buf := &bytes.Buffer{} // io.Writer

	printMD5(in, buf)

	want := "34d1f91fb2e514b8576fab1a75a89a6b"
	got := buf.String()

	if got != want {
		t.Errorf("printMD5()=%s want %s", got, want)
	}
}