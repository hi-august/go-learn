package main

import (
	"fmt"
	// "runtime"
	"sync"
	"time"
)

type Op struct {
	key int
	val int
}

var lock sync.Mutex

var m1 map[int]int
var m2 map[int]int
var max int = 50000

func update_map_by_mutex(i int) {
	lock.Lock()
	// fmt.Printf("sync idx: %d\n", i)
	// println(runtime.NumGoroutine())
	m1[i] = i
	if len(m1) == max {
		fmt.Printf("%s mutex finish\n", time.Now())
	}
	lock.Unlock()
}

var ch chan Op

func update_map_by_chan(i int) {
	// 发送Op结构体到ch通道
	ch <- Op{key: i, val: i}
}

func wait_for_chan(m map[int]int) {
	for {
		select {
		case op := <-ch:
			// fmt.Printf("chan idx is %d\n", op.val)
			m[op.key] = op.val
			if len(m2) == max {
				fmt.Printf("%s chan finish\n", time.Now())
				return
			}
		}
	}
}

func main() {

	// 初始化map,m1&m2
	m1 = make(map[int]int, max)
	m2 = make(map[int]int, max)
	// 初始化通道
	ch = make(chan Op)
	// 开启通道,等待接收数据
	go wait_for_chan(m2)
	for i := 0; i < max; i++ {
		// 并发执行,写操作
		// go关键词并发不受控制,却数量不受控制
		go update_map_by_chan(i)
		go update_map_by_mutex(i)
	}

	time.Sleep(time.Second * 1)
}
