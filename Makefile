build:
	go build -o ./bin/GO-BLOCKCHAIN

run: build
	./bin/GO-BLOCKCHAIN

test:
	go test ./...