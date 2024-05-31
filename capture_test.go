package winscreencap_test

import (
	"fmt"
	"image/png"
	"os"
	"testing"

	"github.com/chixm/winscreencap"
)

func TestDesktopCapture(t *testing.T) {
	// Capture DesktopWindow and save it as screenshot.png
	hwnd, err := winscreencap.GetDesktopWindow()
	if err != nil {
		t.Error(err)
		return
	}

	img, err := winscreencap.CaptureWindow(hwnd, 0)
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

	hwnd, err := winscreencap.FindWindowByName(`game`)
	if err != nil {
		t.Error(err)
		return
	}

	img, err := winscreencap.CaptureWindow(hwnd)
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
	hwnd, err := winscreencap.GetActiveWindow()
	if err != nil {
		t.Error(err) // when debug, No window should be active
		return
	}

	img, err := winscreencap.CaptureWindow(hwnd, 0)
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

func TestWindowCaptureByName(t *testing.T) {

	hwnd, err := winscreencap.FindWindowByName(`game window title`)
	if err != nil {
		t.Error(err)
		return
	}

	img, err := winscreencap.CaptureWindow(hwnd)
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
