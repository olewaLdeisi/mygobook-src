package main

import "fmt"

func main() {
	// 直接创建数组切片，实际会创建一个匿名数组
	mySlice := make([]int, 5, 10)

	fmt.Println("len(mySlice)", len(mySlice))
	fmt.Println("cap(mySlice)", cap(mySlice))

	// 新增元素的方式
	mySlice = append(mySlice, 1, 2, 3)
	fmt.Println("len(mySlice)", len(mySlice))
	fmt.Println("cap(mySlice)", cap(mySlice))
	mySlice2 := []int{8, 9, 10}
	mySlice = append(mySlice, mySlice2...)
	fmt.Println("len(mySlice)", len(mySlice))
	fmt.Println("cap(mySlice)", cap(mySlice))

	// 基于数组切片创建数组切片
	oldSlice := []int{1, 2, 3, 4, 5}
	newSlice := oldSlice[:3]
	fmt.Println("len(newSlice)", len(newSlice))
	fmt.Println("cap(newSlice)", cap(newSlice))

	//newSlice = oldSlice[:5]

	// 内容复制
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	for i := 0; i < len(slice2); i++ {

	}
	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1中的前3个位置
}
