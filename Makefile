build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/simple aws/simple.go

.PHONY: clean
clean:
	rm -rf ./bin

.PHONY: deploy
deploy: build
	sls deploy function -f simple --verbose

.PHONY: install
install: clean build
	sls deploy --verbose

.PHONY: drop
drop:
	sls remove --verbose