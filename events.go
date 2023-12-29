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

const (
	Expose     = int(C.Expose)
	KeyPress   = int(C.KeyPress)
	KeyRelease = int(C.KeyRelease)
)

const (
	// All Event masks.
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
	GrabModeSync             = int(C.GrabModeSync)
	GrabModeAsync            = int(C.GrabModeAsync)
	Mod1Mask                 = int(C.Mod1Mask)
	Mod2Mask                 = int(C.Mod2Mask)
	Mod3Mask                 = int(C.Mod3Mask)
	Mod4Mask                 = int(C.Mod4Mask)
	Mod5Mask                 = int(C.Mod5Mask)
	ShiftMask                = int(C.ShiftMask)
	LockMask                 = int(C.LockMask)
	ControlMask              = int(C.ControlMask)
	AnyModifier              = int(C.AnyModifier)
	Button1Mask              = int(C.Button1Mask)
	Button2Mask              = int(C.Button2Mask)
	Button3Mask              = int(C.Button3Mask)
	Button4Mask              = int(C.Button4Mask)
	Button5Mask              = int(C.Button5Mask)
	ShiftMapIndex            = int(C.ShiftMapIndex)
	LockMapIndex             = int(C.LockMapIndex)
	ControlMapIndex          = int(C.ControlMapIndex)
	Mod1MapIndex             = int(C.Mod1MapIndex)
	Mod2MapIndex             = int(C.Mod2MapIndex)
	Mod3MapIndex             = int(C.Mod3MapIndex)
	Mod4MapIndex             = int(C.Mod4MapIndex)
	Mod5MapIndex             = int(C.Mod5MapIndex)
	PointerWindow            = int(C.PointerWindow)
	InputFocus               = int(C.InputFocus)
	PointerRoot              = int(C.PointerRoot)
	AnyPropertyType          = int(C.AnyPropertyType)
	AnyKey                   = int(C.AnyKey)
	AnyButton                = int(C.AnyButton)
	AllTemporary             = int(C.AllTemporary)
	CurrentTime              = int(C.CurrentTime)
	NoSymbol                 = int(C.NoSymbol)
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

func XNextRequest(display *Display) uint64 {
	displayC := (*C.Display)(display)
	return uint64(C.XNextRequest(displayC))
}

func (ev *tEventType) Type() int {
	return ev.typeCode
}
