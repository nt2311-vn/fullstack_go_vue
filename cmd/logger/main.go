package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	startLogServer()
}

func logHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading body", http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()
	fmt.Println("Received log:", string(body))

	w.WriteHeader(http.StatusOK)
}

func startLogServer() {
	http.HandleFunc("/log", logHandler)
	log.Println("Log server running as port: 8010")
	log.Fatal(http.ListenAndServe(":8010", nil))
}
