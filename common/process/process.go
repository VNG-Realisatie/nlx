// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package process

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

// Process manages closing of processes
type Process struct {
	ctx    context.Context
	cancel context.CancelFunc
}

// NewProcess initializes a new process
func NewProcess() *Process {
	ctx, cancel := context.WithCancel(context.Background())

	p := &Process{
		ctx:    ctx,
		cancel: cancel,
	}

	p.start()

	return p
}

func (p *Process) start() {
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	go func() {
		<-termChan
		p.cancel()
	}()
}

func (p *Process) Wait() {
	<-p.ctx.Done()
}
