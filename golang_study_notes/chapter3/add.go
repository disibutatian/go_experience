package main

import "fmt"

const (
	read byte = 1 << iota
	write
	exec
	freeze
)

func test() {
	a := read | write | freeze
	b := read | freeze | exec
	c := a &^ b
	fmt.Printf("%04b &^ %04b = %04b\n", a, b, c)
}

func main() {
	//fmt.Println("vim-go")
	test()
}
