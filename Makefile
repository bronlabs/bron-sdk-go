.PHONY: $(MAKECMDGOALS)

test:
	go test ./...

generate-keys:
	go run cmd/keygen/main.go

generate:
	go run cmd/generator/main.go bron-open-api-public.json sdk/types sdk/api
