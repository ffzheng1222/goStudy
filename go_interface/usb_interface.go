package main

import (
	"fmt"
)

type Usb interface {
	Start()
	Stop()
}

/* phone  实现Usb的Start Stop接口方法 */
type Phone struct {
	name string
}

func (p Phone) Start() {
    fmt.Printf(p.name)
	fmt.Println("手机开始充电...")
}

func (p Phone) Stop() {
    fmt.Printf(p.name)
	fmt.Println("手机停止充电...")
}

func (p Phone) Call() {
    fmt.Printf(p.name)
	fmt.Println("手机正在打电话...")
}

/* camera  实现Usb的Start Stop接口方法 */
type Camera struct {
	name string
}

func (c Camera) Start() {
    fmt.Printf(c.name)
	fmt.Println("相机开始工作...")
}

func (c Camera) Stop() {
    fmt.Printf(c.name)
	fmt.Println("相机停止工作...")
}

/* Computer   */
type Computer struct {
}

func (computer Computer) working(usb Usb) {
	usb.Start()

	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
	usb.Stop()
}

func main() {
    var usbArray [3]Usb
    usbArray[0] = Phone{"苹果"}
    usbArray[1] = Phone{"小米"}
    usbArray[2] = Camera{"索尼"}

    var computer Computer
    for _, v := range usbArray {
        computer.working(v)
        fmt.Println()
    }
}
