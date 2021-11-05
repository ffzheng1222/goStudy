package main

import (
    "fmt"
)

func main () {
    var i interface{} = "kk"

    j, ok := i.(int)

    if ok {
        fmt.Printf("%T->%d", j, j)
    } else {
        fmt.Println("类型不匹配！")
    }
}
