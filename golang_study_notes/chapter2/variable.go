package main

import "fmt"

var (
	x    int
	a, s = 100, "33"
)

func test() {
	x, y := 1, 2
	x, y = y+3, x+2
	fmt.Println(x, y)
}

func main() {
	test()
	x = 5
	fmt.Printf("addr%p\n", &x)
	x := 8
	fmt.Printf("addr%p\n", &x)
	x, y := 9, 8
	fmt.Printf("addr%p\n", &x)
	fmt.Printf("xxxxx,")
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(a)
	fmt.Println(s)
}
