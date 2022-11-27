.PHONY: test goland

clean:
	go clean -testcache

test:
	go test ./...

goland:
	nix-shell goland.nix

release:
	go build
	./go-release perform
	verion=$$(cat ".version")
	GOPROXY=proxy.golang.org go list -m github.com/activatedio/go-release@$${version}
