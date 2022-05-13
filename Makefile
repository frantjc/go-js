GO ?= go

fmt test:
	@$(GO) $@ ./...
