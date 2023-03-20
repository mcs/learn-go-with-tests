package main

import "fmt"

const englishHelloPrefix = "Hello"
const spanishHelloPrefix = "Hola"
const frenchHelloPrefix = "Bonjour"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	var greeting string
	if language == "Spanish" {
		greeting = spanishHelloPrefix
	} else if language == "French" {
		greeting = frenchHelloPrefix
	} else {
		greeting = englishHelloPrefix
	}
	return greeting + ", " + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
