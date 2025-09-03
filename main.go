package main

import (
	"fmt"
	"slices"
)

func main() {
	animals := []string{"dog", "cat", "dog", "cat"}
	out := []string{}

	for _, animal := range animals {
		if !slices.Contains(out, animal) {
			out = append(out, animal)
		}
	}

	fmt.Printf("Start slice: %v, Out slice: %v", animals, out)
}
