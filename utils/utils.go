package utils

import "testing"

func AssertCorrectMessage(t testing.TB, got, want any) {
	t.Helper()

	if got != want {
		t.Errorf("expected %q but got %q", got, want)
	}
}
