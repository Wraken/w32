// Copyright 2010-2012 The W32 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package w32

import (
	// #include <wtypes.h>
	// #include <winable.h>
	"C"
	"encoding/binary"
	"errors"
	"fmt"

	//"regexp"
	"syscall"
	"unicode/utf8"
	"unsafe"
)

var (
	moduser32 = syscall.NewLazyDLL("user32.dll")

	procAddClipboardFormatListener    = moduser32.NewProc("AddClipboardFormatListener")
	procAdjustWindowRect              = moduser32.NewProc("AdjustWindowRect")
	procAdjustWindowRectEx            = moduser32.NewProc("AdjustWindowRectEx")
	procBeginPaint                    = moduser32.NewProc("BeginPaint")
	procCallNextHookEx                = moduser32.NewProc("CallNextHookEx")
	procCallWindowProc                = moduser32.NewProc("CallWindowProcW")
	procChangeDisplaySettingsEx       = moduser32.NewProc("ChangeDisplaySettingsExW")
	procClientToScreen                = moduser32.NewProc("ClientToScreen")
	procCloseClipboard                = moduser32.NewProc("CloseClipboard")
	procCopyRect                      = moduser32.NewProc("CopyRect")
	procCreateDialogParam             = moduser32.NewProc("CreateDialogParamW")
	procCreateIcon                    = moduser32.NewProc("CreateIcon")
	procCreateWindowEx                = moduser32.NewProc("CreateWindowExW")
	procDefDlgProc                    = moduser32.NewProc("DefDlgProcW")
	procDefWindowProc                 = moduser32.NewProc("DefWindowProcW")
	procDestroyIcon                   = moduser32.NewProc("DestroyIcon")
	procDestroyWindow                 = moduser32.NewProc("DestroyWindow")
	procDialogBoxParam                = moduser32.NewProc("DialogBoxParamW")
	procDispatchMessage               = moduser32.NewProc("DispatchMessageW")
	procDrawIcon                      = moduser32.NewProc("DrawIcon")
	procDrawText                      = moduser32.NewProc("DrawTextW")
	procEmptyClipboard                = moduser32.NewProc("EmptyClipboard")
	procEnableWindow                  = moduser32.NewProc("EnableWindow")
	procEndDialog                     = moduser32.NewProc("EndDialog")
	procEndPaint                      = moduser32.NewProc("EndPaint")
	procEnumChildWindows              = moduser32.NewProc("EnumChildWindows")
	procEnumClipboardFormats          = moduser32.NewProc("EnumClipboardFormats")
	procEnumDisplayMonitors           = moduser32.NewProc("EnumDisplayMonitors")
	procEnumDisplaySettingsEx         = moduser32.NewProc("EnumDisplaySettingsExW")
	procEqualRect                     = moduser32.NewProc("EqualRect")
	procFillRect                      = moduser32.NewProc("FillRect")
	procFindWindowExW                 = moduser32.NewProc("FindWindowExW")
	procFindWindowW                   = moduser32.NewProc("FindWindowW")
	procGetAsyncKeyState              = moduser32.NewProc("GetAsyncKeyState")
	procGetClassName                  = moduser32.NewProc("GetClassNameW")
	procGetClientRect                 = moduser32.NewProc("GetClientRect")
	procGetClipboardData              = moduser32.NewProc("GetClipboardData")
	procGetClipboardFormatName        = moduser32.NewProc("GetClipboardFormatNameW")
	procGetCursorPos                  = moduser32.NewProc("GetCursorPos")
	procGetDC                         = moduser32.NewProc("GetDC")
	procGetDlgItem                    = moduser32.NewProc("GetDlgItem")
	procGetForegroundWindow           = moduser32.NewProc("GetForegroundWindow")
	procGetKeyNameText                = moduser32.NewProc("GetKeyNameTextW")
	procGetKeyState                   = moduser32.NewProc("GetKeyState")
	procGetKeyboardState              = moduser32.NewProc("GetKeyboardState")
	procGetMessage                    = moduser32.NewProc("GetMessageW")
	procGetMonitorInfo                = moduser32.NewProc("GetMonitorInfoW")
	procGetSystemMetrics              = moduser32.NewProc("GetSystemMetrics")
	procGetWindowLong                 = moduser32.NewProc("GetWindowLongW")
	procGetWindowLongPtr              = moduser32.NewProc("GetWindowLongW")
	procGetWindowRect                 = moduser32.NewProc("GetWindowRect")
	procGetWindowText                 = moduser32.NewProc("GetWindowTextW")
	procGetWindowTextLength           = moduser32.NewProc("GetWindowTextLengthW")
	procGetWindowTextW                = moduser32.NewProc("GetWindowTextW")
	procGetWindowThreadProcessId      = moduser32.NewProc("GetWindowThreadProcessId")
	procInflateRect                   = moduser32.NewProc("InflateRect")
	procIntersectRect                 = moduser32.NewProc("IntersectRect")
	procInvalidateRect                = moduser32.NewProc("InvalidateRect")
	procIsClipboardFormatAvailable    = moduser32.NewProc("IsClipboardFormatAvailable")
	procIsDialogMessage               = moduser32.NewProc("IsDialogMessageW")
	procIsRectEmpty                   = moduser32.NewProc("IsRectEmpty")
	procIsWindow                      = moduser32.NewProc("IsWindow")
	procIsWindowEnabled               = moduser32.NewProc("IsWindowEnabled")
	procIsWindowVisible               = moduser32.NewProc("IsWindowVisible")
	procKillTimer                     = moduser32.NewProc("KillTimer")
	procLoadCursor                    = moduser32.NewProc("LoadCursorW")
	procLoadIcon                      = moduser32.NewProc("LoadIconW")
	procMapVirtualKey                 = moduser32.NewProc("MapVirtualKeyExW")
	procMessageBox                    = moduser32.NewProc("MessageBoxW")
	procMonitorFromPoint              = moduser32.NewProc("MonitorFromPoint")
	procMonitorFromRect               = moduser32.NewProc("MonitorFromRect")
	procMonitorFromWindow             = moduser32.NewProc("MonitorFromWindow")
	procMoveWindow                    = moduser32.NewProc("MoveWindow")
	procOffsetRect                    = moduser32.NewProc("OffsetRect")
	procOpenClipboard                 = moduser32.NewProc("OpenClipboard")
	procPeekMessage                   = moduser32.NewProc("PeekMessageW")
	procPostMessage                   = moduser32.NewProc("PostMessageW")
	procPostQuitMessage               = moduser32.NewProc("PostQuitMessage")
	procPrintWindow                   = moduser32.NewProc("PrintWindow")
	procPtInRect                      = moduser32.NewProc("PtInRect")
	procRegisterClassEx               = moduser32.NewProc("RegisterClassExW")
	procRegisterHotKey                = moduser32.NewProc("RegisterHotKey")
	procReleaseCapture                = moduser32.NewProc("ReleaseCapture")
	procReleaseDC                     = moduser32.NewProc("ReleaseDC")
	procRemoveClipboardFormatListener = moduser32.NewProc("RemoveClipboardFormatListener")
	procScreenToClient                = moduser32.NewProc("ScreenToClient")
	procSendInput                     = moduser32.NewProc("SendInput")
	procSendMessage                   = moduser32.NewProc("SendMessageW")
	procSendMessageTimeout            = moduser32.NewProc("SendMessageTimeoutW")
	procSetCapture                    = moduser32.NewProc("SetCapture")
	procSetClipboardData              = moduser32.NewProc("SetClipboardData")
	procSetCursor                     = moduser32.NewProc("SetCursor")
	procSetCursorPos                  = moduser32.NewProc("SetCursorPos")
	procSetFocus                      = moduser32.NewProc("SetFocus")
	procSetForegroundWindow           = moduser32.NewProc("SetForegroundWindow")
	procSetRect                       = moduser32.NewProc("SetRect")
	procSetRectEmpty                  = moduser32.NewProc("SetRectEmpty")
	procSetTimer                      = moduser32.NewProc("SetTimer")
	procSetWinEventHook               = moduser32.NewProc("SetWinEventHook")
	procSetWindowLong                 = moduser32.NewProc("SetWindowLongW")
	procSetWindowLongPtr              = moduser32.NewProc("SetWindowLongW")
	procSetWindowPos                  = moduser32.NewProc("SetWindowPos")
	procSetWindowText                 = moduser32.NewProc("SetWindowTextW")
	procSetWindowsHookEx              = moduser32.NewProc("SetWindowsHookExW")
	procShowWindow                    = moduser32.NewProc("ShowWindow")
	procSubtractRect                  = moduser32.NewProc("SubtractRect")
	procSwapMouseButton               = moduser32.NewProc("SwapMouseButton")
	procToAscii                       = moduser32.NewProc("ToAscii")
	procTranslateAccelerator          = moduser32.NewProc("TranslateAcceleratorW")
	procTranslateMessage              = moduser32.NewProc("TranslateMessage")
	procUnhookWinEvent                = moduser32.NewProc("UnhookWinEvent")
	procUnhookWindowsHookEx           = moduser32.NewProc("UnhookWindowsHookEx")
	procUnionRect                     = moduser32.NewProc("UnionRect")
	procUnregisterHotKey              = moduser32.NewProc("UnregisterHotKey")
	procUpdateWindow                  = moduser32.NewProc("UpdateWindow")
	procVkKeyScanW                    = moduser32.NewProc("VkKeyScanW")
	procVkKeyScanExW                  = moduser32.NewProc("VkKeyScanExW")
	procWaitMessage                   = moduser32.NewProc("WaitMessage")
	setProcessDPIAware                = moduser32.NewProc("SetProcessDPIAware")
)

