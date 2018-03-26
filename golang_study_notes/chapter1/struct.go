package main

import (
	"fmt"
)

type user struct {
	name rune
	age  byte
}

type manger struct {
	user
	title string
}

func main() {
	var m manger

	m.name = 'h'
	m.age = 27
	m.title = "fy"

	fmt.Println(m)
	s := fmt.Sprintf("%+v", m)
	fmt.Println(s)
}
