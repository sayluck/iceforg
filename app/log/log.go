package log

import (
	"context"
	"fmt"
	"iceforg/app/common"
	"iceforg/pkg/config"
	"iceforg/pkg/log"
	"iceforg/pkg/utils"

	"github.com/sirupsen/logrus"
)

type iceLog struct {
	Logger *logrus.Logger
}

var IceLog = new(iceLog)

func LogInit(logConf ...*config.Log) {
	IceLog.Logger = log.GetLogrusLogger(logConf...)
}

func SetLogConfig(logConf ...*config.Log) {
	log.SetLogConfig(IceLog.Logger, logConf...)
}

func prepareLogMsg(c context.Context, arg ...interface{}) string {
	var reqID interface{}
	if c.Value(common.ReqID) != nil {
		reqID = c.Value(common.ReqID)
	} else {
		reqID = utils.GenerateUUID(15)
	}
	return fmt.Sprintf("reqID:%v,detial:%v",
		reqID, arg)
}

func (l *iceLog) Debug(c context.Context, arg ...interface{}) {
	IceLog.Logger.Debug(prepareLogMsg(c, arg))
}

func (l *iceLog) Errorf(c context.Context, format string, arg ...interface{}) {
	IceLog.Logger.Errorf(format, prepareLogMsg(c, arg))
}

func (l *iceLog) Debugf(c context.Context, format string, arg ...interface{}) {
	IceLog.Logger.Debugf(format, prepareLogMsg(c, arg))
}

func (l *iceLog) Error(c context.Context, arg ...interface{}) {
	IceLog.Logger.Error(prepareLogMsg(c, arg))
}
func (l *iceLog) Fatalf(c context.Context, format string, arg ...interface{}) {
	IceLog.Logger.Fatalf(format, prepareLogMsg(c, arg))
}
