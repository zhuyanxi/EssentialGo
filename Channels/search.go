package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Search :
type Search func(query string) Result

// Result :
type Result string

var (
	web   = fakeSearch("web")
	image = fakeSearch("image")
	video = fakeSearch("video")
)

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

func google1(query string) (results []Result) {
	results = append(results, web(query))
	results = append(results, image(query))
	results = append(results, video(query))
	return
}

func google2(query string) (results []Result) {
	c := make(chan Result)
	go func() {
		c <- web(query)
	}()
	go func() {
		c <- image(query)
	}()
	go func() {
		c <- video(query)
	}()
	for i := 0; i < 3; i++ {
		result := <-c
		results = append(results, result)
	}
	return
}

func google3(query string) (results []Result) {
	c := make(chan Result)
	go func() {
		c <- web(query)
	}()
	go func() {
		c <- image(query)
	}()
	go func() {
		c <- video(query)
	}()

	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("time out")
			return
		}
	}
	return
}

func google4(query string) (results []Result) {
	c := make(chan Result)
	go func() {
		c <- first(query, fakeSearch("web1"), fakeSearch("web2"))
	}()
	go func() {
		c <- first(query, fakeSearch("image1"), fakeSearch("image2"))
	}()
	go func() {
		c <- first(query, fakeSearch("video1"), fakeSearch("video2"))
	}()

	timeout := time.After(90 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("time out")
			return
		}
	}
	return
}

// Send requests to multiple replicas, and use the first response.
func first(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) {
		c <- replicas[i](query)
	}
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

func testFakeSearch() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	// results := google1("golang")
	// results := google2("golang")
	// results := google3("golang")
	results := google4("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}

func testFirstSearch() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	result := first("golang",
		fakeSearch("replica 1"),
		fakeSearch("replica 2"))
	elapsed := time.Since(start)
	time.Sleep(time.Second)
	fmt.Println(result)
	fmt.Println(elapsed)
}
