package main
import "testing"
func TestHello(t *testing.T){
	t.Run ("in Japanese", func (t *testing.T){
		got := Hello ("Eldoie", "Japanese")
		want := "Konnichiwa, Eldoie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func (t *testing.T){
		got := Hello ("Eldoie", "French")
		want := "Bonjour, Eldoie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func (t *testing.T){
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("empty string defaults to 'world'",func (t *testing.T){
		got:=Hello("","")
		want:="Hello, World"
		assertCorrectMessage(t,got,want)
	})
}
func assertCorrectMessage(t testing.TB, got, want string){
	t.Helper()
	if got!=want{
		t.Errorf("got %q want %q", got, want)
	}
}
