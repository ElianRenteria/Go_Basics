package main

import (
	"fmt"
)

func main() {
	animals := []string{
		"dog",
		"cat",
	}

	animals = append(animals, "bird")

	// animals = slices.Delete(animals, 0, 1)
	// fmt.Println(animals)
	/*for i:=0;i < len(animals); i++ {
		fmt.Println(animals[i])
	}*/

	/*for index, value := range animals {
		fmt.Println(index, value)
	}*/

	for num := range 10 {
		fmt.Println(num)
	}
}