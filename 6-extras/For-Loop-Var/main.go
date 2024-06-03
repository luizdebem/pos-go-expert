package main

func main() {
	for i := range 10 {
		println(i)
	}

	done := make(chan bool)
	values := []string{"a", "b", "c"}

	for _, v := range values {
		go func() {
			println(v)
			done <- true
		}()
	}

	for range values {
		<-done
	}
}
