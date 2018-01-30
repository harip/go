package main

import (
	"fmt"
	"reflect"
)

var (
	name   = "a"
	course = "a"
	module float64
)

func main() {
	fmt.Println("name is", name, reflect.TypeOf(name))
	fmt.Println(module)
}
