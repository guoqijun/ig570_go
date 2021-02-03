package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"ig570_go/conf"
	"strconv"
	"time"
)

type Gin struct {
	C *gin.Context
}

type ds struct {
}

func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	if data == nil {
		data = ds{}
	}
	msg := GetMsg(errCode)
	beginTime, _ := strconv.ParseInt(g.C.Writer.Header().Get("X-Begin-Time"), 10, 64)

	duration := time.Now().Sub(time.Unix(0, beginTime))
	milliseconds := float64(duration) / float64(time.Millisecond)
	rounded := float64(int(milliseconds*100+.5)) / 100
	roundedStr := fmt.Sprintf("%.3fms", rounded)
	g.C.Writer.Header().Set("X-Response-time", roundedStr)
	g.C.JSON(httpCode, gin.H{
		"code":    errCode,
		"message": msg,
		"data":    data,
	})

	conf.Logger.WithFields(logrus.Fields{
		"msg": "info",
	}).Info("message", "API Response", "code", errCode, "errMsg", msg, "took", roundedStr)

	g.C.Abort()
	return
}
