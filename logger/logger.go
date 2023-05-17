package logger

import (
	"bytes"
	"fmt"
	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

var myLog = new(MyLog)
var base = 0.123456

func Init() {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetFormatter(myLog)

	logrus.SetOutput(colorable.NewColorableStdout())
}

//颜色
const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

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
	_, err := fmt.Fprintf(b, "\x1B[%dm%s %s\x1B[0m \n", color, timestamp, entry.Message)
	return b.Bytes(), err
}

func Info(args ...interface{}) {
	Init()
	logrus.Info(args)
}

func Infof(str string, args ...interface{}) {
	Init()
	logrus.Info(str, args)
}
