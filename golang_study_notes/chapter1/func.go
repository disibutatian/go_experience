package main

import (
	"fmt"
)

type user struct {
	name string
	age  byte
}

func (u user) ToString() string {
	//fmt.Println(u)
	return fmt.Sprintf("%+v", u)
}

type manger struct {
	user
	title string
}

func main() {
	var m manger
	m.name = "hh"
	m.age = 27

	fmt.Println(m.ToString())
}
