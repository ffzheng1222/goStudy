package main

import (
    "encoding/json"
    "fmt"
)

func main() {
	jsonStr := "{\"name\":\"张无忌\", \"age\":18, \"address\":\"九阳神功\"}"

	var heroMap map[string]interface{}

	err := json.Unmarshal([]byte(jsonStr), &heroMap)
	if err != nil {
        fmt.Printf("jsonStr unmarshal map err=%v\n", err)
	}

    fmt.Printf("反序列化后 heroMap=%v\n", heroMap)
}
