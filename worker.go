package main

import (
	"context"
	"database/sql"
	"fmt"
	"jqueue/internal/database"
	"log"
	"sync"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func NewWorker(id int, ctx context.Context, db *database.Queries, ch chan WorkerData, w *sync.WaitGroup, s3 *s3.Client) *Worker {
	return &Worker{
		ID:       id,
		CTX:      ctx,
		DB:       db,
		JobCHan:  ch,
		Wait:     w,
		S3Client: s3,
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
				log.Printf("Worker %d failed to process transaction %v \n", wrkr.ID, job.workID.String())
				return
			}

			//create file
			pdfFIle, err := generatePDF(job)
			if err != nil {
				log.Printf("Worker %v | Failure on generating pdf in Task : %v  | Err: %v \n", wrkr.ID, job.workID.String(), err.Error())
				return
			}

			//store in s3

			err = wrkr.uploadFile(pdfFIle, job.workID.String())
			if err != nil {
				log.Printf("Worker %v| Failed on uploaidng job %v | Err: %v \n", wrkr.ID, job.workID.String(), err.Error())
				return
			}

			link := fmt.Sprintf("https://job-queue-storage.s3.eu-north-1.amazonaws.com/%v.pdf", job.workID.String())

			// change db status
			err = wrkr.DB.CompleteTask(wrkr.CTX, database.CompleteTaskParams{
				Status: "completed",
				Link:   sql.NullString{String: link, Valid: true},
				ID:     job.workID,
			})

			if err != nil {
				log.Printf("Worker %d | Failed on updating DB stauts job %v  | Err: %v \n", wrkr.ID, job.workID, err.Error())
				return
			}
		}
	}

}
