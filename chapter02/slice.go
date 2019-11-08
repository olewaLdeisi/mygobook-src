package main

import "fmt"

func main() {
	// 定义数组, 数组是值类型， 赋值和参数传递都会数据复制
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 定义切片
	var mySlice []int = myArray[:5]

	fmt.Println("Elements of myArray: ")
	for _, v := range myArray {
		fmt.Print(v, " ")
	}

	fmt.Println("\nElements of mySlice: ")

	for _, v := range mySlice {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
