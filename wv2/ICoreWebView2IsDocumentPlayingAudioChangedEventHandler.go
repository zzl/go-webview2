package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 5DEF109A-2F4B-49FA-B7F6-11C39E513328
var IID_ICoreWebView2IsDocumentPlayingAudioChangedEventHandler = syscall.GUID{0x5DEF109A, 0x2F4B, 0x49FA, 
	[8]byte{0xB7, 0xF6, 0x11, 0xC3, 0x9E, 0x51, 0x33, 0x28}}

type ICoreWebView2IsDocumentPlayingAudioChangedEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2, args *win32.IUnknown) com.Error
}

type ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerInterface
}

func (this *ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerInterface)
}

func (this *ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2IsDocumentPlayingAudioChangedEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerImpl) Invoke(sender *ICoreWebView2, args *win32.IUnknown) com.Error {
	var ret com.Error
	return ret
}
type ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerComObj) impl() ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerInterface {
	return this.Impl().(ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerInterface)
}

func (this *ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerComObj) Invoke(sender *ICoreWebView2, args *win32.IUnknown) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2IsDocumentPlayingAudioChangedEventHandlerVtbl *ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerVtbl

func (this *ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
}
	if _pICoreWebView2IsDocumentPlayingAudioChangedEventHandlerVtbl != nil {
		return _pICoreWebView2IsDocumentPlayingAudioChangedEventHandlerVtbl
	}
	_pICoreWebView2IsDocumentPlayingAudioChangedEventHandlerVtbl = &ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:	syscall.NewCallback((*ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2IsDocumentPlayingAudioChangedEventHandlerVtbl
}

func (this *ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerComObj) ICoreWebView2IsDocumentPlayingAudioChangedEventHandler() *ICoreWebView2IsDocumentPlayingAudioChangedEventHandler{
	return (*ICoreWebView2IsDocumentPlayingAudioChangedEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2IsDocumentPlayingAudioChangedEventHandlerComObj(impl ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerInterface, scoped bool) *ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2IsDocumentPlayingAudioChangedEventHandler(impl ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerInterface) *ICoreWebView2IsDocumentPlayingAudioChangedEventHandler {
	return NewICoreWebView2IsDocumentPlayingAudioChangedEventHandlerComObj(impl, true).ICoreWebView2IsDocumentPlayingAudioChangedEventHandler()}

//
type ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerByFuncImpl struct {
	ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerImpl
	handlerFunc func (sender *ICoreWebView2, args *win32.IUnknown) com.Error
}
func (this *ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2, args *win32.IUnknown) com.Error{
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2IsDocumentPlayingAudioChangedEventHandlerByFunc(handlerFunc func (sender *ICoreWebView2, args *win32.IUnknown) com.Error, scoped bool) *ICoreWebView2IsDocumentPlayingAudioChangedEventHandler {
	impl := &ICoreWebView2IsDocumentPlayingAudioChangedEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2IsDocumentPlayingAudioChangedEventHandlerComObj(impl, scoped).ICoreWebView2IsDocumentPlayingAudioChangedEventHandler()
}

