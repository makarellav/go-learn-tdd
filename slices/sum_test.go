package slices

import (
	"fmt"
	"slices"
	"testing"
)

func checkSums(t testing.TB, got, want []int) {
	t.Helper()

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSum(t *testing.T) {
	t.Run("sum some numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("some an empty slice", func(t *testing.T) {
		var numbers []int

		got := Sum(numbers)
		want := 0

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("sum non-empty slices", func(t *testing.T) {
		got := SumAll([]int{1, 9}, []int{3, 4}, []int{-8, 0})
		want := []int{10, 7, -8}

		checkSums(t, got, want)
	})

	t.Run("sum empty slices", func(t *testing.T) {
		got := SumAll([]int{}, []int{}, []int{}, []int{})
		want := []int{0, 0, 0, 0}

		checkSums(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("sum tails for non-empty slices", func(t *testing.T) {
		got := SumAllTails([]int{3, 7, 9}, []int{13, -21, 21, 0}, []int{0, 54, 46, 100, -100})
		want := []int{16, 0, 100}

		checkSums(t, got, want)
	})

	t.Run("sum tails for empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{}, []int{}, []int{})
		want := []int{0, 0, 0, 0}

		checkSums(t, got, want)
	})
}

func ExampleSum() {
	numbersToSum := []int{1, 2, 3, 4, 5}
	sum := Sum(numbersToSum)

	fmt.Println(sum)
	// Output: 15
}

func ExampleSumAll() {
	slicesToSum := [][]int{{1, 2, 3, 4, 5}, {-10, 0, 15}, {100, -99}, {}}
	allSums := SumAll(slicesToSum...)

	fmt.Println(allSums)
	// Output: [15 5 1 0]
}

func ExampleSumAllTails() {
	slicesToSum := [][]int{{1, 2, 3, 4, 5}, {-10, 0, 15}, {100, -99}, {}}
	allTailsSums := SumAllTails(slicesToSum...)

	fmt.Println(allTailsSums)
	// Output: [14 15 -99 0]
}
