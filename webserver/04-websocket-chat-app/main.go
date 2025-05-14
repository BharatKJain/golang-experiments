package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"

	"log/slog"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool) // Connected clients
var broadcast = make(chan []byte)            // Broadcast channel
var mutex = &sync.Mutex{}                    // Protect clients map

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading:", err)
		return
	}
	defer conn.Close()

	mutex.Lock()
	clients[conn] = true
	mutex.Unlock()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			mutex.Lock()
			delete(clients, conn)
			mutex.Unlock()
			break
		}
		broadcast <- message
	}
}

func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		message := <-broadcast

		// Send the message to all connected clients
		mutex.Lock()
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				client.Close()
				delete(clients, client)
			}
		}
		mutex.Unlock()
	}
}

func wsStateHandler(w http.ResponseWriter, r *http.Request) {
	// On demand closure of the websocket connection
	mutex.Lock()
	slog.Info("Connected clients", slog.Any("Clients:", clients))

	for client := range clients {
		slog.Info("Client:", slog.Any("Client IP:", client.RemoteAddr().String()))
		if client.LocalAddr().String() == r.RemoteAddr {
			client.Close()
			delete(clients, client)
			slog.Info("Closed connection", slog.Any("Client:", client))
		}
	}
	mutex.Unlock()
}

func handleHearbeat() {
	for {
		// Send a heartbeat message to all connected clients
		mutex.Lock()
		for client := range clients {
			err := client.WriteMessage(websocket.PingMessage, []byte("heartbeat"))
			if err != nil {
				slog.Error("Error sending heartbeat", slog.Any("Client:", client))
				client.Close()
				delete(clients, client)
			}
		}
		mutex.Unlock()

		// Sleep for a while before sending the next heartbeat
		time.Sleep(2 * time.Second)
	}
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	http.HandleFunc("/wsState", wsStateHandler)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	go handleMessages()
	go handleHearbeat()
	fmt.Println("WebSocket server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
