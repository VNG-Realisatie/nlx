// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package grpcproxy

import "context"

type streamInfo struct {
	fullMethod           string
	organizationName     string
	publicKeyDER         string
	publicKeyFingerprint string
	peerAddr             string
}

type ctxKey int

const streamInfoKey ctxKey = 0

func setStreamInfo(ctx context.Context, info *streamInfo) context.Context {
	return context.WithValue(ctx, streamInfoKey, info)
}

func extractStreamInfo(ctx context.Context) *streamInfo {
	info, ok := ctx.Value(streamInfoKey).(*streamInfo)
	if !ok {
		return nil
	}

	return info
}
