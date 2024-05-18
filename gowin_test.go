package winscreencap_test

import (
	"fmt"
	"image/png"
	"os"
	"testing"

	gowin "github.com/chixm/winscreencap"
	"github.com/lxn/win"
)

func TestDesktopCapture(t *testing.T) {
	// Capture DesktopWindow and save it as screenshot.png
	hwnd := win.GetDesktopWindow()

	img, err := gowin.CaptureWindow(hwnd, 0)
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

func TestWindowCapture(t *testing.T) {
	// Capture DesktopWindow and save it as screenshot.png
	hwnd, err := gowin.FindWindowByName(`game`)
	if err != nil {
		t.Error(err)
		return
	}

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

func TestActiveWindowCapture(t *testing.T) {
	// Capture DesktopWindow and save it as screenshot.png
	hwnd, err := gowin.GetActiveWindow()
	if err != nil {
		t.Error(err) // when debug, No window should be active
		return
	}

	img, err := gowin.CaptureWindow(hwnd, 0)
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
