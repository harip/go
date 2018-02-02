package main

import (
	"fmt"
)

func main() {
	name := "a"
	fmt.Println(name)

	changeNameByRef(&name)

	fmt.Println(name)
}

func changeName(name string) string {
	name = "b"
	fmt.Println(name)
	return name
}
func changeNameByRef(name *string) string {
	*name = "b"
	fmt.Println(*name)
	return *name
}
