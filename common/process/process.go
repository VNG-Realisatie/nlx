package process

import (
	"io"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const maxLevel = 100 // Think even 10 would be enough. Too big number will increase chances of miss use

type process struct {
	mu    *sync.Mutex
	close [maxLevel][]io.Closer
}

// CloseGracefully will append io.Closer interface to the queue of things that should be closed before process exit.
// Entities with lower level will be closed first.
// Entities on higher level (bigger number) will be closed only in case if all lower levels were closed.
// Maximum possible level is 100
func (p *process) CloseGracefully(toClose io.Closer, level uint) error {
	if level > maxLevel {
		return ErrMaxLevelReached
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	p.close[level] = append(p.close[level], toClose)
	return nil
}

// Exit will do shutdown procedure. It will close all dependencies in sequential order, level by level.
func (p *process) Exit(s syscall.Signal) {
	p.mu.Lock()

	// Iterating level by level
	for _, cc := range p.close {
		// Iterating and closing each dependency.
		// TODO: could be improved to run concurrently
		for _, c := range cc {
			err := c.Close()
			if err != nil {
				// TODO: Retry logic
				logger.Error(errors.Wrap(err, "failed to close").Error())
			}
		}
	}

	p.mu.Unlock()
	os.Exit(128 + int(s))
}

// Start will start listening to OS signals. On SIGINT/SIGQUIT/SIGTERM it will start Exit procedure.
func (p *process) Start() {
	// Catch signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	go func() {
		for sig := range sigChan {
			logger = logger.With(zap.String("sig", sig.String())) // potential data race. (updating global var)
			logger.Debug("received a signal")

			switch syssig := sig.(syscall.Signal); syssig {
			case syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM:
				logger.Info("shutting down because of signal")

				p.Exit(syssig)
			default:
				logger.Info("ignoring signal")
			}
		}
	}()
}
