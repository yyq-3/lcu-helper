package logger

import (
	"bytes"
	"fmt"
	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
	"io"
	"math/rand"
	"os"
	"path/filepath"
)

var myLog = new(MyLog)
var upColor int

func Initialize() {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetFormatter(myLog)
	runPath, _ := os.Executable()
	runDir := filepath.Dir(runPath)
	name, _ := filepath.Abs(runDir + "/lcu-helper/access.log")
	err := os.MkdirAll(runDir+"/lcu-helper/", 0744)
	if err != nil {
		Infof("创建日志文件失败,失败原因:%s", err.Error())
		logrus.SetOutput(colorable.NewColorableStdout())
		return
	}
	file, err := os.OpenFile(name, os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		Infof("打开日志文件失败,失败原因:%s", err.Error())
		logrus.SetOutput(colorable.NewColorableStdout())
		return
	} else {
		Infof("日志文件路径: %S", name)
	}
	logrus.SetOutput(io.MultiWriter(colorable.NewColorableStdout(), colorable.NewColorable(file)))
}

type MyLog struct {
}

func (mLog *MyLog) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006/01/02 15:04:05.000")
	msg := fmt.Sprintf("%s %s", timestamp, entry.Message)
	_, err := fmt.Fprintf(b, colorOut(msg))
	return b.Bytes(), err
}

func colorOut(msg string) string {
	// 随机展示颜色
	color := 0
	for {
		color = rand.Intn(7) + 31
		if color != upColor {
			upColor = color
			break
		}
	}
	return fmt.Sprintf("\u001B[%dm%s\u001B[0m\n", color, msg)
}

func Info(args ...interface{}) {
	logrus.Infoln(args)
}

func Infof(str string, args ...interface{}) {
	logrus.Infof(str, args...)
}
