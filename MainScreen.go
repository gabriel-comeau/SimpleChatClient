package main

import (
	"github.com/gabriel-comeau/tbuikit"
	"github.com/nsf/termbox-go"
)

// Holds the event handler callbacks and rectangle calculation callbacks for
// the Connect Screen

//
// Event Handler Callbacks
//

// Handle the escape button pressed on the main screen,
// causing the server to be disconnected from and the
// client to returrn to the connection screen.
func doDisconnect(uiElement, event interface{}) {
	disconnect()
}

// Handles enter key presses for chat input widget
func chatEnterHandler(uiElement, event interface{}) {
	widget, ok := uiElement.(*tbuikit.TextInputWidget)
	if ok {
		netChatChan <- widget.GetBuffer().ReturnAndClear() + "\n"
	}
}

//
// Rectangle Calculation Callbacks
//

// Top title bar
func calculateTopTitleBar() (x1, x2, y1, y2 int) {
	w := tbuikit.GetTermboxWidth()
	x1 = 1
	x2 = w - 1
	y1 = 1
	y2 = TITLE_BAR_HEIGHT
	return
}

// Chat input widget
func calculateChatBufferRect() (x1, x2, y1, y2 int) {
	w, h := termbox.Size()
	x1 = 1
	x2 = w - 1
	y1 = (h / 2) + (h / 4)
	y2 = h - 2
	return
}

// Message display widget
func calculateMessageBufferRect() (x1, x2, y1, y2 int) {
	w, h := termbox.Size()
	x1 = 1
	x2 = w - (w / 4)
	y1 = TITLE_BAR_HEIGHT + 1
	y2 = (h / 2) + (h / 4) - 1
	return
}

// Who's online display widget
func calculateWhoRect() (x1, x2, y1, y2 int) {
	w, h := termbox.Size()
	x1 = w - (w / 4) + 1
	x2 = w - 1
	y1 = TITLE_BAR_HEIGHT + 6
	y2 = (h / 2) + (h / 4) - 1
	return
}

// Label for Who's online display widget
func calculateWhoLabelRect() (x1, x2, y1, y2 int) {
	w := tbuikit.GetTermboxWidth()
	x1 = w - (w / 4) + 1
	x2 = w - 1
	y1 = TITLE_BAR_HEIGHT + 1
	y2 = TITLE_BAR_HEIGHT + 5
	return
}
