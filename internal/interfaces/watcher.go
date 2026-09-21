package interfaces

import (
	"context"
	"errors"
	"path/filepath"
	"strconv"
	"time"

	"golang-cqrs/internal/job"

	"github.com/sirupsen/logrus"
)

type Watcher struct {
	streamsDir string
	app        *job.Application
}

func NewWatcher(streamsDir string, app *job.Application) *Watcher {
	return &Watcher{
		streamsDir: streamsDir,
		app:        app,
	}
}

func (w *Watcher) Start(ctx context.Context, log *logrus.Logger) {
	log.Info("watcher_started")

	seen := make(map[string]struct{})

	for {
		select {
		case <-ctx.Done():
			log.Info("watcher_stopped")
			return

		default:
			// pattern : {streamsDir}/{id}/{id}.m3u8
			pattern := filepath.Join(w.streamsDir, "*", "*.m3u8")
			matches, err := filepath.Glob(pattern)
			if err != nil {
				log.WithError(err).Error("glob_failed")
				continue
			}

			for _, match := range matches {
				if _, ok := seen[match]; ok {
					continue
				}

				seen[match] = struct{}{}

				dir := filepath.Dir(match)
				videoID := filepath.Base(dir)

				id, err := strconv.ParseInt(videoID, 10, 64)
				if err != nil {
					log.WithField("dir", dir).Warn("invalid_video_id")
					continue
				}

				go func(jobID int64) {
					_, err := w.app.Commands.MarDone(ctx, job.MarkDoneCommand{
						JobID: jobID,
					})
					if err != nil {
						if errors.Is(err, job.ErrNotFound) {
							log.WithField("job_id", jobID).Warn("job_not_found")
							return
						}

						log.WithError(err).WithField("job_id", jobID).Error("mark_done_failed")

						delete(seen, match)
						return
					}

					log.WithField("job_id", jobID).Info("job_complete")
				}(id)
			}

			select {
			case <-ctx.Done():
				return
			case <-time.After(2 * time.Second):
			}
		}
	}
}
