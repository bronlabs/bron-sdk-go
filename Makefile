.PHONY:

test:
	go test ./...

generate-keys:
	go run cmd/keygen.go

generate:
	go run cmd/generator.go bron-open-api-public.json sdk/types sdk/api
