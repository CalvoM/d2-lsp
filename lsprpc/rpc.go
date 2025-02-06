package lsprpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/CalvoM/d2-lsp/lsp"
)

type RPCHeader struct {
	ContentLength int    `json:"Content-Length"`
	ContentType   string `json:"Content-Type,omitempty"`
}

// Encodes the messages by adding the header part
// and split content
func Encode(content any) string {
	data, err := json.Marshal(content)
	if err != nil {
		// Handle error gracefully
		panic(err)
	}
	header := RPCHeader{ContentLength: len(data)}
	headerData, err := json.Marshal(header)
	if err != nil {
		// Handle error gracefully
		panic(err)
	}
	return fmt.Sprintf("%s\r\n\r\n%s", headerData, data)
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

func DecodeMessage(msg []byte) (lsp.Method, []byte, error) {
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

	var baseMessage lsp.Request
	if err := json.Unmarshal(content[:contentLength], &baseMessage); err != nil {
		return "", nil, err
	}

	return baseMessage.Method, content[:contentLength], nil
}
