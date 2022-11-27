.PHONY: test goland

clean:
	go clean -testcache

test:
	go test ./...

goland:
	nix-shell goland.nix

