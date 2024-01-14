package main

import (
	"fmt"
)

const (
	spanish = "Spanish"
	french = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
)

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func Hello(name, language string) string {
	prefix := greetingPrefix(language)
	if name == "" {
		name = "World"
	}

	return prefix + name + "!"
}

func main() {
	fmt.Println(Hello("World", ""))
}
