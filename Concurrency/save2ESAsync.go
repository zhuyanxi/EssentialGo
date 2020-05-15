package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func save2ESOne(st time.Time, idx int) {
	defer wg.Done()
	eT2 := time.Since(st)
	time.Sleep(time.Second)
	fmt.Printf("%d cost %s\n", idx, eT2)
}

func save2ESTwo(st time.Time) {
	defer wg.Done()
	eT2 := time.Since(st)
	time.Sleep(time.Second)
	fmt.Printf("cost %s\n", eT2)
}

func save2ESThree(st time.Time, idx int) {
	defer wg.Done()
	eT2 := time.Since(st)
	time.Sleep(time.Second)
	fmt.Printf("%d cost %s\n", idx, eT2)
}

func save2ESAsync(content map[int]string) {
	cLen := len(content)

	var wg sync.WaitGroup
	chContentKey := make(chan int, 1)

	//var mu sync.Mutex

	for i := 0; i < cLen; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			tmp := make(map[int]string)

			key := <-chContentKey
			tmp[key] = content[key]
			//Save2ES(tmp)
			fmt.Printf("key: %d, val: %s\n", key, tmp[key])
		}()
	}

	go func() {
		for key := range content {
			chContentKey <- key
		}
		close(chContentKey)
	}()

	wg.Wait()
	fmt.Printf("%s events were written into ES.\n", strconv.Itoa(cLen))
}
