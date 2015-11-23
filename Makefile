run:
	# Exmaple how to run this.
	go run main.go config.go

generate:
	# specify the name to ensure we regenerate correctly.
	swagger generate client -f swagger.yml -A stash_core_rest_api

build:
	go build ./...


