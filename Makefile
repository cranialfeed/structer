SHELL := /bin/bash -eu -o pipefail

GO_TEST_FLAGS  := -v

PACKAGE_DIRS := $(shell go list ./... 2> /dev/null | grep -v /vendor/)
SRC_FILES    := $(shell find . -name '*.go' -not -path './vendor/*')

