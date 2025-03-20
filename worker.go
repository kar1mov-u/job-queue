package main

import (
	"context"
	"log"
)

func worker(ctx context.Context, jobChan chan WorkerData, id int) {
	for {
		select {
		case <-ctx.Done():
			// log.Printf("Worker with ID : %d is shutting down\n", id)
			return
		case job := <-jobChan:
			log.Printf("Worker with ID : %d is procesing job: %d \n", id, job.workID)

		}
	}

}