func SendMessageTimeout(hwnd HWND, msg uint32, wParam, lParam uintptr, fuFlags, uTimeout uint32, lpdwResult uintptr) uintptr {
	ret, _, _ := procSendMessageTimeout.Call(
		uintptr(hwnd),
		uintptr(msg),
		wParam,
		lParam,
		uintptr(fuFlags),
		uintptr(uTimeout),
		lpdwResult)

	return ret
}

// https://github.com/AllenDang/w32/pull/62/commits/bf59645b86663a54dffb94ca82683cc0610a6de3
func GetClassNameW(hwnd HWND) string {
	buf := make([]uint16, 255)
	procGetClassName.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(255))

	return syscall.UTF16ToString(buf)
}

// https://github.com/AllenDang/w32/pull/62/commits/bf59645b86663a54dffb94ca82683cc0610a6de3
func SetForegroundWindow(hwnd HWND) bool {
	ret, _, _ := procSetForegroundWindow.Call(
		uintptr(hwnd))

	return ret != 0
}

// https://github.com/AllenDang/w32/pull/62/commits/bf59645b86663a54dffb94ca82683cc0610a6de3
func FindWindowExW(hwndParent, hwndChildAfter HWND, className, windowName *uint16) HWND {
	ret, _, _ := procFindWindowExW.Call(
		uintptr(hwndParent),
		uintptr(hwndChildAfter),
		uintptr(unsafe.Pointer(className)),
		uintptr(unsafe.Pointer(windowName)))

	return HWND(ret)
}

