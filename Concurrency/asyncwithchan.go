package main

import (
	"errors"
	"fmt"
	"time"
)

// Holder holds a string value in a concurrency-safe manner
type Holder interface {
	Get() string
	Set(string)
}

type chanHolder struct {
	chSetVal chan string
	chGetVal chan string

	// identify whether the chan is needed to be closed or not
	closeChan chan struct{}
	closeFlag chan bool
}

// NewChanHolder return a new Holder backed by Channels
func NewChanHolder() Holder {
	h := chanHolder{
		chSetVal:  make(chan string),
		chGetVal:  make(chan string),
		closeChan: make(chan struct{}),
		closeFlag: make(chan bool, 1),
	}
	//h.closeFlag <- false
	go h.mux()
	return h
}

func (h chanHolder) mux() {
	var val string
	for {
		// make sure we only accept Set() or Close() when val is empty
		// this means calling Get() will block until a value has been previously Set()
		if val == "" {
			select {
			case <-h.closeChan: // handle closing
				close(h.chSetVal)
				close(h.chGetVal)
				return
			case val = <-h.chSetVal:
				continue
			}
			// val = <-h.chSetVal
			// continue
		}

		select {
		case val = <-h.chSetVal:
		case h.chGetVal <- val:
		case <-h.closeChan: // handle closing, time to clean up
			close(h.chSetVal)
			close(h.chGetVal)
			return
		}
	}
}

func (h chanHolder) Set(s string) {
	// if h.closeFlag {
	// 	fmt.Println("Do not set value after the holder is closed")
	// 	return
	// }
	// h.chSetVal <- s
	select {
	case f := <-h.closeFlag:
		if f {
			fmt.Println("Do not set value after the holder is closed")
		}
	case h.chSetVal <- s:
	}
}

func (h chanHolder) Get() string {
	//return <-h.chGetVal

	select {
	case f := <-h.closeFlag:
		if f {
			return "Can not get value after the holder is closed"
		}
	case s := <-h.chGetVal:
		return s
	}
	return ""
}

var errTimeout = errors.New("timeout waiting for value")

// GetWithTimeout attempts to get the value, or returns errTimeout if getting it takes too long.
func (h chanHolder) GetWithTimeout(d time.Duration) (string, error) {
	select {
	case <-time.After(d):
		return "", errTimeout
	case v := <-h.chGetVal:
		return v, nil
	}
}

func (h chanHolder) Close() {
	//<-h.closeFlag
	h.closeFlag <- true
	close(h.closeChan)
}
