package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Gyaneshwar", "Nepali")
		want := "Namaste, Gyaneshwar"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello World' when an empty string is passed arguments", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// t.Helper() tells the test suite that it's a helper function and reports the line number from actual function call rather than from the helper
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