func FindWindowExS(hwndParent, hwndChildAfter HWND, className, windowName *string) HWND {
	var class *uint16 = nil
	if className != nil {
		class = syscall.StringToUTF16Ptr(*className)
	}
	var window *uint16 = nil
	if windowName != nil {
		window = syscall.StringToUTF16Ptr(*windowName)
	}
	return FindWindowExW(hwndParent, hwndChildAfter, class, window)
}

// https://github.com/AllenDang/w32/pull/62/commits/bf59645b86663a54dffb94ca82683cc0610a6de3
func FindWindowW(className, windowName *uint16) HWND {
	ret, _, _ := procFindWindowW.Call(
		uintptr(unsafe.Pointer(className)),
		uintptr(unsafe.Pointer(windowName)))

	return HWND(ret)
}

func FindWindowS(className, windowName *string) HWND {
	var class *uint16 = nil
	if className != nil {
		class = syscall.StringToUTF16Ptr(*className)
	}
	var window *uint16 = nil
	if windowName != nil {
		window = syscall.StringToUTF16Ptr(*windowName)
	}
	return FindWindowW(class, window)
}

// https://github.com/AllenDang/w32/pull/62/commits/bf59645b86663a54dffb94ca82683cc0610a6de3
func EnumChildWindows(hWndParent HWND, lpEnumFunc WNDENUMPROC, lParam LPARAM) bool {
	ret, _, _ := procEnumChildWindows.Call(
		uintptr(hWndParent),
		uintptr(syscall.NewCallback(lpEnumFunc)),
		uintptr(lParam),
	)

	return ret != 0
}

func GetWindowTextW(hwnd syscall.Handle, str *uint16, maxCount int32) (len int32, err error) {
	r0, _, e1 := syscall.Syscall(procGetWindowTextW.Addr(), 3, uintptr(hwnd), uintptr(unsafe.Pointer(str)), uintptr(maxCount))
	len = int32(r0)
	if len == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func GetForegroundWindow() (hwnd syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procGetForegroundWindow.Addr(), 0, 0, 0, 0)
	if e1 != 0 {
		err = error(e1)
		return
	}
	hwnd = syscall.Handle(r0)
	return
}

func RegisterClassEx(wndClassEx *WNDCLASSEX) ATOM {
	ret, _, _ := procRegisterClassEx.Call(uintptr(unsafe.Pointer(wndClassEx)))
	return ATOM(ret)
}

func LoadIcon(instance HINSTANCE, iconName *uint16) HICON {
	ret, _, _ := procLoadIcon.Call(
		uintptr(instance),
		uintptr(unsafe.Pointer(iconName)))

	return HICON(ret)

}

func LoadIconS(instance HINSTANCE, iconName *string) HICON {
	var icon *uint16 = nil
	if iconName != nil {
		icon = syscall.StringToUTF16Ptr(*iconName)
	}
	return LoadIcon(instance, icon)
}

func LoadCursor(instance HINSTANCE, cursorName *uint16) HCURSOR {
	ret, _, _ := procLoadCursor.Call(
		uintptr(instance),
		uintptr(unsafe.Pointer(cursorName)))

	return HCURSOR(ret)

}

func LoadCursorS(instance HINSTANCE, cursorName *string) HCURSOR {
	var cursor *uint16 = nil
	if cursorName != nil {
		cursor = syscall.StringToUTF16Ptr(*cursorName)
	}
	return LoadCursor(instance, cursor)
}

func ShowWindow(hwnd HWND, cmdshow int) bool {
	ret, _, _ := procShowWindow.Call(
		uintptr(hwnd),
		uintptr(cmdshow))

	return ret != 0

}

func UpdateWindow(hwnd HWND) bool {
	ret, _, _ := procUpdateWindow.Call(
		uintptr(hwnd))
	return ret != 0
}

func CreateWindowEx(exStyle uint, className, windowName *uint16,
	style uint, x, y, width, height int, parent HWND, menu HMENU,
	instance HINSTANCE, param unsafe.Pointer) HWND {
	ret, _, _ := procCreateWindowEx.Call(
		uintptr(exStyle),
		uintptr(unsafe.Pointer(className)),
		uintptr(unsafe.Pointer(windowName)),
		uintptr(style),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(parent),
		uintptr(menu),
		uintptr(instance),
		uintptr(param))

	return HWND(ret)
}

func CreateWindowExS(exStyle uint, className, windowName *string,
	style uint, x, y, width, height int, parent HWND, menu HMENU,
	instance HINSTANCE, param unsafe.Pointer) HWND {
	var class *uint16 = nil
	if className != nil {
		class = syscall.StringToUTF16Ptr(*className)
	}
	var window *uint16 = nil
	if windowName != nil {
		window = syscall.StringToUTF16Ptr(*windowName)
	}
	return CreateWindowEx(
		exStyle,
		class,
		window,
		style,
		x,
		y,
		width,
		height,
		parent,
		menu,
		instance,
		param,
	)
}

func AdjustWindowRectEx(rect *RECT, style uint, menu bool, exStyle uint) bool {
	ret, _, _ := procAdjustWindowRectEx.Call(
		uintptr(unsafe.Pointer(rect)),
		uintptr(style),
		uintptr(BoolToBOOL(menu)),
		uintptr(exStyle))

	return ret != 0
}

