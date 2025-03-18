package main

import (
	"database/sql"
	"jqueue/internal/database"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()
	DB_URL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", DB_URL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)

	cfg := Config{DB: dbQueries}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Workinf"))
	})
	router.Post("/job", cfg.PostJob)
	router.Get("/jobs", cfg.GetTasks)

	srv := http.Server{
		Handler: router,
		Addr:    ":9090",
	}

	log.Println("listening on :9090")
	srv.ListenAndServe()
}
