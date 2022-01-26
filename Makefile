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

package-event:
	rm -f dist/event_main.zip
	GOOS=linux arch=amd64 go build -o dist/main cmd/event_main/event_main.go
	cd dist && zip event_main.zip main && rm -f main

package-http:
	rm -f dist/http_main.zip
	GOOS=linux arch=amd64 go build -o dist/main cmd/http_main/http_main.go
	cd dist && zip http_main.zip main && rm -f main

package-eventbridge:
	rm -f dist/eventbridge_main.zip
	GOOS=linux arch=amd64 go build -o dist/main cmd/eventbridge_main/eventbridge_main.go
	cd dist && zip eventbridge_main.zip main && rm -f main
