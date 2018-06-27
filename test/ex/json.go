package main

import (
    "github.com/bitly/go-simplejson"
    "encoding/json"
    "fmt"
)

func main() {
    var s = make(map[string]interface{})
    s["username"] = "august"
    s["sex"] = "male"
    s["age"] = 23
    // 编码到json数据
    result, _ := json.Marshal(s)
    fmt.Print(string(result)) // 需string
    // var r interface{}
    // err := json.Unmarshal([]byte(result), &r)
    // fmt.Println(r)

    // json解码
    js, err := simplejson.NewJson(result)

    if err != nil {
        panic(err)
    }
    fmt.Printf("\n%v", js)

    // js, err := simplejson.NewJson()
}
