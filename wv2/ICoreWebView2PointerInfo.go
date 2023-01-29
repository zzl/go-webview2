package wv2

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// E6995887-D10D-4F5D-9359-4CE46E4F96B9
var IID_ICoreWebView2PointerInfo = syscall.GUID{0xE6995887, 0xD10D, 0x4F5D,
	[8]byte{0x93, 0x59, 0x4C, 0xE4, 0x6E, 0x4F, 0x96, 0xB9}}

type ICoreWebView2PointerInfo struct {
	win32.IUnknown
}

func NewICoreWebView2PointerInfo(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2PointerInfo {
	if pUnk == nil {
		return nil
	}
	p := (*ICoreWebView2PointerInfo)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2PointerInfo) IID() *syscall.GUID {
	return &IID_ICoreWebView2PointerInfo
}

func (this *ICoreWebView2PointerInfo) GetPointerKind(pointerKind *uint32) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pointerKind)))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) SetPointerKind(pointerKind uint32) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(pointerKind))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) GetPointerId(pointerId *uint32) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pointerId)))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) SetPointerId(pointerId uint32) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(pointerId))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) GetFrameId(frameId *uint32) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(frameId)))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) SetFrameId(frameId uint32) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(frameId))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) GetPointerFlags(pointerFlags *uint32) com.Error {
	addr := (*this.LpVtbl)[9]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pointerFlags)))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) SetPointerFlags(pointerFlags uint32) com.Error {
	addr := (*this.LpVtbl)[10]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(pointerFlags))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) GetPointerDeviceRect(pointerDeviceRect *TagRECT) com.Error {
	addr := (*this.LpVtbl)[11]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pointerDeviceRect)))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) SetPointerDeviceRect(pointerDeviceRect TagRECT) com.Error {
	addr := (*this.LpVtbl)[12]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), (uintptr)(unsafe.Pointer(&pointerDeviceRect)))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) GetDisplayRect(displayRect *TagRECT) com.Error {
	addr := (*this.LpVtbl)[13]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(displayRect)))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) SetDisplayRect(displayRect TagRECT) com.Error {
	addr := (*this.LpVtbl)[14]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), (uintptr)(unsafe.Pointer(&displayRect)))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) GetPixelLocation(pixelLocation *TagPOINT) com.Error {
	addr := (*this.LpVtbl)[15]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pixelLocation)))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) SetPixelLocation(pixelLocation TagPOINT) com.Error {
	addr := (*this.LpVtbl)[16]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&pixelLocation)))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) GetHimetricLocation(himetricLocation *TagPOINT) com.Error {
	addr := (*this.LpVtbl)[17]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(himetricLocation)))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) SetHimetricLocation(himetricLocation TagPOINT) com.Error {
	addr := (*this.LpVtbl)[18]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&himetricLocation)))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) GetPixelLocationRaw(pixelLocationRaw *TagPOINT) com.Error {
	addr := (*this.LpVtbl)[19]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pixelLocationRaw)))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) SetPixelLocationRaw(pixelLocationRaw TagPOINT) com.Error {
	addr := (*this.LpVtbl)[20]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&pixelLocationRaw)))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) GetHimetricLocationRaw(himetricLocationRaw *TagPOINT) com.Error {
	addr := (*this.LpVtbl)[21]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(himetricLocationRaw)))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) SetHimetricLocationRaw(himetricLocationRaw TagPOINT) com.Error {
	addr := (*this.LpVtbl)[22]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&himetricLocationRaw)))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) GetTime(time *uint32) com.Error {
	addr := (*this.LpVtbl)[23]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(time)))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) SetTime(time uint32) com.Error {
	addr := (*this.LpVtbl)[24]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(time))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) GetHistoryCount(historyCount *uint32) com.Error {
	addr := (*this.LpVtbl)[25]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(historyCount)))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) SetHistoryCount(historyCount uint32) com.Error {
	addr := (*this.LpVtbl)[26]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(historyCount))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) GetInputData(inputData *int32) com.Error {
	addr := (*this.LpVtbl)[27]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(inputData)))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) SetInputData(inputData int32) com.Error {
	addr := (*this.LpVtbl)[28]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(inputData))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) GetKeyStates(keyStates *uint32) com.Error {
	addr := (*this.LpVtbl)[29]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(keyStates)))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) SetKeyStates(keyStates uint32) com.Error {
	addr := (*this.LpVtbl)[30]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(keyStates))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) GetPerformanceCount(performanceCount *uint64) com.Error {
	addr := (*this.LpVtbl)[31]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(performanceCount)))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) SetPerformanceCount(performanceCount uint64) com.Error {
	addr := (*this.LpVtbl)[32]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(performanceCount))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) GetButtonChangeKind(buttonChangeKind *int32) com.Error {
	addr := (*this.LpVtbl)[33]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(buttonChangeKind)))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) SetButtonChangeKind(buttonChangeKind int32) com.Error {
	addr := (*this.LpVtbl)[34]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(buttonChangeKind))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) GetPenFlags(penFlags *uint32) com.Error {
	addr := (*this.LpVtbl)[35]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(penFlags)))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) SetPenFlags(penFlags uint32) com.Error {
	addr := (*this.LpVtbl)[36]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(penFlags))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) GetPenMask(penMask *uint32) com.Error {
	addr := (*this.LpVtbl)[37]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(penMask)))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) SetPenMask(penMask uint32) com.Error {
	addr := (*this.LpVtbl)[38]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(penMask))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) GetPenPressure(penPressure *uint32) com.Error {
	addr := (*this.LpVtbl)[39]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(penPressure)))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) SetPenPressure(penPressure uint32) com.Error {
	addr := (*this.LpVtbl)[40]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(penPressure))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) GetPenRotation(penRotation *uint32) com.Error {
	addr := (*this.LpVtbl)[41]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(penRotation)))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) SetPenRotation(penRotation uint32) com.Error {
	addr := (*this.LpVtbl)[42]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(penRotation))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) GetPenTiltX(penTiltX *int32) com.Error {
	addr := (*this.LpVtbl)[43]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(penTiltX)))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) SetPenTiltX(penTiltX int32) com.Error {
	addr := (*this.LpVtbl)[44]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(penTiltX))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) GetPenTiltY(penTiltY *int32) com.Error {
	addr := (*this.LpVtbl)[45]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(penTiltY)))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) SetPenTiltY(penTiltY int32) com.Error {
	addr := (*this.LpVtbl)[46]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(penTiltY))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) GetTouchFlags(touchFlags *uint32) com.Error {
	addr := (*this.LpVtbl)[47]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(touchFlags)))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) SetTouchFlags(touchFlags uint32) com.Error {
	addr := (*this.LpVtbl)[48]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(touchFlags))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) GetTouchMask(touchMask *uint32) com.Error {
	addr := (*this.LpVtbl)[49]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(touchMask)))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) SetTouchMask(touchMask uint32) com.Error {
	addr := (*this.LpVtbl)[50]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(touchMask))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) GetTouchContact(touchContact *TagRECT) com.Error {
	addr := (*this.LpVtbl)[51]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(touchContact)))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) SetTouchContact(touchContact TagRECT) com.Error {
	addr := (*this.LpVtbl)[52]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), (uintptr)(unsafe.Pointer(&touchContact)))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) GetTouchContactRaw(touchContactRaw *TagRECT) com.Error {
	addr := (*this.LpVtbl)[53]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(touchContactRaw)))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) SetTouchContactRaw(touchContactRaw TagRECT) com.Error {
	addr := (*this.LpVtbl)[54]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), (uintptr)(unsafe.Pointer(&touchContactRaw)))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) GetTouchOrientation(touchOrientation *uint32) com.Error {
	addr := (*this.LpVtbl)[55]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(touchOrientation)))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) SetTouchOrientation(touchOrientation uint32) com.Error {
	addr := (*this.LpVtbl)[56]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(touchOrientation))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) GetTouchPressure(touchPressure *uint32) com.Error {
	addr := (*this.LpVtbl)[57]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(touchPressure)))
	return com.Error(ret)
}

func (this *ICoreWebView2PointerInfo) SetTouchPressure(touchPressure uint32) com.Error {
	addr := (*this.LpVtbl)[58]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(touchPressure))
	return com.Error(ret)
}
