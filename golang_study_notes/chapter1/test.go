package main

import (
//"errors"
//"fmt"
)

func test(a, b int) {
	defer println("dispose...")

	println(a / b)
}

func main() {
	a, b := 100, 2
	test(a, b)
}
