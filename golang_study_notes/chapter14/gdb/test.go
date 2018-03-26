package main

import (
	"fmt"
	"time"
)

func count(c chan int) {
	for i := 0; i < 2; i++ {
		time.Sleep(time.Second * 1)
		c <- i
	}

	close(c)
}

func main() {
	msg := "start "
	fmt.Println(msg)

	c := make(chan int)
	go count(c)

	for value := range c {
		fmt.Println("count : ", value)
	}
}
