package main

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

func TestPrintOnce(t *testing.T) {
	t1 := time.Now()
	var wg1 sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg1.Add(1)
		go func() {
			defer wg1.Done()
			printOnce("123")
		}()
	}
	wg1.Wait()
	elapsed := time.Since(t1)
	fmt.Println("Time: ", elapsed)
}

func TestPrintOnceMutex(t *testing.T) {
	t1 := time.Now()
	var wg1 sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		wg1.Add(1)
		go func() {
			defer wg1.Done()
			m.Lock()
			printOnce("123")
			m.Unlock()
		}()
	}
	wg1.Wait()
	elapsed := time.Since(t1)
	fmt.Println("Time: ", elapsed)
}

func TestPrintOnceChan(t *testing.T) {
	t1 := time.Now()
	chPrint = make(chan bool)
	var wg1 sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg1.Done()
			printConOnce("123", chPrint)
		}()
	}
	chPrint <- true
	for i := 0; i < 99; i++ {
		chPrint <- false
	}

	wg1.Wait()
	elapsed := time.Since(t1)
	fmt.Println("Time: ", elapsed)
}

func BenchmarkContextSwitch(b *testing.B) {
	var wgT sync.WaitGroup
	begin := make(chan struct{})
	c := make(chan struct{})

	var token struct{}
	sender := func() {
		defer wgT.Done()
		<-begin // 1
		log.Println("Sender Begin")
		for i := 0; i < b.N; i++ {
			c <- token // 2
		}
	}
	receiver := func() {
		defer wgT.Done()
		<-begin // 1
		log.Println("Reciever Begin")
		for i := 0; i < b.N; i++ {
			<-c // 3
		}
	}

	wgT.Add(2)
	go sender()
	go receiver()
	b.StartTimer() // 4
	close(begin)   // 5
	wgT.Wait()
}
