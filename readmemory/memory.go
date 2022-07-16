package readmemory

import (
	"encoding/binary"
	"math"
	"path/filepath"
	"strconv"
	"unsafe"

	"github.com/0xrawsec/golang-win32/win32"
	kernel32 "github.com/0xrawsec/golang-win32/win32/kernel32"
	windows "golang.org/x/sys/windows"
)

func MemoryReadInit(pid uint32, targetModuleFilename string) (int64, bool) {
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

func readMemoryAt(handle windows.Handle, procReadProcessMemory *windows.Proc, address int64) float32 {
	var (
		data   [4]byte
		length uint32
	)

	procReadProcessMemory.Call(
		uintptr(handle),
		uintptr(address),
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(len(data)),
		uintptr(unsafe.Pointer(&length)),
	)

	bits := binary.LittleEndian.Uint32(data[:])
	float := math.Float32frombits(bits)
	return float
}

func readMemoryAtByte8(handle windows.Handle, procReadProcessMemory *windows.Proc, address int64) uint64 {
	var (
		data   [8]byte
		length uint32
	)

	procReadProcessMemory.Call(
		uintptr(handle),
		uintptr(address),
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(len(data)),
		uintptr(unsafe.Pointer(&length)),
	)

	byte8 := binary.LittleEndian.Uint64(data[:])
	return byte8
}

func ReadMemoryAtByte1(handle windows.Handle, procReadProcessMemory *windows.Proc, address int64) byte {
	var (
		data   [1]byte
		length uint32
	)

	procReadProcessMemory.Call(
		uintptr(handle),
		uintptr(address),
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(len(data)),
		uintptr(unsafe.Pointer(&length)),
	)

	return data[0]
}

func ReadMemoryAtByte2(handle windows.Handle, procReadProcessMemory *windows.Proc, address int64) uint16 {
	var (
		data   [2]byte
		length uint32
	)

	procReadProcessMemory.Call(
		uintptr(handle),
		uintptr(address),
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(len(data)),
		uintptr(unsafe.Pointer(&length)),
	)

	return binary.LittleEndian.Uint16(data[:])
}

type staticPointer struct {
	baseOffset int64
	offsets    []string
}

func sumHex(aHex string, bHex string) string {
	aDecimal, _ := strconv.ParseInt(aHex, 16, 0)
	bDecimal, _ := strconv.ParseInt(bHex, 16, 0)
	resultDecimal := aDecimal + bDecimal
	resultHex := strconv.FormatInt(resultDecimal, 16)
	return resultHex
}
