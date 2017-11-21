package surgeon

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

// TRANSPORTS ********
// decodeGetSurgeonRequest exposes our service to the world (in this case) via JSON over HTTP
func decodeGetSurgeonRequest(_ context.Context, r *http.Request) (interface{}, error) {
	log.Println("decodeGetSurgeonRequest")
	var request getSurgeonRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeGetSurgeonResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	log.Println("encodeGetSurgeonResponse")
	return json.NewEncoder(w).Encode(response)
}
