//          Copyright 2016 Vitali Baumtrok
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)

// Binding of Xlib (version 11, release 6.7).
package xlib

// #cgo LDFLAGS: -lX11
// #include <stdlib.h>
// #include <X11/Xlib.h>
// #include "xlib.h"
import "C"
import (
	"unsafe"
)

type Display C.Display
type Screen C.Screen
type Window C.Window
type Bool C.Bool

func strConcat(a []interface{}) string {
	str := ""
	for _, strPart := range a {
		switch s := strPart.(type) {
		case string:
			str += s
		}
	}
	return str
}

func XOpenDisplay(displayNameParts ...interface{}) *Display {
	if len(displayNameParts) == 0 {
		display := C.XOpenDisplay(nil)
		return (*Display)(display)

	} else {
		displayNameComplete := strConcat(displayNameParts)
		if len(displayNameComplete) > 0 {
			displayNameCompleteC := C.CString(displayNameComplete)
			display := C.XOpenDisplay(displayNameCompleteC)
			C.free(unsafe.Pointer(displayNameCompleteC))
			return (*Display)(display)

		} else {
			display := C.XOpenDisplay(nil)
			return (*Display)(display)
		}
	}
}

func XCloseDisplay(display *Display) {
	displayC := (*C.Display)(display)
	C.XCloseDisplay(displayC)
}

func XDisplayString(display *Display) string {
	displayC := (*C.Display)(display)
	displayNameC := C.XDisplayString(displayC)
	displayName := C.GoString(displayNameC)
	C.free(unsafe.Pointer(displayNameC))
	return displayName
}

func XScreenCount(display *Display) int {
	displayC := (*C.Display)(display)
	screenCount := C.XScreenCount(displayC)
	return int(screenCount)
}

func XScreenOfDisplay(display *Display, screenNumber int) *Screen {
	displayC := (*C.Display)(display)
	screen := C.XScreenOfDisplay(displayC, C.int(screenNumber))
	return (*Screen)(screen)
}

func XWidthOfScreen(screen *Screen) int {
	screenC := (*C.Screen)(screen)
	width := C.XWidthOfScreen(screenC)
	return int(width)
}

func XHeightOfScreen(screen *Screen) int {
	screenC := (*C.Screen)(screen)
	height := C.XHeightOfScreen(screenC)
	return int(height)
}

func XDefaultScreenOfDisplay(display *Display) *Screen {
	displayC := (*C.Display)(display)
	defaultScreen := C.XDefaultScreenOfDisplay(displayC)
	return (*Screen)(defaultScreen)
}

func XRootWindowOfScreen(screen *Screen) Window {
	screenC := (*C.Screen)(screen)
	rootWindow := C.XRootWindowOfScreen(screenC)
	return Window(rootWindow)
}

func XCreateSimpleWindow(display *Display, parent Window, x, y int, width, height, borderWidth uint, border, background uint64) Window {
	displayC := (*C.Display)(display)
	windowC := (C.Window)(parent)
	xC := C.int(x)
	yC := C.int(y)
	widthC := C.uint(width)
	heightC := C.uint(height)
	borderWidthC := C.uint(borderWidth)
	borderC := C.ulong(border)
	backgroundC := C.ulong(background)
	window := C.XCreateSimpleWindow(displayC, windowC, xC, yC, widthC, heightC, borderWidthC, borderC, backgroundC)
	return Window(window)
}

func XMapWindow(display *Display, window Window) {
	displayC := (*C.Display)(display)
	windowC := (C.Window)(window)
	C.XMapWindow(displayC, windowC)
}

func XSelectInput(display *Display, window Window, eventMask int64) {
	displayC := (*C.Display)(display)
	windowC := (C.Window)(window)
	eventMaskC := (C.long)(eventMask)
	C.XSelectInput(displayC, windowC, eventMaskC)
}

func XGrabKey(display *Display, keycode int, modifiers int, grab_window Window, owner_events Bool, pointer_mode int, keyboard_mode int) int {
	displayC := (*C.Display)(display)
	keycodeC := (C.int)(keycode)
	modifiersC := (C.uint)(modifiers)
	grab_windowC := (C.Window)(grab_window)
	owner_eventsC := (C.Bool)(owner_events)
	pointer_modeC := (C.int)(pointer_mode)
	keyboard_modeC := (C.int)(keyboard_mode)
	status := C.XGrabKey(displayC, keycodeC, modifiersC, grab_windowC, owner_eventsC, pointer_modeC, keyboard_modeC)
	return int(status)

}

func XRaiseWindow(display *Display, window Window) {
	displayC := (*C.Display)(display)
	windowC := (C.Window)(window)
	C.XRaiseWindow(displayC, windowC)
}

func XGetWindowAttributes(display *Display, window Window) *C.XWindowAttributes {
	displayC := (*C.Display)(display)
	windowC := (C.Window)(window)
	windowAttributes := new(C.XWindowAttributes)
	C.XGetWindowAttributes(displayC, windowC, windowAttributes)
	return windowAttributes
}

func XMoveResizeWindow(display *Display, window Window, x, y int, width, height uint) {
	displayC := (*C.Display)(display)
	windowC := (C.Window)(window)
	xC := (C.int)(x)
	yC := (C.int)(y)
	widthC := (C.uint)(width)
	heightC := (C.uint)(height)
	C.XMoveResizeWindow(displayC, windowC, xC, yC, widthC, heightC)
}

func XKeysymToKeycode(display *Display, keysym uint64) int {
	displayC := (*C.Display)(display)
	keysymC := (C.KeySym)(keysym)
	keycode := C.XKeysymToKeycode(displayC, keysymC)
	return int(keycode)
}

func XStringToKeysym(string string) uint64 {
	stringC := C.CString(string)
	keysym := C.XStringToKeysym(stringC)
	C.free(unsafe.Pointer(stringC))
	return uint64(keysym)
}
