build:
	go build -o dist/qaper main.go

start:
	${MAKE} build
	pkill -KILL -f "dist/qaper [s]erver" || echo "kill the old process"
	./dist/qaper server

test:
	go run main.go join

format:
	go fmt $(go list ./... | grep -v /vendor/)

.PHONY: build
.PHONY: start
.PHONY: test
.PHONY: format
