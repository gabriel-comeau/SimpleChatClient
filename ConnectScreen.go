package main

import (
	"github.com/nsf/termbox-go"
	"github.com/gabriel-comeau/termbox-uikit"
)

// Holds the event handler callbacks and rectangle calculation callbacks for
// the Connect Screen

//
// Event Handler Callbacks
//

// Switches between selectable widgets on the screen when fired.
func switchConnectScreenActiveWidget(uiElement interface{}, event interface{}) {
	scr, ok := uiElement.(*termbox-uikit.Screen)
	if ok {
		scr.SelectNextWidget()
	}
}

func doConnectButtonPress(uiElement interface{}, event interface{}) {
	connect()
}

// When the quit button is pressed, cause the UI to shutdown.
func doQuitButtonPress(uiElement interface{}, event interface{}) {
	chatUi.Shutdown()
}

//
// Rectangle Calculation Callbacks
//

// Message Box
func calcMsgWidgetRect() (x1, x2, y1, y2 int) {
	w, h := termbox.Size()
	x1 = w / 4
	x2 = x1 + 60
	y2 = h / 4
	y1 = y2 - 6
	return
}

// Hostname Input Field
func calcHostFieldRect() (x1, x2, y1, y2 int) {
	w, h := termbox.Size()
	x1 = w/2 - 10
	x2 = w/2 + 10
	y2 = h - h/3 //30
	y1 = y2 - 2  // 28
	return
}

// Port Input Field
func calcPortFieldRect() (x1, x2, y1, y2 int) {
	w, h := termbox.Size()
	x1 = w/2 - 10
	x2 = w/2 + 10
	y2 = h - h/3 - 4 // 24
	y1 = y2 - 2      // 22
	return
}

// Nickname Input Field
func calcNickFieldRect() (x1, x2, y1, y2 int) {
	w, h := termbox.Size()
	x1 = w/2 - 10
	x2 = w/2 + 10
	y2 = h - h/3 - 8 // 22
	y1 = y2 - 2
	return
}

// Hostname Label
func calcHostLabelRect() (x1, x2, y1, y2 int) {
	w, h := termbox.Size()
	x1 = w/2 - 20
	x2 = x1 + 10
	y2 = h - h/3 //30
	y1 = y2 - 2  // 28
	return
}

// Port Input Field
func calcPortLabelRect() (x1, x2, y1, y2 int) {
	w, h := termbox.Size()
	x1 = w/2 - 20
	x2 = x1 + 10
	y2 = h - h/3 - 4 // 24
	y1 = y2 - 2      // 22
	return
}

// Nickname Input Field
func calcNickLabelRect() (x1, x2, y1, y2 int) {
	w, h := termbox.Size()
	x1 = w/2 - 20
	x2 = x1 + 10
	y2 = h - h/3 - 8 // 22
	y1 = y2 - 2
	return
}

// Connect Button
func calConnectButtonRect() (x1, x2, y1, y2 int) {
	w, h := termbox.Size()
	x1 = w/2 - 40
	x2 = w/2 - 20
	y2 = h - h/3 + 6 //34
	y1 = y2 - 4
	return
}

// Quit Button
func calQuitButtonRect() (x1, x2, y1, y2 int) {
	w, h := termbox.Size()
	x1 = w/2 + 20
	x2 = w/2 + 40
	y2 = h - h/3 + 6 //34
	y1 = y2 - 4
	return
}
