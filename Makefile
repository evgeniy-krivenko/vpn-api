MAKE_PATH=$(GOPATH)/bin:/bin:/usr/bin:/usr/local/bin:$PATH

.DEFAULT_GOAL := gen


.PHONY: gen
gen:
	@protoc --go_out=./gen --go_grpc_out=./gen proto/*.proto
			@for dir in $(CURDIR)/gen/*; do \
			  cd $$dir && \
			  go mod init && go mod tidy; \
			done

gen-protoc:
	@protoc --go_out=./api --go_opt=paths=import \
	 --go_grpc_out=./api --go_grpc_opt=paths=import \
	 proto/*.proto


