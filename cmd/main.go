package main

import (
	"context"
	"fmt"
	"golang-cqrs/internal/config"
	"golang-cqrs/internal/cqrs"
	"golang-cqrs/internal/infrastructure/db"
	"golang-cqrs/internal/infrastructure/logger"
	"golang-cqrs/internal/infrastructure/repository"
	"golang-cqrs/internal/interfaces"
	"golang-cqrs/internal/job"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	// event bus
	eventBus := cqrs.NewEventBus()
	registerEventLoggers(eventBus, log)

	// application
	app := job.NewApplication(repo, eventBus)

	// watcher
	watcher := interfaces.NewWatcher(cfg.StreamsDir, app)
	go watcher.Start(ctx, log)

	// http server
	httpServer := interfaces.NewHTTPServer(app, cfg.UploadDir, cfg.StreamsDir, log)

	// graceful shutdown
	go func() {
		<-ctx.Done()
		log.Info("shutting_down")

		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer shutdownCancel()

		if err := httpServer.Stop(shutdownCtx); err != nil {
			log.WithError(err).Error("http_shutdown_error")
		}

		log.Info("shutdown_complete")
	}()

	log.Info("video_orchestrator_started")

	// start http server
	if err := httpServer.Start(cfg.HTTPAddr); err != nil && err != http.ErrServerClosed {
		log.WithError(err).Fatal("http_start_error")
	}
}

func registerEventLoggers(eventBus *cqrs.EventBus, log *logrus.Logger) {
	cqrs.Subscribe(eventBus, func(e job.Created) {
		log.WithField("job_id", e.Job.ID).Info("event_job_created")
	})
	cqrs.Subscribe(eventBus, func(e job.Completed) {
		log.WithField("job_id", e.JobID).Info("event_job_completed")
	})
}
