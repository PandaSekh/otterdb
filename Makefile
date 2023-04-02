build:
	go build -o bin/main main.go

compile:
	GOOS=linux GOARCH=amd64 go build -o bin/main-linux-amd64 main.go
	GOOS=windows GOARCH=amd64 go build -o bin/main-windows-amd64 main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/app-amd64-darwin app.go
	GOOS=darwin GOARCH=arm64 go build -o bin/app-arm64-darwin app.go

clean:
	rm -rf bin

run:
	go run .

coverage:
	go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out

benchmark:
	go test ./... -bench=.