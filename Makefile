BIN=build
LOG=log

build:
	go build -o ${BIN}/sqncr .

run.mcp:
	go run . mcp

run.tui:
	go run . tui

test:
	go test ./...

debug.mcp:
	dlv debug . -- mcp --dev --http

debug.mcp.prod:
	dlv debug . -- mcp --http

install:
	go install .