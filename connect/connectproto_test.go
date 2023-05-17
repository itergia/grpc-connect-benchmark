//go:build connectproto

package connect

import (
	"context"
	"testing"
)

func BenchmarkConnectProto(b *testing.B) {
	ctx := context.Background()

	runTest(ctx, b)
}
