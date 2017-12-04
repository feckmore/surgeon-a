package surgeon

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// TRANSPORTS ********
// decodeGetSurgeonRequest exposes our service to the world (in this case) via JSON over HTTP
func decodeGetSurgeonRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return nil, err
	}
	request := getSurgeonRequest{id}

	return request, nil
}

func encodeGetSurgeonResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
