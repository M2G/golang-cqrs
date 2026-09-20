package job

import (
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
