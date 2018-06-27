package common

import (
	// "strings"
	"github.com/go-ini/ini"
	"github.com/jie123108/glog"
	"time"
)

var Config UConfig

type UConfig struct {
	LogLevel        string
	Listen          string
	ReadTimeOut     time.Duration
	WriteTimeOut    time.Duration //http read write time out
	MgoUrl          string
	MgoDb           string
	MgoFilminfoCol  string
	MgoFilmmakerCol string
	MgoTimeout      time.Duration
}

func LoadConfig(configFilename string) *UConfig {
	cfg, err := ini.Load(configFilename)
	if err != nil {
		glog.Errorf("LoadConfig failed! err: %s", err)
		return nil
	}
	secLog, err := cfg.GetSection("log")
	if err != nil {
		glog.Errorf("LoadConfig failed! err: %s", err)
		return nil
	}
	LogLevel := secLog.Key("LogLevel").MustString("INFO")
	glog.SetLevelString(LogLevel)
	glog.Infof("Log Level Set to : %s", LogLevel)
	Config.LogLevel = LogLevel

	secSvr, err := cfg.GetSection("server")
	if err != nil {
		glog.Errorf("LoadConfig failed! err: %s", err)
		return nil
	}
	Config.Listen = secSvr.Key("Listen").MustString(":90")
	Config.ReadTimeOut = secSvr.Key("ReadTimeOut").MustDuration(time.Minute * 15)
	Config.WriteTimeOut = secSvr.Key("WriteTimeOut").MustDuration(time.Minute * 15)

	mgoDb, err := cfg.GetSection("mgo")
	if err != nil {
		glog.Errorf("LoadConfig failed! err: %s", err)
		return nil
	}
	Config.MgoUrl = mgoDb.Key("Url").MustString("mongodb://127.0.0.1:27017")
	Config.MgoTimeout = mgoDb.Key("Timeout").MustDuration(time.Minute * 15)
	Config.MgoFilminfoCol = mgoDb.Key("FilminfoCol").MustString("filminfo")
	Config.MgoFilmmakerCol = mgoDb.Key("FilmmakerCol").MustString("filmmaker")
	Config.MgoDb = mgoDb.Key("Db").MustString("test")
	return &Config
}
