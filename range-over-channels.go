package main

import (
	"fmt"
)

func main() {
	queue := make(chan string, 2)
	queue <- "One"
	queue <- "Two"
	close(queue)

	for val := range queue {
		fmt.Println(val)
	}
}
