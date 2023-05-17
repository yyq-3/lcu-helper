package util

import (
	"syscall"
	"unsafe"
)

/**
 * @Author Yongqi.Yang
 * @Date $ $
 **/

type PROCESSENTRY32 struct {
	dwSize              uint32    // 结构大小
	cntUsage            uint32    // 此进程的引用计数
	th32ProcessID       uint32    // 进程id
	th32DefaultHeapID   uintptr   // 进程默认堆id
	th32ModuleID        uint32    // 进程模块id
	cntThreads          uint32    // 进程的线程数
	th32ParentProcessID uint32    // 父进程id
	pcPriClassBase      uint32    // 线程优先权
	dwFlags             uint32    // 保留
	szExeFile           [260]byte // 进程全名
}

func ProcessIsRun(processName string) bool {
	return GetProcessList(processName)
	//processList, _ := process.Processes()
	//for _, proc := range processList {
	//	name, _ := proc.Name()
	//	if name == processName {
	//		return true
	//	}
	//}
	//return false
}

func GetProcessList(processName string) bool {
	/*
	   CreateToolhelp32Snapshot
	       指定快照中包含的系统内容，这个参数能够使用下列数值（常量）中的一个或多个。
	       TH32CS_INHERIT(0x80000000)      - 声明快照句柄是可继承的。
	       TH32CS_SNAPALL                  - 在快照中包含系统中所有的进程和线程。
	       TH32CS_SNAPHEAPLIST(0x00000001) - 在快照中包含在th32ProcessID中指定的进程的所有的堆。
	       TH32CS_SNAPMODULE(0x00000008)   - 在快照中包含在th32ProcessID中指定的进程的所有的模块。
	       TH32CS_SNAPPROCESS(0x00000002)  - 在快照中包含系统中所有的进程。
	       TH32CS_SNAPTHREAD(0x00000004)   - 在快照中包含系统中所有的线程。
	       H32CS_SNAPALL = (TH32CS_SNAPHEAPLIST | TH32CS_SNAPPROCESS | TH32CS_SNAPTHREAD | TH32CS_SNAPMODULE)
	   th32ProcessID
	       指定将要快照的进程ID。如果该参数为0表示快照当前进程。该参数只有在设置了TH32CS_SNAPHEAPLIST或者TH32CS_SNAPMODULE后才有效，在其他情况下该参数被忽略，所有的进程都会被快照。
	*/
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	CreateToolhelp32Snapshot := kernel32.NewProc("CreateToolhelp32Snapshot")
	pHandle, _, _ := CreateToolhelp32Snapshot.Call(uintptr(0x2), uintptr(0x0))
	CloseHandle := kernel32.NewProc("CloseHandle")
	defer CloseHandle.Call(pHandle)
	if int(pHandle) == -1 {
		return false
	}

	Process32Next := kernel32.NewProc("Process32Next")
	for {
		var proc PROCESSENTRY32
		proc.dwSize = uint32(unsafe.Sizeof(proc))
		if rt, _, _ := Process32Next.Call(pHandle, uintptr(unsafe.Pointer(&proc))); int(rt) == 1 {
			var temp []byte
			for _, b := range proc.szExeFile {
				if b != 0 {
					temp = append(temp, b)
				}
			}
			if processName == Byte2str(temp) {
				return true
			}
		} else {
			break
		}
	}
	return false
}
