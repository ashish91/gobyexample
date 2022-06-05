package main

import (
	"fmt"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			i, has := <-jobs
			if has {
				fmt.Println("Received job", i)
			} else {
				fmt.Println("Received all jobs")
				done <- true
				return
			}
		}
	}()


	for i := 0; i < 3; i++ {
		jobs <- i
		fmt.Println("Sent job", i)
	}
  close(jobs)
  fmt.Println("sent all jobs")

  <-done
}
