.PHONY: test goland release

VERSION ?= $(shell cat ./.version)

clean:
	go clean -testcache

test:
	go test ./...

goland:
	nix-shell goland.nix

release:
	go build
	./go-release perform

refreshsum:
	GOPROXY=proxy.golang.org go list -m github.com/activatedio/go-release@$(VERSION)
