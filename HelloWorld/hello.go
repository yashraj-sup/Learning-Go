package main
import "fmt"
const (
	japanese = "Japanese"
	french = "French"
	spanish = "Spanish"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
	japaneseHelloPrefix = "Konnichiwa, "
)
func Hello(name string,language string) string{
	if name ==""{
		name= "World"
	}
	return greetingPrefix(language) + name
}
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish :
		prefix = spanishHelloPrefix
	case french :
		prefix = frenchHelloPrefix
	case japanese :
		prefix = japaneseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return 
}
	func main(){
	fmt.Println(Hello("world","English"))
}



