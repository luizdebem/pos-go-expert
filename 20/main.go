package main

// Condicionais

func main() {
	a := 1
	b := 2
	c := 3

	// if a > b { // > >= == <= <
	// 	println(a)
	// } else {
	// 	println(a)
	// }

	// if a > b || c > a { // && ||
	// 	println("hello world")
	// }

	switch a {
	case 1:
		println(a)
	case 2:
		println(b)
	case 3:
		println(c)
	default:
		println(c)
	}

}
