# drone-ansibrest

[![Go Doc](https://godoc.org/github.com/fuku2014/drone-ansibrest?status.svg)](http://godoc.org/github.com/fuku2014/drone-ansibrest)
[![Go Report](https://goreportcard.com/badge/github.com/fuku2014/drone-ansibrest)](https://goreportcard.com/report/github.com/fuku2014/drone-ansibrest)

Drone plugin to execute ansible. For the
usage information and a listing of the available options please take a look at
[the docs](DOCS.md).

## Build

Build the binary with the following commands:

```
go build
go test
```

## Docker

Build the docker image with the following commands:

```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo
docker build --rm=true -t plugins/ansibrest .
```

Please note incorrectly building the image for the correct x64 linux and with
GCO disabled will result in an error when running the Docker image:

```
docker: Error response from daemon: Container command
'/bin/drone-ansibrest' not found or does not exist..
```

## Usage

Execute from the working directory:

```
docker run --rm \
  -e PLUGIN_ENDPOINT=<endpoint> \
  -e PLUGIN_PLAYBOOK=<playbook> \
  -e PLUGIN_INVENTORY=<inventry> \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  plugins/ansibrest
```
