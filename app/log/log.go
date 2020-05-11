package log

import (
	"bytes"
	"fmt"
	"iceforg/app/common"
	"iceforg/pkg/config"
	"iceforg/pkg/log"
	"io/ioutil"

	"github.com/gin-gonic/gin"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func LogInit(logConf ...*config.Log) {
	Log = log.GetLogrusLogger(logConf...)
}

func SetStartBaseLog(c *gin.Context) string {
	if c.Request.Method == common.MethodGet {
		return fmt.Sprintf("reqID:%s,URL:%s",
			c.GetString(common.ReqID),
			c.Request.RequestURI)
	}
	bodyStr := ""
	body, err := ioutil.ReadAll(c.Request.Body)
	if err == nil {
		bodyStr = string(body)
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return fmt.Sprintf("reqID:%s,URL:%s,Body:%+v",
		c.GetString(common.ReqID),
		c.Request.RequestURI,
		bodyStr)
}

func SetEndBaseLog(c *gin.Context) string {
	return fmt.Sprintf("reqID:%s,Status:%d",
		c.GetString(common.ReqID),
		c.Writer.Status())
}
