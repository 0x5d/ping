package api

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Listen(port int) error {
	r := mux.NewRouter()
	r.Path("/api/ping").Methods(http.MethodPost).HandlerFunc(pingsHandler)
	return http.ListenAndServe(":"+strconv.Itoa(port), r)
}

func pingsHandler(w http.ResponseWriter, r *http.Request) {
	// Publish on the broker.
}