func AdjustWindowRect(rect *RECT, style uint, menu bool) bool {
	ret, _, _ := procAdjustWindowRect.Call(
		uintptr(unsafe.Pointer(rect)),
		uintptr(style),
		uintptr(BoolToBOOL(menu)))

	return ret != 0
}

func DestroyWindow(hwnd HWND) bool {
	ret, _, _ := procDestroyWindow.Call(
		uintptr(hwnd))

	return ret != 0
}

func DefWindowProc(hwnd HWND, msg uint32, wParam, lParam uintptr) uintptr {
	ret, _, _ := procDefWindowProc.Call(
		uintptr(hwnd),
		uintptr(msg),
		wParam,
		lParam)

	return ret
}

func DefDlgProc(hwnd HWND, msg uint32, wParam, lParam uintptr) uintptr {
	ret, _, _ := procDefDlgProc.Call(
		uintptr(hwnd),
		uintptr(msg),
		wParam,
		lParam)

	return ret
}

func PostQuitMessage(exitCode int) {
	procPostQuitMessage.Call(
		uintptr(exitCode))
}

func PrintWindow(hwnd HWND, hdc HDC, nFlags uint32) int {
	ret, _, _ := procPrintWindow.Call(
		uintptr(hwnd),
		uintptr(hdc),
		uintptr(nFlags),
	)

	return int(ret)
}

func GetMessage(msg *MSG, hwnd HWND, msgFilterMin, msgFilterMax uint32) int {
	ret, _, _ := procGetMessage.Call(
		uintptr(unsafe.Pointer(msg)),
		uintptr(hwnd),
		uintptr(msgFilterMin),
		uintptr(msgFilterMax))

	return int(ret)
}

func TranslateMessage(msg *MSG) bool {
	ret, _, _ := procTranslateMessage.Call(
		uintptr(unsafe.Pointer(msg)))

	return ret != 0

}

func DispatchMessage(msg *MSG) uintptr {
	ret, _, _ := procDispatchMessage.Call(
		uintptr(unsafe.Pointer(msg)))

	return ret

}

func SendMessage(hwnd HWND, msg uint32, wParam, lParam uintptr) uintptr {
	ret, _, _ := procSendMessage.Call(
		uintptr(hwnd),
		uintptr(msg),
		wParam,
		lParam)

	return ret
}

func PostMessage(hwnd HWND, msg uint32, wParam, lParam uintptr) bool {
	ret, _, _ := procPostMessage.Call(
		uintptr(hwnd),
		uintptr(msg),
		wParam,
		lParam)

	return ret != 0
}

func WaitMessage() bool {
	ret, _, _ := procWaitMessage.Call()
	return ret != 0
}

func SetWindowText(hwnd HWND, text string) {
	procSetWindowText.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))))
}

func GetWindowTextLength(hwnd HWND) int {
	ret, _, _ := procGetWindowTextLength.Call(
		uintptr(hwnd))

	return int(ret)
}

func GetWindowText(hwnd HWND) string {
	textLen := GetWindowTextLength(hwnd) + 1

	buf := make([]uint16, textLen)
	procGetWindowText.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(textLen))

	return syscall.UTF16ToString(buf)
}

func GetWindowRect(hwnd HWND) *RECT {
	var rect RECT
	procGetWindowRect.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&rect)))

	return &rect
}

func MoveWindow(hwnd HWND, x, y, width, height int, repaint bool) bool {
	ret, _, _ := procMoveWindow.Call(
		uintptr(hwnd),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(BoolToBOOL(repaint)))

	return ret != 0

}

func ScreenToClient(hwnd HWND, x, y int) (X, Y int, ok bool) {
	pt := POINT{X: int32(x), Y: int32(y)}
	ret, _, _ := procScreenToClient.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&pt)))

	return int(pt.X), int(pt.Y), ret != 0
}

func CallWindowProc(preWndProc uintptr, hwnd HWND, msg uint32, wParam, lParam uintptr) uintptr {
	ret, _, _ := procCallWindowProc.Call(
		preWndProc,
		uintptr(hwnd),
		uintptr(msg),
		wParam,
		lParam)

	return ret
}

func SetWindowLong(hwnd HWND, index int, value uint32) uint32 {
	ret, _, _ := procSetWindowLong.Call(
		uintptr(hwnd),
		uintptr(index),
		uintptr(value))

	return uint32(ret)
}

func SetWindowLongPtr(hwnd HWND, index int, value uintptr) uintptr {
	ret, _, _ := procSetWindowLongPtr.Call(
		uintptr(hwnd),
		uintptr(index),
		value)

	return ret
}

func GetWindowLong(hwnd HWND, index int) int32 {
	ret, _, _ := procGetWindowLong.Call(
		uintptr(hwnd),
		uintptr(index))

	return int32(ret)
}

func GetWindowLongPtr(hwnd HWND, index int) uintptr {
	ret, _, _ := procGetWindowLongPtr.Call(
		uintptr(hwnd),
		uintptr(index))

	return ret
}

func EnableWindow(hwnd HWND, b bool) bool {
	ret, _, _ := procEnableWindow.Call(
		uintptr(hwnd),
		uintptr(BoolToBOOL(b)))
	return ret != 0
}

func IsWindowEnabled(hwnd HWND) bool {
	ret, _, _ := procIsWindowEnabled.Call(
		uintptr(hwnd))

	return ret != 0
}

func IsWindowVisible(hwnd HWND) bool {
	ret, _, _ := procIsWindowVisible.Call(
		uintptr(hwnd))

	return ret != 0
}

