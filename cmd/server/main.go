package main

import (
	"github.com/t0mmylam/chat-cli/pkg/chat"
)

func Serve(server *chat.Server) {
	http.Handle("/", websocket.Server{
		Handler: server.Serve(),
		Handshake: nil,
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
