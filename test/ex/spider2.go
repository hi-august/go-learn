package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	// "strings"
	"../../util"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"reflect"
	"sync"
	"time"
)

const (
	localMogoURL = "mongodb://127.0.0.1:27017/filminfo"
	timeOut      = time.Second * 30
	DB           = "test"
	COL          = "filminfo"
)

type Dao struct {
	session *mgo.Session
	coll    *mgo.Collection
	test    string
}

var filmInfoDao *Dao

func newDao(mongo_url, db, col string, timeout time.Duration) *Dao {
	set := new(Dao)
	set.session = InitMgo(mongo_url, timeout)
	set.test = "hi"
	set.session.SetSafe(&mgo.Safe{})
	set.coll = set.session.DB(db).C(col)
	return set
}

func InitMgo(mongo_url string, timeout time.Duration) *mgo.Session {
	client, err := mgo.DialWithTimeout(mongo_url, timeout)
	if err != nil {
		util.Show("open mongodb(%s) failed! err: %v", mongo_url, err)
		return nil
	} else {
		util.Show("open mongodb(%s) ok!", mongo_url)
	}
	return client.Copy()
}

type FilmInfo struct {
	DoubanId int      `bson:"doubanid" json:"doubanid" validate:"required"` // validate取空值
	Url      string   `bson:"url" json:"url"`
	Cookies  []Cookie `bson:"cookies"`
	Imdb     string   `bson:"imdb" json:"imdb", validate:"required"`
}

type Cookie struct {
	Name  string `bson:"name" json:"name"`
	Value string `bson:"value" json:"value"`
}

func get_url(idx int) (url string) {
	url = "https://movie.douban.com/subject/" + strconv.Itoa(idx) + "/"
	return
}

func show(args ...interface{}) {
	fmt.Println(args)
}

func http_get(url string) (filminfo *FilmInfo) {
	// url := get_url(idx)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0")
	res, _ := client.Do(req)
	defer res.Body.Close()
	// util.GetMembers(req)
	cks := make([]Cookie, 0, 20)
	// cks := make(map[interface{}]interface{})
	cookies := res.Cookies()
	var ck Cookie
	for _, cookie := range cookies {
		ck.Name = cookie.Name
		ck.Value = cookie.Value
		cks = append(cks, ck)
		// fmt.Println(cookie, ck)
	}
	// cks["a"] = 1
	// cks[1] = "b"
	// fmt.Println(cks, req.URL, req.Method, req.Header)
	filminfo = &FilmInfo{Url: url, Cookies: cks}
	// show(cks, req.URL, req.Method, req.Header, 111111111)
	// helpRead(res)
	return
}

func qshell(qshellGroup *sync.WaitGroup, idx int) {
	defer qshellGroup.Done()
	// url := "https://movie.douban.com/subject/" + strconv.Itoa(idx) + "/"
	url := get_url(idx)
	filminfo := http_get(url)
	// filminfo.DoubanId = idx
	selector := bson.M{"doubanid": idx}
	filmInfoDao.coll.Upsert(selector, filminfo)
	// res, _ := http.Get(url)
	// helpRead(res)
	show(*filminfo, 111)
}

func main() {
	startTime := time.Now().Unix()
	filmInfoDao = newDao(localMogoURL, DB, COL, timeOut)
	t := reflect.TypeOf(*filmInfoDao)
	util.Show(t.Name(), t.Kind()) // 获取对象名称,类型
	// v := reflect.ValueOf(*filmInfoDao)
	for i := 0; i < t.NumField(); i++ {
		// val := v.Field(i).Interface() // 结构体小写会报错
		f := t.Field(i)
		util.Show(f.Name, f.Type, 111) // 取结构体内名称类型
	}
	defer filmInfoDao.session.Close()
	fmt.Println(startTime)
	ids := []int{27032266, 26805324, 25966044}
	for i, idx := range ids {
		fmt.Println(i, idx)
	}
	// ids = append(ids, 123)
	// ids = append(ids, 456)
	// ids = append(ids, 789)
	qshellGroup := new(sync.WaitGroup)
	for index := 0; index < len(ids); index++ {
		qshellGroup.Add(1)
		time.Sleep(1 * time.Second)
		go qshell(qshellGroup, ids[index])
	}
	qshellGroup.Wait()
	fmt.Printf("total used time:%d second", time.Now().Unix()-startTime)

}

func helpRead(resp *http.Response) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ERROR2!: ", err)
	}
	fmt.Println(string(body), resp.Request.URL)
}
