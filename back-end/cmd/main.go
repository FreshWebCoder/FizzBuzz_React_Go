package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/cors"
	"github.com/gorilla/mux"

	"github.com/starikode430/bunzz-test/cmd/apps"
	"github.com/starikode430/bunzz-test/cmd/config"
)

func main() {
	time.Local = time.UTC
	appEnv := os.Getenv("APP_ENV")

	// Instantiate all app configs
	newConfig, err := config.NewConfig(appEnv, "./config")
	if err != nil {
		panic(err)
	}

	// Create a Global Router with middleware
	appRouter := mux.NewRouter()
	appRouter.Use(cors.Handler(cors.Options{
		AllowedOrigins:   newConfig.Cors.AllowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	// Run apps on separate threads to separate application workload
	go apps.FizzbuzzApplication(appRouter, newConfig.App)
	// Run the router
	addr := fmt.Sprintf(":%d", newConfig.App.Port)
	log.Printf("Server Running on port %d", newConfig.App.Port)
	//nolint:gosec //reason: will be fixed in the future.
	log.Fatal(http.ListenAndServe(addr, appRouter))
}
