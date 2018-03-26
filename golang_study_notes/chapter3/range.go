package main

import (
	"fmt"
)

func test_range() {
	data := [3]int{10, 20, 30}

	//for i, x := range data {
	//	fmt.Printf("addr: %p \n", &x)
	//	fmt.Printf("addr: %p \n", &data[0])
	//	if i == 0 {
	//		data[0] += 100
	//		data[1] += 200
	//		data[2] += 300
	//	}
	//	fmt.Printf("x: %d \n", data[2])

	//	fmt.Printf("x: %d, data: %d\n", x, data[i])
	//}
	k := data[:]
	for i, x := range k {
		//fmt.Printf("addr: %p \n", &k[0])
		//if i == 0 {
		//	k[0] += 100
		//	k[1] += 200
		//	k[2] += 300
		//}
		//k = append(k, 90)
		//k = append(k, 90)
		//k = append(k, 90)
		//fmt.Printf("addr: %p \n", k)
		//k[2] = 5
		//if i == 6 {
		//	break
		//}

		fmt.Printf("x: %d, data: %d\n", x, k[i])
	}
	//	for i, x := range data[:] {
	//		fmt.Printf("addr: %p \n", &data)
	//		if i == 0 {
	//			data[0] += 100
	//			data[1] += 200
	//			data[2] += 300
	//		}
	//
	//		fmt.Printf("x: %d, data: %d\n", x, data[i])
	//	}
	//s := data[:]
	//fmt.Printf("addr: %p \n", s)
	//fmt.Println(s)
	//s = append(s, 90)
	//fmt.Printf("addr: %p \n", &data)
	//fmt.Printf("addr: %p \n", s)
	//fmt.Println(s)
	//fmt.Println(data)
	//s[0] = 4
	//fmt.Println(s)
	//fmt.Println(data)

}

func main() {
	test_range()
}
