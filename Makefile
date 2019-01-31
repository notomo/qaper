build:
	go build -o dist/qaper main.go

start:
	${MAKE} build
	pkill -KILL -f "dist/qaper [s]erver" || echo "kill the old process"
	./dist/qaper server -config ./example/config.toml

test:
	go run main.go join -bookid=1
	go run main.go question
	go run main.go answer -body=answer1

format:
	go fmt $(go list ./... | grep -v /vendor/)

.PHONY: build
.PHONY: start
.PHONY: test
.PHONY: format
