# Connect vs gRPC Go Benchmark

A trivial echo RPC server in Go using both [Connect](https://connect.build/docs/go/getting-started/) and [gRPC](https://grpc.io/docs/languages/go/quickstart/).

```console
$ make
Test binary sizes:
14214982 connect.test
14175988 grpc.test

go test -bench . ./...
goos: linux
goarch: amd64
pkg: github.com/itergia/grpc-connect-benchmark/connect
cpu: Intel(R) Core(TM) i7-6600U CPU @ 2.60GHz
BenchmarkEcho/connect-4                     5668            210448 ns/op
BenchmarkEcho/grpc-4                        4292            234059 ns/op
BenchmarkEcho/grpcWeb-4                     4264            279406 ns/op
PASS
ok      github.com/itergia/grpc-connect-benchmark/connect       5.334s
?       github.com/itergia/grpc-connect-benchmark/connect/gen/benchmark/v1      [no test files]
?       github.com/itergia/grpc-connect-benchmark/connect/gen/benchmark/v1/benchmarkv1connect   [no test files]
goos: linux
goarch: amd64
pkg: github.com/itergia/grpc-connect-benchmark/grpc
cpu: Intel(R) Core(TM) i7-6600U CPU @ 2.60GHz
BenchmarkEcho-4             9664            116506 ns/op
PASS
ok      github.com/itergia/grpc-connect-benchmark/grpc  1.150s
?       github.com/itergia/grpc-connect-benchmark/grpc/gen/proto        [no test files]
```

The variability is high (gRPC sometimes goes up to 140 µs/op,) but the relative ordering seems stable.
Stripping the binaries yields no relative difference.
