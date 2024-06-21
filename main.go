package main

import (
	"fmt"
	"time"
)

func worker(ch chan int) {
	time.Sleep(time.Microsecond)
	ch <- 42
}

func main() {
	ch := make(chan int)
	go worker(ch)
	value := <-ch
	fmt.Println("recieved", value)
}
