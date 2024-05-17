package gowin_test

import (
	"fmt"
	"image/png"
	"os"
	"testing"

	"github.com/chixm/gowin"
	"github.com/lxn/win"
)

func TestWindowCapture(t *testing.T) {
	// Capture DesktopWindow and save it as screenshot.png
	hwnd := win.GetDesktopWindow()

	img, err := gowin.CaptureWindow(hwnd)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	file, err := os.Create("screenshot.png")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Screenshot saved to screenshot.png")
}
