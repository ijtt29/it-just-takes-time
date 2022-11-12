package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const frenchhHelloPrefix = "Bonjour "
const spanishHelloPrefix = "Hola "
const englishHelloPrefix = "Hello "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchhHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("World", ""))
}
