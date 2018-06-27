package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Person struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	Name       string
	Phone      string
	UpdateTime time.Time
}

var (
	IsDrop = true
)

func main() {
	session, err := mgo.Dial("127.0.0.1") // 创建一个新的会话
	// 错误处理
	if err != nil {
		panic(err)
	}
	// 在函数退出前调用关闭会话
	defer session.Close()

	// 删除数据库
	// if IsDrop {
	// err = session.DB("test").DropDatabase()
	// if err != nil {
	// panic(err)
	// }
	// }

	// 创建数据库及集合
	c := session.DB("test").C("people")

	// index := mgo.Index{
	// Key:           []string{"name", "phone"},
	// Unique:        true, // 唯一
	// DropDups:      true, // 重复会删除
	// Background:    true, // 避免长时间占用锁
	// Sparse:        true, // 稀疏索引(为空值时不进入索引),避免duplicate key error报错
	// }
	// 创建索引
	// err = c.EnsureIndex(index)
	// if err != nil {
	// panic(err)
	// }
	// update&&insert
	selector := bson.M{"name": "august"}
	change := bson.M{"$set": bson.M{"phone": "10001", "updateTime": time.Now()}} // 字段首字母小写
	changeInfo, err := c.Upsert(selector, change)                                // upsert返回两个值
	if err != nil {
		panic(err)
	}
	fmt.Println(changeInfo) // {1 0 1 <nil>} 第一个1表示修改成功, 全为0表示新添加的内容

	result := Person{} // query one
	err = c.Find(bson.M{"name": "august"}).One(&result)
	fmt.Println(result)
	var n int
	n, err = c.Find(bson.M{"name": "august"}).Count() // count
	if err != nil {
		panic(err)
	}
	fmt.Println(n)

	var results []Person // query all
	// err = c.Find(nil).Sort("name").All(&results)
	// if err != nil {
	// panic(err)
	// }
	iter := c.Find(nil).Sort("name").Skip(0).Limit(3).Iter() // limit, skip(一般可以用来做分页)
	err = iter.All(&results)
	if err != nil {
		panic(nil)
	}
	fmt.Println(results)
}
