.PHONY:

test:
	go test ./...

generate-keys:
	go run cmd/keygen/main.go

generate:
	go run scripts/update-version.go && go run cmd/generator/main.go bron-open-api-public.json sdk/types sdk/api

.PHONY: generate

update-version:
	go run scripts/update-version.go

.PHONY: update-version

build:
	go run scripts/update-version.go && go build ./...

.PHONY: build

publish:
	go run scripts/update-version.go
	go build ./...
	git add VERSION sdk/version/version.go
	git commit -am "Release v$(shell cat VERSION)"
	git push origin master
	git tag v$(shell cat VERSION)
	git push origin v$(shell cat VERSION)

.PHONY: publish
