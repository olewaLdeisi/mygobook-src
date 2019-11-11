package main

import "sync"

// channel的传递
type PipeData struct {
	value   int
	handler func(int) int
	next    chan int
}

// channel的传递
func handle(queue chan *PipeData) {
	for data := range queue {
		data.next <- data.handler(data.value)
	}
}

// 多核并行化
type Vector []float64

// 分配给每个CPU的计算任务
func (v Vector) DoSome(i, n int, u Vector, c chan int) {
	for ; i < n; i++ {
		//v[i] += u.Op(v[i])
	}
	c <- 1 // 发信号给任务管理器表示计算已完成
}

const NCPU = 16 // 假设有16核
func (v Vector) DoAll(u Vector) {
	c := make(chan int, NCPU) // 用于接收每个CPU的任务完成信号

	for i := 0; i < NCPU; i++ {
		go v.DoSome(i*len(v)/NCPU, (i+1)*len(v)/NCPU, u, c)
	}

	// 等待所有CPU的任务完成
	for i := 0; i < NCPU; i++ {
		<-c // 获取到一个数据， 表示一个CPU计算完成
	}
	// 到这里表示所有计算已经结束
}

func main() {
	// channel基本语法
	/*
		var ch chan int				//
		var m map[string] chan bool // map，元素是bool型的channel

		ch := make(chan int)

		// ch <- value 将value写入channel
		// value := <- ch	从channel读取值赋给value
	*/

	// select
	/*
		select {
		case <- chan1:
		case chan2 <- 1:
		default:
		}
	*/

	// channel缓冲
	/*
		c := make(chan int, 1024)
		// 循环读取
		for i := range c {
			fmt.Println("Received:", i)
		}
	*/

	// 超时机制
	/*
		timeout := make(chan bool, 1)
		go func() {
			time.Sleep(1e9) // 等待1秒钟
			timeout <- true
		}()
		select {
		case <-ch:
			// 从ch中读取到数据
		case <-timeout:
			// 一直没有从ch中读取到数据，但从timeout中读取到了数据
		}
	*/

	// 单向channel，限制权限，是一种概念
	/*
		var ch1 chan int
		var ch2 chan<- float64
		var ch3 <-chan int
	*/
	/*
		ch4 := make(chan int)
		ch5 := <-chan int(ch4)
		ch6 := chan<- int(ch4)

		func Parse(ch <-chan int) {
			for value := range ch {
				fmt.Println("Parsing value", value)
			}
		}
	*/

	// 关闭channel
	// close(ch)
	// 判断channel是否关闭
	// x, ok := <-ch

}

// 全局唯一性操作
var a string
var once sync.Once

func setup() {
	a = "Hello, world"
}
func doprint() {
	once.Do(setup)
	print(a)
}
func twoprint() {
	go doprint()
	go doprint()
}
