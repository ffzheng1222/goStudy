package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var heroSlice []map[string]interface{}

	h1 := make(map[string]interface{})
	h1["name"] = "张无忌"
	h1["age"] = 25
	h1["address"] = "冰火岛"
	heroSlice = append(heroSlice, h1)

	h2 := make(map[string]interface{})
	h2["name"] = "张三丰"
	h2["age"] = 89
	h2["address"] = [2]string{"武当山", "夏威夷"}
	heroSlice = append(heroSlice, h2)

	data, err := json.Marshal(&heroSlice)
	if err != nil {
		fmt.Printf("序列号错误 err=%v\n", err)
	}

	fmt.Printf("序列化后=%v\n", string(data))
}
