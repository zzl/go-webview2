package wv2

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 3D6B6CF2-AFE1-44C7-A995-C65117714336
var IID_ICoreWebView2DownloadOperation = syscall.GUID{0x3D6B6CF2, 0xAFE1, 0x44C7,
	[8]byte{0xA9, 0x95, 0xC6, 0x51, 0x17, 0x71, 0x43, 0x36}}

type ICoreWebView2DownloadOperation struct {
	win32.IUnknown
}

func NewICoreWebView2DownloadOperation(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2DownloadOperation {
	if pUnk == nil {
		return nil
	}
	p := (*ICoreWebView2DownloadOperation)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2DownloadOperation) IID() *syscall.GUID {
	return &IID_ICoreWebView2DownloadOperation
}

func (this *ICoreWebView2DownloadOperation) Add_BytesReceivedChanged(eventHandler *ICoreWebView2BytesReceivedChangedEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2DownloadOperation) Remove_BytesReceivedChanged(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

func (this *ICoreWebView2DownloadOperation) Add_EstimatedEndTimeChanged(eventHandler *ICoreWebView2EstimatedEndTimeChangedEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2DownloadOperation) Remove_EstimatedEndTimeChanged(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

func (this *ICoreWebView2DownloadOperation) Add_StateChanged(eventHandler *ICoreWebView2StateChangedEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2DownloadOperation) Remove_StateChanged(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

func (this *ICoreWebView2DownloadOperation) GetUri(uri *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[9]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)))
	return com.Error(ret)
}

func (this *ICoreWebView2DownloadOperation) GetContentDisposition(contentDisposition *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[10]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(contentDisposition)))
	return com.Error(ret)
}

func (this *ICoreWebView2DownloadOperation) GetMimeType(mimeType *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[11]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(mimeType)))
	return com.Error(ret)
}

func (this *ICoreWebView2DownloadOperation) GetTotalBytesToReceive(totalBytesToReceive *int64) com.Error {
	addr := (*this.LpVtbl)[12]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(totalBytesToReceive)))
	return com.Error(ret)
}

func (this *ICoreWebView2DownloadOperation) GetBytesReceived(bytesReceived *int64) com.Error {
	addr := (*this.LpVtbl)[13]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(bytesReceived)))
	return com.Error(ret)
}

func (this *ICoreWebView2DownloadOperation) GetEstimatedEndTime(estimatedEndTime *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[14]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(estimatedEndTime)))
	return com.Error(ret)
}

func (this *ICoreWebView2DownloadOperation) GetResultFilePath(resultFilePath *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[15]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(resultFilePath)))
	return com.Error(ret)
}

func (this *ICoreWebView2DownloadOperation) GetState(downloadState *int32) com.Error {
	addr := (*this.LpVtbl)[16]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(downloadState)))
	return com.Error(ret)
}

func (this *ICoreWebView2DownloadOperation) GetInterruptReason(interruptReason *int32) com.Error {
	addr := (*this.LpVtbl)[17]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(interruptReason)))
	return com.Error(ret)
}

func (this *ICoreWebView2DownloadOperation) Cancel() com.Error {
	addr := (*this.LpVtbl)[18]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return com.Error(ret)
}

func (this *ICoreWebView2DownloadOperation) Pause() com.Error {
	addr := (*this.LpVtbl)[19]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return com.Error(ret)
}

func (this *ICoreWebView2DownloadOperation) Resume() com.Error {
	addr := (*this.LpVtbl)[20]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return com.Error(ret)
}

func (this *ICoreWebView2DownloadOperation) GetCanResume(canResume *int32) com.Error {
	addr := (*this.LpVtbl)[21]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(canResume)))
	return com.Error(ret)
}

