//go:build h2cserver

package grpc

import (
	"context"
	"net"
	"net/http"
	"testing"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	benchmarkproto "github.com/itergia/grpc-connect-benchmark/grpc/gen/proto"
)

func BenchmarkEchoH2CServer(b *testing.B) {
	ctx := context.Background()

	var srv echoServer
	s := grpc.NewServer()
	benchmarkproto.RegisterEchoServiceServer(s, srv)

	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		b.Fatalf("Listen failed: %v", err)
	}
	defer l.Close()
	go http.Serve(l, h2c.NewHandler(s, &http2.Server{}))

	dialAndRunTest(ctx, b, l.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
}
