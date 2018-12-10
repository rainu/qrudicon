build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/qrudicon aws/lambda.go

.PHONY: clean
clean:
	rm -rf ./bin

.PHONY: deploy
deploy: build
	sls deploy function -f qrudicon --verbose

.PHONY: install
install: clean build
	sls deploy --verbose

.PHONY: drop
drop:
	sls remove --verbose