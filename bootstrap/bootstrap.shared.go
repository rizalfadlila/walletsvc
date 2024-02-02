package bootstrap

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (b *bootstrap) errorHandler(errChan <-chan error) {
	select {
	case s := <-b.signals():
		log.Info().Interface("signal", s)
	case err := <-errChan:
		log.Err(err).Msg("boostrap received error")
	}

	b.gracefullyExit()
}

func (b *bootstrap) signals() <-chan os.Signal {
	done := make(chan os.Signal, 2)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)
	signal.Notify(done, os.Interrupt, syscall.SIGINT)
	return done
}

func (b *bootstrap) gracefullyExit() {
	ctx, cancel := context.WithTimeout(context.Background(), b.getTimeDurationTimeout())
	defer cancel()

	done := make(chan bool)

	go func() {
		b.deps.Close(ctx)
		done <- true
	}()

	select {
	case <-ctx.Done():
		log.Err(errors.New("timeout waiting all processes to stop")).Msg("error gracefully exit")
	case <-done:
		log.Info().Msg("gracefully exit 👋🏻")
	}
}

func (b *bootstrap) getTimeDurationTimeout() time.Duration {
	timeout := b.deps.Config.Server.GracefulTimeout
	timeDuration, err := time.ParseDuration(timeout)
	if err != nil {
		timeDuration, _ = time.ParseDuration(defaultGracefulTimeout)
	}

	return timeDuration
}
