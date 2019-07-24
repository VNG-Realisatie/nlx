// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package process

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// ShutdownFunc functions should gracefully close a process
type ShutdownFunc func() error

type shutdown struct {
	shutDownFunc  ShutdownFunc
	closeComplete chan bool
}

// Process manages closing of processes
type Process struct {
	mu                *sync.Mutex
	close             []ShutdownFunc
	logger            *zap.Logger
	ShutdownComplete  chan bool
	ShutdownRequested chan bool
}

// NewProcess initializes a new process
func NewProcess(logger *zap.Logger) *Process {
	logger.Debug("setting up process")
	// Currently we develop for linux, binaries might be compiled for other platforms
	// but we should warn the user about possible unexpected behavior
	if runtime.GOOS != "linux" {
		logger.Warn("detected non-linux OS, program might behave unexpected", zap.String("os", runtime.GOOS))
	}
	p := &Process{
		mu:                &sync.Mutex{},
		logger:            logger,
		close:             make([]ShutdownFunc, 0),
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
	p.close = append(p.close, toClose)
}

// closeLoop will close all dependencies in reversed sequential order.
func (p *Process) closeLoop() {
	p.mu.Lock()
	close(p.ShutdownRequested)
	for i := len(p.close) - 1; i >= 0; i-- {
		err := p.close[i]()
		// TODO: retry logic
		for err != nil {
			p.logger.Error(errors.Wrap(err, "failed to close").Error())
		}
	}

	err := p.logger.Sync()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to sync zap logger: %v\n", err)
	}
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
