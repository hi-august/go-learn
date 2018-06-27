package main

import (
	"fmt"
	// "sync"
)

func main() {
	// wg := new(sync.WaitGroup)
	// wg.Add添加或者减少等待gocortinue数量
	// wg.Done相当于Add(-1)
	// wg.Wait执行阻塞,直到WaitGroup数量为0

	// channel相比于waitgroup不仅可以实现协程的同步
	// 还可以控制协程的数量
	fmt.Println("ok")
	// channel是一个阻塞管道,是自动阻塞
	// channel分为两类,一为无缓冲,二为缓冲信道
	// 使用带缓冲的channel时要注意放入和取出数据的速率
	// 带缓冲的channel只有在缓冲区满的时候阻塞
	data := make(chan int, 3) // 数据交换队列,缓存信道,缓冲区可以存储3个元素
	exit := make(chan bool)   // 退出通知
	// 调度器不能保证多个goroutine的执行次序,
	// 而且进程退出时不会等待他们结束
	// 默认情况下仅允许一个系统线程服务于gocortinue
	// 标准库runtime.GOMAXPROCS可修改,实现多核并发,而不仅是并发
	data <- 1 // 发送数据到data channel
	data <- 2
	data <- 3
	// data <- 4 // 此行就会发生阻塞,之前的三个已经填满了data的channel
	fmt.Println(len(data))
	go func() {
		for d := range data { // 在缓存未空前不会阻塞
			fmt.Println(d)
		}
		fmt.Println("recv over.")
		exit <- true // 发出退出通知
	}()
	data <- 5
	data <- 6
	close(data) // 关闭队列

	fmt.Println("send over.")
	<-exit // 等待退出通知, ?如何实现
}
