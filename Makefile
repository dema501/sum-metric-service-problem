clean:
	@rm -rf dist

build: clean
	@mkdir dist

	GOOS=linux GOARCH=amd64 go build -o dist/main-cli src/cmd/cli/main.go
	GOOS=linux GOARCH=amd64 go build -o dist/main-rest src/cmd/rest/main.go

run-cli: clean
	go run -race src/cmd/cli/main.go

run-rest: clean
	go run -race src/cmd/rest/main.go

lint:
	golangci-lint run -v --config .golangci.yml ./...

test:
	go test ./... --cover