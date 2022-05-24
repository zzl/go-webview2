package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// E9632730-6E1E-43AB-B7B8-7B2C9E62E094
var IID_ICoreWebView2_8 = syscall.GUID{0xE9632730, 0x6E1E, 0x43AB, 
	[8]byte{0xB7, 0xB8, 0x7B, 0x2C, 0x9E, 0x62, 0xE0, 0x94}}

type ICoreWebView2_8 struct {
	ICoreWebView2_7
}

func NewICoreWebView2_8(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2_8 {
	 if pUnk == nil {
		return nil;
	}
	p := (*ICoreWebView2_8)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2_8) IID() *syscall.GUID {
	return &IID_ICoreWebView2_8
}

func (this *ICoreWebView2_8) Add_IsMutedChanged(eventHandler *ICoreWebView2IsMutedChangedEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[81]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2_8) Remove_IsMutedChanged(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[82]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

func (this *ICoreWebView2_8) GetIsMuted(value *int32) com.Error {
	addr := (*this.LpVtbl)[83]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2_8) SetIsMuted(value int32) com.Error {
	addr := (*this.LpVtbl)[84]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(value))
	return com.Error(ret)
}

func (this *ICoreWebView2_8) Add_IsDocumentPlayingAudioChanged(eventHandler *ICoreWebView2IsDocumentPlayingAudioChangedEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[85]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2_8) Remove_IsDocumentPlayingAudioChanged(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[86]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

func (this *ICoreWebView2_8) GetIsDocumentPlayingAudio(value *int32) com.Error {
	addr := (*this.LpVtbl)[87]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

