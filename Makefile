
build:
	@go build -C sim -o ../bin/8086-sim

clean:
	@go clean
	@rm -rf bin/*

SRC_DIRS ?= .
SRC_FILES := $(shell find $(SRC_DIRS) -name *.go)
format:
	@for f in $(SRC_FILES) ; do \
		go fmt $$f ;\
	done

