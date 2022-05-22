package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"syscall"
	"unsafe"
)

// 76ECEACB-0462-4D94-AC83-423A6793775E
var IID_ICoreWebView2 = syscall.GUID{0x76ECEACB, 0x0462, 0x4D94, 
	[8]byte{0xAC, 0x83, 0x42, 0x3A, 0x67, 0x93, 0x77, 0x5E}}

type ICoreWebView2 struct {
	win32.IUnknown
}

func NewICoreWebView2(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2 {
	p := (*ICoreWebView2)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2) IID() *syscall.GUID {
	return &IID_ICoreWebView2
}

func (this *ICoreWebView2) GetSettings(settings **ICoreWebView2Settings) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(settings)))
	if com.CurrentScope != nil {
		com.CurrentScope.Add(unsafe.Pointer(&(*settings).IUnknown))
	}
	return com.Error(ret)
}

func (this *ICoreWebView2) GetSource(uri *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)))
	return com.Error(ret)
}

func (this *ICoreWebView2) Navigate(uri string) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(uri)))
	return com.Error(ret)
}

func (this *ICoreWebView2) NavigateToString(htmlContent string) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(htmlContent)))
	return com.Error(ret)
}

func (this *ICoreWebView2) Add_NavigationStarting(eventHandler *ICoreWebView2NavigationStartingEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2) Remove_NavigationStarting(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

func (this *ICoreWebView2) Add_ContentLoading(eventHandler *ICoreWebView2ContentLoadingEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[9]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2) Remove_ContentLoading(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[10]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

func (this *ICoreWebView2) Add_SourceChanged(eventHandler *ICoreWebView2SourceChangedEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[11]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2) Remove_SourceChanged(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[12]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

func (this *ICoreWebView2) Add_HistoryChanged(eventHandler *ICoreWebView2HistoryChangedEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[13]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2) Remove_HistoryChanged(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[14]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

func (this *ICoreWebView2) Add_NavigationCompleted(eventHandler *ICoreWebView2NavigationCompletedEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[15]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2) Remove_NavigationCompleted(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[16]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

func (this *ICoreWebView2) Add_FrameNavigationStarting(eventHandler *ICoreWebView2NavigationStartingEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[17]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2) Remove_FrameNavigationStarting(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[18]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

func (this *ICoreWebView2) Add_FrameNavigationCompleted(eventHandler *ICoreWebView2NavigationCompletedEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[19]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2) Remove_FrameNavigationCompleted(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[20]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

func (this *ICoreWebView2) Add_ScriptDialogOpening(eventHandler *ICoreWebView2ScriptDialogOpeningEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[21]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2) Remove_ScriptDialogOpening(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[22]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

func (this *ICoreWebView2) Add_PermissionRequested(eventHandler *ICoreWebView2PermissionRequestedEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[23]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2) Remove_PermissionRequested(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[24]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

func (this *ICoreWebView2) Add_ProcessFailed(eventHandler *ICoreWebView2ProcessFailedEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[25]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2) Remove_ProcessFailed(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[26]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

func (this *ICoreWebView2) AddScriptToExecuteOnDocumentCreated(javaScript string, handler *ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandler) com.Error {
	addr := (*this.LpVtbl)[27]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(javaScript)), uintptr(unsafe.Pointer(handler)))
	return com.Error(ret)
}

func (this *ICoreWebView2) RemoveScriptToExecuteOnDocumentCreated(id string) com.Error {
	addr := (*this.LpVtbl)[28]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(id)))
	return com.Error(ret)
}

func (this *ICoreWebView2) ExecuteScript(javaScript string, handler *ICoreWebView2ExecuteScriptCompletedHandler) com.Error {
	addr := (*this.LpVtbl)[29]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(javaScript)), uintptr(unsafe.Pointer(handler)))
	return com.Error(ret)
}

func (this *ICoreWebView2) CapturePreview(imageFormat int32, imageStream *win32.IStream, handler *ICoreWebView2CapturePreviewCompletedHandler) com.Error {
	addr := (*this.LpVtbl)[30]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(imageFormat), uintptr(unsafe.Pointer(imageStream)), uintptr(unsafe.Pointer(handler)))
	return com.Error(ret)
}

func (this *ICoreWebView2) Reload() com.Error {
	addr := (*this.LpVtbl)[31]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return com.Error(ret)
}

func (this *ICoreWebView2) PostWebMessageAsJson(webMessageAsJson string) com.Error {
	addr := (*this.LpVtbl)[32]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(webMessageAsJson)))
	return com.Error(ret)
}

func (this *ICoreWebView2) PostWebMessageAsString(webMessageAsString string) com.Error {
	addr := (*this.LpVtbl)[33]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(webMessageAsString)))
	return com.Error(ret)
}

