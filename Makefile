TEST ?= ./...

GOCMD=$(if $(shell which richgo),richgo,go)


build:
	$(GOCMD) build -o dist/server ./cmd/server/main.go

build-ci:
	mkdir -p ./dist
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOCMD) build -o dist/server ./cmd/server/main.go

test:
	ENV=test $(GOCMD) test -v $(TEST)

test-watch:
	reflex -s --decoration=none -r \.go$$ -- make test TEST=$(TEST)
	ENV=test $(GOCMD) test -v $(TEST)

dev:
	reflex -s --decoration=none -r '\.(go|html)$$' -- $(GOCMD) run ./cmd/server/main.go

release:
	@bash -c "$$(curl -s https://raw.githubusercontent.com/escaletech/releaser/master/tag-and-push.sh)"
