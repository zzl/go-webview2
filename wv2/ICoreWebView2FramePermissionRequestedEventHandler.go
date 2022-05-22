package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 845D0EDD-8BD8-429B-9915-4821789F23E9
var IID_ICoreWebView2FramePermissionRequestedEventHandler = syscall.GUID{0x845D0EDD, 0x8BD8, 0x429B, 
	[8]byte{0x99, 0x15, 0x48, 0x21, 0x78, 0x9F, 0x23, 0xE9}}

type ICoreWebView2FramePermissionRequestedEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2FramePermissionRequestedEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2Frame, args *ICoreWebView2PermissionRequestedEventArgs2) com.Error
}

type ICoreWebView2FramePermissionRequestedEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2FramePermissionRequestedEventHandlerInterface
}

func (this *ICoreWebView2FramePermissionRequestedEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2FramePermissionRequestedEventHandlerInterface)
}

func (this *ICoreWebView2FramePermissionRequestedEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2FramePermissionRequestedEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2FramePermissionRequestedEventHandlerImpl) Invoke(sender *ICoreWebView2Frame, args *ICoreWebView2PermissionRequestedEventArgs2) com.Error {
	var ret com.Error
	return ret
}
type ICoreWebView2FramePermissionRequestedEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2FramePermissionRequestedEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2FramePermissionRequestedEventHandlerComObj) impl() ICoreWebView2FramePermissionRequestedEventHandlerInterface {
	return this.Impl().(ICoreWebView2FramePermissionRequestedEventHandlerInterface)
}

func (this *ICoreWebView2FramePermissionRequestedEventHandlerComObj) Invoke(sender *ICoreWebView2Frame, args *ICoreWebView2PermissionRequestedEventArgs2) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2FramePermissionRequestedEventHandlerVtbl *ICoreWebView2FramePermissionRequestedEventHandlerVtbl

func (this *ICoreWebView2FramePermissionRequestedEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2FramePermissionRequestedEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
}
	if _pICoreWebView2FramePermissionRequestedEventHandlerVtbl != nil {
		return _pICoreWebView2FramePermissionRequestedEventHandlerVtbl
	}
	_pICoreWebView2FramePermissionRequestedEventHandlerVtbl = &ICoreWebView2FramePermissionRequestedEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:	syscall.NewCallback((*ICoreWebView2FramePermissionRequestedEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2FramePermissionRequestedEventHandlerVtbl
}

func (this *ICoreWebView2FramePermissionRequestedEventHandlerComObj) ICoreWebView2FramePermissionRequestedEventHandler() *ICoreWebView2FramePermissionRequestedEventHandler{
	return (*ICoreWebView2FramePermissionRequestedEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2FramePermissionRequestedEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2FramePermissionRequestedEventHandlerComObj(impl ICoreWebView2FramePermissionRequestedEventHandlerInterface, scoped bool) *ICoreWebView2FramePermissionRequestedEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2FramePermissionRequestedEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2FramePermissionRequestedEventHandler(impl ICoreWebView2FramePermissionRequestedEventHandlerInterface) *ICoreWebView2FramePermissionRequestedEventHandler {
	return NewICoreWebView2FramePermissionRequestedEventHandlerComObj(impl, true).ICoreWebView2FramePermissionRequestedEventHandler()}

//
type ICoreWebView2FramePermissionRequestedEventHandlerByFuncImpl struct {
	ICoreWebView2FramePermissionRequestedEventHandlerImpl
	handlerFunc func (sender *ICoreWebView2Frame, args *ICoreWebView2PermissionRequestedEventArgs2) com.Error
}
func (this *ICoreWebView2FramePermissionRequestedEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2Frame, args *ICoreWebView2PermissionRequestedEventArgs2) com.Error{
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2FramePermissionRequestedEventHandlerByFunc(handlerFunc func (sender *ICoreWebView2Frame, args *ICoreWebView2PermissionRequestedEventArgs2) com.Error, scoped bool) *ICoreWebView2FramePermissionRequestedEventHandler {
	impl := &ICoreWebView2FramePermissionRequestedEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2FramePermissionRequestedEventHandlerComObj(impl, scoped).ICoreWebView2FramePermissionRequestedEventHandler()
}

