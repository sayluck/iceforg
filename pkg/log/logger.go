package log

import (
	"iceforg/pkg/config"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	prettyPrint bool   = false
	level       string = "warn"
)

func GetLogrusLogger(logConf ...*config.Log) *logrus.Logger {

	log := logrus.New()

	if len(logConf) > 0 {
		prettyPrint = logConf[0].PrettyPrint
		level = logConf[0].Level
	}

	log.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: prettyPrint,
	})
	log.SetLevel(getLogLevel(level))

	return log
}

func SetLogConfig(log *logrus.Logger, logConf ...*config.Log) {
	if len(logConf) > 0 {
		prettyPrint = logConf[0].PrettyPrint
		level = logConf[0].Level
	}

	log.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: prettyPrint,
	})
	log.SetLevel(getLogLevel(level))
}

func getLogLevel(level string) logrus.Level {
	switch strings.ToLower(level) {
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	case "fatal":
		return logrus.FatalLevel
	case "panic":
		return logrus.PanicLevel
	case "trace":
		return logrus.TraceLevel
	default:
		return logrus.InfoLevel
	}
}
