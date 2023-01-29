package main

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-webview2/wv2"
	"github.com/zzl/go-win32api/v2/win32"
	"strings"
	"syscall"
	"unsafe"
)

// Global variables

// The main window class name.
const szWindowClass = "DesktopApp"

// The string that appears in the application's title bar.
const szTitle = "WebView sample"

var hInst win32.HINSTANCE

// Pointer to WebViewController
var webviewController *wv2.ICoreWebView2Controller

// Pointer to WebView window
var webviewWindow *wv2.ICoreWebView2

//
func main() {

	com.Initialize()

	globalScope := com.NewScope()
	defer globalScope.Leave()

	wv2.EnsureRuntimeAvailable()

	// Store instance handle in our global variable
	hInst, _ = win32.GetModuleHandle(nil)

	var wcex win32.WNDCLASSEX
	wcex.CbSize = uint32(unsafe.Sizeof(wcex))
	wcex.Style = win32.CS_HREDRAW | win32.CS_VREDRAW
	wcex.LpfnWndProc = syscall.NewCallback(WndProc)
	wcex.HInstance = hInst
	wcex.HCursor, _ = win32.LoadCursor(0, win32.IDC_ARROW)
	wcex.HbrBackground = win32.HBRUSH(win32.COLOR_WINDOW + 1)
	wcex.LpszClassName = win32.StrToPwstr(szWindowClass)

	if ret, _ := win32.RegisterClassEx(&wcex); ret == 0 {
		win32.MessageBox(0, win32.StrToPwstr("Call to RegisterClassEx failed!"),
			win32.StrToPwstr("Windows Desktop Guided Tour"), 0)
		return
	}

	hWnd, _ := win32.CreateWindowEx(0,
		win32.StrToPwstr(szWindowClass),
		win32.StrToPwstr(szTitle),
		win32.WS_OVERLAPPEDWINDOW,
		win32.CW_USEDEFAULT, win32.CW_USEDEFAULT,
		1200, 900,
		0,
		0,
		hInst,
		nil,
	)

	if hWnd == 0 {
		win32.MessageBox(0, win32.StrToPwstr("Call to CreateWindow failed!"),
			win32.StrToPwstr("Windows Desktop Guided Tour"), 0)
		return
	}

	win32.ShowWindow(hWnd, win32.SW_SHOW)
	win32.UpdateWindow(hWnd)

	// <-- WebView2 sample code starts here -->
	// Step 3 - Create a single WebView within the parent window
	// Locate the browser and set up the environment for WebView
	wv2.CreateCoreWebView2EnvironmentWithOptions("", "", nil,
		wv2.NewICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerByFunc(
			func(errorCode com.Error, env *wv2.ICoreWebView2Environment) com.Error {
				// Create a CoreWebView2Controller and get the associated CoreWebView2
				// whose parent is the main window hWnd
				env.CreateCoreWebView2Controller(hWnd,
					wv2.NewICoreWebView2CreateCoreWebView2ControllerCompletedHandlerByFunc(
						func(errorCode com.Error, controller *wv2.ICoreWebView2Controller) com.Error {
							defer com.NewScope().Leave()
							if controller != nil {
								webviewController = controller
								globalScope.AddComPtr(webviewController, true)
								webviewController.GetCoreWebView2(&webviewWindow)
								globalScope.AddComPtr(webviewWindow, true)
							}
							// Add a few settings for the webview
							// The demo step is redundant since the values are the default settings
							var settings *wv2.ICoreWebView2Settings
							webviewWindow.GetSettings(&settings)
							settings.SetIsScriptEnabled(win32.TRUE)
							settings.SetIsWebMessageEnabled(win32.TRUE)
							settings.SetAreDefaultScriptDialogsEnabled(win32.TRUE)

							// Resize WebView to fit the bounds of the parent window
							var bounds wv2.TagRECT
							win32.GetClientRect(hWnd, (*win32.RECT)(&bounds))
							webviewController.SetBounds(bounds)

							// Schedule an async task to navigate to Bing
							webviewWindow.Navigate("https://www.bing.com/")

							// Step 4 - Navigation events
							// register an ICoreWebView2NavigationStartingEventHandler to cancel any non-https navigation
							var token wv2.EventRegistrationToken
							webviewWindow.Add_NavigationStarting(
								wv2.NewICoreWebView2NavigationStartingEventHandlerByFunc(
									func(sender *wv2.ICoreWebView2,
										args *wv2.ICoreWebView2NavigationStartingEventArgs) com.Error {
										var pwszUri win32.PWSTR
										args.GetUri(&pwszUri)
										uri := win32.PwstrToStr(pwszUri)
										if !strings.HasPrefix(uri, "https") {
											args.SetCancel(1)
										}
										win32.CoTaskMemFree(unsafe.Pointer(pwszUri))
										return com.Error(win32.S_OK)
									}, true), &token)

							// Step 5 - Scripting
							// Schedule an async task to add initialization script that freezes the Object object
							webviewWindow.AddScriptToExecuteOnDocumentCreated("Object.freeze(Object);", nil)
							// Schedule an async task to get the document URL
							webviewWindow.ExecuteScript("window.document.URL;", wv2.NewICoreWebView2ExecuteScriptCompletedHandlerByFunc(
								func(errorCode com.Error, resultObjectAsJson string) com.Error {
									url := resultObjectAsJson
									_ = url //doSomethingWithURL(URL);
									return com.Error(win32.S_OK)
								}, true))

							// Step 6 - Communication between host and web content
							// Set an event handler for the host to return received message back to the web content
							webviewWindow.Add_WebMessageReceived(
								wv2.NewICoreWebView2WebMessageReceivedEventHandlerByFunc(
									func(webview *wv2.ICoreWebView2,
										args *wv2.ICoreWebView2WebMessageReceivedEventArgs) com.Error {
										var message win32.PWSTR
										args.TryGetWebMessageAsString(&message)
										// processMessage(&message);
										webview.PostWebMessageAsString(win32.BstrToStr(message))
										win32.CoTaskMemFree(unsafe.Pointer(message))
										return com.Error(win32.S_OK)
									}, true), &token)

							// Schedule an async task to add initialization script that
							// 1) Add an listener to print message from the host
							// 2) Post document URL to the host
							webviewWindow.AddScriptToExecuteOnDocumentCreated(
								"window.chrome.webview.addEventListener('message', event => alert(event.data));"+
									"window.chrome.webview.postMessage(window.document.URL);",
								nil)
							return com.Error(win32.S_OK)
						}, true))
				return com.Error(win32.S_OK)
			}, true))

	var msg win32.MSG
	for {
		ret, _ := win32.GetMessage(&msg, 0, 0, 0)
		if ret == 0 {
			break
		}
		win32.TranslateMessage(&msg)
		win32.DispatchMessage(&msg)
	}

	println("?")
}

func WndProc(hWnd win32.HWND, message uint32,
	wParam win32.WPARAM, lParam win32.LPARAM) win32.LRESULT {
	//greeting = "Hello, Windows desktop!"
	switch message {
	case win32.WM_SIZE:
		if webviewController != nil {
			var bounds wv2.TagRECT
			win32.GetClientRect(hWnd, (*win32.RECT)(&bounds))
			webviewController.SetBounds(bounds)
		}
		break
	case win32.WM_DESTROY:
		win32.PostQuitMessage(0)
		break
	default:
		return win32.DefWindowProc(hWnd, message, wParam, lParam)
	}
	return 0
}
