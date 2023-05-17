//go:generate mkdir -p gen
//go:generate sh -c "protoc --plugin=\"$(go env GOPATH)/bin/protoc-gen-go-grpc\" --plugin=\"$(go env GOPATH)/bin/protoc-gen-go\" --go_out=gen --go_opt=paths=source_relative --go-grpc_out=gen --go-grpc_opt=paths=source_relative proto/*.proto"
package grpc

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"net"
	"testing"
	"time"

	"google.golang.org/grpc"

	benchmarkproto "github.com/itergia/grpc-connect-benchmark/grpc/gen/proto"
)

func dialAndRunTest(ctx context.Context, b *testing.B, addr string, dopts ...grpc.DialOption) {
	cc, err := grpc.DialContext(ctx, addr, dopts...)
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

func newSelfSignedCertificate() (tls.Certificate, error) {
	tmpl := x509.Certificate{
		SerialNumber:          big.NewInt(1),
		BasicConstraintsValid: true,
		NotAfter:              time.Now().Add(time.Hour),
		Subject:               pkix.Name{CommonName: "test-cert"},
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses:           []net.IP{net.ParseIP("::1")},
	}

	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return tls.Certificate{}, err
	}
	rawCert, err := x509.CreateCertificate(rand.Reader, &tmpl, &tmpl, key.Public(), key)
	if err != nil {
		return tls.Certificate{}, err
	}

	return tls.Certificate{
		Certificate: [][]byte{rawCert},
		PrivateKey:  key,
	}, nil
}
