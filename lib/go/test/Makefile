.PHONY: test
test:
	CGO_ENABLED=0 go test -tags=no_cgo ./...

.PHONY: ci
ci: test
