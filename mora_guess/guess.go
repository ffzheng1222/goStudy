package main

import (
    "fmt"
    "math/rand"
)

func existIn(str1 [][]string, str2 []string) int {
    for _, item := range str1 {
        if item[0] == str2[0] && item[1] == str2[1] {
            return 1
        }
    }
    return 0
}

func main() {
    var person string
    guessList := []string{"石头", "剪刀", "布"}
    win := [][]string{{"布", "石头"}, {"石头", "剪刀"}, {"剪刀", "布"}}

    for {
        /********************************************************************
        rand.Intn () 函数是个伪随机函数，不管运行多少次都只会返回同样的随机数，因为它默认的资源就是单一值，所以必须调用 rand.Seed (),
        并且传入一个变化的值作为参数，如 time.Now().UnixNano() , 就是可以生成时刻变化的值.
        **********************************************************************/
        fmt.Println("请输入 '石头, 剪刀, 布'")
        _, _ = fmt.Scanf("%s\n", &person)

        num := rand.Intn(len(guessList))
        computer := guessList[num]
        fmt.Println(computer)

        // 构造一个游戏双方呈现结果的slice
        input := []string{computer, person}
        if computer == person {
            fmt.Println("平手！")
            fmt.Println()
        } else if existIn(win, input) > 0 {
            fmt.Println("电脑获胜！")
            fmt.Println()
        } else {
            fmt.Println("人获胜！")
            fmt.Println()
            break
        }
    }
}
