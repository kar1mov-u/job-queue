package main

import (
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func createS3Client(region string) (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return nil, err
	}

	return s3.NewFromConfig(cfg), nil
}

func (wrkr *Worker) uploadFile(file io.Reader, filename string) error {

	_, err := wrkr.S3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("job-queue-storage"),
		Key:    aws.String(fmt.Sprintf("%v.pdf", filename)),
		Body:   file,
	})
	if err != nil {
		return err
	}

	return nil
}
