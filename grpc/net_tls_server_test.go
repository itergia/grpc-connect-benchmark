//go:build nettlsserver

package grpc

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	benchmarkproto "github.com/itergia/grpc-connect-benchmark/grpc/gen/proto"
)

func BenchmarkEchoNetTLSServer(b *testing.B) {
	ctx := context.Background()

	var srv echoServer
	s := grpc.NewServer()
	benchmarkproto.RegisterEchoServiceServer(s, srv)
	scert, err := newSelfSignedCertificate()
	if err != nil {
		b.Fatalf("newSelfSignedCertificate failed: %v", err)
	}
	hs := http.Server{
		TLSConfig: &tls.Config{
			Certificates:       []tls.Certificate{scert},
			InsecureSkipVerify: true,
		},
		Handler: s,
	}

	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		b.Fatalf("Listen failed: %v", err)
	}
	defer l.Close()
	go hs.ServeTLS(l, "", "")

	dialAndRunTest(ctx, b, l.Addr().String(), grpc.WithTransportCredentials(credentials.NewTLS(hs.TLSConfig)), grpc.WithBlock())
}
