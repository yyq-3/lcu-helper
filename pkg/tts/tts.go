package tts

import (
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	"lcu-helper/pkg/logger"
)

var voice *ole.IDispatch
var ff *ole.IDispatch

func Init() {
	err := ole.CoInitialize(0)
	if err != nil {
		logger.Infof("语音助手初始化失败，失败原因:%s", err.Error())
	}
	unknown, _ := oleutil.CreateObject("SAPI.SpVoice")
	voice, _ = unknown.QueryInterface(ole.IID_IDispatch)
	saveFile, _ := oleutil.CreateObject("SAPI.SpFileStream")
	ff, _ = saveFile.QueryInterface(ole.IID_IDispatch)
}

func Speak(body string) {
	// 打开wav文件
	//oleutil.CallMethod(ff, "Open", "D:\\mygo\\aa.wav", 3, true)
	// 设置voice的AudioOutputStream属性，必须是PutPropertyRef，如果是PutProperty就无法生效
	oleutil.PutPropertyRef(voice, "AudioOutputStream", ff)
	//设置语速
	oleutil.PutProperty(voice, "Rate", 0)
	//设置音量
	oleutil.PutProperty(voice, "Volume", 200)
	//说话
	oleutil.CallMethod(voice, "Speak", body)
	//停止说话
	//oleutil.CallMethod(voice, "Pause")
	//恢复说话
	//oleutil.CallMethod(voice, "Resume")
	//等待结束
	oleutil.CallMethod(voice, "WaitUntilDone", 1000000)

}

// Exit 程序退出关闭流
func Exit() {
	//关闭文件
	_, err := oleutil.CallMethod(ff, "Close")
	if err != nil {
		logger.Infof("语音助手退出失败，失败原因：%s", err.Error())
	}
	ff.Release()
	voice.Release()
	ole.CoUninitialize()
}
