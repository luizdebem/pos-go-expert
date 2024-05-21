package main

import "fmt"

func main() {
	var x interface{} = 10
	var y interface{} = "Hello World."
	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("O tipo é %T e o valor é %v\n", t, t)
}

// usar moderadamente.
// isso existe porém hoje em dia tem generics.
