package wv2

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// D7175BA2-BCC3-11EB-8529-0242AC130003
var IID_ICoreWebView2ClientCertificateRequestedEventHandler = syscall.GUID{0xD7175BA2, 0xBCC3, 0x11EB,
	[8]byte{0x85, 0x29, 0x02, 0x42, 0xAC, 0x13, 0x00, 0x03}}

type ICoreWebView2ClientCertificateRequestedEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2ClientCertificateRequestedEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2, args *ICoreWebView2ClientCertificateRequestedEventArgs) com.Error
}

type ICoreWebView2ClientCertificateRequestedEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2ClientCertificateRequestedEventHandlerInterface
}

func (this *ICoreWebView2ClientCertificateRequestedEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2ClientCertificateRequestedEventHandlerInterface)
}

func (this *ICoreWebView2ClientCertificateRequestedEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2ClientCertificateRequestedEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2ClientCertificateRequestedEventHandlerImpl) Invoke(sender *ICoreWebView2, args *ICoreWebView2ClientCertificateRequestedEventArgs) com.Error {
	var ret com.Error
	return ret
}

type ICoreWebView2ClientCertificateRequestedEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2ClientCertificateRequestedEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2ClientCertificateRequestedEventHandlerComObj) impl() ICoreWebView2ClientCertificateRequestedEventHandlerInterface {
	return this.Impl().(ICoreWebView2ClientCertificateRequestedEventHandlerInterface)
}

func (this *ICoreWebView2ClientCertificateRequestedEventHandlerComObj) Invoke(sender *ICoreWebView2, args *ICoreWebView2ClientCertificateRequestedEventArgs) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2ClientCertificateRequestedEventHandlerVtbl *ICoreWebView2ClientCertificateRequestedEventHandlerVtbl

func (this *ICoreWebView2ClientCertificateRequestedEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2ClientCertificateRequestedEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
	}
	if _pICoreWebView2ClientCertificateRequestedEventHandlerVtbl != nil {
		return _pICoreWebView2ClientCertificateRequestedEventHandlerVtbl
	}
	_pICoreWebView2ClientCertificateRequestedEventHandlerVtbl = &ICoreWebView2ClientCertificateRequestedEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:       syscall.NewCallback((*ICoreWebView2ClientCertificateRequestedEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2ClientCertificateRequestedEventHandlerVtbl
}

func (this *ICoreWebView2ClientCertificateRequestedEventHandlerComObj) ICoreWebView2ClientCertificateRequestedEventHandler() *ICoreWebView2ClientCertificateRequestedEventHandler {
	return (*ICoreWebView2ClientCertificateRequestedEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2ClientCertificateRequestedEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2ClientCertificateRequestedEventHandlerComObj(impl ICoreWebView2ClientCertificateRequestedEventHandlerInterface, scoped bool) *ICoreWebView2ClientCertificateRequestedEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2ClientCertificateRequestedEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2ClientCertificateRequestedEventHandler(impl ICoreWebView2ClientCertificateRequestedEventHandlerInterface) *ICoreWebView2ClientCertificateRequestedEventHandler {
	return NewICoreWebView2ClientCertificateRequestedEventHandlerComObj(impl, true).ICoreWebView2ClientCertificateRequestedEventHandler()
}

//
type ICoreWebView2ClientCertificateRequestedEventHandlerByFuncImpl struct {
	ICoreWebView2ClientCertificateRequestedEventHandlerImpl
	handlerFunc func(sender *ICoreWebView2, args *ICoreWebView2ClientCertificateRequestedEventArgs) com.Error
}

func (this *ICoreWebView2ClientCertificateRequestedEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2, args *ICoreWebView2ClientCertificateRequestedEventArgs) com.Error {
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2ClientCertificateRequestedEventHandlerByFunc(handlerFunc func(sender *ICoreWebView2, args *ICoreWebView2ClientCertificateRequestedEventArgs) com.Error, scoped bool) *ICoreWebView2ClientCertificateRequestedEventHandler {
	impl := &ICoreWebView2ClientCertificateRequestedEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2ClientCertificateRequestedEventHandlerComObj(impl, scoped).ICoreWebView2ClientCertificateRequestedEventHandler()
}

