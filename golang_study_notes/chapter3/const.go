package main

import "fmt"

const s = 3

func test() {
	//a := 2.0 << 3
	//fmt.Printf("%T, %v\n", a, a)

	//var s uint = 3
	//var b int = 1.0 << s
	b := 1.0 << s
	fmt.Printf("%T, %v\n", b, b)
}

func main() {
	fmt.Println("vim-go")
	test()
}
