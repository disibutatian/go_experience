package goroutine

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func chan_varible() {

	done := make(chan struct{})
	c := make(chan string)

	go func() {
		defer close(done)

		// range 可以保证close()前数据都处理完
		for value := range c {
			fmt.Println(value)
		}
	}()

	c <- "v1"
	c <- "v2"

	close(c)

	// 读取阻塞,需要等待有数据输入或者close(done) 发出信号关闭
	<-done
}

func chan_direction() {
	var wg sync.WaitGroup
	wg.Add(2)

	c := make(chan int)
	var send chan<- int = c
	var recv <-chan int = c

	go func() {
		defer wg.Done()
		for value := range recv {
			fmt.Println(value)
		}

		// 也可以使用原先的通道进行数据的读取
		//for value := range c {
		//	fmt.Println(value)
		//}
	}()

	go func() {
		defer wg.Done()
		// 发送端控制 chan 关闭
		//defer close(c)
		// 和close(c) 是等价的
		defer close(send)
		for i := 2; i < 7; i++ {
			send <- i // 向发送管道输入数据
		}
	}()

	wg.Wait()
}

func chan_sem() {
	var wg sync.WaitGroup
	sem := make(chan struct{}, 2)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			//用来释放资源
			defer func() {
				<-sem
			}()

			sem <- struct{}{}
			fmt.Println("i : ", i)
			fmt.Println(time.Now())
			time.Sleep(time.Second)

		}(i)
	}

	wg.Wait()
}

func chan_time() {
	//实现超时设置
	go func() {
		for {
			select {
			case <-time.After(time.Second * 5):
				fmt.Println("timeout")
				os.Exit(0)
			}
		}
	}()

	go func() {
		tick := time.Tick(time.Second)
		for {
			select {
			case <-tick:
				fmt.Println(time.Now())
			}
		}
	}()

	<-(chan struct{})(nil)

}

const (
	max     = 50000000
	block   = 500
	bufsize = 100
)

func chan_performance_1() {
	done := make(chan struct{})
	c := make(chan int, bufsize)
	go func() {
		defer close(done)
		count := 0
		for value := range c {
			count += value
		}
	}()
	for i := 0; i < max; i++ {
		c <- i
	}
	close(c)
	<-done
}

func chan_performance_block() {
	done := make(chan struct{})
	c := make(chan [block]int, bufsize)
	go func() {
		defer close(done)
		count := 0
		for ch := range c {
			for _, value := range ch {
				count += value
			}
		}
	}()

	for i := 0; i < max; i += block {
		var tmp_arr [block]int
		for j := 0; j < block && i+j < max; j++ {
			tmp_arr[j] = i + j
		}
		c <- tmp_arr
	}

	close(c)
	<-done
}

type data struct {
	sync.Mutex //匿名互斥锁
}

func (d *data) test_mutex(s string) {
	d.Lock()
	defer d.Unlock()

	for i := 0; i < 5; i++ {
		fmt.Printf("%v %v\n", s, i)
		//time.Sleep(time.Second)
	}
}

func chan_mutex() {
	var wg sync.WaitGroup
	wg.Add(2)
	d := &data{}

	go func() {
		defer wg.Done()
		d.test_mutex("read")
	}()

	go func() {
		defer wg.Done()
		d.test_mutex("write")
	}()

	//等待所有任务执行结束
	wg.Wait()

}
