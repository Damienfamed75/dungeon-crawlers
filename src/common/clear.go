package common

import (
	"os"
	"os/exec"
	"runtime"
)

func initClear() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// Clear is used for clearing the terminal
// window for your desired operating system.
func Clear() {
	if clear == nil {
		initClear()
	}

	// runtime.GOOS is to check the GO Operating System.
	clsFunc, ok := clear[runtime.GOOS]
	if ok {
		clsFunc()
	} else {
		panic("Your platform isn't supported")
	}
}
