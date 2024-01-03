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
type WindowAttributes C.XWindowAttributes
type Cursor C.Cursor

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

func XGrabButton(display *Display, button int, modifiers int, grab_window Window, owner_events Bool, event_mask int64, pointer_mode int, keyboard_mode int, confine_to Window, cursor Cursor) int {
	displayC := (*C.Display)(display)
	buttonC := (C.uint)(button)
	modifiersC := (C.uint)(modifiers)
	grab_windowC := (C.Window)(grab_window)
	owner_eventsC := (C.Bool)(owner_events)
	event_maskC := (C.uint)(event_mask)
	pointer_modeC := (C.int)(pointer_mode)
	keyboard_modeC := (C.int)(keyboard_mode)
	confine_toC := (C.Window)(confine_to)
	cursorC := (C.Cursor)(cursor)
	status := C.XGrabButton(displayC, buttonC, modifiersC, grab_windowC, owner_eventsC, event_maskC, pointer_modeC, keyboard_modeC, confine_toC, cursorC)
	return int(status)
}

func XDefaultRootWindow(display *Display) Window {
	displayC := (*C.Display)(display)
	rootWindow := C.XDefaultRootWindow(displayC)
	return Window(rootWindow)
}

func XUngrabKey(display *Display, keycode int, modifiers int, grab_window Window) {
	displayC := (*C.Display)(display)
	keycodeC := (C.int)(keycode)
	modifiersC := (C.uint)(modifiers)
	grab_windowC := (C.Window)(grab_window)
	C.XUngrabKey(displayC, keycodeC, modifiersC, grab_windowC)
}

func XUngrabButton(display *Display, button int, modifiers int, grab_window Window) {
	displayC := (*C.Display)(display)
	buttonC := (C.uint)(button)
	modifiersC := (C.uint)(modifiers)
	grab_windowC := (C.Window)(grab_window)
	C.XUngrabButton(displayC, buttonC, modifiersC, grab_windowC)
}

func XGrabPointer(display *Display, grab_window Window, owner_events Bool, event_mask int64, pointer_mode int, keyboard_mode int, confine_to Window, cursor Cursor, time uint64) int {
	displayC := (*C.Display)(display)
	grab_windowC := (C.Window)(grab_window)
	owner_eventsC := (C.Bool)(owner_events)
	event_maskC := (C.uint)(event_mask)
	pointer_modeC := (C.int)(pointer_mode)
	keyboard_modeC := (C.int)(keyboard_mode)
	confine_toC := (C.Window)(confine_to)
	cursorC := (C.Cursor)(cursor)
	timeC := (C.Time)(time)
	status := C.XGrabPointer(displayC, grab_windowC, owner_eventsC, event_maskC, pointer_modeC, keyboard_modeC, confine_toC, cursorC, timeC)
	return int(status)
}

func XUngrabPointer(display *Display, time uint64) {
	displayC := (*C.Display)(display)
	timeC := (C.Time)(time)
	C.XUngrabPointer(displayC, timeC)
}

func XWarpPointer(display *Display, src_window Window, dest_window Window, src_x, src_y int, src_width, src_height uint, dest_x, dest_y int) {
	displayC := (*C.Display)(display)
	src_windowC := (C.Window)(src_window)
	dest_windowC := (C.Window)(dest_window)
	src_xC := (C.int)(src_x)
	src_yC := (C.int)(src_y)
	src_widthC := (C.uint)(src_width)
	src_heightC := (C.uint)(src_height)
	dest_xC := (C.int)(dest_x)
	dest_yC := (C.int)(dest_y)
	C.XWarpPointer(displayC, src_windowC, dest_windowC, src_xC, src_yC, src_widthC, src_heightC, dest_xC, dest_yC)
}

func XQueryPointer(display *Display, window Window) (root_return Window, child_return Window, root_x_return, root_y_return, win_x_return, win_y_return int, mask_return uint) {
	displayC := (*C.Display)(display)
	windowC := (C.Window)(window)
	var root_returnC C.Window
	var child_returnC C.Window
	var root_x_returnC C.int
	var root_y_returnC C.int
	var win_x_returnC C.int
	var win_y_returnC C.int
	var mask_returnC C.uint
	C.XQueryPointer(displayC, windowC, &root_returnC, &child_returnC, &root_x_returnC, &root_y_returnC, &win_x_returnC, &win_y_returnC, &mask_returnC)
	root_return = Window(root_returnC)
	child_return = Window(child_returnC)
	root_x_return = int(root_x_returnC)
	root_y_return = int(root_y_returnC)
	win_x_return = int(win_x_returnC)
	win_y_return = int(win_y_returnC)
	mask_return = uint(mask_returnC)
	return
}

func XUnmapWindow(display *Display, window Window) {
	displayC := (*C.Display)(display)
	windowC := (C.Window)(window)
	C.XUnmapWindow(displayC, windowC)
}

func XDestroyWindow(display *Display, window Window) {
	displayC := (*C.Display)(display)
	windowC := (C.Window)(window)
	C.XDestroyWindow(displayC, windowC)
}

