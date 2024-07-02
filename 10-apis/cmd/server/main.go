// Golang Standards Project Layout no GitHub.

package main

import "10-apis/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBHost)
}
