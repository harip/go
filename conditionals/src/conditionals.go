package main

import (
	"fmt"
)

func main() {

	if a, b := 1, 2; a > b {
		fmt.Println("1>2")
	} else if b < a {
		fmt.Println("2>1")
	} else {
		fmt.Println("2=1")
	}
}
