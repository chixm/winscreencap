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

// list to put window infos in
var list []*WindowInfo

// get all visible window list
func GetWindowList() []*WindowInfo {
	m.Lock()
	defer m.Unlock()
	list = make([]*WindowInfo, 0)
	// callback function
	EnumWindows(enumWindowsProc, 0)
	return list
}
