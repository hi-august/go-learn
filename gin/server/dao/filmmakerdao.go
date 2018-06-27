package dao

import (
	"go-learn/gin/server/common"
	// "github.com/jie123108/glog"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	// "time"
	// "os"
)

type FilmmakerFields struct {
	ID          bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	Name        string        `json:"name" bson:"name"`
	Fmid_db     int64         `json:"fmid_db" bson:"fmid_db"`
	Fmid        int64         `json:"fmid" bson:"fmid"`
	Update_time uint32        `json:"update_time" bson:'update_time'`
	Create_time uint32        `json:"create_time" bson:"create_time"`
	Status      int           `json:"status" bson:"status"`
	Pingfen     float64       `json:"pinfen" bson:"Pingfen"`
	Imdb_no     string        `json:"imdb_no" bson:"imdb_no"`
	Occupation  []int         `json:"occupation" bson:"occupation"`
	Areatype    string        `json:"areatype": bson:"areatype"`
	Horoscope   string        `json:"horoscope" bson:"horoscope"`
	Birthplace  string        `json:"birthplace" bson:"birthplace"`
	Sex         int           `json:"sex" bson:"sex"`
	Desc        string        `json:"desc" bson:"desc"`
	Films       []FilmsFields `json:"films" bson:"films"`
}

type FilmsFields struct {
	Doubanid   int    `json:"doubanid" bson:"doubanid"`
	Occupation int    `json:"occupation" bson:"occupation"`
	Name       string `json:"name" bson:"name"`
	Fid        int    `json:"fid" bson:"fid"`
	Year       int    `json:"year" bson:"year"`
}

func NewFilmmakerDao() *NewDao {
	dao := new(NewDao)
	// common.Show(*Config)
	dao.session = common.InitMgo(Config.MgoUrl, Config.MgoTimeout)
	dao.session.SetSafe(&mgo.Safe{})
	dao.coll = dao.session.DB(Config.MgoDb).C(Config.MgoFilmmakerCol)
	return dao
}

func (this *NewDao) FilmmakerList(limit, skip_num int) (result []*FilmmakerFields, total int, err error) {
	selector := bson.M{}
	result = make([]*FilmmakerFields, 0, 20)
	query := this.coll.Find(selector).Skip(skip_num).Limit(limit)
	total, err = query.Count()
	err = query.All(&result)
	return
}

func (this *NewDao) FilmmakerDetail(fmid_db int) (result *FilmmakerFields, err error) {
	selector := bson.M{}
	if fmid_db != 0 {
		selector["fmid_db"] = fmid_db
	}
	result = &FilmmakerFields{}
	err = this.coll.Find(selector).One(result)
	return
}
