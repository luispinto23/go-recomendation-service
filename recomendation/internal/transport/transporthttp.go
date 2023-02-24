package transport

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/luispinto23/go-recommendation-service/internal/recommendation"
)

func NewMux(recHandler recommendation.Handler) *mux.Router {
	m := mux.NewRouter()
	m.HandleFunc("/recommendation", recHandler.GetRecommendation).Methods(http.MethodGet)
	return m
}
