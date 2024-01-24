package smi

import (
	"testing"
)

func checkValue(t testing.TB, got float64, want float64) {
	t.Helper()

	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		name      string
		shape     Shape
		perimeter float64
	}{
		{"Rectangle", Rectangle{10, 15}, 50},
		{"Circle", Circle{5}, 31.41592653589793},
		{"Triangle", Triangle{10, 12, 15}, 37},
	}

	for _, tt := range perimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			checkValue(t, tt.shape.Perimeter(), tt.perimeter)
		})
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		area  float64
	}{
		{"Rectangle", Rectangle{15, 5}, 75},
		{"Circle", Circle{10}, 314.1592653589793},
		{"Triangle", Triangle{3, 4, 5}, 6},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			checkValue(t, tt.shape.Area(), tt.area)
		})
	}
}
