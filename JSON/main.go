package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	//testAnonymousStructs()

	// i := 1
	// defer logNum(i) // deferred function call: logNum(1)
	// fmt.Println("First main statement")
	// i++
	// defer logNum(i)     // deferred function call: logNum(2)
	// defer logNum(i * i) // deferred function call: logNum(4)
	// fmt.Printf("Num in main %d\n", i)

	// for i := 0; i < 2; i++ {
	// 	defer func() {
	// 		fmt.Printf("%d\n", i)
	// 	}()
	// }

	chIn := make(chan int)
	chOut := make(chan int)
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go sqrtWorker(chIn, chOut)
	}

	go func() {
		chIn <- 2
		chIn <- 4
		close(chIn)
	}()

	go func() {
		wg.Wait()
		close(chOut)
	}()

	for sqrt := range chOut {
		fmt.Printf("Got sqrt: %d\n", sqrt)
	}
}

func sqrtWorker(chIn chan int, chOut chan int) {
	fmt.Printf("sqrtWorker started\n")
	for i := range chIn {
		sqrt := i * i
		chOut <- sqrt
	}
	fmt.Printf("sqrtWorker finished\n")
	wg.Done()
}

func logNum(i int) {
	fmt.Printf("Num %d\n", i)
}
