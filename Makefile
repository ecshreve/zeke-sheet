mod-init:
	go mod init github.com/ecshreve/zeke-sheet

mod-tidy:
	go mod tidy
	
build:
	go build -o bin/zeke-sheet github.com/ecshreve/zeke-sheet/cmd/zeke-sheet

install:
	go install -i github.com/ecshreve/zeke-sheet/cmd/zeke-sheet

run-only:
	bin/zeke-sheet

run: build run-only

test:
	go test github.com/ecshreve/zeke-sheet/...

testv:
	go test -v github.com/ecshreve/zeke-sheet/...

testc:
	go test -race -coverprofile=coverage.txt -covermode=atomic github.com/ecshreve/zeke-sheet/...