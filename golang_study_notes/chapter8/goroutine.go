package goroutine

import (
	"fmt"
	//"math"
	"sync"
	"time"
)

var c int

func count() int {
	c += 1
	return c
}

func goroutine_copy() {
	a := 10

	go func(a, c int) {
		time.Sleep(time.Second)
		fmt.Printf("go a : %v, c : %v\n", a, c)
	}(a, count())

	a += 10
	fmt.Printf("main a : %v, c : %v\n", a, count())
	time.Sleep(time.Second * 3)

}

func goroutine_wait() {
	exit := make(chan struct{})

	go func() {
		fmt.Println("go wait")
		//close(exit)
		exit <- struct{}{}
	}()

	<-exit
	fmt.Println("main exit")
}

func count_int() {
	var x int
	//for i := 0; i < math.MaxInt32; i++ {
	for i := 0; i < 3; i++ {
		x += i
	}

	fmt.Printf("count_int : %v\n", x)

}

func test_count(n int) {
	for i := 0; i < n; i++ {
		count_int()
	}
}

func test_concurrency_count(n int) {
	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		//执行匿名函数，并发执行
		go func() {
			count_int()
			wg.Done()
		}()
		fmt.Println("sort : ", i)
	}
	wg.Wait()
}
