.PHONY: $(MAKECMDGOALS)

test:
	go test ./...

generate-keys:
	go run cmd/keygen/main.go

generate:
	go run cmd/generator/main.go bron-open-api-public.json sdk/types sdk/api

build:
	go build ./...

.PHONY: build

publish:
	go build ./...
	git add sdk/version/version.go
	git commit -am "Release v$(shell grep 'SDK_VERSION =' sdk/version/version.go | cut -d'"' -f2)"
	git push origin master
	git tag v$(shell grep 'SDK_VERSION =' sdk/version/version.go | cut -d'"' -f2)
	git push origin v$(shell grep 'SDK_VERSION =' sdk/version/version.go | cut -d'"' -f2)