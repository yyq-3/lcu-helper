package logger

import (
	"bytes"
	"fmt"
	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
	"math/rand"
	"os"
	"time"
)

var myLog = new(MyLog)
var base = 0.123456

func initialize() {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetFormatter(myLog)
	logrus.SetOutput(colorable.NewColorableStdout())
	file, err := os.OpenFile("./lcu-helper/access.log", 1, os.ModeAppend)
	if err != nil {
		return
	}
	logrus.SetOutput(colorable.NewColorable(file))
}

type MyLog struct {
	Color int
}

func (mLog *MyLog) Format(entry *logrus.Entry) ([]byte, error) {
	// 随机展示颜色
	f := base * float64(time.Now().Nanosecond())
	rand.Seed(int64(f))
	myLog.Color = rand.Intn(7) + 31
	var color = mLog.Color
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006/01/02 15:04:05.000")
	_, err := fmt.Fprintf(b, "\x1b[%dm%s %s\x1b[0m\n", color, timestamp, entry.Message)
	return b.Bytes(), err
}

func Info(args ...interface{}) {
	initialize()
	logrus.Infoln(args)
}

func Infof(str string, args ...interface{}) {
	initialize()
	logrus.Infof(str, args...)
}
