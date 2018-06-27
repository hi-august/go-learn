package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	filminfodao "go-learn/gin/server/dao"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	// "reflect"
)

func main() {
	// debug环境
	gin.SetMode(gin.DebugMode) // gin.ReleaseMode
	// 注册一个路由
	router := gin.Default()
	// 注册接口
	fmt.Println("starting server...")
	v1 := router.Group("/v1")
	{
		v1.GET("/ping", ping)
		v1.GET("/status", status)
		v1.POST("/post", param_post_form)
		// http://127.0.0.1:8000/get?user=august&pwd=123456
		v1.GET("/get", param_get_query)
		v1.GET("/dao", dao)
		v1.POST("/post2", param_get_query)
	}
	// 1.绑定http server
	router.Run("127.0.0.1:8000") // listen and server on 0.0.0.0:8000
	// 2.
	// http.ListenAndServe(":8000", r)
	// 3. 自定义HTTP服务器配置
	// server := &http.Server{
	// Addr:           ":8000",
	// Handler:        r,
	// ReadTimeout:    10 * time.Second,
	// WriteTimeout:   10 * time.Second,
	// MaxHeaderBytes: 1 << 20,
	// }
	// server.ListenAndServe()
	// 无缝重启、停机, fvbock/endless 来替换默认的 ListenAndServe
}

func param_get_query(c *gin.Context) {
	// 2. Query接受参数
	user := c.Query("user")
	pwd := c.Query("pwd")
	c.JSON(200, gin.H{
		"message": "ok",
		"user":    user,
		"pwd":     pwd,
	})
}
func param_post_form(c *gin.Context) {
	// PostForm接受客户端参数
	// python可以用requests data或者body
	// 1. PostForm
	user := c.PostForm("user")
	pwd := c.PostForm("pwd")
	c.JSON(200, gin.H{
		"message": "ok",
		"user":    user,
		"pwd":     pwd,
	})
}
func ping(c *gin.Context) {
	// 返回一条json数据
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func status(c *gin.Context) {
	// 返回一个字符串
	c.String(200, "ok")
}

func dao(c *gin.Context) {
	// 没有取到doubanid,默认为0
	doubanid, _ := strconv.Atoi(c.Query("doubanid"))
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	if page < 1 {
		page = 1
	}
	if limit > 10 {
		limit = 10
	} else if limit < 1 {
		limit = 10
	}
	page_num := (page - 1) * limit
	selector := bson.M{}
	if doubanid != 0 {
		selector["doubanid"] = doubanid
	}
	fmt.Println(page, limit, page_num, doubanid)
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	col := session.DB("filminfo").C("filminfo")

	retLst := make([]*filminfodao.FilminfoFields, 0, 20) // 取内存地址放入retlst
	// retLst := make([]Filminfo, 0, 20) // 取值放入retlst
	err = col.Find(selector).Skip(page_num).Limit(limit).All(&retLst)
	// fmt.Println(retLst)
	c.JSON(200, gin.H{
		"data": retLst,
		"ok":   true,
	})
}
