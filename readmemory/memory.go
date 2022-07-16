package readmemory

import (
	"encoding/binary"
	"path/filepath"
	"unsafe"

	"github.com/0xrawsec/golang-win32/win32"
	kernel32 "github.com/0xrawsec/golang-win32/win32/kernel32"
	windows "golang.org/x/sys/windows"
)

func memoryReadInit(pid uint32, targetModuleFilename string) (int64, bool) {
	win32handle, _ := kernel32.OpenProcess(0x0010|windows.PROCESS_VM_READ|windows.PROCESS_QUERY_INFORMATION, win32.BOOL(0), win32.DWORD(pid))
	moduleHandles, _ := kernel32.EnumProcessModules(win32handle)
	for _, moduleHandle := range moduleHandles {
		s, _ := kernel32.GetModuleFilenameExW(win32handle, moduleHandle)
		if filepath.Base(s) == targetModuleFilename {
			info, _ := kernel32.GetModuleInformation(win32handle, moduleHandle)
			baseAddress := int64(info.LpBaseOfDll)
			return baseAddress, true
		}
	}
	return 0, false
}

func (handle *Handle) ReadMemoryAtByte8(address int64) uint64 {
	var (
		data   [8]byte
		length uint32
	)

	handle.procReadProcessMemory.Call(
		uintptr(handle.processHandle),
		uintptr(handle.baseAddress+address),
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(len(data)),
		uintptr(unsafe.Pointer(&length)),
	)

	byte8 := binary.LittleEndian.Uint64(data[:])
	return byte8
}

func (handle *Handle) ReadMemoryAtByte1(address int64) byte {
	var (
		data   [1]byte
		length uint32
	)

	handle.procReadProcessMemory.Call(
		uintptr(handle.processHandle),
		uintptr(handle.baseAddress+address),
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(len(data)),
		uintptr(unsafe.Pointer(&length)),
	)

	return data[0]
}

func (handle *Handle) ReadMemoryAtByte2(address int64) uint16 {
	var (
		data   [2]byte
		length uint32
	)

	handle.procReadProcessMemory.Call(
		uintptr(handle.processHandle),
		uintptr(handle.baseAddress+address),
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(len(data)),
		uintptr(unsafe.Pointer(&length)),
	)

	return binary.LittleEndian.Uint16(data[:])
}
