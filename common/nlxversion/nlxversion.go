package nlxversion

import (
	"context"

	"go.nlx.io/nlx/common/version"

	"google.golang.org/grpc/metadata"
)

// NlxVersion contains the version for a component
type NlxVersion struct {
	Version   string
	Component string
}

// WithNlxVersionFromContext reads the NLX version headers and passes them to the closure
func WithNlxVersionFromContext(ctx context.Context, f func(nlxVersion NlxVersion)) {
	nlxVersion := NlxVersion{
		Version:   "unknown",
		Component: "unknown",
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		contextNlxVersion, ok := firstString(md.Get("nlx-version"))
		if ok && len(contextNlxVersion) > 0 {
			nlxVersion.Version = contextNlxVersion
		}

		contextNlxComponent, ok := firstString(md.Get("nlx-component"))
		if ok && len(contextNlxComponent) > 0 {
			nlxVersion.Component = contextNlxComponent
		}
	}

	f(nlxVersion)
}

// NewContext returns a context with the NLX version metadata set
func NewContext(nlxComponent string) context.Context {
	return metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{
		"NLX-Component": nlxComponent,
		"NLX-Version":   version.BuildVersion,
	}))
}

func firstString(strings []string) (string, bool) {
	if len(strings) > 0 {
		return strings[0], true
	}

	return "", false
}
