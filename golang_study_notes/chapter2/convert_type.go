package main

import (
	"fmt"
)

type st struct {
	x int
	y string
}

func main() {
	type data int
	var d data = 10

	var x int = 5
	println(d == data(x))

	s := st{32, "fsd"}
	var k st = s
	var p st = st{423, "fd"}
	fmt.Println(s)
	fmt.Println(k)
	fmt.Println(p)
}
