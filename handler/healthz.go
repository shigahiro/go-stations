package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/TechBowl-japan/go-stations/model"
)

// A HealthzHandler implements health check endpoint.
func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	healthResponse := model.HealthzResponse{Message: "OK"}
	http.HandleFunc("/healthz", HealthzHandler)
	if err := json.NewEncoder(w).Encode(healthResponse); err != nil {
		log.Println(err)
	}
}

// NewHealthzHandler returns HealthzHandler based http.Handler.
// func NewHealthzHandler() *HealthzHandler {
// 	return &HealthzHandler{}
// }

// ServeHTTP implements http.Handler interface.
