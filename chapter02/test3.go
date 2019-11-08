package main

import "fmt"

func test(arr []int) {
	fmt.Println(arr)
}

func main() {
	var arr []int
	arr = []int{1, 2, 3, 4, 5}
	test(arr)
}
