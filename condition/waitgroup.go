package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	waitBySleep()
	waitByChannel()
	waitByWG()
}

func waitBySleep() {
	for i := 0; i < 10; i++ {
		go fmt.Printf("waitBySleep i = %d\n", i)
	}
	time.Sleep(time.Second)
	fmt.Println()
}

func waitByChannel() {
	ch := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Printf("waitByChannel i = %d\n", i)
			ch <- true
		}(i)
	}
	for i := 0; i < 10; i++ {
		<-ch
	}
	fmt.Println()
}

func waitByWG() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Printf("waitByWG i = %d\n", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
