package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "strings"
)

func main()  {
    conn, err := net.Dial("tcp", "127.0.0.1:8999")
    if err != nil {
        fmt.Println("client err=", err)
        return
    }

    reader := bufio.NewReader(os.Stdin)

    for {
        //从终端读取一行用户输入，并准备发送给服务器
        line, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("readString err=", err)
        }

        //如果用户输入的是 exit就退出
        line = strings.Trim(line, "\r\n")
        if line == "exit" {
            fmt.Println("客户端退出..")
            break
        }

        //再将line 发送给 服务器
        _, err = conn.Write([]byte(line + "\n"))
        if err != nil {
            fmt.Println("conn.Write err=", err)
        }
    }
}
