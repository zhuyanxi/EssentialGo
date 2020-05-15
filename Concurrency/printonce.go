package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	testSet    = make(map[string]int, 0)
	conTestSet = ConcurrentMap{
		m: make(map[string]int, 0),
	}

	asyncSafeMap = sync.Map{}
)

func printOnce(num string) {
	if _, exist := testSet[num]; !exist {
		fmt.Println(num)
	}
	testSet[num] = 1
	printCounter++
	time.Sleep(2 * time.Millisecond)
}

func printConOnce(num string, ch chan bool) {
	if isPrint := <-ch; isPrint {
		if _, exist := testSet[num]; !exist {
			fmt.Println(num)
			testSet[num] = 1
		}
	}
	printCounter++
	time.Sleep(2 * time.Millisecond)
}

func printOnceAsync(num string) {
	if _, exist := asyncSafeMap.Load(num); !exist {
		fmt.Println(num)
	}
	asyncSafeMap.Store(num, 1)
	time.Sleep(2 * time.Millisecond)
}
