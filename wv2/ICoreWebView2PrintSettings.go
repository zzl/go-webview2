package wv2

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 377F3721-C74E-48CA-8DB1-DF68E51D60E2
var IID_ICoreWebView2PrintSettings = syscall.GUID{0x377F3721, 0xC74E, 0x48CA,
	[8]byte{0x8D, 0xB1, 0xDF, 0x68, 0xE5, 0x1D, 0x60, 0xE2}}

type ICoreWebView2PrintSettings struct {
	win32.IUnknown
}

func NewICoreWebView2PrintSettings(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2PrintSettings {
	if pUnk == nil {
		return nil
	}
	p := (*ICoreWebView2PrintSettings)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2PrintSettings) IID() *syscall.GUID {
	return &IID_ICoreWebView2PrintSettings
}

func (this *ICoreWebView2PrintSettings) GetOrientation(orientation *int32) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(orientation)))
	return com.Error(ret)
}

func (this *ICoreWebView2PrintSettings) SetOrientation(orientation int32) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(orientation))
	return com.Error(ret)
}

func (this *ICoreWebView2PrintSettings) GetScaleFactor(scaleFactor *float64) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(scaleFactor)))
	return com.Error(ret)
}

func (this *ICoreWebView2PrintSettings) SetScaleFactor(scaleFactor float64) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(scaleFactor))
	return com.Error(ret)
}

func (this *ICoreWebView2PrintSettings) GetPageWidth(pageWidth *float64) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pageWidth)))
	return com.Error(ret)
}

func (this *ICoreWebView2PrintSettings) SetPageWidth(pageWidth float64) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(pageWidth))
	return com.Error(ret)
}

func (this *ICoreWebView2PrintSettings) GetPageHeight(pageHeight *float64) com.Error {
	addr := (*this.LpVtbl)[9]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pageHeight)))
	return com.Error(ret)
}

func (this *ICoreWebView2PrintSettings) SetPageHeight(pageHeight float64) com.Error {
	addr := (*this.LpVtbl)[10]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(pageHeight))
	return com.Error(ret)
}

func (this *ICoreWebView2PrintSettings) GetMarginTop(marginTop *float64) com.Error {
	addr := (*this.LpVtbl)[11]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(marginTop)))
	return com.Error(ret)
}

func (this *ICoreWebView2PrintSettings) SetMarginTop(marginTop float64) com.Error {
	addr := (*this.LpVtbl)[12]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(marginTop))
	return com.Error(ret)
}

func (this *ICoreWebView2PrintSettings) GetMarginBottom(marginBottom *float64) com.Error {
	addr := (*this.LpVtbl)[13]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(marginBottom)))
	return com.Error(ret)
}

func (this *ICoreWebView2PrintSettings) SetMarginBottom(marginBottom float64) com.Error {
	addr := (*this.LpVtbl)[14]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(marginBottom))
	return com.Error(ret)
}

func (this *ICoreWebView2PrintSettings) GetMarginLeft(marginLeft *float64) com.Error {
	addr := (*this.LpVtbl)[15]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(marginLeft)))
	return com.Error(ret)
}

func (this *ICoreWebView2PrintSettings) SetMarginLeft(marginLeft float64) com.Error {
	addr := (*this.LpVtbl)[16]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(marginLeft))
	return com.Error(ret)
}

func (this *ICoreWebView2PrintSettings) GetMarginRight(marginRight *float64) com.Error {
	addr := (*this.LpVtbl)[17]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(marginRight)))
	return com.Error(ret)
}

func (this *ICoreWebView2PrintSettings) SetMarginRight(marginRight float64) com.Error {
	addr := (*this.LpVtbl)[18]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(marginRight))
	return com.Error(ret)
}

func (this *ICoreWebView2PrintSettings) GetShouldPrintBackgrounds(shouldPrintBackgrounds *int32) com.Error {
	addr := (*this.LpVtbl)[19]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shouldPrintBackgrounds)))
	return com.Error(ret)
}

func (this *ICoreWebView2PrintSettings) SetShouldPrintBackgrounds(shouldPrintBackgrounds int32) com.Error {
	addr := (*this.LpVtbl)[20]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(shouldPrintBackgrounds))
	return com.Error(ret)
}

func (this *ICoreWebView2PrintSettings) GetShouldPrintSelectionOnly(shouldPrintSelectionOnly *int32) com.Error {
	addr := (*this.LpVtbl)[21]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shouldPrintSelectionOnly)))
	return com.Error(ret)
}

func (this *ICoreWebView2PrintSettings) SetShouldPrintSelectionOnly(shouldPrintSelectionOnly int32) com.Error {
	addr := (*this.LpVtbl)[22]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(shouldPrintSelectionOnly))
	return com.Error(ret)
}

func (this *ICoreWebView2PrintSettings) GetShouldPrintHeaderAndFooter(shouldPrintHeaderAndFooter *int32) com.Error {
	addr := (*this.LpVtbl)[23]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(shouldPrintHeaderAndFooter)))
	return com.Error(ret)
}

func (this *ICoreWebView2PrintSettings) SetShouldPrintHeaderAndFooter(shouldPrintHeaderAndFooter int32) com.Error {
	addr := (*this.LpVtbl)[24]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(shouldPrintHeaderAndFooter))
	return com.Error(ret)
}

func (this *ICoreWebView2PrintSettings) GetHeaderTitle(headerTitle *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[25]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(headerTitle)))
	return com.Error(ret)
}

func (this *ICoreWebView2PrintSettings) SetHeaderTitle(headerTitle string) com.Error {
	addr := (*this.LpVtbl)[26]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(headerTitle)))
	return com.Error(ret)
}

func (this *ICoreWebView2PrintSettings) GetFooterUri(footerUri *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[27]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(footerUri)))
	return com.Error(ret)
}

func (this *ICoreWebView2PrintSettings) SetFooterUri(footerUri string) com.Error {
	addr := (*this.LpVtbl)[28]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(footerUri)))
	return com.Error(ret)
}
