package main

// go build -gcflags "-N -l" gdb.go
// 生成可执行文件
// gdb gdb 进入调试模式
// b 20 在20行断点
// run 执行
// info locals 查看所有变量
// p count 查看变量p
// go get github.com/derekparker/delve/cmd/dlv
// make install
// dlv debug gdb_dlv.go
// b main.main 断点
// p msg 查看变量

import (
	"fmt"
	"time"
)

type Person struct {
    name string
    age int
}

func test() (p *Person) {
    p = &Person{} // p为指针
    p.name = "august"
    p.age = 23
    return
}

func counting(c chan<- int) {
	for i := 0; i < 10; i++ {
		time.Sleep(2 * time.Second)
		c <- i
	}
	close(c)
}

func main() {
	msg := "Starting main"
	fmt.Println(msg)
    p := test()
    fmt.Println(p)
	bus := make(chan int)
	msg = "starting a gofunc"
	go counting(bus)
	for count := range bus {
		fmt.Println("count:", count)
	}
}
