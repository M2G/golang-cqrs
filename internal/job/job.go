package job

import (
	"context"
	"errors"
)

type VideoJob struct {
	ID       int64
	Filename string
}

type JobStatus string

const (
	StatusPending JobStatus = "PENDING"
	StatusDone    JobStatus = "DONE"
)

type VideoJobDetail struct {
	VideoJob
	Status JobStatus
}

var ErrNotFound = errors.New("job not found")

type Repository interface {
	CreateJob(ctx context.Context, filename string) (VideoJob, error)
	MarkDone(ctx context.Context, id int64) error
	Exists(ctx context.Context, id int64) (bool, error)
	GetJobByID(ctx context.Context, id int64) (VideoJobDetail, error)
}
