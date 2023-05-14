TESTS = connect grpc

.PHONY: all
all: $(foreach test,$(TESTS),$(test).test $(test).test.stripped prof/$(test).test.cpuprofile prof/$(test).test.cpuprofile.svg)
	@echo
	@echo Test binary sizes:
	@stat -c '%s %n' *.test
	@echo
	@echo Stripped test binary sizes:
	@stat -c '%s %n' *.test.stripped

.PHONY: clean
clean:
	rm -fr prof $(foreach test,$(TESTS),$(test).test $(test).test.stripped)

.PHONY: %.test
%.test:
	go test -c -bench . ./$(subst .test,,$@)

prof/%.test.cpuprofile: %.test
	@mkdir -p prof
	go test -bench . -cpuprofile $@ ./$(subst .test,,$^)

%.cpuprofile.svg: %.cpuprofile
	go tool pprof -svg -output $@ -lines -sample_index=cpu $^

%.stripped: %
	strip -o $@ $^