func XFree(data unsafe.Pointer) {
	C.XFree(data)
}

func XFlush(display *Display) {
	displayC := (*C.Display)(display)
	C.XFlush(displayC)
}

func XStoreName(display *Display, window Window, name string) {
	displayC := (*C.Display)(display)
	windowC := (C.Window)(window)
	nameC := C.CString(name)
	C.XStoreName(displayC, windowC, nameC)
	C.free(unsafe.Pointer(nameC))
}

func XCreateFontCursor(display *Display, shape uint) Cursor {
	displayC := (*C.Display)(display)
	shapeC := (C.uint)(shape)
	cursor := C.XCreateFontCursor(displayC, shapeC)
	return Cursor(cursor)
}

func XDefineCursor(display *Display, window Window, cursor Cursor) {
	displayC := (*C.Display)(display)
	windowC := (C.Window)(window)
	cursorC := (C.Cursor)(cursor)
	C.XDefineCursor(displayC, windowC, cursorC)
}

func XUndefineCursor(display *Display, window Window) {
	displayC := (*C.Display)(display)
	windowC := (C.Window)(window)
	C.XUndefineCursor(displayC, windowC)
}

func XCreateGC(display *Display, drawable Window, mask uint64, values *C.XGCValues) C.GC {
	displayC := (*C.Display)(display)
	drawableC := (C.Drawable)(drawable)
	gc := C.XCreateGC(displayC, drawableC, C.ulong(mask), values)
	return gc
}

func XFreeGC(display *Display, gc C.GC) {
	displayC := (*C.Display)(display)
	C.XFreeGC(displayC, gc)
}

func XSetForeground(display *Display, gc C.GC, foreground uint64) {
	displayC := (*C.Display)(display)
	foregroundC := (C.ulong)(foreground)
	C.XSetForeground(displayC, gc, foregroundC)
}

func XSetBackground(display *Display, gc C.GC, background uint64) {
	displayC := (*C.Display)(display)
	backgroundC := (C.ulong)(background)
	C.XSetBackground(displayC, gc, backgroundC)
}

func XSetLineAttributes(display *Display, gc C.GC, line_width uint, line_style int, cap_style int, join_style int) {
	displayC := (*C.Display)(display)
	line_widthC := (C.uint)(line_width)
	line_styleC := (C.int)(line_style)
	cap_styleC := (C.int)(cap_style)
	join_styleC := (C.int)(join_style)
	C.XSetLineAttributes(displayC, gc, line_widthC, line_styleC, cap_styleC, join_styleC)
}

func XDrawLine(display *Display, drawable Window, gc C.GC, x1, y1, x2, y2 int) {
	displayC := (*C.Display)(display)
	drawableC := (C.Drawable)(drawable)
	x1C := (C.int)(x1)
	y1C := (C.int)(y1)
	x2C := (C.int)(x2)
	y2C := (C.int)(y2)
	C.XDrawLine(displayC, drawableC, gc, x1C, y1C, x2C, y2C)
}

func XDrawRectangle(display *Display, drawable Window, gc C.GC, x, y int, width, height uint) {
	displayC := (*C.Display)(display)
	drawableC := (C.Drawable)(drawable)
	xC := (C.int)(x)
	yC := (C.int)(y)
	widthC := (C.uint)(width)
	heightC := (C.uint)(height)
	C.XDrawRectangle(displayC, drawableC, gc, xC, yC, widthC, heightC)
}

func XFillRectangle(display *Display, drawable Window, gc C.GC, x, y int, width, height uint) {
	displayC := (*C.Display)(display)
	drawableC := (C.Drawable)(drawable)
	xC := (C.int)(x)
	yC := (C.int)(y)
	widthC := (C.uint)(width)
	heightC := (C.uint)(height)
	C.XFillRectangle(displayC, drawableC, gc, xC, yC, widthC, heightC)
}

func XQlength(display *Display) int {
	displayC := (*C.Display)(display)
	length := C.XQLength(displayC)
	return int(length)
}

func XRootWindow(display *Display, screenNumber int) Window {
	displayC := (*C.Display)(display)
	screenNumberC := (C.int)(screenNumber)
	rootWindow := C.XRootWindow(displayC, screenNumberC)
	return Window(rootWindow)
}

func XServerVendor(display *Display) string {
	displayC := (*C.Display)(display)
	serverVendorC := C.XServerVendor(displayC)
	serverVendor := C.GoString(serverVendorC)
	C.free(unsafe.Pointer(serverVendorC))
	return serverVendor
}

func XVendorRelease(display *Display) int {
	displayC := (*C.Display)(display)
	vendorRelease := C.XVendorRelease(displayC)
	return int(vendorRelease)
}

func XDefaultScreen(display *Display) int {
	displayC := (*C.Display)(display)
	screenNumber := C.XDefaultScreen(displayC)
	return int(screenNumber)
}

func XDefaultDepth(display *Display, screenNumber int) int {
	displayC := (*C.Display)(display)
	screenNumberC := (C.int)(screenNumber)
	depth := C.XDefaultDepth(displayC, screenNumberC)
	return int(depth)
}
