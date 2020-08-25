package nlxversion_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"

	"go.nlx.io/nlx/common/nlxversion"
)

func TestFromContext(t *testing.T) {
	ctx := context.Background()
	v := nlxversion.NewFromGRPCContext(ctx)

	assert.Equal(t, "unknown", v.Component)
	assert.Equal(t, "unknown", v.Version)

	ctx = metadata.NewIncomingContext(ctx, metadata.Pairs(
		"nlx-component", "test",
		"nlx-version", "1.0",
		"nlx-component", "foo",
		"nlx-version", "2.0",
	))
	v = nlxversion.NewFromGRPCContext(ctx)

	assert.Equal(t, "test", v.Component)
	assert.Equal(t, "1.0", v.Version)
}

func TestNewGRPCContext(t *testing.T) {
	type key struct{}

	ctx := context.WithValue(context.Background(), key{}, "parent")
	newCtx := nlxversion.NewGRPCContext(ctx, "test")

	md, ok := metadata.FromOutgoingContext(newCtx)
	assert.Equal(t, true, ok)
	assert.Equal(t, "test", md.Get("NLX-Component")[0])
	assert.Equal(t, "unknown", md.Get("NLX-Version")[0])

	assert.Equal(t, "parent", newCtx.Value(key{}))
}
