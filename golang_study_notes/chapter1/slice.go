package main

import (
	"fmt"
)

func main() {
	//append_slice()
	//copy_slice()
	//add_slice()
	cap_slice()
}

func cap_slice() {
	s := make([]string, 0, 8)
	fmt.Printf("len%v, cap%v\n", len(s), cap(s))
	s = s[:(cap(s))]
	fmt.Printf("len%v, cap%v\n", len(s), cap(s))
}
func add_slice() {
	osa := make([]string, 0)
	sa := &osa

	for i := 0; i < 10; i++ {
		*sa = append(*sa, fmt.Sprintf("s%v", i))
		fmt.Printf("addr of osa:%p , \taddr:%p \t content:%v\n", osa, sa, sa)
	}
	fmt.Printf("addr of osa:%p,\taddr:%p \t content:%v\n", osa, sa, sa)

}

func copy_slice() {
	s := make([]string, 0)
	for i := 0; i < 10; i++ {
		s = append(s, fmt.Sprintf("s%v", i))
	}

	fmt.Println(s)
	var cc = 0
	da := make([]string, 0, 10)
	//cc := copy(da, s)
	copy(da, s)
	fmt.Printf("copy to da(len=%d)\t%v\n", len(da), da)
	da = make([]string, 7)
	cc = copy(da, s[:])
	fmt.Printf("copy to da(len=%d)\tcopied=%d\t%v\n", len(da), cc, da)
	//da = make([]string, 10)
	cc = copy(da[5:], s[8:])
	fmt.Printf("copy to da(len=%d)\tcopied=%d\t%v\n", len(da), cc, da)

}

func append_slice() {
	var s []string
	//fmt.Println(s)

	//s := make([]string)
	fmt.Printf("length%v \taddr%p \tisnil:%v\n", len(s), s, s == nil)
	print("func print", s)

	for i := 1; i < 10; i++ {
		s = append(s, fmt.Sprintf("s%d", i))
	}

	fmt.Printf("[ local print ]\t:\tlength:%v\taddr:%p\tisnil:%v\n", len(s), s, s == nil)
	print("after append", s)

	index := 5
	rear := append([]string{}, s[index:]...)
	s = append(s[0:index], "inserted")
	s = append(s, rear...)
	print("after insert", s)
}

func print(msg string, ss []string) {
	fmt.Printf("[%20s]\t:\tlength:%v\taddr:%p\tisnil:%v\tcontent:%v", msg, len(ss), ss, ss == nil, ss)
	fmt.Println()
}
