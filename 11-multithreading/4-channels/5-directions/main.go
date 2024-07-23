package main

import "fmt"

// ⠀⠀⠀⠀⠀⠀⣀⣠⣀⡀⠀⠀⠀⠀⠀⠀
// ⠀⠀⠀⠀⠀⣾⡿⠛⠿⢿⣇⠀⠀⠀⠀⠀
// ⠀⠀⠀⠀⠀⡏⢤⠄⣴⠎⣿⠀⠀⠀⠀⠀
// ⠀⠀⠀⠀⠀⠧⢀⣒⣆⠰⠿⠀⠀⠀⠀⠀
// ⠀⠈⠛⢶⣤⣄⠢⣍⣠⠜⠀⣠⣤⣤⠶⠆
// ⠀⣠⣴⣶⣾⣿⣿⡆⠙⠙⣿⣿⣿⣶⣤⠀
// ⣼⣿⣿⣿⡿⠟⠋⠀⠠⠺⢿⣿⣿⣿⣿⣦
// ⡞⠡⠋⠁⠀⠀⣠⣾⣧⠀⠀⠙⢿⡟⠻⣿
// ⠇⠀⠀⢀⣤⣾⣿⣿⣿⣷⡄⠀⠀⠉⠀⠿
// ⠀⠑⠂⣿⣿⣿⡏⣯⢑⣿⣿⣦⡀⠀⠐⠀
func receba(name string, hello chan<- string) {
	hello <- name
}

func ler(data <-chan string) {
	fmt.Println(<-data)
}

func main() {
	hello := make(chan string)
	go receba("José", hello)
	ler(hello)
}
