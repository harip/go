package main

import (
	"fmt"
	"strings"
)

func main() {
	module := "function basics"
	author := "abc def"
	fmt.Println(converter(module, author))
	fmt.Println(module)
}

func converter(module, author string) (s1, s2 string) {
	module = strings.Title(module)
	author = strings.Title(author)
	return module, author
}