func SetFocus(hwnd HWND) HWND {
	ret, _, _ := procSetFocus.Call(
		uintptr(hwnd))

	return HWND(ret)
}

func InvalidateRect(hwnd HWND, rect *RECT, erase bool) bool {
	ret, _, _ := procInvalidateRect.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(rect)),
		uintptr(BoolToBOOL(erase)))

	return ret != 0
}

func GetClientRect(hwnd HWND) *RECT {
	var rect RECT
	ret, _, _ := procGetClientRect.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&rect)))

	if ret == 0 {
		panic(fmt.Sprintf("GetClientRect(%d) failed", hwnd))
	}

	return &rect
}

func GetDC(hwnd HWND) HDC {
	ret, _, _ := procGetDC.Call(
		uintptr(hwnd))

	return HDC(ret)
}

func ReleaseDC(hwnd HWND, hDC HDC) bool {
	ret, _, _ := procReleaseDC.Call(
		uintptr(hwnd),
		uintptr(hDC))

	return ret != 0
}

func SetCapture(hwnd HWND) HWND {
	ret, _, _ := procSetCapture.Call(
		uintptr(hwnd))

	return HWND(ret)
}

func ReleaseCapture() bool {
	ret, _, _ := procReleaseCapture.Call()

	return ret != 0
}

func GetWindowThreadProcessId(hwnd HWND) (HANDLE, int) {
	var processId int
	ret, _, _ := procGetWindowThreadProcessId.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&processId)))

	return HANDLE(ret), processId
}

func MessageBox(hwnd HWND, text, caption string, flags uint) int {
	ret, _, _ := procMessageBox.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(caption))),
		uintptr(flags))

	return int(ret)
}

func GetSystemMetrics(index int) int {
	ret, _, _ := procGetSystemMetrics.Call(
		uintptr(index))

	return int(ret)
}

func CopyRect(dst, src *RECT) bool {
	ret, _, _ := procCopyRect.Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src)))

	return ret != 0
}

func EqualRect(rect1, rect2 *RECT) bool {
	ret, _, _ := procEqualRect.Call(
		uintptr(unsafe.Pointer(rect1)),
		uintptr(unsafe.Pointer(rect2)))

	return ret != 0
}

func InflateRect(rect *RECT, dx, dy int) bool {
	ret, _, _ := procInflateRect.Call(
		uintptr(unsafe.Pointer(rect)),
		uintptr(dx),
		uintptr(dy))

	return ret != 0
}

func IntersectRect(dst, src1, src2 *RECT) bool {
	ret, _, _ := procIntersectRect.Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src1)),
		uintptr(unsafe.Pointer(src2)))

	return ret != 0
}

func IsRectEmpty(rect *RECT) bool {
	ret, _, _ := procIsRectEmpty.Call(
		uintptr(unsafe.Pointer(rect)))

	return ret != 0
}

func OffsetRect(rect *RECT, dx, dy int) bool {
	ret, _, _ := procOffsetRect.Call(
		uintptr(unsafe.Pointer(rect)),
		uintptr(dx),
		uintptr(dy))

	return ret != 0
}

func PtInRect(rect *RECT, x, y int) bool {
	pt := POINT{X: int32(x), Y: int32(y)}
	ret, _, _ := procPtInRect.Call(
		uintptr(unsafe.Pointer(rect)),
		uintptr(unsafe.Pointer(&pt)))

	return ret != 0
}

func SetRect(rect *RECT, left, top, right, bottom int) bool {
	ret, _, _ := procSetRect.Call(
		uintptr(unsafe.Pointer(rect)),
		uintptr(left),
		uintptr(top),
		uintptr(right),
		uintptr(bottom))

	return ret != 0
}

func SetRectEmpty(rect *RECT) bool {
	ret, _, _ := procSetRectEmpty.Call(
		uintptr(unsafe.Pointer(rect)))

	return ret != 0
}

func SubtractRect(dst, src1, src2 *RECT) bool {
	ret, _, _ := procSubtractRect.Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src1)),
		uintptr(unsafe.Pointer(src2)))

	return ret != 0
}

func UnionRect(dst, src1, src2 *RECT) bool {
	ret, _, _ := procUnionRect.Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src1)),
		uintptr(unsafe.Pointer(src2)))

	return ret != 0
}

func CreateDialog(hInstance HINSTANCE, lpTemplate *uint16, hWndParent HWND, lpDialogProc uintptr) HWND {
	ret, _, _ := procCreateDialogParam.Call(
		uintptr(hInstance),
		uintptr(unsafe.Pointer(lpTemplate)),
		uintptr(hWndParent),
		lpDialogProc,
		0)

	return HWND(ret)
}

func DialogBox(hInstance HINSTANCE, lpTemplateName *uint16, hWndParent HWND, lpDialogProc uintptr) int {
	ret, _, _ := procDialogBoxParam.Call(
		uintptr(hInstance),
		uintptr(unsafe.Pointer(lpTemplateName)),
		uintptr(hWndParent),
		lpDialogProc,
		0)

	return int(ret)
}

func GetDlgItem(hDlg HWND, nIDDlgItem int) HWND {
	ret, _, _ := procGetDlgItem.Call(
		uintptr(unsafe.Pointer(hDlg)),
		uintptr(nIDDlgItem))

	return HWND(ret)
}

func DrawIcon(hDC HDC, x, y int, hIcon HICON) bool {
	ret, _, _ := procDrawIcon.Call(
		uintptr(unsafe.Pointer(hDC)),
		uintptr(x),
		uintptr(y),
		uintptr(unsafe.Pointer(hIcon)))

	return ret != 0
}

