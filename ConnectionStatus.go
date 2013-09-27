package main

import (
	"sync"
)

// A simple "atomic" boolean type to make sure
// that checks to see if the app is connected are
// thread safe.
type ConnectionStatus struct {
	isConnected bool
	lock        *sync.RWMutex
}

// Initialize - set up the mutex
func (this *ConnectionStatus) Init() {
	this.lock = new(sync.RWMutex)
}

// Get the current status
func (this *ConnectionStatus) Connected() bool {
	this.lock.RLock()
	defer this.lock.RUnlock()
	return this.isConnected
}

// Set the connected flag to true
func (this *ConnectionStatus) Connect() {
	this.lock.Lock()
	this.isConnected = true
	this.lock.Unlock()
}

// Set the connected flag to false
func (this *ConnectionStatus) Disconnect() {
	this.lock.Lock()
	this.isConnected = false
	this.lock.Unlock()
}
