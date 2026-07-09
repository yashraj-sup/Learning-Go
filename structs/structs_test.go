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
	t.Run("rectangles", func(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Area() //method callinggggg
	want := 100.00
		if got != want {
		t.Errorf("got %g want %g", got, want)
		}
	})
	t.Run("circles", func(t *testing.T) {
		circle := Circle{4.0}
		got := circle.Area() //method calling
		want := 50.26548245743669
			if got != want {
			t.Errorf("got %g  want %g", got, want)
			}
	})
}
