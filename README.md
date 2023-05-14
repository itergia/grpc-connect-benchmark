# Connect vs gRPC Go Benchmark

A trivial echo RPC server in Go using both [Connect](https://connect.build/docs/go/getting-started/) and [gRPC](https://grpc.io/docs/languages/go/quickstart/).

```console
$ make
strip -o connect.test.stripped connect.test                                                                                                                                                                 [89/816]
strip -o grpc.test.stripped grpc.test
go test -bench . -cpuprofile connect.test.cpuprofile ./connect
goos: linux
goarch: amd64
pkg: github.com/itergia/grpc-connect-benchmark/connect
cpu: Intel(R) Core(TM) i7-6600U CPU @ 2.60GHz
BenchmarkEcho/connect-4                     7875            160838 ns/op
BenchmarkEcho/grpc-4                        5397            196728 ns/op
BenchmarkEcho/grpcWeb-4                     5672            201658 ns/op
PASS
ok      github.com/itergia/grpc-connect-benchmark/connect       4.639s
go tool pprof -svg -output connect.test.cpuprofile.svg -lines -sample_index=cpu connect.test.cpuprofile
Generating report in connect.test.cpuprofile.svg
go test -bench . -cpuprofile grpc.test.cpuprofile ./grpc
goos: linux
goarch: amd64
pkg: github.com/itergia/grpc-connect-benchmark/grpc
cpu: Intel(R) Core(TM) i7-6600U CPU @ 2.60GHz
BenchmarkEcho-4             9835            127615 ns/op
PASS
ok      github.com/itergia/grpc-connect-benchmark/grpc  1.414s
go tool pprof -svg -output grpc.test.cpuprofile.svg -lines -sample_index=cpu grpc.test.cpuprofile
Generating report in grpc.test.cpuprofile.svg

Test binary sizes:
14218250 connect.test
14175988 grpc.test

Stripped test binary sizes:
9765304 connect.test.stripped
9717720 grpc.test.stripped
```

The variability is high (gRPC sometimes goes up to 140 µs/op,) but the relative ordering seems stable.

## Notes

* The connect benchmark runs all three implementations, so the binary contains three compatibility layers.
  Adding gRPCWeb-Gateway or EnvoyProxy to translate gRPC-Web to gRPC would be a closer comparison in terms of functionality.
