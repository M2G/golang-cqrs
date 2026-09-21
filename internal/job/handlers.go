package job

import (
	"context"
	"fmt"
	"golang-cqrs/internal/cqrs"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreatedJob cqrs.CommandHandler[CreatedJobCommand, VideoJob]
	MarDone    cqrs.CommandHandler[MarkDoneCommand, struct{}]
}

type Queries struct {
	GetJobByID cqrs.QueryHandler[GetJobByIDQuery, VideoJobDetail]
}

func NewApplication(repo Repository, eventBus *cqrs.EventBus) *Application {
	createdJob := func(ctx context.Context, cmd CreatedJobCommand) (VideoJob, error) {
		videoJob, err := repo.CreateJob(ctx, cmd.Filename)
		if err != nil {
			return VideoJob{}, err
		}

		cqrs.Publish(eventBus, Created{
			job: videoJob,
		})

		return videoJob, nil
	}

	marDone := func(ctx context.Context, cmd MarkDoneCommand) (struct{}, error) {
		exists, err := repo.Exists(ctx, cmd.JobID)
		if err != nil {
			return struct{}{}, fmt.Errorf("exists check failed: %w", err)
		}
		if !exists {
			return struct{}{}, ErrNotFound
		}
		if err := repo.MarkDone(ctx, cmd.JobID); err != nil {
			return struct{}{}, err
		}
		cqrs.Publish(eventBus, Completed{
			cmd.JobID,
		})

		return struct{}{}, nil

	}
}
