package main

import (
	"fmt"
	base "gobaselib"
)

func main() {
	fmt.Println(233)
	// 检查文件是否存在
	is_exist := base.IsExist("t.go")
	// 去除后缀
	remove_ext := base.TrimExt("t.go")
	// 计算文件md5值
	file_md5, _ := base.Md5SumFile("t.go", 0)
	// 新建文件夹
	// err := base.MkdirAll("test/test")
	fmt.Println(is_exist, remove_ext, file_md5)
}
