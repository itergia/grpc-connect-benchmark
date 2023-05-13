.PHONY: all
all: connect.test grpc.test
	@echo Test binary sizes:
	@stat -c '%s %n' *.test
	@echo
	go test -bench . ./...

%.test: %/%_test.go
	go test -c -bench . ./$<
