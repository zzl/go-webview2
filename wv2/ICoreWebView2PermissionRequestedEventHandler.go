package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 15E1C6A3-C72A-4DF3-91D7-D097FBEC6BFD
var IID_ICoreWebView2PermissionRequestedEventHandler = syscall.GUID{0x15E1C6A3, 0xC72A, 0x4DF3, 
	[8]byte{0x91, 0xD7, 0xD0, 0x97, 0xFB, 0xEC, 0x6B, 0xFD}}

type ICoreWebView2PermissionRequestedEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2PermissionRequestedEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2, args *ICoreWebView2PermissionRequestedEventArgs) com.Error
}

type ICoreWebView2PermissionRequestedEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2PermissionRequestedEventHandlerInterface
}

func (this *ICoreWebView2PermissionRequestedEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2PermissionRequestedEventHandlerInterface)
}

func (this *ICoreWebView2PermissionRequestedEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2PermissionRequestedEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2PermissionRequestedEventHandlerImpl) Invoke(sender *ICoreWebView2, args *ICoreWebView2PermissionRequestedEventArgs) com.Error {
	var ret com.Error
	return ret
}
type ICoreWebView2PermissionRequestedEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2PermissionRequestedEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2PermissionRequestedEventHandlerComObj) impl() ICoreWebView2PermissionRequestedEventHandlerInterface {
	return this.Impl().(ICoreWebView2PermissionRequestedEventHandlerInterface)
}

func (this *ICoreWebView2PermissionRequestedEventHandlerComObj) Invoke(sender *ICoreWebView2, args *ICoreWebView2PermissionRequestedEventArgs) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2PermissionRequestedEventHandlerVtbl *ICoreWebView2PermissionRequestedEventHandlerVtbl

func (this *ICoreWebView2PermissionRequestedEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2PermissionRequestedEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
}
	if _pICoreWebView2PermissionRequestedEventHandlerVtbl != nil {
		return _pICoreWebView2PermissionRequestedEventHandlerVtbl
	}
	_pICoreWebView2PermissionRequestedEventHandlerVtbl = &ICoreWebView2PermissionRequestedEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:	syscall.NewCallback((*ICoreWebView2PermissionRequestedEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2PermissionRequestedEventHandlerVtbl
}

func (this *ICoreWebView2PermissionRequestedEventHandlerComObj) ICoreWebView2PermissionRequestedEventHandler() *ICoreWebView2PermissionRequestedEventHandler{
	return (*ICoreWebView2PermissionRequestedEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2PermissionRequestedEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2PermissionRequestedEventHandlerComObj(impl ICoreWebView2PermissionRequestedEventHandlerInterface, scoped bool) *ICoreWebView2PermissionRequestedEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2PermissionRequestedEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2PermissionRequestedEventHandler(impl ICoreWebView2PermissionRequestedEventHandlerInterface) *ICoreWebView2PermissionRequestedEventHandler {
	return NewICoreWebView2PermissionRequestedEventHandlerComObj(impl, true).ICoreWebView2PermissionRequestedEventHandler()}

//
type ICoreWebView2PermissionRequestedEventHandlerByFuncImpl struct {
	ICoreWebView2PermissionRequestedEventHandlerImpl
	handlerFunc func (sender *ICoreWebView2, args *ICoreWebView2PermissionRequestedEventArgs) com.Error
}
func (this *ICoreWebView2PermissionRequestedEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2, args *ICoreWebView2PermissionRequestedEventArgs) com.Error{
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2PermissionRequestedEventHandlerByFunc(handlerFunc func (sender *ICoreWebView2, args *ICoreWebView2PermissionRequestedEventArgs) com.Error, scoped bool) *ICoreWebView2PermissionRequestedEventHandler {
	impl := &ICoreWebView2PermissionRequestedEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2PermissionRequestedEventHandlerComObj(impl, scoped).ICoreWebView2PermissionRequestedEventHandler()
}

