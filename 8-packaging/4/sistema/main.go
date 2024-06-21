package main

import (
	"fmt"

	"github.com/google/uuid"
	"luizdebem.com/math"
)

func main() {
	a := math.NewMath(10, 20)
	fmt.Println(a.Add())
	fmt.Println(a.Sub())

	fmt.Println(uuid.New().String())
}
