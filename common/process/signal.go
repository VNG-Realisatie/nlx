package process

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

func setupSignals(ctx context.Context) context.Context {
	// Creating new nested context with CancelFunc to call it on signal
	// All lower level contexts should be inherited from this one to receive cancel command on signal
	newCtx, cancel := context.WithCancel(ctx)

	// Catch signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)
	signal.Notify(sigChan, syscall.SIGQUIT)
	signal.Notify(sigChan, syscall.SIGTERM)
	go func() {
		for sig := range sigChan {
			logger = logger.With(zap.String("sig", sig.String()))
			logger.Debug("received a signal")
			switch syssig := sig.(syscall.Signal); syssig {
			case syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM:
				logger.Info("shutting down because of signal")
				cancel()
			default:
				logger.Info("ignoring signal")
			}
		}
	}()
	return newCtx
}
