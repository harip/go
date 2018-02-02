package main

import (
	"fmt"
)

func main() {
	name := finishes(1, 2, 3, 4, 5, 3, 3, 3)
	fmt.Println(name)
}

func finishes(finishes ...int) int {
	best := finishes[0]
	for _, finish := range finishes {
		if finish > best {
			best = finish
		}
	}
	return best
}
