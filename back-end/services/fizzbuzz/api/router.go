package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewFizzbuzzRouter(appRouter *mux.Router, handler FizzbuzzHandler) {
	appRouter.HandleFunc("/fizzbuzz", handler.Fizzbuzz).
		Methods(http.MethodPost, http.MethodOptions)

	appRouter.HandleFunc("/health-check", handler.HealthCheck)
}
