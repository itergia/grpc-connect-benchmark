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

func BenchmarkEcho(b *testing.B) {
	ctx := context.Background()

	var hdlr echoHandler
	mux := http.NewServeMux()
	mux.Handle(benchmarkv1connect.NewEchoServiceHandler(hdlr))
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		b.Fatalf("Listen failed: %v", err)
	}
	defer l.Close()
	go http.Serve(l, h2c.NewHandler(mux, &http2.Server{}))

	tsts := []struct {
		Name       string
		ClientOpts []connect.ClientOption
	}{
		{"connect", nil},
		{"grpc", []connect.ClientOption{connect.WithGRPC()}},
		{"grpcWeb", []connect.ClientOption{connect.WithGRPCWeb()}},
	}

	for _, tst := range tsts {
		b.Run(tst.Name, func(b *testing.B) {
			client := benchmarkv1connect.NewEchoServiceClient(http.DefaultClient, "http://"+l.Addr().String(), tst.ClientOpts...)
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
		})
	}
}

type echoHandler struct {
	benchmarkv1connect.UnimplementedEchoServiceHandler
}

func (echoHandler) Echo(ctx context.Context, req *connect.Request[benchmarkv1.EchoRequest]) (*connect.Response[benchmarkv1.EchoResponse], error) {
	return connect.NewResponse(&benchmarkv1.EchoResponse{Body: req.Msg.GetBody()}), nil
}
