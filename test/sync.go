package main

import "sync"

var (
	mu      sync.Mutex   // guards balance,互斥锁
	mu2     sync.RWMutex // 读写锁,用于读次数远远多于写次数的场景
	balance int
)

// Mutex会保护共享变量
func Deposit(amount int) {
	mu.Lock()
	balance = balance + amount
	mu.Unlock()
}
func Balance() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}
func Balance2() int {
	mu2.RLock()
	b := balance
	mu2.RUnlock()
	return b
}

func main() {
	balance = 100
}
