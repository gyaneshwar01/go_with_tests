package main

import "fmt"

const (
	spanish            = "Spanish"
	nepali             = "Nepali"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	nepaliHelloPrefix  = "Namaste, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return gettingPrefix(language) + name
}

func gettingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case nepali:
		prefix = nepaliHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", ""))
}
