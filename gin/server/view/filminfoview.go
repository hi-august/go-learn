package view

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"go-learn/gin/server/common"
	"go-learn/gin/server/dao"
)

func FilminfoList(c *gin.Context) {
	// doubanid := c.Query("doubanid")
	// headerform := c.Request.Form
	// doubanid, _ := strconv.Atoi(headerform.Get("doubanid"))
	// doubanid, _ := strconv.Atoi(c.Query("doubanid"))
	// doubanid, _ := common.CheckQueryIntField(c, "doubanid")
	limit, _ := common.CheckQueryIntField(c, "limit")
	page, _ := common.CheckQueryIntField(c, "page")
	if limit <= 0 || limit > 10 {
		limit = 10
	}
	if page < 0 {
		page = 1
	} else if page == 1 {
		page = 0
	}
	skip_num := page * limit
	if skip_num >= 10000 {
		c.JSON(200, gin.H{
			"data":   "",
			"ok":     false,
			"reason": "超过10000",
		})
		return
	}
	filminfodao := dao.NewFilminfoDao()
	defer filminfodao.Free()
	data, total, _ := filminfodao.FilminfoList(limit, skip_num)
	c.JSON(200, common.JsonOk(gin.H{"films": data, "total": total}))
	// c.JSON(200, gin.H{
	// "data":  data,
	// "ok":    true,
	// "total": total,
	// })
}

func FilminfoDetail(c *gin.Context) {
	doubanid, _ := common.CheckQueryIntField(c, "doubanid")
	if doubanid == 0 {
		c.JSON(200, gin.H{
			"data":   "",
			"ok":     false,
			"reason": "args doubanid err",
		})
		return
	}
	filminfodao := dao.NewFilminfoDao()
	defer filminfodao.Free()
	data, err := filminfodao.FilminfoDetail(doubanid)
	// common.Show(err)
	if err != nil {
		c.JSON(200, gin.H{
			"data":   "",
			"ok":     false,
			"reason": err.Error(),
		})
		return
	}
	c.JSON(200, common.JsonOk(gin.H{"film": data}))
}
