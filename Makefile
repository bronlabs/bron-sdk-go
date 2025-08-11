.PHONY: $(MAKECMDGOALS)

test:
	go test ./...

generate-keys:
	go run cmd/keygen/main.go

generate:
	go run cmd/generator/main.go bron-open-api-public.json sdk/types sdk/api

build:
	go build ./...

publish:
	@read -p "Enter version (e.g., 0.1.18): " version; \
	echo "package version" > sdk/version/version.go; \
	echo "" >> sdk/version/version.go; \
	echo "const SDK_VERSION = \"$$version\"" >> sdk/version/version.go; \
	go build ./...; \
	git add sdk/version/version.go; \
	git commit -am "Release v$$version"; \
	git push origin master; \
	git tag v$$version; \
	git push origin v$$version