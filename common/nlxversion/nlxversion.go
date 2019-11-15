package nlxversion

import (
	"context"

	"go.nlx.io/nlx/common/version"

	"google.golang.org/grpc/metadata"
)

func WithNlxVersionFromContext(ctx context.Context, f func(nlxVersion string)) {
	nlxVersion := "unknown"
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if contextNlxVersion, ok := firstString(md.Get("nlx-version")); ok && len(contextNlxVersion) > 0 {
			nlxVersion = contextNlxVersion
		}
	}

	f(nlxVersion)
}

func NewContext(nlxComponent string) context.Context {
	return metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{
		"NLX-Component": nlxComponent,
		"NLX-Version":   version.BuildVersion,
	}))
}

func firstString(strings []string) (string, bool) {
	if len(strings) > 0 {
		return strings[0], true
	} else {
		return "", false
	}
}
