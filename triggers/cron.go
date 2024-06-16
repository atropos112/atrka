// Package triggers provides variety of triggers to initiate the process of sending a message for processing.
package triggers

import (
	"context"
	"errors"
	"time"

	"github.com/adhocore/gronx"
	"github.com/atropos112/atrka/types"
	"github.com/go-co-op/gocron"
)

// CronHandler is a handler that is based on a cron expression
// expression starts with minute e.g. */1 * * * * * - every second
func CronHandler(expression string) func(ctx context.Context, ch chan<- types.Message) error {
	return func(_ context.Context, ch chan<- types.Message) error {
		// Validate the cron expression
		gron := gronx.New()
		if !gron.IsValid(expression) {
			return errors.New("invalid cron expression")
		}

		// Start the cron job
		s := gocron.NewScheduler(time.UTC)
		_, err := s.CronWithSeconds(expression).Do(func() {
			msg := types.CronMessage{Timestamp: time.Now().Unix()}
			ch <- msg
		})

		s.StartBlocking()
		return err
	}
}

// MakeCronTrigger is a helpful function to create a cron trigger using the cron handler
func MakeCronTrigger[P types.Producer](producer P, expression string) types.Trigger[P] {
	return types.NewTrigger(CronHandler(expression), producer)
}
