package main

import "jqueue/internal/database"

type NewJobRequest struct {
	Username string `json:"username"`
}

type Config struct {
	DB *database.Queries
}
