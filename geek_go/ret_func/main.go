package main

import (
	"os"
	"runtime"
)

func main() {
	// The program terminates immediately; deferred functions are not run.
	os.Exit(0)
	// Goexit runs all deferred calls before terminating the goroutine
	runtime.Goexit()
}