func ClientToScreen(hwnd HWND, x, y int) (int, int) {
	pt := POINT{X: int32(x), Y: int32(y)}

	procClientToScreen.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&pt)))

	return int(pt.X), int(pt.Y)
}

func IsDialogMessage(hwnd HWND, msg *MSG) bool {
	ret, _, _ := procIsDialogMessage.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(msg)))

	return ret != 0
}

func IsWindow(hwnd HWND) bool {
	ret, _, _ := procIsWindow.Call(
		uintptr(hwnd))

	return ret != 0
}

func EndDialog(hwnd HWND, nResult uintptr) bool {
	ret, _, _ := procEndDialog.Call(
		uintptr(hwnd),
		nResult)

	return ret != 0
}

func PeekMessage(hwnd HWND, wMsgFilterMin, wMsgFilterMax, wRemoveMsg uint32) (msg MSG, err error) {
	_, _, err = procPeekMessage.Call(
		uintptr(unsafe.Pointer(&msg)),
		uintptr(hwnd),
		uintptr(wMsgFilterMin),
		uintptr(wMsgFilterMax),
		uintptr(wRemoveMsg))

	if !IsErrSuccess(err) {
		return
	}
	err = nil
	return
}

func TranslateAccelerator(hwnd HWND, hAccTable HACCEL, lpMsg *MSG) bool {
	ret, _, _ := procTranslateAccelerator.Call(
		uintptr(hwnd),
		uintptr(hAccTable),
		uintptr(unsafe.Pointer(lpMsg)))

	return ret != 0
}

func SetWindowPos(hwnd, hWndInsertAfter HWND, x, y, cx, cy int, uFlags uint) bool {
	ret, _, _ := procSetWindowPos.Call(
		uintptr(hwnd),
		uintptr(hWndInsertAfter),
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(uFlags))

	return ret != 0
}

func FillRect(hDC HDC, lprc *RECT, hbr HBRUSH) bool {
	ret, _, _ := procFillRect.Call(
		uintptr(hDC),
		uintptr(unsafe.Pointer(lprc)),
		uintptr(hbr))

	return ret != 0
}

func DrawText(hDC HDC, text string, uCount int, lpRect *RECT, uFormat uint) int {
	ret, _, _ := procDrawText.Call(
		uintptr(hDC),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))),
		uintptr(uCount),
		uintptr(unsafe.Pointer(lpRect)),
		uintptr(uFormat))

	return int(ret)
}

func AddClipboardFormatListener(hwnd HWND) bool {
	ret, _, _ := procAddClipboardFormatListener.Call(
		uintptr(hwnd))
	return ret != 0
}

func RemoveClipboardFormatListener(hwnd HWND) bool {
	ret, _, _ := procRemoveClipboardFormatListener.Call(
		uintptr(hwnd))
	return ret != 0
}

func OpenClipboard(hWndNewOwner HWND) bool {
	ret, _, _ := procOpenClipboard.Call(
		uintptr(hWndNewOwner))
	return ret != 0
}

func CloseClipboard() bool {
	ret, _, _ := procCloseClipboard.Call()
	return ret != 0
}

func EnumClipboardFormats(format uint) uint {
	ret, _, _ := procEnumClipboardFormats.Call(
		uintptr(format))
	return uint(ret)
}

func GetClipboardData(uFormat uint) HANDLE {
	ret, _, _ := procGetClipboardData.Call(
		uintptr(uFormat))
	return HANDLE(ret)
}

func SetClipboardData(uFormat uint, hMem HANDLE) HANDLE {
	ret, _, _ := procSetClipboardData.Call(
		uintptr(uFormat),
		uintptr(hMem))
	return HANDLE(ret)
}

func EmptyClipboard() bool {
	ret, _, _ := procEmptyClipboard.Call()
	return ret != 0
}

func GetClipboardFormatName(format uint) (string, bool) {
	cchMaxCount := 255
	buf := make([]uint16, cchMaxCount)
	ret, _, _ := procGetClipboardFormatName.Call(
		uintptr(format),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(cchMaxCount))

	if ret > 0 {
		return syscall.UTF16ToString(buf), true
	}

	return "Requested format does not exist or is predefined", false
}

func IsClipboardFormatAvailable(format uint) bool {
	ret, _, _ := procIsClipboardFormatAvailable.Call(uintptr(format))
	return ret != 0
}

func BeginPaint(hwnd HWND, paint *PAINTSTRUCT) HDC {
	ret, _, _ := procBeginPaint.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(paint)))
	return HDC(ret)
}

func EndPaint(hwnd HWND, paint *PAINTSTRUCT) {
	procEndPaint.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(paint)))
}

func GetKeyboardState(lpKeyState *[]byte) bool {
	ret, _, _ := procGetKeyboardState.Call(
		uintptr(unsafe.Pointer(&(*lpKeyState)[0])))
	return ret != 0
}

func MapVirtualKeyEx(uCode, uMapType uint, dwhkl HKL) uint {
	ret, _, _ := procMapVirtualKey.Call(
		uintptr(uCode),
		uintptr(uMapType),
		uintptr(dwhkl))
	return uint(ret)
}

func GetAsyncKeyState(vKey int) (keyState uint16) {
	ret, _, _ := procGetAsyncKeyState.Call(uintptr(vKey))
	return uint16(ret)
}

