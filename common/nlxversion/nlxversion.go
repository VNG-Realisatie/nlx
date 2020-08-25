package nlxversion

import (
	"context"

	"google.golang.org/grpc/metadata"

	common_version "go.nlx.io/nlx/common/version"
)

// Version contains the version for a component
type Version struct {
	Version   string `db:"version"`
	Component string `db:"component"`
}

// NewFromGRPCContext reads the NLX version from the incomming gRPC context metadata and returns it in a struct
func NewFromGRPCContext(ctx context.Context) Version {
	v := Version{
		Version:   "unknown",
		Component: "unknown",
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		version, ok := firstString(md.Get("nlx-version"))
		if ok && len(version) > 0 {
			v.Version = version
		}

		component, ok := firstString(md.Get("nlx-component"))
		if ok && len(component) > 0 {
			v.Component = component
		}
	}

	return v
}

// NewGRPCContext returns a outgoing gRPC context with the NLX version metadata set
func NewGRPCContext(ctx context.Context, component string) context.Context {
	return metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{
		"NLX-Component": component,
		"NLX-Version":   common_version.BuildVersion,
	}))
}

func firstString(strings []string) (string, bool) {
	if len(strings) > 0 {
		return strings[0], true
	}

	return "", false
}
