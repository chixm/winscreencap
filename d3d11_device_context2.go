package winscreencap

import (
	"syscall"
	"unsafe"

	"github.com/kirides/go-d3d/d3d11"
	"github.com/kirides/go-d3d/dxgi"
)

// expanded ID3D11DeviceContext
// add Map and Unmap method
type ID3D11DeviceContext2 struct {
	vtbl *d3d11.ID3D11DeviceContextVtbl
}

func (obj *ID3D11DeviceContext2) CopyResourceDXGI(dst, src *dxgi.IDXGIResource) int32 {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.CopyResource,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src)),
	)
	return int32(ret)
}
func (obj *ID3D11DeviceContext2) CopyResource2D(dst, src *d3d11.ID3D11Texture2D) int32 {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.CopyResource,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src)),
	)
	return int32(ret)
}
func (obj *ID3D11DeviceContext2) CopySubresourceRegion2D(dst *d3d11.ID3D11Texture2D, dstSubResource, dstX, dstY, dstZ uint32, src *d3d11.ID3D11Texture2D, srcSubResource uint32, pSrcBox *d3d11.D3D11_BOX) int32 {
	ret, _, _ := syscall.Syscall9(
		obj.vtbl.CopySubresourceRegion,
		9,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(dst)),
		uintptr(dstSubResource),
		uintptr(dstX),
		uintptr(dstY),
		uintptr(dstZ),
		uintptr(unsafe.Pointer(src)),
		uintptr(srcSubResource),
		uintptr(unsafe.Pointer(pSrcBox)),
	)
	return int32(ret)
}

func (obj *ID3D11DeviceContext2) CopySubresourceRegion(dst *d3d11.ID3D11Resource, dstSubResource, dstX, dstY, dstZ uint32, src *d3d11.ID3D11Resource, srcSubResource uint32, pSrcBox *d3d11.D3D11_BOX) int32 {
	ret, _, _ := syscall.Syscall9(
		obj.vtbl.CopySubresourceRegion,
		9,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(dst)),
		uintptr(dstSubResource),
		uintptr(dstX),
		uintptr(dstY),
		uintptr(dstZ),
		uintptr(unsafe.Pointer(src)),
		uintptr(srcSubResource),
		uintptr(unsafe.Pointer(pSrcBox)),
	)
	return int32(ret)
}

func (obj *ID3D11DeviceContext2) Map(stagingTexture *d3d11.ID3D11Texture2D, subResource int32, mapType int32, mapFlag int32, mappedSubresource *D3D11_MAPPED_SUBRESOURCE) int32 {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.Map,
		uintptr(unsafe.Pointer(stagingTexture)),
		uintptr(subResource),
		uintptr(mapType),
		uintptr(mapFlag),
		uintptr(unsafe.Pointer(mappedSubresource)),
	)
	return int32(ret)
}

func (obj *ID3D11DeviceContext2) Unmap(stagingTexture *d3d11.ID3D11Texture2D, subResource int32) int32 {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.Unmap,
		uintptr(unsafe.Pointer(stagingTexture)),
		uintptr(subResource),
	)
	return int32(ret)
}

func (obj *ID3D11DeviceContext2) Release() int32 {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.Release,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return int32(ret)
}

const D3D11_MAP_READ = 1

type D3D11_MAPPED_SUBRESOURCE struct {
	PData      []byte
	RowPitch   uint32
	DepthPitch uint32
}
