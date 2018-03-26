package main

import (
	"fmt"
)

func test(x int) func() {
	fmt.Println(&x)

	return func() {
		fmt.Println(&x, x)
	}
}

func main() {
	f := test(100) //只是返回了一个匿名函数
	f()
}
