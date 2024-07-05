.PHONY: all test build run generate clean

test:
	go test -v ./...

build:
	go build -o app .

run: build
	./app

fuzz: 
	go test -fuzz=Fuzz -fuzztime=60s -parallel=4


generate: 
	@node generator

clean:
	@echo "Cleaning up..."
	rm -f app
	rm -f testdata/*.csv
	cd generator && rm -rf node_modules