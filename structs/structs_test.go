package structs
import (
	"testing"
)
func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
func TestArea(t *testing.T) {
	areaTests := []struct{
		name string
		shape Shape
		hasArea float64
	}{
		{name:"Rectangle", shape:Rectangle{Width:10.0, Height:10.0}, hasArea:100.00},
		{name:"Circle", shape:Circle{Radius:4.0}, hasArea: 50.26548245743669},
		{name:"Triangle", shape:Triangle{Base:4, Height: 6}, hasArea: 12.0},
	}
	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf(" %#v got %g want %g", got, tt.shape, tt.hasArea)
			}
		})
	}
}


