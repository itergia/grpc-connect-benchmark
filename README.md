# Connect vs gRPC Go Benchmark

A trivial echo RPC server in Go using both [Connect](https://connect.build/docs/go/getting-started/) and [gRPC](https://grpc.io/docs/languages/go/quickstart/).

```console
$ make
goos: linux
goarch: amd64
pkg: github.com/itergia/grpc-connect-benchmark/connect
cpu: Intel(R) Core(TM) i7-6600U CPU @ 2.60GHz
BenchmarkConnectProto-4             8238            139915 ns/op
PASS
ok      github.com/itergia/grpc-connect-benchmark/connect       2.221s
goos: linux
goarch: amd64
pkg: github.com/itergia/grpc-connect-benchmark/connect
cpu: Intel(R) Core(TM) i7-6600U CPU @ 2.60GHz
BenchmarkGRPCProto-4        6830            169507 ns/op
PASS
ok      github.com/itergia/grpc-connect-benchmark/connect       2.322s
goos: linux
goarch: amd64
pkg: github.com/itergia/grpc-connect-benchmark/connect
cpu: Intel(R) Core(TM) i7-6600U CPU @ 2.60GHz
BenchmarkGRPCWebProto-4             6403            189654 ns/op
PASS
ok      github.com/itergia/grpc-connect-benchmark/connect       2.427s
goos: linux
goarch: amd64
pkg: github.com/itergia/grpc-connect-benchmark/grpc
cpu: Intel(R) Core(TM) i7-6600U CPU @ 2.60GHz
BenchmarkEchoGRPC-4         8624            119151 ns/op
PASS
ok      github.com/itergia/grpc-connect-benchmark/grpc  1.220s
goos: linux
goarch: amd64
pkg: github.com/itergia/grpc-connect-benchmark/grpc
cpu: Intel(R) Core(TM) i7-6600U CPU @ 2.60GHz
BenchmarkEchoNetTLSServer-4         4928            234127 ns/op
PASS
ok      github.com/itergia/grpc-connect-benchmark/grpc  2.327s
goos: linux
goarch: amd64
pkg: github.com/itergia/grpc-connect-benchmark/grpc
cpu: Intel(R) Core(TM) i7-6600U CPU @ 2.60GHz
BenchmarkEchoH2CServer-4            5403            216930 ns/op
PASS
ok      github.com/itergia/grpc-connect-benchmark/grpc  2.322s

Test binary sizes:
14175411 build/grpc_grpcserver.test
14211448 build/connect_connectproto.test
14215711 build/connect_grpcproto.test
14215722 build/connect_grpcwebproto.test
15927035 build/grpc_nettlsserver.test
16150077 build/grpc_h2cserver.test

Stripped test binary sizes:
9717784 build/grpc_grpcserver.test.stripped
9765304 build/connect_connectproto.test.stripped
9765304 build/connect_grpcproto.test.stripped
9765304 build/connect_grpcwebproto.test.stripped
10938936 build/grpc_nettlsserver.test.stripped
11083064 build/grpc_h2cserver.test.stripped
```

## Notes

* Using `-cpuprofile` does nothing for the test file size.
* Stripping doesn't change the relative order.
