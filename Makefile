GOPATH=$(abspath $(shell pwd)/../..)

clean:
	rm -f pod-controller

build: clean
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o pod-controller

