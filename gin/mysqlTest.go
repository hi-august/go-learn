package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
)

func main() {
	db, err := sql.Open("mysql", "root:Cxk!51789@tcp(127.0.0.1:3306)/web?charset=utf8&autocommit=true")
	checkErr(err)
	// 添加&&修改
	// 存在就更新,不存在会插入
	// replace自增id会加1,可能会出现冲突,慎用replace
	stmt, err := db.Prepare("insert user set username=?, password=? on duplicate key update password=?")
	checkErr(err)
	res, err := stmt.Exec("lily", "11", "11")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
	// 删除
	stmt, err = db.Prepare("delete from user where username=?")
	checkErr(err)
	name := "lily"
	res, err = stmt.Exec(name)
	checkErr(err)
	// GetMembers(res)
	ret, err := res.RowsAffected()
	checkErr(err)
	if ret == 1 {
		fmt.Printf("成功删除: %s\n", name)
	}
	// 查询
	// where子句
	// 只返回5个结果
	rows, err := db.Query("select * from user where username='august' limit 5")
	checkErr(err)
	for rows.Next() {
		var id int
		var username string
		var password string
		err := rows.Scan(&id, &username, &password)
		checkErr(err)
		fmt.Println(id, username, password)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

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
