/*
并发：goroutines和channels提供一种简洁的方式来实现并发的计算
Goroutines：并发执行单位，类似轻量级线程，由Go调度运行，用户无须手动分配线程，使用go关键值启动goroutine，非阻塞的，可以高效运行成千上万个goroutine
Channel：用于goroutine之间通信的通道，可以是无缓冲的，也可以是有缓冲的，可以是同步的，也可以是异步的。支持同步和数据共享，避免显式锁机制，使用chan关键字创建，通过<-操作符发送和接收数据。
默认情况下，通道是不带缓冲区的。发送端发送数据，同时必须有接收端接收相应的接收数据。
通过make中第二个参数指定缓冲区大小。由于缓冲区大小是有限的，所以必须接收端接收数据，否则缓冲区一满，数据发送端就无法再发送数据了。
如果通道不带缓冲，发送方会阻塞直接到接收方从通道中接收数据。如果通道带缓冲区，发送方则会阻塞直到发送的值被拷贝到缓冲区内；
如果缓冲区已满，则意味着需要等待直到某个接收方获取到一个值，接收方在有值可以接收之前会一直阻塞
select使得一个goroutine可以同时等待多个通道操作，select会阻塞，直到其中某个case可以继续执行
使用sync.WaitGroup等待多个goroutine完成
Context：用于控制goroutine生命周期：context.WithCancel, context.WithTimeout
Mutex和RWMutex：互斥锁和读写互斥锁
Scheduler：调度器，将Goroutine分配到系统线程中执行，并通过M和P配合高效管理并发

	G：Goroutine
	P：Processor（逻辑处理器）
	M：Machine（系统线程）

2025.4.17 by dralee
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func hello() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello")
		time.Sleep(time.Millisecond * 100)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func channelNoBuffer() {
	s := []int{5, 1, 8, -2, 3, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c
	fmt.Println(x, y, x+y)
}

func channelBuffer() {
	c := make(chan int, 2) // 设置缓存区为2
	// 因为带缓冲区，可同时发送两个数据，而不用立刻去同步读取数据

	c <- 1
	c <- 2

	// 开始读取
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // 关闭通道
}
func fibonacciTest() {
	fmt.Println("fibonacci:")
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c { // 由于fibonacci函数中已经关闭通道，所以这里不会阻塞，读取完就会结束；否则会阻塞
		fmt.Println(i)
	}
}

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func selectChannel() {
	fmt.Println("selectChannel:")
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci2(c, quit)
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // goroutine结束时调用Done方法
	fmt.Printf("worker %d starting\n", id)
	fmt.Printf("worker %d done\n", id)
}

func workerTest() {
	fmt.Println("workerTest:")
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1) // 增加计数器
		go worker(i, &wg)
	}

	wg.Wait() // 等待所有goroutine结束
	fmt.Println("all done")
}

func main() {
	go hello() // 启动goroutine
	for i := 0; i < 10; i++ {
		fmt.Println("main")
		time.Sleep(time.Millisecond * 100)
	}

	channelNoBuffer()
	channelBuffer()

	fibonacciTest()
	selectChannel()

	workerTest()

}
