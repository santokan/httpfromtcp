package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

const inputFile = "messages.txt"

func main() {
	f, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	fmt.Println("Reading messages from file:", inputFile)

	var currentLine string
	for {
		b := make([]byte, 8, 8)
		n, err := f.Read(b)
		if err != nil {
			if currentLine != "" {
				fmt.Printf("read: %s\n", currentLine)
				currentLine = ""
			}
			if errors.Is(err, io.EOF) {
				break
			}
			fmt.Printf("error: %s\n", err.Error())
			break
		}
		str := string(b[:n])
		parts := strings.Split(str, "\n")
		for i := range len(parts) - 1 {
			currentLine += parts[i]
			fmt.Printf("read: %s\n", currentLine)
			currentLine = ""
		}
		currentLine += parts[len(parts)-1]
	}
}
