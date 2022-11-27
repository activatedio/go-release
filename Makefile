.PHONY: test goland

clean:
	go clean -testcache

test:
	go test ./...

goland:
	nix-shell goland.nix

release:
	go get github.com/activatedio/go-release@v0.0.3
	go-release perform


