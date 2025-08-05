.PHONY: build test clean generate-keys

build:
	go build -o bin/bron-sdk-go ./sdk

test:
	go test ./...

clean:
	rm -rf bin/

generate-keys:
	go run cmd/keygen.go

validate-jwk:
	go run cmd/keygen/main.go --validate $(JWK)

generate:
	go run cmd/generator/main.go bron-open-api-public.json src/types src/api

publish:
	go mod tidy
	go test ./...
	git tag v$(VERSION)
	git push origin v$(VERSION)
