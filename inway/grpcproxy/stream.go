// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package grpcproxy

import (
	"context"
	"errors"
	"io"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var clientStreamDesc = &grpc.StreamDesc{
	ServerStreams: true,
	ClientStreams: true,
}

// wrappedServerStream wraps a grpc.ServerStream with a context
type wrappedServerStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (w *wrappedServerStream) Context() context.Context {
	return w.ctx
}

func (p *Proxy) proxyStreamToUpstream(serverStream grpc.ServerStream, clientStream grpc.ClientStream) error {
	serverErrChan := make(chan error, 1)
	clientErrChan := make(chan error, 1)

	go handleServerStream(serverErrChan, serverStream, clientStream)
	go handleClientStream(clientErrChan, clientStream, serverStream)

	for i := 0; i < 2; i++ {
		select {
		case err := <-serverErrChan:
			// Our client has nothing more to send, close the stream to the upstream
			if errors.Is(err, io.EOF) {
				if err = clientStream.CloseSend(); err != nil {
					p.logger.Error("closing client stream", zap.Error(err))
				}

				break
			}

			p.logger.Error("server stream", zap.Error(err))

			return status.Errorf(codes.Internal, "")
		case err := <-clientErrChan:
			// Our upstream has nothing more to send or returned an error
			serverStream.SetTrailer(clientStream.Trailer())

			if errors.Is(err, io.EOF) {
				return nil
			}

			p.logger.Debug("client stream", zap.Error(err))

			return err
		}
	}

	p.logger.Error("unknow error in handler")

	return status.Errorf(codes.Internal, "unknow error")
}

func handleServerStream(ch chan<- error, server grpc.ServerStream, client grpc.ClientStream) {
	msg := &message{}

	for {
		if err := server.RecvMsg(msg); err != nil {
			ch <- err
			break
		}

		if err := client.SendMsg(msg); err != nil {
			ch <- err
			break
		}
	}
}

func handleClientStream(ch chan<- error, client grpc.ClientStream, server grpc.ServerStream) {
	msg := &message{}
	initial := true

	for {
		if err := client.RecvMsg(msg); err != nil {
			ch <- err
			break
		}

		// Send headers on initial message
		if initial {
			initial = false

			md, err := client.Header()
			if err != nil {
				ch <- err
				break
			}

			if err := server.SendHeader(md); err != nil {
				ch <- err
				break
			}
		}

		if err := server.SendMsg(msg); err != nil {
			ch <- err
			break
		}
	}
}
