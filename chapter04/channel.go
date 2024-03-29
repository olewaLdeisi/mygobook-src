package main

import "fmt"

func Count1(ch chan int) {
	fmt.Println("Counting")
	ch <- 1
}

func main() {

	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count1(chs[i])
	}

	for _, ch := range chs {
		<-ch
	}
	/*
		ch := make(chan int, 1)
		for {
			select {
			case ch <- 0:
			case ch <- 1:
			}
			i := <- ch
			fmt.Println("Value received:", i)
		}
	*/
}
