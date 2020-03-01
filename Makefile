all: build run
utest:
		go test ./...
build:
		docker build -t shellmock-dev .
run:
		docker run -it shellmock-dev /bin/bash
test: build
		./run_tests.sh
clean:
		docker rmi shellmock-dev
