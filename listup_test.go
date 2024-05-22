package winscreencap_test

import (
	"fmt"
	"image/png"
	"log"
	"os"
	"testing"

	"github.com/chixm/winscreencap"
	"github.com/lxn/win"
)

func TestListUpWindow(t *testing.T) {
	windowList := winscreencap.GetWindowList()
	for _, v := range windowList {
		log.Println(v.Title, v.Hwnd)
	}
}

// This test will make screenshot files for each window in your PC
// Browser window such as GoogleChrome, Edge black outs screenshot for maybe security reason.
func TestListUpAndTakeShot(t *testing.T) {
	var index int

	windowList := winscreencap.GetWindowList()
	for _, v := range windowList {
		// only visible window
		if !win.IsWindowVisible(v.Hwnd) {
			continue
		}

		flag, err := winscreencap.GetWindowDisplayAffinity(v.Hwnd)
		log.Println(`affinity flag is `, flag)
		if err != nil || flag&winscreencap.WDA_EXCLUDEFROMCAPTURE > 0 {
			continue
		}

		index++

		img, err := winscreencap.CaptureWindow(v.Hwnd)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		log.Println(v.Title+` as screenshot No`, index)

		file, err := os.Create(fmt.Sprintf("screenshot%d.png", index))
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
	}
}
