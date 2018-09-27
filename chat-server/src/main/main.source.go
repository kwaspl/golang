package main

import (
	"chat-server/src/chat"
	"net/http"
)

func main() {
	http.HandleFunc("/api", chat.GetAllChatUsers)
	http.HandleFunc("/ws", chat.WsHandler)

	http.ListenAndServe(":8080", nil)
}
