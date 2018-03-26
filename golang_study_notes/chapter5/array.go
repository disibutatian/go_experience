package main

import (
	"fmt"
)

//数组指针
func array_point_test(a *[2]int) {
	fmt.Printf("addr %p val %v\n", a, a)

}

//
func array_slice(a []int) {

	fmt.Printf("addr %p val %v\n", a, a)
}

func main() {
	a1 := [2]int{10, 20}
	fmt.Printf("addr %p val %v\n", &a1, a1)

	array_point_test(&a1)
	array_slice(a1[:])
}
