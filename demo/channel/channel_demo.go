// channel_demo
package main

import (
	"fmt"
	"time"
)

//只能向chan里写数据
func send(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
}

//只能取channel中的数据
func recv(c <-chan int) {
	for i := range c {
		fmt.Println(i)
	}
}

func main() {
	ch := make(chan int)
	go send(ch)
	go recv(ch)
	time.Sleep(3 * time.Second)

	close(ch)
}
