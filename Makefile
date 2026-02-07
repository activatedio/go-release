.PHONY: test goland release

VERSION ?= $(shell cat ./.version)

fmt:
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/daixiang0/gci@latest
	go fmt ./...
	goimports -w .
	gci -w .

clean:
	go clean -testcache

test:
	go test ./...

goland:
	nix-shell goland.nix

release:
	go build
	./go-release perform
	git push origin
	git push origin --tags
	GOPROXY=proxy.golang.org go list -m github.com/activatedio/go-release@$(VERSION)
