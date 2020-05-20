package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"sync/atomic"
	"time"
)

func main() {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command("go", "version")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}

	var timedOut int32
	timeout := 1 * time.Millisecond
	stopTimer := time.AfterFunc(timeout, func() {
		cmd.Process.Kill()
		atomic.StoreInt32(&timedOut, 1)
	})

	err = cmd.Wait()
	stopTimer.Stop()
	didTimeout := atomic.LoadInt32(&timedOut) != 0
	if didTimeout {
		fmt.Printf("didTimeout: %v, err: %v\n", didTimeout, err)
		return
	}

	// err = cmd.Wait()
	// if err != nil {
	// 	log.Fatalf("cmd.Wait() failed with '%s'\n", err)
	// }
	out := append(stdout.Bytes(), stderr.Bytes()...)
	fmt.Printf("Output:\n%s\n", string(out))
}
