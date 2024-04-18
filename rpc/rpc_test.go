package rpc_test

import (
	"lsp/rpc"
	"testing"
)

func TestEncode(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"testing\":true}"
	actual := rpc.EncodeMessage(map[string]bool{"testing": true})
	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestDecode(t *testing.T) {
	incomingMessage := "Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}"
	method, content, error := rpc.DecodeMessage([]byte(incomingMessage))
	contentLength := len(content)
	if error != nil {
		t.Fatalf("Error: %s", error)
	}

	if contentLength != 15 {
		t.Fatalf("Expected: 15, Actual: %d", contentLength)
	}

	if method != "hi" {
		t.Fatalf("Expected: hi, Actual: %s", method)
	}
}
