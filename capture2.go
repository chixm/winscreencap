package winscreencap

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"syscall"
	"time"

	"github.com/TKMAX777/winapi/dx11"
	"github.com/TKMAX777/winapi/winrt"

	"golang.org/x/sys/windows"
)

var (
	d3d11 = syscall.NewLazyDLL("d3d11.dll")
	dxgi  = syscall.NewLazyDLL("dxgi.dll")
)

func main() {
	// Get handle to the desktop window
	hwnd, _ := GetDesktopWindow()
	if hwnd == 0 {
		fmt.Println("Failed to get handle to the desktop window")
		return
	}

	// Initialize DirectX
	device, context, err := initD3D11()
	if err != nil {
		fmt.Println("Failed to initialize DirectX:", err)
		return
	}
	defer device.Release()
	defer context.Release()

	// Create DXGI output duplicator
	duplicator, err := createDuplicator(device)
	if err != nil {
		fmt.Println("Failed to create DXGI output duplicator:", err)
		return
	}
	defer duplicator.Release()

	// Capture frames
	for i := 0; i < 1; i++ { // Capture only one frame for example
		frame, err := captureFrame(duplicator, context)
		if err != nil {
			fmt.Println("Failed to capture frame:", err)
			return
		}

		// Process frame and save to PNG
		err = processFrame(frame, i)
		if err != nil {
			fmt.Println("Failed to process frame:", err)
			return
		}

		time.Sleep(time.Second / 30)
	}
}

// キャプチャ用デバイスを返す
func initD3D11() (*winrt.IDirect3DDevice, ID3D11DeviceContext, error) {
	var device *winrt.IDirect3DDevice
	var context *winrt.IDirect3DDeviceVtbl

	var featureLevels = []dx11.D3D_FEATURE_LEVEL{
		dx11.D3D_FEATURE_LEVEL_11_0,
		dx11.D3D_FEATURE_LEVEL_10_1,
		dx11.D3D_FEATURE_LEVEL_10_0,
		dx11.D3D_FEATURE_LEVEL_9_3,
		dx11.D3D_FEATURE_LEVEL_9_2,
		dx11.D3D_FEATURE_LEVEL_9_1,
	}

	err := dx11.D3D11CreateDevice(
		nil, dx11.D3D_DRIVER_TYPE_HARDWARE, 0, dx11.D3D11_CREATE_DEVICE_BGRA_SUPPORT|dx11.D3D11_CREATE_DEVICE_DEBUG,
		&featureLevels[0], len(featureLevels),
		dx11.D3D11_SDK_VERSION, &c.deviceDx, nil, nil,
	)
	if err != nil {
		return nil, nil, err
	}

	return device, context, nil
}

func createDuplicator(device *windows.ID3D11Device) (*windows.IDXGIOutputDuplication, error) {
	var adapter *windows.IDXGIAdapter
	hr := device.QueryInterface(&windows.IID_IDXGIDevice, &adapter)
	if windows.FAILED(hr) {
		return nil, fmt.Errorf("QueryInterface for IDXGIDevice failed: %v", hr)
	}
	defer adapter.Release()

	var output *windows.IDXGIOutput
	hr = adapter.EnumOutputs(0, &output)
	if windows.FAILED(hr) {
		return nil, fmt.Errorf("EnumOutputs failed: %v", hr)
	}
	defer output.Release()

	var output1 *windows.IDXGIOutput1
	hr = output.QueryInterface(&windows.IID_IDXGIOutput1, &output1)
	if windows.FAILED(hr) {
		return nil, fmt.Errorf("QueryInterface for IDXGIOutput1 failed: %v", hr)
	}
	defer output1.Release()

	var duplicator *windows.IDXGIOutputDuplication
	hr = output1.DuplicateOutput(device, &duplicator)
	if windows.FAILED(hr) {
		return nil, fmt.Errorf("DuplicateOutput failed: %v", hr)
	}

	return duplicator, nil
}

func captureFrame(duplicator *windows.IDXGIOutputDuplication, context *windows.ID3D11DeviceContext) (*windows.ID3D11Texture2D, error) {
	var frameInfo windows.DXGI_OUTDUPL_FRAME_INFO
	var resource *windows.IDXGIResource

	hr := duplicator.AcquireNextFrame(0, &frameInfo, &resource)
	if hr == windows.DXGI_ERROR_WAIT_TIMEOUT {
		return nil, nil // No new frame available
	} else if windows.FAILED(hr) {
		return nil, fmt.Errorf("AcquireNextFrame failed: %v", hr)
	}
	defer resource.Release()

	// Get the IDXGISurface interface for the captured frame
	var texture *windows.ID3D11Texture2D
	hr = resource.QueryInterface(&windows.IID_ID3D11Texture2D, &texture)
	if windows.FAILED(hr) {
		return nil, fmt.Errorf("QueryInterface for ID3D11Texture2D failed: %v", hr)
	}

	return texture, nil
}

func processFrame(frame *winrt.ID3D11Texture2D, index int) error {
	// Map the frame to CPU memory
	var mappedResource windows.D3D11_MAPPED_SUBRESOURCE
	context.Map(frame, 0, windows.D3D11_MAP_READ, 0, &mappedResource)
	defer context.Unmap(frame, 0)

	// Create an image from the mapped resource
	img := image.NewRGBA(image.Rect(0, 0, int(frameDesc.Width), int(frameDesc.Height)))
	for y := 0; y < int(frameDesc.Height); y++ {
		for x := 0; x < int(frameDesc.Width); x++ {
			offset := y*int(mappedResource.RowPitch) + x*4
			b := mappedResource.Data[offset+0]
			g := mappedResource.Data[offset+1]
			r := mappedResource.Data[offset+2]
			a := mappedResource.Data[offset+3]
			img.SetRGBA(x, y, color.RGBA{R: r, G: g, B: b, A: a})
		}
	}

	// Save the image to a PNG file
	file, err := os.Create(fmt.Sprintf("frame_%d.png", index))
	if err != nil {
		return fmt.Errorf("failed to create PNG file: %v", err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		return fmt.Errorf("failed to encode PNG: %v", err)
	}

	return nil
}
