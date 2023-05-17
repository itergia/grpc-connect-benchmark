//go:build grpcwebproto

package connect

import (
	"context"
	"testing"

	"github.com/bufbuild/connect-go"
)

func BenchmarkGRPCWebProto(b *testing.B) {
	ctx := context.Background()

	runTest(ctx, b, connect.WithGRPCWeb())
}
