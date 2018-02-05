build:
	docker build -t demo-golang-app .
	docker run --publish 6060:8000 --rm --name demo-app-run demo-golang-app
	explorer "http://localhost:6060"

test:
	go test -v ./...

clean:
	rm -rf build

.PHONY: build test lint clean
