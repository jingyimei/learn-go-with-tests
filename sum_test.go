package learn_go_with_tests

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum up a slice", func(t *testing.T) {
		mySlice := []int{1,2,3,4}
		sum := sum(mySlice)
		if sum != 10 {
			t.Errorf("wrong! got %d, want %d given %v", sum, 6, mySlice)
		}
	})

	t.Run("sum up multiple slices", func(t *testing.T) {
		mySlice1 := []int{1,2,3,4}
		mySlice2 := []int{1, 2}
		sum := sumAll(mySlice1, mySlice2)
		want := []int {10, 3}
		if !reflect.DeepEqual(sum, want) {
			t.Errorf("wrong! got %d, want %v given %v and %v", sum, want, mySlice1, mySlice2)
		}
	})

	t.Run("sum up tails", func(t *testing.T) {
		mySlice1 := []int{1,2,3,4}
		mySlice2 := []int{1,2}
		sum := sumAllTails(mySlice1, mySlice2)
		want := []int {9, 2}
		if !reflect.DeepEqual(sum, want) {
			t.Errorf("wrong! got %d, want %v given %v and %v", sum, want, mySlice1, mySlice2)
		}
	})

	t.Run("sum up tails when slice is empty", func(t *testing.T) {
		mySlice1 := []int{}
		mySlice2 := []int{1,2}
		sum := sumAllTails(mySlice1, mySlice2)
		want := []int {0, 2}
		if !reflect.DeepEqual(sum, want) {
			t.Errorf("wrong! got %d, want %v given %v and %v", sum, want, mySlice1, mySlice2)
		}
	})
}


