package goroutine

import (
	"fmt"
)

func test_type() {
	var v1 int = 5
	fmt.Printf("%T->%v\n", v1, v1)
	v2 := float32(v1)
	fmt.Printf("%T->%v\n", v2, v2)
	v3 := float64(v1)
	fmt.Printf("%T->%v\n", v3, v3)

	var1 := new(int32)
	fmt.Printf("%T->%v\n", var1, var1)
	var2 := (*int32)(var1)
	fmt.Printf("%T->%v\n", var2, var2)
}

func test_assert() {
	var i interface{} = 888
	j, ok := i.(int)
	if ok {
		fmt.Printf("assert %v to %T ---%v", i, i, j)
	} else {
		fmt.Println("assert 失败")
	}

}