func GetKeyState(vKey int) (keyState uint16) {
	ret, _, _ := procGetKeyState.Call(uintptr(vKey))
	return uint16(ret)
}

func ToAscii(uVirtKey, uScanCode uint, lpKeyState *byte, lpChar *uint16, uFlags uint) int {
	ret, _, _ := procToAscii.Call(
		uintptr(uVirtKey),
		uintptr(uScanCode),
		uintptr(unsafe.Pointer(lpKeyState)),
		uintptr(unsafe.Pointer(lpChar)),
		uintptr(uFlags))
	return int(ret)
}

func SwapMouseButton(fSwap bool) bool {
	ret, _, _ := procSwapMouseButton.Call(
		uintptr(BoolToBOOL(fSwap)))
	return ret != 0
}

func GetCursorPos() (x, y int, ok bool) {
	pt := POINT{}
	ret, _, _ := procGetCursorPos.Call(uintptr(unsafe.Pointer(&pt)))
	return int(pt.X), int(pt.Y), ret != 0
}

func SetCursorPos(x, y int) bool {
	ret, _, _ := procSetCursorPos.Call(
		uintptr(x),
		uintptr(y),
	)
	return ret != 0
}

func SetCursor(cursor HCURSOR) HCURSOR {
	ret, _, _ := procSetCursor.Call(
		uintptr(cursor),
	)
	return HCURSOR(ret)
}

func CreateIcon(instance HINSTANCE, nWidth, nHeight int, cPlanes, cBitsPerPixel byte, ANDbits, XORbits *byte) HICON {
	ret, _, _ := procCreateIcon.Call(
		uintptr(instance),
		uintptr(nWidth),
		uintptr(nHeight),
		uintptr(cPlanes),
		uintptr(cBitsPerPixel),
		uintptr(unsafe.Pointer(ANDbits)),
		uintptr(unsafe.Pointer(XORbits)),
	)
	return HICON(ret)
}

func DestroyIcon(icon HICON) bool {
	ret, _, _ := procDestroyIcon.Call(
		uintptr(icon),
	)
	return ret != 0
}

func MonitorFromPoint(x, y int, dwFlags uint32) HMONITOR {
	ret, _, _ := procMonitorFromPoint.Call(
		uintptr(x),
		uintptr(y),
		uintptr(dwFlags),
	)
	return HMONITOR(ret)
}

func MonitorFromRect(rc *RECT, dwFlags uint32) HMONITOR {
	ret, _, _ := procMonitorFromRect.Call(
		uintptr(unsafe.Pointer(rc)),
		uintptr(dwFlags),
	)
	return HMONITOR(ret)
}

func MonitorFromWindow(hwnd HWND, dwFlags uint32) HMONITOR {
	ret, _, _ := procMonitorFromWindow.Call(
		uintptr(hwnd),
		uintptr(dwFlags),
	)
	return HMONITOR(ret)
}

func GetMonitorInfo(hMonitor HMONITOR, lmpi *MONITORINFO) bool {
	ret, _, _ := procGetMonitorInfo.Call(
		uintptr(hMonitor),
		uintptr(unsafe.Pointer(lmpi)),
	)
	return ret != 0
}

func EnumDisplayMonitors(hdc HDC, clip *RECT, fnEnum MONITORENUMPROC, dwData LPARAM) bool {
	ret, _, _ := procEnumDisplayMonitors.Call(
		uintptr(hdc),
		uintptr(unsafe.Pointer(clip)),
		syscall.NewCallback(func(Arg1 HMONITOR, Arg2 HDC, Arg3 uintptr, Arg4 LPARAM) LRESULT {
			if fnEnum(Arg1, Arg2, *(*RECT)(unsafe.Pointer(Arg3)), Arg4) {
				return 1
			} else {
				return 0
			}
		}),
		uintptr(dwData),
	)
	return ret != 0
}

func EnumDisplaySettingsEx(szDeviceName *uint16, iModeNum uint32, devMode *DEVMODE, dwFlags uint32) bool {
	ret, _, _ := procEnumDisplaySettingsEx.Call(
		uintptr(unsafe.Pointer(szDeviceName)),
		uintptr(iModeNum),
		uintptr(unsafe.Pointer(devMode)),
		uintptr(dwFlags),
	)
	return ret != 0
}

func ChangeDisplaySettingsEx(szDeviceName *uint16, devMode *DEVMODE, hwnd HWND, dwFlags uint32, lParam uintptr) int32 {
	ret, _, _ := procChangeDisplaySettingsEx.Call(
		uintptr(unsafe.Pointer(szDeviceName)),
		uintptr(unsafe.Pointer(devMode)),
		uintptr(hwnd),
		uintptr(dwFlags),
		lParam,
	)
	return int32(ret)
}

// Synthesizes keystrokes, mouse motions, and button clicks.
// see https://msdn.microsoft.com/en-us/library/windows/desktop/ms646310(v=vs.85).aspx
func SendInput(inputs []INPUT) (err error) {
	var validInputs []C.INPUT

	for _, oneInput := range inputs {
		input := C.INPUT{_type: C.DWORD(oneInput.Type)}

		switch oneInput.Type {
		case INPUT_MOUSE:
			(*MouseInput)(unsafe.Pointer(&input)).mi = oneInput.Mi
		case INPUT_KEYBOARD:
			(*KbdInput)(unsafe.Pointer(&input)).ki = oneInput.Ki
		case INPUT_HARDWARE:
			(*HardwareInput)(unsafe.Pointer(&input)).hi = oneInput.Hi
		default:
			err = errors.New("Unknown input type passed: " + fmt.Sprintf("%d", oneInput.Type))
			return
		}

		validInputs = append(validInputs, input)
	}

	_, _, err = procSendInput.Call(
		uintptr(len(validInputs)),
		uintptr(unsafe.Pointer(&validInputs[0])),
		uintptr(unsafe.Sizeof(C.INPUT{})),
	)
	if !IsErrSuccess(err) {
		return
	}
	err = nil
	return
}

