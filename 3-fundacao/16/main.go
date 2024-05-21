package main

// Type Assertion

func main() {
	var minhaVar interface{} = "Wesley Willians"
	println(minhaVar)          // (0x1001fcd40,0x100208090)
	println(minhaVar.(string)) // Wesley

	res, ok := minhaVar.(int)
	println(res, ok) // 0 false

	res2 := minhaVar.(int) // nesse caso vai dar panic por que não pegamos um ok2 pra validar esse erro.
	println(res2)          // panic: interface conversion: interface {} is string, not int
}
