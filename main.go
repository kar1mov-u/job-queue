package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Workinf"))
	})
	router.Post("/job", PostJob)

	srv := http.Server{
		Handler: router,
		Addr:    ":9090",
	}

	log.Println("listening on :9090")
	srv.ListenAndServe()
}
