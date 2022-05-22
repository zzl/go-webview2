package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// FA504257-A216-4911-A860-FE8825712861
var IID_ICoreWebView2BrowserProcessExitedEventHandler = syscall.GUID{0xFA504257, 0xA216, 0x4911, 
	[8]byte{0xA8, 0x60, 0xFE, 0x88, 0x25, 0x71, 0x28, 0x61}}

type ICoreWebView2BrowserProcessExitedEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2BrowserProcessExitedEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2Environment, args *ICoreWebView2BrowserProcessExitedEventArgs) com.Error
}

type ICoreWebView2BrowserProcessExitedEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2BrowserProcessExitedEventHandlerInterface
}

func (this *ICoreWebView2BrowserProcessExitedEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2BrowserProcessExitedEventHandlerInterface)
}

func (this *ICoreWebView2BrowserProcessExitedEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2BrowserProcessExitedEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2BrowserProcessExitedEventHandlerImpl) Invoke(sender *ICoreWebView2Environment, args *ICoreWebView2BrowserProcessExitedEventArgs) com.Error {
	var ret com.Error
	return ret
}
type ICoreWebView2BrowserProcessExitedEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2BrowserProcessExitedEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2BrowserProcessExitedEventHandlerComObj) impl() ICoreWebView2BrowserProcessExitedEventHandlerInterface {
	return this.Impl().(ICoreWebView2BrowserProcessExitedEventHandlerInterface)
}

func (this *ICoreWebView2BrowserProcessExitedEventHandlerComObj) Invoke(sender *ICoreWebView2Environment, args *ICoreWebView2BrowserProcessExitedEventArgs) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2BrowserProcessExitedEventHandlerVtbl *ICoreWebView2BrowserProcessExitedEventHandlerVtbl

func (this *ICoreWebView2BrowserProcessExitedEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2BrowserProcessExitedEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
}
	if _pICoreWebView2BrowserProcessExitedEventHandlerVtbl != nil {
		return _pICoreWebView2BrowserProcessExitedEventHandlerVtbl
	}
	_pICoreWebView2BrowserProcessExitedEventHandlerVtbl = &ICoreWebView2BrowserProcessExitedEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:	syscall.NewCallback((*ICoreWebView2BrowserProcessExitedEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2BrowserProcessExitedEventHandlerVtbl
}

func (this *ICoreWebView2BrowserProcessExitedEventHandlerComObj) ICoreWebView2BrowserProcessExitedEventHandler() *ICoreWebView2BrowserProcessExitedEventHandler{
	return (*ICoreWebView2BrowserProcessExitedEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2BrowserProcessExitedEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2BrowserProcessExitedEventHandlerComObj(impl ICoreWebView2BrowserProcessExitedEventHandlerInterface, scoped bool) *ICoreWebView2BrowserProcessExitedEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2BrowserProcessExitedEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2BrowserProcessExitedEventHandler(impl ICoreWebView2BrowserProcessExitedEventHandlerInterface) *ICoreWebView2BrowserProcessExitedEventHandler {
	return NewICoreWebView2BrowserProcessExitedEventHandlerComObj(impl, true).ICoreWebView2BrowserProcessExitedEventHandler()}

//
type ICoreWebView2BrowserProcessExitedEventHandlerByFuncImpl struct {
	ICoreWebView2BrowserProcessExitedEventHandlerImpl
	handlerFunc func (sender *ICoreWebView2Environment, args *ICoreWebView2BrowserProcessExitedEventArgs) com.Error
}
func (this *ICoreWebView2BrowserProcessExitedEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2Environment, args *ICoreWebView2BrowserProcessExitedEventArgs) com.Error{
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2BrowserProcessExitedEventHandlerByFunc(handlerFunc func (sender *ICoreWebView2Environment, args *ICoreWebView2BrowserProcessExitedEventArgs) com.Error, scoped bool) *ICoreWebView2BrowserProcessExitedEventHandler {
	impl := &ICoreWebView2BrowserProcessExitedEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2BrowserProcessExitedEventHandlerComObj(impl, scoped).ICoreWebView2BrowserProcessExitedEventHandler()
}

