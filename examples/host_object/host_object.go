package main

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"github.com/zzl/go-com/ole/oleimpl"
	"github.com/zzl/go-webview2/wv2"
	"github.com/zzl/go-win32api/win32"
	"log"
	"net"
	"net/http"
	"syscall"
	"time"
	"unsafe"
)

var globalScope *com.Scope
var webviewController *wv2.ICoreWebView2Controller

func main() {

	com.Initialize()

	globalScope = com.NewScope()
	defer globalScope.Leave()

	hWnd := createMainWindow()
	win32.ShowWindow(hWnd, win32.SW_SHOW)
	win32.UpdateWindow(hWnd)

	serveHttp()

	createWebView(hWnd)
	messageLoop()
}

func createMainWindow() win32.HWND {
	hInst, _ := win32.GetModuleHandle(nil)

	var wcex win32.WNDCLASSEX
	wcex.CbSize = uint32(unsafe.Sizeof(wcex))
	wcex.LpfnWndProc = syscall.NewCallback(WndProc)
	wcex.HInstance = hInst
	wcex.LpszClassName = win32.StrToPwstr("HostObjectTest")

	win32.RegisterClassEx(&wcex)

	hWnd, _ := win32.CreateWindowEx(0,
		wcex.LpszClassName,
		win32.StrToPwstr("Host Object test"),
		win32.WS_OVERLAPPEDWINDOW,
		win32.CW_USEDEFAULT, win32.CW_USEDEFAULT,
		1200, 900, 0, 0, hInst, nil,
	)
	return hWnd
}

func createWebView(hWnd win32.HWND) win32.HRESULT {
	return wv2.CreateCoreWebView2EnvironmentWithOptions("", "", nil,
		wv2.NewICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerByFunc(
			func(errorCode com.Error, env *wv2.ICoreWebView2Environment) com.Error {
				env.CreateCoreWebView2Controller(hWnd,
					wv2.NewICoreWebView2CreateCoreWebView2ControllerCompletedHandlerByFunc(
						createCoreWebView2ControllerCompletedHandler, true))
				return com.Error(win32.S_OK)
			}, true))
}

func createCoreWebView2ControllerCompletedHandler(errorCode com.Error,
	controller *wv2.ICoreWebView2Controller) com.Error {

	webviewController = controller
	globalScope.AddComPtr(webviewController, true)

	defer com.NewScope().Leave()

	var webviewWindow *wv2.ICoreWebView2
	webviewController.GetCoreWebView2(&webviewWindow)

	var settings *wv2.ICoreWebView2Settings
	webviewWindow.GetSettings(&settings)
	settings.SetIsScriptEnabled(win32.TRUE)
	settings.SetIsWebMessageEnabled(win32.TRUE)
	settings.SetAreDefaultScriptDialogsEnabled(win32.TRUE)

	var hWnd win32.HWND
	webviewController.GetParentWindow(&hWnd)
	win32.SendMessage(hWnd, win32.WM_SIZE, 0, 0)

	//
	pDisp := buildHostObject()
	globalScope.AddComPtr(pDisp)
	webviewWindow.AddHostObjectToScript("TestObj", ole.NewVarScoped(pDisp))

	//
	webviewWindow.Navigate("http://127.0.0.1:8080/")

	return com.Error(win32.S_OK)
}

type SubObject struct {
	prop string
}

func (this *SubObject) GetProp() string {
	return this.prop
}

func (this *SubObject) SetProp(prop string) {
	this.prop = prop
}

func (this *SubObject) Func(param string) string {
	return "@SubObject.Func: " + param
}

func buildHostObject() *win32.IDispatch {
	var count = 0

	subObject := &SubObject{}
	pDispSubObject := oleimpl.NewReflectDispatch(subObject)

	return oleimpl.NewFuncMapDispatch(map[string]oleimpl.VariantFunc{
		"sayHello": func(args ...*ole.Variant) *ole.Variant {
			name, _ := args[0].ToString()
			hWnd := win32.GetForegroundWindow()
			win32.MessageBox(hWnd, win32.StrToPwstr("Hello "+name+"!"),
				win32.StrToPwstr("@Host"), win32.MB_ICONINFORMATION)
			return nil
		},
		"getCount": func(args ...*ole.Variant) *ole.Variant {
			count += 1
			return ole.NewVar(count)
		},
	}, map[string]oleimpl.VariantProp{
		"Time": {Get: func(args ...*ole.Variant) *ole.Variant {
			sTime := time.Now().Format(time.Stamp)
			return ole.NewVar(sTime)
		}},
		"SubObject": {Get: func(args ...*ole.Variant) *ole.Variant {
			return ole.NewVar(pDispSubObject)
		}},
	}, func() {
		pDispSubObject.Release()
	})
}

func messageLoop() {
	var msg win32.MSG
	for {
		ret, _ := win32.GetMessage(&msg, 0, 0, 0)
		if ret == 0 {
			break
		}
		win32.TranslateMessage(&msg)
		win32.DispatchMessage(&msg)
	}
}

func WndProc(hWnd win32.HWND, message uint32,
	wParam win32.WPARAM, lParam win32.LPARAM) win32.LRESULT {
	switch message {
	case win32.WM_SIZE:
		if webviewController != nil {
			var bounds wv2.TagRECT
			win32.GetClientRect(hWnd, (*win32.RECT)(&bounds))
			webviewController.SetBounds(bounds)
		}
		return 0
	case win32.WM_DESTROY:
		win32.PostQuitMessage(0)
		return 0
	default:
		return win32.DefWindowProc(hWnd, message, wParam, lParam)
	}
}

func serveHttp() {
	html := `
	<html><head></head><body>
	<h1>Host Object test</h1>
	<style>ul {list-style: none; padding: 0;} li {margin: 1em;}</style>
	<ul>
		<li><button onclick="sayHello();">Say Hello</button></li>
		<li><button onclick="getCount();">Get Count</button></li>
		<li><button onclick="showTime();">Show Time</button></li>
		<li><button onclick="setSubObjectProp();">Set SubObject Prop</button></li>
		<li><button onclick="getSubObjectProp();">Get SubObject Prop</button></li>
		<li><button onclick="callSubObjectFunc();">Call SubObject Func</button></li>
	</ul>
	<script>
	var testObj = chrome.webview.hostObjects.sync.TestObj;
	function sayHello() {
		testObj.SayHello('World');
	}
	function getCount() {
		alert(testObj.getCount()); 
	}
	function showTime() {
		alert(testObj.Time); 
	}
	function setSubObjectProp() {
		let prop = prompt("Please enter sub object prop", "");
		if(prop != null) {
			testObj.subObject.prop = prop;
		}
	}
	function getSubObjectProp() {
		alert(testObj.subObject.prop);
	}
	function callSubObjectFunc() {
		alert(testObj.subObject.func('HELLO'));
	}
	</script>
	</body></html>`
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(html))
	})
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	server := &http.Server{}
	go server.Serve(listener)
}
