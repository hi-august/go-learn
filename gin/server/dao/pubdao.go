package dao

import (
	"github.com/jie123108/glog"
	"go-learn/gin/server/common"
	mgo "gopkg.in/mgo.v2"
	"os"
)

type NewDao struct {
	session *mgo.Session
	coll    *mgo.Collection
}

func (this *NewDao) Free() {
	this.session.Close()
}

func InitConfig() (Config *common.UConfig) {
	// cfg := "upload.ini"
	cfg := "server/upload.ini"
	Config = common.LoadConfig(cfg)
	common.ShowType(*Config)
	glog.Errorf("cfg path %s", cfg)
	if Config == nil {
		os.Exit(1)
	}
	return
}

var Config = InitConfig()
