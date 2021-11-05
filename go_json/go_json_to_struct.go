package main

import (
    "encoding/json"
    "fmt"
)

type Hero struct {
    Name string
    Age int
    Birthday string
    Sal float64
    Skill string
}


func main()  {
    //说明：str在项目的开发中，是通过网络传输获取到，或者读取文件获取到
    jsonStr := "{\"Name\":\"张三丰\", \"Age\":98, \"Birthday\":\"2001-09-21\", \"Sal\":3800.85, \"Skill\":\"武当剑法\"}"


    var heroStruct Hero
    err := json.Unmarshal([]byte(jsonStr), &heroStruct)
    if err != nil {
        fmt.Printf("jsonStr unmarshal struct err=%v\n", err)
    }

    fmt.Printf("反序列化后 heroStruct=%v \nheroStruct.Name=%v\n", heroStruct, heroStruct.Name)
}
