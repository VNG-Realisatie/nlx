package nlxversion

import (
	"context"

	"google.golang.org/grpc/metadata"

	"go.nlx.io/nlx/common/version"
)

// NlxVersion contains the version for a component
type NlxVersion struct {
	Version   string `db:"version"`
	Component string `db:"component"`
}

// GetNlxVersionFromContext reads the NLX version from the context metadata and returns it in a struct
func GetNlxVersionFromContext(ctx context.Context) NlxVersion {
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

	return nlxVersion
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
