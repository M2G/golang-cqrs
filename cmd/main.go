package main

import (
	"context"
	"fmt"
	"golang-cqrs/internal/config"
	"golang-cqrs/internal/cqrs"
	"golang-cqrs/internal/infrastructure/db"
	"golang-cqrs/internal/infrastructure/logger"
	"golang-cqrs/internal/infrastructure/repository"
	"golang-cqrs/internal/job"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

func main() {
	// configuration
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// logger
	log := logger.New(cfg.LogLevel)
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGINT)
	defer cancel()

	// database
	pool, err := db.NewPool(ctx, cfg.DatabaseURL)
	if err != nil {
		log.WithError(err).Fatal("failed to connect to database")
	}
	defer pool.Close()

	// repository
	repo := repository.New(pool)

}