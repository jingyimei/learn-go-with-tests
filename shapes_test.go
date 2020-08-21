package learn_go_with_tests

import (
	"math"
	"testing"
)

func compareResult(got float64, want float64, t *testing.T) {
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestPerimeter(t *testing.T) {
	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Perimeter()
		want := 40.0
		compareResult(got, want, t)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Perimeter()
		want := math.Pi * 2 * 10.0
		compareResult(got, want, t)
	})
}

func TestArea(t *testing.T) {
	areaTests := [] struct {
		shape Shape
		want float64
	} {
		{Rectangle{10, 10}, 100.0},
		{Circle{10},314.1592653589793},
		{Triangle{12.0, 6.0}, 36.0},
	}
	
	for _, tt := range(areaTests) {
		area := tt.shape.Area()
		compareResult(area, tt.want, t)
	}
}
