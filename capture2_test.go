package winscreencap_test

import (
	"testing"
	"time"

	"github.com/chixm/winscreencap"
)

func TestCapture2(t *testing.T) {

	hwnd, _ := winscreencap.FindWindowByName(`game window title`)

	handler := winscreencap.CaptureHandler{}

	err := handler.StartCapture(hwnd)

	if err != nil {
		t.Error(err)
	}

	time.Sleep(time.Minute)
}
