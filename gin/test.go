package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

func GetMembers(i interface{}) {
	// 获取 i 的类型信息
	t := reflect.TypeOf(i)
	for {
		// 进一步获取 i 的类别信息
		if t.Kind() == reflect.Struct {
			// 只有结构体可以获取其字段信息
			fmt.Printf("\n%-8v %v 个字段:\n", t, t.NumField())
			// 进一步获取 i 的字段信息
			for i := 0; i < t.NumField(); i++ {
				fmt.Println(t.Field(i).Name)
			}
		}
		// 任何类型都可以获取其方法信息
		fmt.Printf("\n%-8v %v 个方法:\n", t, t.NumMethod())
		// 进一步获取 i 的方法信息
		for i := 0; i < t.NumMethod(); i++ {
			fmt.Println(t.Method(i).Name)
		}
		if t.Kind() == reflect.Ptr {
			// 如果是指针，则获取其所指向的元素
			t = t.Elem()
		} else {
			// 否则上面已经处理过了，直接退出循环
			break
		}
	}
}

func main() {
	// 调用最基本的GET,并获得返回值
	resp, _ := http.Get("http://127.0.0.1:8000/v1/ping")
	// GetMembers(resp.Request)
	helpRead(resp)
	resp, _ = http.Get("http://127.0.0.1:8000/v1/status")
	helpRead(resp)

	resp, _ = http.Get("http://127.0.0.1:8000/v1/get?user=august&pwd=123456")
	helpRead(resp)
	// 调用最基本的POST,并获得返回值
	resp, _ = http.Post("http://127.0.0.1:8000/v1/post2?user=august&pwd=123456", "application/x-www-form-urlencoded", strings.NewReader(""))
	helpRead(resp)
	resp, _ = http.Post("http://127.0.0.1:8000/v1/post", "application/x-www-form-urlencoded", strings.NewReader("user=august&pwd=123456"))
	helpRead(resp)

	resp, _ = http.Get("http://127.0.0.1:8000/v1/get")
	helpRead(resp)
}

// 封装了fmt,ioutil方法
func helpRead(resp *http.Response) {
	// 关闭http连接
	// 在函数退出时调用
	defer resp.Body.Close()
	// 用ioutil读取
	body, err := ioutil.ReadAll(resp.Body)
	// 错误处理
	if err != nil {
		fmt.Println("ERROR2!: ", err)
	}
	fmt.Println(string(body), resp.Request.URL)
}
