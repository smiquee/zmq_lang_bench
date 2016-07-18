package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"os"
)

func main() {
	zock_in, _ := zmq.NewSocket(zmq.PAIR)
	zock_out, _ := zmq.NewSocket(zmq.PAIR)
	defer zock_in.Close()
	defer zock_out.Close()

	if len(os.Args) != 3 {
		fmt.Println("Give 2 arguments! 'url_socket_in url_socket_out'")
		os.Exit(1)
	}

	zock_in.Connect(os.Args[1])
	zock_out.Bind(os.Args[2])

	for {
		msg, _ := zock_in.Recv(0)
		zock_out.Send(msg, 0)
	}
}
