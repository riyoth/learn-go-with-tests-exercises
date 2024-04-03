package main

import "fmt"

const (
	englishHelloPrefix = "Hello "
	english = "English"
	spanishHelloPrefix = "Hola "
	spanish = "Spanish"
	frenchHelloPrefix = "Bonjour "
	french = "French"
)

const defaultName = "world"

func Hello(name string, language string) string {
	if name == "" {
		name = defaultName
	}

	return greetingPrefix(language)+ name
}
func greetingPrefix(language string) (prefix string){

	switch language {
		case english:
			prefix = englishHelloPrefix
		case french:
			prefix = frenchHelloPrefix
		case spanish:
			prefix = spanishHelloPrefix
		default:
			prefix = englishHelloPrefix
	}
	return

}

func main() {
	fmt.Println(Hello("chris", ""))
}
