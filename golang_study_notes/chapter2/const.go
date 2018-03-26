package main

import "fmt"

const (
	x = 100
	y = "100"
)

type color byte

const (
	black color = iota //自定义类型
	red
	blue
)

func test(c color) {
	fmt.Println(c)
}

func main() {
	test(red)
	test(100)

	xx := 2
	yy := color(xx) //可以使用color 作为类型变量
	test(yy)
}
