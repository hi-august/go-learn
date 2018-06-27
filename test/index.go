package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Hello struct{}

func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "Hello!")
	fmt.Println("hello")
}

func main() {
	var h Hello
	// http.ListenAndServe("localhost:4000", h)
	s := &http.Server{
		Addr:           ":4000",
		Handler:        h,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		log.Println(s.ListenAndServe())
		log.Println("server shutdown")
	}()
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	log.Println(<-ch)
	// 关闭server
	// 执行 Shutdown 时如果传 nil，并且有未完成的请求，会报错
	// 正确的方式是执行 Shutdown 时传入一个非 nil 的 context.Context
	log.Println(s.Shutdown(nil))

	time.Sleep(time.Second * 5)
	log.Println("done.")
}
