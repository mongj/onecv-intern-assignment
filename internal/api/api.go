package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Paylod     json.RawMessage `json:"payload"`
	StatusCode int             `json:"status_code"`
}

type Handler = func(http.ResponseWriter, *http.Request) ([]byte, int, error)

// HTTPHandler converts the internal Handler type into a standard http.HandlerFunc.
func HTTPHandler(handler Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		response, code, err := handler(w, r)
		if err != nil {
			if code >= 500 {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"error": "Internal Server Error"}`))
				log.Println(err)
			} else {
				w.WriteHeader(code)
				w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
			}
		} else {
			w.WriteHeader(code)
			w.Write(response)
		}
	}
}
