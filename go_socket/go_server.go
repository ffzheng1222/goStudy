package main

import (
    "fmt"
    "net"
)


func process(conn net.Conn)  {
    defer conn.Close()

    //这里我们循环的接收客户端发送的数据
    for {
        buf := make([]byte, 1024)
        n, err := conn.Read(buf)
        if err != nil {
            fmt.Printf("客户端退出 err=%v", err)
            return
        }
        //3. 显示客户端发送的内容到服务器的终端
        fmt.Print(string(buf[:n]))
    }
}


func main()  {
    fmt.Println("服务器开始监听....")

    listen, err := net.Listen("tcp", "127.0.0.1:8999")
    if err != nil {
        fmt.Println("listen err=", err)
        return
    }
    defer listen.Close()

    for {
        fmt.Println("等待客户端来链接....")
        connect, err := listen.Accept()
        if err != nil {
            fmt.Println("Accept() err=", err)
        } else {
            fmt.Printf("Accept() suc connect=%v 客户端ip=%v\n", connect, connect.RemoteAddr().String())
        }

        //这里准备其一个协程，为客户端服务
        go process(connect)
    }

}
