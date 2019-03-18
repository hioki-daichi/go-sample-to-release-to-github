.PHONY: clean
clean:
	rm -rf dist
	rm -rf statik

.PHONY: statik
statik:
	statik -src="./public"

.PHONY: run
run:
	go run main.go

.PHONY: dry-release
dry-release:
	@make clean
	@make statik
	goreleaser release --skip-publish

.PHONY: release
release:
	@make clean
	@make statik
	goreleaser release
