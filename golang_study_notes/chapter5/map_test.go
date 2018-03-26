package main

import (
	"fmt"
	"sync"
	"testing"
	//"time"
)

type user struct {
	name string
	age  byte
}

func map_test() {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 1

	if v, ok := m["b"]; ok {
		fmt.Println(v)
	}

	fmt.Println(m)

	//初始化表达语句
	m2 := map[int]struct {
		x int
	}{
		1: {100},
	}

	fmt.Println(m2)

}

func map_modify_value() {
	m := map[int]user{
		1: {"tom", 20},
		2: {"hh", 27},
	}
	v := m[1]
	v.age = v.age + 10
	fmt.Println(m)
	m[1] = v
	fmt.Println(m)

	//指针方式修改值
	m2 := map[int]*user{
		1: &user{"tom", 20},
		2: &user{"hh", 27},
	}
	fmt.Println(m2[1])
	m2[1].age += 15
	fmt.Println(m2[1])

}

func map_add() {
	m := make(map[int]int)

	for i := 0; i < 10; i++ {
		m[i] = i + 10
	}

	for k := range m {
		fmt.Println(m[k])
		if k == 1 {
			m[100] = 1000
		}
	}
	fmt.Println(m)

}

func map_concurrent() {
	m := make(map[string]int)
	var lock sync.RWMutex //使用读写锁进行并发控制

	go func() {
		for {
			lock.Lock()
			m["a"] = 1
			lock.Unlock()
			//time.Sleep(time.Millisecond)
		}
	}()

	go func() {
		for {
			lock.RLock()
			_ = m["a"]
			lock.RUnlock()
			//time.Sleep(time.Millisecond)
		}
	}()
}

func test() map[int]int {
	m := make(map[int]int)
	for i := 0; i < 1000; i++ {
		m[i] = i
	}

	return m
}

func testInit() map[int]int {
	m := make(map[int]int, 1000)
	for i := 0; i < 1000; i++ {
		m[i] = i
	}

	return m
}

func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test()
	}
}

func BenchmarkTestInit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testInit()
	}
}

func main() {
	map_test()
	//map_modify_value()
	//map_add()
	//map_concurrent()
	//select {}
}
