package main

import (
	"fmt"
	"math"
)

func main() {
	a, b := 145.8, 432.8
	c := math.Min(a, b)
	fmt.Println("min value is", c)
}
