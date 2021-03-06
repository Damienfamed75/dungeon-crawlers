package common

import (
	"os"
	"os/exec"
	"runtime"
)

var goos = func() string { return runtime.GOOS }

// initClear is used to initialize the clear
// map of functions for clearing the console
// in the desired operating systems.
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
	clsFunc, ok := clear[goos()]
	if ok {
		clsFunc()
	} else {
		panic("Your platform isn't supported")
	}
}
