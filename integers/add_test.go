package integers

import (
	"fmt"
	"go-learn-tdd/utils"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("2 + 2", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		utils.AssertCorrectMessage(t, sum, expected)
	})

	t.Run("-1 + -1", func(t *testing.T) {
		sum := Add(-1, -1)
		expected := -2

		utils.AssertCorrectMessage(t, sum, expected)
	})

	t.Run("-36 + 36", func(t *testing.T) {
		sum := Add(-36, 36)
		expected := 0

		utils.AssertCorrectMessage(t, sum, expected)
	})
}

func ExampleAdd() {
	sum := Add(10, 15)
	fmt.Println(sum)
	// Output: 25
}
