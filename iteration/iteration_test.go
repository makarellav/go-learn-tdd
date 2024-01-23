package iteration

import (
	"fmt"
	"go-learn-tdd/utils"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat 'a' 5 times", func(t *testing.T) {
		repeated := Repeat(5, "a")
		expected := "aaaaa"

		utils.AssertCorrectMessage(t, repeated, expected)
	})

	t.Run("repeat 'b' 0 times", func(t *testing.T) {
		repeated := Repeat(0, "b")
		expected := ""

		utils.AssertCorrectMessage(t, repeated, expected)
	})

	t.Run("repeat 'c' -1 time", func(t *testing.T) {
		repeated := Repeat(-1, "c")
		expected := ""

		utils.AssertCorrectMessage(t, repeated, expected)
	})
}

func ExampleRepeat() {
	repeated := Repeat(10, "t")
	fmt.Println(repeated)

	// Output: tttttttttt
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat(10000, "a")
	}
}

func BenchmarkRepeatStringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatStringBuilder(10000, "a")
	}
}
