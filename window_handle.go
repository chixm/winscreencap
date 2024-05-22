package winscreencap

import (
	"errors"
	"fmt"
	"syscall"

	"github.com/lxn/win"
)

// functions to Find Windows to Capture

func FindWindowByName(windowName string) (win.HWND, error) {
	str, err := syscall.UTF16PtrFromString(windowName)
	if err != nil {
		return 0, err
	}
	hwnd := win.FindWindow(nil, str)
	if hwnd == 0 {
		return 0, fmt.Errorf(`window named %s not found`, windowName)
	}
	return hwnd, nil
}

func GetActiveWindow() (win.HWND, error) {
	hwnd := win.GetActiveWindow()
	if hwnd == 0 {
		return 0, errors.New(`no active window found`)
	}
	return hwnd, nil
}

// for desktop image capture
func GetDesktopWindow() (win.HWND, error) {
	hwnd := win.GetDesktopWindow()
	if hwnd == 0 {
		return 0, errors.New(`failed to get window handle of desktop`)
	}
	return hwnd, nil
}
