package rectangle

import (
	"fmt"
	"math"
)

func init() {
	fmt.Println("rectangle package initialized")
}
func A(x, y float64) float64 {
	return x * y
}

func P(x, y float64) float64 {
	return math.Sqrt((x * x) + (y * y))
}
