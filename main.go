package main

import (
	"bufio"
	"fmt"
	"github.com/gabriel-comeau/SimpleChatCommon"
	"github.com/gabriel-comeau/tbuikit"
	"github.com/nsf/termbox-go"
	"net"
	"strings"
)

var (
	chatBuffer       *tbuikit.TextInputBuffer
	hostFieldBuffer  *tbuikit.TextInputBuffer
	portFieldBuffer  *tbuikit.TextInputBuffer
	nickFieldBuffer  *tbuikit.TextInputBuffer
	connectMsgBuffer *tbuikit.ColorizedStringBuffer
	messageBuffer    *tbuikit.ColorizedStringBuffer
	whoBuffer        *tbuikit.ColorizedStringBuffer
	chatUi           *tbuikit.UI
	mainScreen       *tbuikit.Screen
	connectScreen    *tbuikit.Screen
	netChatChan      chan string
	netMessageChan   chan string
	connected        *ConnectionStatus
	conn             net.Conn
)

// Create new buffers - these are what hold the data that will be
// sent/recieved from the network and printed to the screen via
// the widgets.
func initBuffers() {
	chatBuffer = new(tbuikit.TextInputBuffer)
	hostFieldBuffer = new(tbuikit.TextInputBuffer)
	portFieldBuffer = new(tbuikit.TextInputBuffer)
	nickFieldBuffer = new(tbuikit.TextInputBuffer)
	hostFieldBuffer.SetLength(32)
	portFieldBuffer.SetLength(8)
	nickFieldBuffer.SetLength(24)
	messageBuffer = new(tbuikit.ColorizedStringBuffer)
	whoBuffer = new(tbuikit.ColorizedStringBuffer)
	connectMsgBuffer = new(tbuikit.ColorizedStringBuffer)
	messageBuffer.Prepare(64)
	connectMsgBuffer.Prepare(64)
}

// Automatically trigger the creation of the buffers and widgets when program starts
func init() {
	initBuffers()
	initUi()

	connected = new(ConnectionStatus)
	connected.Init()

	netMessageChan = make(chan string)
	netChatChan = make(chan string)
}

func main() {
	quitChan := make(chan bool)

	// Start the UI up here
	go chatUi.Start(quitChan)
	for {
		quitSig := <-quitChan
		if quitSig {
			break
		}
	}
}

// Establishes the connection to the chat serrver and gets the network-handler
// go routines running.
func connect() {
	conn = nil
	host := hostFieldBuffer.GetLines(0, 1)[0]
	port := portFieldBuffer.GetLines(0, 1)[0]
	nick := nickFieldBuffer.GetLines(0, 1)[0]
	connection, err := getServerConnection(host, port, nick)

	// display the error in the message window of the connect screen
	if err != nil {
		connectMsg := new(tbuikit.ColorizedString)
		connectMsg.Color = termbox.ColorRed
		connectMsg.Text = err.Error()
		connectMsgBuffer.Add(connectMsg)
		return
	}

	conn = connection
	connected.Connect()

	// no error so clear the message buffer so when we come back to it there aren't
	// any old, no longer relevant messages
	connectMsgBuffer.Clear()

	// Switch screens to the main app screen
	mainScreen.Activate()
	connectScreen.Deactivate()

	if nick != "" {
		setNickName(nick)
	}

	// Start the network handler goroutines
	go addChatMessagesToBuffer(netMessageChan)
	go getNetworkMessages(netMessageChan)
	go sendChatMessages(netChatChan)
}

// Handles disconnection from the server
func disconnect() {

	conn.Close()

	if connected.Connected() {
		connected.Disconnect()

		// Prepare a message for the user in the connect screen message window
		disconnectMsg := new(tbuikit.ColorizedString)
		disconnectMsg.Color = termbox.ColorRed
		disconnectMsg.Text = "Disconnected from server!"
		connectMsgBuffer.Add(disconnectMsg)

		// Clear the main screen's buffers so on reconnect it gets cleaned up
		messageBuffer.Clear()
		chatBuffer.Clear()

		// Switch the screens
		mainScreen.Deactivate()
		connectScreen.Activate()
	}
}

// Establishes connection to the chat server
func getServerConnection(host, port, nick string) (net.Conn, error) {

	connStr := fmt.Sprintf("%v:%v", host, port)
	connection, err := net.Dial("tcp", connStr)
	if err != nil {
		return nil, err
	}

	return connection, nil
}

// Sets user nickname if it was provided on the connect
// screen
func setNickName(nick string) {
	nickCmd := fmt.Sprintf("/nick %v\n", nick)
	conn.Write([]byte(nickCmd))
}

// Read the messages being passed over the channel and push
// them into the message buffer so they can be displayed via the widget
func addChatMessagesToBuffer(netMessageChan chan string) {
	for {
		if connected.Connected() {
			rcvdStr := <-netMessageChan
			msg := SimpleChatCommon.Unpack(rcvdStr)
			if SimpleChatCommon.IsWhoMessage(msg) {
				processWhoMessage(msg)
			} else {
				messageBuffer.Add(msg)
			}
		}
	}
}

// Read the network connection and push anything coming into it
// into the channel
func getNetworkMessages(netMessageChan chan string) {
	b := bufio.NewReader(conn)
	for {
		if connected.Connected() {
			line, err := b.ReadString('\n')
			if err != nil {
				disconnect()
				break
			}
			netMessageChan <- line
		}
	}
}

// Push the contents of the chat buffer over the wire
//
// This means that when a user sees their own messages in the
// message buffer, it is that message returning from the server.
func sendChatMessages(netChatChan chan string) {
	for {
		if connected.Connected() {
			msg := <-netChatChan
			conn.Write([]byte(msg))
		}
	}
}

// Takes the "whos online" message from the server, splits it into the appropriate parts
// and writes it to the who buffer.  Clears the who buffer first
func processWhoMessage(msg *tbuikit.ColorizedString) {

	colors := new(SimpleChatCommon.ColorList)

	// Each who message starts with the special token, so strip it off
	raw := msg.Text
	trimmed := strings.Trim(raw, " \n")
	parts := strings.Split(trimmed, SimpleChatCommon.WHO_MESSAGE_TOKEN)
	if len(parts) > 0 {
		whoBuffer.Clear()
		for _, part := range parts {
			col := colors.FromColor(msg.Color)
			bufMsg := SimpleChatCommon.Create(part, col)
			whoBuffer.Add(bufMsg)
		}
	}
}
