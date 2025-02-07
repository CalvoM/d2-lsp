package lsp

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Encodes the messages by adding the header part
// and split content
func Encode(content any) string {
	data, err := json.Marshal(content)
	if err != nil {
		// Handle error gracefully
		panic(err)
	}
	return fmt.Sprintf("Content-Length:%d\r\n\r\n%s", len(data), data)
}

func Split(data []byte, atEOF bool) (advance int, token []byte, err error) {
	header, content, found := bytes.Cut(data, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		// To continue reading more data
		return 0, nil, nil
	}
	contentLengthBytes := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return 0, nil, err
	}

	if len(content) < contentLength {
		return 0, nil, nil
	}

	totalLength := len(header) + len("\r\n\r\n") + contentLength
	return totalLength, data[:totalLength], nil
}

func DecodeMessage(msg []byte) (Method, []byte, error) {
	header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return "", nil, errors.New("Did not find separator")
	}

	// Content-Length: <number>
	contentLengthBytes := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return "", nil, err
	}

	var baseMessage Request
	if err := json.Unmarshal(content[:contentLength], &baseMessage); err != nil {
		return "", nil, err
	}

	return baseMessage.Method, content[:contentLength], nil
}

func sendResponse(message any) {
	encodedData := Encode(message)
	os.Stdout.Write([]byte(encodedData))
}
