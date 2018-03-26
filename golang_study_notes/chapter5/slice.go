package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func slice_test() {
	var a []int
	b := []int{2, 3}

	fmt.Println(b[0])

	fmt.Printf("a: %#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&a)))
	fmt.Printf("b: %#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&b)))
}

func main() {
	slice_test()
}
