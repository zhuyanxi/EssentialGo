package main

import "fmt"

func goroutineA(a <-chan int) {
	val := <-a
	fmt.Println("G1 received data: ", val)
	return
}

func goroutineB(b <-chan int) {
	val := <-b
	fmt.Println("G2 received data: ", val)
	return
}
func goB(b int, c int) {
	fmt.Printf("G2 received data: %d; square: %d\n", b, c)
	return
}

func testGoAB() {
	count := 100
	ch := make(chan int, 1)
	//ch2 := make(chan int, 1)
	semaphore := 100
	round := count / semaphore
	if count%semaphore != 0 {
		round++
	}
	for i := 0; i < semaphore; i++ {
		//wg.Add(1)
		go func() {
			//defer wg.Done()
			for i := 0; i < round; i++ {
				goroutineB(ch)
				//val := <-ch
				//val2 := <-ch2
				//goB(val, val2)
			}
		}()
	}
	for i := 0; i < count; i++ {
		ch <- i
		//ch2 <- i * i
		// go func(j int) {
		// 	ch <- j
		// 	ch2 <- j * j
		// }(i)
	}
}
