package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		channel1 <- "One"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		channel2 <- "Two"
	}()

	for i := 0; i < 2; i++ {
		select {
			case message1 := <-channel1:
				fmt.Println("Received on Channel1", message1)
			case message2 := <-channel2:
				fmt.Println("Received on Channel2", message2)
		}
	}
}
