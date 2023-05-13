# Connect vs gRPC Go Benchmark

A trivial echo RPC server in Go using both [Connect](https://connect.build/docs/go/getting-started/) and [gRPC](https://grpc.io/docs/languages/go/quickstart/).

```console
$ make
Test binary sizes:
14214982 connect.test
14175988 grpc.test

Stripped test binary sizes:
9765304 connect.test.stripped
9717720 grpc.test.stripped

go test -bench . ./...
goos: linux
goarch: amd64
pkg: github.com/itergia/grpc-connect-benchmark/connect
cpu: Intel(R) Core(TM) i7-6600U CPU @ 2.60GHz
BenchmarkEcho/connect-4                     5032            224454 ns/op
BenchmarkEcho/grpc-4                        4710            250859 ns/op
BenchmarkEcho/grpcWeb-4                     3321            301222 ns/op
PASS
ok      github.com/itergia/grpc-connect-benchmark/connect       5.318s
?       github.com/itergia/grpc-connect-benchmark/connect/gen/benchmark/v1      [no test files]
?       github.com/itergia/grpc-connect-benchmark/connect/gen/benchmark/v1/benchmarkv1connect   [no test files]
goos: linux
goarch: amd64
pkg: github.com/itergia/grpc-connect-benchmark/grpc
cpu: Intel(R) Core(TM) i7-6600U CPU @ 2.60GHz
BenchmarkEcho-4             8924            118307 ns/op
PASS
ok      github.com/itergia/grpc-connect-benchmark/grpc  1.081s
?       github.com/itergia/grpc-connect-benchmark/grpc/gen/proto        [no test files]
```

The variability is high (gRPC sometimes goes up to 140 µs/op,) but the relative ordering seems stable.
