.PHONY: docs
REPO:=github.com/zlyzol/tendermint-0.32.3/libs/events

docs:
	@go get github.com/davecheney/godoc2md
	godoc2md $(REPO) > README.md

test:
	go test -v ./...
