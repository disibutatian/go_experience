package goroutine

import (
	"fmt"
	"sync"
	"time"
)

func select_timeout() {

	timeout := make(chan struct{})
	c := make(chan int, 1)

	go func() {
		time.Sleep(time.Second)
		fmt.Println("wait over")
		close(timeout)
	}()

	//c <- 5

	select {
	case value := <-c:
		fmt.Println("value: ", value)
	case <-timeout:
		fmt.Println("timeout")
	}

}

func select_multi() {
	var wg sync.WaitGroup
	wg.Add(2)
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		defer wg.Done()

		var (
			name  string
			ok    bool
			value int
		)

		for {
			select {
			//ok 是bool 型变量
			case value, ok = <-c1:
				name = "c1"
			case value, ok = <-c2:
				name = "c2"
			}
			//ok
			if !ok {
				//在这里作判断可以避免关闭通道还0值被读出来
				break
			}

			fmt.Println(name, value)
		}
	}()

	go func() {
		defer wg.Done()
		defer close(c1)
		defer close(c2)

		for i := 0; i < 10; i++ {
			select {
			case c1 <- i:
			case c2 <- i:
			}
		}
	}()
	wg.Wait()
}

func select_produce() {
	done := make(chan struct{})
	data := []chan int{
		make(chan int, 3), //初始化一个数据
	}

	go func() {
		defer close(done)
		for i := 0; i < 11; i++ {
			select {
			case data[len(data)-1] <- i:
			default:
				data = append(data, make(chan int, 3))
			}
		}
	}()

	<-done
	//read the slice data
	for i := 0; i < len(data); i++ {
		close(data[i])
		for value := range data[i] { // 可以在close() 同时及时关闭数据输出
			fmt.Println("value :", value)
		}
	}

}

type receiver struct {
	sync.WaitGroup
	data chan int
}

func newReceiver() *receiver {
	//init receiver
	r := &receiver{
		data: make(chan int),
	}

	r.Add(1)
	go func() {
		defer r.Done() //依赖r.data 的关闭
		for value := range r.data {
			fmt.Println("data :", value)
		}
	}()

	return r

}

func create_use_receiver() {
	r := newReceiver()
	r.data <- 1
	r.data <- 2

	close(r.data)
	r.Wait()
}
