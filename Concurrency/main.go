package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

// declare global WaitGroup value is the easiest way to make it visible to all functions and methods
var wg sync.WaitGroup
var sm sync.Mutex

var (
	// used to limit the number of goroutines
	// a buffered channel is a natural semaphore
	semaphoreSize = runtime.NumCPU()

	mu                 sync.Mutex
	totalTasks         int
	curConcurrentTasks int
	maxConcurrentTasks int
)

var (
	testMap           = make(map[int]int)
	testConcurrentMap = ConcurrentMap{
		m: make(map[string]int),
	}
	totalRead = int64(0)

	v = atomic.Value{}

	chPrint      = make(chan bool)
	printCounter = 0
)

func worker(ch chan int, chQuit chan struct{}) {
	for {
		select {
		case v := <-ch:
			testMap[v] = v * v
			fmt.Printf("Got value %d\n", v)
		case <-chQuit:
			fmt.Printf("Signalled on quit channel. Finishing\n")
			chQuit <- struct{}{}
			return
		}
	}
}

func consumer(task, id int) {
	tt := time.Now()
	//mu.Lock()
	totalRead++
	//atomic.AddInt64(&totalRead, 1)
	//mu.Unlock()
	//testConcurrentMap.Add(strconv.Itoa(task), task*2)
	//time.Sleep(300 * time.Microsecond)
	//testMap[task] = task * 2

	fmt.Printf("finish task: %d by worker %d; cost: %s\n", task, id, time.Since(tt))
}

func main() {
	t1 := time.Now()
	concurrencysafemap := NewConcurrencySafeMap()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			concurrencysafemap.Add(strconv.Itoa(j), []byte(strconv.Itoa(j*j)))
		}(i)
	}

	// var result []string
	// holderStruct := NewChanHolderStruct()
	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		tmp := holderStruct.Get()
	// 		fmt.Println(tmp)
	// 		result = append(result, tmp)
	// 	}()
	// }
	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go func(j int) {
	// 		defer wg.Done()
	// 		holderStruct.Set(strconv.Itoa(j), []byte(strconv.Itoa(j*j)))
	// 	}(i)
	// }

	// holder1 := NewChanHolder().(chanHolder)
	// holder1.Close()
	// //holder1.Set("100")
	// fmt.Println(holder1.Get())

	// holder := NewChanHolderOne()
	// go func() {
	// 	fmt.Println(holder.Get())
	// }()
	//fmt.Println(holder.Get())
	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		fmt.Println(holder.Get())
	// 	}()
	// }
	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go func(j int) {
	// 		defer wg.Done()
	// 		holder.Set(strconv.Itoa(j))
	// 	}(i)
	// }
	// for i := 0; i < 10; i++ {
	// 	holder.Set(strconv.Itoa(i))
	// }
	// wg.Add(2)
	// go func() {
	// 	defer wg.Done()
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Println(holder.Get())
	// 	}
	// }()
	// go func() {
	// 	defer wg.Done()
	// 	for i := 0; i < 10; i++ {
	// 		holder.Set(strconv.Itoa(i))
	// 	}
	// }()

	// for i := 0; i < 10; i++ {
	// 	holder.Set(strconv.Itoa(i))
	// 	fmt.Println(holder.Get())
	// }
	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go func(j int) {
	// 		defer wg.Done()
	// 		holder.Set(strconv.Itoa(j))
	// 		fmt.Println(holder.Get())
	// 		// fmt.Println(holder.(chanHolder).GetWithTimeout(2 * time.Second))
	// 	}(i)
	// }

	// time.Sleep(2 * time.Second)
	// holder.Set("zhuyanxi1")
	// fmt.Println(holder.GetWithTimeout(2 * time.Second))

	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		//printConOnce("100", chPrint)
	// 		//printOnceAsync("100")
	// 		printOnce("100")
	// 	}()
	// }
	// chPrint <- true
	// for i := 0; i < 99; i++ {
	// 	chPrint <- false
	// }
	// for i := 0; i < 100; i++ {
	// 	printOnce("100")
	// }

	// ip := []string{"10.132.123.220", "10.134.41.91", "78.55.179.151"}
	// result := region(ip...)
	// fmt.Println(result)

	//cronjob()

	// timer := make(chan int)
	// select {
	// case <-time.After(1 * time.Second):
	// 	fmt.Println("do 15s job and exit")
	// case <-timer:
	// 	fmt.Println("exit after 15s")
	// }
	// ch <- struct{}{}

	// ch := make(chan int)
	// routines := 1000
	// go func() {
	// 	for i := 0; i < routines; i++ {
	// 		go func(id int) {
	// 			for {
	// 				consumer(<-ch, id)
	// 			}
	// 		}(i)
	// 	}
	// }()
	// go func(id int) {
	// 	for {
	// 		consumer(<-ch, id)
	// 	}
	// }(1)
	// for i := 0; i < 500000; i++ {
	// 	ch <- i
	// }
	// fmt.Println("Map lengh: ", testConcurrentMap.Count())
	// fmt.Println("totalRead: ", totalRead)

	//doTimeConsumingTask()
	wg.Wait()
	// sort.Strings(result)
	// fmt.Println(result)
	fmt.Println(concurrencysafemap)
	elapsed := time.Since(t1)
	fmt.Println("Time: ", elapsed)
	fmt.Println("main finished")
}

func doWork() {
	chIn := make(chan int)
	chOut := make(chan int)

	// launch 2 workers : each worker is running independently and concurrently
	for i := 0; i < 2; i++ {
		// increment wg counter : must be done in the main goroutine
		// because there is no guarantee that a newly started goroutine will execute before due to memory model guarantee
		wg.Add(1)

		go sqrtWorker(chIn, chOut)
	}

	go func() {
		for i := 2; i < 8; i++ {
			chIn <- i
		}
		close(chIn)
	}()

	go func() {
		// wait the counter to reach 0, which indicates all workers have finished and it's now safe to close the output channel
		wg.Wait()
		close(chOut)
	}()

	for sqrt := range chOut {
		fmt.Printf("Got sqrt: %d\n", sqrt)
	}
}

func sqrtWorker(chIn chan int, chOut chan int) {
	// just before terminating, decrements wg counter (wg.Add(-1))
	// by using deferred call, make sure that it will be called whenever function ends, no matter how it ends
	defer wg.Done()

	fmt.Printf("sqrtWorker started\n")
	for i := range chIn {
		sqrt := i * i
		chOut <- sqrt
	}
	fmt.Printf("sqrtWorker finished\n")

}

func mult(x, y int) {
	fmt.Printf("%d * %d = %d\n", x, y, x*y)
}

func timeConsumingTask() {
	mu.Lock()
	totalTasks++
	curConcurrentTasks++
	if curConcurrentTasks > maxConcurrentTasks {
		maxConcurrentTasks = curConcurrentTasks
	}
	mu.Unlock()

	// in real system this would be a CPU intensive operation
	time.Sleep(10 * time.Millisecond)
	//fmt.Printf("cur concurrent tasks: %d\n", curConcurrentTasks)

	mu.Lock()
	curConcurrentTasks--
	mu.Unlock()
}

func doTimeConsumingTask() {
	// a channel acting as a semaphore to limit number of goroutines active at any given time
	sem := make(chan int, semaphoreSize)

	for i := 0; i < 32; i++ {
		// acquire semaphore
		sem <- i
		wg.Add(1)

		go func() {
			timeConsumingTask()

			curNum := <-sem
			fmt.Printf("Current channel number: %d\n", curNum)
			if len(sem) == semaphoreSize {
				fmt.Println("===========================")
			}
			wg.Done()
		}()
	}
}
