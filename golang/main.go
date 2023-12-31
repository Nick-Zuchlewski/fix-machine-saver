package main

import (
	"time"
)

func main() {
	slaveIds := []byte{10, 97}

	trivibe := NewTrivibe()
	trivibe.Configure("10.224.10.117", "502", 10000)

	// Connect to trivibe
	Info.Println("Connecting to trivibe")
	if err := trivibe.Connect(); err != nil {
		Error.Println(err)
		return
	}
	defer trivibe.Disconnect()
	Info.Println("Connected to trivibe")

	for {
		for _, slaveId := range slaveIds {
			// Read holding registers
			Info.Printf("Reading holding registers from slave %d\n", slaveId)
			results, err := trivibe.ReadHoldingRegister(slaveId, 176, 3)
			if err != nil {
				Error.Println(err)
			}
			Info.Printf("Results from slave: %v, %d\n", results, slaveId)
			// Sleep for 1 second between slaves
			time.Sleep(100 * time.Millisecond)
		}
		// Sleep for 1 second between loops
		time.Sleep(1000 * time.Millisecond)
	}

}
