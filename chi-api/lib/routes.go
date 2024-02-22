package lib

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func AddRoutes(router chi.Router) {
	router.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome. add /name for personalized hello"))
	})

	router.Get("/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		w.Write([]byte(fmt.Sprintf("welcome %s!", name)))
	})
}
