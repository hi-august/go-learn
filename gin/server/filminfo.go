package main

import (
	"flag"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
	"github.com/jie123108/glog"
	"go-learn/gin/server/common"
	"go-learn/gin/server/view"
	base "gobaselib"
	"net/http"
	"os"
	"runtime"
)

func main() {
	var cfg string
	flag.StringVar(&cfg, "ucfg", "upload.ini", "ini upload config filename")
	flag.Parse()
	defer glog.Flush()
	glog.Errorf("############# Build Time: %s #############", base.BuildTime)

	Config := common.LoadConfig(cfg)
	if Config == nil {
		os.Exit(1)
	}

	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	r := gin.Default()
	// r.Use(access.TokenCheck())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/filminfo/list", view.FilminfoList)
	r.GET("/filminfo/detail", view.FilminfoDetail)
	r.GET("/filmmaker/list", view.FilmmakerList)
	r.GET("/filmmaker/detail", view.FilmmakerDetail)
	// r.POST("/operation/transcode/cancel_aliyun_upload", transcode.AliUploadCancel)
	s := &http.Server{
		Addr:         Config.Listen,
		Handler:      r,
		ReadTimeout:  Config.ReadTimeOut,
		WriteTimeout: Config.WriteTimeOut,
	}
	err := gracehttp.Serve(s)
	if err != nil {
		glog.Errorf("start server error:%s", err.Error())
	}
}