// Simplifies SendInput for Keyboard related keys. Supports alphanumeric
func SendInputString(input string) (err error) {
	var inputs []INPUT
	b := make([]byte, 3)

	/*reg, err := regexp.Compile("^[a-zA-Z0-9]+")
	if err != nil {
		return
	}
	input = reg.ReplaceAllString(input, "")*/

	for _, rune := range input {

		utf8.EncodeRune(b, rune)
		vk := binary.LittleEndian.Uint16(b)
		fmt.Println(vk)

		input := INPUT{
			Type: INPUT_KEYBOARD,
			Ki: KEYBDINPUT{
				WVk:     vk,
				DwFlags: 0x0002 | 0x0008,
				Time:    200,
			},
		}
		inputs = append(inputs, input)
	}
	err = SendInput(inputs)
	return
}

func SetWindowsHookEx(idHook int, lpfn HOOKPROC, hMod HINSTANCE, dwThreadId DWORD) HHOOK {
	ret, _, _ := procSetWindowsHookEx.Call(
		uintptr(idHook),
		uintptr(syscall.NewCallback(lpfn)),
		uintptr(hMod),
		uintptr(dwThreadId),
	)
	return HHOOK(ret)
}

// Lifted from https://github.com/kbinani/win/blob/b749091b14a8a4867e0fc93567d5c6af7b360e8f/user32.go#L5215
func SetWinEventHook(eventMin DWORD, eventMax DWORD, hmodWinEventProc HMODULE, pfnWinEventProc HOOKPROC, idProcess DWORD, idThread DWORD, dwFlags DWORD) HHOOK {

	ret, _, _ := procSetWinEventHook.Call(
		uintptr(eventMin),
		uintptr(eventMax),
		uintptr(hmodWinEventProc),
		uintptr(syscall.NewCallback(pfnWinEventProc)),
		uintptr(idProcess),
		uintptr(idThread),
		uintptr(dwFlags),
	)
	return HHOOK(ret)
}

// Lifted from https://github.com/kbinani/win/blob/b749091b14a8a4867e0fc93567d5c6af7b360e8f/user32.go#L5548
func UnhookWinEvent(hWinEventHook HHOOK) bool {
	ret, _, _ := procUnhookWinEvent.Call(
		uintptr(hWinEventHook))
	return ret != 0
}

func UnhookWindowsHookEx(hhk HHOOK) bool {
	ret, _, _ := procUnhookWindowsHookEx.Call(
		uintptr(hhk),
	)
	return ret != 0
}

func CallNextHookEx(hhk HHOOK, nCode int, wParam WPARAM, lParam LPARAM) LRESULT {
	ret, _, _ := procCallNextHookEx.Call(
		uintptr(hhk),
		uintptr(nCode),
		uintptr(wParam),
		uintptr(lParam),
	)
	return LRESULT(ret)
}

// Defines a system-wide hotkey.
// See https://msdn.microsoft.com/en-us/library/windows/desktop/ms646309(v=vs.85).aspx
func RegisterHotKey(hwnd HWND, id int, fsModifiers uint, vkey uint) (err error) {
	_, _, err = procRegisterHotKey.Call(
		uintptr(hwnd),
		uintptr(id),
		uintptr(fsModifiers),
		uintptr(vkey),
	)
	if !IsErrSuccess(err) {
		return
	}
	err = nil
	return
}

// Defines a system-wide hotkey.
// See https://msdn.microsoft.com/en-us/library/windows/desktop/ms646309(v=vs.85).aspx
func UnregisterHotKey(hwnd HWND, id int) (err error) {
	_, _, err = procUnregisterHotKey.Call(
		uintptr(hwnd),
		uintptr(id),
	)
	if !IsErrSuccess(err) {
		return
	}
	err = nil
	return
}

// Translates a character to the corresponding virtual-key code and shift state for the current keyboard.
// See https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-vkkeyscanw
func VkKeyScanW(char uint16) int16 {
	ret, _, _ := procVkKeyScanW.Call(
		uintptr(char),
	)
	return int16(ret)
}

// Translates a character to the corresponding virtual-key code and shift state.
// See https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-vkkeyscanexw
func VkKeyScanExW(char uint16, hkl HKL) int16 {
	ret, _, _ := procVkKeyScanExW.Call(
		uintptr(char),
		uintptr(hkl),
	)
	return int16(ret)
}

// Sets the process-default DPI awareness to system-DPI awareness.
// See https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setprocessdpiaware
func SetProcessDPIAware() bool {
	ret, _, _ := setProcessDPIAware.Call()
	return ret != 0
}

func GetKeyNameText(scanCode uint32, extended, doNotCare bool) (string, bool) {
	buf := make([]uint16, 255)
	var lParam = uint64(scanCode)
	lParam <<= 16
	if extended {
		lParam |= 1 << 24
	}

	if doNotCare {
		lParam |= 1 << 25
	}

	ret, _, _ := procGetKeyNameText.Call(
		uintptr(lParam),
		uintptr(unsafe.Pointer(&buf[0])),
		255,
	)

	return syscall.UTF16ToString(buf), ret != 0
}
