package main

import (
	"bytes"
	"io"
	"testing"
)

func TestParseEvents(t *testing.T) {
	buf1 := bytes.NewBufferString(`[{"key":"hoge","value":"foo"},{"key":"bar","value":"baz"}]`)
	buf2 := bytes.NewBufferString(`{"key":"hoge","value":"foo"}\n`)
	_, err := ParseEvents(buf1)

	if err != nil {
		t.Fatal("error raised")
	}
	_, err = ParseEvents(buf2)

	if err != nil && err != io.EOF {
		t.Fatalf("error raised: %v", err)
	}
}
