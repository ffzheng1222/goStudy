package main

import "fmt"

var PACE int = 10

func main() {
	p := new(int)

	for i := 0; i < 10; i++ {
		*p = i
		DoOperation(p, PACE, increase)
		fmt.Printf("main increase：%d\n", *p)
		DoOperation(p, PACE, decrease)
		fmt.Printf("main decrease：%d\n", *p)
	}
}

func DoOperation(x *int, y int, f func(a *int, b int)) {
	f(x, y)
}

func increase(a *int, b int) {
	*a += b
}

func decrease(a *int, b int) {
	*a -= b
}
