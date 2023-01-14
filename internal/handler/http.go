package handler

import (
	"context"
	"encoding/json"
	"ltgc-api-create-user/internal/entity"
	"net/http"

	epkit "github.com/go-kit/kit/endpoint"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPServer(_ context.Context, endpoint epkit.Endpoint) http.Handler {
	r := mux.NewRouter()

	r.Use(commonMiddleware)

	r.Methods("POST").Path("/user").Handler(httptransport.NewServer(endpoint, DecodeRequest, encodeResponse))

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// DecodeRequest request decode
func DecodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var appRequest entity.Request

	if err := json.NewDecoder(r.Body).Decode(&appRequest); err != nil {
		return nil, err
	}

	return &appRequest, nil
}
