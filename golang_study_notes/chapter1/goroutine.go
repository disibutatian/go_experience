package main

import (
	"fmt"
	"time"
)

func task(id int) {
	for i := 0; i < 4; i++ {
		fmt.Printf("id:%v, i:%v\n", id, i)
		time.Sleep(time.Second * 1)
	}
}

func consumer(data chan int, done chan bool) {
	for x := range data {
		fmt.Println("recv", x)
	}

	done <- true
}

func producer(data chan int) {
	for i := 0; i < 4; i++ {
		data <- i
		time.Sleep(time.Second)
	}

	close(data)
}

func main() {
	//go task(1)
	//go task(2)
	//time.Sleep(time.Second * 4)
	done := make(chan bool)
	data := make(chan int)

	go consumer(data, done)
	go producer(data)

	<-done
}
