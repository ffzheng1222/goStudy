package main

import (
    "fmt"
    "net/http"
)

func main(){
    //监听协议
    http.HandleFunc("/", HelloWorldHandler)
    http.HandleFunc("/user/login", UserLoginHandler)

    //监听服务
    err := http.ListenAndServe("127.0.0.1:8089", nil)
    if err != nil {
        fmt.Println("Serve 服务器错误！")
    }
}


func HelloWorldHandler(w http.ResponseWriter, r *http.Request)  {
    fmt.Println("r.Method = ", r.Method)
    fmt.Println("r.URL = "   , r.URL)
    fmt.Println("r.Header = ", r.Header)
    fmt.Println("r.Body = "  , r.Body)
    fmt.Fprintf(w, "Hello World!\n")
}


func UserLoginHandler(response http.ResponseWriter, request *http.Request)  {
    fmt.Println("request.Method = ", request.Method)
    fmt.Println("request.URL = "   , request.URL)
    fmt.Println("request.Header = ", request.Header)
    fmt.Println("request.Body = "  , request.Body)
    fmt.Println("UserLoginHandler Hello!")
    fmt.Fprintf(response, "Login Success!\n")
}
