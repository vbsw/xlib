
//          Copyright 2016 Vitali Baumtrok
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


// Binding of Xlib (version 11, release 6.7).
package xlib


// #cgo LDFLAGS: -lX11
// #include <stdlib.h>
// #include <X11/Xlib.h>
import "C"
import (
	"unsafe"
)

type Display C.Display
type Screen C.Screen

func strConcat ( a []interface{} ) string {
	str := ""
	for _, strPart := range a {
		switch s := strPart.(type) {
			case string:
				str += s
		}
	}
	return str
}

func XOpenDisplay ( displayNameParts ...interface{} ) *Display {
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

func XCloseDisplay ( display *Display ) {
	displayC := (*C.Display)(display)
	C.XCloseDisplay(displayC)
}

func XDisplayString ( display *Display ) string {
	displayC := (*C.Display)(display)
	displayNameC := C.XDisplayString(displayC)
	displayName := C.GoString(displayNameC)
	C.free(unsafe.Pointer(displayNameC))
	return displayName
}

func XScreenCount ( display *Display ) int {
	displayC := (*C.Display)(display)
	screenCount := C.XScreenCount(displayC)
	return int(screenCount)
}

func XScreenOfDisplay ( display *Display, screenNumber int ) *Screen {
	displayC := (*C.Display)(display)
	screen := C.XScreenOfDisplay(displayC, C.int(screenNumber))
	return (*Screen)(screen)
}

func XWidthOfScreen ( screen *Screen ) int {
	screenC := (*C.Screen)(screen)
	width := C.XWidthOfScreen(screenC)
	return int(width)
}

func XHeightOfScreen ( screen *Screen ) int {
	screenC := (*C.Screen)(screen)
	height := C.XHeightOfScreen(screenC)
	return int(height)
}

