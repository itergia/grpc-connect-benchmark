//go:build grpcserver

package grpc

import (
	"context"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	benchmarkproto "github.com/itergia/grpc-connect-benchmark/grpc/gen/proto"
)

func BenchmarkEchoGRPC(b *testing.B) {
	ctx := context.Background()

	var srv echoServer
	s := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	benchmarkproto.RegisterEchoServiceServer(s, srv)

	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		b.Fatalf("Listen failed: %v", err)
	}
	defer l.Close()
	go s.Serve(l)

	dialAndRunTest(ctx, b, l.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
}
