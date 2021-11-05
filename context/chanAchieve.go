package main

import (
	"fmt"
	"time"
)

func main() {
	message := make(chan int, 5)
	done := make(chan bool)

	defer close(message)
	// consume
	go func() {
		ticker := time.NewTicker(time.Second * 1)
		for _ = range ticker.C {
			select {
			case <-done:
				fmt.Println("child go routine interrupt ...")
				return
			default:
				fmt.Printf("send message: %d\n", <-message)
			}
		}
	}()

	// product
	for i := 0; i < 5; i++ {
		message <- i
	}

	time.Sleep(time.Second * 5)
	close(done)
	time.Sleep(time.Second * 1)
	fmt.Println("main process exit!")
}
