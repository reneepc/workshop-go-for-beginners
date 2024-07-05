.PHONY: all test build run generate clean

test:
	go test -v ./...

build:
	go build -o app .

run: build
	./app

generate: 
	@node generator

clean:
	@echo "Cleaning up..."
	rm -f app
	rm -f testdata/*.csv
	cd generator && rm -rf node_modules