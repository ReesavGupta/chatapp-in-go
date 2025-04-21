package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ReesavGupta/chatapp-in-go/src/ws"
)

type JoinRequestBody struct { // needs to be in caps to be exported
	Name string `json:"name"` // will not be filled during JSON unmarshaling if it was name
	Room string `json:"room"`
}

func RootHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)

	decoder.DisallowUnknownFields() // !important to disallow unknown fields

	defer r.Body.Close()

	var joinReq JoinRequestBody

	if err := decoder.Decode(&joinReq); err != nil {
		http.Error(w, "Invalid request: "+err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("this is the name:%v, and room :%v", joinReq.Name, joinReq.Room)
}

func HandleWsConnection(hub *ws.Hub) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ws.ServeWs(hub, w, r)
	}
}
