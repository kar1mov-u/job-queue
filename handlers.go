package main

import (
	"database/sql"
	"encoding/json"
	"jqueue/internal/database"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (cfg *Config) PostTransac(w http.ResponseWriter, r *http.Request) {
	data := TransactionData{}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		respondWithErr(w, 500, err.Error())
		return
	}
	dbObj, err := cfg.DB.CreateTask(r.Context(), "pending")
	if err != nil {
		respondWithErr(w, 500, err.Error())
	}
	cfg.Channel <- WorkerData{transsaction: data, workID: dbObj}

	respondWithJson(w, 200, map[string]uuid.UUID{"id": dbObj})

}

func (cfg *Config) GetTransacs(w http.ResponseWriter, r *http.Request) {
	start := r.URL.Query().Get("from")
	end := r.URL.Query().Get("to")
	startTime, err := time.Parse(time.RFC3339, start)
	if err != nil {
		startTime = time.Now().AddDate(0, 0, -7)
	}
	endTime, err := time.Parse(time.RFC3339, end)
	if err != nil {
		endTime = time.Now()
	}

	tasksDb, err := cfg.DB.GetTasks(r.Context(), database.GetTasksParams{
		CreatedAt:   sql.NullTime{Time: startTime, Valid: true},
		CreatedAt_2: sql.NullTime{Time: endTime, Valid: true},
	})
	if err != nil {
		respondWithErr(w, 500, err.Error())
	}
	respondWithJson(w, 200, tasksDb)

}
