package main

import (
	"fmt"
	"sync"
	"time"
)

/***
A method is a function tied to a type, usually a struct but it works with any type.
***/

type Person struct {
	FirstName string
	LastName  string
}

func (p *Person) PrintFullName() {
	fmt.Printf("%s %s\n", p.FirstName, p.LastName)
}

func main() {
	// p := &Person{
	// 	"John",
	// 	"Doe",
	// }
	// p.PrintFullName()

	t1 := time.Now()

	//doQueryIPS()
	doQueryIPSCurrency()

	elapsed := time.Since(t1)
	fmt.Println("Time: ", elapsed)
	fmt.Println("main finished")
}

func doQueryIPS() {
	lenIPs := 50800
	for i := 0; i < lenIPs; i += 100 {
		queryIPS(i, lenIPs)
	}
}

func doQueryIPSCurrency() {
	lenIPs := 50800
	var wg sync.WaitGroup
	chStartIdx := make(chan int, 1)
	threadN := 200
	for i := 0; i < threadN; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for start := range chStartIdx {
				//idx := j
				if start+100 < lenIPs {
					time.Sleep(10 * time.Millisecond)
					fmt.Println("Read ", start, " - ", start+100)
				} else {
					time.Sleep(10 * time.Millisecond)
					fmt.Println("Read ", start, " - ", lenIPs)
				}
			}
		}()
	}

	go func() {
		for i := 0; i < lenIPs; i += 100 {
			chStartIdx <- i
		}
		close(chStartIdx)
	}()
	wg.Wait()
}

func queryIPS(start, len int) {
	if start+100 < len {
		time.Sleep(10 * time.Millisecond)
		fmt.Println("Read ", start, " - ", start+100)
	} else {
		time.Sleep(10 * time.Millisecond)
		fmt.Println("Read ", start, " - ", len)
	}
}
