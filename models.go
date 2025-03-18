package main

import "database/sql"

type NewJobRequest struct {
	Username string `json:"username"`
}

type Config struct {
	Db *sql.DB
}
