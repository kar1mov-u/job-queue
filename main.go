package main

import (
	"context"
	"database/sql"
	"jqueue/internal/database"
	"log"
	"net/http"
	"os"
	"os/signal"

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
	ch := make(chan WorkerData, 10)
	cfg := Config{
		DB:      dbQueries,
		Channel: ch,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := 1; i < 6; i++ {
		go worker(ctx, cfg.Channel, i)
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Workinf"))
	})
	router.Post("/transaction", cfg.PostTransac)
	router.Get("/transactions", cfg.GetTransacs)

	srv := http.Server{
		Handler: router,
		Addr:    ":9090",
	}
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		<-stop
		log.Println("Killing all workers...")
		cancel()
		close(cfg.Channel)
		os.Exit(1)
	}()

	log.Println("listening on :9090")
	srv.ListenAndServe()
}
