package main

import (
	"fmt"
	"time"
)

func main() {
	f()
	fmt.Println("end")
}

func f() {
	defer func() { //必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("xiaorui.cc start")
		if err := recover(); err != nil {
			fmt.Println(err) //这里的err其实就是panic传入的内容，"bug"
		}
		fmt.Println("xiaorui.cc end")
	}()
	for {
		fmt.Println("1")
		a := []string{"a", "b"}
		fmt.Println(a[3]) // 越界访问，肯定出现异常
		panic("bug")      // 上面已经出现异常了,所以肯定走不到这里了。
		fmt.Println("4")  //不会运行的.
		time.Sleep(1 * time.Second)
	}
}
