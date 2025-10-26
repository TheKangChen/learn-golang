package main

import (
	"fmt"
	"log"
	"runtime"
)

func main() {
	// Logging with `runtime.Caller()`
	customLog("Start operation")

	// Logging with log package to achieve the same result
	log.SetFlags(log.Llongfile)
	log.Print("Start operation")

	// There's also `log/slog` for structured logging
}

func customLog(message string) {
	// The `runtime` package provides an interface to Go's runtime system.
	// `runtime.Caller()` allows you to introspect the call stack of the currently executing goroutine.
	_, file, line, ok := runtime.Caller(1)
	if ok {
		fmt.Printf("%s:%d: %s\n", file, line, message)
	} else {
		fmt.Printf("%s\n", message)
	}
}
