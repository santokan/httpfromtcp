package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

const inputFile = "messages.txt"

func main() {
	f, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	// read messages 8bits at a time and print them
	for {
		b := make([]byte, 8, 8)
		n, err := f.Read(b)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			fmt.Printf("error: %s\n", err.Error())
			break
		}
		str := string(b[:n])
		fmt.Printf("read: %s\n", str)
	}
}
