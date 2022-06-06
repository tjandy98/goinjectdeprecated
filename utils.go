package main

import (
	"github.com/JamesHovious/w32"
	"golang.org/x/sys/windows"
	"unsafe"
)

func decodeUtf16ToString(encoded [260]uint16) string {
	end := 0
	for {
		if encoded[end] == 0 {
			break
		}
		end++
	}
	return windows.UTF16ToString(encoded[:end])
}

func getProcessId(processName string) int {
	processId := 0
	handle := w32.CreateToolhelp32Snapshot(w32.TH32CS_SNAPPROCESS, 0)
	if handle != w32.INVALID_HANDLE {
		var processEntry w32.PROCESSENTRY32
		processEntry.Size = uint32(unsafe.Sizeof(processEntry))
		if w32.Process32First(handle, &processEntry) {
			for {
				if decodeUtf16ToString(processEntry.ExeFile) == processName {
					processId = int(processEntry.ProcessID)
					break
				}
				if !w32.Process32Next(handle, &processEntry) {
					break
				}
			}
		}
	}
	w32.CloseHandle(handle)
	return processId
}

func openProcess(processId int) w32.HANDLE {
	handle, _ := w32.OpenProcess(w32.PROCESS_ALL_ACCESS, false, uint32(processId))
	return handle
}
