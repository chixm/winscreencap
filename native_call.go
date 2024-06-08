package winscreencap

import (
	"syscall"
	"unsafe"

	"github.com/lxn/win"
)

// github.com/lxn/win does not include EnumWindows
// EnumWindows can list up all the windows
// see page below for details
// https://learn.microsoft.com/ja-jp/windows/win32/api/winuser/nf-winuser-enumwindows

var (
	modUser32                    = syscall.NewLazyDLL("user32.dll")
	procEnumWindows              = modUser32.NewProc("EnumWindows")
	procGetWindowTextW           = modUser32.NewProc("GetWindowTextW")
	procGetWindowTextLengthW     = modUser32.NewProc("GetWindowTextLengthW")
	procGetWindowDisplayAffinity = modUser32.NewProc("GetWindowDisplayAffinity")
	procPrintWindow              = modUser32.NewProc("PrintWindow")
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

func enumWindowsProc(hwnd win.HWND, lParam uintptr) uintptr {
	if win.IsWindowVisible(hwnd) {
		title, err := getWindowText(hwnd)
		if err != nil {
			//logger.Errorln(`Error on getting Window Title`, err)
			return 1
		}
		list = append(list, &WindowInfo{
			Hwnd:  hwnd,
			Title: title,
		})
	}
	return 1 // Continue enumeration
}

func getWindowText(hwnd win.HWND) (string, error) {
	length, _, err := procGetWindowTextLengthW.Call(uintptr(hwnd))
	if length == 0 {
		return "", err
	}
	buf := make([]uint16, length+1)
	_, _, err = procGetWindowTextW.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(len(buf)),
	)
	if err != nil && err.Error() != "The operation completed successfully." {
		return "", err
	}
	return syscall.UTF16ToString(buf), nil
}

const (
	WDA_NONE               = 0x00000000
	WDA_MONITOR            = 0x00000001
	WDA_EXCLUDEFROMCAPTURE = 0x00000011
)

// Some Windows are Disabled to Capture
func GetWindowDisplayAffinity(hwnd win.HWND) (uint32, error) {
	var affinity uint32
	ret, _, err := procGetWindowDisplayAffinity.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&affinity)),
	)
	if ret == 0 {
		return 0, err
	}
	return affinity, nil
}

func printWindow(hwnd win.HWND, hdc win.HDC) error {
	var PW_RENDERERFULLCONTEXT uint32 = 2
	ret, _, err := procPrintWindow.Call(
		uintptr(hwnd),
		uintptr(hdc),
		uintptr(PW_RENDERERFULLCONTEXT),
	)
	if ret == 0 {
		return err
	}
	return nil
}
