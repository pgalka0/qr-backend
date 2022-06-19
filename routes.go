package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"log"
	"net/http"
	"qr-backend/utils"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func socketHandler(w http.ResponseWriter, r *http.Request) {
	// Check Origin
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// Upgrade to websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("[Error while upgrading to websocket]: ", err)
		return
	}
	defer ws.Close()

	log.Println("New Connection")

	// Read messages
	utils.Reader(ws)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"hello": "world"})
}

func setupRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/qr", socketHandler)
	r.HandleFunc("/", homeHandler)

	handler := cors.Default().Handler(r)

	http.Handle("/", handler)
}
