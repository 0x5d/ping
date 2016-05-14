package api

import (
	"fmt"
	"net/http"
	"strconv"
)

var listenersCh chan chan []byte

func Listen(port int, listeners chan chan []byte) error {
	listenersCh = listeners
	http.HandleFunc("/api/pings", pingsHandler)
	return http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

func pingsHandler(w http.ResponseWriter, r *http.Request) {
	message := make(chan []byte)
	if r.Method != http.MethodPost {
		w.WriteHeader(404)
		return
	}
	listenersCh <- message
	w.Write([]byte(fmt.Sprintf("%s\n", <-message)))
}
