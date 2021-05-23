package main

import "testing"

func TestHello(t *testing.T)  {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		// By t.Helper() when test fails the line number reported will be in our function call rather than inside our test helper.
		t.Helper()
		if got != want {
			t.Errorf("got %q, want: %q", got, want)
		}
	}

	t.Run("saying hello world to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello world, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello world to empty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})
}