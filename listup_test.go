package winscreencap_test

import (
	"log"
	"testing"

	"github.com/chixm/winscreencap"
)

func TestListUpWindow(t *testing.T) {
	windowList := winscreencap.WindowList()

	for _, v := range windowList {
		log.Println(v.Title, v.Hwnd)
	}
}
