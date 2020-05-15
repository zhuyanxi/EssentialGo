package main

type ConcurrencySafeMap struct {
	m   map[string][]byte
	chV chan map[string][]byte
	// chGet chan struct{}
	// chAdd chan struct{}
}

func NewConcurrencySafeMap() *ConcurrencySafeMap {
	cm := &ConcurrencySafeMap{
		m:   make(map[string][]byte),
		chV: make(chan map[string][]byte),
		// chGet: make(chan struct{}),
		// chAdd: make(chan struct{}),
	}
	go cm.mux()
	return cm
}

func (cm *ConcurrencySafeMap) mux() {
	for {
		select {
		case <-cm.chV:
		}
	}
}

func (cm *ConcurrencySafeMap) Add(key string, value []byte) {

	dest := cm.m
	dest[key] = value
	cm.chV <- dest
	cm.m = dest
}

func (cm *ConcurrencySafeMap) Get(key string) ([]byte, bool) {
	bm := <-cm.chV
	cm.m = bm
	if b, ok := bm[key]; ok {
		return b, ok
	}
	return nil, false
}

func (cm *ConcurrencySafeMap) Remove(key string) {
	if _, ok := cm.m[key]; ok {
		delete(cm.m, key)
	}
}
