package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 609302AD-0E36-4F9A-A210-6A45272842A9
var IID_ICoreWebView2FrameNavigationCompletedEventHandler = syscall.GUID{0x609302AD, 0x0E36, 0x4F9A, 
	[8]byte{0xA2, 0x10, 0x6A, 0x45, 0x27, 0x28, 0x42, 0xA9}}

type ICoreWebView2FrameNavigationCompletedEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2FrameNavigationCompletedEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2Frame, args *ICoreWebView2NavigationCompletedEventArgs) com.Error
}

type ICoreWebView2FrameNavigationCompletedEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2FrameNavigationCompletedEventHandlerInterface
}

func (this *ICoreWebView2FrameNavigationCompletedEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2FrameNavigationCompletedEventHandlerInterface)
}

func (this *ICoreWebView2FrameNavigationCompletedEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2FrameNavigationCompletedEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2FrameNavigationCompletedEventHandlerImpl) Invoke(sender *ICoreWebView2Frame, args *ICoreWebView2NavigationCompletedEventArgs) com.Error {
	var ret com.Error
	return ret
}
type ICoreWebView2FrameNavigationCompletedEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2FrameNavigationCompletedEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2FrameNavigationCompletedEventHandlerComObj) impl() ICoreWebView2FrameNavigationCompletedEventHandlerInterface {
	return this.Impl().(ICoreWebView2FrameNavigationCompletedEventHandlerInterface)
}

func (this *ICoreWebView2FrameNavigationCompletedEventHandlerComObj) Invoke(sender *ICoreWebView2Frame, args *ICoreWebView2NavigationCompletedEventArgs) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2FrameNavigationCompletedEventHandlerVtbl *ICoreWebView2FrameNavigationCompletedEventHandlerVtbl

func (this *ICoreWebView2FrameNavigationCompletedEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2FrameNavigationCompletedEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
}
	if _pICoreWebView2FrameNavigationCompletedEventHandlerVtbl != nil {
		return _pICoreWebView2FrameNavigationCompletedEventHandlerVtbl
	}
	_pICoreWebView2FrameNavigationCompletedEventHandlerVtbl = &ICoreWebView2FrameNavigationCompletedEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:	syscall.NewCallback((*ICoreWebView2FrameNavigationCompletedEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2FrameNavigationCompletedEventHandlerVtbl
}

func (this *ICoreWebView2FrameNavigationCompletedEventHandlerComObj) ICoreWebView2FrameNavigationCompletedEventHandler() *ICoreWebView2FrameNavigationCompletedEventHandler{
	return (*ICoreWebView2FrameNavigationCompletedEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2FrameNavigationCompletedEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2FrameNavigationCompletedEventHandlerComObj(impl ICoreWebView2FrameNavigationCompletedEventHandlerInterface, scoped bool) *ICoreWebView2FrameNavigationCompletedEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2FrameNavigationCompletedEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2FrameNavigationCompletedEventHandler(impl ICoreWebView2FrameNavigationCompletedEventHandlerInterface) *ICoreWebView2FrameNavigationCompletedEventHandler {
	return NewICoreWebView2FrameNavigationCompletedEventHandlerComObj(impl, true).ICoreWebView2FrameNavigationCompletedEventHandler()}

//
type ICoreWebView2FrameNavigationCompletedEventHandlerByFuncImpl struct {
	ICoreWebView2FrameNavigationCompletedEventHandlerImpl
	handlerFunc func (sender *ICoreWebView2Frame, args *ICoreWebView2NavigationCompletedEventArgs) com.Error
}
func (this *ICoreWebView2FrameNavigationCompletedEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2Frame, args *ICoreWebView2NavigationCompletedEventArgs) com.Error{
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2FrameNavigationCompletedEventHandlerByFunc(handlerFunc func (sender *ICoreWebView2Frame, args *ICoreWebView2NavigationCompletedEventArgs) com.Error, scoped bool) *ICoreWebView2FrameNavigationCompletedEventHandler {
	impl := &ICoreWebView2FrameNavigationCompletedEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2FrameNavigationCompletedEventHandlerComObj(impl, scoped).ICoreWebView2FrameNavigationCompletedEventHandler()
}

