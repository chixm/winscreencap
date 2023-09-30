package gowin

import (
	"github.com/go-ole/go-ole"
)

const namespace = `Windows.Graphics.Capture`

const staticClass = `Direct3D11CaptureFramePool`

func CallDirect3D11CaptureFramePool() error {
	// IJsonObjectStaticsの生成(戻り値の型はIInspectable)
	inspectable, err := ole.RoGetActivationFactory(namespace+`.`+staticClass, ole.NewGUID(`{2289F159-54DE-45D8-ABCC-22603FA066A0}`))
	if err != nil {
		return err
	}
	logger.Infoln(`succeeded to get inspectible`, inspectable)
	return nil
}
