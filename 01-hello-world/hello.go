package main

import "fmt"

const englishHelloPrefix = "Hello"
const spanishHelloPrefix = "Hola"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	var greeting string
	if language == "Spanish" {
		greeting = spanishHelloPrefix
	} else {
		greeting = englishHelloPrefix
	}
	return greeting + ", " + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
