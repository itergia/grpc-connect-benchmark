//go:generate go run github.com/bufbuild/buf/cmd/buf@latest generate
package connect

import (
	"context"
	"net"
	"net/http"
	"testing"

	"github.com/bufbuild/connect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	benchmarkv1 "github.com/itergia/grpc-connect-benchmark/connect/gen/benchmark/v1"
	"github.com/itergia/grpc-connect-benchmark/connect/gen/benchmark/v1/benchmarkv1connect"
)

func runTest(ctx context.Context, b *testing.B, copts ...connect.ClientOption) {
	var hdlr echoHandler
	_, h := benchmarkv1connect.NewEchoServiceHandler(hdlr, connect.WithCompressMinBytes(1<<10))
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		b.Fatalf("Listen failed: %v", err)
	}
	defer l.Close()
	go http.Serve(l, h2c.NewHandler(h, &http2.Server{}))

	client := benchmarkv1connect.NewEchoServiceClient(http.DefaultClient, "http://"+l.Addr().String(), copts...)
	req := &benchmarkv1.EchoRequest{Body: "hello world"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp, err := client.Echo(
			ctx,
			connect.NewRequest(req),
		)
		if err != nil {
			b.Fatalf("Echo failed: %v", err)
		}
		if resp.Msg.GetBody() != req.Body {
			b.Errorf("Echo didn't echo: got %q, want %q", resp.Msg.GetBody(), req.Body)
		}
	}
}

type echoHandler struct {
	benchmarkv1connect.UnimplementedEchoServiceHandler
}

func (echoHandler) Echo(ctx context.Context, req *connect.Request[benchmarkv1.EchoRequest]) (*connect.Response[benchmarkv1.EchoResponse], error) {
	return connect.NewResponse(&benchmarkv1.EchoResponse{Body: req.Msg.GetBody()}), nil
}
