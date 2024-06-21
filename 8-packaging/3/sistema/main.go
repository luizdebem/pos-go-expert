package main

import (
	"fmt"

	"luizdebem.com/math"
)

func main() {
	a := math.NewMath(10, 20)
	fmt.Println(a.Add())
	fmt.Println(a.Sub())
}
