# shellmock
A shell mocking utility to easily create fake executables and assert on executable calls

**Please only run this in a virtualized testing environment because it could replace arbitrary executables**

## How to use it
```bash
$ docker run -it ubuntu /bin/bash
# then you are in a container. very safe.
$ apt update && apt install -y curl jq  # needed for github-binary-grabbing one-liner
$ curl -s https://api.github.com/repos/hwang381/shellmock/releases/latest | jq -r '.assets[].browser_download_url' | curl -L -s "$(cat /dev/stdin)" -o /bin/shellmock
$ chmod +x /bin/shellmock
$ export SHELLMOCK_ENABLE=1  # mandatory. turn off safety vault built into the program.
$ shellmock setup /usr/local/bin/mongod
$ /usr/local/bin/mongod arg1 arg2
$ /usr/local/bin/mongod arg3 arg4
$ shellmock verify /usr/local/bin/mongod arg1 arg2
$ echo $?
0  # verified first call
$ shellmock verify /usr/local/bin/mongod arg3 arg4
$ echo $?
0  # verified second call
$ shellmock verifyNoInteraction /usr/local/bin/mongod
$ echo $?
0  # verified no more calls were made
```

## Development
Prerequisites: `go 1.13` and `docker`

```bash
# unit tests
make utest

# build and run interactively in a container
make run

# build and run integration tests
make test
```
