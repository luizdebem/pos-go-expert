package main

func main() {
	// Isso dá deadlock: enchendo o channel duas vezes.
	// ch := make(chan string)
	// ch <- "hello"
	// ch <- "world"
	// println(<-ch)
	// println(<-ch)

	// Isso não dá deadlock: channel de buffer 2 (size).
	ch := make(chan string, 2)
	ch <- "hello"
	ch <- "world"
	println(<-ch)
	println(<-ch)
}
