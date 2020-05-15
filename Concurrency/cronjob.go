package main

import (
	"fmt"
	"time"
)

func cronjob() {
	ticker := time.Tick(1 * time.Second)
	for {
		select {
		case <-ticker:
			totalRead++
			if totalRead > 10 {
				return
			}
			fmt.Println("run 1s cronjob:", totalRead)
		}
	}
}
