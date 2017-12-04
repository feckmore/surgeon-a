package surgeon

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewHandler(svc Servicer) http.Handler {
	r := mux.NewRouter()

	r.Handle("/surgeon/{id}", MakeGetSurgeonHandler(svc)).Methods("GET")

	return r
}
