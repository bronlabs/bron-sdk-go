.PHONY: $(MAKECMDGOALS)

VERSION ?= $(shell git describe --tags --abbrev=0 | sed 's/v//' | awk -F. '{$$NF = $$NF + 1;} 1' | sed 's/ /./g')

test:
	go test ./...

generate-keys:
	go run cmd/keygen/main.go

generate:
	go run cmd/generator/main.go bron-open-api-public.json sdk/types sdk/api

build:
	go build ./...

publish:
	echo "package http\n\nconst SDK_VERSION = \"${VERSION}\"" > sdk/http/version.go;
	@make build
	git add sdk/http/version.go;
	git commit -am "Release v${VERSION}";
	git push origin master;
	git tag -a v${VERSION} -m "Release v${VERSION}";
	git push origin v${VERSION};
