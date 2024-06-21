package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)

	for value := range ch {
		fmt.Println(value)
	}
}
