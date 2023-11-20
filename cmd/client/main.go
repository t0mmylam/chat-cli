package main

import (
	"github.com/t0mmylam/chat-cli/pkg/chat"
	"github.com/t0mmylam/chat-cli/pkg/transport"
)

func main() {
	server := chat.NewServer()
	transport.Serve(server)
}