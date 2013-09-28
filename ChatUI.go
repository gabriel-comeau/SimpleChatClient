package main

import (
	"github.com/nsf/termbox-go"
	"github.com/gabriel-comeau/tbuikit"
)

// How far down the top bar with the title goes
const TITLE_BAR_HEIGHT = 5

// Instantiate all of the UI objects here and assign them to their callbacks
func initUi() {
	chatUi = new(tbuikit.UI)
	mainScreen = new(tbuikit.Screen)
	connectScreen = new(tbuikit.Screen)

	initMainScreenWidgets()
	initConnectScreenWidgets()

	connectScreen.Activate()

	chatUi.AddScreen(mainScreen)
	chatUi.AddScreen(connectScreen)
}

// Builds up the widgets for the "main" screen of the application
func initMainScreenWidgets() {
	topTitleBar := tbuikit.CreateLabelWidget("TermboxUI Chat - ESC to disconnect - /help for commands", true, tbuikit.CENTER,
		termbox.ColorYellow, termbox.ColorDefault, termbox.ColorWhite, calculateTopTitleBar)

	chatInputWidget := tbuikit.CreateTextInputWidget(true, termbox.ColorWhite,
		termbox.ColorDefault, termbox.ColorDefault, calculateChatBufferRect, chatBuffer, true, true)

	messageWidget := tbuikit.CreateColorizedTextWidget(termbox.ColorWhite, termbox.ColorWhite,
		termbox.ColorDefault, calculateMessageBufferRect, messageBuffer)

	chatInputWidget.UseDefaultKeys(true)
	chatInputWidget.AddSpecialKeyCallback(termbox.KeyEnter, chatEnterHandler)

	mainScreen.AddWidget(topTitleBar)
	mainScreen.AddWidget(messageWidget)
	mainScreen.AddWidget(chatInputWidget)

	mainScreen.AddSpecialKeyCallback(termbox.KeyEsc, doDisconnect)
}

// Builds up the widgets for the connection screen
func initConnectScreenWidgets() {

	connectTitleBar := tbuikit.CreateLabelWidget("TermboxUI Chat - tab to switch fields", true, tbuikit.CENTER,
		termbox.ColorYellow, termbox.ColorDefault, termbox.ColorWhite, calculateTopTitleBar)

	hostFieldWidget := tbuikit.CreateTextInputWidget(true, termbox.ColorWhite, termbox.ColorWhite, termbox.ColorGreen, calcHostFieldRect, hostFieldBuffer, true, false)
	portFieldWidget := tbuikit.CreateTextInputWidget(true, termbox.ColorWhite, termbox.ColorWhite, termbox.ColorGreen, calcPortFieldRect, portFieldBuffer, true, false)
	nickFieldWidget := tbuikit.CreateTextInputWidget(true, termbox.ColorWhite, termbox.ColorWhite, termbox.ColorGreen, calcNickFieldRect, nickFieldBuffer, true, true)

	hostFieldLabel := tbuikit.CreateLabelWidget("Host:", false, tbuikit.BOTTOM_RIGHT, termbox.ColorWhite, termbox.ColorDefault, termbox.ColorDefault, calcHostLabelRect)
	portFieldLabel := tbuikit.CreateLabelWidget("Port:", false, tbuikit.BOTTOM_RIGHT, termbox.ColorWhite, termbox.ColorDefault, termbox.ColorDefault, calcPortLabelRect)
	nickFieldLabel := tbuikit.CreateLabelWidget("Nickname:", false, tbuikit.BOTTOM_RIGHT, termbox.ColorWhite, termbox.ColorDefault, termbox.ColorDefault, calcNickLabelRect)

	connectMsgWidget := tbuikit.CreateColorizedTextWidget(termbox.ColorBlue, termbox.ColorWhite, termbox.ColorDefault, calcMsgWidgetRect, connectMsgBuffer)

	connectButton := tbuikit.CreateButtonWidget("Connect", tbuikit.CENTER, termbox.ColorWhite,
		termbox.ColorBlue, termbox.ColorDefault, termbox.ColorDefault, termbox.ColorWhite,
		termbox.ColorGreen, calConnectButtonRect, true, false)

	quitButton := tbuikit.CreateButtonWidget("Quit", tbuikit.CENTER, termbox.ColorWhite,
		termbox.ColorRed, termbox.ColorDefault, termbox.ColorDefault, termbox.ColorWhite,
		termbox.ColorGreen, calQuitButtonRect, true, false)

	quitButton.AddSpecialKeyCallback(termbox.KeyEnter, doQuitButtonPress)
	connectScreen.AddSpecialKeyCallback(termbox.KeyEsc, doQuitButtonPress)
	connectButton.AddSpecialKeyCallback(termbox.KeyEnter, doConnectButtonPress)

	hostFieldWidget.UseDefaultKeys(true)
	portFieldWidget.UseDefaultKeys(true)
	nickFieldWidget.UseDefaultKeys(true)

	connectScreen.AddSpecialKeyCallback(termbox.KeyTab, switchConnectScreenActiveWidget)

	// These are the selectable widgets, the order they appear in affects
	// the Screen.SelectNext() order
	connectScreen.AddWidget(nickFieldWidget)
	connectScreen.AddWidget(portFieldWidget)
	connectScreen.AddWidget(hostFieldWidget)
	connectScreen.AddWidget(connectButton)
	connectScreen.AddWidget(quitButton)

	// These aren't selectable so this doesn't matter
	connectScreen.AddWidget(connectTitleBar)
	connectScreen.AddWidget(hostFieldLabel)
	connectScreen.AddWidget(portFieldLabel)
	connectScreen.AddWidget(nickFieldLabel)
	connectScreen.AddWidget(connectMsgWidget)

}
