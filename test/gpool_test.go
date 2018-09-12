package main

import (
	"runtime"
	"testing"
	"time"
	"utils"
)

func TestPool(t *testing.T) {
	// 协程池,控制数量
	pool := utils.New(10)
	// 设置并行数量
	runtime.GOMAXPROCS(runtime.NumCPU())
	println(runtime.NumCPU())
	// goroutine数量
	// t.Errorf("NumGoroutine count: %d", runtime.NumGoroutine())
	for i := 0; i < 100; i++ {
		pool.Add(1)
		go func() {
			time.Sleep(time.Second)
			println(runtime.NumGoroutine())
			pool.Done()
		}()
	}
	pool.Wait()
	// t.Errorf("NumGoroutine count: %d", runtime.NumGoroutine())
}
