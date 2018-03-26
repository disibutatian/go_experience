package pkg

//package main

import (
	"fmt"
	"study/chapter9/internal"
	//内部包的使用必须使用绝对路径
	//"./internal"
	"study/chapter9/test_dir"
	//"./test_dir"
	"unsafe" //用于指针类型转换
)

func test_private() {
	//当出现引用包出现失败时，有可能是路径错误，还有就是定义的包编译不能通过
	d := test.CreateData()
	fmt.Println(d)
	d.Y = 100

	//k := (*struct{ x int })(unsafe.Pointer(d))
	k := (*struct {
		x int
		y int
	})(unsafe.Pointer(d))

	k.x = 200
	k.y = 200
	fmt.Println(d)
}

func test_internal() {
	fmt.Println(inner.I)
}

//func main() {
//	test_private()
//	test_internal()
//}
