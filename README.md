# GO-WebView2

Use Microsoft Edge WebView2 Control excel in golang.

## Source code

Most of the code in this library is generated with [GO-TlbImp](https://github.com/zzl/go-tlbimp).

```go-tlbimp.exe -tlb "C:\Downloads\microsoft.web.webview2.1.0.1210.39\WebView2.tlb" -out-dir C:\Project\go-webview2\wv2```

The "microsoft.web.webview2.1.0.1210.39" folder is extracted from the Microsoft.Web.WebView2 nuget package.

(Here are the steps: download microsoft.web.webview2.xxx.nupkg from https://www.nuget.org/packages/Microsoft.Web.WebView2, 
rename to microsoft.web.webview2.xxx.zip, then unzip. I don't know if there exists a more formal way to do so.)

## Dependencies

Before you run an application that uses webview2, 
[the WebView2 runtime](https://developer.microsoft.com/en-us/microsoft-edge/webview2/#download-section) 
must be installed beforehand.

Besides, the WebView2Loader.dll(provided in the bin folder, or you can find it in the extracted nuget package)
should be copied into the same folder as the executable, or reside in the working directory.

## Examples
The [hello_webview](https://github.com/zzl/go-webview2/blob/main/examples/hello_webview/hello_webview.go) sample
is deliberately structured in a way to resemble the 
[Win32 Getting Started sample](https://github.com/MicrosoftEdge/WebView2Samples/blob/main/GettingStartedGuides/Win32_GettingStarted/HelloWebView.cpp)
from Microsoft. As you can see, there can be very little difference between the two, except for the syntax.

