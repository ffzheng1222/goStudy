package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var heroMap map[string]interface{}

	heroMap = make(map[string]interface{})
	heroMap["name"] = "张无忌"
	heroMap["age"] = 22
	heroMap["address"] = "冰火岛"

	data, err := json.Marshal(&heroMap)
	if err != nil {
		fmt.Printf("序列号错误 err=%v\n", err)
	}

	fmt.Printf("序列化后=%v\n", string(data))
}
