package main

import (
	"log"
	"net/http"

	"github.com/ReesavGupta/chatapp-in-go/src/controllers"
	"github.com/ReesavGupta/chatapp-in-go/src/ws"
)

func main() {

	hub := ws.NewHub()
	go hub.Run()

	mux := http.NewServeMux()
	mux.HandleFunc("/", controllers.RootHandler)
	mux.HandleFunc("/init-ws", controllers.HandleWsConnection(hub))

	log.Fatal(http.ListenAndServe(":8080", mux))
}
