package main

import "fmt"

const englishHelloPrefix = "Hello"
const spanishHelloPrefix = "Hola"
const frenchHelloPrefix = "Bonjour"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + ", " + name
}

func greetingPrefix(language string) string {
	var greeting string
	switch language {
	case "Spanish":
		greeting = spanishHelloPrefix
	case "French":
		greeting = frenchHelloPrefix
	default:
		greeting = englishHelloPrefix
	}
	return greeting
}

func main() {
	fmt.Println(Hello("world", ""))
}
