package gowin

import (
	"image"
	"unsafe"

	"github.com/lxn/win"
)

func CaptureWindow(hwnd win.HWND) (image.Image, error) {
	var rect win.RECT
	win.GetWindowRect(hwnd, &rect)

	width := rect.Right - rect.Left
	height := rect.Bottom - rect.Top

	hdc := win.GetDC(hwnd)
	defer win.ReleaseDC(hwnd, hdc)

	memDC := win.CreateCompatibleDC(hdc)
	defer win.DeleteDC(memDC)

	bitmap := win.CreateCompatibleBitmap(hdc, width, height)
	defer win.DeleteObject(win.HGDIOBJ(bitmap))

	win.SelectObject(memDC, win.HGDIOBJ(bitmap))
	win.BitBlt(memDC, 0, 0, width, height, hdc, 0, 0, win.SRCCOPY)

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

	win.GetDIBits(memDC, bitmap, 0, uint32(height), &img.Pix[0], &bitmapInfo, win.DIB_RGB_COLORS)

	return img, nil
}
