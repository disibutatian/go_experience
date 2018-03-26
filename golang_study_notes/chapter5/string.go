package main

import "fmt"

func test() {
	s := ` ni hao \r\n
		fdjs`
	fmt.Println(s)

	s = "ab" +
		"cd"
	fmt.Println(s)

}

func main() {
	test()

	//fmt.Println("vim-go")
}
