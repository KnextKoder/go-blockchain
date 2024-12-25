build:
	go build -o ./bin/GO-BLOCKCHAIN

run:
	./bin/GO-BLOCKCHAIN

test:
	go test -v ./...