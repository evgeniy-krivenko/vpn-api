MAKE_PATH=$(GOPATH)/bin:/bin:/usr/bin:/usr/local/bin:$PATH

.DEFAULT_GOAL := gen


.PHONY: gen
gen:
	@protoc --go_out=. --go_grpc_out=. proto/*.proto
			@for dir in $(CURDIR)/api/*; do \
			  cd $$dir && \
			  go mod init && go mod tidy; \
			done

gen-protoc:
	@protoc --go_out=. --go_opt=paths=import \
	 --go_grpc_out=. --go_grpc_opt=paths=import \
	 proto/connection.proto


