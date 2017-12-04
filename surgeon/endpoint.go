package surgeon

import (
	"context"
	"log"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

// getSurgeonRequest defines the shape of the incoming data for a getSurgeon request
type getSurgeonRequest struct {
	ID int `json:"id"`
}

// getSurgeonResponse defines the shape of the outgoing data for a response to a getSurgeon request
type getSurgeonResponse struct {
	Surgeon *Surgeon `json:"surgeon"`
	Err     string   `json:"err,omitempty"` // errors don't JSON-marshal, so use string
}

// MakeGetSurgeonHandler returns a server that fulfills the http Handler interface, tying together
// the request, service method & response
func MakeGetSurgeonHandler(svc Servicer) http.Handler {
	return httptransport.NewServer(makeGetSurgeonEndpoint(svc), decodeGetSurgeonRequest, encodeGetSurgeonResponse)
}

// makeGetSurgeonEndpoint is an adapter to convert the StringServicer's Uppercase method to a Go-Kit endpoint
func makeGetSurgeonEndpoint(svc Servicer) endpoint.Endpoint {
	log.Println("makeGetSurgeonEndpoint", svc)
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		log.Println("getSurgeonHandler: request.(getSurgeonRequest)", ctx.Err(), request)
		req, ok := request.(getSurgeonRequest)
		if !ok {
			panic("invalid request")
		}
		log.Println("getSurgeonHandler: send to service", ID(req.ID))
		surgeon, err := svc.GetSurgeonByID(ctx, ID(req.ID))
		if err != nil {
			return getSurgeonResponse{surgeon, err.Error()}, nil
		}

		log.Println("returning getSurgeonResponse from getSurgeonHandler:", surgeon)
		return getSurgeonResponse{surgeon, ""}, nil
	}
}
