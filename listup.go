package winscreencap

import (
	"sync"

	"github.com/lxn/win"
)

type WindowInfo struct {
	Hwnd  win.HWND
	Title string
}

// mutex for goroutine safe
var m = sync.Mutex{}

var list []*WindowInfo

func enumWindowsProc() {
	// callback function
	EnumWindows(EnumWindowsProc, 0)
}

func WindowList() []*WindowInfo {
	m.Lock()
	defer m.Unlock()
	list = make([]*WindowInfo, 0)
	enumWindowsProc()
	return list
}
