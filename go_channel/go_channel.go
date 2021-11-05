package  main

import (
    "fmt"
    "time"
)

func main(){
    ch := make(chan int, 3)
    fmt.Println("len = ", len(ch),  "cap = ", cap(ch))

    go func() {
        for i := 0; i < 8; i++ {
            ch <- i
            fmt.Println("子go程: i = ", i)
            fmt.Println("len = ", len(ch),  "cap = ", cap(ch))
        }
    }()

    time.Sleep(time.Second * 3)

    for i := 0; i < 8; i++ {
        num := <- ch
        fmt.Println("主go程读到: num = ", num)
    }

}
