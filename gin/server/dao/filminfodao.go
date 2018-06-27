package dao

import (
	"go-learn/gin/server/common"
	// "github.com/jie123108/glog"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	// "time"
	// "os"
)

type FilminfoFields struct {
	ID          bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	Name        string        `json:"name" bson:"name"`
	Doubanid    int64         `json:"doubanid" bson:"doubanid"`
	Fid         int64         `json:"fid" bson:"fid"`
	Update_time uint32        `json:"update_time" bson:'update_time'`
	Status      int           `json:"status" bson:"status"`
	Pingfen     float64       `json:"pinfen" bson:"Pingfen"`
	Imdb        string        `json:"imdb" bson:"imdb"`
	Imdb_rating float64       `json:"imdb_rating" bson:"imdb_rating"`
	Pingfen_num int           `json:"pingfen_num" bson:"pingfen_num"`
}

type FilminfoDetailFields struct {
	// FilminfoFields          // 接口不显示FilminfoFields???
	ID             bson.ObjectId     `bson:"_id,omitempty" json:"_id"`
	Name           string            `json:"name" bson:"name"`
	Doubanid       int64             `json:"doubanid" bson:"doubanid"`
	Fid            int64             `json:"fid" bson:"fid"`
	Update_time    uint32            `json:"update_time" bson:'update_time'`
	Create_time    uint32            `json:"create_time" bson:"create_time"`
	Status         int               `json:"status" bson:"status"`
	Pingfen        float64           `json:"pinfen" bson:"Pingfen"`
	Imdb           string            `json:"imdb" bson:"imdb"`
	Imdb_rating    float64           `json:"imdb_rating" bson:"imdb_rating"`
	Pingfen_num    int               `json:"pingfen_num" bson:"pingfen_num"`
	Pubdate        int               `bson:"pubdate" json:"pubdate"` // json返回到接口小写,bson存储到mongo转为小写
	Alias          []string          `bson:"alias" json:"alias"`
	Reason         string            `bson:"reason" json:"reason"`
	Pub_status     int               `bson:"pub_status" json:"pub_status"`
	Category       int               `bson:"category" json:"category"`
	Award_record   string            `bson:"award_record" json:"award_record"`
	Episodes       int               `bson:"episodes" json:"episodes"`
	Desc           string            `bson:"desc" json:"desc"`
	Pingfen_detail map[string]string `bson:"pingfen_detail" json:"pingfen_detail"`
}

func NewFilminfoDao() *NewDao {
	dao := new(NewDao)
	dao.session = common.InitMgo(Config.MgoUrl, Config.MgoTimeout)
	dao.session.SetSafe(&mgo.Safe{})
	dao.coll = dao.session.DB(Config.MgoDb).C(Config.MgoFilminfoCol)
	return dao
}

func (this *NewDao) FilminfoList(limit, skip_num int) (result []*FilminfoFields, total int, err error) {
	selector := bson.M{}
	result = make([]*FilminfoFields, 0, 20)
	query := this.coll.Find(selector).Skip(skip_num).Limit(limit)
	total, err = query.Count()
	err = query.All(&result)
	return
}

func (this *NewDao) FilminfoDetail(doubanid int) (result *FilminfoDetailFields, err error) {
	selector := bson.M{}
	if doubanid != 0 {
		selector["doubanid"] = doubanid
	}
	result = &FilminfoDetailFields{}
	common.ShowType(*result)
	err = this.coll.Find(selector).One(result)
	return
}
