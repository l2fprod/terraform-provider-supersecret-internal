default: build

build:
	go build ./...

test:
	go test ./...

generate:
	go generate ./...
	go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-name supersecret --rendered-provider-name supersecret

lint:
	golangci-lint run ./...

snapshot:
	GPG_FINGERPRINT='' goreleaser release --snapshot --clean --skip=sign --release-notes release-notes/v0.1.0.md

.PHONY: build test generate lint snapshot
