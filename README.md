# winscreencap
Golang Windows Screen Capture Library


## Install
```
go get github.com/chixm/winscreencap
```

## Choose Window and Capture Image
How to save the window named "game" as an png image screenshot.png

```
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
```