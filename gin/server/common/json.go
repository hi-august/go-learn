package common

import (
	"github.com/gin-gonic/gin"
)

func JsonOk(data gin.H) gin.H {
	return gin.H{"ok": true, "reason": "", "data": data}
}

func JsonFail(reason string) gin.H {
	return gin.H{"ok": false, "reason": reason}
}

func JsonFailEx(reason string, data gin.H) gin.H {
	return gin.H{"ok": false, "reason": reason, "data": data}
}
