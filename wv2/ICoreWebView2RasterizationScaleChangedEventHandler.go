package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 9C98C8B1-AC53-427E-A345-3049B5524BBE
var IID_ICoreWebView2RasterizationScaleChangedEventHandler = syscall.GUID{0x9C98C8B1, 0xAC53, 0x427E, 
	[8]byte{0xA3, 0x45, 0x30, 0x49, 0xB5, 0x52, 0x4B, 0xBE}}

type ICoreWebView2RasterizationScaleChangedEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2RasterizationScaleChangedEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2Controller, args *com.UnknownClass) com.Error
}

type ICoreWebView2RasterizationScaleChangedEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2RasterizationScaleChangedEventHandlerInterface
}

func (this *ICoreWebView2RasterizationScaleChangedEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2RasterizationScaleChangedEventHandlerInterface)
}

func (this *ICoreWebView2RasterizationScaleChangedEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2RasterizationScaleChangedEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2RasterizationScaleChangedEventHandlerImpl) Invoke(sender *ICoreWebView2Controller, args *com.UnknownClass) com.Error {
	var ret com.Error
	return ret
}
type ICoreWebView2RasterizationScaleChangedEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2RasterizationScaleChangedEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2RasterizationScaleChangedEventHandlerComObj) impl() ICoreWebView2RasterizationScaleChangedEventHandlerInterface {
	return this.Impl().(ICoreWebView2RasterizationScaleChangedEventHandlerInterface)
}

func (this *ICoreWebView2RasterizationScaleChangedEventHandlerComObj) Invoke(sender *ICoreWebView2Controller, args *com.UnknownClass) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2RasterizationScaleChangedEventHandlerVtbl *ICoreWebView2RasterizationScaleChangedEventHandlerVtbl

func (this *ICoreWebView2RasterizationScaleChangedEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2RasterizationScaleChangedEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
}
	if _pICoreWebView2RasterizationScaleChangedEventHandlerVtbl != nil {
		return _pICoreWebView2RasterizationScaleChangedEventHandlerVtbl
	}
	_pICoreWebView2RasterizationScaleChangedEventHandlerVtbl = &ICoreWebView2RasterizationScaleChangedEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:	syscall.NewCallback((*ICoreWebView2RasterizationScaleChangedEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2RasterizationScaleChangedEventHandlerVtbl
}

func (this *ICoreWebView2RasterizationScaleChangedEventHandlerComObj) ICoreWebView2RasterizationScaleChangedEventHandler() *ICoreWebView2RasterizationScaleChangedEventHandler{
	return (*ICoreWebView2RasterizationScaleChangedEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2RasterizationScaleChangedEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2RasterizationScaleChangedEventHandlerComObj(impl ICoreWebView2RasterizationScaleChangedEventHandlerInterface, scoped bool) *ICoreWebView2RasterizationScaleChangedEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2RasterizationScaleChangedEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2RasterizationScaleChangedEventHandler(impl ICoreWebView2RasterizationScaleChangedEventHandlerInterface) *ICoreWebView2RasterizationScaleChangedEventHandler {
	return NewICoreWebView2RasterizationScaleChangedEventHandlerComObj(impl, true).ICoreWebView2RasterizationScaleChangedEventHandler()}

//
type ICoreWebView2RasterizationScaleChangedEventHandlerByFuncImpl struct {
	ICoreWebView2RasterizationScaleChangedEventHandlerImpl
	handlerFunc func (sender *ICoreWebView2Controller, args *com.UnknownClass) com.Error
}
func (this *ICoreWebView2RasterizationScaleChangedEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2Controller, args *com.UnknownClass) com.Error{
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2RasterizationScaleChangedEventHandlerByFunc(handlerFunc func (sender *ICoreWebView2Controller, args *com.UnknownClass) com.Error, scoped bool) *ICoreWebView2RasterizationScaleChangedEventHandler {
	impl := &ICoreWebView2RasterizationScaleChangedEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2RasterizationScaleChangedEventHandlerComObj(impl, scoped).ICoreWebView2RasterizationScaleChangedEventHandler()
}

