package api

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var responsesCh chan http.ResponseWriter

func Listen(port int, responses chan http.ResponseWriter) error {
	responsesCh = responses
	r := mux.NewRouter()
	r.Path("/api/ping").Methods(http.MethodPost).HandlerFunc(pingsHandler)
	return http.ListenAndServe(":"+strconv.Itoa(port), r)
}

func pingsHandler(w http.ResponseWriter, _ *http.Request) {
	responsesCh <- w
}
