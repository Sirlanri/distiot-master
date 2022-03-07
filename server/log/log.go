package log

import (
	"os"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

//记录log
var Log = &logrus.Logger{}

func init() {
	//控制log级别
	level := logrus.DebugLevel
	Log = &logrus.Logger{
		Out:   os.Stderr,
		Level: level,
		Formatter: &prefixed.TextFormatter{
			DisableColors:   true,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
			ForceFormatting: true,
		},
	}
}
