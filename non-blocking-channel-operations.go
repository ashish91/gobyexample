package main

import (
	"fmt"
)

func main() {
	messages_channel := make(chan string)
	signal_channel := make(chan bool)

	select {
	case msg := <-messages_channel:
		fmt.Println("Received message", msg)
	default:
		fmt.Println("No message received")
	}

	msg := "Hi"
	select {
	case messages_channel <- msg:
		fmt.Println("Sent message", msg)
	default:
		fmt.Println("No message sent")
	}

	select {
		case msg := <-messages_channel:
      fmt.Println("received message", msg)
		case sig := <-signal_channel:
      fmt.Println("received signal", sig)
		default:
			fmt.Println("No activity")
	}
}
