package util

import (
	"bytes"
	"runtime"
	"strconv"
)

// Ref: Google AI
func GetGID() uint64 {
	b := make([]byte, 64)
	// Write the stack trace to the buffer
	b = b[:runtime.Stack(b, false)]
	// The trace begins with "goroutine <id> ["
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	id, _ := strconv.ParseUint(string(b), 10, 64)
	return id
}
