//go:build grpcproto_upstream_client

package connect

import (
	"context"
	"net"
	"net/http"
	"strings"
	"testing"

	"github.com/bufbuild/connect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/itergia/grpc-connect-benchmark/connect/gen/benchmark/v1/benchmarkv1connect"
	benchmarkproto "github.com/itergia/grpc-connect-benchmark/grpc/gen/proto"
)

func BenchmarkGRPCProtoWithUpstreamClient(b *testing.B) {
	ctx := context.Background()

	var hdlr echoHandler
	_, h := benchmarkv1connect.NewEchoServiceHandler(hdlr, connect.WithCompressMinBytes(1<<10))
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		b.Fatalf("Listen failed: %v", err)
	}
	defer l.Close()
	go http.Serve(l, h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.Replace(r.URL.Path, "benchmark.", "benchmark.v1.", -1)
		h.ServeHTTP(w, r)
	}), &http2.Server{}))

	cc, err := grpc.DialContext(ctx, l.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		b.Fatalf("DialContext failed: %v", err)
	}
	defer cc.Close()
	client := benchmarkproto.NewEchoServiceClient(cc)
	req := &benchmarkproto.EchoRequest{Body: "hello world"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp, err := client.Echo(ctx, req)
		if err != nil {
			b.Fatalf("Echo failed: %v", err)
		}
		if resp.GetBody() != req.Body {
			b.Errorf("Echo didn't echo: got %q, want %q", resp.GetBody(), req.Body)
		}
	}
}
