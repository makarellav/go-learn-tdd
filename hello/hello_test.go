package hello

import (
	"go-learn-tdd/utils"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("to a person", func(t *testing.T) {
		got := Hello("Vladyslav", "")
		want := "Hello, Vladyslav!"

		utils.AssertCorrectMessage(t, got, want)
	})

	t.Run("empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"

		utils.AssertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Andrew", "Spanish")
		want := "Hola, Andrew!"

		utils.AssertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Max", "French")
		want := "Bonjour, Max!"

		utils.AssertCorrectMessage(t, got, want)
	})
}
