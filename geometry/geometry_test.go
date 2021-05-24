package main

import "testing"

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want float64
	} {
		{Rectangle{10.0, 15.0}, 150.0},
		{Circle{10}, 314.1592653589793},
	}

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	for _, testCase := range areaTests {
		checkArea(t, testCase.shape, testCase.want)
	}
}