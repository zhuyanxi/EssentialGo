package main

import "time"

// HolderOne holds a string value in a concurrency-safe manner
type HolderOne interface {
	Get() string
	Set(string)
}

type chanHolderOne struct {
	chVal chan string

	// identify whether the chan is needed to be closed or not
	// chClose chan struct{}
}

// NewChanHolderOne return a new Holder backed by Channels
func NewChanHolderOne() HolderOne {
	h := chanHolderOne{
		chVal: make(chan string),
		//chClose: make(chan struct{}),
	}
	//go h.mux()
	return h
}

func (h chanHolderOne) mux() {
	var tmp string
	for {
		select {
		case tmp = <-h.chVal:
			h.chVal <- tmp
		}
	}
}

func (h chanHolderOne) Set(s string) {
	time.Sleep(2 * time.Millisecond)
	h.chVal <- s
}

func (h chanHolderOne) Get() string {
	time.Sleep(2 * time.Millisecond)
	return <-h.chVal
}

// var errTimeoutOne = errors.New("timeout waiting for value")

// // GetWithTimeout attempts to get the value, or returns errTimeout if getting it takes too long.
// func (h chanHolderOne) GetWithTimeout(d time.Duration) (string, error) {
// 	select {
// 	case <-time.After(d):
// 		return "", errTimeout
// 	case v := <-h.chGetVal:
// 		return v, nil
// 	}
// }

// func (h chanHolder) Close() {
// 	close(h.closeChan)
// }
