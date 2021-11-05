package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
    ///ch := make(chan bool)
	baseCtx := context.Background()
	ctx := context.WithValue(baseCtx, "tony", "25")
    //go func(c context.Context) {
    //    fmt.Println(c.Value("tony"))
    //}(ctx)

	timeOutCtx, cancel := context.WithTimeout(ctx, time.Second * 6)
	defer cancel()
	go func(timeCtx context.Context) {
		ticker := time.NewTicker(time.Second * 1)
		for _ = range ticker.C {
			select {
			case <-timeCtx.Done():
				fmt.Println("child go routine interrupt ...")
				return
			default:
				fmt.Println("enter default")
			}
		}
	}(timeOutCtx)

    go func(c context.Context) {
        fmt.Println(c.Value("tony"))
    }(ctx)

    select {
    case <- timeOutCtx.Done():
        fmt.Println("main process exit!")
    }
}
