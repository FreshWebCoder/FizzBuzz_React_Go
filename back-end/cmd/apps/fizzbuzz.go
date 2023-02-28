package apps

import (
	"github.com/gorilla/mux"

	"github.com/starikode430/bunzz-test/cmd/config"
	"github.com/starikode430/bunzz-test/services/fizzbuzz/api"
	"github.com/starikode430/bunzz-test/services/fizzbuzz/fizzbuzz"
)

func FizzbuzzApplication(appRouter *mux.Router, appConfig config.AppConfig) {
	service := fizzbuzz.NewFizzbuzzService(appConfig)
	handler := api.NewHandler(service)
	api.NewFizzbuzzRouter(appRouter, handler)
}
