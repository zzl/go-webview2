package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// E79908BF-2D5D-4968-83DB-263FEA2C1DA3
var IID_ICoreWebView2FrameNavigationStartingEventHandler = syscall.GUID{0xE79908BF, 0x2D5D, 0x4968, 
	[8]byte{0x83, 0xDB, 0x26, 0x3F, 0xEA, 0x2C, 0x1D, 0xA3}}

type ICoreWebView2FrameNavigationStartingEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2FrameNavigationStartingEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2Frame, args *ICoreWebView2NavigationStartingEventArgs) com.Error
}

type ICoreWebView2FrameNavigationStartingEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2FrameNavigationStartingEventHandlerInterface
}

func (this *ICoreWebView2FrameNavigationStartingEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2FrameNavigationStartingEventHandlerInterface)
}

func (this *ICoreWebView2FrameNavigationStartingEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2FrameNavigationStartingEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2FrameNavigationStartingEventHandlerImpl) Invoke(sender *ICoreWebView2Frame, args *ICoreWebView2NavigationStartingEventArgs) com.Error {
	var ret com.Error
	return ret
}
type ICoreWebView2FrameNavigationStartingEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2FrameNavigationStartingEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2FrameNavigationStartingEventHandlerComObj) impl() ICoreWebView2FrameNavigationStartingEventHandlerInterface {
	return this.Impl().(ICoreWebView2FrameNavigationStartingEventHandlerInterface)
}

func (this *ICoreWebView2FrameNavigationStartingEventHandlerComObj) Invoke(sender *ICoreWebView2Frame, args *ICoreWebView2NavigationStartingEventArgs) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2FrameNavigationStartingEventHandlerVtbl *ICoreWebView2FrameNavigationStartingEventHandlerVtbl

func (this *ICoreWebView2FrameNavigationStartingEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2FrameNavigationStartingEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
}
	if _pICoreWebView2FrameNavigationStartingEventHandlerVtbl != nil {
		return _pICoreWebView2FrameNavigationStartingEventHandlerVtbl
	}
	_pICoreWebView2FrameNavigationStartingEventHandlerVtbl = &ICoreWebView2FrameNavigationStartingEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:	syscall.NewCallback((*ICoreWebView2FrameNavigationStartingEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2FrameNavigationStartingEventHandlerVtbl
}

func (this *ICoreWebView2FrameNavigationStartingEventHandlerComObj) ICoreWebView2FrameNavigationStartingEventHandler() *ICoreWebView2FrameNavigationStartingEventHandler{
	return (*ICoreWebView2FrameNavigationStartingEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2FrameNavigationStartingEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2FrameNavigationStartingEventHandlerComObj(impl ICoreWebView2FrameNavigationStartingEventHandlerInterface, scoped bool) *ICoreWebView2FrameNavigationStartingEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2FrameNavigationStartingEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2FrameNavigationStartingEventHandler(impl ICoreWebView2FrameNavigationStartingEventHandlerInterface) *ICoreWebView2FrameNavigationStartingEventHandler {
	return NewICoreWebView2FrameNavigationStartingEventHandlerComObj(impl, true).ICoreWebView2FrameNavigationStartingEventHandler()}

//
type ICoreWebView2FrameNavigationStartingEventHandlerByFuncImpl struct {
	ICoreWebView2FrameNavigationStartingEventHandlerImpl
	handlerFunc func (sender *ICoreWebView2Frame, args *ICoreWebView2NavigationStartingEventArgs) com.Error
}
func (this *ICoreWebView2FrameNavigationStartingEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2Frame, args *ICoreWebView2NavigationStartingEventArgs) com.Error{
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2FrameNavigationStartingEventHandlerByFunc(handlerFunc func (sender *ICoreWebView2Frame, args *ICoreWebView2NavigationStartingEventArgs) com.Error, scoped bool) *ICoreWebView2FrameNavigationStartingEventHandler {
	impl := &ICoreWebView2FrameNavigationStartingEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2FrameNavigationStartingEventHandlerComObj(impl, scoped).ICoreWebView2FrameNavigationStartingEventHandler()
}

