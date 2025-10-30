package main

import (
	"bytes"
	"fmt"
)

func main() {
	var b bytes.Buffer

	// Initializing buffers
	var _ *bytes.Buffer = new(bytes.Buffer)
	data := []byte("Hello Golang")
	var _ *bytes.Buffer = bytes.NewBuffer(data) // Creating and initializing buffer with content

	// Writing strings to a buffer
	// The buffer grows dynamically as you write to it
	b.WriteString("abc")
	b.WriteString("def") // You can concatanate strings by continuously writing strings to a buffer

	// Printing a buffer
	fmt.Fprintf(&b, " A number: %d, a string %s\n", 10, "bird")
	b.WriteString("[done]")

	// Converting bytes in buffer to string
	fmt.Println(b.String())
	// Reading from a buffer
}
