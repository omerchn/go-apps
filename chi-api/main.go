package main

import (
	"chi-api/lib"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Route("/", lib.AddRoutes)

	addr := lib.GetServerAddr(8080)
	fmt.Printf("Starting server at address %v \n", addr)
	err := http.ListenAndServe(addr, router)

	if err != nil {
		panic("Server failed to start")
	}
}
