package main

import (
	"log"
	"net/http"

	"github.com/ReesavGupta/chatapp-in-go/src/controllers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", controllers.RootHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
