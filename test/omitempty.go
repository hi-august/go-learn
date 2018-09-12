package main

import (
	"encoding/json"
	"fmt"
)

// omitempty为空则不输出此字段
type Project struct {
	Name string `json:"name"`
	Url  string `json:"url"`
	Docs string `json:"docs,omitempty"`
}

func main() {
	p1 := Project{
		Name: "CleverGo高性能框架",
		Url:  "https://github.com/headwindfly/clevergo",
	}

	data, err := json.Marshal(p1)
	if err != nil {
		panic(err)
	}

	// p1 没有为Docs赋值，这里打印出来不会出现Docs的字段
	fmt.Printf("%s\n", data)

	p2 := Project{
		Name: "CleverGo高性能框架",
		Url:  "https://github.com/headwindfly/clevergo",
		Docs: "https://github.com/headwindfly/clevergo/tree/master/docs",
	}

	data2, err := json.Marshal(p2)
	if err != nil {
		panic(err)
	}

	// p2 则会打印所有
	fmt.Printf("%s\n", data2)

}
