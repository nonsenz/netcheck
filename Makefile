BASEDIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
NETWORK = docker_default

build:
	docker run --rm -v "$(BASEDIR)":/usr/src/netcheck -w /usr/src/netcheck -e GOOS=linux -e CGO_ENABLED=0 golang:alpine go build -a -installsuffix cgo -o ./docker/nc nc.go
	docker build -t netcheck ./docker/
	rm -f ./docker/nc

clean:
	rm -f ./docker/nc