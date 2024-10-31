build:
	go build -o bin/app ./cmd/app
	go build -o bin/tester ./cmd/tester

run: build
	./bin/app
	./bin/tester