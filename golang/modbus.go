package main

import (
	"fmt"
	"sync"
	"time"

	modbus "github.com/goburrow/modbus"
)

type trivibe struct {
	sync.Mutex
	handler   *modbus.TCPClientHandler
	client    modbus.Client
	connected bool
}

// NewTrivibe provides a new instance of trivibe
func NewTrivibe() *trivibe {
	return &trivibe{}
}

func (t *trivibe) Configure(ip, port string, timeout uint64) (err error) {
	defer t.Unlock()
	t.Lock()

	// Configure handle
	t.handler = modbus.NewTCPClientHandler(fmt.Sprintf("%s:%s", ip, port))
	t.handler.Timeout = time.Duration(timeout) * time.Millisecond
	return
}

// Connect connects to the trivibe
func (t *trivibe) Connect() (err error) {
	defer t.Unlock()
	t.Lock()

	// Connect and create client
	if err = t.handler.Connect(); err != nil {
		return
	}

	t.client = modbus.NewClient(t.handler)
	t.connected = true

	return
}

// Disconnect disconnects from the trivibe
func (t *trivibe) Disconnect() (err error) {
	defer t.Unlock()
	t.Lock()

	// Destroy client and handle
	err = t.handler.Close()
	t.client = nil
	t.connected = false
	return
}

// ReadHoldingRegister reads a holding register from the trivibe
func (t *trivibe) ReadHoldingRegister(slaveId byte, address uint16, quantity uint16) (result []byte, err error) {
	defer t.Unlock()
	t.Lock()

	t.handler.SlaveId = slaveId
	result, err = t.client.ReadHoldingRegisters(address, quantity)
	return
}
