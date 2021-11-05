package main

import "fmt"

func main() {
	var mouth byte

	fmt.Print("请输入月份：")
	fmt.Scanln(&mouth)

	switch mouth {
	case 3, 4, 5:
		fmt.Println("春天")
	case 6, 7, 8:
		fmt.Println("夏天")
	case 9, 10, 11:
		fmt.Println("秋天")
	case 12, 1, 2:
		fmt.Println("冬天")
	default:
		fmt.Println("输入有误...")
	}
}
