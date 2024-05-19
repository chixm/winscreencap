package winscreencap

import (
	"syscall"

	"github.com/lxn/win"
)

// github.com/lxn/win does not include EnumWindows
// EnumWindows can list up all the windows
// see page below for details
// https://learn.microsoft.com/ja-jp/windows/win32/api/winuser/nf-winuser-enumwindows

var (
	modUser32                = syscall.NewLazyDLL("user32.dll")
	procEnumWindows          = modUser32.NewProc("EnumWindows")
	procGetWindowTextW       = modUser32.NewProc("GetWindowTextW")
	procGetWindowTextLengthW = modUser32.NewProc("GetWindowTextLengthW")
)

type WNDENUMPROC func(hwnd win.HWND, lParam uintptr) uintptr

func EnumWindows(enumFunc WNDENUMPROC, lParam uintptr) error {
	ret, _, err := procEnumWindows.Call(
		syscall.NewCallback(enumFunc),
		lParam,
	)
	if ret == 0 {
		return err
	}
	return nil
}

func EnumWindowsProc(hwnd win.HWND, lParam uintptr) uintptr {
	if win.IsWindowVisible(hwnd) {
		logger.Infoln(`found visible window hwnd:`, hwnd)
	}
	return 1 // Continue enumeration
}
