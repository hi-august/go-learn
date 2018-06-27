package main

import (
    "fmt"
    "io"
    "os"
    "bufio"
)

func dealWithFile() {
    // 打开读取的文件
    fi, err := os.Open("ex1.go")
    // 如果有错误panic
    if err != nil {
        panic(err)
    }
    // 在函数退出之前执行
    defer func() {
        if err := fi.Close(); err != nil {
            panic(err)
        }
    }()
    // 创建一个文件
    fo, err := os.Create("out.txt")
    // 如果有错误panic
    if err != nil {
        panic(err)
    }
    // 在函数退出之前执行
    defer func() {
        if err := fo.Close(); err != nil {
            panic(err)
        }
    }()
    // 1.
    // buf := make([]byte, 50)
    // for {
        // n, err := fi.Read(buf)
        // fmt.Println(n)
        // if err != nil {
            // panic(err)
        // }
        // if n == 0 {
            // break
        // }
        // // 执行写入out.txt文件
        // if _, err := fo.Write(buf[:n]); err != nil {
            // panic(err)
        // }
    // }
    // 2. bufio
    r := bufio.NewReader(fi)
    w := bufio.NewWriter(fo)
    buf := make([]byte, 50)
    for {
        n, err := r.Read(buf)
        fmt.Println(n)
        // io.EOF可以认为是最后一行
        if err != nil && err != io.EOF {
            panic(err)
        }
        if n == 0 {
            // 如果是最后一行,break
            break
        }
        if _, err := w.Write(buf[:n]); err != nil {
            panic(err)
        }
    }
    if err = w.Flush(); err != nil {
        panic(err)
    }
}

func main() {
    dealWithFile()
}
