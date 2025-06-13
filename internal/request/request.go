package request

import (
	"fmt"
	"io"
	"strings"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	lines := string(b)
	// TODO: finish this 4-2
	return nil, nil
}

func parseRequestLine(line string) (RequestLine, error) {
	parts := strings.Split(line, " ")
	if len(parts) != 3 {
		return RequestLine{}, fmt.Errorf("invalid request line: %s", line)
	}

	// verify that method is using only capitalized alphabetic characters
	if len(parts[0]) == 0 || strings.ToUpper(parts[0]) != parts[0] {
		return RequestLine{}, fmt.Errorf("invalid method in request line: %s", parts[0])
	}

	return RequestLine{
		Method:        parts[0],
		RequestTarget: parts[1],
		HttpVersion:   parts[2],
	}, nil
}
