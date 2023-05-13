.PHONY: all
all: connect.test grpc.test connect.test.stripped grpc.test.stripped
	@echo Test binary sizes:
	@stat -c '%s %n' *.test
	@echo
	@echo Stripped test binary sizes:
	@stat -c '%s %n' *.test.stripped
	@echo
	go test -bench . ./...

.PHONY: %.test
%.test:
	go test -c -bench . ./$(subst .test,,$@)

%.stripped: %
	strip -o $@ $^
