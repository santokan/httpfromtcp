package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	l, err := net.Listen("tcp", ":42069")
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}
	defer l.Close()

	fmt.Println("Listening on :42069")
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Printf("error: %s\n", err.Error())
			os.Exit(1)
		}
		fmt.Println("Accepted connection from", conn.RemoteAddr())

		linesChan := getLinesChannel(conn)
		for line := range linesChan {
			fmt.Printf("read: %s\n", line)
		}
		fmt.Println("Connection closed by", conn.RemoteAddr())
	}
}

func getLinesChannel(conn net.Conn) <-chan string {
	lines := make(chan string)

	go func() {
		defer close(lines)

		var currentLine string
		for {
			b := make([]byte, 8)
			n, err := conn.Read(b)
			if err != nil {
				if currentLine != "" {
					lines <- currentLine
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
				lines <- currentLine
				currentLine = ""
			}
			currentLine += parts[len(parts)-1]
		}
	}()

	return lines
}
