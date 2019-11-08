package main

import "fmt"

func GetName() (firstName, lastName, nickName string) {
	return "May", "Chan", "Chibi Maruko"
}

func myfunc() {
	i := 0
HERE:
	fmt.Println(i)
	i++
	if i < 10 {
		goto HERE
	}
}

func myfunc2(args ...int) {
	// 语法糖，实际上args是个切片
	fmt.Println(args)
	for _, arg := range args {
		fmt.Print(arg, " ")
	}
	fmt.Println()
}

func myfunc3(args ...int) {
	myfunc2(args...)
	myfunc2(args[1:]...)
}

func main() {
	// 变量声明
	/*
		var v1 int
		var v2 string
		var v3 [10]int	// 数组
		var v4 []int	// 数组切片
		var v5 struct{	// 结构体
			f int
		}
		var v6 *int		// 指针
		var v7 map[string]int	// map，key为string类型，value为int类型
		var v8 func(a int) int
		var (
			v9 int
			v10 string
		)
	*/

	// 变量初始化
	/*
		var v1 int =10
		var v2 = 10	// 编译器自动推导v2类型
		v3 := 10	// 编译器自动推导v3类型，不能用于声明全局变量，不能是已经被声明过的
	*/

	// 变量赋值
	/*
		var v10 int
		v10 = 123
		i, j = j, i
	*/

	// 多重返回和匿名变量
	/*
		_, _, nickName := GetName()
	*/

	// 常量
	/*
		const Pi float64 = 3.14159265358979323846
		const zero = 0.0	// 无类型浮点常量
		const (
			size int64 = 1024
			eof = -1	// 无类型整型常量
		)
		const u, v float32 = 0, 3
		const a, b, c = 3, 4, "foo"
		// a = 3, b = 4, c = "foo" ， 无类型整型和字符串常量
		const mask = 1 << 3	// 可以是编译器运算的常量表达式
		// const Home = os.GetEnv("HOME") 这种运行期才能返回结果，会编译报错
	*/

	// 预定义常量
	/*
		const (	// iota遇到一个const会被重置为0
			c0 = iota	// c0 == 0
			c1 = iota	// c1 == 1
			c2 = iota	// c2 == 2
		)
		const (
			c3 = iota	// c3 == 0
			c4 	// c4 == 1
			c5 	// c5 == 2
		)
		const (
			a = 1 << iota	// a == 1
			b = 1 << iota	// b == 2
			c = 1 << iota	// c == 4
		)
		const (
			d = 1 << iota	// d == 1
			e  	// e == 2
			f 	// f == 4
		)
	*/

	// 字符串
	/*
		str := "Hello, 世界"
		n := len(str)
		for i := 0; i < n; i++ {
			ch := str[i]
			fmt.Println(i, ch)
		}
		for key, value := range str {
			fmt.Println(key, value)
		}
	*/

	// 流程控制-if
	/*
		// 必须有花括号
		if a < 5 {
		} else {
		}
	*/

	// 流程控制-switch
	/*
		i := 2
		switch i {
		case 0:
			fmt.Println("0")
		case 1:
			fmt.Println("1")
		case 2:
			// 输出3
			fmt.Println("before")
			fallthrough
			// fallthrough后添加语句会编译报错
			//fmt.Println("after")
		case 3:
			fmt.Println("3")
		case 4, 5, 6:
			fmt.Println("4, 5, 6")
		default:
			fmt.Println("Default")
		}
		Num := 5
		// switch后表达式非必须
		switch  {
		case 0 <= Num && Num <= 3:
			fmt.Println("0-3")
		case 4 <= Num && Num <= 6:
			fmt.Println("4-6")
		case 7 <= Num && Num <= 9:
			fmt.Println("7-9")
		}
	*/

	// 循环
	/*
			sum1 := 0
			for i := 0; i < 10 ; i++ {
				sum1 += i
			}
			fmt.Println(sum1)

			sum := 0
			for {
				sum++
				if sum > 100 {
					break
				}
			}
			fmt.Println(sum)

			a := []int{1, 2, 3, 4, 5, 6}
			for i, j := 0, len(a) - 1; i < j; i, j = i + 1, j - 1 {
				a[i], a[j] = a[j], a[i]
			}
			for _, value := range a {
				fmt.Print(value, " ")
			}
			fmt.Println()

		JLoop:
			for j := 0; j < 5; j++ {
				for i := 0; i < 10; i++ {
					if i > 5 {
						break JLoop
					}
					fmt.Println(i)
				}
			}
	*/

	// goto
	//myfunc()

	// 不定参数
	/*
		myfunc2(5,6,2,4)
		//var a = []int{1,66,7}
		//	//fmt.Println(a)
		// 不定参数的传递
		myfunc3(1,2,5,7)
	*/
}
