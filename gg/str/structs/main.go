package main

import (
	"fmt"

	"structs/computer"
)

func main() {
	var spec computer.Spec
	spec.Maker = "apple"
	spec.Price = 1000
	fmt.Println(spec)
}
