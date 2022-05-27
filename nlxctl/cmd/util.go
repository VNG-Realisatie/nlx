package cmd

import (
	"context"
	"strings"

	"github.com/spf13/viper"
	"google.golang.org/grpc/metadata"
)

func splitConfigString(configString string) []string {
	return strings.Split(configString, "\n---\n")
}

func newRequestContext() context.Context {
	ctx := context.Background()
	md := metadata.New(map[string]string{"Authorization": viper.GetString(accessToken)})
	ctx = metadata.NewOutgoingContext(ctx, md)

	return ctx
}
