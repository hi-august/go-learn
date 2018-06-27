package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
	"strconv"
)

func Show(args ...interface{}) {
	fmt.Println(args)
}

func ShowType(args interface{}) {
	t := reflect.TypeOf(args)
	Show(t.Name(), t.Kind())
}

func StringToInt(str string) (value int, err error) {
	value, err = strconv.Atoi(str)
	return value, err
}

func IntToString(value int) (strvalue string) {
	strvalue = strconv.Itoa(value)
	return
}

func CheckQueryIntField(c *gin.Context, key string) (value int, err error) {
	strvalue := c.Query(key)
	if 0 == len(strvalue) {
		value = 0
		return
	}
	value, err = StringToInt(strvalue)
	return
}
