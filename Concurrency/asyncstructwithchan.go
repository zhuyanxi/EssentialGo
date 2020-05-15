package main

import "time"

// HolderStruct holds a string value in a concurrency-safe manner
// type HolderStruct interface {
// 	Get() string
// 	Set(string, []byte)
// }

type item struct {
	key   string
	value interface{}
}

type ChanHolderStruct struct {
	// value  *item
	chFlag chan *item

	// identify whether the chan is needed to be closed or not
	// chClose chan struct{}
}

// NewChanHolderStruct return a new Holder backed by Channels
func NewChanHolderStruct() *ChanHolderStruct {
	h := &ChanHolderStruct{
		// value:  &item{},
		chFlag: make(chan *item),
		//chClose: make(chan struct{}),
	}
	return h
}

func (h *ChanHolderStruct) Set(key string, s []byte) {
	time.Sleep(2 * time.Millisecond)
	i := &item{
		key:   key,
		value: s,
	}
	// h.value = i
	h.chFlag <- i
}

func (h *ChanHolderStruct) Get() string {
	time.Sleep(2 * time.Millisecond)
	i := <-h.chFlag
	// i := h.value
	return i.key + ":" + string(i.value.([]byte))
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
