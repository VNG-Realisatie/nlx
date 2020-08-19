// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package process

import (
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"runtime"
	"sync"
	"syscall"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// ShutdownFunc functions should gracefully close a process
type ShutdownFunc func() error

// Process manages closing of processes
type Process struct {
	mu                *sync.Mutex
	cleanup           []ShutdownFunc
	logger            *zap.Logger
	ShutdownComplete  chan bool
	ShutdownRequested chan bool
}

// NewProcess initializes a new process
func NewProcess(logger *zap.Logger) *Process {
	logger.Debug("setting up process")

	p := &Process{
		mu:                &sync.Mutex{},
		logger:            logger,
		cleanup:           make([]ShutdownFunc, 0),
		ShutdownComplete:  make(chan bool),
		ShutdownRequested: make(chan bool),
	}

	p.start()

	return p
}

// CloseGracefully will append ShutdownFunc and a channel for closing to
// the queue of things that should be closed before process exit.
func (p *Process) CloseGracefully(toClose ShutdownFunc) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.cleanup = append(p.cleanup, toClose)
}

func (p *Process) closeWithinTime(c1 chan string) {
	go func() {
		for i := len(p.cleanup) - 1; i >= 0; i-- {
			shutdown := p.cleanup[i]
			err := shutdown()
			// TODO: retry logic
			for err != nil {
				p.logger.Error(errors.Wrap(err, "failed to close").Error())
			}
		}
		c1 <- "done"
	}()
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

// closeLoop will close all dependencies in reversed sequential order.
func (p *Process) closeLoop() {
	p.mu.Lock()
	close(p.ShutdownRequested)

	c1 := make(chan string, 1)
	go p.closeWithinTime(c1)

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(5 * time.Second):
		p.logger.Warn(
			"timeout while trying to shutdown",
			zap.Int("shutdownlist", len(p.cleanup)))
	}

	p.syncLogger()

	close(p.ShutdownComplete)
	fmt.Fprint(os.Stdout, "shutdown complete\n")
}

// ExitGracefully start shutdown procedure manually.
func (p *Process) ExitGracefully() {
	p.closeLoop()
	os.Exit(0)
}

// Exit will start shutdown procedure on os signal.
func (p *Process) exit(s syscall.Signal) {
	p.closeLoop()
	os.Exit(128 + int(s))
}

// Start will start listening to OS signals. On SIGINT/SIGQUIT/SIGTERM it will start Exit procedure.
func (p *Process) start() {
	// Catch signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	go func() {
		for sig := range sigChan {
			logger := p.logger.With(zap.String("sig", sig.String()))
			logger.Debug("received a signal")

			switch syssig := sig.(syscall.Signal); syssig {
			case syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM:
				logger.Info("shutting down because of signal")

				p.exit(syssig)
			default:
				logger.Info("ignoring signal")
			}
		}
	}()
}

func (p *Process) syncLogger() {
	p.logger.Debug("syncing logger")
	err := p.logger.Sync()
	if err == nil {
		return
	}

	if pathErr, ok := err.(*os.PathError); ok {
		if pathErr.Op == "sync" && ignoreSyncError(pathErr.Err) {
			return
		}
	}

	fmt.Fprintf(os.Stderr, "failed to sync zap logger: %v\n", err)
}

func ignoreSyncError(err error) bool {
	switch err {
	case
		// On MacOS and on Linux when closing /dev/stdout or /dev/stderr
		syscall.ENOTTY, syscall.EINVAL:
		return true
	}

	return false
}
