package wv2

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// F9614724-5D2B-41DC-AEF7-73D62B51543B
var IID_ICoreWebView2Controller3 = syscall.GUID{0xF9614724, 0x5D2B, 0x41DC,
	[8]byte{0xAE, 0xF7, 0x73, 0xD6, 0x2B, 0x51, 0x54, 0x3B}}

type ICoreWebView2Controller3 struct {
	ICoreWebView2Controller2
}

func NewICoreWebView2Controller3(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2Controller3 {
	if pUnk == nil {
		return nil
	}
	p := (*ICoreWebView2Controller3)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2Controller3) IID() *syscall.GUID {
	return &IID_ICoreWebView2Controller3
}

func (this *ICoreWebView2Controller3) GetRasterizationScale(scale *float64) com.Error {
	addr := (*this.LpVtbl)[28]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(scale)))
	return com.Error(ret)
}

func (this *ICoreWebView2Controller3) SetRasterizationScale(scale float64) com.Error {
	addr := (*this.LpVtbl)[29]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(scale))
	return com.Error(ret)
}

func (this *ICoreWebView2Controller3) GetShouldDetectMonitorScaleChanges(value *int32) com.Error {
	addr := (*this.LpVtbl)[30]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2Controller3) SetShouldDetectMonitorScaleChanges(value int32) com.Error {
	addr := (*this.LpVtbl)[31]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(value))
	return com.Error(ret)
}

func (this *ICoreWebView2Controller3) Add_RasterizationScaleChanged(eventHandler *ICoreWebView2RasterizationScaleChangedEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[32]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2Controller3) Remove_RasterizationScaleChanged(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[33]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

func (this *ICoreWebView2Controller3) GetBoundsMode(boundsMode *int32) com.Error {
	addr := (*this.LpVtbl)[34]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(boundsMode)))
	return com.Error(ret)
}

func (this *ICoreWebView2Controller3) SetBoundsMode(boundsMode int32) com.Error {
	addr := (*this.LpVtbl)[35]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(boundsMode))
	return com.Error(ret)
}

