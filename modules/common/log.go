package common

import (
	"iceforg/pkg/config"
	"iceforg/pkg/log"

	"github.com/sirupsen/logrus"
)

type iceLog struct {
	logger *logrus.Logger
}

var IceLog = new(iceLog)

func LogInit(logConf ...*config.Log) {
	IceLog.logger = log.GetLogrusLogger(logConf...)
}

func SetLogConfig(logConf ...*config.Log) {
	log.SetLogConfig(IceLog.logger, logConf...)
}

func (l *iceLog) Debug(arg ...interface{}) {
	IceLog.logger.Debug(arg)
}

func (l *iceLog) Debugf(format string, arg ...interface{}) {
	IceLog.logger.Debugf(format, arg)
}

func (l *iceLog) Info(arg ...interface{}) {
	IceLog.logger.Info(arg)
}

func (l *iceLog) Infof(format string, arg ...interface{}) {
	IceLog.logger.Infof(format, arg)
}

func (l *iceLog) Error(arg ...interface{}) {
	IceLog.logger.Error(arg)
}

func (l *iceLog) Fatalf(format string, arg ...interface{}) {
	IceLog.logger.Fatalf(format, arg)
}
