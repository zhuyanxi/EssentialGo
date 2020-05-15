package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bxcodec/faker/v3"
)

func main() {
	startTime := time.Now()

	var arr1 []string
	for i := 0; i < 100; i++ {
		arr1 = append(arr1, faker.IPv4())
	}

	notHit := make(map[string]struct{})
	for i := 0; i < 100; i++ {
		notHit[faker.IPv4()] = struct{}{}
	}

	var arr2 []string
	for _, val := range arr1 {
		if _, ok := notHit[val]; ok {
			arr2 = append(arr2, val)
		}
	}

	elapsed := time.Since(startTime)
	fmt.Println(strings.Join(arr1, ","))
	fmt.Println(len(arr1), ":", elapsed)
	fmt.Println(len(arr2), ":", elapsed)
}
