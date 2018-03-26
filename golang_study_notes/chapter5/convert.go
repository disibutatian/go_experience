package main

import (
	"fmt"
)

func test() {
	r := '我'

	s := string(r)
	b := byte(r)

	//s2 := string(b)
	//r2 := rune(b)
	r2 := rune(r)

	//fmt.Println(s, b, s2, r2)
	fmt.Println(s, b, r2)
}

func main() {
	test()

	m := map[string]int{
		"abc": 123,
	}

	key := []byte("abc")
	x, ok := m[string(key)]

	println(x, ok)
}
