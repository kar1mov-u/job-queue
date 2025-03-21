package main

import (
	"context"
	"jqueue/internal/database"
	"sync"

	"github.com/google/uuid"
)

type TransactionData struct {
	SenderName      string `json:"sender"`
	SenderCard      string `json:"sender_card"`
	ReceiverName    string `json:"receiver_name"`
	ReceiverCard    string `json:"receiver_card"`
	TransactionDate string `json:"transaction_date"`
	ReceiptDate     string `json:"receipt_data"`
	TransactionId   string `json:"transaction_id"`
	Amout           string `json:"amount"`
	Commision       string `json:"commision"`
	Total           string `json:"total"`
}
type WorkerData struct {
	transsaction TransactionData
	workID       uuid.UUID
}

type Config struct {
	DB      *database.Queries
	Channel chan WorkerData
}

type Worker struct {
	ID      int
	CTX     context.Context
	DB      *database.Queries
	JobCHan chan WorkerData
	Wait    *sync.WaitGroup
}
