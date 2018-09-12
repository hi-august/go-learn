package main

import (
	"fmt"
	"reflect"
)

func main() {
	u := User{Name: "august", Age: 22}
	t := reflect.TypeOf(u)
	fmt.Println(t, u)

	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		fmt.Println(sf.Tag.Get("json"), ",", sf.Tag.Get("bson"))
	}

}

type User struct {
	Name string `json:"name" bson:"b_name"`
	Age  int    `json:"age" bson:"b_age"`
}
