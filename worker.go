package main

import (
	"context"
	"jqueue/internal/database"
	"log"
	"sync"
)

func NewWorker(id int, ctx context.Context, db *database.Queries, ch chan WorkerData, w *sync.WaitGroup) *Worker {
	return &Worker{
		ID:      id,
		CTX:     ctx,
		DB:      db,
		JobCHan: ch,
		Wait:    w,
	}
}

func (wrkr *Worker) Run() {
	defer wrkr.Wait.Done()
	for {
		select {
		case <-wrkr.CTX.Done():
			log.Printf("Worker with ID : %d is shutting down\n", wrkr.ID)
			return
		case job := <-wrkr.JobCHan:

			log.Printf("Worker with ID : %d is procesing job: %v \n", wrkr.ID, job.workID.String())

			//change the db status
			err := wrkr.DB.UpdateStatus(wrkr.CTX, database.UpdateStatusParams{
				ID:     job.workID,
				Status: "in_progress",
			})
			if err != nil {
				//do some logic for failed jobs
				log.Printf("Worker %d failed to process transaction %v", wrkr.ID, job.workID.String())
			}

			//create file
			//store in s3
			// change db status

		}
	}

}
