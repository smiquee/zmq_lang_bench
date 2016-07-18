all: identity_c identity_go

godeps:
	GOPATH=${PWD} go get github.com/pebbe/zmq4
	GOPATH=${PWD} go install github.com/pebbe/zmq4

identity_go:
	GOPATH=${PWD} go install identity_go

identity_c:
	mkdir -p bin
	cc identity_c.c -o bin/identity_c -lzmq
