package main

import (
	"fmt"
	"github.com/JamesHovious/w32"
	"log"
	"syscall"
	"time"
)

func main() {

	var processInput string
	var dllInput string

	fmt.Print("Enter process name e.g notepad.exe: ")
	fmt.Scanln(&processInput)
	fmt.Print("Enter dll path e.g C:\\Users\\test.dll: ")
	fmt.Scanln(&dllInput)

	processId := 0

	for {
		processId = getProcessId(processInput)
		if processId != 0 {
			fmt.Println("Found process with pid:", processId)
			break
		}
		time.Sleep(2 * time.Second)
	}

	var dllPath = dllInput

	handleProcess := openProcess(processId)
	dwSize := len(dllPath)

	loc, err := w32.VirtualAllocEx(handleProcess, 0, dwSize, w32.MEM_RESERVE|w32.MEM_COMMIT, w32.PAGE_EXECUTE_READWRITE)
	if err == nil {

		err := w32.WriteProcessMemory(handleProcess, uint32(loc), []byte(dllPath), uint(dwSize))

		test, _ := w32.ReadProcessMemory(handleProcess, uint32(loc), uint(dwSize))
		if string(test) != dllPath {
			fmt.Println(string(test))
			fmt.Println("DLL Path not found in target process memory")
		}

		moduleKernel, _ := syscall.LoadLibrary("kernel32.dll")
		lpLoadLibrary, err := syscall.GetProcAddress(moduleKernel, "LoadLibraryA")

		if err != nil {
			log.Panic(err)
		}

		handleThread, _, _ := w32.CreateRemoteThread(handleProcess, nil, 0, uint32(lpLoadLibrary), loc, 0)
		time.Sleep(5 * time.Second)

		w32.CloseHandle(handleThread)
	}

}
