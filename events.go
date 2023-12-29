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

//"unsafe"

const (
	Expose     = int(C.Expose)
	KeyPress   = int(C.KeyPress)
	KeyRelease = int(C.KeyRelease)
)

const (
	NoEventMask              = int64(C.NoEventMask)
	KeyPressMask             = int64(C.KeyPressMask)
	KeyReleaseMask           = int64(C.KeyReleaseMask)
	ButtonPressMask          = int64(C.ButtonPressMask)
	ButtonReleaseMask        = int64(C.ButtonReleaseMask)
	EnterWindowMask          = int64(C.EnterWindowMask)
	LeaveWindowMask          = int64(C.LeaveWindowMask)
	PointerMotionMask        = int64(C.PointerMotionMask)
	PointerMotionHintMask    = int64(C.PointerMotionHintMask)
	Button1MotionMask        = int64(C.Button1MotionMask)
	Button2MotionMask        = int64(C.Button2MotionMask)
	Button3MotionMask        = int64(C.Button3MotionMask)
	Button4MotionMask        = int64(C.Button4MotionMask)
	Button5MotionMask        = int64(C.Button5MotionMask)
	ButtonMotionMask         = int64(C.ButtonMotionMask)
	KeymapStateMask          = int64(C.KeymapStateMask)
	ExposureMask             = int64(C.ExposureMask)
	VisibilityChangeMask     = int64(C.VisibilityChangeMask)
	StructureNotifyMask      = int64(C.StructureNotifyMask)
	ResizeRedirectMask       = int64(C.ResizeRedirectMask)
	SubstructureNotifyMask   = int64(C.SubstructureNotifyMask)
	SubstructureRedirectMask = int64(C.SubstructureRedirectMask)
	FocusChangeMask          = int64(C.FocusChangeMask)
	PropertyChangeMask       = int64(C.PropertyChangeMask)
	ColormapChangeMask       = int64(C.ColormapChangeMask)
	OwnerGrabButtonMask      = int64(C.OwnerGrabButtonMask)
	GrabModeSync             = int64(C.GrabModeSync)
	GrabModeAsync            = int64(C.GrabModeAsync)
)

type XEvent interface {
	Type() int
}

type tEventType struct {
	typeCode int
}

type XKeyEvent struct {
	tEventType
	Serial       uint64
	SendEvent    bool
	Display      *Display
	Window       Window
	Root         Window
	Subwindow    Window
	Time         uint64
	X, Y         int
	XRoot, YRoot int
	State        uint
	KeyCode      uint
	SameScreen   bool
}

func XNextEvent(display *Display) XEvent {
	displayC := (*C.Display)(display)
	var xeventC C.XEvent
	var xeventTypeC C.int
	C.XNextEvent(displayC, &xeventC)
	C.xlib_xevent_type(&xeventC, &xeventTypeC)

	switch xeventTypeC {
	case C.KeyPress:
		return newXKeyEvent(&xeventC, xeventTypeC)
	case C.KeyRelease:
		return newXKeyEvent(&xeventC, xeventTypeC)
	}
	return nil
}

func newXKeyEvent(xeventC *C.XEvent, xeventTypeC C.int) *XKeyEvent {
	xKeyEvent := new(XKeyEvent)
	var serialC C.ulong
	var sendEventC C.Bool
	var displayC *C.Display
	var windowC C.Window
	var rootC C.Window
	var subwindowC C.Window
	var timeC C.Time
	var xC, yC C.int
	var xRootC, yRootC C.int
	var stateC C.uint
	var keyCodeC C.uint
	var sameScreenC C.Bool
	C.xlib_xkeyevent_values(xeventC, &serialC, &sendEventC, &displayC, &windowC, &rootC, &subwindowC,
		&timeC, &xC, &yC, &xRootC, &yRootC, &stateC, &keyCodeC, &sameScreenC)
	xKeyEvent.typeCode = int(xeventTypeC)
	xKeyEvent.Serial = uint64(serialC)
	xKeyEvent.SendEvent = (sendEventC != 0)
	xKeyEvent.Display = (*Display)(displayC)
	xKeyEvent.Window = Window(windowC)
	xKeyEvent.Root = Window(rootC)
	xKeyEvent.Subwindow = Window(subwindowC)
	xKeyEvent.Time = uint64(timeC)
	xKeyEvent.X = int(xC)
	xKeyEvent.Y = int(yC)
	xKeyEvent.XRoot = int(xRootC)
	xKeyEvent.YRoot = int(yRootC)
	xKeyEvent.State = uint(stateC)
	xKeyEvent.KeyCode = uint(keyCodeC)
	xKeyEvent.SameScreen = (sameScreenC != 0)
	return xKeyEvent
}

func (this *tEventType) Type() int {
	return this.typeCode
}
