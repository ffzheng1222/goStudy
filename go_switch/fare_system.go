package main

import "fmt"

func main () {
    var mount, age byte
    var price float64 = 60

    fmt.Print("请输入游玩的月份：")
    fmt.Scanln(&mount)
    fmt.Print("请输入游客的年龄：")
    fmt.Scanln(&age)

    if mount >= 4 && mount <= 10 {
        if age >= 18 && age <= 60 {
            fmt.Printf("%v月 年龄:%v  票价:%v\n", mount, age, price)
        } else if age < 18 {
            fmt.Printf("%v月 年龄:%v  票价:%v\n", mount, age, price/2)
        } else {
            fmt.Printf("%v月 年龄:%v  票价:%v\n", mount, age, price/3)
        }
    } else {
        if age >= 18 && age <= 60 {
            fmt.Printf("%v月 年龄:%v  票价:40\n", mount, age)
        } else {
            fmt.Printf("%v月 年龄:%v  票价:20\n", mount, age)
        }
    }

}
