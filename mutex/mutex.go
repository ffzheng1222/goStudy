package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// go lock()
	go rlock()
	go wlock()
	time.Sleep(time.Second)
}

func lock() {
	l := sync.Mutex{}
	for i := 0; i < 3; i++ {
		l.Lock()
		defer l.Unlock()
		fmt.Printf("lock: i = %d\n", i)
	}
	fmt.Println()
}

func rlock() {
	l := sync.RWMutex{}
	for i := 0; i < 3; i++ {
		l.RLock()
		defer l.RUnlock()
		fmt.Printf("rlock: i = %d\n", i)
	}
	fmt.Println()
}

func wlock() {
	for i := 0; i < 3; i++ {
		go func(i int) {
			gl := sync.RWMutex{}
			gl.Lock()
			defer gl.Unlock()
			fmt.Printf("wlock: i = %d\n", i)
		}(i)
	}

	//l := sync.RWMutex{}
	//for i := 0; i < 3; i++ {
	//    l.Lock()
	//    //defer l.Unlock()
	//    fmt.Printf("wlock: i = %d\n", i)
	//    l.Unlock()
	//}
	fmt.Println()
}
