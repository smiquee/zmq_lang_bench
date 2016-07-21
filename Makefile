all: identity_c identity_go sender_go receiver_go

godeps:
	GOPATH=${PWD} go get github.com/pebbe/zmq4
	GOPATH=${PWD} go install github.com/pebbe/zmq4

identity_go:
	GOPATH=${PWD} go install identity_go

sender_go:
	GOPATH=${PWD} go install sender_go

receiver_go:
	GOPATH=${PWD} go install receiver_go

identity_c:
	mkdir -p bin
	cc identity_c.c -o bin/identity_c -lzmq

new_zmq:
	GOPATH=${PWD} go install new_zmq
