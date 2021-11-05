package main

import (
	"encoding/json"
	"fmt"
)

type HeroJson struct {
	Name     string
	Age      int
	Birthday string
	Sal      float64
	Skill    string
}

func main() {
	hero := HeroJson{
		Name:     "张三丰",
		Age:      88,
		Birthday: "2009-11-11",
		Sal:      8000.0,
		Skill:    "教武当剑法!",
	}

	data, err := json.Marshal(&hero)
	if err != nil {
		fmt.Printf("序列号错误 err=%v\n", err)
	}

	fmt.Printf("序列化后=%v\n", string(data))
}
