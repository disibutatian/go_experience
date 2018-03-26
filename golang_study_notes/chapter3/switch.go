package main

import "fmt"

func test_switch() {
	switch score := 90; true {
	case score < 60:
		fmt.Println("不及格\n")
	case score < 80 && score >= 60:
		fmt.Println("及格\n")
	case score <= 100 && score >= 80:
		fmt.Println("优秀\n")
		fallthrough
	default:
		fmt.Println("error\n")
	}
}

func main() {
	test_switch()
}
