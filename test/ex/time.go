package main

import (
    "time"
    "fmt"
    "reflect"
    _ "../../util"
)

func main() {
    now := time.Now() // 返回本地时间
    fmt.Println(now)
    fmt.Println(now.Local())
    fmt.Println(now.Year()) // 年份
    fmt.Println(now.Month(), now.Day) // 月/日
    fmt.Println(now.Hour(), now.Minute(), now.Second()) // 时分秒
    fmt.Println(now.Weekday()) // 返回星期几
    now_str := now.Format("2006-01-02 15:04:05") // Time转为字符串
    tp2 := reflect.TypeOf(now)
    // util.GetMembers(tp)
    time.Sleep(3*time.Second) // 暂停3秒
    fmt.Println(reflect.TypeOf(now_str).String() == "string") // 字符串比较
    fmt.Println(tp2.Name() == "Time")
}

