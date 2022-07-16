package readmemory

import (
	windows "golang.org/x/sys/windows"
)

type Handle struct {
	processHandle         windows.Handle
	baseAddress           int64
	procReadProcessMemory *windows.Proc
}

func NewHandle(processName string) Handle {
	var handle Handle

	pid, _ := bindDefaultProcess(processName)
	handle.processHandle, _ = windows.OpenProcess(0x0010|windows.PROCESS_VM_READ|windows.PROCESS_QUERY_INFORMATION, false, pid)
	handle.procReadProcessMemory = windows.MustLoadDLL("kernel32.dll").MustFindProc("ReadProcessMemory")
	baseAddress, _ := memoryReadInit(pid, processName)
	handle.baseAddress = baseAddress

	return handle
}
