build:
	go build -o dist/qaper main.go

start:
	${MAKE} build
	pkill -KILL -f "dist/qaper [s]erver" || echo "kill the old process"
	./dist/qaper server -config ./example/config.toml

test:
	${MAKE} build
	./dist/qaper join -bookid=1
	./dist/qaper question
	./dist/qaper answer -body=answer1

format:
	go fmt $(go list ./... | grep -v /vendor/)

.PHONY: build
.PHONY: start
.PHONY: test
.PHONY: format
