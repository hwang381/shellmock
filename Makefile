all: build
utest:
		go test ./...
build:
		docker build -t shellmock-dev .
run: build
		docker run -it shellmock-dev /bin/bash
test: build
		./run_tests.sh
clean:
		docker rmi shellmock-dev
build-linux:
		GOOS=linux GOARCH=amd64 go build
