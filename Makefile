.PHONY: build clean run

build: clean
				mkdir dist
				go build -trimpath -o dist/app ./cmd/markov/main.go

clean:
			 rm -rf dist

run:
			 PRETTY=true go run --race ./cmd/markov/main.go --cache_path=.cache

test:
			 go test --race ./...

fmt: clean
			GOFUMPT_SPLIT_LONG_LINES=on gofumpt -l -w .

lint:
			golangci-lint run -E goimports -E gofumpt