func (this *ICoreWebView2) Add_WebMessageReceived(handler *ICoreWebView2WebMessageReceivedEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[34]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(handler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2) Remove_WebMessageReceived(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[35]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

func (this *ICoreWebView2) CallDevToolsProtocolMethod(methodName string, parametersAsJson string, handler *ICoreWebView2CallDevToolsProtocolMethodCompletedHandler) com.Error {
	addr := (*this.LpVtbl)[36]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(methodName)), uintptr(win32.StrToPointer(parametersAsJson)), uintptr(unsafe.Pointer(handler)))
	return com.Error(ret)
}

func (this *ICoreWebView2) GetBrowserProcessId(value *uint32) com.Error {
	addr := (*this.LpVtbl)[37]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2) GetCanGoBack(canGoBack *int32) com.Error {
	addr := (*this.LpVtbl)[38]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(canGoBack)))
	return com.Error(ret)
}

func (this *ICoreWebView2) GetCanGoForward(canGoForward *int32) com.Error {
	addr := (*this.LpVtbl)[39]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(canGoForward)))
	return com.Error(ret)
}

func (this *ICoreWebView2) GoBack() com.Error {
	addr := (*this.LpVtbl)[40]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return com.Error(ret)
}

func (this *ICoreWebView2) GoForward() com.Error {
	addr := (*this.LpVtbl)[41]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return com.Error(ret)
}

func (this *ICoreWebView2) GetDevToolsProtocolEventReceiver(eventName string, receiver **ICoreWebView2DevToolsProtocolEventReceiver) com.Error {
	addr := (*this.LpVtbl)[42]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(eventName)), uintptr(unsafe.Pointer(receiver)))
	if com.CurrentScope != nil {
		com.CurrentScope.Add(unsafe.Pointer(&(*receiver).IUnknown))
	}
	return com.Error(ret)
}

func (this *ICoreWebView2) Stop() com.Error {
	addr := (*this.LpVtbl)[43]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return com.Error(ret)
}

func (this *ICoreWebView2) Add_NewWindowRequested(eventHandler *ICoreWebView2NewWindowRequestedEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[44]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2) Remove_NewWindowRequested(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[45]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

func (this *ICoreWebView2) Add_DocumentTitleChanged(eventHandler *ICoreWebView2DocumentTitleChangedEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[46]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2) Remove_DocumentTitleChanged(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[47]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

func (this *ICoreWebView2) GetDocumentTitle(title *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[48]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(title)))
	return com.Error(ret)
}

func (this *ICoreWebView2) AddHostObjectToScript(name string, object *ole.Variant) com.Error {
	addr := (*this.LpVtbl)[49]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(name)), uintptr(unsafe.Pointer(object)))
	return com.Error(ret)
}

func (this *ICoreWebView2) RemoveHostObjectFromScript(name string) com.Error {
	addr := (*this.LpVtbl)[50]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(name)))
	return com.Error(ret)
}

func (this *ICoreWebView2) OpenDevToolsWindow() com.Error {
	addr := (*this.LpVtbl)[51]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return com.Error(ret)
}

func (this *ICoreWebView2) Add_ContainsFullScreenElementChanged(eventHandler *ICoreWebView2ContainsFullScreenElementChangedEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[52]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2) Remove_ContainsFullScreenElementChanged(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[53]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

func (this *ICoreWebView2) GetContainsFullScreenElement(containsFullScreenElement *int32) com.Error {
	addr := (*this.LpVtbl)[54]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(containsFullScreenElement)))
	return com.Error(ret)
}

func (this *ICoreWebView2) Add_WebResourceRequested(eventHandler *ICoreWebView2WebResourceRequestedEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[55]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2) Remove_WebResourceRequested(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[56]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

func (this *ICoreWebView2) AddWebResourceRequestedFilter(uri string, resourceContext int32) com.Error {
	addr := (*this.LpVtbl)[57]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(uri)), uintptr(resourceContext))
	return com.Error(ret)
}

func (this *ICoreWebView2) RemoveWebResourceRequestedFilter(uri string, resourceContext int32) com.Error {
	addr := (*this.LpVtbl)[58]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(uri)), uintptr(resourceContext))
	return com.Error(ret)
}

func (this *ICoreWebView2) Add_WindowCloseRequested(eventHandler *ICoreWebView2WindowCloseRequestedEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[59]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2) Remove_WindowCloseRequested(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[60]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

