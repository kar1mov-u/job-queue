package main

import (
	"context"
	"database/sql"
	"fmt"
	"jqueue/internal/database"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"

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
	//Create SQLC instance
	dbQueries := database.New(db)
	//queue to tasks
	ch := make(chan WorkerData, 10)
	cfg := Config{
		DB:      dbQueries,
		Channel: ch,
	}

	//context to workers
	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s3Client, err := createS3Client(os.Getenv("AWS_REGION"))
	if err != nil {
		log.Fatal("Failed on creating s3 client: ", err)
	}

	//initaiting workers
	workers := make([]*Worker, 5)
	for i := 1; i < 6; i++ {
		wg.Add(1)
		fmt.Println("starting worker ", i)
		workers[i-1] = NewWorker(i, ctx, dbQueries, cfg.Channel, &wg, s3Client)
		go workers[i-1].Run()
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Workinf"))
	})
	router.Post("/transaction", cfg.PostTransac)
	router.Get("/transactions", cfg.GetTransacs)
	router.Get("/transactions/{TaskId}", cfg.GetTask)

	srv := http.Server{
		Handler: router,
		Addr:    ":9090",
	}

	// handling shutdonw
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		<-stop
		log.Println("Killing all workers...")
		cancel()
		close(cfg.Channel)
		wg.Wait()
		os.Exit(0)
	}()

	log.Println("listening on :9090")
	srv.ListenAndServe()
}
