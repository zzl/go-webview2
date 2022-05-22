package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 3C067F9F-5388-4772-8B48-79F7EF1AB37C
var IID_ICoreWebView2SourceChangedEventHandler = syscall.GUID{0x3C067F9F, 0x5388, 0x4772, 
	[8]byte{0x8B, 0x48, 0x79, 0xF7, 0xEF, 0x1A, 0xB3, 0x7C}}

type ICoreWebView2SourceChangedEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2SourceChangedEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2, args *ICoreWebView2SourceChangedEventArgs) com.Error
}

type ICoreWebView2SourceChangedEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2SourceChangedEventHandlerInterface
}

func (this *ICoreWebView2SourceChangedEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2SourceChangedEventHandlerInterface)
}

func (this *ICoreWebView2SourceChangedEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2SourceChangedEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2SourceChangedEventHandlerImpl) Invoke(sender *ICoreWebView2, args *ICoreWebView2SourceChangedEventArgs) com.Error {
	var ret com.Error
	return ret
}
type ICoreWebView2SourceChangedEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2SourceChangedEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2SourceChangedEventHandlerComObj) impl() ICoreWebView2SourceChangedEventHandlerInterface {
	return this.Impl().(ICoreWebView2SourceChangedEventHandlerInterface)
}

func (this *ICoreWebView2SourceChangedEventHandlerComObj) Invoke(sender *ICoreWebView2, args *ICoreWebView2SourceChangedEventArgs) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2SourceChangedEventHandlerVtbl *ICoreWebView2SourceChangedEventHandlerVtbl

func (this *ICoreWebView2SourceChangedEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2SourceChangedEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
}
	if _pICoreWebView2SourceChangedEventHandlerVtbl != nil {
		return _pICoreWebView2SourceChangedEventHandlerVtbl
	}
	_pICoreWebView2SourceChangedEventHandlerVtbl = &ICoreWebView2SourceChangedEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:	syscall.NewCallback((*ICoreWebView2SourceChangedEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2SourceChangedEventHandlerVtbl
}

func (this *ICoreWebView2SourceChangedEventHandlerComObj) ICoreWebView2SourceChangedEventHandler() *ICoreWebView2SourceChangedEventHandler{
	return (*ICoreWebView2SourceChangedEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2SourceChangedEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2SourceChangedEventHandlerComObj(impl ICoreWebView2SourceChangedEventHandlerInterface, scoped bool) *ICoreWebView2SourceChangedEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2SourceChangedEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2SourceChangedEventHandler(impl ICoreWebView2SourceChangedEventHandlerInterface) *ICoreWebView2SourceChangedEventHandler {
	return NewICoreWebView2SourceChangedEventHandlerComObj(impl, true).ICoreWebView2SourceChangedEventHandler()}

//
type ICoreWebView2SourceChangedEventHandlerByFuncImpl struct {
	ICoreWebView2SourceChangedEventHandlerImpl
	handlerFunc func (sender *ICoreWebView2, args *ICoreWebView2SourceChangedEventArgs) com.Error
}
func (this *ICoreWebView2SourceChangedEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2, args *ICoreWebView2SourceChangedEventArgs) com.Error{
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2SourceChangedEventHandlerByFunc(handlerFunc func (sender *ICoreWebView2, args *ICoreWebView2SourceChangedEventArgs) com.Error, scoped bool) *ICoreWebView2SourceChangedEventHandler {
	impl := &ICoreWebView2SourceChangedEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2SourceChangedEventHandlerComObj(impl, scoped).ICoreWebView2SourceChangedEventHandler()
}

