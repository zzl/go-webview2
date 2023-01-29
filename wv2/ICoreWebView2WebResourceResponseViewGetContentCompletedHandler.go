package wv2

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 875738E1-9FA2-40E3-8B74-2E8972DD6FE7
var IID_ICoreWebView2WebResourceResponseViewGetContentCompletedHandler = syscall.GUID{0x875738E1, 0x9FA2, 0x40E3,
	[8]byte{0x8B, 0x74, 0x2E, 0x89, 0x72, 0xDD, 0x6F, 0xE7}}

type ICoreWebView2WebResourceResponseViewGetContentCompletedHandler struct {
	win32.IUnknown
}

type ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(errorCode com.Error, content *win32.IStream) com.Error
}

type ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerInterface
}

func (this *ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerInterface)
}

func (this *ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2WebResourceResponseViewGetContentCompletedHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerImpl) Invoke(errorCode com.Error, content *win32.IStream) com.Error {
	var ret com.Error
	return ret
}

type ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerComObj) impl() ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerInterface {
	return this.Impl().(ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerInterface)
}

func (this *ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerComObj) Invoke(errorCode com.Error, content *win32.IStream) uintptr {
	return (uintptr)(this.impl().Invoke(errorCode, content))
}

var _pICoreWebView2WebResourceResponseViewGetContentCompletedHandlerVtbl *ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerVtbl

func (this *ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
	}
	if _pICoreWebView2WebResourceResponseViewGetContentCompletedHandlerVtbl != nil {
		return _pICoreWebView2WebResourceResponseViewGetContentCompletedHandlerVtbl
	}
	_pICoreWebView2WebResourceResponseViewGetContentCompletedHandlerVtbl = &ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:       syscall.NewCallback((*ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerComObj).Invoke),
	}
	return _pICoreWebView2WebResourceResponseViewGetContentCompletedHandlerVtbl
}

func (this *ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerComObj) ICoreWebView2WebResourceResponseViewGetContentCompletedHandler() *ICoreWebView2WebResourceResponseViewGetContentCompletedHandler {
	return (*ICoreWebView2WebResourceResponseViewGetContentCompletedHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2WebResourceResponseViewGetContentCompletedHandlerComObj(impl ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerInterface, scoped bool) *ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2WebResourceResponseViewGetContentCompletedHandler(impl ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerInterface) *ICoreWebView2WebResourceResponseViewGetContentCompletedHandler {
	return NewICoreWebView2WebResourceResponseViewGetContentCompletedHandlerComObj(impl, true).ICoreWebView2WebResourceResponseViewGetContentCompletedHandler()
}

//
type ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerByFuncImpl struct {
	ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerImpl
	handlerFunc func(errorCode com.Error, content *win32.IStream) com.Error
}

func (this *ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerByFuncImpl) Invoke(errorCode com.Error, content *win32.IStream) com.Error {
	return this.handlerFunc(errorCode, content)
}

func NewICoreWebView2WebResourceResponseViewGetContentCompletedHandlerByFunc(handlerFunc func(errorCode com.Error, content *win32.IStream) com.Error, scoped bool) *ICoreWebView2WebResourceResponseViewGetContentCompletedHandler {
	impl := &ICoreWebView2WebResourceResponseViewGetContentCompletedHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2WebResourceResponseViewGetContentCompletedHandlerComObj(impl, scoped).ICoreWebView2WebResourceResponseViewGetContentCompletedHandler()
}

