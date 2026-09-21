package cqrs

import "context"

type CommandHandler[C any, R any] func(ctx context.Context, cmd C) (R, error)

type QueryHandler[C any, R any] func(ctx context.Context, cmd C) (R, error)
