package iteration
import (
	"strings"
	"testing"
)

func Repeat(character string, repeatCount int) string{
	var repeated strings.Builder
	for i :=0; i< repeatCount; i++{
		repeated.WriteString(character)
	}
	return repeated.String()
}
func TestRepeatSeventimes(t *testing.T){
	repeated := Repeat("a",7)
	expected := "aaaaaaa"
	if repeated != expected{
		t.Errorf("expected %q but got %q", expected , repeated)
	}
}
func TestRepeatThreetimes(t *testing.T){
	repeated := Repeat("a" , 3)
	expected := "aaa"
	if repeated != expected{
		t.Errorf("expected %q but got %q", expected , repeated)
	}
}
func TestRepeat(t *testing.T) {
	repeated := Repeat("a" , 5)
	expected := "aaaaa"
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected , repeated)
	}
}
func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat ("a" , 5)
	}
}
