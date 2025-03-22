package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"jqueue/internal/database"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
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
	//send data to queue
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

	respTasks := []respBody{}
	for _, taskDb := range tasksDb {
		task := respBody{
			ID:          taskDb.ID.String(),
			CreatedAt:   taskDb.CreatedAt.Time,
			CompletedAt: taskDb.CompletedAt.Time,
			Status:      taskDb.Status,
			Link:        taskDb.Link.String,
		}
		respTasks = append(respTasks, task)
	}
	respondWithJson(w, 200, respTasks)

}

func (cfg *Config) GetTask(w http.ResponseWriter, r *http.Request) {
	taskIdStr := chi.URLParam(r, "TaskId")
	taskId, err := uuid.Parse(taskIdStr)
	if err != nil {
		respondWithErr(w, http.StatusBadRequest, "ID should be valid UUID")
		return
	}
	dbTask, err := cfg.DB.GetTask(r.Context(), taskId)
	if errors.Is(err, sql.ErrNoRows) {
		respondWithErr(w, 404, "no task found")
		return
	} else if err != nil {
		log.Println("Error on hadnler GetTask: ", err)
		respondWithErr(w, 500, err.Error())
		return
	}
	resp := respBody{
		ID:          dbTask.ID.String(),
		CreatedAt:   dbTask.CreatedAt.Time,
		CompletedAt: dbTask.CompletedAt.Time,
		Status:      dbTask.Status,
		Link:        dbTask.Link.String,
	}
	respondWithJson(w, 200, resp)
}
