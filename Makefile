connect_variants = connectproto grpcproto grpcproto_upstream_client grpcwebproto
grpc_variants = grpcserver nettlsserver h2cserver

.PHONY: all
all:
	@mkdir -p build
	@set -e ; for variant in $(connect_variants); do \
		go test -o build/connect_$$variant.test -cpuprofile prof/connect_$$variant.test.cpuprofile -bench . -tags $$variant ./connect ;\
		strip -o build/connect_$$variant.test.stripped build/connect_$$variant.test ;\
	done
	@set -e ; for variant in $(grpc_variants); do \
		go test -o build/grpc_$$variant.test -cpuprofile prof/grpc_$$variant.test.cpuprofile -bench . -tags $$variant ./grpc ;\
		strip -o build/grpc_$$variant.test.stripped build/grpc_$$variant.test ;\
	done
	@echo
	@echo Test binary sizes:
	@stat -c '%s %n' build/*.test | sort -k1,1n
	@echo
	@echo Stripped test binary sizes:
	@stat -c '%s %n' build/*.test.stripped | sort -k1,1n

.PHONY: clean
clean:
	rm -fr build prof
