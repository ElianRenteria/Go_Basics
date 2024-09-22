package main

import (
	"fmt"
)

const hello = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World!"
	}
	return hello+name
}

func main() {
	fmt.Println(Hello("Elian"))
}