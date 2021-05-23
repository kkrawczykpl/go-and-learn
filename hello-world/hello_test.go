package main

import "testing"

func TestHello(t *testing.T)  {
	got := Hello("Chris")
	want := "Hello world, Chris"

	if got != want {
		t.Errorf("got %q, want: %q", got, want)
	}
}