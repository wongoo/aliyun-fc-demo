version := v1.0.0

format:
	goimports -w -l .
	go fmt ./...
	gofumpt -w .

license-check:
	# go install github.com/lsm-dev/license-header-checker/cmd/license-header-checker@latest
	license-header-checker -v -a -r apache-license.txt . go

check: license-check
	# golangci-lint run

test:
	go test ./... -coverprofile=coverage.txt -covermode=atomic

build: format check test
	GOOS=linux arch=amd64 go build -o dist/main cmd/event_main/event_main.go
	cd dist && zip event_main.zip main && rm -f main

	GOOS=linux arch=amd64 go build -o dist/main cmd/http_main/http_main.go
	cd dist && zip http_main.zip main && rm -f main

