package winscreencap_test

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"

	"github.com/chixm/winscreencap"
	"github.com/lxn/win"
)

func TestD3DCapture(t *testing.T) {
	hwnd := win.FindWindow(nil, win.StringToBSTR(`Prime Video for Windows`))

	var handler = winscreencap.CaptureHandler{}

	err := handler.StartCapture(hwnd)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer handler.Close()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
