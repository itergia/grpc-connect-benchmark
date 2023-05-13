//go:generate mkdir -p gen
//go:generate sh -c "protoc --plugin=\"$(go env GOPATH)/bin/protoc-gen-go-grpc\" --plugin=\"$(go env GOPATH)/bin/protoc-gen-go\" --go_out=gen --go_opt=paths=source_relative --go-grpc_out=gen --go-grpc_opt=paths=source_relative proto/*.proto"
package grpc

import (
	"context"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	benchmarkproto "github.com/itergia/grpc-connect-benchmark/grpc/gen/proto"
)

func BenchmarkEcho(b *testing.B) {
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

type echoServer struct {
	benchmarkproto.UnimplementedEchoServiceServer
}

func (echoServer) Echo(ctx context.Context, req *benchmarkproto.EchoRequest) (*benchmarkproto.EchoResponse, error) {
	return &benchmarkproto.EchoResponse{Body: req.GetBody()}, nil
}
