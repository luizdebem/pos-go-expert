package main

import (
	"fmt"
	"time"
)

type Message struct {
	id  int
	Msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)

	go func() {
		msg := Message{1, "hello from rabbitmq"}
		c1 <- msg
		time.Sleep(time.Second * 2)
	}()

	go func() {
		msg := Message{1, "hello from kafka"}
		c2 <- msg
		time.Sleep(time.Second * 1)
	}()

	for {
		select {
		case msg := <-c1:
			fmt.Printf("Received from RabbitMQ: %v\n", msg.Msg)

		case msg := <-c2:
			fmt.Printf("Received from Kafka: %v\n", msg.Msg)

		case <-time.After(time.Second * 3):
			println("timeout")

			// default:
			// 	println("default")
		}
	}

}
