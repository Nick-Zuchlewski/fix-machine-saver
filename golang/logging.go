package main

import (
	"io/ioutil"
	"log"
	"os"
)

// Global loggers since didnt feel like writing my own and didnt want
// to use a complex lib. Based loosely on ardanlabs.
var (
	Trace *log.Logger
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
)

// Using the magic init function
// Logs use UTC
// https://tutorialedge.net/golang/the-go-init-function/#:~:text=The%20init%20Function,-In%20Go%2C%20the&text=These%20init()%20functions%20can,should%20pay%20close%20attention%20to.
func init() {
	Trace = log.New(ioutil.Discard,
		"TRACE: ",
		log.Ldate|log.Ltime|log.LUTC|log.Lshortfile)

	Info = log.New(os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime|log.LUTC)

	Warn = log.New(os.Stdout,
		"WARN: ",
		log.Ldate|log.Ltime|log.LUTC)

	Error = log.New(os.Stderr,
		"ERROR: ",
		log.Ldate|log.Ltime|log.LUTC)
}
