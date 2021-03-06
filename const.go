package wke

// #include "wke.h"
import "C"

type MouseFlags uint

const (
	DefaultAgent = "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"
)
const (
	MF_LBUTTON MouseFlags = C.WKE_LBUTTON
	MF_RBUTTON MouseFlags = C.WKE_RBUTTON
	MF_SHIFT   MouseFlags = C.WKE_SHIFT
	MF_CONTROL MouseFlags = C.WKE_CONTROL
	MF_MBUTTON MouseFlags = C.WKE_MBUTTON
)

type KeyFlags uint

const (
	KF_EXTENDED KeyFlags = C.WKE_EXTENDED
	KF_REPEAT   KeyFlags = C.WKE_REPEAT
)

type MouseMsg uint

const (
	MM_MOUSEMOVE     MouseMsg = C.WKE_MSG_MOUSEMOVE
	MM_LBUTTONDOWN   MouseMsg = C.WKE_MSG_LBUTTONDOWN
	MM_LBUTTONUP     MouseMsg = C.WKE_MSG_LBUTTONUP
	MM_LBUTTONDBLCLK MouseMsg = C.WKE_MSG_LBUTTONDBLCLK
	MM_RBUTTONDOWN   MouseMsg = C.WKE_MSG_RBUTTONDOWN
	MM_RBUTTONUP     MouseMsg = C.WKE_MSG_RBUTTONUP
	MM_RBUTTONDBLCLK MouseMsg = C.WKE_MSG_RBUTTONDBLCLK
	MM_MBUTTONDOWN   MouseMsg = C.WKE_MSG_MBUTTONDOWN
	MM_MBUTTONUP     MouseMsg = C.WKE_MSG_MBUTTONUP
	MM_MBUTTONDBLCLK MouseMsg = C.WKE_MSG_MBUTTONDBLCLK
	MM_MOUSEWHEEL    MouseMsg = C.WKE_MSG_MOUSEWHEEL
)
