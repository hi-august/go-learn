package view

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"go-learn/gin/server/common"
	"go-learn/gin/server/dao"
)

func FilmmakerList(c *gin.Context) {
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
	filmmakerdao := dao.NewFilmmakerDao()
	defer filmmakerdao.Free()
	data, total, _ := filmmakerdao.FilmmakerList(limit, skip_num)
	c.JSON(200, common.JsonOk(gin.H{"fminfo": data, "total": total}))
}

func FilmmakerDetail(c *gin.Context) {
	fmid_db, _ := common.CheckQueryIntField(c, "fmid_db")
	if fmid_db == 0 {
		c.JSON(200, gin.H{
			"data":   "",
			"ok":     false,
			"reason": "args fmid_db err",
		})
		return
	}
	filmmakerdao := dao.NewFilmmakerDao()
	defer filmmakerdao.Free()
	data, err := filmmakerdao.FilmmakerDetail(fmid_db)
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
