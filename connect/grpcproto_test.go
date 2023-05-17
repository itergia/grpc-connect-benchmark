//go:build grpcproto

package connect

import (
	"context"
	"testing"

	"github.com/bufbuild/connect-go"
)

func BenchmarkGRPCProto(b *testing.B) {
	ctx := context.Background()

	runTest(ctx, b, connect.WithGRPC())
}
