package common

import (
	"database/sql"
	"github.com/jie123108/glog"
	// "github.com/ziutek/mymysql/godrv"
	_ "github.com/go-sql-driver/mysql"
)

var g_db *sql.DB

func InitDb(mysql_url string) *sql.DB {
	if g_db != nil {
		return g_db
	}

	var err error
	g_db, err = sql.Open("mysql", mysql_url)
	if err != nil {
		glog.Fatalf("open mysql(%s) failed! err: %v", mysql_url, err)
		return nil
	} else {
		glog.Infof("open mysql(%s) ok!", mysql_url)
	}
	// TODO: 数据库连接池相关参数设置。
	// g_db.SetConnMaxLifetime(d)
	g_db.SetMaxIdleConns(50)
	g_db.SetMaxOpenConns(100)

	return g_db
}
