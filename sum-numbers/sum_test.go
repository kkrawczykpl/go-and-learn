package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got int, want int) {
		// By t.Helper() when test fails the line number reported will be in our function call rather than inside our test helper.
		t.Helper()
		if got != want {
			t.Errorf("Got: %d, want: %d", got, want)
		}
	}

	t.Run("Sum up slice of 10 numbers 1-10", func(t *testing.T) {
		numbers := []int{1,2,3,4,5,6,7,8,9,10}

		got := Sum(numbers)
		want := 55
		assertCorrectMessage(t, got, want)
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}