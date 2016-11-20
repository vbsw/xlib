
//          Copyright 2016 Vitali Baumtrok
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)

#ifndef GOXLIB_H
#define GOXLIB_H

extern void xlib_xevent_type ( const XEvent *const xevent, int *const xevent_type_return );
extern void xlib_xkeyevent_values ( const XEvent *const xevent,
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
                                     Bool *const same_screen_return);

#endif /* GOXLIB_H */
