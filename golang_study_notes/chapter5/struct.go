package main

import (
	"fmt"
)

type node struct {
	_    int
	id   int
	next *node
}

type file struct {
	name string
	attr struct {
		x int
		y int
	}
}
type int_new int

type anonymous struct {
	int_new
}

func struct_test() {
	n1 := node{
		id: 1,
	}
	n2 := node{
		id:   2,
		next: &n1,
	}

	fmt.Println(n1, n2)
	// 直接匿名定义结构体
	n3 := struct {
		x int
		y int
	}{
		x: 5,
		y: 5,
	}

	//作为字段类型
	f := file{
		name: "hh",
		//attr: {
		//	x: 1,
		//	y: 1,
		//},
	}

	f.attr.x = 8
	f.attr.y = 8
	fmt.Println(n3, f)

	//
	an := anonymous{int_new: 7}
	//an.int_new = 6
	fmt.Println(an)

}

func struct_chan() {
	c := make(chan struct{})

	go func() {
		fmt.Println("hello")
		c <- struct{}{}
	}()

	<-c
	fmt.Println("world")

}

func main() {
	struct_test()
	//struct_chan()
}
