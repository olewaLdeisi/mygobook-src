package main

import "fmt"

func main() {
	var j int = 5

	a := func() func() {
		var i int = 10
		return func() {
			fmt.Printf("i, j: %d, %d\n", i, j)
		}
	}() // 括号删除则执行无输出

	a()
	j *= 2
	a()
}
