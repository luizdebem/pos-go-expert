package main

func panic1() {
	panic("panic1")
}

func panic2() {
	panic("panic2")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			if r == "panic1" {
				println("recover: panic1")
			}
			if r == "panic2" {
				println("recover: panic2")
			}
		}
	}()

	panic1()
	panic2()
	panic("outro")
}
