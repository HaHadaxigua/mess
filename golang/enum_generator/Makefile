.PHONY: all fmt

GO ?= GO111MODULE=on go

fmt:
	-$(GO) fmt ./...

gen:
	-$(GO) get github.com/abice/go-enum
	-$(GO) generate ./enum/xxx.go
	-$(GO) mod tidy

