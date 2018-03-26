package main

import "fmt"

func test() *int {
	a := 100
	return &a
}

func main() {
	var a *int = test()
	fmt.Println(a, *a)
}
