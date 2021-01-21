.PHONY: build

# This Makefile is a simple example that demonstrates usual steps to build a binary that can be run in the same
# architecture that was compiled in. The "ldflags" in the build assure that any needed dependency is included in the
# binary and no external dependencies are needed to run the service.

GOSAMPLEAPI_VERSION=$(shell git describe --always --long --dirty --tags)
BIN_NAME=app

build:
	go build -a -ldflags="[pattern=]args list" -o ${BIN_NAME}
	@echo "You can now use ./${BIN_NAME}"