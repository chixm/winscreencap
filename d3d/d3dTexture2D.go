package d3d

import (
	"syscall"
	"unsafe"
)

var (
	IID_ID3D11Texture2D = syscall.GUID{0x6f15AAF2, 0xd208, 0x4e89, [8]byte{0x9a, 0xc8, 0x43, 0x6e, 0x36, 0x3e, 0x32, 0x2d}}
)

type ID3D11Texture2D struct {
	lpVtbl *ID3D11Texture2DVtbl
}

type ID3D11Texture2DVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
	GetDesc        uintptr
}

type D3D11_TEXTURE2D_DESC struct {
	Width          uint32
	Height         uint32
	MipLevels      uint32
	ArraySize      uint32
	Format         uint32
	SampleDesc     DXGI_SAMPLE_DESC
	Usage          uint32
	BindFlags      uint32
	CPUAccessFlags uint32
	MiscFlags      uint32
}

type DXGI_SAMPLE_DESC struct {
	Count   uint32
	Quality uint32
}

func (obj *ID3D11Texture2D) GetDesc(desc *D3D11_TEXTURE2D_DESC) {
	syscall.SyscallN(
		obj.lpVtbl.GetDesc,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(desc)),
		0,
	)
}
