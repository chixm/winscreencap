package winscreencap

import (
	"fmt"
	"image"
	"image/color"

	"github.com/kirides/go-d3d"
	"github.com/kirides/go-d3d/d3d11"
)

func captureFrameToImage(frame *d3d11.ID3D11Texture2D) (*image.RGBA, error) {
	// Get description of the frame
	var desc d3d11.D3D11_TEXTURE2D_DESC
	frame.GetDesc(&desc)

	// Create a staging texture to read the data
	desc.Usage = d3d11.D3D11_USAGE_STAGING
	desc.BindFlags = 0
	desc.CPUAccessFlags = d3d11.D3D11_CPU_ACCESS_READ

	// Assuming you have a Direct3D device and context initialized
	var device *d3d11.ID3D11Device
	var context *ID3D11DeviceContext2
	var stagingTexture *d3d11.ID3D11Texture2D

	hr := device.CreateTexture2D(&desc, &stagingTexture)
	if d3d.HRESULT(hr).Failed() {
		return nil, fmt.Errorf("failed to create staging texture")
	}
	defer stagingTexture.Release()

	context.CopyResource2D(stagingTexture, frame)

	// Map the staging texture
	var mappedResource D3D11_MAPPED_SUBRESOURCE
	hr = context.Map(stagingTexture, 0, D3D11_MAP_READ, 0, &mappedResource) // Map method exists in D3D11DeviceContext but not implemented yet
	if d3d.HRESULT(hr).Failed() {
		return nil, fmt.Errorf("failed to map staging texture")
	}
	defer context.Unmap(stagingTexture, 0)

	// Create an image.RGBA and copy the data
	img := image.NewRGBA(image.Rect(0, 0, int(desc.Width), int(desc.Height)))
	for y := 0; y < int(desc.Height); y++ {
		row := mappedResource.PData[y*int(mappedResource.RowPitch) : (y+1)*int(mappedResource.RowPitch)]
		for x := 0; x < int(desc.Width); x++ {
			offset := x * 4
			r := row[offset+2]
			g := row[offset+1]
			b := row[offset+0]
			a := row[offset+3]
			img.SetRGBA(x, y, color.RGBA{R: r, G: g, B: b, A: a})
		}
	}

	return img, nil
}
