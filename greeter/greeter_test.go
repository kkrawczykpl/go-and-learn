package main

import (
	"bytes"
	"testing"
)

func TestGreeter(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris!"

	if got != want {
		t.Errorf("Got: %q, want: %q", got, want)
	}
}