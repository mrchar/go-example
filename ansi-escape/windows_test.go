// Package main 该文件演示如何在windows下使能ansi打印
package main

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"syscall"
	"testing"
)

// onlyWindows 必须在windows下运行，否则会异常退出
func onlyWindows() {
	if runtime.GOOS != "windows" {
		panic(errors.New("only windows"))
	}
}

// enableAnsi 在windows下使能ansi
// 具体细节可以访问：https://docs.microsoft.com/en-us/windows/console/setconsolemode?redirectedfrom=MSDN
// 参考：https://github.com/sirupsen/logrus/issues/496
func enableAnsi() {
	onlyWindows()
	handle := syscall.Handle(os.Stdout.Fd())
	kernel32DLL := syscall.NewLazyDLL("kernel32.dll")
	setConsoleModeProc := kernel32DLL.NewProc("SetConsoleMode")
	setConsoleModeProc.Call(uintptr(handle), 0x0001|0x0002|0x0004)
}

func TestPrintln(t *testing.T) {
	// 将打印出错误的AnsiEscape
	fmt.Println("\033[1;31mRed\033[0m")
	// 使能windows的ansi控制符
	enableAnsi()
	// 将识别AnsiEscape控制符，并打印出红色字体
	fmt.Println("\033[1;31mRed\033[0m")
}
