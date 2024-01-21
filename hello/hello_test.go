package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("to a person", func(t *testing.T) {
		got := Hello("Vladyslav", "")
		want := "Hello, Vladyslav!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Andrew", "Spanish")
		want := "Hola, Andrew!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Max", "French")
		want := "Bonjour, Max!"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
