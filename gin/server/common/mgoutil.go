package common

import (
	"github.com/jie123108/glog"
	mgo "gopkg.in/mgo.v2"
	"time"
)

var g_mgo *mgo.Session

// [mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
// mongodb://myuser:mypass@localhost:40001,otherhost:40001/mydb
// https://godoc.org/labix.org/v2/mgo#Dial
func InitMgo(mongo_url string, timeout time.Duration) *mgo.Session {
	if g_mgo == nil {
		var err error
		g_mgo, err = mgo.DialWithTimeout(mongo_url, timeout)
		if err != nil {
			glog.Fatalf("open mongodb(%s) failed! err: %v", mongo_url, err)
			return nil
		} else {
			glog.Infof("open mongodb(%s) ok!", mongo_url)
		}
	}
	return g_mgo.Copy()
}
