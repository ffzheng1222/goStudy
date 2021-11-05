package main

import (
    "fmt"
    "math/rand"
    "time"
)

type Orderinfo struct {
	id int
	name string
}

// RandString 生成随机字符串
func RandString(len int) string {
    bytes := make([] byte, len)
    for i:=0; i<len; i++{
        b := rand.Intn(26) + 65
        bytes[i] = byte(b)
    }
    return string(bytes)
}

func product(in chan<- Orderinfo, oi []Orderinfo) {
    for _, orderinfo := range oi {
        orderinfo.id = rand.Intn(100)
        orderinfo.name = RandString(10)
        in <- orderinfo
    }
    close(in)
}


func consume(out <-chan Orderinfo, oi []Orderinfo) {
    for orderinfo := range out {
        fmt.Printf("订单id：%d, 订单名：%s\n", orderinfo.id, orderinfo.name)
    }
}


func main() {
    rand.Seed(time.Now().UnixNano())
    OrderinfoSlice := make([]Orderinfo, 10)
	ch := make(chan Orderinfo)

	go product(ch, OrderinfoSlice)
	consume(ch, OrderinfoSlice)
}


