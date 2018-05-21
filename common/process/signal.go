package process

import (
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

func setupSignals() {
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
				// TODO(GeertJohan): #205 run through Exit() function which runs shutdown/cleanup functions (ExitFunc's) and tries to perform graceful shutdown. #205
				os.Exit(128 + int(syssig))
			default:
				logger.Info("ignoring signal")
			}
		}
	}()
}
