package main

import (
    "encoding/json"
    "fmt"
)

func main() {
	jsonStr := "[{\"address\":\"北京\", \"age\":8, \"name\":\"tom\"}," +
		"{\"address\":[\"河南\",\"上海\"], \"age\":18, \"name\":\"mary\"}]"

	var addrSlice []map[string]interface{}
	//注意：反序列化map,不需要make，因为make操作被封装到Unmarshal函数
	err := json.Unmarshal([]byte(jsonStr), &addrSlice)
	if err != nil {
        fmt.Printf("jsonStr unmarshal slice err=%v\n", err)
	}

    fmt.Printf("反序列化后 addrSlice=%v\n", addrSlice)
}
