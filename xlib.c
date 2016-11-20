
//          Copyright 2016 Vitali Baumtrok
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


#include <X11/Xlib.h>

void xlib_xevent_type ( const XEvent *const xevent, int *const xevent_type_return ) {
	xevent_type_return[0] = xevent[0].type;
}

void xlib_xkeyevent_values ( const XEvent *const xevent,
                              unsigned long *const serial_return,
                              Bool *const send_event_return,
                              Display **const display_return,
                              Window *const window_return,
                              Window *const root_return,
                              Window *const subwindow_return,
                              Time *const time_return,
                              int *const x_return,
                              int *const y_return,
                              int *const x_root_return,
                              int *const y_root_return,
                              unsigned int *const state_return,
                              unsigned int *const keycode_return,
                              Bool *const same_screen_return ) {
	serial_return[0] = xevent[0].xkey.serial;
	send_event_return[0] = xevent[0].xkey.send_event;
	display_return[0] = xevent[0].xkey.display;
	window_return[0] = xevent[0].xkey.window;
	subwindow_return[0] = xevent[0].xkey.subwindow;
	time_return[0] = xevent[0].xkey.time;
	x_return[0] = xevent[0].xkey.x;
	y_return[0] = xevent[0].xkey.y;
	x_root_return[0] = xevent[0].xkey.x_root;
	y_root_return[0] = xevent[0].xkey.y_root;
	state_return[0] = xevent[0].xkey.state;
	keycode_return[0] = xevent[0].xkey.keycode;
	same_screen_return[0] = xevent[0].xkey.same_screen;
}
