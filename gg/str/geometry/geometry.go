// geometry.go

package main

import (
	"fmt"
	"geometry/rectangle"
	"log"
)

// 包级别变量
var c, d float64 = 6, 7

func init() {
	println("main package initialized")
	if c < 0 {
		log.Fatal("c is less than zero")
	}
	if d < 0 {
		log.Fatal("d is less than zero")
	}
}
func main() {
	//var c, d float64 = 6, 7
	fmt.Println("Geometrical shape properties")
	fmt.Printf("area of rectangle %.2f", rectangle.A(c, d))
}
