package admin

import (
	"golang.org/x/sys/windows"
	"lcu-helper/pkg/logger"
	"os"
	"syscall"
)

/*
	本包主要用来获取Windows操作系统管理员身份
*/

// IsAdmin 判断是否管理员运行
func IsAdmin() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	return err == nil
}

func WithAdminRun() {
	if !IsAdmin() {
		AllowAdmin()
	} else {
		logger.Initialize()
		logger.Info("已获取管理员身份")
	}
}

func AllowAdmin() {
	// 指定要进行的操作
	operation := "runas"
	operationPtr, _ := syscall.UTF16PtrFromString(operation)
	// 获取当前运行exe完整文件名
	exeFileName, _ := os.Executable()
	logger.Info(exeFileName)
	exeFileNamePtr, _ := syscall.UTF16PtrFromString(exeFileName)
	cwd, _ := os.Getwd()
	cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
	// 显式运行
	var showCmd int32 = 1
	err := windows.ShellExecute(0, operationPtr, exeFileNamePtr, nil, cwdPtr, showCmd)
	if err != nil {
		logger.Info(err)
	}
	os.Exit(-2)
}
