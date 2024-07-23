// Channels

// - Fazer comunicação entre threads
// - Segurança para uma thread saber o momento em que ela pode trabalhar com um determinado dado.

package main

// Thread 1 (main)
func main() {
	// Thread 1
	channel := make(chan string) // Canal vazio

	// Thread 2
	go func() {
		channel <- "Hello World!" // Canal preenchido.
	}()

	// Thread 1
	msg := <-channel // Canal esvazia.
	println(msg)

}
