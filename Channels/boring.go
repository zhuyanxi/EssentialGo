package main

import (
	"fmt"
	"math/rand"
	"time"
)

func testBoring1() {
	joe := boring("Joe!")
	ann := boring("Ann")
	for i := 0; i < 5; i++ {
		// fmt.Printf("You say: %q\n", <-joe)
		fmt.Println(<-ann)
		fmt.Println(<-joe)
	}
	fmt.Println("You are boring; I am leaving.")
}

func testBoring2() {
	c := fanIn(boring("Joe!"), boring("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You are boring; I am leaving.")
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

type message struct {
	str  string
	wait chan bool
}
