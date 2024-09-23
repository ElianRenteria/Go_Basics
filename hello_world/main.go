package main

import (
	"fmt"
)

const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

type Salutation struct {
	name string
	language string
}

func Hello(salutation Salutation) string {
	if salutation.name == "" {
		salutation.name = "World!"
	} 
	/*if salutation.language == "Spanish" {
		return "Hola, "+salutation.name
	}
	if salutation.language == "French" {
		return "Bonjour, "+salutation.name
	}
	return hello+salutation.name*/
	return greetingPrefix(salutation.language) + salutation.name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello(Salutation{name:"Elian"}))
}