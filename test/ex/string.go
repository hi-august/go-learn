package main

import (
	"bytes"
	"fmt"
	s "strings"
)

var p = fmt.Println
var buffer bytes.Buffer

func main() {
	p("contains", s.Contains("test", "es"))         //包含
	p("substring", "test"[:2])                      //切片
	p("count", s.Count("test", "t"))                //统计个数
	p("hasprefix", s.HasPrefix("test", "te"))       //前缀
	p("hassuffix", s.HasSuffix("test", "st"))       //后缀
	p("index", s.Index("test", "s"))                //第一个出现位置
	p("lastindex", s.LastIndex("test", "t"))        //最后一个位置
	p("join", s.Join([]string{"a", "b"}, "-"))      //拼接
	p("repeat", s.Repeat("a", 5))                   //重复5次
	p("replace", s.Replace("fooooo", "o", "0", -1)) //替换字符,-1表示替换到末尾
	p("split", s.Split("a-b-c-d", "-"))             //分割字符串
	p("tolower", s.ToLower("TEST"))                 //转成小写
	p("toupper", s.ToUpper("test"))                 //转成大写
	p("len", len("hello"))                          //字符串长度
	p("char", "test"[1])                            //101
	fmt.Printf("%c", "test"[1])                     //获取字符串元素
	s2 := "name: "
	s3 := "august"
	buffer.WriteString(s2)
	buffer.WriteString(s3)
	fmt.Println(buffer.String())
}
