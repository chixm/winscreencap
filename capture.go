package winscreencap

import (
	"errors"
	"fmt"
	"image"
	"syscall"
	"unsafe"

	"github.com/lxn/win"
)

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

// CaptureOptions
type Options uint32

const (
	WithWindowFrame = Options(1 << iota) // 1 Include Windows Frame in png image
	NoAlphaFix                           // 2 Do not automatically fix even alpha value is zero,
)

// Capture all the window
func CaptureWindow(hwnd win.HWND, options ...Options) (image.Image, error) {
	// JOIN options
	var option Options
	for _, v := range options {
		option = option & v
	}

	var rect win.RECT
	if option&WithWindowFrame > 0 {
		win.GetWindowRect(hwnd, &rect)
	} else {
		if !win.GetClientRect(hwnd, &rect) {
			return nil, errors.New(`failed to get client area of hwnd`)
		}
	}

	width := rect.Right - rect.Left
	height := rect.Bottom - rect.Top

	hdc := win.GetDC(hwnd)
	defer win.ReleaseDC(hwnd, hdc)

	memDC := win.CreateCompatibleDC(hdc)
	defer win.DeleteDC(memDC)

	bitmap := win.CreateCompatibleBitmap(hdc, width, height)
	defer win.DeleteObject(win.HGDIOBJ(bitmap))

	win.SelectObject(memDC, win.HGDIOBJ(bitmap))
	if !win.BitBlt(memDC, 0, 0, width, height, hdc, rect.Top, rect.Left, win.SRCCOPY) {
		return nil, errors.New(`failed to BitBlt screen`)
	}

	img := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))

	bitmapInfo := win.BITMAPINFO{
		BmiHeader: win.BITMAPINFOHEADER{
			BiSize:        uint32(unsafe.Sizeof(win.BITMAPINFOHEADER{})),
			BiWidth:       int32(width),
			BiHeight:      -int32(height),
			BiPlanes:      1,
			BiBitCount:    32,
			BiCompression: win.BI_RGB,
		},
	}

	bgra := make([]byte, width*height*4)

	if win.GetDIBits(memDC, bitmap, 0, uint32(height), &bgra[0], &bitmapInfo, win.DIB_RGB_COLORS) == int32(0) {
		return nil, fmt.Errorf(`can not get DIBits of window`)
	}

	// Convert BGRA to RGBA
	for y := 0; y < int(height); y++ {
		for x := 0; x < int(width); x++ {
			i := (y*int(width) + x) * 4
			img.Pix[i+0] = bgra[i+2] // R
			img.Pix[i+1] = bgra[i+1] // G
			img.Pix[i+2] = bgra[i+0] // B
			if bgra[i+3] == 0 && (option&NoAlphaFix == 0) {
				img.Pix[i+3] = 255
			} else {
				img.Pix[i+3] = bgra[i+3] // A
			}
		}
	}
	return img, nil
}