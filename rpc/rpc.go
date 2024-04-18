package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

func EncodeMessage(msg any) string {
	content, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)
}

type BaseMessage struct {
	Method string `json:"method"`
}

func DecodeMessage(msg []byte) (string, []byte, error) {
	header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return "", nil, errors.New("Did not find separator")
	}
	contentLengthBytes := header[len("Content-Length: "):]
	contentLength, error := strconv.Atoi(string(contentLengthBytes))
	if error != nil {
		return "", nil, error
	}

	_ = content
	baseMessage := BaseMessage{}

	if error := json.Unmarshal(content, &baseMessage); error != nil {
		return "", nil, error
	}

	return baseMessage.Method, content[:contentLength], nil
}

func Split(data []byte, _ bool) (advance int, token []byte, err error) {
	header, content, found := bytes.Cut(data, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return 0, nil, nil
	}
	contentLengthBytes := header[len("Content-Length: "):]
	contentLength, error := strconv.Atoi(string(contentLengthBytes))
	if error != nil {
		return 0, nil, error
	}

	if len(content) < contentLength {
		return 0, nil, nil
	}

	totalLength := len(content) + 4 + len(header)
	return totalLength, data[:totalLength], nil
}
