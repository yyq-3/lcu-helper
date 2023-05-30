package pkg

import (
	"fmt"
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

func Speak() {
	ole.CoInitialize(0)
	unknown, _ := oleutil.CreateObject("SAPI.SpVoice")
	voice, _ := unknown.QueryInterface(ole.IID_IDispatch)
	saveFile, _ := oleutil.CreateObject("SAPI.SpFileStream")
	ff, _ := saveFile.QueryInterface(ole.IID_IDispatch)
	// 打开wav文件
	//oleutil.CallMethod(ff, "Open", "D:\\mygo\\aa.wav", 3, true)
	// 设置voice的AudioOutputStream属性，必须是PutPropertyRef，如果是PutProperty就无法生效
	oleutil.PutPropertyRef(voice, "AudioOutputStream", ff)
	//设置语速
	oleutil.PutProperty(voice, "Rate", 0)
	//设置音量
	oleutil.PutProperty(voice, "Volume", 200)
	//说话
	oleutil.CallMethod(voice, "Speak", "玩家A击杀了玩家B完成了一次双杀")
	//oleutil.CallMethod(voice, "Speak", "bb", 1)
	//停止说话
	//oleutil.CallMethod(voice, "Pause")
	//恢复说话
	//oleutil.CallMethod(voice, "Resume")
	//等待结束
	oleutil.CallMethod(voice, "WaitUntilDone", 1000000)
	//关闭文件
	oleutil.CallMethod(ff, "Close")
	fmt.Printf("lalallalala\n")
	ff.Release()
	voice.Release()
	ole.CoUninitialize()
}
