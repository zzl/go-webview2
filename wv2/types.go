package wv2

import (
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
)

// struct __MIDL___MIDL_itf_webview2_0005_0001_0001
type MIDL___MIDL_itf_webview2_0005_0001_0001__ struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}

// struct COREWEBVIEW2_PHYSICAL_KEY_STATUS
type COREWEBVIEW2_PHYSICAL_KEY_STATUS struct {
	RepeatCount   uint32
	ScanCode      uint32
	IsExtendedKey int32
	IsMenuKeyDown int32
	WasKeyDown    int32
	IsKeyReleased int32
}

// struct tagRECT
type TagRECT struct {
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

// struct EventRegistrationToken
type EventRegistrationToken struct {
	Value int64
}

// alias wireHWND
type WireHWND = *RemotableHandle_

// struct _RemotableHandle
type RemotableHandle_ struct {
	FContext int32
	U        uintptr
}

// struct _LARGE_INTEGER
type LARGE_INTEGER_ struct {
	QuadPart int64
}

// struct _ULARGE_INTEGER
type ULARGE_INTEGER_ struct {
	QuadPart uint64
}

// struct tagSTATSTG
type TagSTATSTG struct {
	PwcsName          win32.PWSTR
	Type              uint32
	CbSize            ULARGE_INTEGER_
	Mtime             FILETIME_
	Ctime             FILETIME_
	Atime             FILETIME_
	GrfMode           uint32
	GrfLocksSupported uint32
	Clsid             syscall.GUID
	GrfStateBits      uint32
	Reserved          uint32
}

// struct _FILETIME
type FILETIME_ struct {
	DwLowDateTime  uint32
	DwHighDateTime uint32
}

// struct tagPOINT
type TagPOINT struct {
	X int32
	Y int32
}

// alias wireHICON
type WireHICON = *RemotableHandle_

// struct COREWEBVIEW2_COLOR
type COREWEBVIEW2_COLOR struct {
	A byte
	R byte
	G byte
	B byte
}

